# Day 27: 90DaysofDevOps

## From Automated to Automatic - Event-Driven Infrastructure Management with Ansible 

**Daniel Bodky**
- [Twitter](https://twitter.com/d_bodky)
- [LinkedIn](https://linkedin.com/in/daniel-bodky)
- [Blog](https://dbodky.me)

## Overview

A universal truth and recurring theme in the DevOps world is automation. From providing infrastructure to testing code to deploying to production, many parts of the DevOps lifecycle get automated already. One popular technology for managing infrastructure and configuration in an automated way is Ansible, but are we fully utilizing its capabilities yet?

This presentation will give a broad overview of Ansible and its architecture and use-cases, before exploring a relatively new feature, Event-driven Ansible (EDA). Analzying applications of event-driven Ansible, participants will see that automated management is nice, but automatic management is awesome, not just regarding DevOps principles, but also in terms of reaction times, the human tendency for minor mistakes, and toil for operators.

Participants will get first-hand insights into Ansible, its strengths, weaknesses, and the potential of event-driven automation within the DevOps world.

## Demos

<details>

<summary>Prerequisites</summary>

### Ansible Inventory

> [!NOTE]
> For this inventory file to work, you need to create VMs accordingly and adjust the IP addresses to fit your lab environment.

Ansible utilizes so-called inventories to manage a list of hosts and groups of hosts. Below is the inventory for the demo environment used in this presentation.

```yaml
hosts:
  webservers:
    hosts:
      webshop.example.com:  # Ubuntu
        ansible_host: 192.168.1.10
        webserver: apache2
      company.example.com:  # Ubuntu
        ansible_host: 192.168.1.11
        webserver: nginx
      internal.example.com:  # CentOS Stream
        ansible_host: 192.168.1.12
        webserver: httpd
```

You can copy-paste this inventory into a file called `hosts.yml` and use it for the following demos.

</details>

<details>

<summary>Lab 1: Ansible Basics</summary>

### Demo 1: Ansible Basics

#### Ansible from the CLI via `ansible`

The first example installs a webserver on all hosts in the `webservers` group. The installed webserver is defined as a **host variable** in the inventory file `hosts.yml` (*see above*).

```console
ansible \
   webservers  \
  -i hosts.yml \
  -m package   \
  -a 'name="{{ webserver }}"'
```

#### Ansible from the CLI via `ansible-playbook`

The second example utilizes the following **playbook** to **install** and **start** the defined webserver on all hosts in the `webservers` group.

```yaml
---
- name: Install webservers
  hosts: webservers
  vars:
    package: "{{ webserver }}"
  become: true
  tasks:
    - name: Install webserver
      ansible.builtin.package:
        name: "{{ package }}"
        state: present

    - name: Start webserver
      ansible.builtin.service:
          name: "{{ package }}"
          state: started
```

Save this playbook as `playbook.yml` and run it with the following command.

```console
ansible-playbook \
  -i hosts.yml \
    playbook.yml
```

You will see a separated output for each task in the playbook. In the end, you should be able to access the webserver on each host in the `webservers` group.

> [!TIP]
> Ansible is **idempotent** - try running the playbook again and see how the output differs.

</details>

<details>

<summary>Lab 2: Event-driven Ansible and Generic Webhooks</summary>

### Demo 2: Event-driven Ansible and Generic Webhooks

#### Prerequisites

For this demo, we will use `localhost` as the target host. Therefore, we need to adjust our inventory file `hosts.yml` accordingly:

```yaml
hosts:
  localhost: {}

The first demo of event-driven Ansible shows how to use a generic webhook to trigger a playbook run. Copy the following rulebook into a file called `rulebook.yml`.

```yaml
- name: Listen to webhook events
  hosts: all
  sources:
    - ansible.eda.webhook:
        host: 0.0.0.0
        port: 5000
  rules:
    - name: Debug event output
      condition: event.payload.greeting is defined
      action:
        debug:
          msg: "Hello {{ event.payload.greeting }}!"

    - name: Greet stranger
      condition: 1 == 1  # default case
      action:
        debug:
          msg: Hello World!
```

#### Start the EDA server

To start the EDA server, run the following command.

```console
ansible-rulebook \
  -i hosts.yml \
  --rulebook rulebook.yml
```

#### Trigger the webhook

Once the EDA server is running, we can open a second terminal session and double-check that it is listening on the correct port:

```console
netstat -lntup | grep 5000
```

Now, we can trigger the webhook from our second terminal session using `curl`, first with empty input:

```console
curl \
  -H "Content-Type: application/json" \
  -d '{}' \
  http://localhost:5000/endpoint
```

If we switch over to the first terminal session, we should see the output of the second rule, which is the default case:

```console
Hello World!
```

Now, we can trigger the webhook again, this time with a payload:

```console
curl \
  -H "Content-Type: application/json" \
  -d '{"greeting": "Daniel"}' \
  http://localhost:5000/endpoint
```

If we switch over to the first terminal session again, we should see the output of the first rule, which is the case for a defined `greeting` in the payload:

```console
Hello Daniel!
```

</details>

## Resources

- [Ansible Documentation](https://docs.ansible.com/)
- [Installing Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible)
- [Ansible Galaxy](https://galaxy.ansible.com/)
- [EDA Documentation](https://ansible.readthedocs.io/projects/rulebook/en/stable/introduction.html)
- [Installing and Running EDA](https://ansible.readthedocs.io/projects/rulebook/en/stable/installation.html)
