# Understanding Authentication, Role-Based Access Control (RBAC) and Auditing in Red Hat OpenShift: Control and Secure Your Cluster

In this post, we are going to look at expanding access to our clusters to other users, we've already tackled the more complex subject of Security Contexts, which looks at the permissions that workloads have when they run on the platform, now we turn this focus to the users of the platform.

## Back to Basics: Authentication, Authorization, Accounting

AAA stands for Authentication, Authorization, and Accounting. These are three essential components of IT systems' security and access control. AAA ensures that users are who they claim to be, have the right permissions to access resources, and keep track of their activities within the system. Let's break down each component:

- **Authentication**: This is the process of validating a user's identity when they attempt to access an IT system. Authentication typically involves requesting a user to provide their credentials, such as a username and password, or using other mechanisms like multi-factor authentication (MFA) and single sign-on (SSO).

- **Authorization**: After a user has been authenticated, authorization determines the level of access they have within the system. It controls which resources a user can access and what actions they can perform. This is usually managed by defining roles or permissions and assigning them to users or groups.

- **Accounting**: Also referred to as auditing, is the process of recording and monitoring user activities within an IT system. This includes logging information like when a user logs in, what actions they perform, and what resources they access. Accounting is essential for security, compliance, and troubleshooting purposes.

Together, these three components form the AAA framework that helps IT administrators manage and secure access to their systems, ensuring that only authorized users can access resources and that their actions are logged for auditing purposes.
## Back to Red Hat Openshift: Introduction to Authentication, Access Control and Auditing

Authentication in Red Hat OpenShift is the process of validating the identity of a user or system attempting to access the cluster. OpenShift supports various identity providers (IdPs) for authentication, such as LDAP, GitHub, GitLab, Google, Keystone, and many more. The authentication process in OpenShift 4 is handled by the OAuth server, which is responsible for managing OAuth tokens and interacting with the configured identity providers.

Role-Based Access Control (RBAC) (Authorization) is an essential security feature in Red Hat OpenShift that enables administrators to manage and control access to resources within a cluster.

Red Hat OpenShift uses the Kubernetes audit logging mechanism to perform accounting. The audit logs record requests made to the Kubernetes API server, providing detailed information about the user, resource, action, and outcome. The audit logs are essential for understanding user behavior, detecting security incidents, and meeting compliance requirements.

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
- attributes - The first non-empty attribute is used. At least one attribute is required. If none of the listed attribute have a value, authentication fails. Defined attributes are retrieved as raw, allowing for binary values to be used.
  - Ensure the values provided appear as attributes in your domain controller, this is usually a cause of failures.
