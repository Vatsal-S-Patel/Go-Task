# Go Tasks

## Task-3 : Sudoku Solver


### Problem Statement

9x9 grid, your task is to check whether the given sudoku is solvable or not. The input for your program will be a 9x9 matrix where empty cells are represented by 0, and the filled cells are represented by their respective digits.


### Instructions 

1. Create a Golang program that takes the input Sudoku grid and returns whether the given sudoku is solvable or not.
2. If it fails then also return the row and column for which it fails
3. Test your solution with different Sudoku puzzles to validate its correctness.


### Solution

Approach : Recursion and Backtracking

The steps to solve or validate Sudoku is as given below:
1. Firstly it checks that sudoku is solvable or not. To check that, it checks for the right digit for empty place which is denoted by 0.
2. To find that given Sudoku is solvable or not, it will place all 1 to 9 digit and checks that it is right or clashing. 
3. It will iterate through all the sudoku board, row by row and checks for 0 and iterate through 1 to 9 and if there is no digit fitting for that place, it will backtrack and again try with other digit placing at that place
4. If place is not empty then it will continue to check other places denoted with 0 by iterating the sudoku board row by row and again place all digit from 1 to 9 and continue.
5. Hence, using above approach, we can solve the Sudoku and also validate it.


### How to Run Code ?

1. Clone this Repo
2. Go to `task-3` branch
3. `cd task-3` to switch directory 
4. `go run main.go` to execute this file