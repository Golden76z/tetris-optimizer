# Tetris-Optimizer

## How to run the programm

```go
go run . [filename]
```
Files are stored in the example folder.  
To run the very first file, you can use the command:
```go
go run . badexample00
```
You don't need to add the **.txt** at the end of the filename.

## Check

The check function will check if the file contains only good tetrominos and if the format is good or not.  
Here's what we need to check to make sure our tetrominos are valid:  
- They need to have exactly 4 blocks, anything lower or higher is a bad tetrominos.  
- They need to have at least 6 links, otherwise that means that the blocks are not linked.  

Good example:
```txt
.###
...#
....
....
```

Bad example:
```txt
.###
...#
...#
....
```
(5 blocks instead of 4)

## Formatting

We get rid of all the empty lines and empty columns in the tetrominos.  

Before formatting:
```txt
....
.##.
.##.
....
```

After formatting:
```txt
##
##
```

## Replace with colors

We then replace all the # with colors, making sure that each tetrominos have a unique color so that we can distinguish them.  

Before replacing them:
```txt
##
##
```

After replacing:
```txt
🟨🟨
🟨🟨
```

## Square size

We then have a function that will give us the minimum size of the square we have to check.  
For example if we have 12 tetrominos, 12 x 4 = 48.  
So the minimum square is 7 by 7 because 7 x 7 = 49.  

## Solver

We then use the backtracking Algorithm to find the best solution to fit all the tetrominos of the file into the minimum square possible.  
It place a tetromino if it can, otherwise it goes to the next tetromino until he finds the solution.  
If no solutions were found on the current square, we increase the size of the square and try again.  

 ### MIT License

Copyright (c) [2024]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.