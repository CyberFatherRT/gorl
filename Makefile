gorl: pkg/**/*.go main.go
	@mkdir -p build
	go build -o build/gorl main.go

run: main.go build/gorl
	./build/gorl
