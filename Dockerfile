FROM golang:1.21.0

WORKDIR /app
COPY . .

EXPOSE 8080
RUN go build -o main .

CMD ["/app/main"]
