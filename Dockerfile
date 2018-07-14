FROM golang:1.8

RUN mkdir -p /data

WORKDIR /go/src/github.com/teyushen/star-go
COPY . .

RUN go get -d -v github.com/teyushen/star-go
RUN go install -v github.com/teyushen/star-go

WORKDIR /data

ENTRYPOINT ["star-go"]
#CMD ["star-go"]
