# vaticwella
[![Build Status](https://travis-ci.org/horvatic/vaticwella.svg?branch=master)](https://travis-ci.org/horvatic/vaticwella) [![Coverage Status](https://coveralls.io/repos/github/horvatic/vaticwella/badge.svg?branch=master)](https://coveralls.io/github/horvatic/vaticwella?branch=master)

A service to process images using image processing techniques, and return the edited image.

To build: (main directory) go build -o vaticwella main.go

To test: (top level directory) go test ./...

Service has: 
	Super Sampling,	Median filter

Running the Service:

Post an image to localhost:8080

ss = super sampling

mf = mean filter

sa = search area


Example: 

localhost:8080/ss?sa=10

localhost:8080/mf?sa=1

