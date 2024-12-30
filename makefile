SERVICES_DIR = ./services
FEDERATED_GRAPH_SERVICE_DIR = $(SERVICES_DIR)/federated-graph
INVENTORY_SERVICE_DIR = $(SERVICES_DIR)/inventory
ORDER_SERVICE_DIR = $(SERVICES_DIR)/order
PRODUCT_SERVICE_DIR = $(SERVICES_DIR)/product
TOOLS_DIR = ./internal/tools

FEDERATED_GRAPH_SERVICE_NAME = federated-graph-service
INVENTORY_SERVICE_NAME = inventory-service
PRODUCT_SERVICE_NAME = product-service
ORDER_SERVICE_NAME = order-service

BASE_BUILDER = base-graphkubel-builder
BASE_PRODUCTION = base-graphkubel-production
FEDERATED_GRAPH_IMAGE = $(FEDERATED_GRAPH_SERVICE_NAME):latest
INVENTORY_IMAGE = $(INVENTORY_SERVICE_NAME):latest
PRODUCT_IMAGE = $(PRODUCT_SERVICE_NAME):latest
ORDER_IMAGE = $(ORDER_SERVICE_NAME):latest

ALPINE_VERSION = 3.14

WORKDIR = /app

.PHONY: all
all: build

# Build
.PHONY: build
build: build-federated-graph build-product build-inventory build-order 

.PHONY: build-base-builder
build-base-builder:
	@echo "Building base builder image..."
	docker build -t $(BASE_BUILDER) -f Dockerfile.builder .

.PHONY: build-base-production
build-base-production:
	@echo "Building base production image..."
	docker build --build-arg ALPINE_VERSION=$(ALPINE_VERSION) -t $(BASE_PRODUCTION) -f Dockerfile.production .
	
.PHONY: build-base-images
build-base-images: build-base-builder build-base-production

.PHONY: build-federated-graph
build-federated-graph: 
	@echo "Building Federated-Graph Service..."
	cd $(FEDERATED_GRAPH_SERVICE_DIR) && docker build -t $(FEDERATED_GRAPH_IMAGE) -f Dockerfile.federated-graph .

.PHONY: build-inventory
build-inventory: build-base-images
	@echo "Building Inventory Service..."
	cd $(INVENTORY_SERVICE_DIR) && docker build  -t $(INVENTORY_IMAGE) -f Dockerfile.inventory .

.PHONY: build-order
build-order: build-base-images
	@echo "Building Order Service..."
	cd $(ORDER_SERVICE_DIR) && docker build --build-arg ALPINE_VERSION=$(ALPINE_VERSION) -t $(ORDER_IMAGE) -f Dockerfile.order .

.PHONY: build-product
build-product: build-base-images
	@echo "Building Product Service..."
	cd $(PRODUCT_SERVICE_DIR) && go generate && docker build --build-arg ALPINE_VERSION=$(ALPINE_VERSION) -t $(PRODUCT_IMAGE) -f Dockerfile.product .

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

.PHONY: install-tools
install-tools:
	@echo "Installing tools..."
	chmod +x $(TOOLS_DIR)/install-tools.sh && $(TOOLS_DIR)/install-tools.sh

.PHONY: migrate-db
migrate-db:
	@echo -e "\033[1;35mMigrating order and product databases...\033[0m"
	./migrateDb.sh

# Default target
.PHONY: default
default: all
