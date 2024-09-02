FROM golang:1.23.0 AS builder

WORKDIR /build
COPY go.* .
RUN go mod download
COPY . .
RUN go build -o ./ims cmd/imsapi/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/ims ./ims

ENV DB_USER=postgres
ENV DB_PORT=5434
ENV DB_PASSWORD=password
ENV DB_HOST=host.docker.internal
ENV DB_NAME=ims
ENV JWT_KEY=ultra-super-secret

CMD ["/app/ims"]