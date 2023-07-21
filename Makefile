ifndef CURDIR
$(error CURDIR is not set)
endif
$(info Current directory: $(CURDIR))

include date.mk

PACKAGE  				= ${shell pwd | rev | cut -f1 -d'/' - | rev}
DATE    				?= $(shell date +%Y-%m-%d_%I:%M:%S%p)
GITHASH 				= $(shell git rev-parse HEAD)
VERSION 				?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
							cat $(PACKAGE)/.version 2> /dev/null || echo v0)
PKGS     				= $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))
TESTPKGS 				= $(shell env GO111MODULE=on $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
BIN      				= $(GOPATH)/bin
DOCKER_BUILD_CONTEXT	=.
DOCKER_FILE_PATH		=Dockerfile
GO      				= go
TIMEOUT 				= 300
V 						= 0
Q 						= $(if $(filter 1,$V),,@)
M 						= $(shell printf "\033[34;1m▶\033[0m")
os = $(shell uname)

export GO111MODULE=on
export MY_POD_NAMESPACE=development

$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) building $(REPOSITORY)…)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=off GOPATH=$$tmp GOBIN=$(BIN) $(GO) get $(REPOSITORY) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOLINT = $(BIN)/golint
$(BIN)/golint: REPOSITORY=golang.org/x/lint/golint
$(info GOLINT: $(GOLINT))

GOCOVMERGE = $(BIN)/gocovmerge
$(BIN)/gocovmerge: REPOSITORY=github.com/wadey/gocovmerge
$(info GOCOVMERGE: $(GOCOVMERGE))

GOCOV = $(BIN)/gocov
$(BIN)/gocov: REPOSITORY=github.com/axw/gocov/...
$(info GOCOV: $(GOCOV))

GOCOVXML = $(BIN)/gocov-xml
$(BIN)/gocov-xml: REPOSITORY=github.com/AlekSi/gocov-xml
$(info GOCOVXML: $(GOCOVXML))

GO2XUNIT = $(BIN)/go2xunit
$(BIN)/go2xunit: REPOSITORY=github.com/tebeka/go2xunit
$(info GO2XUNIT: $(GO2XUNIT))


########################################################################################################################
##########                                                                                                    ##########
########## (~˘▾˘)~  (~˘▾˘)~  (~˘▾˘)~  (~˘▾˘)~  (~˘▾˘)~  RECIPES  ~(˘▾˘~)  ~(˘▾˘~)  ~(˘▾˘~)  ~(˘▾˘~)  ~(˘▾˘~)  ##########
##########                                                                                                    ##########
########################################################################################################################

# Removes and recreates a directory
clean-directory:
	rm -rf ${DIRECTORY}
	mkdir -p ${DIRECTORY}

.DEFAULT_GOAL := build
build: all

.PHONY: all
all: fmt test $(BIN) ; $(info $(M) building executable…) @ ## Build program binary
	$Q $(GO) build -tags release -ldflags '-X main.GitComHash=$(GITHASH) -X main.BuildStamp=$(DATE)' -o bin/application cmd/server/main.go

docker: all ; $(info $(M) building container...) @ ## Build docker container
	@docker build $(DOCKER_BUILD_ARGS) -t $(PACKAGE):$(GITHASH) $(DOCKER_BUILD_CONTEXT) -f $(DOCKER_FILE_PATH)
	@docker tag $(PACKAGE):$(GITHASH) $(PACKAGE):latest

#Testing v1.13
COVERAGE_FILE:=$(CURDIR)/test/coverage.$(DATE_TIME).out
COVERAGE_HTML_FILE:=$(COVERAGE_FILE).html
$(info Coverage Profile Output : $(COVERAGE_FILE))
$(info Coverage HTML Output: $(COVERAGE_HTML_FILE))

.PHONY: test-coverage-report
test-coverage-report:
	$(info Calling Go tests...)
	$(GO) test -timeout=5s -bench=. -run=. -cover -coverprofile=$(COVERAGE_FILE) -covermode=count -v ./...
	$(info Calling Go HTML coverage converter...)
	$(GO) tool cover -html=$(COVERAGE_FILE) -o=$(COVERAGE_HTML_FILE)
	$(info Go test coverage calls issued.)

#Testing tools
TEST_TARGETS := test-default test-bench test-short test-verbose test-race test-cover
.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
test-cover:   ARGS=-cover        ## Run test with basic coverage
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: fmt lint ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q $(GO) test -timeout $(TIMEOUT)s $(ARGS) $(TESTPKGS)

test-xml: fmt lint | $(GO2XUNIT) ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests with xUnit output
	$Q mkdir -p test
	$Q 2>&1 $(GO) test -timeout 20s -v $(TESTPKGS) | tee test/tests.output
	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml
COVERAGE_MODE = atomic
COVERAGE_DIR := $(CURDIR)/test/coverage.$(DATE_TIME)
COVERAGE_PROFILE = $(COVERAGE_DIR)/profile.out
COVERAGE_XML = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML = $(COVERAGE_DIR)/index.html

.PHONY: test-coverage test-coverage-tools
test-coverage-tools: | $(GOCOVMERGE) $(GOCOV) $(GOCOVXML)
test-coverage:
	fmt lint test-coverage-tools ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)/coverage
	$Q for pkg in $(TESTPKGS); do \
		$(GO) test \
			-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $$pkg | \
					grep '^$(PACKAGE)/' | \
					tr '\n' ',')$$pkg \
			-covermode=$(COVERAGE_MODE) \
			-coverprofile="$(COVERAGE_DIR)/coverage/`echo $$pkg | tr "/" "-"`.cover" $$pkg ;\
	done
	$Q $(GOCOVMERGE) $(COVERAGE_DIR)/coverage/*.cover > $(COVERAGE_PROFILE)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)


#lint
.PHONY: lint
lint: ; exit 0

#fmt
.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	$Q $(GO) fmt ./...

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf $(BIN)
	@rm -rf test/tests.* test/coverage.*

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: vendor
vendor: ; $(info $(M) running go mod vendor…) @ ## Run go mod vendor
	$Q $(GO) mod vendor

run: docker
	  docker run -p 8000:8000 "${PACKAGE}:$(GITHASH)"
