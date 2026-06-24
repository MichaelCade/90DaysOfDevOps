# Day 90 - Fighting fire with fire: Why we cannot always prevent technical issues with more tech
[![Watch the video](thumbnails/day90.png)](https://www.youtube.com/watch?v=PJsBQGKkn60)

 To summarize, the goal is to create effective documentation that empowers your team to respond effectively to various situations. You've identified two main types of documentation: developer documentation (internal team, focusing on tool usage and source code) and user documentation (external users, focusing on product usage).

In Car Engineering (CAR E), you want to document known knowns (things understood and aware of), known unknowns (issues understood but not yet fully comprehended), known unknowns (understood but not yet aware), and unknown unknowns (unknown and uncomprehended issues). To understand the current state of your infrastructure, map services dependencies (internal and external) and plan experiments with defined goals, components, expected results, and factors affecting your hypothesis.

Projects such as CNCF's Mesh and Litmos CS can help automate experiments and post-mortem reviews, which are crucial for learning from past incidents and improving future responses. Postmortems also serve as an opportunity to enhance documentation on incident resolution.

The key takeaway is to document your work consistently, whether you're working with your own infrastructure or that of an organization. Sharing your notes publicly can be beneficial to others who may join your team in the future. Remember, there's no right or wrong when it comes to writing—it's better to write things down incorrectly than for someone to try a command based on your past notes that no longer works.

Thank you for joining this presentation. For more content from me, visit an.com. Shout out to Michael, the organizer of 90 Days of DevOps, for having me here. I hope this was useful for you, and I look forward to seeing you on my YouTube channel and potentially at a conference in person. Have an amazing day!

**Whom are you writing for?**
The audience for documentation can be developers (e.g., within a team or open-source project) or users (e.g., end-users of a service). When writing for developers, assume an existing knowledge level, while for users, provide more detailed explanations. Consider the stages when the product is supposed to be used (e.g., installation, upgrade, or use new features).

**What are the goals?**
Documentation aims to reach specific goals, such as providing setup and configuration guides or describing implementation scenarios. Tutorials typically have a narrower scope, serving a specific use case.

**Technical Solutions:**
The presentation mentions KGBT (Knowledge Graph-Based Tool) for scanning Kubernetes clusters and triaging issues. This tool helps enrich documentation with AI-powered insights. Other technical solutions include documentation frameworks like Diet Taco's framework, which models documentation after KGBT and Canonical's model.

**C Engineering Experiments:**
In C engineering, there are known-knowns (things we're aware of and understand), known-unknowns (aware but don't understand), unknown-knowns (understand but not aware), and unknown-unknowns (neither aware nor understand). To gain an understanding of the current state of infrastructure, map services, dependencies, both internal and external.

**Post-Mortem Reviews:**
After running experiments, conduct post-mortem reviews to understand which components have been tested and how to solve incidents in the future. This allows for sharing knowledge with others and enhancing documentation on resolving incidents.

**Main Tip:**
The presenter's main tip is to start writing down what you're doing, whether working with your own home cluster or organization's infrastructure. Share notes publicly to help others and gain value from documenting experiences.
