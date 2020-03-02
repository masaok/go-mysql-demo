APP := demo

all: clean build

build: $(APP).go
	go build $(APP).go

clean:
	rm $(APP)

hello:
	go run hello.go

run: $(APP).go
	go run $(APP).go
