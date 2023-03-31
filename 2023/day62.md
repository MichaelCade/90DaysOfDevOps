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

### Installing the Compliance Operator

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
### Reviewing the profiles and the rules

We can see the installed profiles by running the below command, each profile has the product name that it applies to added as a prefix to the profile’s name. ocp4-e8 applies the Essential 8 benchmark to the OpenShift Container Platform product, while rhcos4-e8 applies the Essential 8 benchmark to the Red Hat Enterprise Linux CoreOS (RHCOS) product.

````sh
$ oc get -n openshift-compliance profiles.compliance
NAME                 AGE
ocp4-cis             107m
ocp4-cis-node        107m
ocp4-e8              107m
ocp4-high            107m
ocp4-high-node       107m
ocp4-moderate        107m
ocp4-moderate-node   107m
ocp4-nerc-cip        107m
ocp4-nerc-cip-node   107m
ocp4-pci-dss         107m
ocp4-pci-dss-node    107m
rhcos4-e8            107m
rhcos4-high          107m
rhcos4-moderate      107m
rhcos4-nerc-cip      107m
````

We can view the details of a profile by running:

````sh
$ oc get -n openshift-compliance -oyaml profiles.compliance ocp4-cis
````
````yaml
apiVersion: compliance.openshift.io/v1alpha1
description: This profile defines a baseline that aligns to the Center for Internet
  Security® Red Hat OpenShift Container Platform 4 Benchmark™, V1.1. This profile
  includes Center for Internet Security® Red Hat OpenShift Container Platform 4 CIS
  Benchmarks™ content. Note that this part of the profile is meant to run on the Platform
  that Red Hat OpenShift Container Platform 4 runs on top of. This profile is applicable
  to OpenShift versions 4.6 and greater.
id: xccdf_org.ssgproject.content_profile_cis
kind: Profile
metadata:
  annotations:
    compliance.openshift.io/image-digest: pb-ocp477wpm
    compliance.openshift.io/product: redhat_openshift_container_platform_4.1
    compliance.openshift.io/product-type: Platform
  creationTimestamp: "2023-03-31T09:09:52Z"
  generation: 1
  labels:
    compliance.openshift.io/profile-bundle: ocp4
  name: ocp4-cis
  namespace: openshift-compliance
  ownerReferences:
  - apiVersion: compliance.openshift.io/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: ProfileBundle
    name: ocp4
    uid: 19c2e4a5-094f-416a-a06b-eb0598e39618
  resourceVersion: "12971302"
  uid: 6dc3cca4-5649-43ae-8c46-614f82fd6744
rules:
- ocp4-accounts-restrict-service-account-tokens
- ocp4-accounts-unique-service-account
- ocp4-api-server-admission-control-plugin-alwaysadmit
- ocp4-api-server-admission-control-plugin-alwayspullimages
- ocp4-api-server-admission-control-plugin-namespacelifecycle
- ocp4-api-server-admission-control-plugin-noderestriction
- ocp4-api-server-admission-control-plugin-scc
- ocp4-api-server-admission-control-plugin-securitycontextdeny
- ocp4-api-server-admission-control-plugin-service-account
- ocp4-api-server-anonymous-auth
- ocp4-api-server-api-priority-flowschema-catch-all
- ocp4-api-server-api-priority-gate-enabled
- ocp4-api-server-audit-log-maxbackup
- ocp4-api-server-audit-log-maxsize
- ocp4-api-server-audit-log-path
- ocp4-api-server-auth-mode-no-aa
- ocp4-api-server-auth-mode-node
- ocp4-api-server-auth-mode-rbac
- ocp4-api-server-basic-auth
- ocp4-api-server-bind-address
...
- ocp4-scc-limit-privileged-containers
- ocp4-scc-limit-process-id-namespace
- ocp4-scc-limit-root-containers
- ocp4-scheduler-no-bind-address
- ocp4-secrets-consider-external-storage
- ocp4-secrets-no-environment-variables
- ocp4-version-detect-in-hypershift
- ocp4-version-detect-in-ocp
title: CIS Red Hat OpenShift Container Platform 4 Benchmark
````



### Running a Scan

Now that we have the operator installed, we have two key configurations:

