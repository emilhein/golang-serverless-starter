[![Go Report Card](https://goreportcard.com/badge/github.com/emilhein/golang-serverless-starter)](https://goreportcard.com/report/github.com/emilhein/golang-serverless-starter)
[![Build Status](https://travis-ci.org/emilhein/golang-serverless-starter.svg?branch=master)](https://travis-ci.org/emilhein/golang-serverless-starter)

# golang-starter

A starter project for playing around with Golang.

### Serverless function

This project shows how you can setup automatic deployment to AWS using the serverless framework and Travis CI

The serverless function simply echos back what it recieves

## Travis setup

To have travis deploy on your AWS account when you commit ot the master branch remember to set the _AWS_ACCESS_KEY_ID_ & _AWS_SECRET_ACCESS_KEY_ environment variables in Travis
