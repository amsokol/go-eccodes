IMAGE_BUILDER_NAME = amsokol/go-eccodes/centos-builder
IMAGE_BUILDER_VERSION = latest

IMAGE_EXAMPLE1_NAME = amsokol/go-eccodes/example-get-all-messages-from-file
IMAGE_EXAMPLE1_VERSION = latest

IMAGE_EXAMPLE2_NAME = amsokol/go-eccodes/example-get-messages-from-file-by-index
IMAGE_EXAMPLE2_VERSION = latest

build: build-examples

test: run-example-get-all-messages-from-file-docker run-example-get-messages-from-file-by-index-docker

build-examples: build-example-get-all-messages-from-file-docker build-example-get-messages-from-file-by-index-docker

build-builder-docker:
	docker build -t $(IMAGE_BUILDER_NAME):$(IMAGE_BUILDER_VERSION) -f CentOS-Builder.Dockerfile .

build-example-get-all-messages-from-file-docker: build-builder-docker
	docker build -t $(IMAGE_EXAMPLE1_NAME):$(IMAGE_EXAMPLE1_VERSION) -f example-get-all-messages-from-file.Dockerfile .

build-example-get-messages-from-file-by-index-docker: build-builder-docker
	docker build -t $(IMAGE_EXAMPLE2_NAME):$(IMAGE_EXAMPLE2_VERSION) -f example-get-messages-from-file-by-index.Dockerfile .

run-example-get-all-messages-from-file-docker: build-example-get-all-messages-from-file-docker
	docker run --name go-eccodes-example-get-all-messages-from-file --rm $(IMAGE_EXAMPLE1_NAME):$(IMAGE_EXAMPLE1_VERSION)

run-example-get-messages-from-file-by-index-docker: build-example-get-messages-from-file-by-index-docker
	docker run --name go-eccodes-example-get-messages-from-file-by-index --rm $(IMAGE_EXAMPLE2_NAME):$(IMAGE_EXAMPLE2_VERSION)
