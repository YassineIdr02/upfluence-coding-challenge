PHONY_TARGETS=server dev stop install
.PHONY: $(PHONY_TARGETS)

CLIENT_DIR=client
SERVER_DIR=server
CLIENT_PORT=5173
SERVER_PORT=8080

install:
	@echo "Installing client dependencies..."
	cd $(CLIENT_DIR) && npm install

server:
	cd $(SERVER_DIR) && go run .

dev:
	@echo "Starting server and client..."
	cd $(SERVER_DIR) && go run . & \
	cd $(CLIENT_DIR) && npm run dev

stop:
	@echo "Stopping server and client..."
	@pkill -f "go run" || true
	@pkill -f "vite" || true