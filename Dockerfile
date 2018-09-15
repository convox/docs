FROM golang:1.11

RUN curl -Ls https://github.com/mattgreen/watchexec/releases/download/1.8.6/watchexec-1.8.6-x86_64-unknown-linux-gnu.tar.gz | \
    tar -C /usr/bin --strip-components 1 -xz

WORKDIR /go/src/github.com/convox/docs
COPY . /go/src/github.com/convox/docs
RUN go install .

CMD ["bin/web"]