- bindDN - DN to use to bind during the search phase. Must be set if bindPassword is defined.
- insecure - When true, no TLS connection is made to the server. When false, ldaps:// URLs connect using TLS, and ldap:// URLs are upgraded to TLS.
- url - ensure you add the port, and that your search OU/CN path is correct, as well as the [search filter attribute](https://docs.openshift.com/container-platform/4.12/authentication/identity_providers/configuring-ldap-identity-provider.html#identity-provider-about-ldap_configuring-ldap-identity-provider), in this example it is ```sAMAccountName```

> If a CR does not exist, oc apply creates a new CR and might trigger the following warning: Warning: oc apply should be used on resources created by either oc create --save-config or oc apply. In this case you can safely ignore this warning.


4. Now you can login to the OpenShift cluster as an LDAP user either via the ```oc login``` CLI command, or via the Console UI. You can logout with ```oc logout```

![oc login](/2023/images/Day61%20-%20Authentication%20-%20Role-Based%20Access%20Control%20and%20Auditing%20in%20Red%20Hat%20OpenShift/oc%20login.jpg)

Final note on this subject which caught me out. Once you log out, you will remain in the ```oc context``` of that user you've logged in, and if you view your ```KUBECONFIG``` file, you will see that user's context has now been added. So for me, I wanted to get back to using my kubeadmin account, but I had to run ```oc config use-context admin``` command.

Now we have a way for our users to log into the platform, we can start to look at controlling what they can do on the platform.

## What is Role-Based Access Control (RBAC)?

RBAC is a method of managing permissions by assigning roles to users, groups, or service accounts. In OpenShift, roles are sets of rules that define the actions (verbs) allowed on specific resources (API objects). By granting roles to users or groups, you control their access to cluster resources based on the principle of least privilege, ensuring that users have only the necessary permissions to perform their tasks. Here, Red Hat OpenShift is just consuming the native features of Kubernetes, concerning RBAC. However, it's important to understand this area, as it then features in the enterprise features OpenShift brings, such as Projects and pipelines to name but a few.

## RBAC in Red Hat OpenShift

In OpenShift, RBAC is implemented through a set of built-in roles and custom roles, which can be assigned to users, groups, and service accounts. The key components of RBAC in OpenShift are:

- Roles: A role is a collection of policies that define a set of permissions, including what actions users can perform on resources. These are applied to a namespace.
- ClusterRoles: Similar to roles, but these permissions are applicable cluster-wide, rather than in a specific namespace.
- RoleBindings: These are objects that associate roles with users, groups, or service accounts, granting them the permissions defined by the role

The below image shows the relationships between the different components.

![Red Hat OpenShift - Relationships between cluster roles, local roles, cluster role bindings, local role bindings, users, groups and service accounts](/2023/images/Day61%20-%20Authentication%20-%20Role-Based%20Access%20Control%20and%20Auditing%20in%20Red%20Hat%20OpenShift/Red%20Hat%20OpenShift%20-%20Access%20Control%20-%20Relationship%20between%20cluster%20roles%2C%20local%20roles%20and%20role%20bindings.png)

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
- apiGroups: ["project.openshift.io"]
  resources: ["projects"]
  verbs: ["get"]
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

3. Now, we need to bind the custom role to the user. In this example, let's assume the user's username is "johndoe". Create a RoleBinding by saving the following YAML content to a file named "app-mgr-rolebinding.yaml":

````yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: app-mgr-binding
  namespace: rbac-test-ns
subjects:
- kind: User
  name: test
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: app-mgr
  apiGroup: rbac.authorization.k8s.io
````

Create the RoleBinding by applying the YAML file:
As a alternative to add a user to an existing role, we can use the following command line:

````sh
$ oc adm policy add-role-to-user <role> <user> -n <project>

# For my example above, once I created the "app-mgr" role in the rbac-test-ns namespace I can run the following to create the roleBinding
oc adm policy add-role-to-user app-mgr test -n rbac-test-ns
````
We can also see who can perform actions against certain resources too by using the command:

````sh
oc adm policy who-can {verb} {resource}

# for example
oc adm policy who-can create pods -n rbac-test-ns

# below is the output (shortened) you can see my test user at the bottom
resourceaccessreviewresponse.authorization.openshift.io/<unknown> 

Namespace: rbac-test-ns
Verb:      create
Resource:  pods

Users:  system:admin
        system:serviceaccount:kube-system:daemon-set-controller
        system:serviceaccount:kube-system:job-controller
        .......
        system:serviceaccount:rbac-test-ns:deployer
        test
Groups: system:cluster-admins
        system:masters

````

With these steps, we have successfully set up a brand new RBAC for the user "test" (from my LDAP setup and testing earlier) to deploy and manage an application in the "rbac-test-ns" project. The user can now interact with the cluster, creating deployments, managing services, and configuring resources within the scope of the "app-mgr" role in the "rbac-test-ns" namespace.

## How Accounting Works in Red Hat OpenShift 4

Here are the main components of accounting in Red Hat OpenShift:

- Audit Policy: The audit policy determines which requests should be logged and the level of detail to include in the logs. You can configure the audit policy using a YAML file that specifies rules for each type of resource and action.

- Audit Backend: The audit backend is responsible for processing and storing the audit logs. Red Hat OpenShift supports two types of audit backends: log backend and webhook backend. The log backend writes logs to a file on the API server node, while the webhook backend sends logs to an external HTTP(S) endpoint.

- Log Retention and Rotation: OpenShift provides mechanisms for managing audit log retention and rotation to ensure that logs do not consume excessive disk space. Log rotation settings can be configured to control the maximum log file size and the number of old log files to keep.

To enable and configure audit logging in OpenShift, administrators can create an audit policy, configure the audit backend, and set log retention and rotation settings. Once configured, the audit logs can be used to monitor user activities, identify security issues, and comply with regulatory requirements.

Auditing is conducted at the API Server level, as it captures all requests coming into the server. This means that the audit logs by default will be stored on each of the control-plane nodes.

Each audit log will contain the following fields:

- level - The audit level at which the event was generated.

- auditID - A unique audit ID, generated for each request.

- stage - The stage of the request handling when this event instance was generated.

- requestURI - The request URI is sent by the client to a server.

- verb - The Kubernetes verb associated with the request. For non-resource requests, this is the lowercase HTTP method.

- user - The authenticated user information.

- impersonatedUser - Optional. The impersonated user information, if the request is impersonating another user.

- sourceIPs - Optional. The source IPs, from where the request originated and any intermediate proxies.

- userAgent - Optional. The user agent string is reported by the client. Note that the user agent is provided by the client, and must not be trusted.

- objectRef - Optional. The object reference this request is targeted at. This does not apply to List-type requests, or non-resource requests.

- responseStatus - Optional. The response status is populated even when the ResponseObject is not a Status type. For successful responses, this will only include the code. For non-status type error responses, this will be auto-populated with the error message.

- requestObject - Optional. The API object from the request is in JSON format. The RequestObject is recorded as is in the request (possibly re-encoded as JSON), prior to version conversion, defaulting, admission or merging. It is an external versioned object type, and might not be a valid object on its own. This is omitted for non-resource requests and is only logged at the request level and higher.

- responseObject - Optional. The API object is returned in the response, in JSON format. The ResponseObject is recorded after conversion to the external type and serialized as JSON. This is omitted for non-resource requests and is only logged at the response level.

- requestReceivedTimestamp - The time that the request reached the API server.

- stageTimestamp - The time that the request reached the current audit stage.

- annotations - Optional. An unstructured key-value map is stored with an audit event that may be set by plugins invoked in the request serving chain, including authentication, authorization and admission plugins. Note that these annotations are for the audit event, and do not correspond to the ```metadata.annotations``` of the submitted object. Keys should uniquely identify the informing component to avoid name collisions, for example, podsecuritypolicy.admission.k8s.io/policy. Values should be short. Annotations are included in the metadata level.

Below is an example of output from the Kubernetes API Server from the official documentation:

````json
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"ad209ce1-fec7-4130-8192-c4cc63f1d8cd","stage":"ResponseComplete","requestURI":"/api/v1/namespaces/openshift-kube-controller-manager/configmaps/cert-recovery-controller-lock?timeout=35s","verb":"update","user":{"username":"system:serviceaccount:openshift-kube-controller-manager:localhost-recovery-client","uid":"dd4997e3-d565-4e37-80f8-7fc122ccd785","groups":["system:serviceaccounts","system:serviceaccounts:openshift-kube-controller-manager","system:authenticated"]},"sourceIPs":["::1"],"userAgent":"cluster-kube-controller-manager-operator/v0.0.0 (linux/amd64) kubernetes/$Format","objectRef":{"resource":"configmaps","namespace":"openshift-kube-controller-manager","name":"cert-recovery-controller-lock","uid":"5c57190b-6993-425d-8101-8337e48c7548","apiVersion":"v1","resourceVersion":"574307"},"responseStatus":{"metadata":{},"code":200},"requestReceivedTimestamp":"2020-04-02T08:27:20.200962Z","stageTimestamp":"2020-04-02T08:27:20.206710Z","annotations":{"authorization.k8s.io/decision":"allow","authorization.k8s.io/reason":"RBAC: allowed by ClusterRoleBinding \"system:openshift:operator:kube-controller-manager-recovery\" of ClusterRole \"cluster-admin\" to ServiceAccount \"localhost-recovery-client/openshift-kube-controller-manager\""}}
````

