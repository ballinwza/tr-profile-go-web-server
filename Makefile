# ------ Golang Command ------ #
run-os:
	GO_ENV=development go run main.go rooms.go routes.go stats.go

run-win:
	set GO_ENV=development && go run main.go rooms.go routes.go stats.go

i:
	go mod tidy

format:
	gofmt -l .
