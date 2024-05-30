build/gorl: pkg/**/*.go main.go
	@mkdir -p build
	go build -o $@ main.go

.PHONY: run
run: ./build/gorl
	$^
