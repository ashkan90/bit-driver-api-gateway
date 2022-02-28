FROM golang:1.17-alpine as build

WORKDIR /app

VOLUME ["/app"]

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux CGO_ENABLED=0 go build -o /bit-driver-api-gateway ./cmd/

FROM alpine
COPY --from=build /bit-driver-api-gateway ./app

COPY ./services.yaml ./
EXPOSE 8080

CMD [ "./app", "--proxy-services", "services.yaml" ]