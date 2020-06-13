FROM goland:1.14

ENV GO111MODULE=on

WORKDIR /go/src/github.com/alitaso345/zatsu/grpc-web-sample2/server
COPY . .
RUN go get github.com/pilu/fresh
CMD ["fresh"]
EXPOSE 9090