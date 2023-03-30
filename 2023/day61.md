# Understanding Authentication, Role-Based Access Control (RBAC) and Auditing in Red Hat OpenShift: Control and Secure Your Cluster

In this post, we are going to look at expanding access to our clusters to other users, we've already tackled the more complex subject of Security Contexts, which looks at the permissions that workloads have when they run on the platform, now we turn this focus to the users of the platform.

## Back to Basics: Authentication, Authorization, Accounting

AAA stands for Authentication, Authorization, and Accounting. These are three essential components of IT systems' security and access control. AAA ensures that users are who they claim to be, have the right permissions to access resources, and keep track of their activities within the system. Let's break down each component:

- **Authentication**: This is the process of validating a user's identity when they attempt to access an IT system. Authentication typically involves requesting a user to provide their credentials, such as a username and password, or using other mechanisms like multi-factor authentication (MFA) and single sign-on (SSO).

- **Authorization**: After a user has been authenticated, authorization determines the level of access they have within the system. It controls which resources a user can access and what actions they can perform. This is usually managed by defining roles or permissions and assigning them to users or groups.

- **Accounting**: Also referred to as auditing, is the process of recording and monitoring user activities within an IT system. This includes logging information like when a user logs in, what actions they perform, and what resources they access. Accounting is essential for security, compliance, and troubleshooting purposes.

Together, these three components form the AAA framework that helps IT administrators manage and secure access to their systems, ensuring that only authorized users can access resources and that their actions are logged for auditing purposes.
## Back to Red Hat Openshift: Introduction to Authentication, Access Control and Auditing

Authentication in Red Hat OpenShift is the process of validating the identity of a user or system attempting to access the cluster. OpenShift 4 supports various identity providers (IdPs) for authentication, such as LDAP, GitHub, GitLab, Google, Keystone, and many more. The authentication process in OpenShift 4 is handled by the OAuth server, which is responsible for managing OAuth tokens and interacting with the configured identity providers.

Role-Based Access Control (RBAC) (Authorization) is an essential security feature in Red Hat OpenShift that enables administrators to manage and control access to resources within a cluster. In this blog post, we will explore RBAC in OpenShift, discuss its importance for securing your cluster, and provide examples using a real-world scenario.

Red Hat uses the Kubernetes audit logging mechanism to perform accounting. The audit logs record requests made to the Kubernetes API server, providing detailed information about the user, resource, action, and outcome. The audit logs are essential for understanding user behavior, detecting security incidents, and meeting compliance requirements.

## Configuring an Identity provider for Authorization

Before we dig into any RBAC configurations, we need to start by giving our OpenShift cluster some form of identity provider that allows our users to login. In my previous posts, I was simply using the ```kubeadmin``` god mode account that is generated at bootstrap.

