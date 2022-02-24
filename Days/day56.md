## Understanding Infrastructure as a Code aka IaC and Terraform Basics

An Infrastructure in common language would be considered as the final output obtained after processing multiple different kinds of inputs. </br>

For instance: An Infrastructure for a car mainly includes Wheels, Engine, Steering, Clutch, brake, Accelerator and so on. </br>

In similar way, for creating a cloud infrastructure, there is also some of its main building blocks,</br>


# What is terraform? </br>

Terraform is an infrastructure as code (IaC) tool that allows you to build, change and version infrastructure safely and efficiently. </br>
This includes low-level components such as compute instances, storage, and networking, as well as high-level components such as DNS entries, SaaS features, etc. </br>
- Terraform can manage both existing service providers and custom in-house solutions. 
- For writing configuration in terraform we use HCL.
- Hashicorp Configuration language (HCL)
- HCL is a toolkit for creating structured configuration languages that are both human and machine-friendly, for use with command-line tools.
- Although intended to be generally useful, it is primarily targeted towards devops tools, servers, etc.

#### A simple HCL code would look something like this,

```
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }
  required_version = ">= 0.14.9"
}
provider "aws" {
  profile = "default"
  region  = "us-west-2"
}
resource "aws_instance" "app_server" {
  ami           = "ami-830c94e3"
  instance_type = "t2.micro"

  tags = {
    Name = "MyFirstTerraformConfiguration"
  }
}
```

#### Key Components of Cloud Infrastructure
- Compute

Compute services are for running diverse workloads on the AWS Compute platform.
Most importantly it helps to lower the infrastructure costs and accelerate innovation on the world’s most reliable, secure, and capable cloud.

We are mainly going to use EC2 services in our upcoming blogs.

- EC2

Amazon Elastic Compute Cloud (Amazon EC2) is a web service that provides secure, resizable compute capacity in the cloud. It is designed to make web-scale computing easier for developers.
resource "aws_instance" "example" {
  ami           = var.AMIS[var.AWS_REGION]
  instance_type = "t2.micro"
}

- Network and Content Delivery
Use cases for Network and Content Delivery </br>

- Networking foundations : It helps us to quickly set up, secure, and monitor our network.
- Application networking : Provides traditional and modern applications with improved security, availability, performance, and streamlined monitoring.
- Edge networking : Deliver your data with single-digit millisecond latency.
- Hybrid connectivity : Create fast, secure, and reliable connections between your on-premises and AWS networks.


We are mainly going to use VPC services in our upcoming blogs. </br>

- VPC
VPC stands for Virtual Private Cloud, using VPC we define and provision a logically isolated network for our AWS resources.
```
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_support   = "true"
  enable_dns_hostnames = "true"
  enable_classiclink   = "false"
  tags = {
    Name = "main"
  }
}
```

- Storage

Storage is a complete range of services for us to store, access, govern, and analyse our data to reduce costs, increase agility, and accelerate innovation.
With using cloud storage we can reduce the storage costs.

- S3

Amazon Simple Storage Service (Amazon S3) is an object storage service that offers industry-leading scalability, data availability, security, and performance. </br>
```
 resource "aws_s3_bucket" "b" {
  bucket = "nitin1455"
}
```

- Security, Identity and Compliance
Security, Identity and Compliance are for securing our workloads and applications in the cloud. </br>

#### Some of it’s use cases are :

- Data protection
- Identity & access management
- Network & application protection
- Threat detection & continuous monitoring
- Compliance & data privacy

- IAM

IAM is for securely managing access to services and resources.

Using IAM, we can create and manage AWS users and groups, and use permissions to allow and deny their access to AWS resources.
```
resource "aws_iam_group" "administrators" {
  name = "administrators"
}
```
```
resource "aws_iam_policy_attachment" "administrators-attach" {
  name       = "administrators-attach"
  groups     = [aws_iam_group.administrators.name]
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}
```
# user
```
resource "aws_iam_user" "admin1" {
  name = "admin1"
}
```
