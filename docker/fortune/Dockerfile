FROM alpine:3.10.2

RUN apk add fortune
ADD fortuneloop.sh /bin/fortuneloop.sh

ENTRYPOINT sh /bin/fortuneloop.sh

