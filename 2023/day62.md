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

To read the descriptions of one of the rules that form the policy:

````sh
$ oc get -n openshift-compliance rules ocp4-api-server-anonymous-auth -oyaml
````

````yaml
apiVersion: compliance.openshift.io/v1alpha1
checkType: Platform
description: |-
  By default, anonymous access to the OpenShift API is enabled, but at the same time, all requests must be authorized. If no authentication mechanism is used, the request is assigned the system:anonymous virtual user and the system:unauthenticated virtual group. This allows the authorization layer to determin which requests, if any, is an anonymous user authorized to make. To verify the authorization rules for anonymous requests run the following:

  $ oc describe clusterrolebindings

  and inspect the bidnings of the system:anonymous virtual user and the system:unauthenticated virtual group. To test that an anonymous request is authorized to access the readyz endpoint, run:

  $ oc get --as="system:anonymous" --raw='/readyz?verbose'

  In contrast, a request to list all projects should not be authorized:

  $ oc get --as="system:anonymous" projects
id: xccdf_org.ssgproject.content_rule_api_server_anonymous_auth
instructions: |-
  Run the following command to view the authorization rules for anonymous requests:
  $ oc describe clusterrolebindings
  Make sure that there exists at least one clusterrolebinding that binds
  either the system:unauthenticated group or the system:anonymous
  user.
  To test that an anonymous request is authorized to access the readyz
  endpoint, run:
  $ oc get --as="system:anonymous" --raw='/readyz?verbose'
  In contrast, a request to list all projects should not be authorized:
  $ oc get --as="system:anonymous" projects
kind: Rule
metadata:
  annotations:
    compliance.openshift.io/image-digest: pb-ocp477wpm
    compliance.openshift.io/rule: api-server-anonymous-auth
    control.compliance.openshift.io/CIS-OCP: 1.2.1
    control.compliance.openshift.io/NERC-CIP: CIP-003-8 R6;CIP-004-6 R3;CIP-007-3
      R6.1
    control.compliance.openshift.io/NIST-800-53: CM-6;CM-6(1)
    control.compliance.openshift.io/PCI-DSS: Req-2.2
    policies.open-cluster-management.io/controls: 1.2.1,CIP-003-8 R6,CIP-004-6 R3,CIP-007-3
      R6.1,CM-6,CM-6(1),Req-2.2
    policies.open-cluster-management.io/standards: CIS-OCP,NERC-CIP,NIST-800-53,PCI-DSS
  creationTimestamp: "2023-03-31T09:09:53Z"
  generation: 1
  labels:
    compliance.openshift.io/profile-bundle: ocp4
  name: ocp4-api-server-anonymous-auth
  namespace: openshift-compliance
  ownerReferences:
  - apiVersion: compliance.openshift.io/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: ProfileBundle
    name: ocp4
    uid: 19c2e4a5-094f-416a-a06b-eb0598e39618
  resourceVersion: "12971338"
  uid: 12db5786-4ff6-4e80-90e0-2b370541f6e1
rationale: When enabled, requests that are not rejected by other configured authentication
  methods are treated as anonymous requests. These requests are then served by the
  API server. If you are using RBAC authorization, it is generally considered reasonable
  to allow anonymous access to the API Server for health checks and discovery purposes,
  and hence this recommendation is not scored. However, you should consider whether
  anonymous discovery is an acceptable risk for your purposes.
severity: medium
title: Ensure that anonymous requests to the API Server are authorized
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

Let's look at the first scan in more detail, as we also want to capture the label attached, so we can get the full results to view.

````sh
$ oc get compliancescan -n openshift-compliance ocp4-cis -o yaml
````

````yaml
oc get compliancescan -n openshift-compliance ocp4-cis -o yaml
apiVersion: compliance.openshift.io/v1alpha1
kind: ComplianceScan
metadata:
  creationTimestamp: "2023-03-31T10:49:41Z"
  finalizers:
  - scan.finalizers.compliance.openshift.io
  generation: 1
  labels:
    compliance.openshift.io/suite: cis-compliance
  name: ocp4-cis
  namespace: openshift-compliance
  ownerReferences:
  - apiVersion: compliance.openshift.io/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: ComplianceSuite
    name: cis-compliance
    uid: c40db43f-0635-4731-b537-5ad3fc08cc06
  resourceVersion: "13035410"
  uid: 9d5b10b8-f67c-44f3-b13d-0765cb037091
