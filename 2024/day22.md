# Day 22 - Test in Production with Kubernetes and Telepresence
[![Watch the video](thumbnails/day22.png)](https://www.youtube.com/watch?v=-et6kHmK5MQ)

 To summarize, Telepresence is an open-source tool that allows developers to test their code changes in a Kubernetes environment without committing, building Docker images, or deploying. It works by redirecting incoming requests from a service in a remote Kubernetes cluster to the local machine where you're testing. This is achieved through global interception mode (for all requests) and personal interception mode (for specific request headers).

To set it up:
1. Configure your local setup.
2. Install Telepresence on your Kubernetes cluster.
3. Test the whole thing.

Details can be found in this blog post: arab.medium.com/telepresence-kubernetes-540f95a67c74

Telepresence makes the feedback loop shorter for testing on Kubernetes, especially with microservices where it's difficult to run everything locally due to dependencies. With Telepresence, you can mark just one service and run it on your local machine for easier testing and debugging.

**Summary:**
The speaker shares their experience with using a staging environment to test code before deploying it to production. They mention how they missed a column in their code, which broke the staging environment, but was caught before reaching production. The speaker introduces Telepresence, an open-source tool that allows developers to automatically deploy and test their code on a local machine, without committing changes or running CI/CD pipelines.

**Key Points:**

1. Importance of having a staging environment for testing code.
2. How missing a column in the code can break the staging environment.
3. Introduction to Telepresence as a solution to improve the development process.
4. Benefits of using Telepresence, including:
	* Shorter feedback loop
	* Ability to test and debug services locally
	* Open-source and community-driven

**Purpose:**
The speaker aims to share their experience with using a staging environment and introducing Telepresence as a tool to improve the development process. The purpose is to educate developers about the importance of testing code before deploying it to production and provide a solution to make this process more efficient and effective.

