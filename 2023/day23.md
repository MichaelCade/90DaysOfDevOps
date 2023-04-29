# Artifacts Scan

In the previous two days we learned why and how to scan container images.

However, usually our infrastructure consists of more than just container images.
Yes, our services will run as containers, but around them we can also have other artifacts like:

- Kubernetes manifests
- Helm templates
- Terraform code

For maximum security, you would be scanning all the artifacts that you use for your environment, not only your container images.

The reason for that is that even if you have the most secure Docker images with no CVEs,
but run then on an insecure infrastructure with bad Kubernetes configuration,
then your environment will not be secure.

**Each system is as secure as its weakest link.**

Today we are going to take a look at different tools for scanning artifacts different than container images.

## Kubernetes manifests

Scanning Kubernetes manifests can expose misconfigurations and security bad practices like:

- running containers as root
- running containers with no resource limits
- giving too much and too powerful capabilities to the containers
- hardcoding secrets in the templates, etc.

All of these are part of the security posture of our Kubernetes workloads, and having a bad posture in security is just as bad as having a bad posture in real-life.

One popular open-source tool for scanning Kubernetes manifests is [KubeSec](https://kubesec.io/).

It outputs a list of misconfiguration.

For example, this Kubernetes manifest taken from their docs has a lot of misconfigurations like missing memory limits, running as root, etc.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: kubesec-demo
spec:
  containers:
    - name: kubesec-demo
      image: gcr.io/google-samples/node-hello:1.0
      securityContext:
        runAsNonRoot: false
```

Let's scan it and look at the results.

```shell
$ kubesec scan kubesec-test.yaml
[
  {
    "object": "Pod/kubesec-demo.default",
    "valid": true,
    "message": "Passed with a score of 0 points",
    "score": 0,
    "scoring": {
      "advise": [
        {
          "selector": ".metadata .annotations .\"container.seccomp.security.alpha.kubernetes.io/pod\"",
          "reason": "Seccomp profiles set minimum privilege and secure against unknown threats"
        },
        {
          "selector": ".spec .serviceAccountName",
          "reason": "Service accounts restrict Kubernetes API access and should be configured with least privilege"
        },
        {
          "selector": "containers[] .securityContext .runAsNonRoot == true",
          "reason": "Force the running image to run as a non-root user to ensure least privilege"
        },
        {
          "selector": ".metadata .annotations .\"container.apparmor.security.beta.kubernetes.io/nginx\"",
          "reason": "Well defined AppArmor policies may provide greater protection from unknown threats. WARNING: NOT PRODUCTION READY"
        },
        {
          "selector": "containers[] .resources .requests .memory",
          "reason": "Enforcing memory requests aids a fair balancing of resources across the cluster"
        },
        {
          "selector": "containers[] .securityContext .runAsUser -gt 10000",
          "reason": "Run as a high-UID user to avoid conflicts with the host's user table"
        },
        {
          "selector": "containers[] .resources .limits .cpu",
          "reason": "Enforcing CPU limits prevents DOS via resource exhaustion"
        },
        {
          "selector": "containers[] .resources .requests .cpu",
          "reason": "Enforcing CPU requests aids a fair balancing of resources across the cluster"
        },
        {
          "selector": "containers[] .securityContext .readOnlyRootFilesystem == true",
          "reason": "An immutable root filesystem can prevent malicious binaries being added to PATH and increase attack cost"
        },
        {
          "selector": "containers[] .securityContext .capabilities .drop",
          "reason": "Reducing kernel capabilities available to a container limits its attack surface"
        },
        {
          "selector": "containers[] .resources .limits .memory",
          "reason": "Enforcing memory limits prevents DOS via resource exhaustion"
        },
        {
          "selector": "containers[] .securityContext .capabilities .drop | index(\"ALL\")",
          "reason": "Drop all capabilities and add only those required to reduce syscall attack surface"
        }
      ]
    }
  }
]
```

As we see it produced 12 warnings about thing in this manifests we would want to change.
Each warning has an explanation telling us WHY we need to fix it.

### Others

Other such tools include [kube-bench](https://github.com/aquasecurity/kube-bench), [kubeaudit](https://github.com/Shopify/kubeaudit) and [kube-score](https://github.com/zegl/kube-score).

They work in the same or similar manner.
You give them a resource to analyze and they output a list of things to fix.

They can be used in a CI setup.
Some of them can also be used as [Kubernetes validating webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/), and can block resources from being created if they violate a policy.

## Helm templates

[Helm](https://helm.sh/) templates are basically templated Kubernetes resources that can be reused and configured with different values.

There are some tools like [Snyk](https://docs.snyk.io/products/snyk-infrastructure-as-code/scan-kubernetes-configuration-files/scan-and-fix-security-issues-in-helm-charts) that have *some* support for scanning Helm templates for misconfigurations the same way we are scanning Kubernetes resources.

However, the best way to approach this problem is to just scan the final templated version of your Helm charts.
E.g. use the `helm template` to substitute the templated values with actual ones and just scan that via the tools provided above.

## Terraform

The most popular tool for scanning misconfigurations in Terraform code is [tfsec](https://github.com/aquasecurity/tfsec).

It uses static analysis to spot potential issues in your code.

It support multiple cloud providers and points out issues specific to the one you are using.

For example, it has checks for [using the default VPC in AWS](https://aquasecurity.github.io/tfsec/v1.28.1/checks/aws/ec2/no-default-vpc/),
[hardcoding secrets in the EC2 user data](https://aquasecurity.github.io/tfsec/v1.28.1/checks/aws/ec2/no-secrets-in-launch-template-user-data/),
or [allowing public access to your ECR container images](https://aquasecurity.github.io/tfsec/v1.28.1/checks/aws/ecr/no-public-access/).

It allow you to enable/disable checks and to ignore warnings via inline comments.

It also allows you to define your own policies via [Rego](https://www.openpolicyagent.org/docs/latest/policy-language/).

## Summary

A Secure SDLC would include scanning of all artifacts that end up in our production environment, not just the container images.

Today we learned how to scan non-container artifacts like Kubernetes manifests, Helm charts and Terraform code.
The tools we looked at are free and open-source and can be integrated into any workflow or CI pipeline.
See you on [Day 24](day24.md).
