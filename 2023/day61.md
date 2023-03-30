# Understanding Role-Based Access Control (RBAC) in Red Hat OpenShift: Control and Secure Your Cluster

## Introduction

Role-Based Access Control (RBAC) is an essential security feature in Red Hat OpenShift that enables administrators to manage and control access to resources within a cluster. In this blog post, we will explore RBAC in OpenShift, discuss its importance for securing your cluster, and provide examples using a real-world scenario.

### What is Role-Based Access Control (RBAC)?

RBAC is a method of managing permissions by assigning roles to users, groups, or service accounts. In OpenShift, roles are sets of rules that define the actions (verbs) allowed on specific resources (API objects). By granting roles to users or groups, you control their access to cluster resources based on the principle of least privilege, ensuring that users have only the necessary permissions to perform their tasks.

### RBAC in Red Hat OpenShift

RBAC is a security mechanism that grants permissions to users based on their roles within the organization. It provides a way to control what resources users can access and what actions they can perform within a cluster.

In OpenShift, RBAC is implemented through a set of built-in roles and custom roles, which can be assigned to users, groups, and service accounts. The key components of RBAC in OpenShift are:

    Roles: A role is a collection of policies that define a set of permissions, including what actions users can perform on resources.
    ClusterRoles: Similar to roles, but these permissions are applicable cluster-wide, rather than in a specific namespace.
    RoleBindings: These are objects that associate roles with users, groups, or service accounts, granting them the permissions defined by the role

To set up a new RBAC for a user to deploy and manage an application in a new project called "my-first-app", follow the steps below:

Step 1: Create the project
First, we need to create the "my-first-app" project. Run the following command:

sh

$ oc new-project my-first-app

Step 2: Create a custom role
Create a new custom role that grants the necessary permissions for deploying and managing applications. In this example, we will create a role called "app-manager". Save the following YAML content to a file named "app-manager-role.yaml":

yaml

kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: app-manager
  namespace: my-first-app
rules:
- apiGroups: [""]
  resources: ["pods", "services", "endpoints", "persistentvolumeclaims", "configmaps", "secrets"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
- apiGroups: ["apps"]
  resources: ["deployments", "replicasets", "statefulsets"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]

Create the custom role by applying the YAML file:

sh

$ oc apply -f app-manager-role.yaml

Step 3: Create a RoleBinding
Now, we need to bind the custom role to the user. In this example, let's assume the user's username is "johndoe". Create a RoleBinding by saving the following YAML content to a file named "app-manager-rolebinding.yaml":

yaml

kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: app-manager-binding
  namespace: my-first-app
subjects:
- kind: User
  name: johndoe
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: app-manager
  apiGroup: rbac.authorization.k8s.io

Create the RoleBinding by applying the YAML file:

sh

$ oc apply -f app-manager-rolebinding.yaml

With these steps, we have successfully set up a brand new RBAC for the user "johndoe" to deploy and manage an application in the "my-first-app" project. The user can now interact with the cluster, creating deployments, managing services, and configuring resources within the scope of the "app-manager" role in the "my-first-app" namespace.

## Resources
- 
- Red Hat OpenShift Documentation - [Using RBAC to define and apply permissions](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.12/html/authentication_and_authorization/using-rbac)