FROM golang:alpine AS build
WORKDIR /go/src/todo
COPY . .
WORKDIR /go/src/todo/src
RUN go build -o main

FROM alpine:latest
COPY --from=build /go/src/todo /todo
ENTRYPOINT /todo
COPY ./public ./public
EXPOSE 3000
CMD ["./main"]