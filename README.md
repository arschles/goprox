# goprox

[![Build Status](https://travis-ci.org/arschles/goprox.svg?branch=master)](https://travis-ci.org/arschles/goprox)
[![Go Report Card](http://goreportcard.com/badge/arsches/goprox)](http://goreportcard.com/report/arschles/goprox)
[![Docker Repository on Quay](https://quay.io/repository/arschles/goprox/status "Docker Repository on Quay")](https://quay.io/repository/arschles/goprox)

A proxy server for Go packages, for you to run yourself

# Alpha
This project is alpha status. It's in early development but is ready for testing. See below for more information

# Go Package Management Today

Go package management is fairly sophisticated on the client side. Tools like [godep](https://github.com/tools/godep) and [glide](https://github.com/Masterminds/glide) (to name only 2 popular ones) are very good at capturing your dependencies, traversing dependency trees, locking versions and otherwise generally ensuring that your builds are deterministic. If you use these tools to manage your dependencies, you can write your code with the "soft" guarantee that the code underneath will not change regardless of where or when it's built.

## The Catch

I called it a "soft" guarantee above because dependency management tools, along with `go get` itself, downloads your dependency code directly from hosted source code management systems like GitHub. This strategy has two major shortcomings:

1. If the hosted system is unavailable for any reason, your build will fail because you won't be able to download your dependency
2. If the author of a library modifies or deletes a tag or modifies or deletes a commit, the code you depended on will change or disappear

As you'd expect, the first shortcoming can result in intermittent failures and can be annoying. It happened on a large scale when [Google Code shut down](https://www.reddit.com/r/golang/comments/42r1j7/codegooglecom_is_down_all_packages_hosted_there/).

The second shortcoming, however, can create much more subtle failures. If someone changes a commit, your code might not compile. Even worse, it might compile but work differently.

(Author's note: problem #1 is not limited to the Go community. [NPM](https://www.npmjs.com/) had a high-profile similar [issue](http://blog.npmjs.org/post/141577284765/kik-left-pad-and-npm)).

# How Goprox Fixes These Problems

There are three main ways to solve the problems listed above:

1. Trust the maintainers of the code you depend on to be reasonable (don't change tags or commits, don't delete their repository, etc...)
2. Rely on a trusted third party that stores dependency code for you, provides uptime guarantees, tells you when/how it will change, etc...
3. Store the code yourself somewhere

Goprox is software that lets you do #3.

Specifically, it's a server that:

- Serves `go get` requests
- Redirects those requests to its built-in git server (via the `<meta>` tag approach - see `go help importpath` for details on how this works)
- When a request comes into the git server, `goprox` downloads the package code from a specified S3 bucket and responds with the requested Git Sha.
  - If the package code doesn't exist, it fills the S3 bucket first and then responds with the code

# OLD:
# Why?

If you use a Go package management tool (like [glide](https://github.com/Masterminds/glide)) that can download your dependencies for you (in glide's case, this is with a `glide install`), then you rely on the repositories that host your packages to always be available.

Recent (as of this writing) events have proven that you can't rely on repositories to always be available. [Google Code has shut down](http://google-opensource.blogspot.com/2015/03/farewell-to-google-code.html), and BitBucket and Github have outages sometimes, for example.

If you have many developers or robots (such as CI tools) working on your codebase at a given time, an outage can slow or halt development because the developers or robots can't download the dependencies they need to build your codebase. Sometimes, in these cases, you have to wait until the repository is back up. Other times, you have to switch dependencies.

`goprox` is a server that you run which serves Go packages from _your_ S3 bucket. It is not a centralized Go package repository - you run the server yourself, and you store the code that it serves.
