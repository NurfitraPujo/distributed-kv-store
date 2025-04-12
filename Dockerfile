FROM golang:1.23.5 AS build

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o kvs

FROM scratch

COPY --from=build /src/kvs .

EXPOSE 8080

CMD ["/kvs"]