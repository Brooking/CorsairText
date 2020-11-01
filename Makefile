init:

gen: init
	go get github.com/golang/mock/mockgen
	go generate ./...

run: init
	go run main.go

test: init
	go test -cover -race ./...

build: init
	go build -o bin/corsairtext

debug: init
	go build -gcflags='all=-N -l' -o bin/corsairtextdebug