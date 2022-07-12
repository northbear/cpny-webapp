APP_VERSION    := 0.0.1
IMAGE_REGISTRY := $(CNPY_ECR_REPOSITORY)
CNPY_APP_NAME  ?= webapp
IMAGE_NAME     := $(CNPY_APP_NAME)/web

EXECUTABLE      := build/web
IMAGE_FULL_NAME := $(if $(IMAGE_REGISTRY), $(IMAGE_REGISTRY)/$(IMAGE_NAME), $(IMAGE_NAME))

.PHONY: prepare clean all build-image push-image

all: prepare $(EXECUTABLE)

prepare:
	[ -d build ] || mkdir -p build

$(EXECUTABLE): main.go
	go build -o $(EXECUTABLE) -ldflags="-X main.APP_VERSION=$(APP_VERSION)" $<

build-image:
	buildah build --rm --no-cache -t $(IMAGE_FULL_NAME):$(APP_VERSION) .

push-image: build-image
	buildah push $(IMAGE_FULL_NAME):$(APP_VERSION)

clean:
	rm -rf build
