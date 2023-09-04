FROM golang:bullseye

WORKDIR /app

CMD ["./myapp"]


# TODO: Add a multi-stage build to reduce the size of the image
#
# FROM golang:bullseye as BUILD
# WORKDIR /app
# RUN GOARCH=amd64 GOOS=linux go build -o myapp ./cmd/server_http
# FROM alpine:edge as alpine
# WORKDIR /app
# COPY --from=BUILD /app/myapp .
# RUN chmod +x /app/myapp
# CMD ["./myapp"]