To view the audit logs:

1. List the available audit logs on the control-plane nodes

````sh
$ oc adm node-logs --role=master --path=openshift-apiserver/

# Example output
90days-ocp-72ptq-master-0 audit-2023-03-18T13-32-30.141.log
90days-ocp-72ptq-master-0 audit-2023-03-20T18-52-05.290.log
90days-ocp-72ptq-master-0 audit-2023-03-22T23-59-32.898.log
90days-ocp-72ptq-master-0 audit-2023-03-25T05-18-17.982.log
90days-ocp-72ptq-master-0 audit-2023-03-27T10-18-07.255.log
90days-ocp-72ptq-master-0 audit-2023-03-29T15-39-12.983.log
90days-ocp-72ptq-master-0 audit.log
90days-ocp-72ptq-master-1 audit-2023-03-18T13-40-29.849.log
90days-ocp-72ptq-master-1 audit-2023-03-20T18-48-37.329.log
90days-ocp-72ptq-master-1 audit-2023-03-23T00-03-57.031.log
90days-ocp-72ptq-master-1 audit-2023-03-25T05-14-28.573.log
90days-ocp-72ptq-master-1 audit-2023-03-27T10-25-27.601.log
90days-ocp-72ptq-master-1 audit-2023-03-29T15-34-14.462.log
90days-ocp-72ptq-master-1 audit.log
90days-ocp-72ptq-master-2 audit-2023-03-18T13-15-35.530.log
90days-ocp-72ptq-master-2 audit-2023-03-20T18-22-40.880.log
90days-ocp-72ptq-master-2 audit-2023-03-22T23-18-33.055.log
90days-ocp-72ptq-master-2 audit-2023-03-25T04-18-56.637.log
90days-ocp-72ptq-master-2 audit-2023-03-27T09-44-50.243.log
90days-ocp-72ptq-master-2 audit-2023-03-29T14-39-01.789.log
90days-ocp-72ptq-master-2 audit.log
````

