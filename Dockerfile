FROM alpine:3.5

ENV VSCALECLI_VERSION=v0.1

RUN apk add --no-cache curl unzip

WORKDIR /app

RUN curl -L https://github.com/lestex/vscalecli/releases/download/{VSCALECLI_VERSION}/vscalecli

ENTRYPOINT ["./vscale"]
