# image with development tools
FROM xena/go:1.10
ENV GOPATH /root/go
RUN apk --no-cache add git protobuf retool make
COPY . /root/go/src/github.com/horseville/horseville
WORKDIR /root/go/src/github.com/horseville/horseville
RUN make build

# runner image
FROM xena/alpine
COPY --from=0 /root/go/src/github.com/horseville/horseville/bin/ /usr/local/bin/
CMD /usr/local/bin/horsevilled
