FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache just git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN just build

FROM gcr.io/distroless/static-debian13
WORKDIR /

COPY --from=builder /app/bin/{{ cookiecutter.bin_name }} /{{ cookiecutter.bin_name }}

USER nonroot:nonroot

ENTRYPOINT [ "/{{ cookiecutter.bin_name }}" ]
