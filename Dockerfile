FROM golang:1.11.5-stretch
RUN go version

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ENV GOPATH=/home
ENV APPPATH=/home/src/unio-logger

WORKDIR ${APPPATH}
COPY ./ ${APPPATH}/
COPY ./Gopkg.toml ${APPPATH}/Gopkg.toml
RUN dep ensure

EXPOSE 8080