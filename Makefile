gorl: cmd/**/*.go
	@mkdir -p build
	go build -o build/gorl cmd/main.go

run: main.go build/gorl
	./build/gorl
