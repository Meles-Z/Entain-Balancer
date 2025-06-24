# ======================
# Docker Compose Commands
# ======================
.PHONY: up down build restart logs populate_db wipe_db

build:
	@echo "Rebuilding containers..."
	@docker-compose up --build -d

up:
	@echo "Starting containers..."
	@docker-compose up -d

down:
	@echo "Stopping containers..."
	@docker-compose down


restart: down up  

logs:
	@docker-compose logs -f

# ======================
# Database Management
# ======================
populate_db:
	@echo "Populating database..."
	@go run scripts/populate_db/populate.go

wipe_db:
	@echo "Wiping database..."  # Fixed echo message (was "Populating")
	@go run scripts/wipe_db/wipe.go

# ======================
# Help
# ======================
help:
	@echo "Available targets:"
	@echo "  up        - Start containers"
	@echo "  down      - Stop containers"
	@echo "  build     - Rebuild and start containers"
	@echo "  restart   - Restart containers (down + up)"
	@echo "  logs      - Follow container logs"
	@echo "  populate_db - Seed the database"
	@echo "  wipe_db   - Clear the database"