spec:
  content: ssg-ocp4-ds.xml
  contentImage: registry.redhat.io/compliance/openshift-compliance-content-rhel8@sha256:c4bf5b2b20ff538adbc430b7ee993fbd7c291203a9810534005148304e3b169b
  maxRetryOnTimeout: 3
  profile: xccdf_org.ssgproject.content_profile_cis
  rawResultStorage:
    nodeSelector:
      node-role.kubernetes.io/master: ""
    pvAccessModes:
    - ReadWriteOnce
    rotation: 3
    size: 1Gi
    tolerations:
    - effect: NoSchedule
      key: node-role.kubernetes.io/master
      operator: Exists
    - effect: NoExecute
      key: node.kubernetes.io/not-ready
      operator: Exists
      tolerationSeconds: 300
    - effect: NoExecute
      key: node.kubernetes.io/unreachable
      operator: Exists
      tolerationSeconds: 300
    - effect: NoSchedule
      key: node.kubernetes.io/memory-pressure
      operator: Exists
  scanTolerations:
  - operator: Exists
  scanType: Platform
  showNotApplicable: false
  strictNodeScan: true
  timeout: 30m
status:
  conditions:
  - lastTransitionTime: "2023-03-31T10:51:20Z"
    message: Compliance scan run is done running the scans
    reason: NotRunning
    status: "False"
    type: Processing
  - lastTransitionTime: "2023-03-31T10:51:20Z"
    message: Compliance scan run is done and has results
    reason: Done
    status: "True"
    type: Ready
  phase: DONE
  remainingRetries: 3
  result: NON-COMPLIANT
  resultsStorage:
    name: ocp4-cis
    namespace: openshift-compliance
  warnings: |-
    could not fetch /apis/flowcontrol.apiserver.k8s.io/v1alpha1/flowschemas/catch-all: the server could not find the requested resource
    could not fetch /apis/logging.openshift.io/v1/namespaces/openshift-logging/clusterlogforwarders/instance: the server could not find the requested resource
    could not fetch /apis/apps/v1/namespaces/openshift-sdn/daemonsets/sdn: daemonsets.apps "sdn" not found
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch "/api/v1/nodes/NODE_NAME/proxy/configz": the server could not find the requested resource
    could not fetch /apis/hypershift.openshift.io/v1beta1/namespaces/clusters/hostedclusters/None: the server could not find the requested resource
    Kubelet configs for 90days-ocp-72ptq-master-1 are not consistent with role master, Diff: [{"op":"replace","path":"/address","value":"192.168.200.183"}] of KubeletConfigs for master role will not be saved.
    Kubelet configs for 90days-ocp-72ptq-worker-x7v4j are not consistent with role worker, Diff: [{"op":"replace","path":"/address","value":"192.168.200.194"}] of KubeletConfigs for worker role will not be saved.
    Kubelet configs for 90days-ocp-72ptq-master-2 are not consistent with role master, Diff: [{"op":"add","path":"/address","value":"192.168.200.181"}] of KubeletConfigs for master role will not be saved.
    Kubelet configs for 90days-ocp-72ptq-master-0 are not consistent with role control-plane, Diff: [{"op":"replace","path":"/address","value":"192.168.200.185"}] of KubeletConfigs for control-plane role will not be saved.
    Kubelet configs for 90days-ocp-72ptq-master-2 are not consistent with role control-plane, Diff: [{"op":"add","path":"/address","value":"192.168.200.181"}] of KubeletConfigs for control-plane role will not be saved.
    Kubelet configs for 90days-ocp-72ptq-worker-5cgp8 are not consistent with role worker, Diff: [{"op":"add","path":"/address","value":"192.168.200.187"}] of KubeletConfigs for worker role will not be saved.
````

