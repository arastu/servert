FROM quay.io/brianredbeard/corebox

MAINTAINER Touhid Arastu "<touhid.arastu@gmail.com>"

COPY servert /go/bin/
ENTRYPOINT /go/bin/servert

EXPOSE 8080
