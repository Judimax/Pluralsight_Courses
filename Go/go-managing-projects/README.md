# Course Overview

## Course Overview

# Getting Started with Go Packages

## Introduction

## Prerequisites &amp; Scenario Setup

## Introduction to Packages in Go

## Demo: Packages in Go

## Scoping: Package vs. Public Scoping

## Demo: Scoping: Package vs. Public Scoping
![1689947446933](image/README/1689947446933.png)

## Initializing of Package
* go does not requires a init function
![1689947988332](image/README/1689947988332.png)


## Demo: Initializing of Package
go-managing-projects\02\demos\demo9\main.go
 func init runs before main,
 * runs things in lexical order
 go-managing-projects\02\demos\demo11

* all the init fn get changed before the main function, if all values are defined main function runs last


## Summary
* uppercase is a pulbic scope can be accessed by other values

# Exploring Go Modules


## Introduction to Modules in Go
![1689951221347](image/README/1689951221347.png)

## Demo: Modules in Go
* if the files is in the same directory no need to import
* go.sum is like a package-lock.json,

## Understanding Versioning
![1689952141992](image/README/1689952141992.png)
* programs are compatilbe if major and minor versions are the same

## Managing Dependency
![1689952254809](image/README/1689952254809.png)
* go mod tidy removes unused import and adds needed import

## Demo: Managing Dependency with Versioning
* to downgrade a go version
```go
go get github.com/pioz/faker@v0.1.1
```

* to upgrade to latest,
```ps1
go mod -u
```

* to remove unneeded packages, and add needed packages
```
go mod tidy
```

* you can use github as your repo for versioning

## Implementing Vendoring
* all deps can be bundle and deployed

## Demo: Implementing Vendoring
![1689952909000](image/README/1689952909000.png)
* to make a package vendor so it does not need the internet
```ps1
go mod vendor -v
```

## Dependency Management Tool - dep
![1689953270206](image/README/1689953270206.png)

##  Demo: Dependency Management Tool - dep

## Summary

# Embedding Resource in Go


## Introduction

## Advantages of Embedding Resources

## Demo: Embedding Single File in a Project

## Demo: Embedding Multiple Files in a Web Project
04\demos\ps.m4\demo2\main.go
  <h1>Advantages of using go:embed</h1>
  <ul>
    <li>Simplified distribution    </li>
    <li>Improved reliability    </li>
    <li>Enhanced security    </li>
    <li>Improved performance    </li>
    <li>Increased portability    </li>
    <li>Fun to use    </li>
  </ul>

## Summary
