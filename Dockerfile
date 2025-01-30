## development #################################################################

FROM golang:1.22 AS development

RUN apt-get update && apt-get -y install curl software-properties-common && curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get update && apt-get -y install nodejs

RUN curl -Ls https://github.com/mattgreen/watchexec/releases/download/1.8.6/watchexec-1.8.6-x86_64-unknown-linux-gnu.tar.gz | \
    tar -C /usr/bin --strip-components 1 -xz

ENV DEVELOPMENT=true

WORKDIR /go/src/github.com/convox/docs/webpack
COPY webpack/node_modules ./node_modules
RUN npm rebuild

WORKDIR /go/src/github.com/convox/docs
COPY . .
RUN make build

CMD ["bin/web"]

## package #####################################################################

FROM golang:1.22 AS package

RUN apt-get update && apt-get -y install curl software-properties-common && curl -sL https://deb.nodesource.com/setup_14.x | bash -
RUN apt-get update && apt-get -y install nodejs

RUN go install github.com/gobuffalo/packr@latest

ENV PACKAGE=true

WORKDIR /go/src/github.com/convox/docs

COPY --from=development /go/src/github.com/convox/docs .
RUN make -B build

## production ##################################################################

FROM ubuntu:18.04 AS production

RUN apt-get update && apt-get -y install ca-certificates

ENV DEVELOPMENT=false
ENV GOPATH=/go

WORKDIR /

COPY bin/web /bin/

COPY --from=package /go/bin/web /go/bin/web

RUN groupadd -r docs && useradd -r -g docs docs
USER docs

CMD ["/bin/web"]
