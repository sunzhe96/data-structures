# Data Structures

This projects is my implementations of some basic data structures. The implementation is in the Go programming language. 

## Data Type

The data structures in this repository can only be used to store integer data which is the "int" type in golang.

## List of the data structures implemented

- [X] Linked Lists
- [X] Stacks
- [X] Quenes
- [X] Binary Search Trees
- [X] Hash Tables
- [X] Binary Search Trees
- [] Priority Quenes
- [] Disjoint Sets
- [] Binary Search Trees
- [] AVL Trees
- [] Splay Trees

## Linked lists 

| API                               | function                                                                      |  
|-----------------------------------|-------------------------------------------------------------------------------|  
| func New() *List                  | returns a pointer to an empty list                                            |  
|                                   |                                                                               |  
| func (l *List) PushFront(key int) | adds a key to the front of l                                                  |  
|                                   |                                                                               |  
| func (l *List) TopFront() int     | returns the key of the front item. the program will panic if l is empty       |  
|                                   |                                                                               |  
| func (l *List) PopFront() int     | removes the key which is the front item. the program will panic if l is empty |  
|                                   |                                                                               |  
| func (l *List) PushBack() int     | adds a key to the back of l                                                   |  
|                                   |                                                                               |  
| func (l *List) TopBack() int      | returns the key of the back item                                              |  
|                                   |                                                                               |  
| func (l *List) Popback() int      | remove the key of the back item                                               |  
|                                   |                                                                               |  
| func (l *List) Find(key int) bool | reports whether key is in l                                                   |  
|                                   |                                                                               |  
| func (l *List) Erase(key int)     | removes key from l                                                            |  
|                                   |                                                                               |  
| func (l *List) Empty() bool       | reports wheter l is empty                                                     |  
|                                   |                                                                               |  
| func (l *List) Print()            | prints l                                                                      |  

## Stacks 

| API                           | function                                        |  
|-------------------------------|-------------------------------------------------|  
| func New() *Stack             | returns a pointer to an empty stack             |  
|                               |                                                 |  
| func (s *Stack) Push(key int) | adds key to s                                   |  
|                               |                                                 |  
| func (s *Stack) Pop() int     | removes and returns the most recently-added key |  
|                               |                                                 |  
| func Top(s *Stack) int        | returns the most recently-added key             |  
|                               |                                                 |  
| func Empty(s *Stack) int      | reports whether s is empty                      |  
|                               |                                                 |  
| func (s *Stack) Print()       | prints s                                        |  

## Quenes 

| API                           | function                                       |  
|-------------------------------|------------------------------------------------|  
| func New() *Quene             | returns a pointer to an empty qunene           |  
|                               |                                                |  
| func (q *Quene) EnQuene()     | adds key to q                                  |  
|                               |                                                |  
| func (q *Quene) DeQuene() int | removes and returns the least recent-added key |  
|                               |                                                |  
| func (q *Quene) Empty() bool  | reports whether q is empty                     |  

## Hash Tables 

| API                                    | Operation                                         |  
|----------------------------------------|---------------------------------------------------|  
| func New() HashTable                   | returns an empty hash table                       |  
|                                        |                                                   |  
| func (h *HashTable) Search(Key)        | prints the corresponding value                    |  
|                                        |                                                   |  
| func (h *HashTable) Insert(Key, Value) | inserts a key-value pair into h                   |  
|                                        |                                                   |  
| func (h *HashTable) Delete(Key)        | deletes the key-value pair by key                 |  
|                                        |                                                   |  
| func (h *HashTable) Print()            | prints the hash table by each slot                |  

## Binary Search Trees

| API                                        | function                                          |  
|--------------------------------------------|---------------------------------------------------|  
| func NewNode(k int) *Node                  | returns a pointer to a node with key equal to k   |  
|                                            |                                                   |  
| func NewBST() *BST                         | returns a pointer to an empty BST                 |  
|                                            |                                                   |  
| func (t *BST) Print()                      | prints t                                          |  
|                                            |                                                   |  
| func (t *BST) Search(k int) *Node          | returns a pointer to the node with key equal to k |  
|                                            |                                                   |  
| func (t *BST) Min() *Node                  | returns the node with minimum key in t            |  
|                                            |                                                   |  
| func (t *BST) Max() *Node                  | returns the node with maximum key in t            |  
|                                            |                                                   |  
| func (n *Node) Successor() *Node           | returns a pointer to the successor of n           |  
|                                            |                                                   |  
| func (t *Node) Predecessor() *Node         | returns a pointer to the predecessor of n         |  
|                                            |                                                   |  
| func (t *BST) Insert(n *Node)              | inserts n into t                                  |  
|                                            |                                                   |  
| func (t *BST) transplant(u *Node, v *Node) | replaces u with v                                 |  
|                                            |                                                   |  
| func (t *BST) Delete(n *Node)              | deletes n                                         |  

