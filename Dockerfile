FROM golang:latest
WORKDIR /go/src/github.com/L1ghtman2k/ScoreTrak
COPY pkg/ pkg/
COPY cmd/ cmd/
COPY go.mod go.mod
COPY go.sum go.sum
COPY test/ test/
RUN go mod tidy
RUN go build -o master cmd/master/main.go
RUN go build -o worker cmd/worker/main.go
RUN chmod +x master worker

#Set Context Path as ScoreTrak directory