# cropit

This repository consists of three Go packages:

- [crop](github.com/amitsaha/cropit/crop): This package is a thin wrapper around the
  [cutter](https://godoc.org/github.com/oliamb/cutter) package and exports a single function ``Crop()``
- [cropit](github.com/amitsaha/cropit/cropit): This package is a program which uses the above package to
  implement a basic image cropping program
- [server](github.com/amitsaha/cropit/server): This package is a HTTP server using the ``crop`` pacakge
  and thus a HTTP client can send images to it and get the cropped images back.
  
[![GoDoc](https://godoc.org/github.com/amitsaha/cropit?status.svg)](https://godoc.org/github.com/amitsaha/cropit)
