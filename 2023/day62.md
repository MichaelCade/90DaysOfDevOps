# Compliance and Vulnerability Scanning provided by Red Hat OpenShift Operators

Many components make up the security posture of a well-designed and architected platform. For a Red Hat OpenShift platform, this will include implementing the AAA model we covered on [Day 61](/2023/day61.md), container security, certificates to validate access to the platform and between the workloads that run and communicate on and with the platform, data encryption, vulnerability scanning, compliance auditing and remediation, as just a shortlisted example.

In this post, I'm going to focus on just two areas of security inside of Red Hat OpenShift; Compliance and Vulnerability Scanning.

# Red Hat OpenShift Compliance Operator

## Compliance Overview

In the context of Kubernetes, compliance refers to the adherence of OpenShift deployments to various industry standards, regulations, and best practices. Compliance is essential for organizations operating in regulated industries like healthcare, finance, or government sectors, as well as for those who prioritize security and privacy.

OpenShift, as a Kubernetes-based container orchestration platform, provides several features and tools to help organizations achieve and maintain compliance. These features include:

- Security and access control: OpenShift provides robust security features like Role-Based Access Control (RBAC), network policies, and Security Context Constraints (SCCs) to manage access to cluster resources and ensure the secure deployment and operation of applications.

- Auditing and monitoring: OpenShift's built-in auditing and monitoring capabilities make it easy to track user activities, resource usage, and system events. This information is crucial for detecting and responding to security incidents, meeting regulatory requirements, and troubleshooting issues.

- Image and container security: OpenShift's integrated container registry, image signing, and image scanning features help ensure the integrity and security of container images. Additionally, OpenShift enforces security best practices through resource constraints and pod security policies.

- Encrypted communication: OpenShift supports TLS/SSL for secure communication between cluster components, as well as between the cluster and external clients. This helps protect sensitive data in transit.

- Compliance Operator: Red Hat OpenShift provides the Compliance Operator, with an OpenShift-native tool that helps organizations evaluate and enforce compliance policies. The Compliance Operator leverages the OpenSCAP framework and can be configured to meet specific regulatory requirements or security standards, such as NIST, PCI-DSS, HIPAA, or GDPR.

Achieving and maintaining compliance in OpenShift involves configuring the platform according to industry standards, regulations, and best practices, continuously monitoring and auditing the environment, and promptly addressing any identified issues. By utilising the capabilities and resources offered by OpenShift, organisations have the opportunity to establish secure and regulation-compliant ecosystems for their applications and workloads, ensuring optimal performance and adherence to industry standards.

In this post, we'll be focusing on the Compliance Operator, which empowers the platform administrators by allowing them to define the desired compliance state for their cluster. It offers a comprehensive overview of discrepancies between the current and target states, as well as actionable insights to address these gaps effectively.

