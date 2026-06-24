# OpenShift Projects - Creation, Configuration and Governance

## Understanding OpenShift Projects: How They Differ from Kubernetes Namespaces

Red Hat OpenShift adds many features to simplify and enhance the management of Kubernetes clusters. One such feature is OpenShift Projects, which are similar to Kubernetes Namespaces but with added benefits tailored to the enterprise environment. In this post, we will explore the concept of OpenShift Projects, how they differ from Kubernetes Namespaces, and provide examples of creating and configuring Projects.

### OpenShift Projects: A Brief Overview

OpenShift Projects are an abstraction layer built on top of Kubernetes Namespaces. They provide a convenient way to organize and manage resources within an OpenShift cluster, and they offer additional features such as:

- Simplified multi-tenancy: Projects enable better isolation between users and teams, ensuring that each group works within its own environment without impacting others.
- Access control: Projects facilitate role-based access control (RBAC), allowing administrators to define and manage user permissions at the project level.
- Resource quotas and limits: Projects support setting resource quotas and limits to prevent overconsumption of cluster resources by individual projects.

## Creating and Configuring an OpenShift Project

Let's walk through the process of creating and configuring an OpenShift Project.

1. Create a new project:

To create a new project, use the oc new-project command:

````sh

$ oc new-project my-sample-project --description="My Sample OpenShift Project" --display-name="Sample Project"
````
This command creates a new project called my-sample-project with a description and display name.

2. Switch between projects:

You can switch between projects using the oc project command:

````sh
$ oc project my-sample-project
````

This command sets the active project to my-sample-project.

3. Configure resource quotas:

You can apply resource quotas to your project to limit the consumption of resources. Create a file called resource-quota.yaml with the following content:

````yaml

apiVersion: v1
kind: ResourceQuota
metadata:
  name: my-resource-quota
spec:
  hard:
    requests.cpu: "2"
    requests.memory: 2Gi
    limits.cpu: "4"
    limits.memory: 4Gi
````
## Adding Resource Quotas to projects

To apply the resource quota to your project:

````sh
$ oc apply -f resource-quota.yaml -n my-sample-project
````

This command applies the resource quota to the my-sample-project, limiting the total CPU and memory consumption for the project.

5. Configure role-based access control (RBAC):

To manage access control for your project, you can define and assign roles to users. For example, create a file called developer-role.yaml with the following content:

````yaml

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: developer
rules:
- apiGroups: [""]
  resources: ["pods", "services", "configmaps", "persistentvolumeclaims"]
  verbs: ["create", "get", "list", "watch", "update", "delete"]
  ````

Apply the role to your project:

````sh
$ oc apply -f developer-role.yaml -n my-sample-project
````

Now, you can grant the developer role to a specific user:

````sh
$ oc policy add-role-to-user developer my-user -n my-sample-project
````
This command grants the developer role to my-user in the my-sample-project.

## Adding SCC to a project

Remember in [Day 59](/2023/day59.md), we covered the Security Context Contraints, and how they provide security against the workloads we run inside the cluster, in the examples I provided, we fixed the security violation of the workload (pod) by ensuring the Service Account that it uses, is added to the correct SCC policy.

In this example, I'm going to which SCC at the project level, so that any workloads deployed to this project, conform to the correct policy.

1. Create a new project and change to that projects context

````sh
$ oc new-project scc-ns-test

$ oc project ssc-ns-test
````

2. Create a file called ```ngnix.yaml``` with the below content

````yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
````
2. Deploy an Ngnix Deployment to this project, and watch for the failure

````sh
$ oc apply -f ngnix.yaml

