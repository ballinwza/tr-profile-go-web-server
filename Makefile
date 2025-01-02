# ------ Golang Command ------ #
run:
	go run main.go rooms.go routes.go stats.go

i:
	go mod tidy

format:
	gofmt -l .
