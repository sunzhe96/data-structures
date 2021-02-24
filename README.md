# Data Structures

This repository contains my implementations of the data structures taught in a Data Structures Course of Coursera. The implementation is in the Go programming language. 

## Data Type

The data structures implemented in this repository can only be used to store integer data which is the "int" type in golang.

##  The list of data structures I have implemented.

- [X] Linked Lists
- [X] Stacks
- [X] Quenes
<<<<<<< HEAD
- [X] Hash Tables
- [] Trees
=======
- [X] Binary Search Trees
>>>>>>> docs
- [] Priority Quenes
- [] Disjoint Sets
- [] Binary Search Trees
- [] AVL Trees
- [] Splay Trees

## Linked lists 

| API                                   | function                     |
|---------------------------------------|------------------------------|
| func (l *List) PushFront(key int) int | add key to front             |
|                                       |                              |
| func (l *List) TopFront() int         | return key of the front item |
|                                       |                              |
| func (l *List) PopFront() int         | remove the front item        |
|                                       |                              |
| func (l *List) PushBack() int         | add item to back             |
|                                       |                              |
| func (l *List) TopBack() int          | return key of back item      |
|                                       |                              |
| func (l *List) Popback() int          | remove back item             |
|                                       |                              |
| func (l *List) Find(key int) bool     | is key in list?              |
|                                       |                              |
| func (l *List) Erase(key int)         | remove key from list         |
|                                       |                              |
| func (l *List) Empty() bool           | is list empty?               |
|                                       |                              |
| func (l *List) PrintList()            | print the list               |

## Stacks 

| API     | function                                    |  
|---------|---------------------------------------------|  
| Push()  | adds key to collection                      |  
|         |                                             |  
| Pop()   | removes and returns most recently-added key |  
|         |                                             |  
| Top()   | returns most recently-added key             |  
|         |                                             |  
| Empty() | are there any elements?                     |  

## Quenes 

| API       | function                                       |  
|-----------|------------------------------------------------|  
| EnQuene() | Adds key to the collection                     |  
|           |                                                |  
| DeQuene() | removes and returns the least recent-added key |  
|           |                                                |  
| Empty()   | are there any elements?                        |  


<<<<<<< HEAD
## Hash Tables 

| API                | Operation                                        |  
|--------------------|--------------------------------------------------|  
| Search(Key)        | return the corresponding value                   |  
|                    |                                                  |  
| Insert(Key, Value) | insert a key-value pair in the dictionary        |  
|                    |                                                  |  
| Delete(Key)        | delete the key-value pair                        |  
|                    |                                                  |  
| MakeHashTable()    | initialize a hashtable and return the hash table |  
|                    |                                                  |  
| PrintTable()       | print the hash table by each slot                |  

=======
## Binary Search Trees

| API       | function                                       |  
|-----------|------------------------------------------------|  
| EnQuene() | Adds key to the collection                     |  
|           |                                                |  
| DeQuene() | removes and returns the least recent-added key |  
|           |                                                |  
| Empty()   | are there any elements?                        |  
>>>>>>> docs
