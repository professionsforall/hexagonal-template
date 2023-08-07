
MOD_PATH=github.com/professionsforall/hexagonal-template

GOCMD := $(shell [ -z "$(GOCMD)" ] && echo "go" || echo "$(GOCMD)")
GOVET=$(GOCMD) vet
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
UPX=/usr/bin/upx
DOCKER_NAME=$(BUILD_TYPE).$(PROJECT_NAME)
DOCKER_LOCAL_NAME=docker.local:5000/$(BUILD_TYPE).$(PROJECT_NAME)
GRPC=protoc
DOCKER=docker
GIT=/usr/bin/git
export GOSUMDB=off
export GOPRIVATE="git.pio.ir/*"

RELEASE?=$(shell cat VERSION || echo "none")
COMMIT?=$(shell git rev-parse --short HEAD)
ts := $(shell /bin/date "+%Y%m%d%H%M%S")
BUILD_TYPE := $(shell [ -z "$(BUILD_TYPE)" ] && echo "dev" || echo "$(BUILD_TYPE)")

all: mkbuild get tidy swag build deps

mkbuild:
	@mkdir -p build.$(BUILD_TYPE)
debug:
	@$(GOBUILD) -tags netgo -tags $(BUILD_TYPE) -gcflags="-N -l" -ldflags "-X $(MOD_PATH)/internal.BuildDate=$(ts) -X $(MOD_PATH)/internal.Version=$(RELEASE) -X $(MOD_PATH)/internal.Commit=$(COMMIT) -X $(MOD_PATH)/internal.AppName=$(PROJECT_NAME)  -X $(MOD_PATH)/internal.ProjectDomainName=$(PROJECT_DOMAIN_NAME)  -X $(MOD_PATH)/internal.BuildType=$(BUILD_TYPE) -X $(MOD_PATH)/internal.ProjectNickname=$(PROJECT_NICK_NAME)" -o build.$(BUILD_TYPE)/$(PROJECT_NAME) -v
redebug:
	@$(GOBUILD) -tags netgo -tags $(BUILD_TYPE) -gcflags="-N -l" -ldflags "-X $(MOD_PATH)/internal.BuildDate=$(ts) -X $(MOD_PATH)/internal.Version=$(RELEASE) -X $(MOD_PATH)/internal.Commit=$(COMMIT) -X $(MOD_PATH)/internal.AppName=$(PROJECT_NAME)  -X $(MOD_PATH)/internal.ProjectDomainName=$(PROJECT_DOMAIN_NAME)  -X $(MOD_PATH)/internal.BuildType=$(BUILD_TYPE) -X $(MOD_PATH)/internal.ProjectNickname=$(PROJECT_NICK_NAME)" -a -o build.$(BUILD_TYPE)/$(PROJECT_NAME) -v
build:
	@$(GOBUILD) -tags netgo -tags $(BUILD_TYPE) -ldflags='-extldflags "-static" -s -w'  -o build.$(BUILD_TYPE)/$(PROJECT_NAME) -v
rebuild:
	@$(GOBUILD) -tags netgo -tags $(BUILD_TYPE) -ldflags='-extldflags "-static" -s -w -X $(MOD_PATH)/internal.BuildDate=$(ts) -X $(MOD_PATH)/internal.Version=$(RELEASE) -X $(MOD_PATH)/internal.Commit=$(COMMIT) -X $(MOD_PATH)/internal.AppName=$(PROJECT_NAME) -X $(MOD_PATH)/internal.ProjectDomainName=$(PROJECT_DOMAIN_NAME)   -X $(MOD_PATH)/internal.BuildType=$(BUILD_TYPE) -X $(MOD_PATH)/internal.ProjectNickname=$(PROJECT_NICK_NAME)' -a -o build.$(BUILD_TYPE)/$(PROJECT_NAME) -v
