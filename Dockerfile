FROM alpine:3.5

RUN apk add --update bash curl && rm -rf /var/cache/apk/*

ADD Keystore /bin/Keystore

EXPOSE 8080

RUN mkdir keys

CMD ["/bin/Keystore"]