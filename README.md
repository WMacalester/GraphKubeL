# GraphKubeL
A short exploration into using kubernetes and graphql in a microservice architecture

## Overview
 A backend system for a simple e-commerce platform that allows users to browse products, place orders, and check inventory levels. Each service (e.g., product catalog, orders, inventory) is a separate microservice, deployed and orchestrated on Kubernetes, and exposed through a unified GraphQL API. The following are a list of steps / goals to guide this exploration

## Goals
### Microservices:
#### Product Service:
* Manages the catalog of products.
* Allows querying for products, categories, descriptions, prices, and stock levels.

#### Order Service:
* Handles placing and tracking customer orders.
* Enables querying for order history, statuses, and details about specific orders.

#### Inventory Service:
* Tracks inventory levels across warehouses.
* Updates stock levels when orders are placed or products are restocked.

#### User Service (Optional stretch goal):
* Manages user accounts, authentication, and user preferences.
* Provides endpoints to query user details, saved addresses, etc.
* Learning Opportunities:

### GraphQL
* Use GraphQL as a gateway to unify data from the Product, Order, and Inventory services.
* Query the API to retrieve products with their stock status, and place an order in one request.
#### Stretch goals
* Use GraphQL subscriptions to notify clients when an order status changes or when product stock is updated.
* Federate using apollo server

### Kubernetes Orchestration (Locally):
* Run each microservice (Product, Order, Inventory) locally on a Kubernetes cluster using Minikube, Kind (Kubernetes in Docker), or K3s.
* Learn how to define Kubernetes manifests (YAML files) for Deployments, Services, and ConfigMaps to orchestrate your microservices locally.
* Simulate scaling and service discovery by running multiple replicas of a service in your local Kubernetes environment.
* Use kubectl to manage, monitor, and inspect the status of your services while developing locally.

#### Service Communication:
* Use Kubernetes Services to allow your microservices (e.g., Product, Order, Inventory) to communicate with each other.
* Learn to configure internal service-to-service communication and expose external endpoints via Kubernetes Ingress for testing your GraphQL API locally.

#### Environment-Specific Configuration:
* Use Kubernetes ConfigMaps and Secrets to manage configurations for each microservice locally.
* Learn how to inject configuration and secrets (e.g., database URLs, API keys) into your containers in a local environment.

#### Testing Locally in Kubernetes:
* Create local integration tests for your microservices using the local Kubernetes cluster.
* Use tools like Telepresence to run a single service locally while still connecting it to a Kubernetes cluster for testing interactions with other services.
* Use Tilt or Skaffold for rapid iteration in Kubernetes.

#### Local API Gateway/Ingress:
* Set up a local Ingress controller (e.g., NGINX Ingress) to route traffic to the correct microservices within your local Kubernetes environment.
* Learn how to configure local load balancing and traffic routing to the appropriate service behind your GraphQL API.
* Experiment with exposing your services externally on your local machine, and simulate traffic routing between services (e.g., accessing the Product API or Order API through the GraphQL gateway).
* Optionally, explore tools like Istio for service mesh, traffic routing, and monitoring.

#### Local Monitoring and Logging:
* Learn how to monitor and log your services locally using tools like Prometheus (for monitoring) and Grafana (for visualization) in your local Kubernetes environment.
* Set up ELK Stack (Elasticsearch, Logstash, Kibana) or Loki for logging microservices locally, capturing logs, and troubleshooting locally.

* Run monitoring and logging tools inside your local Kubernetes cluster to see the health of your microservices.
* Observe performance metrics, error logs, and resource utilization in real-time while testing your GraphQL queries.

#### Data Storage in Local Kubernetes Cluster:
* Run local databases (e.g., PostgreSQL, MongoDB, Redis) inside Kubernetes pods, with persistent volumes set up for each service.
* Learn how to connect microservices to their respective data stores locally, simulating production-like environments.
* Experiment with persistence by configuring PersistentVolumeClaims (PVCs) in Kubernetes locally to ensure data is stored correctly across restarts.
