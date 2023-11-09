# Red Hat OpenShift - Architecture

## Base Operating System - Red Hat CoreOS

[In 2018 Red Hat acquired](https://www.redhat.com/en/about/press-releases/red-hat-acquire-coreos-expanding-its-kubernetes-and-containers-leadership) a small startup called CoreOS, which designed and produced the cloud-native focused operating system "CoreOS".

This product was folding into the portfolio as "Red Hat CoreOS" (RHCOS), to become their primary operating system for Red Hat OpenShift. In fact, RHCOS is only supported for use as part of a Red Hat OpenShift environment. Meaning you cannot just spin an instance to use as your general purpose linux. This is different to competitors such as VMware, whom are the guardians of [PhotonOS](https://vmware.github.io/photon/assets/files/html/3.0/Introduction.html), which is the base OS for their Tanzu Kubernetes Grid platform, as well as many other of their software appliances, and available for general use by the community as an open-source project.

The most important note about RHCOS, is that it is the only supported operating system for the Red Hat OpenShift Control Plane (a.k.a Master) nodes. For Compute Plane (a.k.a Worker) nodes, you have the choice to deploy either RHCOS or[ Red Hat Enterprise Linux (RHEL)](https://www.redhat.com/en/technologies/linux-platforms/enterprise-linux) as the operating system. You can read about the RHCOS [key features here](https://docs.openshift.com/container-platform/4.12/architecture/architecture-rhcos.html#rhcos-key-features_architecture-rhcos).

RHCOS is designed as a minimal user-interaction platform from a configuration standpoint. It is not [encouraged to directly configure](https://docs.openshift.com/container-platform/4.12/architecture/architecture-rhcos.html#rhcos-configured_architecture-rhcos) a RHCOS instance, as it's management comes from the Red Hat OpenShift platform itself, meaning any configuarations would be controlled via Kubernetes Objects. 

Bringing a RHCOS machine online is enabled through [Ignition](https://docs.openshift.com/container-platform/4.12/architecture/architecture-rhcos.html#rhcos-about-ignition_architecture-rhcos), a utility designed to manipulate the disks during initial configuration of the machine. This runs at first boot, and typically will be used to configure disk partitions for the machine, and provide the necessary details for the machine to connect ot the bootstrapping machine in the environment, from which it will recieve it's configuration to become part of the Red Hat OpenShift Cluster.

- When booting an RHCOS machine, you provide an ignition file, which contains the various components/settings for first boot. This is a JSON file which is base64 encoded. Typically for a Cluster Node, this will be details on how to connect to the bootstrap machine. For the bootstrap machine, when building a new Red Hat OpenShift cluster, this ignition file will contain details such as secrets/locations to pull the necessary Red Hat software and configuration manifests for the Red Hat OpenShift Cluster.
- Ignition files are typically created from the install-config.yaml file used with the OpenShift-Install CLI tool.

To summarize this section, RHCOS brings the following, as a specifically designed cloud native operating system:

- Ignition, for first boot configuration of machines.
- CRI-O, a Kubernetes native container runtime implementation. (OpenShift 3.x (Deprecated version) installations used Docker).
- Kubelet, a kubernetes node agent for launching and monitoring containers.

## Product Architecture

At a basic level, Red Hat OpenShift if built ontop of the open-source platform, Kubernetes. Meaning that all components you have learned about from this base platform, are apparent and available in the Red Hat OpenShift platform.

If you haven't visited the [#90DaysOfDevOps - Kubernetes section](https://github.com/MichaelCade/90DaysOfDevOps/blob/main/2022.md#kubernetes), then I urge you to do so, before continuing with this section on Red Hat OpenShift. 

![Red Hat OpenShift - Product Architect](images/Day57%20-%20Red%20Hat%20OpenShift%20Architecture/Red%20Hat%20OpenShift%20-%20Product%20Architecture.png)

Ontop of the Kubernetes platform, Red Hat then delivers it's enterprise sauce sprinkled around to help make your cloud native strategy a sucess:

- Ability to deploy to variety of clouds or your on-premises datacenter.
- Integrating Red Hat Technologies from Red Hat Enterprise Linux
- Open Source development mode. This means source code is available in public software repositories.

Red Hat believes that although Kubernetes is a great platform for managing your applicaitons, it doesn't do a great job of platform-level requirements management (think supporting services to make your apps work) or deployment process handling. Therefore they layer ontop they additional components to give you a full enterprise ready Kubernetes platform. 

- Custom Operating system based on Red Hat Enterprise Linux (RHCOS)(See above).
- Simplified installation and lifecycle management at a cluster platform level (See below).
- Operators - serving as platform foundation, removing the need to upgrade operating systems and control-plane applications, leading to a simplier lifecycle management. This is based on [OperatorHub](https://operatorhub.io/). (See below for an explanation).
- Red Hat Quay Container Image Registry - based on [Quay.io](quay.io)
- Other Enhancements to the base Kubernetes platform include:
  - Software Defined Networking
  - Authentication
  - Log aggregation
  - Montioring
  - Routing

And finally to round off, you can interact with a Red Hat OpenShift Cluster, either via a "Comprehensive" web console, or the custom [OpenShift CLI tool ```oc```](https://docs.openshift.com/container-platform/4.12/cli_reference/openshift_cli/getting-started-cli.html), which is a mix of ```kubectl```, ```kubeadm``` and some specific CLI for Red Hat OpenShift.

The below image nicely finishes off this section covering the product and it's components and why you would potentially choose Red Hat OpenShift over a vanilla Kubernetes platform.

- Simplification of creating and managing a cluster
- Built in tooling for the Application Developer to create and deploy their applications, with workload lifecycle management included, such as ability to monitor and scale those applications.

![Red Hat OpenShift Container Platform Lifecycle](images/Day57%20-%20Red%20Hat%20OpenShift%20Architecture/OpenShift%20Container%20Platform%20lifecycle.png)

For a further deep dive into the control plane architecture, you can read the [official documentation here](https://docs.openshift.com/container-platform/4.12/architecture/control-plane.html).

### What is an Operator?

A little side note here, to explain what an Operator is in regards to a Kubernetes platform. It was created by CoreOS, and hence it's significance in the Red Hat OpenShift Platform.

> Operators implement and automate common Day-1 (installation, configuration, etc) and Day-2 (re-configuration, update, backup, failover, restore, etc.) activities in a piece of software running inside your Kubernetes cluster, by integrating natively with Kubernetes concepts and APIs. We call this a Kubernetes-native application. With Operators you can stop treating an application as a collection of primitives like Pods, Deployments, Services or ConfigMaps, but instead as a single object that only exposes the knobs that make sense for the application.

[Source](https://operatorhub.io/what-is-an-operator)

# Installation methods for Red Hat OpenShift

Red Hat OpenShift Container Platform (OCP) offers the flexibility of:

- Deploy a cluster on provisioned infrastructure and the cluster it maintains.
- Deploy a cluster on infrastructure that you prepare and maintain.

And with these options, there are two types of installation methods/deployment methods for clusters:

- Installer provisioned infrastructure (IPI)
- User provisioned infrastructure (UPI)

There is a third method, which is Agent-based, providing the flexbility of UPI, driven by the Assisted Installer (AI) tool. 

Either method, IPI or UPI is driven from the ```openshift-install``` installation program, which is a CLI tool provided for Linux and Mac Operating systems only. 

The installation program will generate the necessary components to build a cluster such as the Ignition files for bootstrap, master and worker machines. It will further monitor the installation for known targets that an installation must achieve for a successful deployment of a cluster, and provide error handling in the event of a failed cluster deployment, by collecting the necessary troubleshooting logs. 

To visualise bringing all these moving parts together, I have provided the below image from the Red Hat OpenShift documentation.

A cluster definition is created in a special file called ```install-config.yaml```, this file contains the following information:

- Cluster name
- Base domain (FDQN for the network where the cluster will run)
- Details to pull the software from an image registry (Pull Secret)
- SSH Key (so you can access the nodes for troubleshooting)
- Specific Infrastructure platform details (Login details, which networks and storage to use, for example)
- Workload Customizations, such what instance types to use for your Control Plane (Master) and Compute Plane (Worker) nodes.

There is also additional files which may be stored along side the root of the ```install-config.yaml``` in a folder called ```manifests``` these are additional files which can be configured to assist the bootstrapping of a cluster to integrate with your infrastructure, such as your Networking platform.

Once you have all of these files, by running the ```openshift-install``` CLI tool, this will create the ignition files for your boostrap, control plane, and compute plane nodes. Returning to the earlier descriptions of RHCOS, these files contain the first boot information to configure the Operation System and start the process of building a consistent Kubernetes cluster with minimal to no interaction.

![OpenShift Container Platform installation targets and dependencies](images/Day57%20-%20Red%20Hat%20OpenShift%20Architecture/OpenShift%20Container%20Platform%20installation%20targets%20and%20dependencies.png)

## Installer provisioned infrastructure (IPI)

This is the default installation method, and preferred by Red Hat for their customers to initiate a cluster installation, as it provides a reference architectural deployment out of the box.

The ```openshift-install``` CLI tool can act as it's own installation wizard, presenting you with a number of queries for the values it needs to deploy to your choosen platform. You can also customize the installation process to support more advanced scenerios, such as the number of machines deployed, instance type/size, CIDR range for the Kubernetes service network, 

The main point here, is that the installation software provisions the underlying infrastructure for the cluster.

By using an IPI installation method, the provisioned cluster then has the further ability to continue to manage all aspects of the cluster and provisioned infrastructure going forward from a lifecycle management point of view. For example, if you scale the number of compute plane (worker) nodes in your cluster, the OpenShift Container Platform can interact with the underlying platform (for example, AWS, VMware vSphere) to create the new virtual machines and bootstrap them to the cluster.


## User provisioned infrastructure (UPI)

With a UPI method, the OpenShift Container Platform will be installed to infrastucture that you have provided. The installation software will still be used to generate the assets needed to provision the cluster, however you will manually build the nodes and provide the necessary ignition to bring the nodes online. You must also manage the infrastructure supporting cluster resources such as:

- Any infrastructure the nodes are deployed to/upon (Azure Stack Hub, vSphere, IBM Power, Bare Metal, for example)
- Load Balancers
- Cluster networking, including DNS records and required subnets
- Storage for both cluster infrastructure and workloads

When using a UPI installation, you have the option to deploy your Compute Plane (worker) nodes as Red Hat Enterprise Linux machines.

## Assisted Installer

As mentioned earlier, the assisted installer is a kind of hybrid of the UPI method, but offers the hosting of the installation artifacts and removing the need for a bootstrap machine, essentially you provision/install your nodes from a live boot cd, which has the necessary configuration to bring up your node and pull down the rest of the hosted files from a known location. 

You can find out more from this [Red Hat blog post](How to use the OpenShift Assisted Installer), or [official documentation](https://docs.openshift.com/container-platform/4.12/installing/installing_on_prem_assisted/installing-on-prem-assisted.html)

# Installation Process

A temporary bootstrap machine is provisioned using IPI or UPI, which contains the necessary information to build the OpenShift cluster itself (which becomes the permanent control plane nodes). Once the control plane is online, the control plane will initiate the creation of the compute plane (worker) nodes.

![Creating the bootstrap, control plane, and compute machines](images/Day57%20-%20Red%20Hat%20OpenShift%20Architecture/Creating%20the%20bootstrap%20control%20plane%20and%20compute%20machines.png)

Once the control plane is initialised, the bootstrap machine is destroyed. If you are manually provisioning the platform (UPI), then you complete a number of the provisioning steps manually.

>Bootstrapping a cluster involves the following steps:
>    1. The bootstrap machine boots and starts hosting the remote resources required for the control plane machines to boot. (Requires manual intervention if you provision the infrastructure)
>    2. The bootstrap machine starts a single-node etcd cluster and a temporary Kubernetes control plane.
>    3. The control plane machines fetch the remote resources from the bootstrap machine and finish booting. (Requires manual intervention if you provision the infrastructure)
>    4. The temporary control plane schedules the production control plane to the production control plane machines.
>    5. The Cluster Version Operator (CVO) comes online and installs the etcd Operator. The etcd Operator scales up etcd on all control plane nodes.
>    6. The temporary control plane shuts down and passes control to the production control plane.
>    7. The bootstrap machine injects OpenShift Container Platform components into the production control plane.
>    8. The installation program shuts down the bootstrap machine. (Requires manual intervention if you provision the infrastructure)
>    9. The control plane sets up the compute nodes.
>    10. The control plane installs additional services in the form of a set of Operators.
>
>The result of this bootstrapping process is a running OpenShift Container Platform cluster. The cluster then downloads and configures remaining components needed for the day-to-day operation, including the creation of compute machines in supported environments.

# Summary

We have covered the components that make up a Red Hat OpenShift Container Platform environment, why they are important to the environment, and what enteprise features they bring over a vanilla Kubernetes environment. We then dived into the methods available to deploy an OpenShift Cluster and the process that a Cluster build undertakes. 

In [Day 58](../day58.md) will cover the steps to install Red Hat OpenShift to a VMware vSphere environment.

# Resources

- [Glossary of common terms for OpenShift Container Platform architecture](https://docs.openshift.com/container-platform/4.12/architecture/index.html#openshift-architecture-common-terms_architecture-overview)
