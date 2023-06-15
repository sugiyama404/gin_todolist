# For development phase
FROM golang:1.19-buster as development
WORKDIR /app

ENV HOST dbserver
ENV DBNAME todoproject
ENV USERNAME goMySql1
ENV PASSWORD goMySql1

COPY ./app /app

# build phase
FROM golang:1.19-buster as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./app .

RUN go build -o main ./cmd/...

# deploy phase
FROM gcr.io/distroless/static-debian11:nonroot
WORKDIR /

COPY --from=build /app/main /main
USER nonroot
EXPOSE 8080

CMD [ "/main" ]
