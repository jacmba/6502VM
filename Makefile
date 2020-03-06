.PHONY: build test tdd run clean

build:
	go build -o ./build/6502vm ./src

test:
	go test -v ./src/...

tdd:
	looper src

run:
	go run ./src/main.go

clean:
	rm build -rf
