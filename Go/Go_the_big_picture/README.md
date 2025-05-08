# Course Overview

## Course Overview

# Why Go?

## Introduction

## Why Create Go?
* language of the cloud
* efficient, ease

## Guiding Principles
* simplicity as a core value
* just use 1 way to do things
* strong commitment backward compatibility

## Language Characteristics
* go is cross platform
* rapid compliation

## Primary Use Cases

## Summary

# Programming with Go

## Introduction
* what would C look like if we made it today
* we dont hold ourselves to standards

## Language Characteristics and Evolution

## Demo: Hello, World
to start a new app
```go
go mod init demo
```
* to work with string
```go
import "fmt"
```

## Demo: A Simple Web Service

## Demo: CLI Application
* error monitor is impt if we get an err, err var is defined
```go
	f, err := os.Open("./log.txt")
```
crash the app
```go
		log.Fatal(err)
```

* closes when the main fn closes
```go
	defer f.Close()
```

* to run the cli command
```go
 go run . -level DEBUG
```
## Summary

# Leveraging the Go Ecosystem
* vsc and goland from jetbrains
## Introduction

## Development Tools


## Common Frameworks
* go kit, microservice framework
* gin, API framework
* gorrila tookkit, API framework
* cobra CLI framework
* std lib CLI framework
* docker Cloud framework


## Community Resources
* go.dev
effective_go - how you go about in go
* there is a go playground
* std library


## Overview of Go's Website

## Learning Resources

## Online Community
https://github.com/golang/go/wiki#the-go-community

## Conferences

## Summary
