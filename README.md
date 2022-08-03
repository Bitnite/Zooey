# Zooey

> Zooey Lang

</br>


<h3> Zooey can do almost anything that a basic programming language does: </h3>

### Types

Zooey has the following data types: `null`, `bool`, `int`, `str`, `array`,
`hash`,`fn`

Type      | Syntax                                    | 
--------- | ----------------------------------------- | 
null      | `null`                                    |
bool      | `true false`                              |
int       | `0 24 1654 -10`                           | 
str       | `"blablabla string here"`                 | 
array     | `[] [1, 2] [1, 2, 3]`                     |
hash      | `{} {"a": 1} {"a": 1, "b": 2}`            |


### Variable Bindings
> You can use 'owo (var) :=: (expression) 
```
owo myVar :=: 5
```

### Arithmetic Expressions
```
>> owo myVar :=: 10
>> owo anotherVar :=: myVar * 2
>> (myVar + anotherVar) / 2 - 3
12
```

### Conditional Expressions

Zooey supports `if` and `else`:

```
>> owo a :=: 10
>> owo b :=: a * 2
>> owo c :=: if b > a { 99 } else { 100 }
>> c
> 99

>> a :=: 10
>> b :=: 10
>> c :=: if b >= a { 99 } else { 100 }
>> c
> 99
```
## Installation

TODO

## Usage example
### Variable Declaration
```
owo myVariable :=: 5
owo myFunction :=: fn(x){x + 1}
owo myHash :=: {"name": "myName", "otherName : "randomName"}
owo myList :=: [1,3,4,5, fn(x){x + 1}]

myList[4](5)
> 6

myHash["name"]
>myName

map([1,2,3,4],fn(x){x * 2})
>[2,4,6,8]

sum(map([1,2,3,4],fn(x){x * 2}))
>20
```

### While Loops

Zooey supports only one looping construct, the `while` loop:

```
owo x :=: 3
owo myList :=: [1,2,3,4,5]

while (x > 0) {
 owo myList :=: map(myList,fn(x){x * 2}) 
 plsShow(myList) 
 owo x :=: x - 1
}

// [2,4,6,8,10]
// [4,8,12,16,20]
// [8,16,24,32,40]
```

### Strings

```
>> owo makeGreeter :=: fn(greeting) { fn(name) { greeting + " " + name + "!" } }
>> owo hello :=: makeGreeter("Hello")
>> hello("mellum")
Hello mellum!
```

### Arrays

```sh
>> owo myArray :=: ["Thorsten", "Ball", 28, fn(x) { x * x }]
>> myArray[0]
Thorsten
>> myArray[4 - 2]
28
>> myArray[3](2)
4
```
### Hashes

```sh
>> owo myHash :=: {"name": "Jimmy", "age": 72, true: "yes, a boolean", 99: "correct, an integer"}
>> myHash["name"]
Jimmy
>> myHash["age"]
72
>> myHash[true]
yes, a boolean
>> myHash[99]
correct, an integer
```
### Builtin functions
- `len(iterable)`
  Returns the length of the iterable (`str`, `array` or `hash`).
- `first(iterable)`
  Returns the first element of the array.
- `last(iterable)`
  Returns the last element of the array.
- `push(array, element)`
  Add an element to the array.
- `replace(string,element_to_replace,element_to_put)`
  Replace something inside a string.
- `plsShow(element)`
  Print something in the screen.
- `show(element, element ...)`
  Concatenate all elements
- `Strcomp(string,string)`
  Compare the pointers of two strings.
- `String_Upcase(string)`
  Upcase an entire string.
- `String_Split(string)`
  Split a string into array



## Meta

Antonio Mello Babo – [@MelloTonio](https://github.com/MelloTonio/)
Stefano Martins Giordano– [@Giordano26](https://github.com/Giordano26/)
???????????????? - [??????]




