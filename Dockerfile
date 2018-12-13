FROM golang:1.11

# Go envs
ENV GOPATH /go/
ENV APP_HOME $GOPATH/src/github.com/vavarodrigues/api/
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

# Go Deps
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Project Deps
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

COPY . $APP_HOME

EXPOSE 8080

CMD ["api"]