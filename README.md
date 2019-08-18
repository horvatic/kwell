# Vaticwella
[![Build Status](https://travis-ci.org/horvatic/vaticwella.svg?branch=master)](https://travis-ci.org/horvatic/vaticwella) 
[![Coverage Status](https://coveralls.io/repos/github/horvatic/vaticwella/badge.svg?branch=master)](https://coveralls.io/github/horvatic/vaticwella?branch=master)


## Overview
A service to process images using image processing techniques, and return the edited image.
Service has: 
	Super Sampling,	Median filter, binary filter

To build: (go to main directory) go build -o vaticwella main.go

To test: (top level directory) go test ./...

## Docker
Dockerhub image:
https://hub.docker.com/r/horvatic/vaticwella

To build a new docker image:
(go to main directory) docker build -t vaticwella .

To run image:
docker run --publish 8080:8080 --name vaticwella --rm vaticwella

## Running The Service

host/{filter}?sa={searchArea}&runInParallel={ToRunInParallel}

Post an image to localhost:8080

ss = super sampling

mf = mean filter

bf = binary filter

sa (int) = search area

runInParallel (bool) = Using parallel processing

supported image formats: gif, jpeg, png

Example: 

localhost:8080/mf?sa=10&runInParallel=false

localhost:8080/bf?sa=10&runInParallel=true

