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

> [!NOTE]
> The below content is a copy of the [lab repository's](https://github.com/mocdaniel/lab-event-driven-ansible) README for convenience.

---
# Event-Driven Ansible Lab

This is a lab designed to demonstrate Ansible and how Event-Driven Ansible (**EDA**) builds on top of its capabilities.

The setup is done with Ansible, too. It will install **Ansible, EDA, Prometheus**, and **Alertmanager** on a VM to demonstrate some of the capabilities of EDA.

## Prerequisites

To follow along with this lab in its entirety, you will need three VMs:

> [!NOTE]
> If you want to skip Ansible basics and go straight to EDA, you'll need just the `eda-controller.example.com` VM and can skip the others.

| VM name            | OS          |
|--------------------|-------------|
| eda-controller.example.com | CentOS/Rocky 8.9 |
| company.example.com        | CentOS/Rocky 8.9 |
| webshop.example.com       | Ubuntu 22.04     |

**You'll need to be able to SSH to each of these VMs as root using SSH keys.**

## Lab Setup

### Clone the repository and create a Python virtual environment

```bash
git clone https://github.com/mocdaniel/lab-event-driven-ansible.git
cd lab-event-driven-ansible
python3 -m venv .venv
source .venv/bin/activate
```

### Install Ansible and other dependencies

```bash
pip install -r requirements.txt
```

### Create the inventory file

```yaml
---
# hosts.yml
webservers:
  hosts:
    webshop.example.com:
      ansible_host: <ip-address>
      webserver: apache2
    company.example.com:
      ansible_host: <ip-address>
      webserver: httpd
eda_controller:
  hosts:
    eda-controller.example.com:
      ansible_host: <ip-address>
```

### Install Needed Roles and Collections

```bash
ansible-galaxy install -r requirements.yml
```

### Run the Setup Playbook

After you created the inventory file and filled in the IP addresses, you can run the setup playbook:

```bash
ansible-playbook playbooks/setup.yml
```

> [!CAUTION]
> Due to a known bug with Python on MacOS, you need to run `export NO_PROXY="*"` on MacOS before running the playbook

---

## Demos

### Lab 1: Ansible Basics

<details>

<summary>Ansible from the CLI via ansible</summary>

#### Ansible from the CLI via `ansible`

The first example installs a webserver on all hosts in the `webservers` group. The installed webserver is defined as a **host variable** in the inventory file `hosts.yml` (*see above*).

```console
ansible \
   webservers  \
  -m package   \
  -a 'name="{{ webserver }}"' \
  --one-line
```

Afterwards, we can start the webserver on all hosts in the `webservers` group.

```console
ansible \
   webservers  \
  -m service   \
  -a 'name="{{ webserver }}" state=started' \
  --one-line
```

Go on and check if the web servers are running on the respective hosts.

> [!HINT]
> Ansible is **idempotent** - try running the commands again and see how the output differs.

</details>

<details>

<summary>Ansible from the CLI via ansible-playbook</summary>

#### Ansible from the CLI via `ansible-playbook`

The second example utilizes the following **playbook** to **gather** and **display information** for all hosts in the `webservers` group, utilizing the **example** role from the lab repository.

```yaml
---
- name: Example role
  hosts: webservers
  gather_facts: false
  vars:
    greeting: "Hello World!"
  pre_tasks:
    - name: Say Hello
      ansible.builtin.debug:
        msg: "{{ greeting }}"
  roles:
    - role: example
  post_tasks:
    - name: Say goodbye
      ansible.builtin.debug:
        msg: Goodbye!
```

```console
ansible-playbook \
    playbooks/example.yml
```

</details>

### Lab 2: Event-Driven Ansible

<details>

<summary>Receive Generic Events via Webhook</summary>

#### Receive Generic Events via Webhook

If you followed the setup instructions for the EDA lab, you should already have a running EDA instance on the `eda-controller.example.com` VM.

If you navigate to `/etc/edacontroller/rulebook.yml` on the VM, you'll see the following rulebook:

```yaml
---
- name: Listen to webhook events
  hosts: all
  sources:
    - ansible.eda.webhook:
        host: 0.0.0.0
        port: 5000
  rules:
    - name: Debug event output
      condition: 1 == 1
      action:
        debug:
          msg: "{{ event }}"

- name: Listen to Alertmanager alerts
  hosts: all
  sources:
    - ansible.eda.alertmanager:
        host: 0.0.0.0
        port: 9000
        data_alerts_path: alerts
        data_host_path: labels.instance
        data_path_separator: .
  rules:
    - name: Restart MySQL server
      condition: event.alert.labels.alertname == 'MySQL not running' and event.alert.status == 'firing'
      action:
        run_module:
          name: ansible.builtin.service
          module_args:
            name: mysql
            state: restarted
    - name: Debug event output
      condition: 1 == 1
      action:
        debug:
          msg: "{{ event }}"

```

For this part of the lab, the **first rule** is the one we're interested in: It listens to a generic webhook on port `5000` and prints the event's **metadata** to its logs.

To test this, we can use the `curl` command to send a `POST` request to the webhook `/endpoint` from the VM itself:

```console
curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"foo": "bar"}' \
  http://localhost:5000/endpoint
```

If you now check the logs of the EDA controller, you should see the following output:

```console
journalctl -fu eda-controller

Jan 11 16:35:29 eda-controller ansible-rulebook[56882]: {'payload': {'foo': 'bar'}, 'meta': {'endpoint': 'endpoint',
'headers': {'Host': 'localhost:5000', 'User-Agent': 'curl/7.76.1', 'Accept': '*/*', 'Content-Length': '21',
'Content-Type': 'application/x-www-form-urlencoded'}, 'source': {'name': 'ansible.eda.webhook', 'type': 'ansible.eda.webhook'},
'received_at': '2024-01-11T15:35:29.798401Z', 'uuid': '6ebf8dd2-60a2-455a-9383-97b81f535366'}}
```

A rule that always evaluates to `true` is not very useful, so let's change the rule to only print the the value of `foo` if the `foo` key is present in the event's payload, and `no foo :(` otherwise:

```yaml
---
- name: Listen to webhook events
  hosts: all
  sources:
    - ansible.eda.webhook:
        host: 0.0.0.0
        port: 5000
  rules:
    - name: Foo
      condition: event.payload.foo is defined
      action:
        debug:
          msg: "{{ event.payload.foo }}"
    - name: No foo
      condition: 1 == 1
      action:
        debug:
          msg: "no foo :("
```

Send the same `curl` request again and check the logs, you should see a line saying `bar` now.

Let's also try a `curl` request with a different payload:

```console
curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"bar": "baz"}' \
  http://localhost:5000/endpoint
```

This time, the output should be `no foo :(`.

</details>

<details>

<summary>Restarting Services Automatically with EDA</summary>

#### Restarting Services Automatically with EDA

The last lab is more of a demo - it shows how you can use EDA to automatically react on events observed by **Prometheus** and **Alertmanager**.

For this demo, the second **ruleset** in our rulebook is the one we're interested in:

```yaml
- name: Listen to Alertmanager alerts
  hosts: all
  sources:
    - ansible.eda.alertmanager:
        host: 0.0.0.0
        port: 9000
        data_alerts_path: alerts
        data_host_path: labels.instance
        data_path_separator: .
  rules:
    - name: Restart MySQL server
      condition: event.alert.labels.alertname == 'MySQL not running' and event.alert.status == 'firing'
      action:
        run_playbook:
          name: ./playbook.yml
    - name: Debug event output
      condition: 1 == 1
      action:
        debug:
          msg: "{{ event }}"
```

With this rule, we can restart our MySQL server if it's not running! But how do we get the event to trigger? With **Prometheus** and **Alertmanager**!

When you ran the setup playbook, it installed **Prometheus** and **Alertmanager** on the `eda-controller.example.com` VM. You can access the **Prometheus** UI at `http://<eda-controller-ip>:9090` and the **Alertmanager** UI at `http://<eda-controller-ip>:9093`.

It also installed a **Prometheus exporter** for the **MySQL** database that runs on the server.

With this setup, we can now shut down our MySQL server and see what happens - make sure to watch the output of the EDA controller's logs:

```console
systemctl stop mysql
journalctl -fu edacontroller
```


Within 30-90 seconds, you should see EDA running our **playbook** and restarting the MySQL server. You can track that process by watching the Prometheus/Alertmanager UIs for firing alerts.

Once you see the playbook being executed in the logs, you can check the MySQL state once more:

```console
systemctl status mysql
```

MySQL should be up and running again!
</details>
