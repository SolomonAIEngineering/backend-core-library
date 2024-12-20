# Copyright The Solomon AI Authors
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

TOOLS_MOD_DIR := ./internal/tools

ALL_DOCS := $(shell find . -name '*.md' -type f | sort)
ALL_GO_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | sort)
OTEL_GO_MOD_DIRS := $(filter-out $(TOOLS_MOD_DIR), $(ALL_GO_MOD_DIRS))
ALL_COVERAGE_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | grep -E -v '^./example|^$(TOOLS_MOD_DIR)' | sort)

GO = go
TIMEOUT = 60

.DEFAULT_GOAL := precommit

.PHONY: precommit ci
precommit: fmt test vanity-import-fix misspell go-mod-tidy test-default
ci: vanity-import-check build test-default check-clean-work-tree test-coverage
VERSION:=$(shell grep 'VERSION' version.go | awk '{ print $$4 }' | tr -d '"')
PATCH_VERSION?=$(shell grep 'VERSION' version.go | awk -F'[".]' '{ printf "%s.%s.%d", $$2, $$3, $$4+1 }')
MINOR_VERSION?=$(shell grep 'VERSION' version.go | awk -F'[".]' '{ printf "%s.%d.0", $$2, $$3+1 }')
MAJOR_VERSION?=$(shell grep 'VERSION' version.go | awk -F'[".]' '{ printf "%d.0.0", $$2+1 }')
# Tools

TOOLS = $(CURDIR)/.tools

$(TOOLS):
	@mkdir -p $@
$(TOOLS)/%: | $(TOOLS)
	cd $(TOOLS_MOD_DIR) && \
	$(GO) build -o $@ $(PACKAGE)

MULTIMOD = $(TOOLS)/multimod
$(TOOLS)/multimod: PACKAGE=go.opentelemetry.io/build-tools/multimod

SEMCONVGEN = $(TOOLS)/semconvgen
$(TOOLS)/semconvgen: PACKAGE=go.opentelemetry.io/build-tools/semconvgen

CROSSLINK = $(TOOLS)/crosslink
$(TOOLS)/crosslink: PACKAGE=go.opentelemetry.io/build-tools/crosslink

SEMCONVKIT = $(TOOLS)/semconvkit
$(TOOLS)/semconvkit: PACKAGE=go.opentelemetry.io/otel/$(TOOLS_MOD_DIR)/semconvkit

DBOTCONF = $(TOOLS)/dbotconf
$(TOOLS)/dbotconf: PACKAGE=go.opentelemetry.io/build-tools/dbotconf

GOLANGCI_LINT = $(TOOLS)/golangci-lint
$(TOOLS)/golangci-lint: PACKAGE=github.com/golangci/golangci-lint/cmd/golangci-lint

MISSPELL = $(TOOLS)/misspell
$(TOOLS)/misspell: PACKAGE=github.com/client9/misspell/cmd/misspell

GOCOVMERGE = $(TOOLS)/gocovmerge
$(TOOLS)/gocovmerge: PACKAGE=github.com/wadey/gocovmerge

STRINGER = $(TOOLS)/stringer
$(TOOLS)/stringer: PACKAGE=golang.org/x/tools/cmd/stringer

PORTO = $(TOOLS)/porto
$(TOOLS)/porto: PACKAGE=github.com/jcchavezs/porto/cmd/porto

GOJQ = $(TOOLS)/gojq
$(TOOLS)/gojq: PACKAGE=github.com/itchyny/gojq/cmd/gojq

.PHONY: tools
tools: $(CROSSLINK) $(DBOTCONF) $(GOLANGCI_LINT) $(MISSPELL) $(GOCOVMERGE) $(STRINGER) $(PORTO) $(GOJQ) $(SEMCONVGEN) $(MULTIMOD) $(SEMCONVKIT)

# Build

.PHONY: generate build

generate: $(OTEL_GO_MOD_DIRS:%=generate/%)
generate/%: DIR=$*
generate/%: | $(STRINGER) $(PORTO)
	@echo "$(GO) generate $(DIR)/..." \
		&& cd $(DIR) \
		&& PATH="$(TOOLS):$${PATH}" $(GO) generate ./... && $(PORTO) -w .

build: generate $(OTEL_GO_MOD_DIRS:%=build/%) $(OTEL_GO_MOD_DIRS:%=build-tests/%)
build/%: DIR=$*
build/%:
	@echo "$(GO) build $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) build ./...

