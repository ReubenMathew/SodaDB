# SodaDB
**Type-Safe NoSQL Key-Value Datastore 🥤**

[![Go Report Card](https://goreportcard.com/badge/github.com/ReubenMathew/SodaDB)](https://goreportcard.com/report/github.com/ReubenMathew/SodaDB)
[![CircleCI](https://circleci.com/gh/ReubenMathew/SodaDB.svg?style=shield)](https://app.circleci.com/pipelines/github/ReubenMathew/SodaDB)

Features
---
- get
- set

Supported Data Types
---
- Integer (32 Bit) (INT)
- String (STR)
- Character (CHAR)
- Byte (BYTE)

Usage
---
### Set
#### Single Value
```
soda> set [datatype] [value] [key]
```
#### List
```
soda> set [datatype] [values] [key]
```
***Example*** 
```
soda> set INT (1,4,2,5,2) foo // single
soda> set INT 300 bar // list
```


### Get
*All keys are stored as strings*
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



**Written with 💙**



Inspired by: 
- https://notes.eatonphil.com/database-basics.html
- https://jeffknupp.com/blog/2014/09/01/what-is-a-nosql-database-learn-by-writing-one-in-python/

