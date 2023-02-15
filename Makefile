
run:
	go run .

build:
	go build .
	sudo cp ./yo /usr/local/bin

.PHONY: run build
.SILENT: run build 