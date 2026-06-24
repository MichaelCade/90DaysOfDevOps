# Day 52: Identity and Access Management (IAM)

As cloud computing continues to gain popularity, more and more organizations are turning to cloud platforms to manage their infrastructure. However, with this comes the need to ensure proper security measures are in place to protect data and resources. One of the most critical tools for managing security in AWS is Identity and Access Management (IAM).

## What is AWS IAM?
|![](images/day52-1.png)| 
|:-:|
| <i>IAM is (1) WHO (2) CAN ACCESS (3) WHAT</i>|


AWS IAM is a web service that allows you to manage users and their access to AWS resources. With IAM, you can create and manage AWS users and groups, control access to AWS resources, and set permissions that determine what actions users can perform on those resources. IAM provides fine-grained access control, which means that you can grant or deny permissions to specific resources at a granular level.

IAM is an essential tool for securing your AWS resources. Without it, anyone with access to your AWS account would have unrestricted access to all your resources. With IAM, you can control who has access to your resources, what actions they can perform, and what resources they can access. IAM also enables you to create and manage multiple AWS accounts, which is essential as large organizations will always have many accounts that will need some level of interaction with each other:

|![](images/day52-2.png)|
|:-:|
| <i>Multi-Account IAM access is essential knowledge</i>|


## How to Get Started with AWS IAM

Getting started with AWS IAM is straightforward. Here are the steps you need to follow:

### Step 1: Create an AWS Account

The first step is to create an AWS account if you don't already have one. We did this on day 50 so you should be good to go 😉

### Step 2: Set up IAM

Once you have an AWS account, you can set up IAM by navigating to the IAM console. The console is where you'll manage IAM users, groups, roles, and policies. 

### Step 3: Create an IAM User

The next step is to create an IAM user. An IAM user is an entity that you create in IAM that represents a person or service that needs access to your AWS resources. When you create an IAM user, you can specify the permissions that the user should have. One of the homework assignments from Day 50 was to [Create an IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html), if you haven't completed that go back and make one now. 

### Step 4: Create an IAM Group

After you've created an IAM user, the next step is to create an IAM group. An IAM group is a collection of IAM users. When you create an IAM group, you can specify the permissions that the group should have. Watch "IAM Basics" and read "IAM User Guide:Getting Started" in the resources section to accomplish this.

### Step 5: Assign Permissions to the IAM Group

Once you've created an IAM group, you can assign permissions to the group. This involves creating an IAM policy that defines the permissions that the group should have. You can then attach the policy to the group. Watch "IAM Tutorial & Deep Dive" and go through the IAM Tutorial in the resources section to accomplish this.

### Step 6: Test the IAM User

After you've assigned permissions to the IAM group, you can test the IAM user to ensure that they have the correct permissions. To do this, you can log in to the AWS Management Console using the IAM user's credentials and attempt to perform the actions that the user should be able to perform.

## Resources:
[IAM Basics](https://youtu.be/iF9fs8Rw4Uo)

[IAM User Guide: Getting started](https://docs.aws.amazon.com/IAM/latest/UserGuide/getting-started.html)

[IAM Video Tutorial & Deep Dive](https://youtu.be/ExjW3HCFG1U)

[IAM Tutorial: Delegate access across AWS accounts using IAM roles](https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)

See you in [Day 53](day53.md).
