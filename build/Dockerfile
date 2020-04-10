# Multi-stage build

#First stage for bulding binary executable
FROM golang:1.14.1 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN make build

#Second stage only to run prepared binary executable
COPY --from=builder /app/main /
COPY --from=builder /app/configs /

EXPOSE 8080

ENTRYPOINT ["./main", "-config-path", "apiserver.toml"]
