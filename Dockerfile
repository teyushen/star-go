FROM golang:1.8

WORKDIR /go/src/github.com/teyushen/star-go
COPY . .

RUN mkdir .star-go && rm -rf .git .gitignore .idea .DS_Store

RUN go get -d -v github.com/teyushen/star-go
RUN go install -v github.com/teyushen/star-go


ENTRYPOINT ["star-go"]
#CMD ["star-go"]
