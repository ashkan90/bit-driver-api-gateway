FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go test ./... -v

RUN go build -o ./bit-driver-api-gateway ./cmd/

EXPOSE 8080

CMD [ "./bit-driver-api-gateway" ]