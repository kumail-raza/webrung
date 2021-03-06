FROM golang:1.13-alpine

RUN set -eux; \
	apk add --no-cache --virtual .build-deps \
		gcc \
		musl-dev \
	;
RUN mkdir -p /go/src/github.com/minhajuddinkhan/webrung/
WORKDIR /go/src/github.com/minhajuddinkhan/webrung/
COPY . .

ENTRYPOINT ["/go/src/github.com/minhajuddinkhan/webrung/docker-entrypoint.sh"]
