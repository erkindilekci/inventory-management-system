FROM golang:1.23.0 AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./ims .

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/ims ./ims
CMD ["/app/ims"]