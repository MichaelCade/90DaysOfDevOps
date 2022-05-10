---
title: '#90DaysOfDevOps - Variables & Constants in Go - Day 11'
published: false
description: 90DaysOfDevOps - Variables & Constants in Go
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048862
---

Before we get into the topics for today I want to give a massive shout out to [Techworld with Nana](https://www.youtube.com/watch?v=yyUHQIec83I) and this fantastic concise journey through the fundamentals of Go. 

On [Day8](day08.md) we set our environment up, on [Day9](day09.md) we walked through the Hello #90DaysOfDevOps code and on [Day10](day10.md)) we looked at our Go workspace and went a little deeper into compiling and running the code.

Today we are going to take a look into Variables, Constants and Data Types whilst writing a new program. 

## Variables & Constants in Go
Let's start by planning our application, I think it would be a good idea to work on a program that tells us how many days we have remained in our #90DaysOfDevOps challenge. 

The first thing to consider here is that as we are building our app and we are welcoming our attendees and we are giving the user feedback on the number of days they have completed we might use the term #90DaysOfDevOps many times throughout the program. This is a great use case to make #90DaysOfDevOps a variable within our program. 

- Variables are used to store values. 
- Like a little box with our saved information or values. 
- We can then use this variable across the program which also benefits that if this challenge or variable changes then we only have to change this in one place. Meaning we could translate this to other challenges we have in the community by just changing that one variable value. 

To declare this in our Go Program we define a value by using a **keyword** for variables. This will live within our `func main` block of code that you will see later. You can find more about [Keywords](https://go.dev/ref/spec#Keywords)here. 

Remember to make sure that your variable names are descriptive. If you declare a variable you must use it or you will get an error, this is to avoid possible dead code, code that is never used. This is the same for packages not used. 

```
var challenge = "#90DaysOfDevOps"
```
With the above set and used as we will see in the next code snippet you can see from the output below that we have used a variable. 

```
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    fmt.Println("Welcome to", challenge "")
}
```
You can find the above code snippet in [day11_example1.go](Go/day11_example1.go)

You will then see from the below that we built our code with the above example and we got the output shown below. 

![](Images/Day11_Go1.png)

We also know that our challenge is 90 days at least for this challenge, but next, maybe it's 100 so we want to define a variable to help us here as well. However, for our program, we want to define this as a constant. Constants are like variables, except that their value cannot be changed within code (we can still create a new app later on down the line with this code and change this constant but this 90 will not change whilst we are running our application)

Adding the `const` to our code and adding another line of code to print this. 

```
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90

    fmt.Println("Welcome to", challenge)
    fmt.Println("This is a", daystotal, "challenge")
}
```
You can find the above code snippet in [day11_example2.go](Go/day11_example2.go)

If we then go through that `go build` process again and run you will see below the outcome.

![](Images/Day11_Go2.png)

Finally, and this won't be the end of our program we will come back to this in [Day12](day12.md) to add more functionality. We now want to add another variable for the number of days we have completed the challenge. 

Below I added `dayscomplete` variable with the number of days completed. 

```
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90
    var dayscomplete = 11

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge and you have completed", dayscomplete, "days")
    fmt.Println("Great work")
}
```
You can find the above code snippet in [day11_example3.go](Go/day11_example3.go)

Let's run through that `go build` process again or you could just use `go run`

![](Images/Day11_Go3.png)

Here are some other examples that I have used to make the code easier to read and edit. We have up till now been using `Println` but we can simplify this by using `Printf` by using `%v` which means we define our variables in order at the end of the line of code. we also use `\n` for a line break. 

I am using `%v` as this uses a default value but there are other options that can be found here in the [fmt package documentation](https://pkg.go.dev/fmt) you can find the code example [day11_example4.go](Go/day11_example4.go)

Variables may also be defined in a simpler format in your code. Instead of defining that it is a `var` and the `type` you can code this as follows to get the same functionality but a nice cleaner and simpler look for your code. This will only work for variables though and not constants.  

```
func main() {
    challenge := "#90DaysOfDevOps"
    const daystotal = 90
```

## Data Types 
In the above examples, we have not defined the type of variables, this is because we can give it a value here and Go is smart enough to know what that type is or at least can infer what it is based on the value you have stored. However, if we want a user to input this will require a specific type. 

We have used Strings and Integers in our code so far. Integers for the number of days and strings are for the name of the challenge. 

It is also important to note that each data type can do different things and behaves differently. For example, integers can multiply where strings do not. 

There are four categories 

- **Basic type**: Numbers, strings, and booleans come under this category.
- **Aggregate type**: Array and structs come under this category.
- **Reference type**: Pointers, slices, maps, functions, and channels come under this category.
- **Interface type**

The data type is an important concept in programming. Data type specifies the size and type of variable values.

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

Go has three basic data types:

- **bool**: represents a boolean value and is either true or false
- **Numeric**: represents integer types, floating-point values, and complex types
- **string**: represents a string value

I found this resource super detailed on data types [Golang by example](https://golangbyexample.com/all-data-types-in-golang-with-examples/)

I would also suggest [Techworld with Nana](https://www.youtube.com/watch?v=yyUHQIec83I&t=2023s) at this point covers in some detail a lot about the data types in Go. 

If we need to define a type in our variable we can do this like so: 

```
var TwitterHandle string 
var DaysCompleted uint
```

Because Go implies variables where a value is given we can print out those values with the following: 

```
fmt.Printf("challenge is %T, daystotal is %T, dayscomplete is %T\n", conference, daystotal, dayscomplete)
```
There are many different types of integer and float types the links above will cover off these in detail. 

- **int** = whole numbers
- **unint** = positive whole numbers
- **floating point types** = numbers that contain a decimal component

## Resources

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s) 
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I) 
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals) 
- [FreeCodeCamp -  Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s) 
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N) 

Next up we are going to start adding some user input functionality to our program so that we are asking how many days have been completed. 

See you on [Day 12](day12.md).
