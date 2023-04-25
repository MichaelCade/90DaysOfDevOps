---
title: '#90DaysOfDevOps - Application Focused Backup - Day 88'
published: false
description: 90DaysOfDevOps - Application Focused Backups
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048749
---

## 애플리케이션 중심 백업

데이터 서비스 또는 데이터베이스와 같은 데이터 집약적인 애플리케이션에 대해서는 이미 [day 85](day85.md)에서 설명한 바 있습니다. 이러한 데이터 서비스의 경우 특히 애플리케이션 일관성과 관련하여 일관성을 관리하는 방법을 고려해야 합니다.

이 글에서는 애플리케이션 데이터를 일관되게 보호하는 것과 관련된 요구 사항을 자세히 살펴보겠습니다.

이를 위해 우리가 선택한 도구는 [Kanister](https://kanister.io/)입니다.

![](/2022/Days/Images/Day88_Data1.png)

### Kanister 소개

Kanister는 Kasten이 만든 오픈소스 프로젝트로, Kubernetes에서 애플리케이션 데이터를 관리(백업 및 복원)할 수 있게 해줍니다. Kanister를 Helm 애플리케이션으로 Kubernetes 클러스터에 배포할 수 있습니다.

Kanister는 Kubernetes 커스텀 리소스를 사용하며, Kanister를 배포할 때 설치되는 주요 커스텀 리소스는 다음과 같습니다.

- `Profile` - 백업을 저장하고 복구할 대상 위치입니다. 가장 일반적으로는 오브젝트 스토리지입니다.
- `Blueprint` - 데이터베이스를 백업 및 복구하기 위해 수행해야 하는 단계가 Blueprint에 유지되어야 합니다.
- `ActionSet` - 대상 백업을 Profile로 이동하고 복원 작업을 수행하는 동작입니다.

### 실행 연습

실습에 들어가기 전에 Kanister가 애플리케이션 데이터를 보호하기 위해 취하는 워크플로우를 살펴보겠습니다. 먼저 컨트롤러를 Helm을 사용하여 Kubernetes 클러스터에 배포하고, Kanister는 해당 네임스페이스 내에 위치합니다. 커뮤니티에서 지원하는 많은 Blueprint를 사용할 수 있으며, 이에 대해서는 곧 자세히 다루겠습니다. 그런 다음 데이터베이스 워크로드가 있습니다.

![](/2022/Days/Images/Day88_Data2.png)

이제 ActionSet을 생성합니다.

![](/2022/Days/Images/Day88_Data3.png)

ActionSet을 사용하면 특정 데이터 서비스에 대해 Blueprint에 정의된 액션을 실행할 수 있습니다.

![](/2022/Days/Images/Day88_Data4.png)

ActionSet은 차례로 Kanister 함수(KubeExec, KubeTask, 리소스 라이프사이클)를 사용하여 백업을 대상 리포지토리(Profile)로 push합니다.

![](/2022/Days/Images/Day88_Data5.png)

해당 작업이 완료/실패하면 해당 상태가 ActionSet에서 업데이트됩니다.

![](/2022/Days/Images/Day88_Data6.png)

### Kanister 배포

이번에도 Minikube 클러스터를 사용하여 애플리케이션 백업을 수행합니다. 이전 세션에서 계속 실행 중이라면 이 클러스터를 계속 사용할 수 있습니다.

이 글을 쓰는 시점에, 우리는 다음 Helm 명령으로 이미지 버전 `0.75.0`을 사용하여 Kubernetes 클러스터에 kanister를 설치합니다.

`helm install kanister --namespace kanister kanister/kanister-operator --set image.tag=0.75.0 --create-namespace`

![](/2022/Days/Images/Day88_Data7.png)

`kubectl get pods -n kanister`를 사용하여 pod가 실행 중인지 확인한 다음 사용자 정의 리소스 정의가 사용 가능한지 확인할 수 있습니다.(Kanister만 설치한 경우 강조 표시된 3이 표시됩니다.)

![](/2022/Days/Images/Day88_Data8.png)

### 데이터베이스 배포하기

Helm을 통해 MySQL을 배포합니다:

```Shell
APP_NAME=my-production-app
kubectl create ns ${APP_NAME}
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install mysql-store bitnami/mysql --set primary.persistence.size=1Gi,volumePermissions.enabled=true --namespace=${APP_NAME}
kubectl get pods -n ${APP_NAME} -w
```

![](/2022/Days/Images/Day88_Data9.png)

초기 데이터로 MySQL 데이터베이스를 채우고 다음을 실행합니다:

```Shell
MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace ${APP_NAME} mysql-store -o jsonpath="{.data.mysql-root-password}" | base64 --decode)
MYSQL_HOST=mysql-store.${APP_NAME}.svc.cluster.local
MYSQL_EXEC="mysql -h ${MYSQL_HOST} -u root --password=${MYSQL_ROOT_PASSWORD} -DmyImportantData -t"
echo MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}
```

### MySQL 클라이언트 생성하기

클라이언트 역할을 할 다른 컨테이너 이미지를 실행합니다.

```Shell
APP_NAME=my-production-app
kubectl run mysql-client --rm --env APP_NS=${APP_NAME} --env MYSQL_EXEC="${MYSQL_EXEC}" --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} --env MYSQL_HOST=${MYSQL_HOST} --namespace ${APP_NAME} --tty -i --restart='Never' --image  docker.io/bitnami/mysql:latest --command -- bash
```

```Shell
참고: 이미 기존 MySQL 클라이언트 pod가 실행 중인 경우 다음 명령으로 삭제하세요.

kubectl delete pod -n ${APP_NAME} mysql-client
```

### MySQL에 데이터 추가하기

```Shell
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

아래와 같은 데이터를 볼 수 있을 것입니다.

![](/2022/Days/Images/Day88_Data10.png)

### Kanister Profile 생성

Kanister는 Blueprint와 이 두 유틸리티를 통해 오브젝트 스토리지 공급자와 상호작용할 수 있는 CLI인 `kanctl`과 또 다른 유틸리티인 `kando`를 제공합니다.

[CLI 다운로드](https://docs.kanister.io/tooling.html#tooling)

이제 Profile 대상과 복원 위치로 사용할 AWS S3 버킷을 생성했습니다. 환경 변수를 사용하여 `kanctl`로 실행하는 명령을 계속 보여드릴 수 있도록 환경 변수를 사용하여 kanister Profile을 생성하겠습니다.

`kanctl create profile s3compliant --access-key $ACCESS_KEY --secret-key $SECRET_KEY --bucket $BUCKET --region eu-west-2 --namespace my-production-app`

![](/2022/Days/Images/Day88_Data11.png)

### Blueprint 시간

[Kanister 예제](https://github.com/kanisterio/kanister/tree/master/examples)에 나열되어 있지 않은 데이터 서비스가 아니라면 처음부터 만들 필요는 없지만, 커뮤니티 기여를 통해 이 프로젝트의 인지도를 높일 수 있습니다.

우리가 사용할 Blueprint는 아래와 같습니다.

```Shell
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

이를 추가하기 위해 `kubectl create -f mysql-blueprint.yml -n kanister` 명령을 사용합니다.

![](/2022/Days/Images/Day88_Data12.png)

### ActionSet을 생성하고 애플리케이션 보호하기

이제 이 애플리케이션에 대한 백업을 정의하는 ActionSet을 사용하여 MySQL 데이터의 백업을 수행하겠습니다. 컨트롤러와 동일한 네임스페이스에 ActionSet을 생성합니다.

`kubectl get profiles.cr.kanister.io -n my-production-app` 명령은 이전에 생성한 Profile을 표시하며, 여기에 여러 Profile을 구성할 수 있으므로 다른 ActionSet에 특정 Profile을 사용할 수 있습니다.

그런 다음 `kanctl`을 사용하여 다음 명령으로 ActionSet을 생성합니다.

`kanctl create actionset --action backup --namespace kanister --blueprint mysql-blueprint --statefulset my-production-app/mysql-store --profile my-production-app/s3-profile-dc5zm --secrets mysql=my-production-app/mysql-store`

위의 명령에서 네임스페이스에 추가한 Blueprint, `my-production-app` 네임스페이스의 statefulset, 그리고 MySQL 애플리케이션에 들어가기 위한 시크릿을 정의하고 있음을 알 수 있습니다.

![](/2022/Days/Images/Day88_Data13.png)

ActionSet 이름을 가져와서 다음 명령 `kubectl --namespace kanister describe actionset backup-qpnqv`를 사용하여 ActionSet의 상태를 확인합니다.

마지막으로, 이제 AWS S3 버킷에 데이터가 있는지 확인할 수 있습니다.

![](/2022/Days/Images/Day88_Data14.png)

### 복원

복원하기 전에 약간의 손상을 입혀야 합니다. 테이블을 drop할 수도 있고, 실수일 수도 있고 그렇지 않을 수도 있습니다.

MySQL pod에 연결합니다.

```Shell
APP_NAME=my-production-app
kubectl run mysql-client --rm --env APP_NS=${APP_NAME} --env MYSQL_EXEC="${MYSQL_EXEC}" --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} --env MYSQL_HOST=${MYSQL_HOST} --namespace ${APP_NAME} --tty -i --restart='Never' --image  docker.io/bitnami/mysql:latest --command -- bash
```

중요 데이터 DB는 `echo "SHOW DATABASES;" | ${mysql_exec}`로 확인할 수 있습니다.

그런 다음 drop하기 위해 `echo "DROP DATABASE myImportantData;" | ${mysql_exec}`를 실행했습니다.

그리고 데이터베이스를 보여주기 위해 몇 번의 시도를 통해 이것이 사라진 것을 확인했습니다.

![](/2022/Days/Images/Day88_Data15.png)

이제 Kanister를 사용하여 `kubectl get actionset -n kanister`를 사용하여 앞서 가져 온 ActionSet 이름을 찾아 중요한 데이터를 다시 가져올 수 있습니다. 그런 다음 `kanctl create actionset -n kanister --action restore --from "backup-qpnqv"`를 사용하여 데이터를 복원하기 위한 복원 ActionSet을 생성합니다.

![](/2022/Days/Images/Day88_Data16.png)

아래 명령어를 사용하여 데이터베이스에 연결하면 데이터가 복구되었는지 확인할 수 있습니다.

```Shell
APP_NAME=my-production-app
kubectl run mysql-client --rm --env APP_NS=${APP_NAME} --env MYSQL_EXEC="${MYSQL_EXEC}" --env MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} --env MYSQL_HOST=${MYSQL_HOST} --namespace ${APP_NAME} --tty -i --restart='Never' --image  docker.io/bitnami/mysql:latest --command -- bash
```

이제 MySQL 클라이언트에 들어가서 `echo "SHOW DATABASES;" | ${MYSQL_EXEC}`를 실행하면 데이터베이스가 다시 돌아온 것을 확인할 수 있습니다. 또한 `"select * from Accounts;" | ${MYSQL_EXEC}`를 실행하여 데이터베이스의 내용을 확인할 수 있으며 중요한 데이터가 복원되었습니다.

![](/2022/Days/Images/Day88_Data17.png)

다음 포스트에서는 Kubernetes 내의 재해 복구에 대해 살펴보겠습니다.

## 자료

- [Kanister Overview - An extensible open-source framework for app-lvl data management on Kubernetes](https://www.youtube.com/watch?v=wFD42Zpbfts)
- [Application Level Data Operations on Kubernetes](https://community.cncf.io/events/details/cncf-cncf-online-programs-presents-cncf-live-webinar-kanister-application-level-data-operations-on-kubernetes/)
- [Kubernetes Backup and Restore made easy!](https://www.youtube.com/watch?v=01qcYSck1c4&t=217s)
- [Kubernetes Backups, Upgrades, Migrations - with Velero](https://www.youtube.com/watch?v=zybLTQER0yY)
- [7 Database Paradigms](https://www.youtube.com/watch?v=W2Z7fbCLSTw&t=520s)
- [Disaster Recovery vs. Backup: What's the difference?](https://www.youtube.com/watch?v=07EHsPuKXc0)
- [Veeam Portability & Cloud Mobility](https://www.youtube.com/watch?v=hDBlTdzE6Us&t=3s)

[Day 89](day89.md)에서 봐요!
