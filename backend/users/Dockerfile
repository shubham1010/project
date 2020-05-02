FROM golang:latest as dev
MAINTAINER shubham jagdhane

ADD . /go/src/github.com/shubham1010/users

WORKDIR /go/src/github.com/shubham1010/users

RUN CGO_ENABLED=0 GOOD=linux go build -a -installsuffix cgo -o main .


FROM alpine:3.9.2 as prod
MAINTAINER shubham jagdhane

WORKDIR /root/

#COPY --from=dev /go/src/github.com/shubham1010/users/containerData /root/containerData

#COPY --from=dev /go/src/github.com/shubham1010/users/main containerData/main
COPY --from=dev /go/src/github.com/shubham1010/users .

EXPOSE 8080

ENTRYPOINT ["./main"]
