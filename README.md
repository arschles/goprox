# Goprox

[![Build Status](https://travis-ci.org/arschles/goprox.svg?branch=master)](https://travis-ci.org/arschles/goprox)
[![Go Report Card](http://goreportcard.com/badge/arschles/goprox)](http://goreportcard.com/report/arschles/goprox)
[![Docker Repository on Quay](https://quay.io/repository/arschles/goprox/status "Docker Repository on Quay")](https://quay.io/repository/arschles/goprox)

Goprox is a tool for managing your own Golang dependencies. Use it to:

1. Curate dependencies for your whole organization
2. Provide a server from which anyone can read and `go get` dependencies on your `GOPATH`
3. Insulate your dependencies from changes upstream in Github/BitBucket

# Alpha
This project is alpha status. It's in early development and not all features are available.
It is available for testing, however. See below for details.

# A Proxy?

Goprox is both a server and a client. The server, `goproxd`, acts as a proxy for the 
packages under your `GOPATH` that everyone can access over the network. The client,
`goprox` is a convenient CLI that talks to `goproxd`.

For example, if someone is running a `goprox` server, you execute the following command
to `go get` [Gorilla Mux](https://godoc.org/github.com/gorilla/mux) from it:

```console
goprox get github.com/gorilla/mux HEAD
```

This command does the following:

- Tells `goprox` to package up the latest version (`HEAD`) of `$GOPATH/src/github.com/gorilla/mux` on its filesystem
- Receives the package
- Unpackages it and puts the contents into `vendor/github.com/gorilla/mux`

# Install

Since this project is in alpha, there are no downloadable binaries. To install it,
you must have the following installed:

1. Go 1.7+
2. git
3. GNU Make

If you have those dependencies installed, simply clone this repository and run:

```console
make build
```

You will then have `./cmd/cli/goprox` and `./cmd/server/goproxd` binaries available.
The former is the CLI and the latter is the server.

# Why Goprox Beats `go get`

As you know, `go get github.com/gorilla/mux` will download the _latest_ version of
the code at https://github.com/gorilla/mux and install it directly into your `$GOPATH`.

This approach is quick and easy, but presents the following problems:

1. Your codebase may not compile later if the authors of Gorilla Mux change the code in an incompatible way
2. Multiple Go projects might need different versions of Gorilla Mux, but there is only
one version in the `$GOPATH`

The first problem is solved by dependency management tools like 
[Glide](https://github.com/Masterminds/glide)
(the tool that this project uses) and [Godep](https://github.com/tools/godep).

The second problem is solved by the
[`vendor` folder](https://blog.gopheracademy.com/advent-2015/vendor-folder/). 
If you're unfamiliar with how the `vendor` folder works, here is a brief overview:

- The `go` toolchain will look for a directory called `vendor` in the current directory as well
as the directories above the current directory
- The toolchain will look for dependencies in the `vendor` directory _first_, before
looking in the `$GOPATH`
- That fact means that you can "freeze" a dependency package's version for just your
project, without changing any other part of the `$GOPATH`
- Your project still has to be in the `$GOPATH`
- In most cases, Go projects put a single vendor directory at the project root
- The aforementioned tools (Glide and Godep) -- and others -- can automatically manage
your `vendor` directory for you

So, we already have tools to solve both of the above problems, so why do we need another?

# Why Goprox Beats Glide, Godep and Other Tools

There is no central package repository for Go dependencies, nor is there a "package format" 
like [Java's JAR](https://docs.oracle.com/javase/8/docs/technotes/guides/jar/jarGuide.html)
or [Rust's Crates](http://doc.crates.io/guide.html).

So, in order for these tools to fetch a dependency for use in your project, they have to
go to the source control repository -- Github in most cases -- and download the code.

That means that the released packages -- the code you depend on -- lives in the same place
as in-development code. That's a fundamental problem, because development, or "bleeding edge"
code is meant to change over time, but your dependencies can't. If they did, your build
will become unstable and your progress will come to a grinding halt.

Let's zoom in on some specific examples of this problem:

1. You depend on version `e406d3a` of a package on GitHub. Three weeks later, the 
package's developer does a big `git rebase` and squashes that commit into another. 
Since your dependency tools can't download the package at that version anymore, your builds
break. Your only option is to somehow determine which new version is most appropriate to use
2. You depend on a package in GitHub, and three weeks later it is deleted. Your build is broken
and you have to hope that someone still has a copy of the package code
3. You depended on a package in Google Code, then 
[Google Code shut down](https://www.engadget.com/2015/03/13/google-code-closing/). Just as
with the previous example, your build is broken and you have to hope that someone still has a
copy of the package code. See 
[here](https://www.reddit.com/r/golang/comments/42r1j7/codegooglecom_is_down_all_packages_hosted_there) 
for more details on how that shutdown affect Go projects specifically

# Why Not Check In Your `vendor` directory?

If you're looking to fix the issues above, you could just check in your vendor directory!
Nobody who builds your code will need to download any dependency management tools (Glide, Godep,
etc...), nor will you risk your build breaking _because you will have no external 
dependencies_. Every dependency you need will be in your repository forever.

However, you'll eventually need to upgrade the code in your vendor directory, and then you're
faced once again with the problems listed in the previous section.

Additionally, if you are writing a library yourself, or it's possible that someone might
`import` your library, they themselves will run into problems importing and using _your_
package. See [this thread](https://groups.google.com/forum/#!msg/golang-nuts/AnMr9NL6dtc/UnyUUKcMCAAJ) for more 
detail on why this is the case.

# REMAINDER STUFF - MODIFY OR REMOVE

(note that even in other languages these problems exist. The node.js community had a big [issue](http://blog.npmjs.org/post/141577284765/kik-left-pad-and-npm) when someone deleted a small package that was very "deep" in many large dependency trees)

## Example Scenario

Let's imagine a scenario that shows how relying on third party repositories can be harmful to your project, and then show how GoProx can help.

Say your project depends on a (fictional) package of mine called `github.com/arschles/gofictional`. You use [glide](https://github.com/Masterminds/glide), so you get repeatable, deterministic builds because you've pinned your version of that package to a known Git SHA.

One day, I update my package, squash it into the commit SHA that you depend on and push the update to my `master` branch. This means that next time you pull dependencies (for example, in your CI system), your project will build with my new update.

At best, your code won't build, but at worst your tests won't catch the change and you'll ship fundamentally different software without knowing it.

If you change the dependency's source to use GoProx, you'll be effectively pulling the code from your S3 bucket, not from GitHub. Even though I squashed commits into the original SHA, the code GoProx will serve won't have it. If you decide to use the updated code, you can tell GoProx to update its cache.
