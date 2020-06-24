FROM golang
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM alpine:latest
COPY --from=0 /app/server /server
ENV PORT 8080
CMD ["/server"]