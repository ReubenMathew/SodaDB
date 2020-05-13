# SodaDB
**A Volatile Concurrent Key-Value Datastore  🥤**

[![Go Report Card](https://goreportcard.com/badge/github.com/ReubenMathew/SodaDB)](https://goreportcard.com/report/github.com/ReubenMathew/SodaDB)

Install
---
Visit the [releases](https://github.com/ReubenMathew/SodaDB/releases) page.

If you don't see one compatible with your operating system, open an *Issue*.

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

Usage
---
### Set
*Note: All keys are stored as strings*
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


**Drivers for Python and Go coming soon!**


Written with 💙



Inspired by: 
- https://notes.eatonphil.com/database-basics.html
- https://jeffknupp.com/blog/2014/09/01/what-is-a-nosql-database-learn-by-writing-one-in-python/

