# SodaDB
**A Beginner-Friendly Concurrent Datastore**

[![Go Report Card](https://goreportcard.com/badge/github.com/ReubenMathew/SodaDB)](https://goreportcard.com/report/github.com/ReubenMathew/SodaDB)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FReubenMathew%2FSodaDB%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/ReubenMathew/SodaDB/goto?ref=master)

Install
---
Visit the [releases](https://github.com/ReubenMathew/SodaDB/releases) page and download the latest release.

If you don't see one compatible with your operating system, open an *Issue*.

*Issue templates for releases coming soon!*

Build from source
---
1. Clone repository
2. Run build.sh in */tools* folder to reproduce releases
or navigate to */src* and run `go build .`

Features
---
*More features will be added in the future*
- get
- set

Supported Data Types
---
- Integer (32 Bit) (INT)
- String (STR)
- Character (CHAR)
- Byte (BYTE)

Using the CLI
--- 
1. Run **SodaDB** executable
2. Choose to launch either the client or server application
3. Starting storing!

Usage
---
### Set
#### Single Value
```
soda> set [datatype] [value] [key]
```
#### List
```
soda> set list [datatype] [values] [key]
```
***Example*** 
```
soda> set INT (1,4,2,5,2) foo // single
soda> set INT 300 bar // list
```


### Get
```
soda> get [key]
```

***Example*** 
```
soda> get foo
(1,4,2,5,2)

soda> get bar
300
```


**Drivers for Python, Go and TypeScript coming soon!**


Written with 💙



Inspired by: 
- https://notes.eatonphil.com/database-basics.html
- https://jeffknupp.com/blog/2014/09/01/what-is-a-nosql-database-learn-by-writing-one-in-python/

