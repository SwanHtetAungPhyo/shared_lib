gen:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./commons/protos/*.proto

run-gate:
	@go run ./gateway/main.go

run-auth:
	@go run ./services/auth-service/main.go

git-sync: 
	@git add . 
	@git commit -m "Sync"