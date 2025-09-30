PHONY_TARGETS=server client dev stop
.PHONY: $(PHONY_TARGETS)

CLIENT_DIR=client
SERVER_DIR=server
CLIENT_PORT=5173
SERVER_PORT=8080

server:
	cd $(SERVER_DIR) && go run .

client:
	cd $(CLIENT_DIR) && npm run dev

dev:
	@echo "Starting backend and frontend..."
	cd $(SERVER_DIR) && go run . & \
	cd $(CLIENT_DIR) && npm run dev

stop:
	@echo "Stopping servers..."
	@pkill -f "go run" || true
	@pkill -f "vite" || true