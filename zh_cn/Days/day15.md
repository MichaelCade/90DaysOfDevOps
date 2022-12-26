---
title: '#90DaysOfDevOps - Linux Commands for DevOps (Actually everyone) - Day 15'
published: false
description: 90DaysOfDevOps - Linux Commands for DevOps (Actually everyone)
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048834
---
## Linux Commands for DevOps (Actually everyone)

I mentioned it [yesterday](day14.md) that we are going to be spending a lot of time in the terminal with some commands to get stuff done. 

I also mentioned that with our vagrant provisioned VM we can use `vagrant ssh` and gain access to our box. You will need to be in the same directory as we provisioned it from.

For SSH you won't need the username and password, you will only need that if you decide to login to the Virtual Box console. 

This is where we want to be as per below: 

![](Images/Day15_Linux1.png)

## Commands 

Obviously I cannot cover all the commands here, there are  pages and pages of documentation that cover these but also if you are ever in your terminal and you just need to understand options to a specific command we have the `man` pages short for manual. We can use this to go through each of the commands we touch on during this post to find out more options for each one. We can run `man man` which will give you the help for manual pages. To escape the man pages you should press `q` for quit. 

![](Images/Day15_Linux2.png)
![](Images/Day15_Linux3.png)

`sudo` If you are familar with Windows and the right click `run as administrator` we can think of `sudo` as very much this. When you run a command with this command you will be running it as `root` it will prompt you for the password before running the command. 

![](Images/Day15_Linux4.png)

For one off jobs like installing applications or services you might need that `sudo command` but what if you have several tasks to deal with and you want to live as `sudo` for a while? This is where you can use `sudo su` again the same as `sudo` once entered you will be prompted for your `root` password. In a test VM like ours this is fine but I would find it very hard for us to be rolling around as `root` for prolonged periods, bad things can happen. To get out of this elevated position you simply type in `exit` 

![](Images/Day15_Linux5.png)

I find myself using `clear` all the time, the `clear` command does exactly what it says it is going to clear the screen of all previous commands, putting your prompt to the top and giving you a nice clean workspace. Windows I think is `cls` in the .mdprompt. 

![](Images/Day15_Linux6.png)

Let's now look at some commands where we can actually create things within our system and then visualise them in our terminal, first of all we have `mkdir` this will allow us to create a folder in our system. With the following command we can create a folder in our home directory called Day15 `mkdir Day15`

![](Images/Day15_Linux7.png)

With `cd` this allows us to change directory, so for us to move into our newly created directory we can do this with `cd Day15` tab can also be used to autocomplete the directory available. If we want to get back to where we started we can use `cd ..`

![](Images/Day15_Linux8.png)

`rmdir` allows for us to remove the directory, if we run `rmdir Day15` then the folder will be removed (note that this will only work if you have nothing in the folder)

![](Images/Day15_Linux9.png)

I am sure we have all done it where we have navigated to the depths of our file system to a directory and not known where we are. `pwd` gives us the print out of the working directory, pwd as much as it looks like password it stands for print working directory. 

![](Images/Day15_Linux10.png)

We know how to create folders and directories but how do we create files? We can create files using the `touch` command if we were to run `touch Day15` this would create a file. Ignore `mkdir` we are going see this again later.

![](Images/Day15_Linux11.png)

`ls` I can put my house on this, you will use this command so many times, this is going to list the all the files and folders in the current directory. Let's see if we can see that file we just created. 

![](Images/Day15_Linux12.png)

How can we find files on our Linux system? `locate` is going to allow us to search our file system. If we use `locate Day15` it will report back that location of the file. Bonus round is that if you know that the file does exist but you get a blank result then run `sudo updatedb` which will index all the files in the file system then run your `locate` again. If you do not have `locate` available to you, you can install it using this command `sudo apt install mlocate` 

![](Images/Day15_Linux13.png)

What about moving files from one location to another? `mv` is going to allow you to move your files. Example `mv Day15 90DaysOfDevOps` will move your file to the 90DaysOfDevOps folder. 

![](Images/Day15_Linux14.png)

We have moved our file but what if we want to rename it now to something else? We can do that using the `mv` command again... WOT!!!? yep we can simply use `mv Day15 day15` to change to upper case or we could use `mv day15 AnotherDay` to change it altogether, now use `ls` to check the file. 

![](Images/Day15_Linux15.png)

Enough is enough, let's now get rid (delete)of our file and maybe even our directory if we have one created. `rm` simply `rm AnotherDay` will remove our file. We will also use quite a bit `rm -R` which will recursively work through a folder or location. We might also use `rm -R -f` to force the removal of all of those files. Spoiler if you run `rm -R -f /` add sudo to it and you can say goodbye to your system....! 

![](Images/Day15_Linux16.png)

We have looked at moving files around but what if I just want to copy files from one folder to another, simply put its very similar to the `mv` command but we use `cp` so we can now say `cp Day15 Desktop` 

![](Images/Day15_Linux17.png)

We have created folders and files but we haven't actually put any contents into our folder, we can add contents a few ways but an easy way is `echo` we can also use `echo` to print out a lot of things in our terminal, I personally use echo a lot to print out system variables to know if they are set or not at least. we can use `echo "Hello #90DaysOfDevOps" > Day15` and this will add this to our file. We can also append to our file using `echo "Commands are fun!" >> Day15` 

