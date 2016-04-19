FROM alpine:3.3
RUN apk -U add git &&  rm -rf /var/cache/apk/*
ADD bin/goprox /bin/goprox
CMD ["/bin/goprox"]
