FROM golang:1.19.1-buster
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY internal/*.go ./internal/
COPY services/*.go ./services/
COPY *.go ./

RUN go build -o ./go-calculator

EXPOSE 8080

CMD [ "./go-calculator" ]
