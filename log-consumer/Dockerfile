FROM golang:1.18 AS base

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN ["go", "mod", "download"]

COPY main.go .

RUN ["go", "build", "-o", "/consumer", "main.go"]

FROM golang:1.18

WORKDIR consumer

COPY --from=base /consumer .

CMD ["./consumer"]