The Compliance Operator provides the ability to describe the required compliance state of a cluster and report overviews of gaps and ways to remediate them. The Compliance Operator assesses compliance of both the Kubernetes API resources of OpenShift Container Platform, as well as the nodes running the cluster. The Compliance Operator uses OpenSCAP, a NIST-certified tool, to scan and enforce security policies provided by the content. You can view details of the out-of-the-box compliance profiles that are [provided here](https://docs.openshift.com/container-platform/4.12/security/compliance_operator/compliance-operator-supported-profiles.html).

## Installing the Compliance Operator

1. We need to create a specially managed OpenShift namespace. Save the below content as ```openshift-compliance-namespace.yaml``` and apply using the ```oc apply -f openshift-compliance-namespace.yaml``` command.

````yaml
apiVersion: v1
kind: Namespace
metadata:
  labels:
    openshift.io/cluster-monitoring: "true"
    pod-security.kubernetes.io/enforce: privileged 
  name: openshift-compliance
````

2. Now we need to define an ```OperatorGroup``` for the Compliance operator. Save the below content as ```openshift-compliance-operator-group.yaml``` and apply using the ```oc apply -f openshift-compliance-operator-group.yaml``` command.

````yaml
apiVersion: operators.coreos.com/v1
kind: OperatorGroup
metadata:
  name: compliance-operator
  namespace: openshift-compliance
spec:
  targetNamespaces:
  - openshift-compliance
````

3. Define and apply a subscription to the Operator. Save the below content as ```openshift-compliance-operator-subscription.yaml``` and apply using the ```oc apply -f openshift-compliance-operator-subscription.yaml``` command.

````yaml
apiVersion: operators.coreos.com/v1alpha1
kind: Subscription
metadata:
  name: compliance-operator-sub
  namespace: openshift-compliance
spec:
  channel: "release-0.1"
  installPlanApproval: Automatic
  name: compliance-operator
  source: redhat-operators
  sourceNamespace: openshift-marketplace
````

4. You can verify the Operator installation with the below command;

````sh
$ oc get csv -n openshift-compliance
$ oc get deploy -n openshift-compliance
````


# Red Hat Quay Container Security Operator

## Vulnerability Scanning Overview

Vulnerability scanning in Red Hat OpenShift refers to the process of inspecting container images for known security issues, such as outdated software packages, misconfigurations, or exposed sensitive information. The goal of vulnerability scanning is to identify and remediate potential security risks in container images before they are deployed to the OpenShift cluster, thus enhancing the overall security posture of the platform.

OpenShift provides several features and tools to facilitate vulnerability scanning:

- Integrated container registry: OpenShift includes a built-in container registry to store and manage container images. The integrated registry allows for a more streamlined and secure process when scanning images for vulnerabilities, as it eliminates the need to rely on external registries.

- ImageStreams: An ImageStream in OpenShift is an abstraction that represents a series of related container images, typically different versions of the same application. ImageStreams simplify the process of tracking and deploying container images, making it easier to apply vulnerability scanning and remediation across multiple versions of an application.

- Image signing and trust: OpenShift supports container image signing and the enforcement of signature-based trust policies. This feature ensures that only trusted and verified images can be deployed to the cluster, helping to prevent the deployment of images with known vulnerabilities.

- Third-party integrations: OpenShift can be easily integrated with external vulnerability scanning tools and platforms, such as Aqua Security, Sysdig, or Twistlock. These tools can be configured to automatically scan container images stored in the OpenShift registry and provide detailed reports on identified vulnerabilities and suggested remediation steps.

- OpenShift Operators: OpenShift supports the use of Operators, which are automated software extensions that manage applications and their components. Operators can be used to deploy and manage vulnerability scanning tools within the OpenShift cluster, ensuring a consistent and automated scanning process. Red Hat provides the ```Red Hat Quay Container Security Operator```, however you can also implement third party scanners such as [Trivy](https://github.com/aquasecurity/trivy) from [Aqua Security](https://aquasec.com/).

By leveraging these features and tools, Red Hat OpenShift enables organizations to perform comprehensive vulnerability scanning on container images, reducing the risk of security breaches and enhancing the overall security of the platform.
# Summary

Whilst for this 2023 edition focusing on DevSecOps, we could have purely spent time focusing on Security and Compliance for Red Hat OpenShift in-depth, I wanted to start at a higher level, understanding why you would choose an enterprise Kubernetes offering, and what features will enhance your cloud-native platform. Hopefully, this has given you a solid understanding of this offering, as well as being able to understand the basics of how to run and operate it. Another area we only touched upon briefly was application deployment, instead focusing on the security posture of deploying workloads, rather than the methods of building and running the applications themselves. This topic of application build, deployment and management requires a whole section on its own.

I urge you to spend time reading through the official documentation for Red Hat OpenShift, it's quite comprehensive with the information you need to fully get to grips and operate the platform.

## Resources

- Red Hat OpenShift Documentation 
    - [OpenShift Container Platform security and compliance](https://docs.openshift.com/container-platform/4.12/security/index.html)
    - [Understanding container security](https://docs.openshift.com/container-platform/4.12/security/container_security/security-understanding.html#security-understanding)
- [Red Hat OpenShift security guide (ebook)](https://www.redhat.com/en/resources/openshift-security-guide-ebook)