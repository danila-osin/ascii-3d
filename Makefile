
build:
	go build -o main.out cmd/main.go

run:
	go run cmd/main.go -w $(w) -h $(h) -fr $(fr) -fa $(fa) -m $(m)

.PHONY: build run