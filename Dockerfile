FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux go build -o /knight


FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /

COPY --from=builder /knight /knight
COPY assets/ assets/
COPY view/ view/

ENV GIN_MODE=release

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/knight" ]