2. To view a specific audit log:

````sh
$ oc adm node-logs <node_name> --path=openshift-apiserver/<log_name>

# Example command
$ oc adm node-logs 90days-ocp-72ptq-master-2 --path=openshift-apiserver/audit-2023-03-29T14-39-01.789.log
````

Example output:

````json
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"5d08341a-8fa2-4fab-8e4a-4fb1c01d2d6e","stage":"ResponseStarted","requestURI":"/apis/template.openshift.io/v1/namespaces/openshift/templates?allowWatchBookmarks=true\u0026resourceVersion=9537849\u0026timeout=5m44s\u0026timeoutSeconds=344\u0026watch=true","verb":"watch","user":{"username":"system:serviceaccount:openshift-cluster-samples-operator:cluster-samples-operator","groups":["system:serviceaccounts","system:serviceaccounts:openshift-cluster-samples-operator","system:authenticated"],"extra":{"authentication.kubernetes.io/pod-name":["cluster-samples-operator-7f8d575897-s7566"],"authentication.kubernetes.io/pod-uid":["8a96caaf-1de9-4cef-915f-ace02c764e52"]}},"sourceIPs":["192.168.200.181","10.130.0.2"],"userAgent":"cluster-samples-operator/v0.0.0 (linux/amd64) kubernetes/$Format","objectRef":{"resource":"templates","namespace":"openshift","apiGroup":"template.openshift.io","apiVersion":"v1"},"responseStatus":{"metadata":{},"code":200},"requestReceivedTimestamp":"2023-03-27T09:44:50.239076Z","stageTimestamp":"2023-03-27T09:44:50.242747Z","annotations":{"authorization.k8s.io/decision":"allow","authorization.k8s.io/reason":"RBAC: allowed by RoleBinding \"shared-resource-viewers/openshift\" of Role \"shared-resource-viewer\" to Group \"system:authenticated\""}}
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"e65876d9-6fca-422c-93a9-110921648d1b","stage":"ResponseComplete","requestURI":"/metrics","verb":"get","user":{"username":"system:serviceaccount:openshift-monitoring:prometheus-k8s","groups":["system:authenticated"]},"sourceIPs":["10.129.2.13"],"userAgent":"Prometheus/2.39.1","responseStatus":{"metadata":{},"code":200},"requestReceivedTimestamp":"2023-03-27T09:44:51.044096Z","stageTimestamp":"2023-03-27T09:44:51.065007Z","annotations":{"authorization.k8s.io/decision":"allow","authorization.k8s.io/reason":"RBAC: allowed by ClusterRoleBinding \"prometheus-k8s\" of ClusterRole \"prometheus-k8s\" to ServiceAccount \"prometheus-k8s/openshift-monitoring\""}}
````

Next, we have the OpenShift OAuth API Server audit logs, which will capture our user interactions

1. List the relevant logs on the control-plane nodes

````sh
oc adm node-logs --role=master --path=oauth-apiserver/

