---
title: '#90DaysOfDevOps - Application Focused Backup - Day 88'
published: false
description: 90DaysOfDevOps - Application Focused Backups
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048749
---
## Application Focused Backups

We have already spent some time talking about data services or data intensive applications such as databases on [Day 85](day85.md). For these data services we have to consider how we manage consistency, especially when it comes application consistency. 

In this post we are going to dive into that requirement around protecting the application data in a consistent manner. 

In order to do this our tool of choice will be [Kanister](https://kanister.io/)

![](Images/Day88_Data1.png)

### Introducing Kanister 

Kanister is an open-source project by Kasten, that enables us to manage (backup and restore) application data on Kubernetes. You can deploy Kanister as a helm application into your Kubernetes cluster. 

Kanister uses Kubernetes custom resources, the main custom resources that are installed when Kanister is deployed are 

- `Profile` - is a target location to store your backups and recover from. Most commonly this will be object storage. 
- `Blueprint` - steps that are to be taken to backup and restore the database should be maintained in the Blueprint
- `ActionSet` - is the motion to move our target backup to our profile as well as restore actions. 
 
### Execution Walkthrough 

Before we get hands on we should take a look at the workflow that Kanister takes in protecting application data. Firstly our controller is deployed using helm into our Kubernetes cluster, Kanister lives within its own namespace. We take our Blueprint of which there are many community supported blueprints available, we will cover this in more detail shortly. We then have our database workload. 

![](Images/Day88_Data2.png)

We then create our ActionSet.   

![](Images/Day88_Data3.png)

The ActionSet allows us to run the actions defined in the blueprint against the specific data service. 

![](Images/Day88_Data4.png)

The ActionSet in turns uses the Kanister functions (KubeExec, KubeTask, Resource Lifecycle) and pushes our backup to our target repository (Profile). 

![](Images/Day88_Data5.png)

If that action is completed/failed the respective status is updated in the Actionset.

![](Images/Day88_Data6.png)

### Deploying Kanister 

Once again we will be using the minikube cluster to achieve this application backup. If you have it still running from the previous session then we can continue to use this. 

At the time of writing we are up to image version `0.75.0` with the following helm command we will install kanister into our Kubernetes cluster. 

`helm install kanister --namespace kanister kanister/kanister-operator --set image.tag=0.75.0 --create-namespace`

![](Images/Day88_Data7.png)

We can use `kubectl get pods -n kanister` to ensure the pod is up and runnnig and then we can also check our custom resource definitions are now available (If you have only installed Kanister then you will see the highlighted 3)

![](Images/Day88_Data8.png)

### Deploy a Database 

Deploying mysql via helm:

```
APP_NAME=my-production-app
kubectl create ns ${APP_NAME}
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install mysql-store bitnami/mysql --set primary.persistence.size=1Gi,volumePermissions.enabled=true --namespace=${APP_NAME}
kubectl get pods -n ${APP_NAME} -w
```
![](Images/Day88_Data9.png)

Populate the mysql database with initial data, run the following:

```
MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace ${APP_NAME} mysql-store -o jsonpath="{.data.mysql-root-password}" | base64 --decode)
MYSQL_HOST=mysql-store.${APP_NAME}.svc.cluster.local
MYSQL_EXEC="mysql -h ${MYSQL_HOST} -u root --password=${MYSQL_ROOT_PASSWORD} -DmyImportantData -t"
echo MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
```

### Create a MySQL CLIENT 
We will run another container image to act as our client

```
APP_NAME=my-production-app
kubectl run mysql-client --rm --env APP_NS=${APP_NAME} --env MYSQL_EXEC="${MYSQL_EXEC}" --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} --env MYSQL_HOST=${MYSQL_HOST} --namespace ${APP_NAME} --tty -i --restart='Never' --image  docker.io/bitnami/mysql:latest --command -- bash
```
```
Note: if you already have an existing mysql client pod running, delete with the command

kubectl delete pod -n ${APP_NAME} mysql-client
```

### Add Data to MySQL

```
echo "create database myImportantData;" | mysql -h ${MYSQL_HOST} -u root --password=${MYSQL_ROOT_PASSWORD}
MYSQL_EXEC="mysql -h ${MYSQL_HOST} -u root --password=${MYSQL_ROOT_PASSWORD} -DmyImportantData -t"
echo "drop table Accounts" | ${MYSQL_EXEC}
echo "create table if not exists Accounts(name text, balance integer); insert into Accounts values('nick', 0);" |  ${MYSQL_EXEC}
echo "insert into Accounts values('albert', 112);" | ${MYSQL_EXEC}
echo "insert into Accounts values('alfred', 358);" | ${MYSQL_EXEC}
echo "insert into Accounts values('beatrice', 1321);" | ${MYSQL_EXEC}
echo "insert into Accounts values('bartholomew', 34);" | ${MYSQL_EXEC}
echo "insert into Accounts values('edward', 5589);" | ${MYSQL_EXEC}
echo "insert into Accounts values('edwin', 144);" | ${MYSQL_EXEC}
echo "insert into Accounts values('edwina', 233);" | ${MYSQL_EXEC}
echo "insert into Accounts values('rastapopoulos', 377);" | ${MYSQL_EXEC}
echo "select * from Accounts;" |  ${MYSQL_EXEC}
exit
```
You should be able to see some data as per below. 

![](Images/Day88_Data10.png)


### Create Kanister Profile

Kanister provides a CLI, `kanctl` and another utility `kando` that is used to interact with your object storage provider from blueprint and both of these utilities. 

[CLI Download](https://docs.kanister.io/tooling.html#tooling)

I have gone and I have created an AWS S3 Bucket that we will use as our profile target and restore location. I am going to be using environment variables so that I am able to still show you the commands I am running with `kanctl` to create our kanister profile. 

`kanctl create profile s3compliant --access-key $ACCESS_KEY --secret-key $SECRET_KEY --bucket $BUCKET --region eu-west-2 --namespace my-production-app`

![](Images/Day88_Data11.png)

### Blueprint time

Don't worry you don't need to create your own one from scratch unless your data service is not listed here in the [Kanister Examples](https://github.com/kanisterio/kanister/tree/master/examples) but by all means community contributions are how this project gains awareness. 

The blueprint we will be using will be the below. 


```
apiVersion: cr.kanister.io/v1alpha1
kind: Blueprint
metadata:
  name: mysql-blueprint
actions:
  backup:
    outputArtifacts:
      mysqlCloudDump:
        keyValue:
          s3path: "{{ .Phases.dumpToObjectStore.Output.s3path }}"
    phases:
    - func: KubeTask
      name: dumpToObjectStore
      objects:
        mysqlSecret:
          kind: Secret
          name: '{{ index .Object.metadata.labels "app.kubernetes.io/instance" }}'
          namespace: '{{ .StatefulSet.Namespace }}'
      args:
        image: ghcr.io/kanisterio/mysql-sidecar:0.75.0
        namespace: "{{ .StatefulSet.Namespace }}"
        command:
        - bash
        - -o
        - errexit
        - -o
        - pipefail
        - -c
        - |
          s3_path="/mysql-backups/{{ .StatefulSet.Namespace }}/{{ index .Object.metadata.labels "app.kubernetes.io/instance" }}/{{ toDate "2006-01-02T15:04:05.999999999Z07:00" .Time  | date "2006-01-02T15-04-05" }}/dump.sql.gz"
          root_password="{{ index .Phases.dumpToObjectStore.Secrets.mysqlSecret.Data "mysql-root-password" | toString }}"
          mysqldump --column-statistics=0 -u root --password=${root_password} -h {{ index .Object.metadata.labels "app.kubernetes.io/instance" }} --single-transaction --all-databases | gzip - | kando location push --profile '{{ toJson .Profile }}' --path ${s3_path} -
          kando output s3path ${s3_path}
  restore:
    inputArtifactNames:
    - mysqlCloudDump
    phases:
    - func: KubeTask
      name: restoreFromBlobStore
      objects:
        mysqlSecret:
          kind: Secret
          name: '{{ index .Object.metadata.labels "app.kubernetes.io/instance" }}'
          namespace: '{{ .StatefulSet.Namespace }}'
      args:
        image: ghcr.io/kanisterio/mysql-sidecar:0.75.0
        namespace: "{{ .StatefulSet.Namespace }}"
        command:
        - bash
        - -o
        - errexit
        - -o
        - pipefail
        - -c
        - |
          s3_path="{{ .ArtifactsIn.mysqlCloudDump.KeyValue.s3path }}"
          root_password="{{ index .Phases.restoreFromBlobStore.Secrets.mysqlSecret.Data "mysql-root-password" | toString }}"
          kando location pull --profile '{{ toJson .Profile }}' --path ${s3_path} - | gunzip | mysql -u root --password=${root_password} -h {{ index .Object.metadata.labels "app.kubernetes.io/instance" }}
  delete:
    inputArtifactNames:
    - mysqlCloudDump
    phases:
    - func: KubeTask
      name: deleteFromBlobStore
      args:
        image: ghcr.io/kanisterio/mysql-sidecar:0.75.0
        namespace: "{{ .Namespace.Name }}"
        command:
        - bash
        - -o
        - errexit
        - -o
        - pipefail
        - -c
        - |
          s3_path="{{ .ArtifactsIn.mysqlCloudDump.KeyValue.s3path }}"
          kando location delete --profile '{{ toJson .Profile }}' --path ${s3_path}
```

To add this we will use the `kubectl create -f mysql-blueprint.yml -n kanister` command 

![](Images/Day88_Data12.png)

### Create our ActionSet and Protect our application 

We will now take a backup of the MySQL data using an ActionSet defining backup for this application. Create an ActionSet in the same namespace as the controller.

`kubectl get profiles.cr.kanister.io -n my-production-app` This command will show us the profile we previously created, we can have multiple profiles configured here so we might want to use specific ones for different ActionSets 

We are then going to create our ActionSet with the following command using `kanctl` 

`kanctl create actionset --action backup --namespace kanister --blueprint mysql-blueprint --statefulset my-production-app/mysql-store --profile my-production-app/s3-profile-dc5zm --secrets mysql=my-production-app/mysql-store`

You can see from the command above we are defining the blueprint we added to the namespace, the statefulset in our `my-production-app` namespace and also the secrets to get into the MySQL application. 

![](Images/Day88_Data13.png)

Check the status of the ActionSet by taking the ActionSet name and using this command `kubectl --namespace kanister describe actionset backup-qpnqv`

Finally we can go and confirm that we now have data in our AWS S3 bucket. 

![](Images/Day88_Data14.png)

### Restore 

We need to cause some damage before we can restore anything, we can do this by dropping our table, maybe it was an accident, maybe it wasn't. 

Connect to our MySQL pod. 

```
APP_NAME=my-production-app
kubectl run mysql-client --rm --env APP_NS=${APP_NAME} --env MYSQL_EXEC="${MYSQL_EXEC}" --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} --env MYSQL_HOST=${MYSQL_HOST} --namespace ${APP_NAME} --tty -i --restart='Never' --image  docker.io/bitnami/mysql:latest --command -- bash
```

You can see that our importantdata db is there with `echo "SHOW DATABASES;" |  ${MYSQL_EXEC}` 

Then to drop we ran `echo "DROP DATABASE myImportantData;" |  ${MYSQL_EXEC}`

And confirmed that this was gone with a few attempts to show our database. 

![](Images/Day88_Data15.png)

We can now use Kanister to get our important data back in business using the `kubectl get actionset -n kanister` to find out our ActionSet name that we took earlier. Then we will create a restore ActionSet to restore our data using `kanctl create actionset -n kanister --action restore --from "backup-qpnqv"`

![](Images/Day88_Data16.png)

We can confirm our data is back by using the below command to connect to our database. 

```
APP_NAME=my-production-app
kubectl run mysql-client --rm --env APP_NS=${APP_NAME} --env MYSQL_EXEC="${MYSQL_EXEC}" --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} --env MYSQL_HOST=${MYSQL_HOST} --namespace ${APP_NAME} --tty -i --restart='Never' --image  docker.io/bitnami/mysql:latest --command -- bash
```
Now we are inside the MySQL Client, we can issue the `echo "SHOW DATABASES;" |  ${MYSQL_EXEC}` and we can see the database is back. We can also issue the `echo "select * from Accounts;" |  ${MYSQL_EXEC}` to check the contents of the database and our important data is restored. 

![](Images/Day88_Data17.png)

In the next post we take a look at Disaster Recovery within Kubernetes. 

## Resources 

- [Kanister Overview - An extensible open-source framework for app-lvl data management on Kubernetes](https://www.youtube.com/watch?v=wFD42Zpbfts)
- [Application Level Data Operations on Kubernetes](https://community.cncf.io/events/details/cncf-cncf-online-programs-presents-cncf-live-webinar-kanister-application-level-data-operations-on-kubernetes/)
- [Kubernetes Backup and Restore made easy!](https://www.youtube.com/watch?v=01qcYSck1c4&t=217s)
- [Kubernetes Backups, Upgrades, Migrations - with Velero](https://www.youtube.com/watch?v=zybLTQER0yY)
- [7 Database Paradigms](https://www.youtube.com/watch?v=W2Z7fbCLSTw&t=520s)
- [Disaster Recovery vs. Backup: What's the difference?](https://www.youtube.com/watch?v=07EHsPuKXc0)
- [Veeam Portability & Cloud Mobility](https://www.youtube.com/watch?v=hDBlTdzE6Us&t=3s)

See you on [Day 89](day89.md)
