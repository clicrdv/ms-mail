FROM golang:1.8

COPY . /go/src/ms-mail
WORKDIR /go/src/ms-mail
EXPOSE 50052
RUN make get
RUN make binary
CMD ["ms-mail"]
