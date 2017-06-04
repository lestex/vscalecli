FROM alpine:3.5

ENV VSCALECLI_VERSION=v0.1.0

RUN apk add --no-cache curl unzip

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

RUN curl -L https://github.com/lestex/vscalecli/archive/{VSCALECLI_VERSION}.zip | unzip

ENTRYPOINT ["./vscale"]