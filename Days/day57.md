## An intro to Terraform 

"Terraform is a tool for building, changing, and versioning infrastructure safely and efficiently" 

The above quote is from HashiCorp, HashiCorp is the company behind Terraform. 

"Terraform is an open-source infrastructure as code software tool that provides a consistent CLI workflow to manage hundreds of cloud services. Terraform codifies cloud APIs into declarative configuration files."

HashiCorp have a great resource in [HashiCorp Learn](https://learn.hashicorp.com/terraform?utm_source=terraform_io&utm_content=terraform_io_hero) which covers all of their products and gives some great walkthrough demos when you are trying to achieve something with Infrastructure as Code. 

All cloud providers and on prem platforms generally give us access to management consoles which enables us to create our resources via a UI, generally these platforms also provide a CLI or API access to also create the same resources but with an API we have the ability to provision fast. 

Infrastructure as Code allows us to hook into those APIs to deploy our resources in a desired state. 

Other tools but not exclusive or exhaustive below. If you have other tools then please share via a PR.  

| Cloud Specific                  | Cloud Agnostic | 
| ------------------------------- | -------------- |
| AWS CloudFormation              | Terraform      | 
| Azure Resource Manager          | Pulumi         | 
| Google Cloud Deployment Manager |                | 

This is another reason why we are using Terraform, we want to be agnostic to the clouds and platforms that we wish to use for our demos but also in general. 

## Terraform Overview 

Terraform is a provisioning focused tool, 




## Terraform Installation 