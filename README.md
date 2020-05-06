# SodaDB
**🥤 Type-Safe NoSQL Key-Value Datastore**

## Usage
### Set
#### Single Value
```
set [datatype] [value] [key]
```
##### Example 
```
set INT 300 foo
```

#### List
```
set [datatype] [values] [key]
```
##### Example 
```
set INT (1,4,2,5,2) bar
```


### Get
```
get [key]
```
##### Example 
```
get bar
```
```bash
(1,4,2,5,2)
```

Inspired by: https://notes.eatonphil.com/database-basics.html

Written with 💙
