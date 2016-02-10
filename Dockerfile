FROM alpine:3.3
RUN apk -U add git
ADD goprox /bin/goprox
CMD ["/bin/goprox"]
