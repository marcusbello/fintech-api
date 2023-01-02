#BUILD
FROM golang:latest as build

COPY . /service

WORKDIR /service

RUN ls

RUN cd /service/cmd && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /http-service

CMD /http-service

# TEST
FROM build as test

# PRODUCTION
FROM alpine:latest as production

RUN apk --no-cache add ca-certificates

COPY --from=build /http-service ./

RUN chmod +x ./http-service

EXPOSE 3030

ENTRYPOINT ["./http-service"]
