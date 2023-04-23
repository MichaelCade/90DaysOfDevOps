---
title: '#90DaysOfDevOps - Hands-On Backup & Recovery - Day 87'
published: false
description: 90DaysOfDevOps - Hands-On Backup & Recovery
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048717
---

## 백업 및 복구 실습

지난 세션에서는 로컬 NAS와 클라우드 기반 오브젝트 스토리지에 중요한 데이터를 백업하는 데 사용한 오픈소스 백업 도구인 [Kopia](https://kopia.io/)에 대해 살펴봤습니다.

이번 섹션에서는 Kubernetes 백업의 세계로 들어가 보겠습니다. 이 플랫폼은 챌린지 초반에 [큰 그림: Kubernetes](/2022/Days/Days/day49.md)에서 다뤘던 플랫폼입니다.

이번에도 Minikube 클러스터를 사용하지만, 이번에는 사용 가능한 몇 가지 애드온을 활용하겠습니다.

### Kubernetes 클러스터 설정

Minikube 클러스터를 설정하기 위해 `minikube start --addons volumesnapshots,csi-hostpath-driver --apiserver-port=6443 --container-runtime=containerd -p 90daysofdevops --kubernetes-version=1.21.2`를 실행하면 백업을 수행할 때 이를 최대한 활용하기 위해 `volumesnapshots` 및 `csi-hostpath-driver`를 사용하고 있는 것을 알 수 있습니다.

이 시점에서는 아직 Kasten K10을 배포하지 않았지만, 클러스터가 가동되면 다음 명령을 실행하여 Kasten K10이 이를 사용할 수 있도록 volumesnapshotclass에 주석을 달려고 합니다.

```Shell
kubectl annotate volumesnapshotclass csi-hostpath-snapclass \
    k10.kasten.io/is-snapshot-class=true
```

또한 다음을 사용하여 기본 저장소 클래스를 표준 기본 저장소 클래스에서 csi-hostpath 저장소 클래스로 변경할 것입니다.

```Shell
kubectl patch storageclass csi-hostpath-sc -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'

kubectl patch storageclass standard -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false"}}}'
```

![](/2022/Days/Images/Day87_Data1.png)

### Kasten K10 배포하기

Kasten Helm 리포지토리를 추가합니다.

`helm repo add kasten https://charts.kasten.io/`

여기에서도 `arkade kasten install k10`을 사용할 수 있지만 데모를 위해 다음 단계를 진행하겠습니다. [자세한 내용](https://blog.kasten.io/kasten-k10-goes-to-the-arkade)

네임스페이스를 생성하고 K10을 배포합니다.(약 5분 정도 소요됨)

`helm install k10 kasten/k10 --namespace=kasten-io --set auth.tokenAuth.enabled=true --set injectKanisterSidecar.enabled=true --set-string injectKanisterSidecar.namespaceSelector.matchLabels.k10/injectKanisterSidecar=true --create-namespace`

![](/2022/Days/Images/Day87_Data1.png)

다음 명령어를 실행하여 pod가 생성되는 것을 확인할 수 있습니다.

`kubectl get pods -n kasten-io -w`

![](/2022/Days/Images/Day87_Data3.png)

포트 포워딩을 통해 K10 대시보드에 접속하고, 새 터미널을 열어 아래 명령을 실행합니다.

`kubectl --namespace kasten-io port-forward service/gateway 8080:8000`

Kasten 대시보드는 `http://127.0.0.1:8080/k10/#/`에서 사용할 수 있습니다.

![](/2022/Days/Images/Day87_Data4.png)

이제 대시보드로 인증하려면 다음 명령어로 얻을 수 있는 토큰이 필요합니다.

```Shell
TOKEN_NAME=$(kubectl get secret --namespace kasten-io|grep k10-k10-token | cut -d " " -f 1)
TOKEN=$(kubectl get secret --namespace kasten-io $TOKEN_NAME -o jsonpath="{.data.token}" | base64 --decode)

echo "Token value: "
echo $TOKEN
```

![](/2022/Days/Images/Day87_Data5.png)

이제 이 토큰을 가져와 브라우저에 입력하면 이메일과 회사 이름을 입력하라는 메시지가 표시됩니다.

![](/2022/Days/Images/Day87_Data6.png)

그러면 Kasten K10 대시보드에 액세스할 수 있습니다.

![](/2022/Days/Images/Day87_Data7.png)

### stateful 애플리케이션 배포

Kubernetes 섹션에서 사용한 stateful 애플리케이션을 사용합니다.

![](/2022/Days/Images/Day55_Kubernetes1.png)

이 애플리케이션의 YAML 구성 파일은 여기에서 찾을 수 있습니다. -> [pacman-stateful-demo.yaml](/2022/Days/Kubernetes/pacman-stateful-demo.yaml)

![](/2022/Days/Images/Day87_Data8.png)

`kubectl get all -n pacman`을 사용하여 다가오는 pod를 확인할 수 있습니다.

![](/2022/Days/Images/Day87_Data9.png)

그런 다음 새 터미널에서 pacman 프론트엔드를 포트 포워드할 수 있습니다. `kubectl port-forward svc/pacman 9090:80 -n pacman`을 실행합니다.

브라우저에서 다른 탭을 열어 http://localhost:9090/ 에접속합니다.

![](/2022/Days/Images/Day87_Data10.png)

시간을 내어 백엔드 MongoDB 데이터베이스에서 높은 점수를 기록하세요.

![](/2022/Days/Images/Day87_Data11.png)

### 높은 점수 보호

이제 데이터베이스에 미션 크리티컬한 데이터가 있으며 이를 잃고 싶지 않습니다. 이 전체 애플리케이션을 보호하기 위해 Kasten K10을 사용할 수 있습니다.

Kasten K10 대시보드 탭으로 돌아가면 Kubernetes 클러스터에 pacman 애플리케이션이 추가되어 애플리케이션 수가 1개에서 2개로 늘어난 것을 볼 수 있습니다.

![](/2022/Days/Images/Day87_Data12.png)

애플리케이션 카드를 클릭하면 클러스터에서 자동으로 검색된 애플리케이션을 볼 수 있습니다.

![](/2022/Days/Images/Day87_Data13.png)

Kasten K10을 사용하면 스토리지 기반 스냅샷을 활용할 수 있을 뿐만 아니라 사본을 오브젝트 스토리지 옵션으로 내보낼 수도 있습니다.

데모에서는 클러스터에 수동 스토리지 스냅샷을 생성한 다음, 고득점 데이터에 불량 데이터를 추가하여 실수로 실수하는 상황을 시뮬레이션해 보겠습니다.

먼저 아래의 수동 스냅샷 옵션을 사용할 수 있습니다.

![](/2022/Days/Images/Day87_Data14.png)

데모에서는 모든 것을 기본값으로 두겠습니다.

![](/2022/Days/Images/Day87_Data15.png)

대시보드로 돌아가면 실행 중인 작업에 대한 상태 보고서가 표시되며, 완료되면 이와 같이 성공적으로 표시됩니다.

![](/2022/Days/Images/Day87_Data16.png)

### 실패 시나리오

이제 애플리케이션에 규범적인 잘못된 변경을 추가하기만 하면 미션 크리티컬 데이터에 치명적인 변경을 수행할 수 있습니다.

아래에서 볼 수 있듯이, 프로덕션 미션 크리티컬 데이터베이스에는 원하지 않는 두 가지 입력이 있습니다.

![](/2022/Days/Images/Day87_Data17.png)

### 데이터 복원

이것은 간단한 데모이며 현실적이지 않은 방식이지만 데이터베이스를 삭제하는 것이 얼마나 쉬운지 보셨나요?

이제 고득점 목록을 실수하기 전의 상태로 조금 더 깔끔하게 정리해 보겠습니다.

애플리케이션 카드와 pacman 탭으로 돌아가면 이제 복원할 수 있는 복원 지점이 하나 생겼습니다.

![](/2022/Days/Images/Day87_Data18.png)

복원을 선택하면 해당 애플리케이션에 연결된 모든 스냅샷과 내보내기를 볼 수 있습니다.

![](/2022/Days/Images/Day87_Data19.png)

복원을 선택하면 사이드 창이 나타나며, 기본 설정을 유지하고 복원을 누릅니다.

![](/2022/Days/Images/Day87_Data20.png)

복원할 것인지 확인합니다.

![](/2022/Days/Images/Day87_Data21.png)

그런 다음 대시보드로 돌아가서 복원 진행 상황을 확인할 수 있습니다. 다음과 같은 화면이 표시됩니다.

![](/2022/Days/Images/Day87_Data22.png)

하지만 더 중요한 것은 미션 크리티컬 애플리케이션에서 고점수 목록이 어떻게 보이는지입니다. 앞서 다룬 대로 pacman으로 포트 포워딩을 다시 시작해야 합니다.

![](/2022/Days/Images/Day87_Data23.png)

매우 간단한 데모이며, 백업과 관련하여 Kasten K10이 달성할 수 있는 것의 표면적인 부분만 다루었습니다. 앞으로 이러한 영역 중 일부에 대해 좀 더 심층적인 비디오 콘텐츠를 제작할 예정입니다. 또한 재해 복구 및 데이터의 이동성과 관련하여 데이터 관리와 관련된 몇 가지 다른 주요 영역을 강조하기 위해 Kasten K10을 사용할 것입니다.

다음에는 애플리케이션 일관성에 대해 살펴보겠습니다.

## 자료

- [Kubernetes Backup and Restore made easy!](https://www.youtube.com/watch?v=01qcYSck1c4&t=217s)
- [Kubernetes Backups, Upgrades, Migrations - with Velero](https://www.youtube.com/watch?v=zybLTQER0yY)
- [7 Database Paradigms](https://www.youtube.com/watch?v=W2Z7fbCLSTw&t=520s)
- [Disaster Recovery vs. Backup: What's the difference?](https://www.youtube.com/watch?v=07EHsPuKXc0)
- [Veeam Portability & Cloud Mobility](https://www.youtube.com/watch?v=hDBlTdzE6Us&t=3s)

[Day 88](day88.md)에서 봐요!
