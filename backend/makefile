
WAIT_FOR_IT = ./wait-for-it.sh

# Define the ports for each service
AUTH_GRPC = 50053
EMAIL_GRPC = 50054

# Target to generate protobuf files
gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	./commons/protos/*.proto

# Target to start the gateway service
run-gate:
	@echo "Starting the gateway service..."
	@go run ./gateway/main.go

# Target to start the auth service
run-auth:
	@echo "Starting the auth service..."
	@PORT=$(AUTH_GRPC) go run ./services/auth-service/main.go

# Target to start the email service
run-email:
	@echo "Starting the email services..."
	@PORT=$(EMAIL_GRPC) go run ./services/email-services/main.go

# Target to start the compose services
compose:
	@echo "Compose up"
	@docker compose up

# Target to stop the compose services
compose-down:
	@echo "Compose down..."
	@docker compose down

# Target to ensure the auth service is ready
wait-auth:
	@$(WAIT_FOR_IT) -h localhost -p $(AUTH_GRPC) -t 60 -- echo "Auth service is ready"

# Target to ensure the email service is ready
wait-email:
	@$(WAIT_FOR_IT) -h localhost -p $(EMAIL_GRPC) -t 60 -- echo "Email service is ready"

# Target to run all services in sequence
run-app:
	@echo "Starting the application"
	@make compose &
	@sleep 10
	@make run-auth &
	@make wait-auth
	@make run-email &
	@make wait-email
	@make run-gate

# Target to sync with git
git-sync:
	@git add .
	@git commit -m "Sync"

git-pull-push:
	@git pull origin main
	@git push origin main