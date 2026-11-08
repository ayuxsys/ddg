default: build

build:
	go build -o apihunter ./

install:
	go install ./