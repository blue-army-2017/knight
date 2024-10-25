FROM golang:1.23 AS builder

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN sqlc generate
RUN GOOS=linux go build -o /knight


FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /

COPY --from=builder /knight /knight
COPY assets/ assets/
COPY view/ view/
COPY schema/ schema/

ENV GIN_MODE=release

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/knight" ]
