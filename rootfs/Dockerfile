FROM alpine:3.3
ENV WEB_BIND_PORT 8080
ENV GIT_PORT $WEB_BIND_PORT
EXPOSE $WEB_BIND_PORT
RUN apk -U add git &&  rm -rf /var/cache/apk/*
ADD bin/goprox /bin/goprox
CMD ["/bin/goprox"]
