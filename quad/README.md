#   Instructions

Write a function ```QuadA``` that prints a **valid** rectangle with a given width of ```x``` and height of ```y```.

The function must draw the rectangles as in the examples.

If ```x``` and y are positive numbers, the program should print the rectangles as seen in the examples, otherwise, the function should print nothing.

Make sure you submit all the necessary files to run the program.

**Expected function**
```
func QuadA(x,y int) {

}
```
#   Usage

Here are possible programs to test your function :

Program #1
```
package main

import "piscine"

func main() {
	piscine.QuadA(5,3)
}
```
And its output :
```
$ go run .
o---o
|   |
o---o
$
```
Program #2
```
package main

import "piscine"

func main() {
	piscine.QuadA(5,1)
}
```
And its output :
```
$ go run .
o---o
$
```
Program #3
```
package main

import "piscine"

func main() {
	piscine.QuadA(1,1)
}
```
And its output :
```
$ go run .
o
$
```
Program #4
```
package main

import "piscine"

func main() {
	piscine.QuadA(1,5)
}
```
And its output :
```
$ go run .
o
|
|
|
o
$
```