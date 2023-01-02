#BUILD
FROM golang:1.19 as build

# Environment variables
ENV PORT=3030
ENV APP_COUCHBASE_HOST=localhost
ENV APP_COUCHBASE_USER=Administrator
ENV APP_COUCHBASE_PASSWORD=couchbase
ENV APP_COUCHBASE_BUCKET=fintech

COPY . /app

WORKDIR /app

RUN cd cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main


CMD /main

# TEST
FROM build as test

# PRODUCTION
FROM alpine:latest as production

# Environment variables
ENV PORT=3030
ENV APP_COUCHBASE_HOST=localhost
ENV APP_COUCHBASE_USER=Administrator
ENV APP_COUCHBASE_PASSWORD=couchbase
ENV APP_COUCHBASE_BUCKET=fintech

RUN apk --no-cache add ca-certificates

COPY --from=build ./main ./

RUN chmod +x ./main

EXPOSE ${PORT}

ENTRYPOINT ["./main"]
