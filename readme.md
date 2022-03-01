<h1 align="center">Bi Taksi Driver API Gateway</h1>

![example workflow](https://github.com/ashkan90/bit-driver-api-gateway/actions/workflows/main.yml/badge.svg)

> The Gateway project takes entire traffic onto it. It handles some basic logics such as Authorization then proxies the incoming request to related service. 

## Introduction

The Gateway project does not implement complete JWT authorization. It has super-simple logic to handle authorization. <br>
Gateway only checks your 'Authorization' header and service waiting it to be 'authenticated: true' otherwise it will block customer request. <br>
Gateway does not implement Application Load Balancer. It leaves Load balancer things to server/cloud provider.  

## Test purpose
To test the gateway, you should run 'driver-location-service' and 'driver-matching-service' firstly;<br>
[driver-location-service](https://github.com/ashkan90/bit-driver-location-service) <br>
[driver-match-service](https://github.com/ashkan90/bit-driver-matching-service)

then start the gateway to check everything is ok.
```shell
docker build -t bit-driver-api-gateway .
docker run -p 4050:8080 bit-driver-api-gateway
```

## Endpoints

```console
GET /match-svc/find-nearest/
GET /location-svc/nearest-driver-location/
```

## Example Usages
```shell
curl --location --request GET 'http://localhost:4050/match-svc/find-nearest/' \
--header 'Authorization: Bearer authenticated: true' \
--header 'Content-Type: application/json' \
--data-raw '{
    "longitude": 40.94289771,
    "latitude": 29.0390297
}'
```

# TODO
- [x] Unit test
- [ ] Circuit-breaker