escape:
	@$(GOBUILD) -tags netgo -tags $(BUILD_TYPE) -gcflags="-m" -ldflags "-s -w -X $(MOD_PATH)/internal.BuildDate=$(ts) -X $(MOD_PATH)/internal.Version=$(RELEASE) -X $(MOD_PATH)/internal.Commit=$(COMMIT) -X $(MOD_PATH)/internal.AppName=$(PROJECT_NAME) -X $(MOD_PATH)/internal.ProjectDomainName=$(PROJECT_DOMAIN_NAME)   -X $(MOD_PATH)/internal.BuildType=$(BUILD_TYPE) -X $(MOD_PATH)/internal.ProjectNickname=$(PROJECT_NICK_NAME)" -o build.$(BUILD_TYPE)/$(PROJECT_NAME) -v
get:
	@$(GOCMD) version
	@$(GOGET)  -tags $(BUILD_TYPE) -v ./...
tidy:
	@$(GOMOD) tidy
test:
	@$(GOTEST) -tags $(BUILD_TYPE) -v ./...
test_shuffle:
	@$(GOTEST) -tags $(BUILD_TYPE) -v -shuffle=on ./...
clean:
	@echo "Are you sure you want to clean the files(untracked files wil be removed too)? (Yes, default no)"
	@read -r response; \
    if [ "$$response" = "Yes" ]; then \
        git clean -fdx; \
        echo "Files cleaned."; \
    else \
        echo "Clean operation aborted."; \
    fi
test_html:
	@echo "Running tests"
	go test -v -shuffle=on -cover ./...
	go test  -covermode=count  -coverprofile=coverage.out  ./...  
	go tool cover -html=coverage.out
	
test_bench :
	@echo "Running benchmarks"
	go test -bench=.  -benchmem  ./... 

test_cpu :
	go test ./$(pn) -bench=. -benchmem -cpuprofile cpu.out -memprofile mem.out -v
	go tool pprof -http=:$(port) cpu.out

test_mem :
	go test ./$(pn) -bench=. -benchmem -cpuprofile cpu.out -memprofile mem.out -v
	go tool pprof -http=:$(port) mem.out

fmt:
	@echo "Formatting code"
	gofmt -l -s -w ./
	goimports -l -w ./

vet:
	@$(GOVET) ./...
deps:
	@[ -f "./deps.sh" ] && ./deps.sh || echo "no deps.sh"
burun: build
	@$(GORUN) -race .
swag:
	@swag init

test_clear :
	@rm -f coverage.out || rm -f cpu.out || rm -f mem.out || rm *.test
compress:
	@[ -f "$(UPX)" ] && $(UPX) -9 build.$(BUILD_TYPE)/$(PROJECT_NAME) || echo "upx not found"
run:
	@gnome-terminal --window --hide-menubar --title="Pio - $(BUILD_TYPE) $(PROJECT_NAME)" -- bash -c 'while  true ; do ./build.$(BUILD_TYPE)/$(PROJECT_NAME) || echo ">>>>>>> ERROR <<<<<<<<<<"; sleep 10; done'

docker-build:
	@$(DOCKER) build -f dockerfile.$(BUILD_TYPE) --rm -t $(DOCKER_NAME):$(RELEASE) .
	@$(DOCKER) build -f dockerfile.$(BUILD_TYPE) --rm -t $(DOCKER_NAME):latest .
docker-rebuild:
	@$(DOCKER) build -f dockerfile.$(BUILD_TYPE) --no-cache --rm -t $(DOCKER_NAME):$(RELEASE) .
	@$(DOCKER) build -f dockerfile.$(BUILD_TYPE) --no-cache --rm -t $(DOCKER_NAME):latest .
docker-push:
	@$(DOCKER) push -f dockerfile.$(BUILD_TYPE) $(DOCKER_NAME):$(RELEASE) .
	@$(DOCKER) push -f dockerfile.$(BUILD_TYPE) $(DOCKER_NAME):latest .
docker-purge:
	@$(DOCKER) image -f dockerfile.$(BUILD_TYPE) rm $(DOCKER_NAME) .
