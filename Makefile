# Variables
DOCKER_IMAGE = redis:latest
CONTAINER_NAME = redis_go_locate
HOST_PORT = 6379
CONTAINER_PORT = 6379
VOLUME_NAME = redis_data
REDIS_VOLUME = $(VOLUME_NAME):/data  # Redis stores data in /data inside the container

# Docker commands
DOCKER_RUN = docker run
DOCKER_STOP = docker stop
DOCKER_RM = docker rm
DOCKER_VOLUME_CREATE = docker volume create
DOCKER_VOLUME_RM = docker volume rm

# Default target: build and run Redis container
.PHONY: all
all: run

# Create a Docker volume if it doesn't exist
.PHONY: create-volume
create-volume:
	@echo "Creating Docker volume: $(VOLUME_NAME)"
	@$(DOCKER_VOLUME_CREATE) $(VOLUME_NAME)

# Run Redis container with volume and port mapping
.PHONY: run
run: create-volume
	@echo "Running Redis container: $(CONTAINER_NAME)"
	@$(DOCKER_RUN) --name $(CONTAINER_NAME) -d -p $(HOST_PORT):$(CONTAINER_PORT) -v $(REDIS_VOLUME) $(DOCKER_IMAGE)

# Stop Redis container
.PHONY: stop
stop:
	@echo "Stopping Redis container: $(CONTAINER_NAME)"
	@$(DOCKER_STOP) $(CONTAINER_NAME)

# Remove Redis container
.PHONY: remove
remove:
	@echo "Removing Redis container: $(CONTAINER_NAME)"
	@$(DOCKER_RM) $(CONTAINER_NAME)

# Clean up: stop and remove container, remove volume
.PHONY: clean
clean: stop remove
	@echo "Removing Redis volume: $(VOLUME_NAME)"
	@$(DOCKER_VOLUME_RM) $(VOLUME_NAME)

# Status: show running Docker containers
.PHONY: status
status:
	@docker ps -a | grep $(CONTAINER_NAME) || echo "No container with name $(CONTAINER_NAME) is running."