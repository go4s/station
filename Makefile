BUILD_VERSION=`git rev-parse --short HEAD`
BUILD_DATETIME=`date +"%Y-%m-%d %H:%M:%S"`
DOCKER_CLI=`podman`
TAG=`debug`

all: .PHONY
	@echo ">> Building:			[$(MODULE_NAME)]"
	@echo ">> Version:			[$(BUILD_VERSION)]"
	@echo ">> DateTime:			[$(BUILD_DATETIME)]"
	@go build -mod=vendor -tags netgo,musl																				\
		-ldflags "-X \"main.BuildVersion=$(BUILD_VERSION)\""															\
		-ldflags "-X \"main.BuildDateTime=$(BUILD_DATETIME)\""															\
		-o /go/bin/$(MODULE_NAME) $(MODULE_PATH)
	@chmod 700 bootstrap_generator.sh && ./bootstrap_generator.sh

image: tidy vendor
	@echo ">> ModuleName:				[$(MODULE_NAME)]"
	@echo ">> ModulePath:				[$(MODULE_PATH)]"
	@$(DOCKER_CLI) build . -t ghcr.io/go4s/$(MODULE_NAME):$(TAG)														\
						--build-arg MODULE_PATH=$(MODULE_PATH)															\
						--build-arg MODULE_NAME=$(MODULE_NAME)

tidy:
	@go mod tidy

vendor:
	@go mod vendor

.PHONY:


