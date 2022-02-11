## The Big Picture - Containers 

We are now starting the next section and this section is going to be focused on containers and in particular we are likely going to be looking into Docker and get into some of the key areas to understand more about Containers. 

I will also be trying to get some hands on here to create our own container that we can use later on during this section but also future sections later on in the challenge. 

As always this first post is going to be focused on the big picture of how we got here and what it all means. 

#History of platforms and application development
#do we want to talk about Virtualisation & Containerisation 

### Why another way to run applications? 

The first thing we have to take a look at is why do we need another way to run our software or applications, well it is just that choice is great, we can run our applications in many different forms, we might see applications deployed on physical hardware with an operating system and a single application deployed there, we might see virtual machine or cloud based IaaS instances running our application which then integrate into a database again in a VM or as PaaS offering in the public cloud. Or we might see our applications running in containers. 

None of the above options are wrong or right, but they each have their reasons to exist and I also strongly believe that none of these are going away. I have seen a lot of content that walks into Containers vs Virtual Machines and there really should not be an argument as that is more like apples vs pears argument where they are both fruit (ways to run our applications) but they are not the same. 

I would also say that if you were starting out and you were developing an application you should lean towards containers simply because and we will get into some of these areas later, but really its about effciency, speed and size. But that also comes with a price, if you have no idea about containers then its going to be a learning curve to force yourself to understand the why and get into that mindset. If you have developed your applications a particular way or you are not in a green field environment then you might have more pain points to deal with before even considering containers. 

We have many different choices then when it comes to downloading a given piece of software, there are a variety of different operating systems that we might be using. And specific instructions for what we need to do to install our applications. More and more recently I am finding that the applications we might have once needed a full server OS, A VM, Physical or cloud instance are now releasing container based versions of their software. I find this massively interesting as this opens the world of containers and then Kubernetes to everyone and not just a focus on application developers. 

As you can probably tell as I have said before, I am not going to advocate that the answer is containers whats the question! But I would like to discuss how this is another option for us to be aware of when we deploy our applications. 

We have had container technology for a long time now, so why now over the last say 10 years has this become popular, and I would say even more popular in the last 5 when we have had containers for decades. It comes down to I believe the challenge containers or I should say images as well to how we distribute our software, because if we just have container technology, then we still will have many of the same problems we've had with software management. 

If we think about Docker as a tool, the reason that it took off, is because of the ecosystem of images that are easy to find and use. Simple to get on your systems and get up and running. another major part to this is consistency across this entire space, of all these different challenges that we face with software. It doesn't matter if it's MongoDB or nodeJS, the process to get either of those up and running will be exactly the same. The process to stop either of those is exactly the same. All of these issues will still exist, but the nice thing is, when we bring good container and image technology together, we now have a single set of tools to help us tackle all of these different problems.

Ok, probably sounds great and exciting but we still need to understand what is a container and now i have mentioned images so let's cover those areas next. 

### What is a container? 

When we run applications on our computer, the web browser or VScode that you are using to read this post. That application is running as a process or what is known as a process. On our laptops or systems we tend to run multiple applications or as we said processes. When we open a new application or click on our application icon this is an application we would like to run, sometimes this application might be a service that we just want to run in the background, our operating system is full of services that are running in the background providing you with the user experience you get with your system. 

That application icon represents a link to an executable so that the operating system can load that executable into memory. Interestingly, that executable is sometimes referred to as an image when we're talking about a process. I also just mentioned image in the last section when it comes to how and why containers and images combined made containers popular in our ecosystem. 





### What is an Image? 




### Where to get started? 




### What does a Container Infrastructure look like? 



## Resources 

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)

See you on [Day 43](day43.md) 