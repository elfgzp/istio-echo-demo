FROM golang:1.14-alpine AS builder
RUN apk add --no-cache curl jq git build-base
WORKDIR /opt
RUN cd /opt
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM godleon/nettools:latest
LABEL maintainer="Elf Gzp <gzp@741424975@gmail.com> (https://elfgzp.cn)"
COPY --from=builder /opt/app ./
RUN chmod +x /app
EXPOSE 8088
CMD ["/app"]