build-tests/%: DIR=$*
build-tests/%:
	@echo "$(GO) build tests $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) list ./... \
		| grep -v third_party \
		| xargs $(GO) test -vet=off -run xxxxxMatchNothingxxxxx >/dev/null

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test
test-default test-race: ARGS=-race
test-bench:   ARGS=-run=xxxxxMatchNothingxxxxx -test.benchtime=1ms -bench=.
test-short:   ARGS=-short
test-verbose: ARGS=-v -race
$(TEST_TARGETS): test
test: $(OTEL_GO_MOD_DIRS:%=test/%)
test/%: DIR=$*
test/%:
	@echo "$(GO) test -timeout $(TIMEOUT)s $(ARGS) $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) list ./... \
		| grep -v third_party \
		| xargs $(GO) test -timeout $(TIMEOUT)s $(ARGS)

COVERAGE_MODE    = atomic
COVERAGE_PROFILE = coverage.out
.PHONY: test-coverage
test-coverage: | $(GOCOVMERGE)
	@set -e; \
	printf "" > coverage.txt; \
	for dir in $(ALL_COVERAGE_MOD_DIRS); do \
	  echo "$(GO) test -coverpkg=go.opentelemetry.io/otel/... -covermode=$(COVERAGE_MODE) -coverprofile="$(COVERAGE_PROFILE)" $${dir}/..."; \
	  (cd "$${dir}" && \
	    $(GO) list ./... \
	    | grep -v third_party \
	    | grep -v 'semconv/v.*' \
	    | xargs $(GO) test -coverpkg=./... -covermode=$(COVERAGE_MODE) -coverprofile="$(COVERAGE_PROFILE)" && \
	  $(GO) tool cover -html=coverage.out -o coverage.html); \
	done; \
	$(GOCOVMERGE) $$(find . -name coverage.out) > coverage.txt

.PHONY: golangci-lint golangci-lint-fix
golangci-lint-fix: ARGS=--fix
golangci-lint-fix: golangci-lint
golangci-lint: $(OTEL_GO_MOD_DIRS:%=golangci-lint/%)
golangci-lint/%: DIR=$*
golangci-lint/%: | $(GOLANGCI_LINT)
	@echo 'golangci-lint $(if $(ARGS),$(ARGS) ,)$(DIR)' \
		&& cd $(DIR) \
		&& $(GOLANGCI_LINT) run --allow-serial-runners $(ARGS)

.PHONY: crosslink
crosslink: | $(CROSSLINK)
	@echo "Updating intra-repository dependencies in all go modules" \
		&& $(CROSSLINK) --root=$(shell pwd) --prune

.PHONY: go-mod-tidy
go-mod-tidy: $(ALL_GO_MOD_DIRS:%=go-mod-tidy/%)
go-mod-tidy/%: DIR=$*
go-mod-tidy/%: | crosslink
	@echo "$(GO) mod tidy in $(DIR)" \
		&& cd $(DIR) \
		&& $(GO) mod tidy -compat=1.19

.PHONY: lint-modules
lint-modules: go-mod-tidy

.PHONY: lint
lint: misspell lint-modules golangci-lint

.PHONY: vanity-import-check
vanity-import-check: | $(PORTO)
	@$(PORTO) --include-internal -l . || echo "(run: make vanity-import-fix)"

.PHONY: vanity-import-fix
vanity-import-fix: | $(PORTO)
	@$(PORTO) --include-internal -w .

.PHONY: misspell
misspell: | $(MISSPELL)
	@$(MISSPELL) -w $(ALL_DOCS)

.PHONY: license-check
license-check:
	@licRes=$$(for f in $$(find . -type f \( -iname '*.go' -o -iname '*.sh' \) ! -path '**/third_party/*' ! -path './.git/*' ) ; do \
	           awk '/Copyright The Solomon AI Authors|generated|GENERATED/ && NR<=3 { found=1; next } END { if (!found) print FILENAME }' $$f; \
	   done); \
	   if [ -n "$${licRes}" ]; then \
	           echo "license header checking failed:"; echo "$${licRes}"; \
	           exit 1; \
	   fi

DEPENDABOT_CONFIG = .github/dependabot.yml
.PHONY: dependabot-check
dependabot-check: | $(DBOTCONF)
	@$(DBOTCONF) verify $(DEPENDABOT_CONFIG) || echo "(run: make dependabot-generate)"

