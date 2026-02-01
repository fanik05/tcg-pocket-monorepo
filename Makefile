# Variables
API_DIR = apps/api
WEB_DIR = apps/web

# Default command
all: help

# Install dependencies for both apps
install:
	@echo "Installing Go dependencies..."
	cd $(API_DIR) && go mod tidy
	@echo "Installing Next.js dependencies..."
	cd $(WEB_DIR) && npm install

# Run the Go backend
api-dev:
	cd $(API_DIR) && go run main.go

# Run the Next.js frontend
web-dev:
	cd $(WEB_DIR) && npm run dev

# Run BOTH at the same time (The "Magic" command)
dev:
	make -j 2 api-dev web-dev

# Help menu
help:
	@echo "Available commands:"
	@echo "  make install  - Install all dependencies"
	@echo "  make dev      - Start both API and Web in development mode"
	@echo "  make api-dev  - Start only the Go backend"
	@echo "  make web-dev  - Start only the Next.js frontend"