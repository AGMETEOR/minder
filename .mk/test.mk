#
# Copyright 2023 Stacklok, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: clean
clean: ## clean up environment
	rm -rf dist/* & rm -rf bin/*

.PHONY: test
test: clean init-examples ## run tests in verbose mode
	go test -json -race -v ./... | gotestfmt

.PHONY: test-silent
test-silent: clean init-examples ## run tests in a silent mode (errors only output)
	go test -json -race -v ./... | gotestfmt -hide "all"

.PHONY: cover
cover: init-examples ## display test coverage
	go test -v -coverprofile=coverage.out -race ./...
	go tool cover -func=coverage.out

.PHONY: lint
lint: ## lint go files
	golangci-lint run
