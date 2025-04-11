FROM golang:alpine
WORKDIR /build
COPY . .
RUN go build -o repos cmd/main.go
CMD [ "./repos" ]
