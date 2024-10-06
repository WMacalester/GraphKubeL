SERVICES_DIR = ./services
INVENTORY_SERVICE_DIR = $(SERVICES_DIR)/inventory
ORDER_SERVICE_DIR = $(SERVICES_DIR)/order
PRODUCT_SERVICE_DIR = $(SERVICES_DIR)/product

PRODUCT_SERVICE_NAME = product-service
INVENTORY_SERVICE_NAME = inventory-service
ORDER_SERVICE_NAME = order-service

BASE_IMAGE = base-graphkubel-image
PRODUCT_IMAGE = $(PRODUCT_SERVICE_NAME):latest
INVENTORY_IMAGE = $(INVENTORY_SERVICE_NAME):latest
ORDER_IMAGE = $(ORDER_SERVICE_NAME):latest

ALPINE_VERSION = 3.14

WORKDIR = /app

.PHONY: all
all: build

# Graphql Generation
generate:
	cd "$(DIRECTORY)" && go run github.com/99designs/gqlgen generate

generate-product: 
	$(MAKE) generate DIRECTORY=$(PRODUCT_SERVICE_DIR)

# Build
.PHONY: build
build: build-product build-inventory build-order

.PHONY: build-base-image
build-base-image:
	@echo "Building Inventory Service..."
	docker build -t $(BASE_IMAGE) -f Dockerfile.base .

.PHONY: build-inventory
build-inventory: build-base-image
	@echo "Building Inventory Service..."
	cd $(INVENTORY_SERVICE_DIR) && docker build --build-arg ALPINE_VERSION=$(ALPINE_VERSION) -t $(INVENTORY_IMAGE) -f Dockerfile.inventory .

.PHONY: build-order
build-order: build-base-image
	@echo "Building Order Service..."
	cd $(ORDER_SERVICE_DIR) && docker build --build-arg ALPINE_VERSION=$(ALPINE_VERSION) -t $(ORDER_IMAGE) -f Dockerfile.order .

.PHONY: build-product
build-product: generate-product build-base-image
	@echo "Building Product Service..."
	cd $(PRODUCT_SERVICE_DIR) && docker build --build-arg ALPINE_VERSION=$(ALPINE_VERSION) -t $(PRODUCT_IMAGE) -f Dockerfile.product .

# Run
.PHONY: run
run: run-product run-inventory run-order

.PHONY: run-inventory
run-inventory:
	@echo "Running Inventory Service..."
	docker run -p 8080:8080 $(INVENTORY_IMAGE)

.PHONY: run-order
run-order:
	@echo "Running Order Service..."
	docker run -p 8081:8080 $(ORDER_IMAGE)

.PHONY: run-product
run-product:
	@echo "Running Product Service..."
	docker run -p 8082:8080 $(PRODUCT_IMAGE)

# Clean
.PHONY: clean
clean:
	@echo "Cleaning up images..."
	docker rmi $(PRODUCT_IMAGE) $(INVENTORY_IMAGE) $(ORDER_IMAGE) || true

# Default target
.PHONY: default
default: all
