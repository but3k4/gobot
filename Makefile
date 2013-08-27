all: install

install:
	go get -v github.com/but3k4/gobot
	go build -o bin/gobot bot.go


clean:
	rm bin/*
