all: docker

docker:
	docker build --tag="joshrendek/ghforkaudit" .
	docker run -v `pwd`:/go/src/github.com/joshrendek/ghforkaudit joshrendek/ghforkaudit

build:
	godep go build -o ghforkaudit_linux