In the output, we can see the warnings from the results straight away. To view the full results we can run the command below using the label from the above output. Alternatively, I can list all failed results across the conducted scans by running the command:

````sh
$ oc get compliancecheckresults -n openshift-compliance -l 'compliance.openshift.io/check-status=FAIL'
NAME                                                          STATUS   SEVERITY
ocp4-cis-api-server-encryption-provider-cipher                FAIL     medium
ocp4-cis-api-server-encryption-provider-config                FAIL     medium
ocp4-cis-audit-log-forwarding-enabled                         FAIL     medium
ocp4-cis-configure-network-policies-namespaces                FAIL     high
ocp4-cis-kubeadmin-removed                                    FAIL     medium
ocp4-cis-node-master-kubelet-enable-protect-kernel-defaults   FAIL     medium
ocp4-cis-node-master-kubelet-enable-protect-kernel-sysctl     FAIL     medium
ocp4-cis-node-worker-kubelet-enable-protect-kernel-defaults   FAIL     medium
ocp4-cis-node-worker-kubelet-enable-protect-kernel-sysctl     FAIL     medium
````

We can look at the individual ```ComplianceRemediation``` in more detail with the command;

````sh
 oc get compliancecheckresults -n openshift-compliance ocp4-cis-audit-log-forwarding-enabled  -o yaml
````

Which provides us the following output:

````yaml
apiVersion: compliance.openshift.io/v1alpha1
description: |-
  Ensure that Audit Log Forwarding Is Enabled
  OpenShift audit works at the API server level, logging all requests coming to the server. Audit is on by default and the best practice is to ship audit logs off the cluster for retention. The cluster-logging-operator is able to do this with the

  ClusterLogForwarders

  resource. The forementioned resource can be configured to logs to different third party systems. For more information on this, please reference the official documentation: https://docs.openshift.com/container-platform/4.6/logging/cluster-logging-external.html
id: xccdf_org.ssgproject.content_rule_audit_log_forwarding_enabled
instructions: |-
  Run the following command:
  oc get clusterlogforwarders instance -n openshift-logging -ojson | jq -r '.spec.pipelines[].inputRefs | contains(["audit"])'
  The output should return true.
kind: ComplianceCheckResult
metadata:
  annotations:
    compliance.openshift.io/rule: audit-log-forwarding-enabled
  creationTimestamp: "2023-03-31T10:51:02Z"
  generation: 1
  labels:
    compliance.openshift.io/check-severity: medium
    compliance.openshift.io/check-status: FAIL
    compliance.openshift.io/scan-name: ocp4-cis
    compliance.openshift.io/suite: cis-compliance
  name: ocp4-cis-audit-log-forwarding-enabled
  namespace: openshift-compliance
  ownerReferences:
  - apiVersion: compliance.openshift.io/v1alpha1
    blockOwnerDeletion: true
    controller: true
    kind: ComplianceScan
    name: ocp4-cis
    uid: 9d5b10b8-f67c-44f3-b13d-0765cb037091
  resourceVersion: "13034914"
  uid: c5b7342d-af35-4944-86bd-8a4f907acccc
rationale: Retaining logs ensures the ability to go back in time to investigate or
  correlate any events. Offloading audit logs from the cluster ensures that an attacker
  that has access to the cluster will not be able to tamper with the logs because
  of the logs being stored off-site.
severity: medium
status: FAIL
````

