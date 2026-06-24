# Day 43 - Programming Language: Python Loops, functions, modules and libraries

Welcome to the second day of Python, and today we will cover some more concepts:
- Loops
- Functions
- Modules and libraries
- File I/O

## Loops (for/while):

Loops are used to repeatedly run a block of code.

### for loop

Using the `for` loop, a piece of code is executed once for each element of a sequence (such as a list, string, or tuple). Here is an example of a for loop that prints each programming language in a list:

``` python
languages = ['Python', 'Go', 'JavaScript']

# for loop
for language in languages:
    print(language)
```

Output
```
Python
Go
JavaScript
```

### while loop

The `while loop` is used to execute a block of code repeatedly as long as a condition is True. Here's an example of a while loop that prints the numbers from 1 to 5:

``` python
i = 1
n = 5

# while loop from i = 1 to 5
while i <= n:
    print(i)
    i = i + 1
```

Output:
```
1
2
3
4
5
```

## Functions
Functions are reusable chunks of code with argument/parameters and return values.
Using the `def` keyword in Python, you can define a function. In your programme, functions can be used to encapsulate complex logic and can be called several times.
Functions can also be used to simplify code and make it easier to read. Here is an illustration of a function that adds two numbers:

``` python
# function has two arguments num1 and num2
def add_numbers(num1, num2):
    sum = num1 + num2
    print('The sum is: ',sum)
```

``` python
# calling the function with arguments to add 5 and 2
add_numbers(5, 2)

# Output: The sum is: 9
```

## Understanding Modules and Importing Libraries:
A module is a file in Python that contains definitions and statements. Modules let you arrange your code and reuse it across many apps.
The Standard Library, a sizable collection of Python modules, offers a wide range of capabilities, such as file I/O, regular expressions, and more.
Additional libraries can be installed using package managers like pip.
You must import a module or library using the import statement in order to use it in your programme. Here is an illustration of how to load the math module and calculate a number's square root using the sqrt() function:

``` python
import math

print(math.sqrt(16)) # 4.0
```

## File I/O
File I/O is used to read and write data to and from files on disk.
The built-in Python function open() can be used to open a file, after which you can read from and write to it using methods like read() and write().
To save system resources, you should always close the file after you are done with it.
An example of reading from a file and printing its content:

``` python
f = open("90DaysOfDevOps.txt", "r")
print(f.read())
f.close()
```

## Exception Handling:

Exceptions are runtime errors that happen when your programme runs into unexpected circumstances, such dividing by zero or attempting to access a list element that doesn't exist.
Using a try/except block, you can manage exceptions in Python. The try block's code is run, and if an exception arises, the except block's code is run to handle the exception.

``` python
try:
  f = open("90DaysOfDevOps.txt")
  try:
    f.write("Python is great")
  except:
    print("Something went wrong when writing to the file")
```

## Conclusion

That is it for today, I will see you tomorrow in [Day 44 | Day 3 of Python!](day44.md).
