build:
	go build -o ./build/6502vm ./src

test:
	go test ./src/...

run:
	go run ./src/main.go

clean:
	rm build -rf
