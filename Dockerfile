FROM alpine:3.5

ENV VSCALECLI_VERSION=v0.1.0

RUN apk add --no-cache curl unzip

WORKDIR /app

RUN curl -L https://github.com/lestex/vscalecli/archive/{VSCALECLI_VERSION}.zip | unzip

ENTRYPOINT ["./vscale"]
