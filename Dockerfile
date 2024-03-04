FROM golang:1.22.0-alpine3.19 AS builder

WORKDIR /app

RUN apk add bash
RUN apk add gcc libc-dev

ENV CGO_ENABLED 1

ADD . .

RUN bash -c assets/download.sh
RUN go build -o knight .


FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/knight .
COPY --from=builder /app/assets/ assets/
COPY --from=builder /app/templates/ templates/

ENTRYPOINT [ "./knight" ]
