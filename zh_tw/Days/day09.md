---
title: '#90DaysOfDevOps - Let''s explain the Hello World code - Day 9'
published: false
description: 90DaysOfDevOps - Let's explain the Hello World code
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048732
---
## 讓我們解釋一下 Hello World 編碼

### GO程式語言如何運作 

在第 8 天，我們瀏覽了您在工作站上安裝 Go 的過程，然後我們創建了我們第一個 Go的應用程序。 
 
在本章節中，我們將更深入地研究代碼並了解更多關於 Go 程式語言的內容。

### 什麼是編譯？ 
在我們進入第6行Hello World代碼之前[6 lines of the Hello World code](Go/hello.go)，我們必需要對編譯有一點了解。 

像我們常用的Python、 Java、Go和C++編程語言都是高階的程試語言。
這意味著它們是人類可辨別的，但是當機器嘗試執行程序時，它需要採用機器可以理解的形式。我們必須將人類可辨別的代碼翻譯成為機器代碼這就稱為編譯。 

![](Images/Day9_Go1.png)

從上面你可以看到我們在第 8 天做了什麼 [Day 8](day08.md)，我們創建了一個簡單的 Hello World main.go檔然後我們使用指令 `go build main.go` 來編譯我們可執行的檔案。 

### 什麼是套件？
套件是在同一目錄中收藏的源碼檔所形成的編譯。我們可以進一步簡化這一點，一個套件是在同一個目錄底下的一堆.go檔案。還記得第 8 天的 Hello 文件檔嗎？如果當您進入更複雜的 Go 程式語言時，您可能會發現你有文件1、文件2、文件3與許多套件所編輯而成的數個.go檔案。 

我們使用套件所以我們可以重複使用其他人的代碼，我們不必從頭開始編寫所有東西。或許我們想擁有一個從長遠來看可以為你節省大量時間與精力的計算器作為我們編程的一部分，在此你可能會找到現有的一個 Go 套件裡，包含你可以導入到代碼中的數學函數。

Go 程式語言鼓勵您將代碼統整在套件中，以便於重新使用和維護源代碼。

### Hello #90DaysOfDevOps 並行 
現在讓我們看一下我們的 Hello #90DaysOfDevOps main.go檔並逐行查看。

![](Images/Day9_Go2.png)

在第一行，你有'main套件'這意味著這個檔案附屬於一個稱作main的資料包。所有.go檔案都必須隸屬於這個套件，它們在初始行應該也有像套件的東西。

一個套件可以任意命名。我們必須在編程的一開頭就下'main'的指令，此動作將會在此套件運行，這就是規則。（還有需要再更加了解的嗎?）

![](Images/Day9_Go3.png)

Whenever we want to compile and execute our code we have to tell the machine where the execution needs to start. We do this by writing a function called main. The machine will look for a function called main to find the entry point of the program. 

A function is a block of code that can do some specific task for and can be used across the program. 

You can declare a function with any name using `func` but in this case we need to name it `main` as this is where the code starts. 

![](Images/Day9_Go4.png)

Next we are going to look at line 3 of our code, the import, this basically means you want to bring in another package to your main program. fmt is a standard package being used here provided by Go, this package contains the `Println()`function and because we have imported this we can use this in line 6. There are a number of standard packages you can include in your program and leverage or reuse them in your code saving you the hassle of having to write from scratch. [Go Standard Library](https://pkg.go.dev/std)

![](Images/Day9_Go5.png)

the `Println()` that we have here is a way in which to write to a standard output to the terminal where ever the executuable has been executed succesfully. Feel free to change the message in between the (). 

![](Images/Day9_Go6.png)

### TLDR

- **Line 1** = This file will be in the package called `main` and this needs to be called `main` because includes the entry point of the program. 
- **Line 3** = For us to use the `Println()` we have to import the fmt package to use this on line 6. 
- **Line 5** = The actual starting point, its the `main` function. 
- **Line 6** = This will let us print "Hello #90DaysOfDevOps" on our system. 

## Resources

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s) 
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I) 
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals) 
- [FreeCodeCamp -  Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s) 
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N) 

See you on [Day 10](day10.md).
