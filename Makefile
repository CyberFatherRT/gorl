build/gorl: main.go
	@mkdir -p build
	go build -o $@ $^

run: main.go build/gorl
	./build/gorl
