# Day 45 - Python: Debugging, testing and Regular expression

Welcome to Day 4 of Python!
Today we will learn about:

- Debugging and testing
- Regular expressions
- Datetime library

Let's start!

## Debugging and testing

Debugging is the process of finding and correcting errors or bugs in code. Python includes a debugger called `pdb` that allows you to step through your code and inspect variables as you go. You can use `pdb` to help you figure out where your code is going wrong and how to fix it.

``` python
import pdb

def add_numbers(x, y):
    result = x + y
    pdb.set_trace() # Start the debugger at this point in the code
    return result

result = add_numbers(2, 3)
print(result)
```

In this example, we define the `add_numbers` function, which adds two numbers and returns the result. To start the debugger at a specific point in the code, we use the pdb.set trace() function (in this case, after the result has been calculated). This enables us to inspect variables and step through the code to figure out what's going on.

In addition to debugging, testing is an important part of programming. It entails creating test cases to ensure that your code is working properly. Python includes a `unittest` module that provides a framework for writing and running test cases.

``` python
import unittest

def is_prime(n):
    if n < 2:
        return False
    for i in range(2, n):
        if n % i == 0:
            return False
    return True

class TestIsPrime(unittest.TestCase):
    def test_is_prime(self):
        self.assertTrue(is_prime(2))
        self.assertTrue(is_prime(3))
        self.assertTrue(is_prime(5))
        self.assertFalse(is_prime(4))

if __name__ == '__main__':
    unittest.main()

```

Output:

``` bash
----------------------------------------------------------------------
Ran 1 test in 0.000s

OK
```

## Regular expressions:

In Python, regular expressions are a powerful tool for working with text data. They enable you to search for and match specific character patterns within a string. Python's `re` module includes functions for working with regular expressions.
For example, you can use regular expressions to search for email addresses within a larger block of text, or to extract specific data from a string that follows a particular pattern.

``` python
import re

# Search for a phone number in a string
text = 'My phone number is 555-7777'
match = re.search(r'\d{3}-\d{4}', text)
if match:
    print(match.group(0))

# Extract email addresses from a string
text = 'My email is example@devops.com, but I also use other@cloud.com'
matches = re.findall(r'\S+@\S+', text)
print(matches)
```

Output:

``` bash
555-7777
['example@devops.com,', 'other@cloud.com']
```

## Datetime library:

As the name suggests, Python's `datetime` module allows you to work with dates and times in your code. It includes functions for formatting and manipulating date and time data, as well as classes for representing dates, times, and time intervals.
The datetime module, for example, can be used to get the current date and time, calculate the difference between two dates, or convert between different date and time formats.

``` python
from datetime import datetime, timedelta

# Get the current date and time
now = datetime.now()
print(now) # Output: 2023-02-17 11:33:27.257712

# Create a datetime object for a specific date and time
date = datetime(2023, 2, 1, 12, 0)
print(date) # Output: 2023-02-01 12:00:00

# Calculate the difference between two dates
delta = now - date
print(delta) # Output: 15 days, 23:33:27.257712
```

Output:

``` bash
2023-02-17 11:33:27.257712
2023-02-01 12:00:00
15 days, 23:33:27.257712
```

## Resources

- [pdb - The Python Debugger](https://docs.python.org/3/library/pdb.html)
- [re - Regular expressions operations](https://docs.python.org/3/library/re.html)
- [datetime - Basic date and time types](https://docs.python.org/3/library/datetime.html)

 See you tomorrow in [Day 46](day46.md).
 