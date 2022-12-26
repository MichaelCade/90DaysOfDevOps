---
title: '#90DaysOfDevOps - Testing, Tools & Alternatives - Day 62'
published: false
description: '90DaysOfDevOps - Testing, Tools & Alternatives'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049053
---

## Testing, Tools & Alternatives

As we close out this section on Infrastructure as Code we must mention testing our code, the various tools available and then some of the alternatives to Terraform to achieve this. As I said at the start of the section my focus was on Terraform because it is firstly free and open source, secondly, it is cross-platform and agnostic to environments. But there are also alternatives out there that should be considered but the overall goal is to make people aware that this is the way to deploy your infrastructure.

### Code Rot

The first area I want to cover in this session is code rot, unlike application code, infrastructure as code might get used and then not for a very long time. Let's take the example that we are going to be using Terraform to deploy our VM environment in AWS, perfect and it works the first time and we have our environment, but this environment doesn't change too often so the code gets left the state possibly or hopefully stored in a central location but the code does not change.

What if something changes in the infrastructure? But it is done out of band, or other things change in our environment.

- Out of band changes
- Unpinned versions
- Deprecated dependencies
- Unapplied changes

### Testing

Another huge area that follows on from code rot and in general is the ability to test your IaC and make sure all areas are working the way they should.

First up there are some built-in testing commands we can take a look at:

| Command              | Description                                                                                |
| -------------------- | ------------------------------------------------------------------------------------------ |
| `terraform fmt`      | Rewrite Terraform configuration files to a canonical format and style.                     |
| `terraform validate` | Validates the configuration files in a directory, referring only to the configuration      |
| `terraform plan`     | Creates an execution plan, which lets you preview the changes that Terraform plans to make |
| Custom validation    | Validation of your input variables to ensure they match what you would expect them to be   |

We also have some testing tools available external to Terraform:

- [tflint](https://github.com/terraform-linters/tflint)

  - Find possible errors
  - Warn about deprecated syntax and unused declarations.
  - Enforce best practices, and naming conventions.

Scanning tools

- [checkov](https://www.checkov.io/) - scans cloud infrastructure configurations to find misconfigurations before they're deployed.
- [tfsec](https://aquasecurity.github.io/tfsec/v1.4.2/) - static analysis security scanner for your Terraform code.
- [terrascan](https://github.com/accurics/terrascan) - static code analyser for Infrastructure as Code.
- [terraform-compliance](https://terraform-compliance.com/) - a lightweight, security and compliance-focused test framework against terraform to enable the negative testing capability for your infrastructure-as-code.
- [snyk](https://docs.snyk.io/products/snyk-infrastructure-as-code/scan-terraform-files/scan-and-fix-security-issues-in-terraform-files) - scans your Terraform code for misconfigurations and security issues

Managed Cloud offering

- [Terraform Sentinel](https://www.terraform.io/cloud-docs/sentinel) - embedded policy-as-code framework integrated with the HashiCorp Enterprise products. It enables fine-grained, logic-based policy decisions, and can be extended to use information from external sources.

Automated testing

- [Terratest](https://terratest.gruntwork.io/) - Terratest is a Go library that provides patterns and helper functions for testing infrastructure

Worth a mention

- [Terraform Cloud](https://cloud.hashicorp.com/products/terraform) - Terraform Cloud is HashiCorp’s managed service offering. It eliminates the need for unnecessary tooling and documentation for practitioners, teams, and organizations to use Terraform in production.

- [Terragrunt](https://terragrunt.gruntwork.io/) - Terragrunt is a thin wrapper that provides extra tools for keeping your configurations DRY, working with multiple Terraform modules, and managing remote state.

- [Atlantis](https://www.runatlantis.io/) - Terraform Pull Request Automation

### Alternatives

We mentioned on Day 57 when we started this section that there were some alternatives and I very much plan on exploring this following on from this challenge.

| Cloud Specific                  | Cloud Agnostic |
| ------------------------------- | -------------- |
| AWS CloudFormation              | Terraform      |
| Azure Resource Manager          | Pulumi         |
| Google Cloud Deployment Manager |                |

I have used AWS CloudFormation probably the most out of the above list and native to AWS but I have not used the others other than Terraform. As you can imagine the cloud-specific versions are very good in that particular cloud but if you have multiple cloud environments then you are going to struggle to migrate those configurations or you are going to have multiple management planes for your IaC efforts.

I think an interesting next step for me is to take some time and learn more about [Pulumi](https://www.pulumi.com/)

From a Pulumi comparison on their site

> "Both Terraform and Pulumi offer the desired state infrastructure as code model where the code represents the desired infrastructure state and the deployment engine compares this desired state with the stack’s current state and determines what resources need to be created, updated or deleted."

The biggest difference I can see is that unlike the HashiCorp Configuration Language (HCL) Pulumi allows for general-purpose languages like Python, TypeScript, JavaScript, Go and .NET.

A quick overview [Introduction to Pulumi: Modern Infrastructure as Code](https://www.youtube.com/watch?v=QfJTJs24-JM) I like the ease and choices you are prompted with and want to get into this a little more.

This wraps up the Infrastructure as code section and next we move on to that little bit of overlap with configuration management in particular as we get past the big picture of configuration management we are going to be using Ansible for some of those tasks and demos.

## Resources

I have listed a lot of resources down below and I think this topic has been covered so many times out there, If you have additional resources be sure to raise a PR with your resources and I will be happy to review and add them to the list.

- [What is Infrastructure as Code? Difference of Infrastructure as Code Tools](https://www.youtube.com/watch?v=POPP2WTJ8es)
- [Terraform Tutorial | Terraform Course Overview 2021](https://www.youtube.com/watch?v=m3cKkYXl-8o)
- [Terraform explained in 15 mins | Terraform Tutorial for Beginners](https://www.youtube.com/watch?v=l5k1ai_GBDE)
- [Terraform Course - From BEGINNER to PRO!](https://www.youtube.com/watch?v=7xngnjfIlK4&list=WL&index=141&t=16s)
- [HashiCorp Terraform Associate Certification Course](https://www.youtube.com/watch?v=V4waklkBC38&list=WL&index=55&t=111s)
- [Terraform Full Course for Beginners](https://www.youtube.com/watch?v=EJ3N-hhiWv0&list=WL&index=39&t=27s)
- [KodeKloud - Terraform for DevOps Beginners + Labs: Complete Step by Step Guide!](https://www.youtube.com/watch?v=YcJ9IeukJL8&list=WL&index=16&t=11s)
- [Terraform Simple Projects](https://terraform.joshuajebaraj.com/)
- [Terraform Tutorial - The Best Project Ideas](https://www.youtube.com/watch?v=oA-pPa0vfks)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)
- [Pulumi - IaC in your favorite programming language!](https://www.youtube.com/watch?v=vIjeiDcsR3Q&t=51s)

See you on [Day 63](day63.md)
