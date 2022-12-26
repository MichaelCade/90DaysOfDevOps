## The Big Picture: DevOps and Networking

As with all sections, I am using open and free training materials and a lot of the content can be attributed to others. In the case of the networking section a large majority of the content shown is from [Practical Networking](https://www.practicalnetworking.net/)'s free [Networking Fundamentals series](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi).  It is mentioned in the resources as well as a link but it's appropriate to highlight this as from a community point of view, I have leveraged this course to help myself understand more about particular areas of technologies. This repository is a repository for my note taking and enabling the community to hopefully benefit from this and the listed resources. 

Welcome to Day 21! We are going to be getting into Networking over the next 7 days, Networking and DevOps are the overarching themes but we will need to get into some of the networking fundamentals as well.

Ultimately as we have said previously DevOps is about a culture and process change within your organisation this as we have discussed can be Virtual Machines, Containers, or Kubernetes but it can also be the network, If we are using those DevOps principles for our infrastructure that has to include the network more to the point from a DevOps point of view you also need to know about the network as in the different topologies and networking tools and stacks that we have available.

I would argue that we should have our networking devices configured using infrastructure as code and have everything automated like we would our virtual machines, but to do that we have to have a good understanding of what we are automating.

### What is NetDevOps | Network DevOps?

You may also hear the terms Network DevOps or NetDevOps. Maybe you are already a Network engineer and have a great grasp on the network components within the infrastructure you understand the elements used around networking such as DHCP, DNS, NAT etc. You will also have a good understanding of the hardware or software-defined networking options, switches, routers etc.

But if you are not a network engineer then we probably need to get foundational knowledge across the board in some of those areas so that we can understand the end goal of Network DevOps.

But in regards to those terms, we can think of NetDevOps or Network DevOps as applying the DevOps Principles and Practices to the network, applying version control and automation tools to the network creation, testing, monitoring, and deployments.

If we think of Network DevOps as having to require automation, we mentioned before about DevOps breaking down the silos between teams. If the networking teams do not change to a similar model and process then they become the bottleneck or even the failure overall.

Using the automation principles around provisioning, configuration, testing, version control and deployment is a great start. Automation is overall going to enable speed of deployment, stability of the networking infrastructure and consistent improvement as well as the process being shared across multiple environments once they have been tested. Such as a fully tested Network Policy that has been fully tested on one environment can be used quickly in another location because of the nature of this being in code vs a manually authored process which it might have been before.
A really good viewpoint and outline of this thinking can be found here. [Network DevOps](https://www.thousandeyes.com/learning/techtorials/network-devops)

## Networking The Basics

Let's forget the DevOps side of things to begin with here and we now need to look very briefly into some of the Networking fundamentals.

### Network Devices

If you prefer this content in video form, check out these videos from Practical Networking:

* [Network Devices - Hosts, IP Addresses, Networks - Networking Fundamentals - Lesson 1a](https://www.youtube.com/watch?v=bj-Yfakjllc&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=1)
* [Network Devices - Hub, Bridge, Switch, Router - Networking Fundamentals - Lesson 1b
](https://www.youtube.com/watch?v=H7-NR3Q3BeI&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=2)

**Host** are any devices which send or receive traffic.

![](Images/Day21_Networking1.png)

**IP Address** the identity of each host.

![](Images/Day21_Networking2.png)

**Network** is what transports traffic between hosts. If we did not have networks there would be a lot of manual movement of data!

A logical group of hosts which require similar connectivity.

![](Images/Day21_Networking3.png)

**Switches** facilitate communication **_within_** a network. A switch forwards data packets between hosts. A switch sends packets directly to hosts.

- Network: A Grouping of hosts which require similar connectivity.
- Hosts on a Network share the same IP address space.

![](Images/Day21_Networking4.png)

**Router** facilitates communication between networks. As we said before that a switch looks after communication within a network a router allows us to join these networks together or at least give them access to each other if permitted.

A router can provide a traffic control point (security, filtering, redirecting) More and more switches also provide some of these functions now.

Routers learn which networks they are attached to. These are known as routes, a routing table is all the networks a router knows about.

A router has an IP address in the networks they are attached to. This IP is also going to be each host's way out of their local network also known as a gateway.

Routers also create the hierarchy in networks I mentioned earlier.

![](Images/Day21_Networking5.png)

## Switches vs Routers

**Routing** is the process of moving data between networks.

- A router is a device whose primary purpose is Routing.

**Switching** is the process of moving data within networks.

- A Switch is a device whose primary purpose is switching.

This is very much a foundational overview of devices as we know there are many different Network Devices such as:

- Access Points
- Firewalls
- Load Balancers
- Layer 3 Switches
- IDS / IPS
- Proxies
- Virtual Switches
- Virtual Routers

Although all of these devices are going to perform Routing and/or Switching.

Over the next few days, we are going to get to know a little more about this list.

- OSI Model
- Network Protocols
- DNS (Domain Name System)
- NAT
- DHCP
- Subnets

## Resources

* [Networking Fundamentals](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)
* [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)

See you on [Day22](day22.md)