# Example output
90days-ocp-72ptq-master-0 audit-2023-03-22T20-22-19.424.log
90days-ocp-72ptq-master-0 audit-2023-03-29T08-44-53.926.log
90days-ocp-72ptq-master-0 audit.log
90days-ocp-72ptq-master-1 audit-2023-03-22T20-34-32.796.log
90days-ocp-72ptq-master-1 audit-2023-03-29T08-41-39.406.log
90days-ocp-72ptq-master-1 audit.log
90days-ocp-72ptq-master-2 audit-2023-03-22T20-08-20.762.log
90days-ocp-72ptq-master-2 audit-2023-03-29T08-10-51.942.log
90days-ocp-72ptq-master-2 audit.log
````

2. View a specific OAuth API Server log:

````sh
$ oc adm node-logs <node_name> --path=oauth-apiserver/<log_name>

# Example command
$ oc adm node-logs 90days-ocp-72ptq-master-2 --path=oauth-apiserver/audit-2023-03-29T08-10-51.942.log
`````

Example output

````json
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"6472c130-436a-4454-9bf9-77b6066d8ce2","stage":"ResponseComplete","requestURI":"/apis/oauth.openshift.io/v1/oauthclients/console","verb":"get","user":{"username":"system:serviceaccount:openshift-console-operator:console-operator","groups":["system:serviceaccounts","system:serviceaccounts:openshift-console-operator","system:authenticated"],"extra":{"authentication.kubernetes.io/pod-name":["console-operator-7cc8457b5b-rpz7f"],"authentication.kubernetes.io/pod-uid":["4b73b593-9cde-4ae4-a8f7-00e7f7ee7902"]}},"sourceIPs":["192.168.200.181","10.129.0.2"],"userAgent":"console/v0.0.0 (linux/amd64) kubernetes/$Format","objectRef":{"resource":"oauthclients","name":"console","apiGroup":"oauth.openshift.io","apiVersion":"v1"},"responseStatus":{"metadata":{},"code":200},"requestReceivedTimestamp":"2023-03-22T23:35:00.970020Z","stageTimestamp":"2023-03-22T23:35:00.972703Z","annotations":{"authorization.k8s.io/decision":"allow","authorization.k8s.io/reason":"RBAC: allowed by ClusterRoleBinding \"console-operator\" of ClusterRole \"console-operator\" to ServiceAccount \"console-operator/openshift-console-operator\""}}
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"db3d3280-fcb9-4368-b7db-17b97bcb9826","stage":"ResponseComplete","requestURI":"/apis/oauth.openshift.io/v1/oauthclients/console","verb":"get","user":{"username":"system:serviceaccount:openshift-console-operator:console-operator","groups":["system:serviceaccounts","system:serviceaccounts:openshift-console-operator","system:authenticated"],"extra":{"authentication.kubernetes.io/pod-name":["console-operator-7cc8457b5b-rpz7f"],"authentication.kubernetes.io/pod-uid":["4b73b593-9cde-4ae4-a8f7-00e7f7ee7902"]}},"sourceIPs":["192.168.200.181","10.129.0.2"],"userAgent":"console/v0.0.0 (linux/amd64) kubernetes/$Format","objectRef":{"resource":"oauthclients","name":"console","apiGroup":"oauth.openshift.io","apiVersion":"v1"},"responseStatus":{"metadata":{},"code":200},"requestReceivedTimestamp":"2023-03-22T23:35:00.976516Z","stageTimestamp":"2023-03-22T23:35:00.980651Z","annotations":{"authorization.k8s.io/decision":"allow","authorization.k8s.io/reason":"RBAC: allowed by ClusterRoleBinding \"console-operator\" of ClusterRole \"console-operator\" to ServiceAccount \"console-operator/openshift-console-operator\""}}
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"b5a99a12-637a-4bb2-b733-1f3812ca76fe","stage":"ResponseComplete","requestURI":"/openapi/v2","verb":"get","user":{"username":"system:aggregator","groups":["system:authenticated"]},"sourceIPs":["10.128.0.2"],"responseStatus":{"metadata":{},"code":304},"requestReceivedTimestamp":"2023-03-22T23:35:21.077019Z","stageTimestamp":"2023-03-22T23:35:21.085460Z","annotations":{"authorization.k8s.io/decision":"allow","authorization.k8s.io/reason":"RBAC: allowed by ClusterRoleBinding \"cluster-status-binding\" of ClusterRole \"cluster-status\" to Group \"system:authenticated\""}}
````

And finally the OAuth server audit logs

1. List the available logs on the control-plane nodes:

````sh
oc adm node-logs --role=master --path=oauth-server/

# Example output
90days-ocp-72ptq-master-0 audit.log
90days-ocp-72ptq-master-1 audit.log
90days-ocp-72ptq-master-2 audit.log
````

2. View a specific log:

````sh
$ oc adm node-logs <node_name> --path=oauth-server/<log_name>

# Example command
$ oc adm node-logs 90days-ocp-72ptq-master-2 --path=oauth-server/audit.log
````

Example output:

````json
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"0ae79d4b-a44e-4f16-aadf-a600f42a63d0","stage":"RequestReceived","requestURI":"/","verb":"head","user":{"username":"system:anonymous","groups":["system:unauthenticated"]},"sourceIPs":["10.131.0.2"],"userAgent":"Go-http-client/1.1","requestReceivedTimestamp":"2023-03-30T14:13:04.446550Z","stageTimestamp":"2023-03-30T14:13:04.446550Z"}
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"0ae79d4b-a44e-4f16-aadf-a600f42a63d0","stage":"ResponseComplete","requestURI":"/","verb":"head","user":{"username":"system:anonymous","groups":["system:unauthenticated"]},"sourceIPs":["10.131.0.2"],"userAgent":"Go-http-client/1.1","responseStatus":{"metadata":{},"status":"Failure","message":"forbidden: User \"system:anonymous\" cannot head path \"/\"","reason":"Forbidden","details":{},"code":403},"requestReceivedTimestamp":"2023-03-30T14:13:04.446550Z","stageTimestamp":"2023-03-30T14:13:04.456751Z","annotations":{"authorization.k8s.io/decision":"forbid","authorization.k8s.io/reason":""}}
{"kind":"Event","apiVersion":"audit.k8s.io/v1","level":"Metadata","auditID":"a151df17-578c-415f-a83e-64753a0d16dc","stage":"RequestReceived","requestURI":"/oauth/authorize?client_id=openshift-challenging-client\u0026code_challenge=L56_-VuTMU7qF36WuKxF6mnA2nj_oEEETPsAdQ-w24I\u0026code_challenge_method=S256\u0026redirect_uri=https%3A%2F%2Foauth-openshift.apps.90days-ocp.simon.local%2Foauth%2Ftoken%2Fimplicit\u0026response_type=code","verb":"get","user":{"username":"system:anonymous","groups":["system:unauthenticated"]},"sourceIPs":["10.131.0.2"],"userAgent":"Go-http-client/1.1","requestReceivedTimestamp":"2023-03-30T14:13:04.489643Z","stageTimestamp":"2023-03-30T14:13:04.489643Z"}
````

You can filter the logs using the ```jq``` tool (See the [jq Manual](https://stedolan.github.io/jq/manual/) for detailed information). An example command would be:

````sh
$ oc adm node-logs 90days-ocp-72ptq-master-2 --path=openshift-apiserver/audit.log | jq 'select(.user.username == "test")'
````

Example output:

````json
{
  "kind": "Event",
  "apiVersion": "audit.k8s.io/v1",
  "level": "Metadata",
  "auditID": "97ce5c38-4e0c-447b-8e9f-a2c8c84ebc19",
  "stage": "ResponseComplete",
  "requestURI": "/apis/project.openshift.io/v1/projectrequests",
  "verb": "list",
  "user": {
    "username": "test",
    "groups": [
      "system:authenticated:oauth",
      "system:authenticated"
    ],
    "extra": {
      "scopes.authorization.openshift.io": [
        "user:full"
      ]
    }
  },
  "sourceIPs": [
    "10.129.0.35",
    "10.129.0.2"
  ],
  "objectRef": {
    "resource": "projectrequests",
    "apiGroup": "project.openshift.io",
    "apiVersion": "v1"
  },
  "responseStatus": {
    "metadata": {},
    "status": "Success",
    "code": 200
  },
  "requestReceivedTimestamp": "2023-03-30T14:42:37.237408Z",
  "stageTimestamp": "2023-03-30T14:42:37.250606Z",
  "annotations": {
    "authorization.k8s.io/decision": "allow",
    "authorization.k8s.io/reason": "RBAC: allowed by ClusterRoleBinding \"basic-users\" of ClusterRole \"basic-user\" to Group \"system:authenticated\""
  }
}
{
  "kind": "Event",
  "apiVersion": "audit.k8s.io/v1",
  "level": "Metadata",
  "auditID": "5868d505-1c46-47ed-b0ab-d53d9878d495",
  "stage": "ResponseStarted",
  "requestURI": "/apis/project.openshift.io/v1/projects?cluster=local-cluster&watch=true",
  "verb": "watch",
  "user": {
    "username": "test",
    "groups": [
      "system:authenticated:oauth",
      "system:authenticated"
    ],
    "extra": {
      "scopes.authorization.openshift.io": [
        "user:full"
      ]
    }
  },
  "sourceIPs": [
    "192.168.200.10",
    "10.129.0.35",
    "10.129.0.2"
  ],
  "userAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
  "objectRef": {
    "resource": "projects",
    "apiGroup": "project.openshift.io",
    "apiVersion": "v1"
  },
  "responseStatus": {
    "metadata": {},
    "code": 101
  },
  "requestReceivedTimestamp": "2023-03-30T14:42:40.166647Z",
  "stageTimestamp": "2023-03-30T14:42:40.170013Z",
  "annotations": {
    "authorization.k8s.io/decision": "allow",
    "authorization.k8s.io/reason": "RBAC: allowed by ClusterRoleBinding \"basic-users\" of ClusterRole \"basic-user\" to Group \"system:authenticated\""
  }
}
````

To wrap up this section, I think it's worth highlighting that keeping the logs on the nodes themselves is not best practice, and you will want to offload these to an external logging service. I won't cover that in detail in this post, but I will provide you a [link to a blog post](https://veducate.co.uk/openshift-forward-logs-log-insight-cloud/) I created previously, showing how you can offload the logs to the VMware Aria Operations for Logs product (previously known as vRealize Log Insight)

# Summary

In this post, we've seen more of how Red Hat OpenShift takes the out-of-the-box functionality from upstream Kubernetes and entwines it with enterprise engineering, such as the ability to easily add an Identity provider for Authorization without having to pull in more open-source components. The AAA model is a basic premise when it comes to managing the security stance of your environment, and before you let users consume the Red Hat OpenShift platform, it's key to understand how this works and start putting in the structure upfront around the likes of Roles and Role Bindings for user privileges.

This post only takes you into the beginnings of setting up RBAC, and we haven't discussed a cohesive strategy of how to implement it within your organisation. It's safe to say least privilege is the best, and the outcome you need to aim for. I've included several links below for you to understand RBAC in Kubernetes and OpenShift in more detail, and I'm sure with a quick search on your favourite search engine, you'll also find lots of posts around how best to implement a least privilege model.

To wrap up this section on Red Hat OpenShift with [Day 62](/2023/day62.md) we will cover compliance and vulnerability scanning provided by Red Hat Operators.

## Resources
- Kubernetes.io 
  - [Authenicating](https://kubernetes.io/docs/reference/access-authn-authz/authentication/)
  - [Using RBAC Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac/)
- Red Hat OpenShift Documentation
  - [Configuring an LDAP identity provider](https://docs.openshift.com/container-platform/4.12/authentication/identity_providers/configuring-ldap-identity-provider.html)
  - [Using RBAC to define and apply permissions](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.12/html/authentication_and_authorization/using-rbac)
  - [Viewing Audit logs](https://docs.openshift.com/container-platform/4.12/security/audit-log-view.html)
  - [Configuring the audit log policy](https://docs.openshift.com/container-platform/4.12/security/audit-log-policy-config.html#audit-log-policy-config)
  - [Monitoring cluster events and logs](https://docs.openshift.com/container-platform/4.12/security/container_security/security-monitoring.html)
- Red Hat Blog - [Multiple ways of Authentication on OpenShift Container Platform (OCP), Part 2](https://cloud.redhat.com/blog/multiple-ways-of-authentication-on-openshift-container-platform-ocp-part-2)
- OpenShift Examples - [Active Directory/LDAP](https://examples.openshift.pub/cluster-configuration/authentication/activedirectory-ldap/)
- vEducate.co.uk - [How to configure Red Hat OpenShift to forward logs to VMware vRealize Log Insight Cloud](https://veducate.co.uk/openshift-forward-logs-log-insight-cloud/)