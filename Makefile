# Variables
DOCKER_IMAGE_NAME = my-go-app
DOCKER_TAG = latest
PORT = 8080
CONTAINER_NAME = go-app-container

# Build the Docker image
.PHONY: build
build:
	@echo "Building the Docker image..."
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

# Run the Docker container
.PHONY: run
run:
	@echo "Running the Docker container..."
	docker run --name $(CONTAINER_NAME) -p $(PORT):$(PORT) -d $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

# Stop and remove the running container
.PHONY: stop
stop:
	@echo "Stopping the Docker container..."
	docker stop $(CONTAINER_NAME) || true
	@echo "Removing the Docker container..."
	docker rm $(CONTAINER_NAME) || true

# Clean up the Docker image and container
.PHONY: clean
clean: stop
	@echo "Removing the Docker image..."
	docker rmi $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) || true

# Rebuild the Docker image and run the container
.PHONY: rebuild
rebuild: clean build run

# View Docker logs
.PHONY: logs
logs:
	@echo "Showing logs for the container..."
	docker logs -f $(CONTAINER_NAME)
