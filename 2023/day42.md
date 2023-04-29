# Day 42 - Programming Language:Introduction to Python

Guido van Rossum created Python, a high-level, interpreted and dynamic programming language, in the late 1980s. It is widely used in range of applications, including web development, devops and data analysis, as well as artificial intelligence and machine learning.

## Installation and Setting up the Environment:

Python is available for download and installation on a variety of platforms, including Windows, Mac, and Linux. Python can be downloaded from [the official website](https://www.python.org/.).
![Python Website](/2023/images/day42-01.png)

Following the installation of Python, you can configure your environment with an Integrated Development Environment (IDE) such as [PyCharm](https://www.jetbrains.com/pycharm/), [Visual Studio Code](https://code.visualstudio.com/), or IDLE (the default IDE that comes with Python).
I personally use Visual Studio Code.

You can also use cloud environment, where you will not have to configure and install python locally, like [Replit](https://replit.com/).
![Replit Website](/2023/images/day42-02.png)

## Basic Data Types:

Python includes a number of built-in data types for storing and manipulating data. The following are the most common ones:

- Numbers: integers, floating-point numbers, and complex numbers
- Strings are character sequences.
- Lists are ordered groups of elements.
- Tuples are ordered immutable collections of elements.
- Dictionaries are collections of key-value pairs that are not ordered.

## Operations and Expressions:

With the above data types, you can perform a variety of operations in Python, including arithmetic, comparison, and logical operations.
Expressions can also be used to manipulate data, such as combining multiple values into a new value.

## Variables:

A variable is declared and assigned a value in Python by using the assignment operator =. The variable is on the left side of the operator, and the value being assigned is on the right, which can be an expression like `2 + 2` or even other variables. As an example:

``` python
a = 7         # assign variable a the value 7
b = x + 3     # assign variable b the value of a plus 3
c = b         # assign variable c the value of b
```

These examples assign numbers to variables, but numbers are only one of the data types supported by Python. There is no type declaration for the variables. This is due to the fact that Python is a dynamically typed language, which means that the variable type is determined by the data assigned to it. The x, y, and z variables in the preceding examples are integer types, which can store both positive and negative whole numbers.

Variable names are case sensitive and can contain any letter, number, or underscore ( ). They cannot, however, begin with a number.
Also, with numbers, strings are among the most commonly used data types. A string is a sequence of one or more characters. Strings are typically declared with single quotation marks, but they can also be declared with double quotation marks:

``` python
a = 'My name is Rishab'
b = "This is also a string"
```

You can add strings to other strings — an operation known as "concatenation" — with the same + operator that adds two numbers:

``` python
x = 'My name is' + ' ' + 'Rishab'
print(x) # outputs: My name is Rishab
```

## Printing to the console:

The print function in Python is one of more than 60 built-in functions. It outputs text to the screen.
Let's see an example of the most famous "Hello World!":

``` python
print('Hello World!')
```

The print argument is a string, which is one of Python's basic data types for storing and managing text. Print outputs a newline character at the end of the line by default, so subsequent calls to print will begin on the next line.

## Resources:

- [Learn Python - Full course by freeCodeCamp](https://youtu.be/rfscVS0vtbw)
- [Python tutorial for beginners by Nana](https://youtu.be/t8pPdKYpowI)
- [Python Crash Course book](https://amzn.to/40NfY45)

See you on [Day 43](day43.md).