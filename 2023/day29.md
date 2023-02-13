# Recap

Last day we discussed why monitoring, logging and auditing are the basics of runtime defense. In short: you cannot protect a live system without knowing what is happening. We built a Minikube cluster yesterday with Prometheus and Grafana. We are continuing to build over this stack today.
Let's start 😎

# Application logging

Application logs are important from many perspective. This is the way operators know what is happening inside applications they run on their infrastrucutre. For the same reason, keeping application logs is important from a security perspective because they provide a detailed record of the system's activity, which can be used to detect and investigate security incidents.

By analyzing application logs, security teams can identify unusual or suspicious activity, such as failed login attempts, access attempts to sensitive data, or other potentially malicious actions. Logs can also help track down the source of security breaches, including when and how an attacker gained access to the system, and what actions they took once inside.

In addition, application logs can help with compliance requirements, such as those related to data protection and privacy. By keeping detailed logs, organizations can demonstrate that they are taking the necessary steps to protect sensitive data and comply with regulations.

Loki is a component in the Grafana stack which collects logs using Promtail for Pods running in the Kubernetes cluster and stores them just as Prometheus does for metrics.

To install Loki with Promtail on your cluster, install the following Helm chart.

```bash
helm install loki --namespace=monitoring grafana/loki-stack
```

This will put a Promtail and a Loki instance in your Minikube and will start collecting logs. Note that this installation in not production grade and it is here to demonstrate the capabilities.

You should be seeing the Pods are ready: 
```bash
$ kubectl get pods | grep loki
loki-0                                               1/1     Running       0             8m25s
loki-promtail-mpwgq                                  1/1     Running       0             8m25s
```

Now go to your Grafana UI (just as we did yesterday):

```bash
kubectl get secret --namespace default grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
minikube service grafana-np --url
```

Take the secret of the admin password (if you haven't changed it already) and print the URL of the service, then go to the URL and log in.

In order to see the logs in Grafana, we need to hook up Loki as a "data source" just as we did yesterday with Prometheus. 

![](images/day29-1.gif)

Now add here a new Loki data source.

The only thing that needs to be changed in the default configuration is the endpoint of the Loki service, in our case it is http://loki:3100, see it below:

![](images/day29-2.png)

Now click "Save & test" and your Grafana should be now connected to Loki.

You can explore your logs in the "Explore" screen (click Explore in the left menu).

To try our centralized logging system, we are going to check when Etcd container did compactization in the last hour.

Choose Loki source on the top of the screen (left of the explore title) and switch from query builder mode (visual builder) to code.

Add the following line in the query field:
```
{container="etcd"} |= `compaction`
```
and click "run query" on the top right part of the screen.

You should see logs in your browser, like this:

![](images/day29-3.png)


Voila! You have a logging system ;-)


# Application behavior monitoring


# Next...

Tomorrow we will continue to the application level. Application logs and behavior monitoring will be in focue. We will continue to use the same setup and go deeper into the rabbit hole 😄
