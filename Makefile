build:
	go build -o todo ./cmd/todo

run:
	go run ./cmd/todo

clean:
	rm -f todo
