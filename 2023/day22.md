# Continuous Image Repository Scan - Container Registries

Yesterday we learned how to integrate container image vulnerability scanning into our CI/CD pipelines.

Today, we are going to take a look at how to enforce that our images are scanned on another level - the container registry.

There are container registries that will automatically scan your container images once you push them.
This ensures that we will have visibility into the number of vulnerabilities for every container image produced by our team.

Let's take a look at few different registries that provide this capability and how we can use it.

## Docker Hub

[Docker Hub](https://hub.docker.com/) is the first container registry.
It was build by the team that created Docker and is still very popular today.

Docker Hub has automatic vulnerability scanner, powered by [Snyk](https://snyk.io/).

This means that, if enabled, when you push an image to Docker Hub it will be automatically scanned and the results with be visible to you in the UI.

You can learn more about how to enable and use this feature from the Docker Hub [docs](https://docs.docker.com/docker-hub/vulnerability-scanning/).

**NOTE:** This feature is not free.
In order to use it you need to have a subscription.

## Harbor

[Harbor](https://goharbor.io/) is an open-source container registry.
Originally developed in VMware, it is now part of the CNCF.

It supports image scanning via [Trivy](https://github.com/aquasecurity/trivy) and/or [Clair](https://github.com/quay/clair).

This is configured during installation.
(Even if you don't enable image scanning during installation, it can always be configured afterwards).
For more info, check out the [docs](https://goharbor.io/docs/2.0.0/administration/vulnerability-scanning/).

## AWS ECR

[AWS ECR](https://aws.amazon.com/ecr/) also supports [image scanning via Clair](https://docs.aws.amazon.com/AmazonECR/latest/userguide/image-scanning-basic.html).

## Azure Container Registry

[Azure Container Registry](https://azure.microsoft.com/en-us/products/container-registry) support [image scanning via Qualys](https://azure.microsoft.com/en-us/updates/vulnerability-scanning-for-images-in-azure-container-registry-is-now-generally-available/).

## GCP

[GCP Container Registry](https://cloud.google.com/container-registry) also support [automatic image scanning](https://cloud.google.com/container-analysis/docs/automated-scanning-howto).

## Policy Enforcement

Just scanning the images and having the results visible in your registry is nice thing to have,
but it would be even better if we have a way to enforce some standards for these images.

In [Day 14](day14.md) we saw how to make `grype` fail a scan if an image has vulnerabilities above a certain severity.

Something like this can also be enforced on the container registry level.

For example, [Harbor](https://goharbor.io/) has the **Prevent vulnerable images from running** option, which when enable does not allow you to pull an image that has vulnerabilities above a certain severity.
If you cannot pull the image, you cannot run it, so this is a good rule to have if you don't want to be running vulnerable images.
Of course, a rule like that can effectively prevent you from deploying something to your environment, so you need to use it carefully.

More about this option and how to enable it in Harbor you can read [here](https://goharbor.io/docs/2.3.0/working-with-projects/project-configuration/).

For more granular control and for unblocking deployments you can configure a [per-project CVE allowlist](https://goharbor.io/docs/2.3.0/working-with-projects/project-configuration/configure-project-allowlist/).
This will allow certain images to run even though they have vulnerabilities.
However, these vulnerabilities would be manually curated and allow-listed by the repo admin.

## Summary

Scanning your container images and having visibility into the number of vulnerabilities inside them is critical for a secure SDLC.

One place to do that is you CI pipeline (as seen in [Day 21](day21.md)).

Another place is your container registry (as seen today).

Both are good options, both have their pros and cons.
It is up to the DevSecOps architect to decide which approach works better for them and their thread model.
See you on [Day 23](day23.md).
