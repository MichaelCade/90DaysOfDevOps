# Day 47 - Automation with Python

Using Python for infrastructure management involves automating the management of IT infrastructure, such as servers, databases, and networking equipment. This can include tasks like provisioning, configuration, and orchestration.
Python is a popular language for infrastructure management, and there are several tools and libraries available to help with this. Some popular tools for infrastructure management that use Python include:

- Fabric
- PyWinRM
- Pulumi

## Fabric

Fabric is a Python library that can be used for streamlining SSH commands and remote execution of scripts, which can be used to automate server configuration and deployment.
Here is an example in which we will be using the Fabric library to connect to a remote server using SSH, and then run the `ls -l` command on the remote server. The output of the command will be printed to the console.

``` python
from fabric import Connection

# Connect to the remote server
c = Connection('user@remote.host')

# Run a command on the remote server
result = c.run('ls -l')

# Print the output of the command
print(result.stdout)
```


## PyWinRM

 A Python library that can be used to automate Windows Remote Management tasks, which can be used to automate Windows server configuration and management.

## Pulumi

Pulumi is a cloud infrastructure as code (CIaC) tool that lets you define and manage cloud resources in a variety of programming languages, including Python.

You can use Pulumi to write Python code to describe your infrastructure as code, and then use the Pulumi CLI to deploy and manage it. Here is an example:

``` python
import pulumi
from pulumi_aws import ec2

# Define an EC2 instance
server = ec2.Instance('server',
    instance_type='t2.micro',
    ami='ami-0c55b159cbfafe1',
    tags={
        'Name': 'cloud-server',
    },
)

# Export the server's IP address
pulumi.export('ip_address', server.public_ip)
```

In this example, we're using the Pulumi Python SDK to define an EC2 instance on AWS. We specify the instance type, the AMI ID, and some tags for the instance, and then export the instance's public IP address. ou can then use the Pulumi CLI to deploy this infrastructure, which will create the EC2 instance on AWS. You can also use the Pulumi CLI to manage your infrastructure over time, making changes and updates as needed.

## Resources:

- [Learn more about Fabric](https://docs.fabfile.org/en/stable/index.html)
- [PyWinRM](https://github.com/diyan/pywinrm)
- [Pulumi - IaC Tool](https://www.pulumi.com/docs/reference/pkg/python/pulumi/)

See you tomorrow in [Day 48](day48.md).
