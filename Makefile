# ------ Golang Command ------ #
run-os:
	GO_ENV=development go run main.go stats.go

run-win:
	set GO_ENV=development
	go run main.go stats.go

i:
	go mod tidy

format:
	gofmt -l .

swag:
	export PATH=$PATH:$HOME/go/bin

swagi:
	swag init