FROM golang:alpine

RUN apk add --update git curl && \
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    
WORKDIR /go/src/github.com/vavaRodrigues/api

COPY . ./

RUN dep ensure && \
    go install

EXPOSE 8080

CMD ["api"]
