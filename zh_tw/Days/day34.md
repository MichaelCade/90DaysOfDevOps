---
title: '#90DaysOfDevOps - Microsoft Azure Hands-On Scenarios - Day 34'
published: false
description: 90DaysOfDevOps - Microsoft Azure Hands-On Scenarios
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048763
---
## Microsoft Azure Hands-On Scenarios

The last 6 days have been focused on Microsoft Azure and the public cloud in general, a lot of this foundation had to contain a lot of theory to understand the building blocks of Azure but also this will nicely translate to the other major cloud providers as well. 

I mentioned at the very beginning about getting a foundational knowledge of the public cloud and choosing one provider to at least begin with, if you are dancing between different clouds then I believe you can get lost quite easily whereas choosing one you get to understand the fundamentals and when you have those it is quite easy to jump into the other clouds and accelerate your learning. 

In this final session, I am going to be picking and choosing my hands-on scenarios from this page here which is a reference created by Microsoft and is used for preparations for the [AZ-104 Microsoft Azure Administrator](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/) 

There are some here such as Containers and Kubernetes that we have not covered in any detail as of yet so I don't want to jump in there just yet. 

In previous posts, we have created most of Modules 1,2 and 3. 

### Virtual Networking 
Following [Module 04](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_04-Implement_Virtual_Networking.html):

I went through the above and changed a few namings for the purpose of #90DaysOfDevOps. I also instead of using the Cloud Shell went ahead and logged in with my new user created on previous days with the Azure CLI on my Windows machine. 

You can do this using the `az login` which will open a browser and let you authenticate to your account. 

I have then created a PowerShell script and some references from the module to use to build out some of the tasks below. You can find the associated files in this folder.
 (Cloud\01VirtualNetworking) 

 Please make sure you change the file location in the script to suit your environment. 

At this first stage we have no virtual network or virtual machines created in our environment, I only have a cloudshell storage location configured in my resource group. 

I first of all run my [PowerShell script](Cloud/01VirtualNetworking/Module4_90DaysOfDevOps.ps1)

 ![](Images/Day34_Cloud1.png)
 
- Task 1: Create and configure a virtual network

 ![](Images/Day34_Cloud2.png)

- Task 2: Deploy virtual machines into the virtual network

 ![](Images/Day34_Cloud3.png)

- Task 3: Configure private and public IP addresses of Azure VMs
  
 ![](Images/Day34_Cloud4.png)

- Task 4: Configure network security groups

![](Images/Day34_Cloud5.png)
![](Images/Day34_Cloud6.png)

- Task 5: Configure Azure DNS for internal name resolution

![](Images/Day34_Cloud7.png)
![](Images/Day34_Cloud8.png)

### Network Traffic Management 
Following [Module 06](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_06-Implement_Network_Traffic_Management.html):

Next walkthrough, from the last one we have gone into our resource group and deleted our resources, if you had not set up the user account like me to only have access to that one resource group you could follow the module changing the name to `90Days*` this will delete all resources and resource group. This will be my process for each of the following lab. 

For this lab I have also created a PowerShell script and some references from the module to use to build out some of the tasks below. You can find the associated files in this folder.
 (Cloud\02TrafficManagement) 


- Task 1: Provision the lab environment

I first of all run my [PowerShell script](Cloud/02TrafficManagement/Mod06_90DaysOfDevOps.ps1)

![](Images/Day34_Cloud9.png)

- Task 2: Configure the hub and spoke network topology

![](Images/Day34_Cloud10.png)

- Task 3: Test transitivity of virtual network peering

For this my 90DaysOfDevOps group did not have access to the Network Watcher because of permissions, I expect this is because Network Watchers are one of those resources that are not tied to a resource group which is where our RBAC was covered for this user. I added the East US Network Watcher contributer role to the 90DaysOfDevOps group. 

![](Images/Day34_Cloud11.png)
![](Images/Day34_Cloud12.png)
![](Images/Day34_Cloud13.png)

^  This is expected, since the two spoke virtual networks are not peered with each other (virtual network peering is not transitive).

- Task 4: Configure routing in the hub and spoke topology

I had another issue here with my account not being able to run the script as my user within the group 90DaysOfDevOps which I am unsure of so I did jump back into my main admin account. The 90DaysOfDevOps group is an owner of everything in the 90DaysOfDevOps Resource Group so would love to understand why I cannot run a command inside the VM? 

![](Images/Day34_Cloud14.png)
![](Images/Day34_Cloud15.png)

I then was able to go back into my michael.cade@90DaysOfDevOps.com account and continue this section. Here we are running the same test again but now with the result being reachable. 

![](Images/Day34_Cloud16.png)

- Task 5: Implement Azure Load Balancer

![](Images/Day34_Cloud17.png)
![](Images/Day34_Cloud18.png)

- Task 6: Implement Azure Application Gateway

![](Images/Day34_Cloud19.png)
![](Images/Day34_Cloud20.png)

### Azure Storage 
Following [Module 07](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_07-Manage_Azure_Storage.html):

For this lab I have also created a PowerShell script and some references from the module to use to build out some of the tasks below. You can find the associated files in this folder.
 (Cloud\03Storage) 

- Task 1: Provision the lab environment

I first of all run my [PowerShell script](Cloud/03Storage/Mod07_90DaysOfDeveOps.ps1)

![](Images/Day34_Cloud21.png)

- Task 2: Create and configure Azure Storage accounts

![](Images/Day34_Cloud22.png)

- Task 3: Manage blob storage

![](Images/Day34_Cloud23.png)

- Task 4: Manage authentication and authorization for Azure Storage

![](Images/Day34_Cloud24.png)
![](Images/Day34_Cloud25.png)

I was a little impatient waiting for this to be allowed but it did work eventually. 

![](Images/Day34_Cloud26.png)


- Task 5: Create and configure an Azure Files shares

On the run command this would not work with michael.cade@90DaysOfDevOps.com so I used my elevated account. 

![](Images/Day34_Cloud27.png)
![](Images/Day34_Cloud28.png)
![](Images/Day34_Cloud29.png)


- Task 6: Manage network access for Azure Storage

![](Images/Day34_Cloud30.png)

### Serverless (Implement Web Apps)
Following [Module 09a](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_09a-Implement_Web_Apps.html):


- Task 1: Create an Azure web app

![](Images/Day34_Cloud31.png)

- Task 2: Create a staging deployment slot

![](Images/Day34_Cloud34.png)

- Task 3: Configure web app deployment settings

![](Images/Day34_Cloud33.png)

- Task 4: Deploy code to the staging deployment slot

![](Images/Day34_Cloud32.png)

- Task 5: Swap the staging slots

![](Images/Day34_Cloud35.png)

- Task 6: Configure and test autoscaling of the Azure web app

This script I am using can be found in (Cloud/05Serverless)

![](Images/Day34_Cloud36.png)

This wraps up the section on Microsoft Azure and the public cloud in general. I will say that I had lots of fun attacking and working through this scenarios. 

## Resources 

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Next we will be diving into version control systems, specifically around git and then also code repository overviews and we will be choosing GitHub as this is my preferred option. 

See you on [Day 35](day35.md) 
