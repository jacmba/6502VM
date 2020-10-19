.PHONY: build test tdd run clean

build:
	go build -o ./build/6502vm ./

test:
	go test -v ./...

tdd:
	looper ./

run:
	go run main.go

clean:
	rm build -rf
