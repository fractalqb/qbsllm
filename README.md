# qbsllm – Slim Logging with Go
[![Build Status](https://travis-ci.org/fractalqb/qbsllm.svg)](https://travis-ci.org/fractalqb/qbsllm)
[![codecov](https://codecov.io/gh/fractalqb/qbsllm/branch/master/graph/badge.svg)](https://codecov.io/gh/fractalqb/qbsllm)
[![Go Report Card](https://goreportcard.com/badge/github.com/fractalqb/qbsllm)](https://goreportcard.com/report/github.com/fractalqb/qbsllm)
[![GoDoc](https://godoc.org/github.com/fractalqb/qbsllm?status.svg)](https://pkg.go.dev/git.fractalqb.de/fractalqb/qbsllm)

`import "git.fractalqb.de/fractalqb/qbsllm"`

---

This lib is deprecated. [qblog](https://codeberg.org/fractalqb/qblog)
is a replacement that features sllm as message format.

This is a logging library that brings together

- [Slim structured logging](https://git.fractalqb.de/fractalqb/sllm)
with
- [leveled logging from qblog](https://codeberg.org/fractalqb/qblog)
and
- the [logging configuration from
  c4hgol](https://codeberg.org/fractalqb/c4hgol)

to make a logging library that can coexists peacefully with other
logging libraries.

## Dependencies
![Dependency Graph](file:depgraph.svg)
