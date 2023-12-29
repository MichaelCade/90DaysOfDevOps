# Reuse, Don't Repeat - Creating an Infrastructure as Code Module Library 

When you've been working with Infrastructure as Code for a while, you'll inevitably start thinking about how you can make your code more reusable. This is especially true if you're working within a team and want to share resources and knowledge. This is where modularisation comes into play - taking the logic you've created and turning it into a reusable module that can be used by anyone who wants to create that kind of resource.

## Why Create Reusable Infrastructure as Code?

There will usually come a point when you are writing some IaC when you realise you've written it before; you might even decide to copy and paste some code between your previous project and your current one. At this point, you think "there has to be a better way", and usually there is, creating modules.

Each IaC language calls them something different, but we will call them modules for these articles. This is taking some logic to deploy some infrastructure, encapsulating it, adding some parameters and outputs and a way to distribute it to the people who want to consume it so that it can be reused over and over. However, before you run off and start creating modules, you need to consider another point. Does it add value?

You can create any module you want, but there is no point if it doesn't add any value. There's no point creating a module in Bicep that wraps the virtual machine resource and doesn't do anything else; you may as well use the virtual machine resource on its own. Any module you create should add some value on top of the already existing resource. In my experience, the value will usually come from one or more of these areas.

### Deploying Multiple Resources

If you have a solution containing multiple different resources that are regularly deployed together, this is an ideal candidate for creating a module. Maybe you have an application you deploy that uses a web app, SQL database, storage account and application insights workspace. You can create a module that provides a simple interface for the end-user deploying this application, who doesn't need any knowledge of underlying components; they need to work with the module's interface. Because your abstracting your logic into this module, it also frees you up to add or change the underlying resources in the module if you need to.

### Simplifying Complex Resources

Some resources are complex and have lots of properties and configuration options. This is great for flexibility, but maybe you always deploy the resource in a certain way, and there are some properties that will always be set the same way. You can use a module to abstract away from the complex interface and present a simpler one to the end-user. Inside your module, you set the properties that don't need to be changed by the end-user. An excellent example of this is a virtual machine. Maybe you always have accelerated networking turned on, you always use a particular operating system version, you don't need to offer any configuration on the network interface, and you always install the BitLocker extensions. You could create a VM module that does all of this and only provides the user choices they might want to change.

### Enforcing Policy and Standards

Most companies will have standards on how cloud infrastructure needs to be deployed and usually involve resources being configured in a certain way. This is an excellent use for modularisation by building these standards as part of the module, so the end-user doesn't even have to think about it. The user knows that they will get their resources deployed in a compliant way if they use the module. For example, a common standard is that a network security group needs to be deployed with specific rules applied, such as allowing traffic from the corporate network, allowing access to monitoring tools, etc. We can create an NSG module that the user can deploy that has these rules built-in and already configured. The user could then layer on top any custom rules they want.

### Providing Utility or Extension

It might be that the benefit of your module is not in the resource that it deploys but that it does some beyond just creating a resource. Maybe your module deploys a web application, but it also runs some additional code to deploy your website into the web app. Another example of a utility module is what I call a "configuration module", a module that doesn't create any resources but contains configuration settings that might be consumed by other modules or by the end-user. We'll look at this approach in more detail in a future article.

## Module Design

Once you've decided what module you want to create and confirmed that you are adding value, it's time to design the module. There are a few tips I would suggest looking at when you plan your module.

- Do one thing well - it can be tempting to build large complex modules that do lot's of things, but most of the time, it's better to create a module that does one specific thing and make them as self-contained and straightforward as possible. This way, users can pick and choose the modules they need and even chain them together.
- Design upfront - before you start coding your module, you should have a good idea of what you want your module to do and what its inputs and outputs look like. This will likely evolve as you create it, but at least you will have a starting point.
- Provide options, but not too many - you want your module to be parameterised so that users can use it in different scenarios and customise it appropriately. Still, you want to make sure you keep your interface as straightforward as possible and only offer parameters that make sense. If 99% of the time the user is going to pick the same value, then maybe don't offer that parameter.
- Version your modules - You will inevitably release new versions of your modules over time, so make sure that you version them appropriately. Users need to know when a module has been updated and to be able to choose when they update, especially if there are breaking changes
- Make them easy to obtain - it's no good having an extensive library of modules if no one can get hold of them. Each IaC language has a different way of distributing modules, and we'll look at that in more detail in future articles. Make sure you use that functionality to provide an easy way for users to install and update your modules
- Keep them updated - it's no good creating a module and then never updating it. The cloud resources your module creates will be updated; your module needs to as well. Make sure you commit to keeping them updated; otherwise, no one is going to use them.
- Create good documentation - for someone to use your module, they need documentation. Make sure it explains things well, is easy to get at and is kept updated.

## Useful Links

- [Demo Source Code](https://github.com/sam-cogan/events/tree/main/90DaysofDevOps24/Examples)
- [Creating a Configuration Module for your Infrastructure as Code](https://samcogan.com/creating-a-configuration-module-for-your-infrastructure-as-code/)