# Day 47 - Automation with Python

Using Python for infrastructure management involves automating the management of IT infrastructure, such as servers, databases, and networking equipment. This can include tasks like provisioning, configuration, and orchestration.
Python is a popular language for infrastructure management, and there are several tools and libraries available to help with this. Some popular tools for infrastructure management that use Python include:

- Fabric
- PyWinRM
- Pulumi

## Fabric

Fabric is a Python library that can be used for streamlining SSH commands and remote execution of scripts, which can be used to automate server configuration and deployment
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