.PHONY: dependabot-generate
dependabot-generate: | $(DBOTCONF)
	@$(DBOTCONF) generate > $(DEPENDABOT_CONFIG)

.PHONY: check-clean-work-tree
check-clean-work-tree:
	@if ! git diff --quiet; then \
	  echo; \
	  echo 'Working tree is not clean, did you forget to run "make precommit"?'; \
	  echo; \
	  git status; \
	  exit 1; \
	fi

SEMCONVPKG ?= "semconv/"
.PHONY: semconv-generate
semconv-generate: | $(SEMCONVGEN) $(SEMCONVKIT)
	[ "$(TAG)" ] || ( echo "TAG unset: missing opentelemetry specification tag"; exit 1 )
	[ "$(OTEL_SPEC_REPO)" ] || ( echo "OTEL_SPEC_REPO unset: missing path to opentelemetry specification repo"; exit 1 )
	$(SEMCONVGEN) -i "$(OTEL_SPEC_REPO)/semantic_conventions/." --only=span -p conventionType=trace -f trace.go -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	$(SEMCONVGEN) -i "$(OTEL_SPEC_REPO)/semantic_conventions/." --only=event -p conventionType=event -f event.go -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	$(SEMCONVGEN) -i "$(OTEL_SPEC_REPO)/semantic_conventions/." --only=resource -p conventionType=resource -f resource.go -t "$(SEMCONVPKG)/template.j2" -s "$(TAG)"
	$(SEMCONVKIT) -output "$(SEMCONVPKG)/$(TAG)" -tag "$(TAG)"

.PHONY: prerelease
prerelease: | $(MULTIMOD)
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) verify && $(MULTIMOD) prerelease -m ${MODSET}

COMMIT ?= "HEAD"
.PHONY: add-tags
add-tags: | $(MULTIMOD)
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) verify && $(MULTIMOD) tag -m ${MODSET} -c ${COMMIT}

run-tests:
	@echo "running core-library tests"
	@echo "----- running core-auth-sdk tests"
	cd ./auth_client && go test && cd ..

.PHONY: add-license
add-license: ## Find all .go files not in the vendor directory and try to write a license notice.
	find . -path ./vendor -prune -o -type f -name "*.go" -print | xargs ./etc/add_license.sh
	# Check for any changes made with -G. to ignore permissions changes. Exit with a non-zero
	# exit code if there is a diff.
	git diff -G. --quiet

changelog:
	git-chglog --output CHANGELOG.md && git add CHANGELOG.md && git commit -m 'updated CHANGELOG.md'

gen:
	cd queue-message-format && make && cd ..

# Define a function to update version
define update-version
	@next="$(1)"; \
	current="$(VERSION)"; \
	echo "version-set: current: $$current, next: $$next"; \
	FILES="./version.go"; \
	for file in $$FILES; do \
		/usr/bin/sed -i '' "s/$$current/$$next/g" $$file; \
	done; \
	echo "Version $$next set in code, deployment, chart and kustomize";
endef

minor-version-set:
	echo hello
	$(call update-version,$(MINOR_VERSION))

major-version-set:
	$(call update-version,$(MAJOR_VERSION))

patch-version-set:
	$(call update-version,$(PATCH_VERSION))

release-minor-version: 
	echo "Releasing $(MINOR_VERSION)"
	git checkout -b release-$(MINOR_VERSION)
	make minor-version-set
	git add .
	git commit -m "bumping version from $(VERSION) to $(MINOR_VERSION)"
	git tag v$(MINOR_VERSION)
	git push --set-upstream origin release-$(MINOR_VERSION)
	git push origin v$(MINOR_VERSION)

release-major-version: 
	echo "Releasing $(MAJOR_VERSION)"
	git checkout -b release-$(MAJOR_VERSION)
	make major-version-set
	git add .
	git commit -m "bumping version from $(VERSION) to $(MAJOR_VERSION)"
	git tag v$(MAJOR_VERSION)
	git push --set-upstream origin release-$(MAJOR_VERSION)
	git push origin v$(MAJOR_VERSION)

release-patch-version:
	echo "Releasing $(PATCH_VERSION)"
	git checkout -b release-$(PATCH_VERSION)
	make patch-version-set
	git add .
	git commit -m "bumping version from $(VERSION) to $(PATCH_VERSION)"
	git tag v$(PATCH_VERSION)
	git push --set-upstream origin release-$(PATCH_VERSION)
	git push origin v$(PATCH_VERSION)

fmt:
	gofmt -l -s -w ./..
