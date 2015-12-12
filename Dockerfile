FROM golang:1.5
RUN go get github.com/tools/godep
WORKDIR /go/src/github.com/joshrendek/ghforkaudit
CMD ["make", "build"]
