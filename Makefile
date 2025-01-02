# ------ Golang Command ------ #
run:
	export GO_ENV=dev
	go run main.go rooms.go routes.go stats.go

i:
	go mod tidy

format:
	gofmt -l .
