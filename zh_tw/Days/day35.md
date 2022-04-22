---
title: '#90DaysOfDevOps - The Big Picture: Git - Version Control - Day 35'
published: false
description: 90DaysOfDevOps - The Big Picture Git - Version Control
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049041
---
## 概述: Git - 版本控制

在開始 Git 工具之前, 我們需要瞭解什麼是版本控制? 為什麼要使用它? 作為 Git 教學的開端, 我們會粗略帶過何謂版本控制和 Git 基礎概念.  

### 甚麼是版本控制? 

由於 Git 並不是唯一的版本控制工具，所以在這裡將會涵蓋版本控制有哪些選項和使用方法. 

版本控制最明顯且重要的功能是可以追蹤專案的歷史. 我們可以利用`git log` 指令回顧倉儲內,過去專案提交的紀錄和註解. 無須擔心,稍後就會說明要如何下指令. 讓我們想想, 如果這是實際的軟體專案, 有著滿滿的原始碼,許多開發者在不同的時間點提交軟體, 不同的作者和審稿人在這裡做紀錄, 我們得以知道發生了什麼事, 何時發生, 又是誰在審稿. 

![](Images/Day35_Git1.png)

版本控制很炫, 就像是你在更動程式碼前手動複製一份檔案一樣. 你可能會為了以防萬一, 註解掉無用的舊程式碼, 而你隨時可以還原. 

![](Images/Day35_Git2.png)

I have started using version control over not just source code but pretty much anything, talks about projects like this (90DaysOfDevOps) because why would you not want that rollback and log of everything that has gone on. 

However, a big disclaimer **Version Control is not a Backup!**

Another benefit of Version Control is the ability to manage multiple versions of a project, Let's create an example, we have a free app that is available on all operating systems and then we have a paid-for app also available on all operating systems. The majority of the code is shared between both applications. We could copy and paste our code each commit to each app but that is going to be very messy especially as you scale your development to more than just one person, also mistakes will be made. 

The premium app is where we are going to have additional features, let's call them premium commits, the free edition will just contain the normal commits. 

The way this is achieved in Version Control is through branching. 

![](Images/Day35_Git3.png)

Branching allows for two code streams for the same app as we stated above. But we will still want new features that land in our source code free version to be in our premium and to achieve this we have something called merging. 

![](Images/Day35_Git4.png)

Now, this same easy but merging can be complicated because you could have a team working on the free edition and you could have another team working on the premium paid for version and what if both change code that affects aspects of the overall code. Maybe a variable gets updated and breaks something. Then you have a conflict that breaks one of the features. Version Control cannot fix the conflicts that are down to you. But version control allows this to be easily managed. 

The primary reason if you have not picked up so far for version control, in general, is the ability to collaborate. The ability to share code amongst developers and when I say code as I said before more and more we are seeing much more use cases for other reasons to use source control, maybe its a joint presentation you are working on with a colleague or a 90DaysOfDevOps challenge where you have the community offering their corrections and updates throughout the project. 

Without version control how did teams of software developers even handle this? I find it hard enough when I am working on my projects to keep track of things. I expect they would split out the code into each functional module. Maybe a little part of the puzzle then was bringing the pieces together and then problems and issues before anything would get released. 

With version control, we have a single source of truth. We might all still work on different modules but it enables us to collaborate better. 

![](Images/Day35_Git5.png)

Another thing to mention here is that it's not just developers that can benefit from Version Control, it's all members of the team to have visibility but also tools all having awareness or leverage, Project Management tools can be linked here, tracking the work. We might also have a build machine for example Jenkins which we will talk about in another module. A tool that Builds and Packages the system, automating the deployment tests and metrics. 

### What is Git? 

Git is a tool that tracks changes to source code or any file, or we could also say Git is an open-source distributed version control system. 

There are many ways in which git can be used on our systems, most commonly or at least for me I have seen it in at the command line, but we also have graphical user interfaces and tools like Visual Studio Code that have git aware operations we can take advantage of. 

Now we are going to run through a high-level overview before we even get Git installed on our local machine. 

Let's take the folder we created earlier. 

![](Images/Day35_Git2.png)

To use this folder with version control we first need to initiate this directory using the `git init command. For now, just think that this command puts our directory as a repository in a database somewhere on our computer. 

![](Images/Day35_Git6.png)

Now we can create some files and folders and our source code can begin or maybe it already has and we have something in here already. We can use the `git add .` command which puts all files and folders in our directory into a snapshot but we have not yet committed anything to that database. We are just saying all files with the `.` are ready to be added.   

![](Images/Day35_Git7.png)

Then we want to go ahead and commit our files, we do this with the `git commit -m "My First Commit"` command. We can give a reason for our commit and this is suggested so we know what has happened for each commit. 

![](Images/Day35_Git8.png)

現在我們可以透過 `git log` 指令, 知道專案的歷史和發生何事.

![](Images/Day35_Git9.png)

我們還能透過 `git status` 指令來檢查倉儲的狀態, 如下圖的前三行, 這表示我們還沒有提交檔案. 如果我們添加一個新檔 samplecode.ps1 , 再次輸入 `git status` 的時候, 會看到我們準備要提交的文件. 

![](Images/Day35_Git10.png)

Add our new file using the `git add samplecode.ps1` command and then we can run `git status` again and see our file is ready to be committed. 

![](Images/Day35_Git11.png)

Then issue `git commit -m "My Second Commit"` command.

![](Images/Day35_Git12.png)

Another `git status` now shows everything is clean again.

![](Images/Day35_Git13.png)

We can then use the `git log` command which shows the latest changes and first commit. 

![](Images/Day35_Git14.png)

If we wanted to see the changes between our commits i.e what files have been added or modified we can use the `git diff b8f8 709a`

![](Images/Day35_Git15.png)

Which then displays what has changed in our case we added a new file. 

![](Images/Day35_Git16.png)

We can also and we will go deeper into this later on but we can jump around our commits i.e we can go time travelling! By using our commit number we can use the `git checkout 709a` command to jump back in time without losing our new file. 

![](Images/Day35_Git17.png)

But then equally we will want to move forward as well and we can do this the same way with the commit number or you can see here we are using the `git switch -` command to undo our operation. 

![](Images/Day35_Git18.png)

摘要 (TLDR) 

- 追蹤專案的歷史
- 管理一個專案內不同版本的程式碼
- 在開發人員在開發人員和更廣泛的團隊和工具之間共享程式碼
- 協調團隊合作
- 哦，還有一些做時間旅行! 


這個章節的內容似乎跳的有些快, 希望在尚未解說 Git 指令的強大之前, 你能對版本控制的功能有大致上的了解. 

下一章,我們將會在你的電腦安裝和設置 Git ,並且更深入的解說 Git 的命令和使用案例, 以及 Git 能做到什麼.

## Resources 

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s) 
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg) 
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s) 
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)

See you on [第 36 天](day36.md) 

