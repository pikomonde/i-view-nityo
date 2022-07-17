run:
	go run cmd/http/app.go
test:
	mockgen -source=repository/repository_interface.go -package=repository > repository/mock_repository.go
	mockgen -source=service/service_interface.go -package=service > service/mock_service.go

	# go test ./... -coverprofile cover.out && go tool cover -func cover.out
	go test ./... -coverprofile cover.out.temp && cat cover.out.temp | grep -v 'mock_*\|\.pb\.' > cover.out && go tool cover -func cover.out