- ScanSettings - This is the Schema for the scansettings API, and therefore you will provide configurations for running the scan, such as where it stores data, configurations for the container that is run on the platform to perform the scan, which components are subject to be scanned, and schedules to run the scan.
- ScanSettingsBinding - This is the Schema for the scansettingbindings API. This is used to bind the ScanSettings configuration to the compliance profiles you want to run against your chosen components.

There is a default ScanSettings called ```default``` supplied when you install the Compliance Operator, which you can run against your system.

````sh
$ oc describe scansettings default -n openshift-compliance

Name:                  default
Namespace:             openshift-compliance
Labels:                <none>
Annotations:           <none>
API Version:           compliance.openshift.io/v1alpha1
Kind:                  ScanSetting
Max Retry On Timeout:  3
Metadata:
  Creation Timestamp:  2023-03-31T09:09:33Z
  Generation:          1
  Resource Version:  12971055
  UID:               78945c1e-c323-40d8-87d9-c571275d58e3
Raw Result Storage:
  Node Selector:
    node-role.kubernetes.io/master:  
  Pv Access Modes:
    ReadWriteOnce
  Rotation:  3
  Size:      1Gi
  Tolerations:
    Effect:              NoSchedule
    Key:                 node-role.kubernetes.io/master
    Operator:            Exists
    Effect:              NoExecute
    Key:                 node.kubernetes.io/not-ready
    Operator:            Exists
    Toleration Seconds:  300
    Effect:              NoExecute
    Key:                 node.kubernetes.io/unreachable
    Operator:            Exists
    Toleration Seconds:  300
    Effect:              NoSchedule
    Key:                 node.kubernetes.io/memory-pressure
    Operator:            Exists
Roles:
  master
  worker
Scan Tolerations:
  Operator:           Exists
Schedule:             0 1 * * *
Show Not Applicable:  false
Strict Node Scan:     true
Timeout:              30m
Events:               <none>
````
You can read more details about the above [configuration here](https://docs.openshift.com/container-platform/4.12/security/compliance_operator/compliance-scans.html#running-compliance-scans_compliance-operator-scans), however, most of the lines are self-explanatory.

There is a second provided ScanSettings called ```default-auto-apply```, which follows the same premise as the above, however` will auto-remediate any findings as part of the scan.

Now we need to create a ScanSettingBinding to the CIS benchmark profiles. Create a YAML file using the below content called ````compliance-scansettingbinding.yaml```` and apply with ````oc apply -f compliance-scansettingbinding.yaml````.

````yaml
apiVersion: compliance.openshift.io/v1alpha1
kind: ScanSettingBinding
metadata:
  name: cis-compliance
  namespace: openshift-compliance
profiles:
  - name: ocp4-cis-node
    kind: Profile
    apiGroup: compliance.openshift.io/v1alpha1
  - name: ocp4-cis
    kind: Profile
    apiGroup: compliance.openshift.io/v1alpha1
settingsRef:
  name: default
  kind: ScanSetting
  apiGroup: compliance.openshift.io/v1alpha1
````

At this stage of the workflow, the Compliance Operator reconciles the ScanSettingBinding object, taking into account both the ```Binding``` and ```Bound``` settings. As a result, it generates a ComplianceSuite object along with the corresponding ComplianceScan objects to streamline the compliance evaluation process.


````sh
$ oc get compliancescan -n openshift-compliance
NAME                   PHASE         RESULT
ocp4-cis               AGGREGATING   NOT-AVAILABLE
ocp4-cis-node-master   AGGREGATING   NOT-AVAILABLE
ocp4-cis-node-worker   AGGREGATING   NOT-AVAILABLE
````
The scans will now work through the scanning phases and finish on the ```DONE``` phase once completed. At this point, you are probably going to find that the result is ```NON-COMPLIANT```. Now we can review the scan results and apply any necessary remediations to bring the cluster into compliance with the profile.

````sh
$ oc get compliancescan -n openshift-compliance
NAME                   PHASE   RESULT
ocp4-cis               DONE    NON-COMPLIANT
ocp4-cis-node-master   DONE    NON-COMPLIANT
ocp4-cis-node-worker   DONE    NON-COMPLIANT
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