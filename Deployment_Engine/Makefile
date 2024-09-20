# Define directories
BACKEND_DIR=./Backend
FRONTEND_DIR=./Frontend
HOSTS_FILE=./hosts.ini

# Define commands for backend (Go)
backend-deps-install:
	@echo "Installing Go dependencies for backend..."
	cd $(BACKEND_DIR) && go mod tidy

backend-deps-delete:
	@echo "Removing Go dependencies for backend..."
	cd $(BACKEND_DIR) && rm -rf $(BACKEND_DIR)/go.sum $(BACKEND_DIR)/go.mod $(BACKEND_DIR)/vendor

backend-deps-reinstall: backend-deps-delete backend-deps-install
	@echo "Reinstalling Go dependencies for backend..."

# Define commands for frontend (React)
frontend-deps-install:
	@echo "Installing React dependencies for frontend..."
	cd $(FRONTEND_DIR) && npm install --force

frontend-deps-delete:
	@echo "Removing React dependencies for frontend..."
	cd $(FRONTEND_DIR) && rm -rf node_modules

frontend-deps-reinstall: frontend-deps-delete frontend-deps-install
	@echo "Reinstalling React dependencies for frontend..."

# Define commands for cleaning and installing all dependencies
clean-all: backend-deps-delete frontend-deps-delete remove-hosts
	@echo "Cleaned up all dependencies and removed hosts.ini."

remove-hosts:
	@if [ -f $(HOSTS_FILE) ]; then \
		echo "Removing hosts.ini file..."; \
		rm $(HOSTS_FILE); \
	else \
		echo "hosts.ini file not found, skipping..."; \
	fi

install-all: backend-deps-install frontend-deps-install
	@echo "Installed all dependencies."

reinstall-all: backend-deps-reinstall frontend-deps-reinstall
	@echo "Reinstalled all dependencies."

backend-run:
	@echo "Running Go backend server..."
	cd $(BACKEND_DIR) && go run .

frontend-run:
	@echo "Running React frontend server..."
	cd $(FRONTEND_DIR) && npm start


# Help message
help:
	@echo "Usage:"
	@echo "  make backend-deps-install   - Install backend dependencies (Go)"
	@echo "  make backend-deps-delete    - Delete backend dependencies (Go)"
	@echo "  make backend-deps-reinstall - Reinstall backend dependencies (Go)"
	@echo "  make frontend-deps-install  - Install frontend dependencies (React)"
	@echo "  make frontend-deps-delete   - Delete frontend dependencies (React)"
	@echo "  make frontend-deps-reinstall- Reinstall frontend dependencies (React)"
	@echo "  make clean-all              - Delete both backend and frontend dependencies and remove hosts.ini"
	@echo "  make install-all            - Install both backend and frontend dependencies"
	@echo "  make reinstall-all          - Reinstall both backend and frontend dependencies"
