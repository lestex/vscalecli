FROM alpine:3.5

ENV VSCALECLI_VERSION=0.1

RUN apk add --no-cache curl bash file

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app

RUN curl -L https://github.com/lestex/vscalecli/releases/download/v${VSCALECLI_VERSION}/vscale -o vscale
RUN chmod +x vscale

ENTRYPOINT ["./vscale"]
