# goprox
A proxy server for Go packages, for you to run yourself

# Work In Progress
This project is under development right now. It does not work as intended until further notice.

# What

`goprox` is a proxy server that:

- Serves `go get` requests
- Redirects those requests to its built-in git server (via the `<meta>` tag approach - see `go help importpath` for details on how this works)
- When a request comes into the git server, `goprox` downloads the package code from a specified S3 bucket and responds with the requested Git Sha.
  - If the package code doesn't exist, it fills the S3 bucket first and then responds with the code

# Why?

If you use a Go package management tool (like [glide](https://github.com/Masterminds/glide)) that can download your dependencies for you (in glide's case, this is with a `glide install`), then you rely on the repositories that host your packages to always be available.

Recent (as of this writing) events have proven that you can't rely on repositories to always be available. [Google Code has shut down](http://google-opensource.blogspot.com/2015/03/farewell-to-google-code.html), and BitBucket and Github have outages sometimes, for example.

If you have many developers or robots (such as CI tools) working on your codebase at a given time, eventually you will hit an outage of some kind, you won't be able to download one or more dependencies, and you'll hit a false negative build failure.

`goprox` serves as a cache that is backed by an S3 bucket, which is highly available and controlled by you and only you.
