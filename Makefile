gorl: main.go
	@mkdir -p build
	go build -o build/gorl main.go

run: main.go
	go run main.go
