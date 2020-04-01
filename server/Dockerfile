################################
# Build the Source Code
################################
FROM golang:1.14.1-alpine3.11 AS build
LABEL maintainer="Nico Alimin <nicoalimin@yahoo.com>"

ENV GO111MODULE="on"
WORKDIR /server
COPY . .
RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/api

################################
# Copy the binary to final image
################################
FROM alpine:3.7

WORKDIR /server

COPY --from=build /server/bin/api /server

# Copy the Swagger docs to final image
RUN mkdir -p /server/docs
COPY --from=build /server/docs /server/docs

EXPOSE 80

CMD ["/server/api"]