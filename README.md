# vaticwella
[![Build Status](https://travis-ci.org/horvatic/vaticwella.svg?branch=master)](https://travis-ci.org/horvatic/vaticwella) 
[![Coverage Status](https://coveralls.io/repos/github/horvatic/vaticwella/badge.svg?branch=master)](https://coveralls.io/github/horvatic/vaticwella?branch=master)

A service to process images using image processing techniques, and return the edited image.

To build: (go to main directory) go build -o vaticwella main.go

To test: (top level directory) go test ./...

Dockerhub image:
https://hub.docker.com/r/horvatic/vaticwella

To build a new docker image:
(go to main directory) docker build -t vaticwella .

To run image:
docker run --publish 8080:8080 --name vaticwella --rm vaticwella

Service has: 
	Super Sampling,	Median filter, binary filter

Running the Service:

Post an image to localhost:8080

ss = super sampling

mf = mean filter

bf = binary filter

sa = search area

supported image formats: gif, jpeg, png

Example: 

localhost:8080/ss?sa=10

localhost:8080/mf?sa=1

