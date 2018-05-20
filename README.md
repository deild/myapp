# Myapp

## Overview

Myapp is an myapp app.

## Folder structure

A basic tree would look like:

    .
    ├── commands
    │   ├── myapp.go
    │   └── myapp_test.go
    ├── script
    │   └── setup
    ├── .gitignore
    ├── .travis.yml
    ├── CONTRIBUTING.md
    ├── Gopkg.lock
    ├── Gopkg.toml
    ├── LICENSE
    ├── README.md
    ├── goreleaser.yml
    ├── magefile.go
    └── main.go

- commands/myapp.go and myapp_test.go: is the “library” of the application and its respective files. Could be more than one file, of course;
- script: script help me start new projects faster;
- .gitignore: some standard gitignore, vendor, binary, etc;
- .travis.yml: tell Travis CI what to do;
- CONTRIBUTING.md: newcommer guide;
- Gopkg.lock and Gopkg.toml: dependencies locks and manifest;
- README.md: what you are reading;
- LICENSE: MIT;
- goreleaser.yml: the GoReleaser configuration;
- magefile.go: contains common tasks for the project, like formating, testing, linting, etc;
- main.go: is the cli entrypoint;

## Starting a new project

To use it, you can simply:

```bash
cd $GOPATH/src/github.com/youruser
git clone git@github.com/deild/myapp.git yourapp
cd myapp
./script/setup youruser YourApp # notice the case on the second arg
```

It is actually a working app (that does nothing), to run it:

```bash
mage vendor
go run main.go -h
```

Now, you create a [GitHub repository](https://help.github.com/articles/create-a-repo/) for your new app and push it:

```bash
git remote add origin https://github.com/youruser/yourapp.git
git push origin master
```

If you check the README file, you’ll see that there are already a few badges on the bottom, but some of them are not working. Let’s fix them!

## Contributing to Myapp

First of all, you can read the [CONTRIBUTING.md](CONTRIBUTING.md) file. It is the _“newcomer guide”_.

[![GitHub Release](https://img.shields.io/github/release/deild/myapp.svg)](https://github.com/deild/myapp/releases/latest)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/deild/myapp)
[![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Travis](https://img.shields.io/travis/deild/myapp.svg)](https://travis-ci.org/deild/myapp)
[![CodeFactor](https://www.codefactor.io/repository/github/deild/myapp/badge)](https://www.codefactor.io/repository/github/deild/myapp)
[![Go Report Card](https://goreportcard.com/badge/github.com/deild/myapp)](https://goreportcard.com/report/github.com/deild/myapp)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-blue.svg)](https://github.com/goreleaser)
[![SemVer](https://img.shields.io/badge/semver-2.0.0-blue.svg)](https://semver.org/)