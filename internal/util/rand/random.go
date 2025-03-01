//
// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package rand contains utility functions largely for unit testing.
// WARNING: Do not use the functions in this package that generate rand / seeds
// for any security related purposes, outside of testing.
package rand

import (
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"sync/atomic"
)

// seedGuard makes sure each call to NewRand is initialized with a unique seed.
var seedGuard atomic.Int64

// NewRand returns a new instance of rand.Rand with a fixed source.
func NewRand(seed int64) *rand.Rand {
	seedGuard.Add(1)
	return rand.New(rand.NewSource(seed + seedGuard.Load()))
}

// RandomInt returns a random integer between min and max.
func RandomInt(min, max int64, seed int64) int64 {
	r := NewRand(seed)
	return min + r.Int63n(max-min+1)
}

// RandomString returns a random string of length n.
func RandomString(n int, seed int64) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := make([]byte, n)
	r := NewRand(seed)
	for i := range s {
		s[i] = letters[r.Intn(len(letters))]
	}
	return string(s)
}

// RandomEmail returns a random email address.
func RandomEmail(seed int64) string {
	return RandomString(10, seed) + "@example.com"
}

// RandomName returns a random name.
func RandomName(seed int64) string {
	return RandomString(10, seed)
}

// RandomURL returns a random URL.
func RandomURL(seed int64) string {
	return "http://" + RandomString(10, seed) + ".com"
}

// RandomPassword returns a random password.
func RandomPassword(length int, seed int64) string {
	// Define character pools
	upperChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars := "abcdefghijklmnopqrstuvwxyz"
	numberChars := "0123456789"
	specialChars := "_.;?&@"

	r := NewRand(seed)

	// Create a slice to hold the password characters
	password := make([]byte, length)

	// Determine the number of characters needed from each pool
	numUpper := 1
	numLower := 1
	numNumber := 1
	numSpecial := 1
	numRemaining := length - numUpper - numLower - numNumber - numSpecial

	// Fill the password with random characters
	fillPasswordChars(r, password, upperChars, numUpper)
	fillPasswordChars(r, password, lowerChars, numLower)
	fillPasswordChars(r, password, numberChars, numNumber)
	fillPasswordChars(r, password, specialChars, numSpecial)
	fillPasswordChars(r, password, upperChars+lowerChars+numberChars+specialChars, numRemaining)

	// Shuffle the password characters
	r.Shuffle(length, func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

func fillPasswordChars(r *rand.Rand, password []byte, charset string, count int) {
	for i := 0; i < count; i++ {
		randomIndex := r.Intn(len(password))
		for password[randomIndex] != 0 {
			randomIndex = r.Intn(len(password))
		}
		password[randomIndex] = getRandomChar(r, charset)
	}
}

func getRandomChar(r *rand.Rand, charset string) byte {
	return charset[r.Intn(len(charset))]
}

// RandomKeypair returns a random RSA keypair
func RandomKeypair(length int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(crand.Reader, length)
	if err != nil {
		return nil, nil
	}
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey
}

// RandomKeypairFile generates a random RSA keypair and writes it to files
func RandomKeypairFile(length int, privateFilePath string, publicFilePath string) error {
	filePriv, err := os.Create(filepath.Clean(privateFilePath))
	if err != nil {
		return err
	}
	defer filePriv.Close()

	filePub, err := os.Create(filepath.Clean(publicFilePath))
	if err != nil {
		return err
	}
	defer filePub.Close()

	privateKey, err := rsa.GenerateKey(crand.Reader, length)
	if err != nil {
		return err
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	err = pem.Encode(filePriv, pemBlock)
	if err != nil {
		return err
	}

	publicKey := &privateKey.PublicKey
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	err = pem.Encode(filePub, publicKeyPEM)
	if err != nil {
		return err
	}
	return nil
}

// RandomPrivateKeyFile generates a random RSA private key and writes it to a file
func RandomPrivateKeyFile(length int, filePath string) error {
	privateKey, err := rsa.GenerateKey(crand.Reader, length)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer file.Close()

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	err = pem.Encode(file, pemBlock)
	if err != nil {
		return err
	}

	return nil
}

// GetRandomPort returns a random port number.
// The binding address should not need to be configurable
// as this is a short-lived operation just to discover a random available port.
// Note that there is a possible race condition here if another process binds
// to the same port between the time we discover it and the time we use it.
// This is unlikely to happen in practice, but if it does, the user will
// need to retry the command.
// Marking a nosec here because we want this to listen on all addresses to
// ensure a reliable connection chance for the client. This is based on lessons
// learned from the sigstore CLI.
func GetRandomPort() (int, error) {
	listener, err := net.Listen("tcp", ":0") // #nosec
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	port := listener.Addr().(*net.TCPAddr).Port
	return port, nil
}
