FROM golang:1.21.4
WORKDIR /aoc2023
COPY . .
RUN go mod download && \
	go build -v -o exercise1
ENV FILE files/input.txt
ENTRYPOINT ["/aoc2023/exercise1"]