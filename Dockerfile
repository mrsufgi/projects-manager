# Builder
FROM golang:1.15-alpine3.12 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

# Build 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

# Runtime
FROM scratch

COPY --chown=0:0 --from=builder /app/build/projects /

ENTRYPOINT ["/projects"]