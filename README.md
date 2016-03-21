# goprox
A proxy server for Go packages, for you to run yourself

# Work In Progress
This project is under (slow) development right now. It won't work as intended for now.

# Why?
There are two reasons why you might want a proxy server:

1. You don't want to check in your `vendor` directory, but want a fast, reliable package repository.
2. You want a documentation server for a specific version of a package

This project aims first to solve (1). The most important goal is reliability. Recently (as of this writing), `code.google.com` went offline for many Go packages. `go get` no longer works on them.

If you can run a proxy yourself that holds a cache of the source, you can still get the code, even if the target of `go get` (i.e. Github, code.google.com, etc...) is down. Finally, many Go dependency tools (like [Glide](https://github.com/Masterminds/glide)) allow you to specify an alternate location for a package, effectively allowing you to alias a canonical package name to your proxy.
