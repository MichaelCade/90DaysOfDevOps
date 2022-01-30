## Microsoft Azure Security Models

Following on from the Microsoft Azure Overview, we are going to start with Azure Security and see where this can help in our day today. For the most part, I have found the built-in roles have been sufficient but knowing that we can create and work with many different areas of authentication and configurations. I have found Microsoft Azure to be quite advanced with its Active Directory background compared to other public clouds. 

## Microsoft Azure Security Models

This is one area that Microsoft Azure seemingly works differently from other public cloud providers, in Azure there is ALWAYS Azure AD. 

### Directory Services 

- Azure Activ Directory hosts the security principles used by Microsoft Azure and other Microsoft cloud services. 
- Authentication is accomplished through protocols such as SAML, WS-Federation, OpenID Connect and OAuth2. 
- Queries are accomplished through REST API called Microsoft Graph API. 
- Tenants have a tenant.onmicrosoft.com default name but can also have custom domain names. 
- Subscriptions are associated with an Azure Active Directory tenant. 

If we think about AWS to compare the equivalent offering would be AWS IAM (Identity & Access Management) Although still very different 

Azure AD Connect provides the ability to replicate accounts from AD to Azure AD. This can also include groups and sometimes objects. This can be granular and filtered. Supports multiple forests and domains. 

It is possible to create cloud accounts in Microsoft Azure Active Directory (AD) but most organisations already have accounted for their users in their own Active Directory being on-premises. 

Azure AD Connect also allows you to not only see Windows AD servers but also other Azure AD, Google and others. This also provides the ability to collaborate with external people and organisations this is called Azure B2B. 

Authentication options between Active Directory Domain Services and Microsoft Azure Active Directory are possible with both identity sync with a password hash.

![](Images/Day30_Cloud1.png)

The passing of the password hash is optional, if this is not used then pass-through authentication is required. 

There is a video linked below that goes into detail about Passthrough authentication. 

[User sign-in with Azure Active Directory Pass-through Authentication](https://docs.microsoft.com/en-us/azure/active-directory/hybrid/how-to-connect-pta)

![](Images/Day30_Cloud2.png)

### Federation 

It's fair to say that if you are using Microsoft 365, Microsoft Dynamics and on-premises Active Directory it is quite easy to understand and integrate into Azure AD for federation. However, you might be using other services outside of the Microsoft ecosystem. 

Azure AD can act as a federation broker to these other Non-Microsoft apps and other directory services. 

This will be seen in the Azure Portal as Enterprise Applications of which there are a large number of options. 

![](Images/Day30_Cloud3.png)

If you scroll down on the enterprise application page you are going to see a long list of featured applications. 

![](Images/Day30_Cloud4.png)

This option also allows for "bring your own" integration, an application you are developing or a non-gallery application. 

I have not looked into this before but I can see that this is quite the feature set when compared to the other cloud providers and capabilities. 

### Role-Based Access Control 

We have already covered on [Day 29](day29.md) the scopes we are going to cover here, we can set our role-based access control according to one of these areas. 

- Subscriptions
- Management Group
- Resource Group 
- Resources 

Roles can be split into three, there are many built-in roles in Microsoft Azure. Those three are: 

- Owner 
- Contributor 
- Reader 

Owner and Contributor are very similar in their boundaries of scope however the owner can change permissions. 

Other roles are specific to certain types of Azure Resources as well as custom roles. 

We should focus on assigning permissions to groups vs users. 

Permissions are inherited. 

If we go back and look at the "90DaysOfDevOps" Resource group we created and check the Access Control (IAM) within you can see we have a list of contributors and a customer User Access Administrator, and we do have a list of owners (But I cannot show this)

![](Images/Day30_Cloud5.png)

We can also check the roles we have assigned here if they are BuiltInRoles and which category they fall under. 

![](Images/Day30_Cloud6.png)

We can also use the check access tab if we want to check an account against this resource group and make sure that the account we wish to have that access to has the correct permissions or maybe we want to check if a user has too much access. 

![](Images/Day30_Cloud7.png)

### Azure Security Centre 

- Microsoft Azure Security Centre provides insight into the security of the entire Azure environment. 

- A single dashboard for the overall security health of all Azure resources. 

- Free and Standard SKU

I have switched to another subscription to view the Azure Security centre and you can see here based on very few resources that I have some recommendations in one place. 

![](Images/Day30_Cloud8.png)

## Resources 

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

See you on [Day 31](day31.md)