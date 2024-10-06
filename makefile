SERVICES_DIR = ./services
INVENTORY_SERVICE_DIR = $(SERVICES_DIR)/inventory
ORDER_SERVICE_DIR = $(SERVICES_DIR)/order
PRODUCT_SERVICE_DIR = $(SERVICES_DIR)/product

PRODUCT_SERVICE_NAME = product-service
INVENTORY_SERVICE_NAME = inventory-service
ORDER_SERVICE_NAME = order-service

PRODUCT_IMAGE = $(PRODUCT_SERVICE_NAME):latest
INVENTORY_IMAGE = $(INVENTORY_SERVICE_NAME):latest
ORDER_IMAGE = $(ORDER_SERVICE_NAME):latest

WORKDIR = /app

.PHONY: all
all: build

# Build
.PHONY: build
build: build-product build-inventory build-order

.PHONY: build-inventory
build-inventory:
	@echo "Building Inventory Service..."
	cd $(INVENTORY_SERVICE_DIR) && docker build -t $(INVENTORY_IMAGE) -f Dockerfile.inventory .

.PHONY: build-order
build-order:
	@echo "Building Order Service..."
	cd $(ORDER_SERVICE_DIR) && docker build -t $(ORDER_IMAGE) -f Dockerfile.order .

.PHONY: build-product
build-product:
	@echo "Building Product Service..."
	cd $(PRODUCT_SERVICE_DIR) && docker build -t $(PRODUCT_IMAGE) -f Dockerfile.product .

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
