# Deploying a Sample Application on Red Hat OpenShift: Handling Security Context Constraints (SCC)

On [Day 58](/day58.md) we finished looking around the developer and administrator interfaces of a newly deployed cluster.

In this submission (Day 59), we will walk through the process of deploying a sample MongoDB application to a newly deployed Red Hat OpenShift cluster. However, this deployment will fail due to the default security context constraints (SCC) in OpenShift. We will explain why the deployment fails, how to resolve the issue, and provide a brief overview of SCC in OpenShift with examples.

## Understanding Security Context Constraints (SCC)

Security context constraints in OpenShift are a security feature that allows administrators to control various aspects of the container runtime, such as user and group IDs, SELinux context, and the use of host resources. In Short, SCCs determine which security settings are allowed or disallowed for containerized applications. By default, OpenShift comes with several predefined SCCs, such as ```restricted```, ```anyuid```, and ```hostaccess```. These SCCs serve as templates for creating custom SCCs to meet specific security requirements.

> Warning: Do not modify the default SCCs. Customizing the default SCCs can lead to issues when some of the platform pods deploy or OpenShift Container Platform is upgraded. Additionally, the default SCC values are reset to the defaults during some cluster upgrades, which discards all customizations to those SCCs. Instead of modifying the default SCCs, create and modify your own SCCs as needed.

For example, the restricted SCC (default for most deployments, or restricted-v2 for new installs of OCP 4.11 and later) does not allow containers to run as root or with privileged access, while the ```anyuid``` SCC permits containers to run with any user ID, including root. By creating custom SCCs and granting them to service accounts or users, administrators can ensure that applications adhere to the desired security policies without compromising functionality.

Security context constraints allow an administrator to control:

    - Whether a pod can run privileged containers with the allowPrivilegedContainer flag

    - Whether a pod is constrained with the allowPrivilegeEscalation flag

    - The capabilities that a container can request

    - The use of host directories as volumes

    - The SELinux context of the container

    - The container user ID

    - The use of host namespaces and networking

    - The allocation of an FSGroup that owns the pod volumes

    - The configuration of allowable supplemental groups

    - Whether a container requires to write access to its root file system

    - The usage of volume types

    - The configuration of allowable seccomp profiles