![](Images/Day15_Linux18.png)

Another one of those commands you will use a lot! `cat` short for concatenate. We can use `cat Day15` to see the contents inside the file. Great for quickly reading those configuration files. 

![](Images/Day15_Linux19.png)

If you have a long complex configuration file and you want or need to find something fast in that file vs reading every line then `grep` is your friend, this will allow us to search your file for a specific word using `cat Day15 | grep "#90DaysOfDevOps"`

![](Images/Day15_Linux20.png)

If you are like me and you use that `clear` command a lot then you might miss some of the commands previously ran, we can use `history` to find out all those commands we have run prior. `history -c` will remove the history. 

When you run `history` and you would like to pick a specific command you can use `!3` to choose the 3rd command in the list. 

You are also able to use `history | grep "Command` to search for something specific. 

On servers to trace back when was a command executed, it can be useful to append the date and time to each command in the history file.

The following system variable controls this behaviour:
```
HISTTIMEFORMAT="%d-%m-%Y %T "
```
You can easily add to your bash_profile:
```
echo 'export HISTTIMEFORMAT="%d-%m-%Y %T "' >> ~/.bash_profile
```
So as useful to allow the history file grow bigger:

```
echo 'export HISTSIZE=100000' >> ~/.bash_profile
echo 'export HISTFILESIZE=10000000' >> ~/.bash_profile
```

![](Images/Day15_Linux21.png)

Need to change your password? `passwd` is going allow us to change our password. Note that when you add your password in like this when it is hidden it will not be shown in `history` however if your command has `-p PASSWORD` then this will be visible in your `history`. 

![](Images/Day15_Linux22.png)

We might also want to add new users to our system, we can do this with `useradd` we have to add the user using our `sudo` command, we can add a new user with `sudo useradd NewUser`

![](Images/Day15_Linux23.png)

Creating a group again requires `sudo` and we can use `sudo groupadd DevOps` then if we want to add our new user to that group we can do this by running `sudo usermod -a -G DevOps` `-a` is add and `-G` is group name. 

![](Images/Day15_Linux24.png)

How do we add users to the `sudo` group, this would be a very rare occassion for this to happen but in order to do this it would be `usermod -a -G sudo NewUser` 

### Permissions 

read, write and execute are the permissions we have on all of our files and folders on our Linux system. 

A full list: 

- 0 = None `---`
- 1 = Execute only `--X`
- 2 = Write only `-W-`
- 3 = Write & Exectute `-WX`
- 4 = Read Only `R--`
- 5 = Read & Execute `R-X`
- 6 = Read & Write `RW-`
- 7 = Read, Write & Execute `RWX`

You will also see `777` or `775` and these represent the same numbers as the list above but each one represents **User - Group - Everyone**

Let's take a look at our file. `ls -al Day15` you can see the 3 groups mentioned above, user and group has read & write but everyone only has read. 

![](Images/Day15_Linux25.png)

We can change this using `chmod` you might find yourself doing this if you are creating binaries a lot on your systems as well and you need to give the ability to execute those binaries. `chmod 750 Day15` now run `ls -al Day15` if you want to run this for a whole folder then you can use `-R` to recursively do that. 

![](Images/Day15_Linux26.png)

What about changing the owner of the file? We can use `chown` for this operation, if we wanted to change the ownership of our `Day15` from user `vagrant` to `NewUser` we can run `sudo chown NewUser Day15` again `-R` can be used. 

![](Images/Day15_Linux27.png)

A command that you will come across is `awk` where this comes in real use is when you have an output that you only need specific data from. like running `who` we get lines with information, but maybe we only need the names. We can run `who | awk '{print $1}'` to get just a list of that first column. 

![](Images/Day15_Linux28.png)

If you are looking to read streams of data from standard input, then generates and executes command lines; meaning it can take output of a command and passes it as argument of another command. `xargs` is a useful tool for this use case. If for example I want a list of all the Linux user accounts on the system I can run. `cut -d: -f1 < /etc/passwd` and get the long list we see below. 

![](Images/Day15_Linux29.png)

If I want to compact that list I can do so by using `xargs` in a command like this `cut -d: -f1 < /etc/passwd | sort | xargs` 

![](Images/Day15_Linux30.png)

I didn't mention the `cut` command either, this allows us to remove sections from each line of a file. It can be used to cut parts of a line by byte position, character and field. The `cut -d " " -f 2 list.txt` command allows us to remove that first letter we have and just display our numbers. There are so many combinations that can be used here with this command, I am sure I have spent too much time trying to use this command when I could have extracted data quicker manually. 

![](Images/Day15_Linux31.png)

Also to note if you type a command and you are no longer with happy with it and you want to start again just hit control + c and this will cancel that line and start you fresh. 

## Resources 

- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

See you on [Day16](day16.md)

This is a pretty heavy list already but I can safely say that I have used all of these commands in my day to day, be it from an administering Linux servers or in my Linux Desktop, it is very easy when you are in Windows or macOS to navigate the UI but in Linux Servers they are not there, everything is done through the terminal. 


















