FROM golang:1.9
MAINTAINER Lei Tian

WORKDIR /go/src

RUN \
    mkdir -p github.com/imSQL && \
    go get -u github.com/go-sql-driver/mysql && \
    go get -u github.com/gin-gonic/gin && \
    go get -u github.com/imSQL/proxysql && \
    git clone https://github.com/imSQL/proxysql_master.git github.com/imSQL/proxysql_master && \
    cd /go/src/github.com/imSQL && \
    go install ./proxysql_master && \
    rm -fr /go/src/github.com/imSQL/proxysql_master

EXPOSE 3333


ENTRYPOINT ["/go/bin/proxysql_master"]