To learn more details about what each of the out-of-the-box default security context constraints does, see [this official documentation page](https://docs.openshift.com/container-platform/4.12/authentication/managing-security-context-constraints.html#default-sccs_configuring-internal-oauth).

![Red Hat OpenShift - oc get scc](/2023/images/day59-Red%20Hat%20OpenShift%20-%20oc%20get%20scc.jpg)

### Anatomy of a Security Context Constraint configuration

SCCs consist of settings and strategies that control the security features that a pod has access to. These settings fall into three categories:

- Controlled by a boolean
  - Fields of this type default to the most restrictive value. For example;
    - ```AllowPrivilegedContainer``` is always set to ```false``` if unspecified.
- Controlled by an allowable set
  - Fields of this type are checked against the set to ensure their value is allowed.
- Controlled by a strategy
  - Items that have a strategy to generate a value provide:
    - A mechanism to generate the value, and
    - A mechanism to ensure that a specified value falls into the set of allowable values.

CRI-O has the following [default list of capabilities](https://github.com/cri-o/cri-o/blob/main/docs/crio.conf.5.md#crioruntime-table) that are allowed for each container of a pod:

````
  default_capabilities = [
	  "CHOWN",
	  "DAC_OVERRIDE",
	  "FSETID",
	  "FOWNER",
	  "SETGID",
	  "SETUID",
	  "SETPCAP",
	  "NET_BIND_SERVICE",
	  "KILL",
  ]
````

You can learn more about Linux capabilities [here](https://linuxera.org/container-security-capabilities-seccomp/) and [here](https://man7.org/linux/man-pages/man7/capabilities.7.html). The containers use the capabilities from this default list, but pod manifest authors (the person writing the application YAML for Kubernetes) can alter the list by requesting additional capabilities or removing some of the default behaviors. To control the capabilities allowed or denied for Pods running in the cluster, use the ```allowedCapabilities```, ```defaultAddCapabilities```, and ```requiredDropCapabilities``` parameters in your SCC to control such requests from the pods.

#### Quick Snippet: configuring a pod with capabilities

You can [specify additional capabilities for your pod](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-capabilities-for-a-container) as per the below example. 

````yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo-4
spec:
  containers:
  - name: sec-ctx-4
    image: gcr.io/google-samples/node-hello:1.0
    securityContext:
      capabilities:
        add: ["NET_ADMIN", "SYS_TIME"]
````

Let's look at some of the default contexts in further detail.

### Example SCC Configurations

1. Restricted SCC:

Denies access to all host features and requires pods to be run with a UID, and SELinux context that are allocated to the namespace.

The restricted-v2 SCC:

    Ensures that pods cannot run as privileged

    Ensures that pods cannot mount host directory volumes

    Requires that a pod is run as a user in a pre-allocated range of UIDs

    Requires that a pod is run with a pre-allocated MCS label

    Allows pods to use any FSGroup

    Allows pods to use any supplemental group

    Ensures that no child process of a container can gain more privileges than its parent (AllowPrivilegeEscalation=False)

You can get this SCC configuration by running ```oc get scc restricted-v2 -o yaml```

````yaml
allowHostDirVolumePlugin: false
allowHostIPC: false
allowHostNetwork: false
allowHostPID: false
allowHostPorts: false
allowPrivilegeEscalation: true
allowPrivilegedContainer: false
allowedCapabilities: null
apiVersion: security.openshift.io/v1
defaultAddCapabilities: null
fsGroup:
  type: MustRunAs
groups: []
kind: SecurityContextConstraints
metadata:
  annotations:
    include.release.openshift.io/ibm-cloud-managed: "true"
    include.release.openshift.io/self-managed-high-availability: "true"
    include.release.openshift.io/single-node-developer: "true"
    kubernetes.io/description: restricted denies access to all host features and requires
      pods to be run with a UID, and SELinux context that are allocated to the namespace.
    release.openshift.io/create-only: "true"
  creationTimestamp: "2023-03-16T09:34:36Z"
  generation: 1
  name: restricted
  resourceVersion: "401"
  uid: 8ced4b4e-7fed-4369-a0b8-da40880f4a3d
priority: null
readOnlyRootFilesystem: false
requiredDropCapabilities:
- KILL
- MKNOD
- SETUID
- SETGID
runAsUser:
  type: MustRunAsRange
seLinuxContext:
  type: MustRunAs
supplementalGroups:
  type: RunAsAny
users: []
volumes:
- configMap
- downwardAPI
- emptyDir
- ephemeral
- persistentVolumeClaim
- projected
- secret
````

2. Privileged SCC:

Allows access to all privileged and host features and the ability to run as any user, any group, any FSGroup, and with any SELinux context.

The privileged SCC allows:

    Users to run privileged pods

    Pods to mount host directories as volumes

    Pods to run as any user

    Pods to run with any MCS label

    Pods to use the host’s IPC namespace

    Pods to use the host’s PID namespace

    Pods to use any FSGroup

    Pods to use any supplemental group

    Pods to use any seccomp profiles

    Pods to request any capabilities

You can get this SCC configuration by running ```oc get scc privileged -o yaml```

````yaml
allowHostDirVolumePlugin: true
allowHostIPC: true
allowHostNetwork: true
allowHostPID: true
allowHostPorts: true
allowPrivilegeEscalation: true
allowPrivilegedContainer: true
allowedCapabilities:
- '*'
allowedUnsafeSysctls:
- '*'
apiVersion: security.openshift.io/v1
defaultAddCapabilities: null
fsGroup:
  type: RunAsAny
groups:
- system:cluster-admins
- system:nodes
- system:masters
kind: SecurityContextConstraints
metadata:
  annotations:
    include.release.openshift.io/ibm-cloud-managed: "true"
    include.release.openshift.io/self-managed-high-availability: "true"
    include.release.openshift.io/single-node-developer: "true"
    kubernetes.io/description: 'privileged allows access to all privileged and host
      features and the ability to run as any user, any group, any fsGroup, and with
      any SELinux context.  WARNING: this is the most relaxed SCC and should be used
      only for cluster administration. Grant with caution.'
    release.openshift.io/create-only: "true"
  creationTimestamp: "2023-03-16T09:34:35Z"
  generation: 1
  name: privileged
  resourceVersion: "398"
  uid: 19a16cc2-ce1f-4037-b70e-49ba261cb599
priority: null
readOnlyRootFilesystem: false
requiredDropCapabilities: null
runAsUser:
  type: RunAsAny
seLinuxContext:
  type: RunAsAny
seccompProfiles:
- '*'
supplementalGroups:
  type: RunAsAny
users:
- system:admin
- system:serviceaccount:openshift-infra:build-controller
volumes:
- '*'
````
Now let's look at some specific items from the above YAML:

- **allowedCapabilities:** - A list of capabilities that a pod can request. An empty list means that none of the capabilities can be requested while the special symbol * allows any capabilities.
- **defaultAddCapabilities: []** - A list of additional capabilities that are added to any pod.
- **fsGroup:** - The FSGroup strategy, dictates the allowable values for the security context.
- **groups** - The groups that can access this SCC.
- **requiredDropCapabilities** A list of capabilities to drop from a pod. Or, specify ALL to drop all capabilities.
- **runAsUser:** - The runAsUser strategy type, which dictates the allowable values for the security context.
- **seLinuxContext:** - The seLinuxContext strategy type, dictates the allowable values for the security context.
- **supplementalGroups** - The supplementalGroups strategy, dictates the allowable supplemental groups for the security context.
- **users:** - The users who can access this SCC.
- **volumes:** - The allowable volume types for the security context. In the example, * allows the use of all volume types.

The users and groups fields on the SCC control which users can access the SCC. By default, cluster administrators, nodes, and the build controller are granted access to the privileged SCC. All authenticated users are granted access to the restricted-v2 SCC.

## Let's deploy a sample application

I'm going to deploy some of the basic components of my [trusty Pac-Man application for Kubernetes](https://github.com/saintdle/pacman-tanzu). The MongoDB deployment, PVC and Secret.

First, I need to create the namespace to place the components in, ```oc create ns pacman```.

Now I apply the below YAML file ```oc apply -f mongo-test.yaml```

````yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    name: mongo
  name: mongo
  namespace: pacman
  annotations:
    source: "https://github.com/saintdle/pacman-tanzu"
spec:
  replicas: 1
  selector:
    matchLabels:
      name: mongo
  template:
    metadata:
      labels:
        name: mongo
    spec:
      initContainers:
      - args:
        - |
          mkdir -p /bitnami/mongodb
          chown -R "1001:1001" "/bitnami/mongodb"
        command:
        - /bin/bash
        - -ec
        image: docker.io/bitnami/bitnami-shell:10-debian-10-r158
        imagePullPolicy: Always
        name: volume-permissions
        resources: {}
        securityContext:
          runAsUser: 0
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /bitnami/mongodb
          name: mongo-db
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext:
        fsGroup: 1001
      serviceAccountName: default
      terminationGracePeriodSeconds: 30
      volumes:
      - name: mongo-db
        persistentVolumeClaim:
          claimName: mongo-storage
      containers:
      - image: bitnami/mongodb:4.4.8
        name: mongo
        env:
        - name: MONGODB_ROOT_PASSWORD
          valueFrom:
            secretKeyRef:
              key: database-admin-password
              name: mongodb-users-secret
        - name: MONGODB_DATABASE
          valueFrom:
            secretKeyRef:
              key: database-name
              name: mongodb-users-secret
        - name: MONGODB_PASSWORD
          valueFrom:
            secretKeyRef:
              key: database-password
              name: mongodb-users-secret
        - name: MONGODB_USERNAME
          valueFrom:
            secretKeyRef:
              key: database-user
              name: mongodb-users-secret
        ports:
        - name: mongo
          containerPort: 27017
        volumeMounts:
          - name: mongo-db
            mountPath: /bitnami/mongodb/
---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: mongo-storage
  namespace: pacman
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: v1
kind: Secret
metadata:
  name: mongodb-users-secret
  namespace: pacman
type: Opaque 
data:
  database-admin-name: Y2x5ZGU=
  database-admin-password: Y2x5ZGU=
  database-name: cGFjbWFu
  database-password: cGlua3k=
  database-user: Ymxpbmt5
````
Once applied, I see the following output:

>Warning: would violate PodSecurity "restricted:v1.24": allowPrivilegeEscalation != false (containers "volume-permissions", "mongo" must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (containers "volume-permissions", "mongo" must set securityContext.capabilities.drop=["ALL"]), runAsNonRoot != true (pod or containers "volume-permissions", "mongo" must set securityContext.runAsNonRoot=true), runAsUser=0 (container "volume-permissions" must not set runAsUser=0), seccompProfile (pod or containers "volume-permissions", "mongo" must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")
>
>deployment.apps/mongo created
>
>secret/mongodb-users-secret created

If I now inspect the deployment and replicaset in the ```pacman``` namespace, we'll see that's stuck, I have no pods running.

````
# oc get all -n pacman

NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mongo   0/1     0            0           3m9s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/mongo-56cc764fb   1         0         0       3m9s
````


## Why the Deployment Fails

The provided Kubernetes application includes an initContainer with the following security context:

````yaml
securityContext:
  runAsUser: 0
````

This configuration means that the initContainer will attempt to run as the root user (UID 0). However, OpenShift's default SCCs restrict the use of the root user for security reasons. As a result, the deployment fails because it violates the default security context constraints. The same is true of the other configuration settings mentioned in the above output as well. Remember in OCP 4.11 and later (new installs), the default SCC is the restricted-v2 policy.

## How to Resolve the Issue

To resolve this issue, we need to modify the deployment configuration to comply with the SCC policies in OpenShift. There are several ways to achieve this, but in this example, we will create a custom SCC that allows the initContainer to run as root. Follow these steps:

1. Create a new custom SCC, and save the below YAML in a file called mongo-custom-scc.yaml:

````yaml
apiVersion: security.openshift.io/v1
kind: SecurityContextConstraints
metadata:
  name: mongo-custom-scc
allowPrivilegedContainer: false
allowHostNetwork: false
allowHostPorts: false
allowHostPID: false
allowHostIPC: false
runAsUser:
  type: RunAsAny
seLinuxContext:
  type: MustRunAs
fsGroup:
  type: RunAsAny
supplementalGroups:
  type: RunAsAny
````

2. Apply the custom SCC to your OpenShift cluster:

````sh
oc apply -f mongo-custom-scc.yaml
````

3. Grant the mongo-custom-scc SCC to the service account that the MongoDB deployment is using:

````sh
oc adm policy add-scc-to-user mongo-custom-scc system:serviceaccount:<namespace>:default

# In my environment, I run:
oc adm policy add-scc-to-user mongo-custom-scc system:serviceaccount:pacman:default
````

Replace <namespace> with the namespace where your MongoDB deployment is located.

4. Redeploy the MongoDB application.

````
# oc scale deploy mongo -n pacman --replicas=0

deployment.apps/mongo scaled

# oc scale deploy mongo -n pacman --replicas=1

deployment.apps/mongo scaled
````
In the real world, the first port of call should always be to work to ensure your containers and applications run with the least privileges necessary and therefore don't need to run as root.

If they do need some sort of privilege, then defining tight RBAC and SCC control in place is key.

# Summary

In this post, we discussed how the default security context constraints in OpenShift can prevent deployments from running as expected. We provided a solution to the specific issue of running an initContainer as root for a MongoDB application. Understanding and managing SCCs in OpenShift is essential for maintaining secure and compliant applications within your cluster.

On [Day 60](/day60.md)](/day60.md), we will look at OpenShift Projects Creation, Configuration and Governance, for example consuming SCC via the project level, and other features of Red Hat OpenShift.

## Resources

- Red Hat OpenShift - [Managing security context constraints](https://docs.openshift.com/container-platform/4.12/authentication/managing-security-context-constraints.html)
- Red Hat Blog - [Managing SCCs in OpenShift](https://cloud.redhat.com/blog/managing-sccs-in-openshift)
- Linuxera - [Capabilities and Seccomp Profiles on Kubernetes](https://linuxera.org/capabilities-seccomp-kubernetes/)
- Red Hat OpenShift - [Important OpenShift changes to Pod Security Standards](https://connect.redhat.com/en/blog/important-openshift-changes-pod-security-standards)
  - OCP 4.11 and later - With the introduction of a new built-in admission controller that enforces the [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/), namespaces and pods can be defined with three different policies: Privileged, Baseline and Restricted. Therefore, pods not configured according to the enforced security standards defined globally, or on the namespace level, will not be admitted and will not run.
- [Pods fail to create due to "allowPrivilegeEscalation: true" in OpenShift 4.11](https://access.redhat.com/solutions/6976492)
- [Using the legacy restricted SCC in OCP 4.11+](https://access.redhat.com/articles/6973044)
- [Role-based access to security context constraints](https://docs.openshift.com/container-platform/4.12/authentication/managing-security-context-constraints.html#role-based-access-to-ssc_configuring-internal-oauth)
  - You can specify SCCs as resources that are handled by RBAC. This allows you to scope access to your SCCs to a certain project or to the entire cluster.
- Kubernetes.io - [Configure a Security Context for a Pod or Container](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/)