To specify an identity provider, will need to create a custom resource (CR) that describes that identity provider and add it to the cluster. For the example of this post, I will detail how to configure a ldap identity provider to validate user names and passwords against an LDAPv3 server, using simple bind authentication. In my lab environment, I already have a Windows Server which is set up as a [domain controller](https://learn.microsoft.com/en-us/windows-server/identity/ad-ds/get-started/virtual-dc/active-directory-domain-services-overview).

In short, when a user logs in to the OpenShift cluster, OpenShift will contact the domain controller via a configured provided domain account, and perform a lookup of the users account, if it is returned, it will try to bind to the LDAP provider with the users account and password provided. If this bind is successful, the user will be authenticated to the OpenShift cluster.

1. On your domain controller, create a user account for the OpenShift Cluster to use for lookups, and ensure this account has the appropriate permissions. For Windows Server domain services, this user account can be a member of the standard ```domain users``` group. It is also a good time to record and have to hand the following information ready;

- LDAP server URL (e.g., ldap://ldap.example.com:389)
- Bind DN and password for the LDAP server (e.g., cn=admin,dc=example,dc=com and password)
  - This password will be [base64](https://en.wikipedia.org/wiki/Base64) encoded in the secret.
- User search base and filter (e.g., ou=users,dc=example,dc=com and (uid=%u))
- Group search base and filter (e.g., ou=groups,dc=example,dc=com and (member=%u))

2. Create a secret that contains the ```bindPassword```, i.e the password of the domain account used to connect to the LDAP server for the user lookups.

````sh
$ oc create secret generic ldap-secret --from-literal=bindPassword=<secret> -n openshift-config 

# my example
oc create secret generic ldap-secret --from-literal=bindPassword=VMware1! -n openshift-config
````

3. Now we apply the configuration that tells the OAuth service about our LDAP server and how to connect to it. I have not performed the extra steps to save the CA cert from the LDAP Server, as I am using the insecure LDAP port 389.

Save the below file as ```ldap-provider.yaml```, change for your user details, and apply to your cluster using ```oc apply -f ldap-provider.yaml```.

````yaml
apiVersion: config.openshift.io/v1
kind: OAuth
metadata:
  name: cluster
spec:
  identityProviders:
  - name: ldapidp 
    mappingMethod: claim 
    type: LDAP
    ldap:
      attributes:
        id: 
        - name
        name: 
        - cn
        preferredUsername: 
        - sAMAccountName
      bindDN: "CN=svc_openshift,OU=Services,OU=Accounts,DC=simon,DC=local" 
      bindPassword: 
        name: ldap-secret
      insecure: true
      url: "ldap://sc-step-01.simon.local:389/CN=Users,DC=simon,DC=local?sAMAccountName"
````

Some notes about the above configuration;
- attributes - A first non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. Defined attributes are retrieved as raw, allowing for binary values to be used.
  - Ensure the values provided appear as attributes in your domain controller, this is usually a cause of failures.
- bindDN - DN to use to bind during the search phase. Must be set if bindPassword is defined.
- insecure - When true, no TLS connection is made to the server. When false, ldaps:// URLs connect using TLS, and ldap:// URLs are upgraded to TLS.
- url - ensure you add the port, and that your search OU/CN path is correct, as well as the [search filter attribute](https://docs.openshift.com/container-platform/4.12/authentication/identity_providers/configuring-ldap-identity-provider.html#identity-provider-about-ldap_configuring-ldap-identity-provider), in this example it is ```sAMAccountName```

> If a CR does not exist, oc apply creates a new CR and might trigger the following warning: Warning: oc apply should be used on resources created by either oc create --save-config or oc apply. In this case you can safely ignore this warning.


4. Now you can login to the OpenShift cluster as an LDAP user either via the ```oc login``` CLI command, or via the Console UI. You can logout with ```oc logout```

![oc login](/2023/images/Day61%20-%20Authentication%20-%20Role-Based%20Access%20Control%20and%20Auditing%20in%20Red%20Hat%20OpenShift/oc%20login.jpg)

Final note on this subject which caught me out. Once you log out, you will remain in the ```oc context``` of that user you've logged in, and if you view your ```KUBECONFIG``` file, you will see that user's context has now been added. So for me, I wanted to get back to using my kubeadmin account, but I had to run ```oc config use-context admin``` command. 

## What is Role-Based Access Control (RBAC)?

RBAC is a method of managing permissions by assigning roles to users, groups, or service accounts. In OpenShift, roles are sets of rules that define the actions (verbs) allowed on specific resources (API objects). By granting roles to users or groups, you control their access to cluster resources based on the principle of least privilege, ensuring that users have only the necessary permissions to perform their tasks. Here, Red Hat OpenShift is just consuming the native features of Kubernetes, concerning RBAC. However, it's important to understand this area, as it then features in the enterprise features OpenShift brings, such as Projects and pipelines to name but a few.

### RBAC in Red Hat OpenShift

In OpenShift, RBAC is implemented through a set of built-in roles and custom roles, which can be assigned to users, groups, and service accounts. The key components of RBAC in OpenShift are:

- Roles: A role is a collection of policies that define a set of permissions, including what actions users can perform on resources.
- ClusterRoles: Similar to roles, but these permissions are applicable cluster-wide, rather than in a specific namespace.
- RoleBindings: These are objects that associate roles with users, groups, or service accounts, granting them the permissions defined by the role

To set up a new RBAC for a user to deploy and manage an application in a new project called "rbac-test-ns", follow the steps below:

1. First, we need to create the "rbac-test-ns" project. Run the following command:

````sh
$ oc new-project rbac-test-ns
````

2. Create a new custom role that grants the necessary permissions for deploying and managing applications. In this example, we will create a role called "app-mgr". Save the following YAML content to a file named "app-mgr-role.yaml":

````yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: app-mgr
  namespace: rbac-test-ns
rules:
- apiGroups: [""]
  resources: ["pods", "services", "endpoints", "persistentvolumeclaims", "configmaps", "secrets"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
- apiGroups: ["apps"]
  resources: ["deployments", "replicasets", "statefulsets"]
  verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
````
Create the custom role by applying the YAML file:

````sh

$ oc apply -f app-mgr-role.yaml
````
Step 3: Create a RoleBinding

Now, we need to bind the custom role to the user. In this example, let's assume the user's username is "johndoe". Create a RoleBinding by saving the following YAML content to a file named "app-mgr-rolebinding.yaml":

````yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: app-mgr-binding
  namespace: rbac-test-ns
subjects:
- kind: User
  name: johndoe
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: app-mgr
  apiGroup: rbac.authorization.k8s.io
````
Create the RoleBinding by applying the YAML file:

````sh
$ oc apply -f app-mgr-rolebinding.yaml
````
With these steps, we have successfully set up a brand new RBAC for the user "johndoe" to deploy and manage an application in the "rbac-test-ns" project. The user can now interact with the cluster, creating deployments, managing services, and configuring resources within the scope of the "app-mgr" role in the "rbac-test-ns" namespace.

## How Accounting Works in Red Hat OpenShift 4

Here are the main components of accounting in OpenShift 4:

    Audit Policy: The audit policy determines which requests should be logged and the level of detail to include in the logs. You can configure the audit policy using a YAML file that specifies rules for each type of resource and action.

    Audit Backend: The audit backend is responsible for processing and storing the audit logs. OpenShift 4 supports two types of audit backends: log backend and webhook backend. The log backend writes logs to a file on the API server node, while the webhook backend sends logs to an external HTTP(S) endpoint.

    Log Retention and Rotation: OpenShift 4 provides mechanisms for managing audit log retention and rotation to ensure that logs do not consume excessive disk space. Log rotation settings can be configured to control the maximum log file size and the number of old log files to keep.

To enable and configure audit logging in OpenShift 4, administrators can create an audit policy, configure the audit backend, and set log retention and rotation settings. Once configured, the audit logs can be used to monitor user activities, identify security issues, and comply with regulatory requirements.
## Resources
- Kubernetes.io - [Using RBAC Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
- Red Hat OpenShift Documentation 
  - [Configuring an LDAP identity provider](https://docs.openshift.com/container-platform/4.12/authentication/identity_providers/configuring-ldap-identity-provider.html)
  - [Using RBAC to define and apply permissions](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.12/html/authentication_and_authorization/using-rbac)