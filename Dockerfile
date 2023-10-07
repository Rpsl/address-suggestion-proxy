##
## Build
##
FROM golang:1.21-bullseye AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod verify

RUN export CGO_ENABLED=0 && go build -ldflags "-s -w" -o /address-suggestion-proxy ./main.go

##
## Production
##
FROM golang:1.21-alpine

WORKDIR /app/

COPY --from=build /address-suggestion-proxy /app/address-suggestion-proxy

ENTRYPOINT ["/app/address-suggestion-proxy"]