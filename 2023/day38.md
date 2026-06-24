# Increase the Security Posture of Your Organization with Dynamic Credentials 

As we talked about yesterday, Vault is commonly used as a platform to consolidate your static, long-lived credentials. However, you're still stuck with the management nightmare of rotating those credentials based on your organization's security policies. This credential rotation "tradition" is a manual, laborious task and is susceptible to errors resulting in application downtime. But…

What if you could get rid of static credentials altogether? If your app only needs database access once a week to run reports, why do you give it a credential that is valid 24/7/365? Wouldn't it be great to have your applications generate dynamic, self-destructing credentials on-demand for the systems needed to function? 

Well, it's all possible using many of the secrets engines available in Vault. 

## Intro to Dynamic Secrets

Dynamic secrets are credentials that are generated on demand on behalf of the requesting Vault client. Rather than simply reading a static credential stored in the KV store, Vault can make an API call to the requested platform, create a credential, and pass the credential back to the user. In the process, Vault attaches a lease (TTL) to the credential, which indicates how long the credential is valid. The Vault client can then use the credential to communicate directly with the platform for its intended function. Once the credential's lease expires, Vault calls back to the platform and deletes the credential, making it invalid.

## Benefits of Using Dynamic Secrets

There are so many benefits of migrating to dynamic credentials. The obvious benefit is not having long-lived credentials that are manually rotated. Because these long-lived credentials are often shared between teams and application stacks, they are more susceptible to misuse or abuse. However, when you migrate to dynamic creds, each application instance can retrieve a unique credential to access the required platform when using dynamic credentials. And if that application instance is terminated (think containerization or auto-scaling), the credential will be invalidated by Vault and not impact other instances or applications in your environment. 

Dynamic secrets also eliminate the manual process of rotating credentials. Rather than rotating the credentials once a year, these highly privileged creds are now rotated once a month, once a week, or once every hour. For example, consider how you might use Terraform to deploy resources on your public or private cloud. You likely create a credential on the target platform and use it repeatedly via environment variables or sensitive workspace variables in TFC/TFE. But why are you giving Terraform highly privileged credentials that are valid 24/7/365 when your Terraform runs only take 15 minutes each day? Switch to dynamic credentials. Using a combination of Vault secrets engine and the Vault provider for Terraform, you can have Terraform generate a dynamic credential for the platform it needs to access for resource deployment or management.

## Configure a Dynamic Secrets Engine

With that long-winded introduction out of the way, let's talk about how we can accomplish this goal of using dynamic credentials. HashiCorp Vault offers a plethora of secrets engines that generate dynamic credentials or data, such as:
•	AWS
•	Azure
•	GCP
•	Active Directory
•	AliCloud
•	Consul
•	Databases (supports about 15 different database platforms)
•	PKI certificates
•	Kubernetes
•	RabbitMQ
•	SSH
•	Terraform Cloud

I think you'd agree that's a lot of options for a single platform.  The cool thing is that you can use as many of these secrets engines as you want, and even enable many of the same type of secrets engine. Let's take a look at how we would enable one of these, using AWS as the example.

Each secrets engine must be enabled on a path, and all interactions with the secrets engine is then done using the path. Because of this, each secrets engine must be enabled on a unique path so Vault knows where to route the request. To enable a secrets engine, use the following structure for the command `vault secrets enable -path=<name> <secrets_engine_type>`. Enabling the AWS secrets engine on the path of `cloud` would look like this:

`vault secrets engine -path=cloud aws`

> Note that if you do not provide the path flag, the secrets engine will be enabled on the default path, which is the same name of the secrets engine type. For example, the AWS secrets engine would be enabled at aws/.

Ok, so we've got a secrets engine enabled, let's start configuring it. Configuration of a secrets engine generally requires a few things. The first is a way for Vault to communicate with the platform (in this case, AWS). Just like any other AWS client, Vault needs credentials to authenticate to the target platform to perform actions. In the case of AWS, you can provide Vault with an Access Key and Secret Key or use an IAM role if Vault was deployed on AWS.

In addition to authentication, Vault needs the proper authorization to perform actions on the platform. For AWS, that equates to having an IAM policy attached to the credentials you will give to Vault. Depending on the type of credential you want to generate (Vault supports 3 types for AWS), the policy should allow permission to create/manage/delete a user, generate keys, manage the keys, etc. These policies are provided by HashiCorp in their documentation. I also have some pre-defined policies for Vault [on my GitHub](https://github.com/btkrausen/hashicorp).

To provide credentials to Vault to access your AWS account, you can use the following command:

```
vault write aws/config/root \
    access_key=AKIAIOSFODNN7EXAMPLE \
   secret_key=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY \
    region=us-east-1
```

> Although you can add your credentials using this configuration, you can also provide them using environment variables as well.

Ok, now Vault has access to our AWS account, it doesn't yet know what credentials to create and the level of access those credentials should have in your AWS account. Each unique requirement for credentials will be defined in a role (that's a Vault role, not to be confused with an AWS role). Roles are used to map a user to a set of permissions on the platform. For example, you can create a role for a developer that provides read-only access to the AWS account, or you can create a role for Terraform that provides a broader set of privileges to create, manage, and delete resources. Roles can be used by many applications, but keep in mind that a different role will be required for each UNIQUE set of permissions you want to give the Vault client.

Let's create our first role. This role will be for our developer who needs read-only access to an AWS so they can obtain logs for troubleshooting an application. The role name will be `developer`.

```
vault write aws/roles/developer \
    policy_arns=arn:aws:iam::aws:policy/AmazonEC2ReadOnlyAccess \
   credential_type=iam_user
```

> In this example, I used the IAM_USER credential type but for production environments, I'd recommend using ASSUME_ROLE so you can have a

When a developer requests credentials, they will be tied to the AWS-managed policy named ` arn:aws:iam::aws:policy/AmazonEC2ReadOnlyAccess`. You can use AWS-managed policies, Customer-Managed policies, or simply provide the policy within the Vault role configuration and Vault will create the inline policy when the user is created.

## Let's Get Credentials

Ok, we have the secrets engine mounted (enabled!), the secrets engine configured to access our account, and now we have a role. The next logical step is to test it and obtain our first set of dynamic credentials. Since the role we created is named `developer`, that will come into play for the path needed to get credentials. To get creds, use the following command:

`vault read aws/creds/developer`

If successful, Vault should have returned a set of credentials that the client can use to interact directly with AWS. Woohoo! 

Keep in mind that when the lease (TTL) expires, Vault will go back and delete the credentials on AWS, permanently invaliding them.

## Day 2 Ops 

The last thing I wanted to touch on was managing the credentials used by Vault to access the platform. The whole purpose of a dynamic secrets engine is to eliminate your static, long-lived credentials. However, the first step here was to provide Vault…..static, long-lived credentials to access the platform? Hmm…sounds somewhat counterproductive, right? What if THAT credential gets compromised or shared?

Fear not, as Vault provides a simple endpoint for secrets engines which allows Vault to quickly rotate that 'root' credential. You can hit this endpoint as often as needed to rotate the credential Vault uses to access the platform. In this example, you can use the following command to rotate the AWS credential we provided above.

`vault write -f aws/config/rotate-root`

Once you run this command, only AWS and Vault know the credentials. No human user has the root credential, and even a Vault administrator can read back the full access key and secret key. This operation can be run as often as needed to meet your internal security policies for credential rotation.

See you on [Day 39](day39.md)

