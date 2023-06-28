build:
	go build -o bin/poker

run:
	./bin/poker

test:
	go test -v ./...

install telnet:
	brew install telnet

tel:
	telnet localhost 999