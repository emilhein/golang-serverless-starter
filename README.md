[![Go Report Card](https://goreportcard.com/badge/github.com/emilhein/golang-starter)](https://goreportcard.com/report/github.com/emilhein/golang-starter)
[![Build Status](https://travis-ci.org/emilhein/golang-starter.svg?branch=master)](https://travis-ci.org/emilhein/golang-starter)

# golang-starter

A starter project for playing around with Golang.

This project shows how you can setup automatic deployment to AWS using the serverless framework and Travis CI

The serverless function simply echos back what it recieves

## Get started

```
go build
.\golang-starter.exe

```

For local development i would use [gin](https://github.com/gin-gonic/gin) and run:

```
cd $GOPATH/src/golang-starter/src/github.com/emilhein/webserver
gin run main.go
```