Warning: would violate PodSecurity "restricted:v1.24": allowPrivilegeEscalation != false (container "nginx" must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (container "nginx" must set securityContext.capabilities.drop=["ALL"]), runAsNonRoot != true (pod or container "nginx" must set securityContext.runAsNonRoot=true), seccompProfile (pod or container "nginx" must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")
deployment.apps/nginx-deployment created
````
As per Day 59's example, the deployment is created, but the pod will not be running.

3. Let's inspect the project's configuration before we continue further

````sh
$ oc get project scc-ns-test -o json
````
````json
{
    "apiVersion": "project.openshift.io/v1",
    "kind": "Project",
    "metadata": {
        "annotations": {
            "openshift.io/description": "",
            "openshift.io/display-name": "",
            "openshift.io/requester": "system:admin",
            "openshift.io/sa.scc.mcs": "s0:c27,c4",
            "openshift.io/sa.scc.supplemental-groups": "1000710000/10000",
            "openshift.io/sa.scc.uid-range": "1000710000/10000"
        },
        "creationTimestamp": "2023-03-29T09:23:18Z",
        "labels": {
            "kubernetes.io/metadata.name": "scc-ns-test",
            "pod-security.kubernetes.io/audit": "restricted",
            "pod-security.kubernetes.io/audit-version": "v1.24",
            "pod-security.kubernetes.io/warn": "restricted",
            "pod-security.kubernetes.io/warn-version": "v1.24"
        },
        "name": "scc-ns-test",
        "resourceVersion": "11247602",
        "uid": "3f720113-1e30-4a3f-b97e-48f88735e510"
    },
    "spec": {
        "finalizers": [
            "kubernetes"
        ]
    },
    "status": {
        "phase": "Active"
    }
}
````

Note that under ```labels``` section, we have several pod-security settings that are specified by default. That is because the most restrictive policy is applied to namespaces and their workloads by default. 

4. Now let's delete our deployment

````sh
$ oc delete -f nginx.yaml
````

5. Let's alter the configuration of this project to consume the ````privileged```` SCC, allowing us to brute force our pod to run. (In the real world, we would create an appropriate SCC and this use, rather than giving workloads god mode type access)

We are going to use the `oc patch` command and pass in the modifications to the labels. There are two ways to achieve this using the patch argument, we can either pass in the changes within the command line in JSON format, or we can pass in a file that is either JSON or YAML content. I'll detail both options below. 

This first command passes JSON content as part of the executions to alert the configuration.

````sh
$ oc patch namespace/scc-ns-test -p '{"metadata":{"labels":{"pod-security.kubernetes.io/audit":"privileged","pod-security.kubernetes.io/enforce":"privileged","pod-security.kubernetes.io/warn":"privileged","security.openshift.io/scc.podSecurityLabelSync":"false"}}}'
````

To break this down further, from the above example showing the Project configuration in JSON, we are altering the "audit", "warn" and "enforce" settings for Pod-Security to have the "privileged" value, and we also add a new label called "security.openshift.io/scc.podSecurityLabelSync" with a value of false. This stops the security admission controller from overwriting our changes. As the default SCC enforced is "restricted".

Rather than including the JSON changes in the same command line, which can get very long if you have a lot of changes, you can simply create a JSON or YAML file, containing content such as below and then apply it using the ```--patch-file``` argument. 

````yaml
metadata:
  labels:
    pod-security.kubernetes.io/audit: privileged
    pod-security.kubernetes.io/enforce: privileged
    pod-security.kubernetes.io/warn: privileged
    security.openshift.io/scc.podSecurityLabelSync: false
````

````sh
oc patch namespace/scc-ns-test  --patch-file ns-patch.yaml
````

5. Now if we inspect our Project, we will see the changes in effect.

````sh
oc get project scc-ns-test -o json
````

```json
{
    "apiVersion": "project.openshift.io/v1",
    "kind": "Project",
    "metadata": {
        "annotations": {
            "openshift.io/description": "",
            "openshift.io/display-name": "",
            "openshift.io/requester": "system:admin",
            "openshift.io/sa.scc.mcs": "s0:c27,c4",
            "openshift.io/sa.scc.supplemental-groups": "1000710000/10000",
            "openshift.io/sa.scc.uid-range": "1000710000/10000"
        },
        "creationTimestamp": "2023-03-29T09:23:18Z",
        "labels": {
            "kubernetes.io/metadata.name": "scc-ns-test",
            "pod-security.kubernetes.io/audit": "privileged",
            "pod-security.kubernetes.io/audit-version": "v1.24",
            "pod-security.kubernetes.io/enforce": "privileged",
            "pod-security.kubernetes.io/warn": "privileged",
            "pod-security.kubernetes.io/warn-version": "v1.24",
            "security.openshift.io/scc.podSecurityLabelSync": "false"
        },
        "name": "scc-ns-test",
        "resourceVersion": "11479286",
        "uid": "3f720113-1e30-4a3f-b97e-48f88735e510"
    },
    "spec": {
        "finalizers": [
            "kubernetes"
        ]
    },
    "status": {
        "phase": "Active"
    }
}
````

6. Redeploy the nginx instances or other containers you've been working with.

# Summary

There is just so much to cover, but hopefully you've now learned that Projects are more than just a Kubernetes Namespace with a different name. One of the areas we didn't cover, is the ability to [control Project creation](https://docs.openshift.com/container-platform/4.12/applications/projects/configuring-project-creation.html) by OpenShift users, either from a governed default template, or simply removing the ability for self-service access to create templates.

On [Day 61](/2023/day61.md), we shall cover the larger subject of RBAC within the cluster, and bring it back to applying access to projects.

## Resources

- Red Hat OpenShift Documentation - Building Applications - [Projects](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.12/html/building_applications/projects#doc-wrapper)