This particular remediation has no automatic remediation, but you can see [another example here](https://docs.openshift.com/container-platform/4.12/security/compliance_operator/compliance-operator-remediation.html#compliance-review_compliance-remediation), whereby a ```MachineConfig``` resource will be applied by the remediation. Applying remediation can take different forms, depending on what the remediation is. Therefore I won't dive into this detail today, I think we've covered enough to understand how to run and view compliance checks against our platform. You can read more about [remediations here](https://docs.openshift.com/container-platform/4.12/security/compliance_operator/compliance-operator-remediation.html#compliance-applying_compliance-remediation). The final point I wanted to highlight on this topic, is the ```oc-compliance``` plugin, which extends the functionality of the ```oc``` CLI tool, you can few the [details and how to use it here](https://docs.openshift.com/container-platform/4.12/security/compliance_operator/oc-compliance-plug-in-using.html).

# Red Hat Quay Container Security Operator

## Vulnerability Scanning Overview

Vulnerability scanning in Red Hat OpenShift refers to the process of inspecting container images for known security issues, such as outdated software packages, misconfigurations, or exposed sensitive information. The goal of vulnerability scanning is to identify and remediate potential security risks in container images before they are deployed to the OpenShift cluster, thus enhancing the overall security posture of the platform.

OpenShift provides several features and tools to facilitate vulnerability scanning:

- Integrated container registry: OpenShift includes a built-in container registry to store and manage container images. The integrated registry allows for a more streamlined and secure process when scanning images for vulnerabilities, as it eliminates the need to rely on external registries.

- ImageStreams: An ImageStream in OpenShift is an abstraction that represents a series of related container images, typically different versions of the same application. ImageStreams simplify the process of tracking and deploying container images, making it easier to apply vulnerability scanning and remediation across multiple versions of an application.

- Image signing and trust: OpenShift supports container image signing and the enforcement of signature-based trust policies. This feature ensures that only trusted and verified images can be deployed to the cluster, helping to prevent the deployment of images with known vulnerabilities.

- Third-party integrations: OpenShift can be easily integrated with external vulnerability scanning tools and platforms, such as Aqua Security, Sysdig, or Twistlock. These tools can be configured to automatically scan container images stored in the OpenShift registry and provide detailed reports on identified vulnerabilities and suggested remediation steps.

- OpenShift Operators: OpenShift supports the use of Operators, which are automated software extensions that manage applications and their components. Operators can be used to deploy and manage vulnerability scanning tools within the OpenShift cluster, ensuring a consistent and automated scanning process. Red Hat provides the ```Red Hat Quay Container Security Operator```, however, you can also implement third-party scanners such as [Trivy](https://github.com/aquasecurity/trivy) from [Aqua Security](https://aquasec.com/).

By leveraging these features and tools, Red Hat OpenShift enables organizations to perform comprehensive vulnerability scanning on container images, reducing the risk of security breaches and enhancing the overall security of the platform.

Focusing on the ```Red Hat Quay Container Security Operator```, this provides the following:

- Watches containers associated with pods on all or specified namespaces

- Queries the container registry where the containers came from for vulnerability information, provided an image’s registry is running image scanning (such as Quay.io or a Red Hat Quay registry with Clair scanning)

- Exposes vulnerabilities via the ImageManifestVuln object in the Kubernetes API

## Installing the Red Hat Quay Container Security Operator

This time I'm going to provide instructions on how to perform these steps in the Red Hat Console interface, rather than CLI.

1. In the administrator view within the console, navigate to Operators > OperatorHub and search for "Red Hat Quay Container Security Operator".
2. Select the tile, and click to install.
3. Confirm the settings
    -  All namespaces and automatic approval strategy are selected, by default.
4. Select Install. The Container Security Operator appears after a few moments on the Installed Operators screen.

When you now browse to the homepage of the Red Hat OpenShit Console, you will see an "Image vulnerability" component on the status tile, which you can select to see high-level information about your estate.

![Red Hat OpenShift Console - Image vulnerability](/2023/images/Day62%20-%20Red%20Hat%20Quay%20Container%20Security%20Operator/Red%20Hat%20Console%20-%20Image%20vulnerability.jpg)

You can click on the circle graph or the namespaces to see more details which takes you to the navigation page of Adminstration (1) > Image Vulnerabilities (2).

Now you can see a list of the vulnerabilities and you can change the project view for all projects or a specific one to curate this list.

![Red Hat OpenShift Console - Administration - Image vulnerabilities](/2023/images/Day62%20-%20Red%20Hat%20Quay%20Container%20Security%20Operator/Red%20Hat%20OpenShift%20Console%20-%20Administration%20-%20Image%20vulnerabilities.jpg)

Clicking the image name (3) will show further manifest details via the OpenShift console, including which pods are affected. Or you can click the Manifest details on the right-hand side (4) which will take you to the Quay Security Scanner report hosted on [Quay.io](https://Quay.io).

Below shows the Image manifest details including each vulnerability found with that image, and links to appropriate documentation, such as CVE information held on [access.redhat.com](https://access.redhat.com
)
![Red Hat OpenShift Console - Administration - Image vulnerabilities - Manifest details](/2023/images/Day62%20-%20Red%20Hat%20Quay%20Container%20Security%20Operator/Red%20Hat%20OpenShift%20Console%20-%20Administration%20-%20Image%20vulnerabilities%20-%20Manifest%20details.jpg)

Below is the ```affected pods``` tab view on the Image manifests page. 

![Red Hat OpenShift Console - Administration - Image vulnerabilities - Manifest details - Affected pods](/2023/images/Day62%20-%20Red%20Hat%20Quay%20Container%20Security%20Operator/Red%20Hat%20OpenShift%20Console%20-%20Administration%20-%20Image%20vulnerabilities%20-%20Manifest%20details%20-%20Affected%20pods.jpg)

And finally the Quay Security Scanner page for one of the images shown in the report. This was for my pacman application, you can see the Quay.io report yourself [here](https://quay.io/repository/ifont/pacman-nodejs-app/manifest/sha256:196ae9a1a33a2d32046a46739779ca273667f1d4f231f8a721e8064c3509405e?tab=vulnerabilities)(free sign up account required).

![Red Hat OpenShift Console - Administration - Image vulnerabilities - Quay Manifest](/2023/images/Day62%20-%20Red%20Hat%20Quay%20Container%20Security%20Operator/Red%20Hat%20OpenShift%20Console%20-%20Administration%20-%20Image%20vulnerabilities%20-%20Quay%20Manifest.jpg)

Of course we can also see information in the command line too, to see all vulnerabilities found, use the command:

```sh
oc get vuln --all-namespaces
NAMESPACE                                          NAME                                                                      AGE
openshift-apiserver-operator                       sha256.01974e4c0e0d112e09bee8fe2625d565d3d62fa42013b38d7ce43d2d40f6057a   20h
openshift-apiserver                                sha256.13640b919950fc648219c528ee7ed30262bae856566fbd6c4cb5e15ffd457d6f   20h
openshift-apiserver                                sha256.8829aefa24dd606d2fe3ff86b97858c07acedae5f5eb3f044c20395762e7c02b   20h
openshift-authentication-operator                  sha256.31b617cec5c22e187cc22da606fc6998ea3529b1b6e8d80d1799c3dc9705997e   20h
openshift-authentication                           sha256.41e06255fc823c0082a74466b69ccfb672947b7075ea43a10e729c5f39314d00   20h
openshift-cloud-controller-manager-operator        sha256.a7856b6371fc4a7ade8a678daca149db6c6a55ee7137d9e308721d2d3bebf364   20h
openshift-cloud-credential-operator                sha256.1986315effe0f3ee415e86df3a87765268ed1da405c7a297c278e1d7030286a4   20h
...
openshift-vsphere-infra                            sha256.ddf81e535cf7a6b2775f3db690ec1e6eaa1c7427a0f9b98ce120d8ad06520440   20h
test-app                                           sha256.196ae9a1a33a2d32046a46739779ca273667f1d4f231f8a721e8064c3509405e   20h
```

You can inspect the details of a specific vulnerability by running the command:

```sh
oc describe vuln -n {namespace} {sha ID of vuln}

# Example

oc describe vuln -n openshift-apiserver-operator sha256.01974e4c0e0d112e09bee8fe2625d565d3d62fa42013b38d7ce43d2d40f6057a
```

Example output:

```yaml
Name:         sha256.01974e4c0e0d112e09bee8fe2625d565d3d62fa42013b38d7ce43d2d40f6057a
Namespace:    openshift-apiserver-operator
Labels:       openshift-apiserver-operator/openshift-apiserver-operator-7bd84bd596-pgpct=true
Annotations:  <none>
API Version:  secscan.quay.redhat.com/v1alpha1
Kind:         ImageManifestVuln
Metadata:
  Creation Timestamp:  2023-03-30T19:01:40Z
  Generation:          17
  Resource Version:  13206497
  UID:               dbec02e5-e4c6-412f-b561-757237844d43
Spec:
  Features:
    Name:     pip
    Version:  9.0.3
    Vulnerabilities:
      Description:     Pip 21.1 updates its dependency 'urllib3' to v1.26.4 due to security issues.
      Metadata:        {"UpdatedBy": "pyupio", "RepoName": "pypi", "RepoLink": "https://pypi.org/simple", "DistroName": "", "DistroVersion": "", "NVD": {"CVSSv3": {"Vectors": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:L/I:L/A:N", "Score": 6.5}}}
      Name:            pyup.io-40291 (CVE-2021-28363)
      Namespace Name:  pyupio
      Severity:        Medium
      Description:     A flaw was found in python-pip in the way it handled Unicode separators in git references. A remote attacker could possibly use this issue to install a different revision on a repository. The highest threat from this vulnerability is to data integrity. This is fixed in python-pip version 21.1.
      Metadata:        {"UpdatedBy": "pyupio", "RepoName": "pypi", "RepoLink": "https://pypi.org/simple", "DistroName": "", "DistroVersion": "", "NVD": {"CVSSv3": {"Vectors": "CVSS:3.1/AV:N/AC:L/PR:L/UI:R/S:U/C:N/I:H/A:N", "Score": 5.7}}}
      Name:            pyup.io-42559 (CVE-2021-3572)
      Namespace Name:  pyupio
      Severity:        Medium
      Description:     Pip before 19.2 allows Directory Traversal when a URL is given in an install command, because a Content-Disposition header can have ../ in a filename, as demonstrated by overwriting the /root/.ssh/authorized_keys file. This occurs in _download_http_url in _internal/download.py.
      Metadata:        {"UpdatedBy": "pyupio", "RepoName": "pypi", "RepoLink": "https://pypi.org/simple", "DistroName": "", "DistroVersion": "", "NVD": {"CVSSv3": {"Vectors": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:H/A:N", "Score": 7.5}}}
      Name:            pyup.io-38765 (CVE-2019-20916)
      Namespace Name:  pyupio
      Severity:        High
    Name:              openssl-libs
    Version:           1:1.1.1k-7.el8_6
    Vulnerabilities:
      Description:  OpenSSL is a toolkit that implements the Secure Sockets Layer (SSL) and Transport Layer Security (TLS) protocols, as well as a full-strength general-purpose cryptography library.

Security Fix(es):

* openssl: X.400 address type confusion in X.509 GeneralName (CVE-2023-0286)

For more details about the security issue(s), including the impact, a CVSS score, acknowledgments, and other related information, refer to the CVE page(s) listed in the References section.
      Fixedby:         1:1.1.1k-8.el8_6
      Link:            https://access.redhat.com/errata/RHSA-2023:1441 https://access.redhat.com/security/cve/CVE-2023-0286
      Metadata:        {"UpdatedBy": "RHEL8-rhel-8.6-eus", "RepoName": "cpe:/o:redhat:rhel_eus:8.6::baseos", "RepoLink": null, "DistroName": "Red Hat Enterprise Linux Server", "DistroVersion": "8", "NVD": {"CVSSv3": {"Vectors": "CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:H/I:N/A:H", "Score": 7.4}}}
      Name:            RHSA-2023:1441: openssl security update (Important)
      Namespace Name:  RHEL8-rhel-8.6-eus
      Severity:        High
    Name:              urllib3
    Version:           1.24.2
    Vulnerabilities:
      Description:     Urllib3 1.26.5 includes a fix for CVE-2021-33503: When provided with a URL containing many @ characters in the authority component, the authority regular expression exhibits catastrophic backtracking, causing a denial of service if a URL were passed as a parameter or redirected to via an HTTP redirect.
      Metadata:        {"UpdatedBy": "pyupio", "RepoName": "pypi", "RepoLink": "https://pypi.org/simple", "DistroName": "", "DistroVersion": "", "NVD": {"CVSSv3": {"Vectors": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H", "Score": 7.5}}}
      Name:            pyup.io-43975 (CVE-2021-33503)
      Namespace Name:  pyupio
      Severity:        High
      Description:     urllib3 before 1.25.9 allows CRLF injection if the attacker controls the HTTP request method, as demonstrated by inserting CR and LF control characters in the first argument of putrequest(). NOTE: this is similar to CVE-2020-26116.
      Metadata:        {"UpdatedBy": "pyupio", "RepoName": "pypi", "RepoLink": "https://pypi.org/simple", "DistroName": "", "DistroVersion": "", "NVD": {"CVSSv3": {"Vectors": "CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:C/C:L/I:L/A:N", "Score": 7.2}}}
      Name:            pyup.io-38834 (CVE-2020-26137)
      Namespace Name:  pyupio
      Severity:        High
    Name:              setuptools
    Version:           39.2.0
    Vulnerabilities:
      Description:     Python Packaging Authority (PyPA) setuptools before 65.5.1 allows remote attackers to cause a denial of service via HTML in a crafted package or custom PackageIndex page. There is a Regular Expression Denial of Service (ReDoS) in package_index.py.
      Metadata:        {"UpdatedBy": "pyupio", "RepoName": "pypi", "RepoLink": "https://pypi.org/simple", "DistroName": "", "DistroVersion": "", "NVD": {"CVSSv3": {"Vectors": "CVSS:3.1/AV:N/AC:H/PR:N/UI:N/S:U/C:N/I:N/A:H", "Score": 5.9}}}
      Name:            pyup.io-52495 (CVE-2022-40897)
      Namespace Name:  pyupio
      Severity:        Medium
  Image:               quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256
  Manifest:            sha256:01974e4c0e0d112e09bee8fe2625d565d3d62fa42013b38d7ce43d2d40f6057a
Status:
  Affected Pods:
    openshift-apiserver-operator/openshift-apiserver-operator-7bd84bd596-pgpct:
      cri-o://dd4fcb700a95c041d19bf7829d3e07516ccf2a36522027f920d76ed0aa57f84c
  Fixable Count:     1
  High Count:        4
  Highest Severity:  High
  Last Update:       2023-03-31 15:31:32.853602342 +0000 UTC
  Medium Count:      3
Events:              <none>
```

# Summary

Whilst for this 2023 edition focusing on DevSecOps, we could have purely spent time focusing on Security and Compliance for Red Hat OpenShift in-depth, I wanted to start at a higher level, understanding why you would choose an enterprise Kubernetes offering, and what features will enhance your cloud-native platform. Hopefully, this has given you a solid understanding of this offering, as well as being able to understand the basics of how to run and operate it. Another area we only touched upon briefly was application deployment, instead focusing on the security posture of deploying workloads, rather than the methods of building and running the applications themselves. This topic of application build, deployment and management requires a whole section on its own.

I urge you to spend time reading through the official documentation for Red Hat OpenShift, it's quite comprehensive with the information you need to fully get to grips and operate the platform.

## Resources

- Red Hat OpenShift Documentation
  - [OpenShift Container Platform security and compliance](https://docs.openshift.com/container-platform/4.12/security/index.html)
  - [Understanding container security](https://docs.openshift.com/container-platform/4.12/security/container_security/security-understanding.html#security-understanding)
  - [Troubleshooting the Compliance Operator](https://docs.openshift.com/container-platform/4.12/security/compliance_operator/compliance-operator-troubleshooting.html)
  - [Running the Red Hat Quay Container Security Operator](https://docs.openshift.com/container-platform/4.12/security/pod-vulnerability-scan.html)
  - [Securing container content](https://docs.openshift.com/container-platform/4.12/security/container_security/security-container-content.html)
- [Red Hat OpenShift security guide (ebook)](https://www.redhat.com/en/resources/openshift-security-guide-ebook)
- YouTube - [ CVE and CVSS explained | Security Detail | Presented by Red Hat](https://www.youtube.com/watch?v=oSyEGkX6sX0)