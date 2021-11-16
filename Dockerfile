FROM golang:1.17
LABEL maintainer="fouraminc@gmail.com"
WORKDIR /build
COPY . /build
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build beer-data-service .

FROM alpine:latest
WORKDIR /app
COPY --from=0 /build/beer-data-service .
EXPOSE 8080
ENTRYPOINT ["./beer-data-service"]
