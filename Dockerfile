FROM alpine:3.5

RUN apk add --update bash curl && rm -rf /var/cache/apk/*

ADD keystore /bin/keystore

EXPOSE 80

RUN mkdir keys

CMD ["/bin/keystore"]
