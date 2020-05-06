# SodaDB
**🥤 Type-Safe NoSQL Key-Value Datastore**

[![Go Report Card](https://goreportcard.com/badge/github.com/ReubenMathew/SodaDB)](https://goreportcard.com/report/github.com/ReubenMathew/SodaDB)

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
set [datatype] [value] [key]
```
***Example*** 
```
set INT 300 foo
```

#### List
```
set [datatype] [values] [key]
```
***Example*** 
```
set INT (1,4,2,5,2) bar
```


### Get
*All keys are stored as strings*
```
get [key]
```
***Example*** 
```
get ba
```
```bash
(1,4,2,5,2)
```
***Example*** 
```
get foo
```
```bash
300
```


**Written with 💙**





Inspired by: 
- https://notes.eatonphil.com/database-basics.html
- https://jeffknupp.com/blog/2014/09/01/what-is-a-nosql-database-learn-by-writing-one-in-python/

