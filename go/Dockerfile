FROM golang:1.13-alpine as builder

COPY . /go
WORKDIR /go/src/sqrt

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sqrt .

FROM scratch

COPY --from=builder /go/src/sqrt/sqrt /

EXPOSE 8000
CMD ["/sqrt"]