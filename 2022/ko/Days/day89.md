---
title: '#90DaysOfDevOps - Disaster Recovery - Day 89'
published: false
description: 90DaysOfDevOps - Disaster Recovery
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048718
---

## 재해 복구

장애 시나리오에 따라 복구 요구 사항이 어떻게 달라지는지 이미 언급했습니다. 화재, 홍수, 피의 시나리오의 경우 대부분 재해 상황으로, 워크로드를 완전히 다른 위치에서 가능한 한 빨리 가동하거나 최소한 복구 시간 목표(RTO)를 거의 0에 가깝게 설정해야 할 수 있습니다.

이는 전체 애플리케이션 스택을 standby 환경으로 복제하는 작업을 자동화할 때만 대규모로 달성할 수 있습니다.

이를 통해 클라우드 리전, 클라우드 제공업체 간 또는 온프레미스와 클라우드 인프라 간에 빠른 페일오버를 수행할 수 있습니다.

지금까지의 주제에 따라 몇 세션 전에 배포하고 구성한 Minikube 클러스터를 사용하여 Kasten K10을 사용하여 이를 달성하는 방법에 대해 집중적으로 살펴보겠습니다.

그런 다음 이론상 어느 위치에서나 standby 클러스터 역할을 할 수 있는 또 다른 Minikube 클러스터를 Kasten K10을 설치하여 생성할 것입니다.

Kasten K10은 또한 실행 중인 Kubernetes 클러스터에 문제가 발생할 경우 카탈로그 데이터를 복제하여 새 클러스터에서 사용할 수 있도록 하는 기능을 내장하고 있습니다. [K10 재해 복구](https://docs.kasten.io/latest/operating/dr.html).

### K10에 오브젝트 스토리지 추가

가장 먼저 해야 할 일은 백업을 위한 대상 위치로 오브젝트 스토리지 버킷을 추가하는 것입니다. 이것은 오프사이트 위치 역할을 할 뿐만 아니라 복구할 재해 복구 소스 데이터로도 활용할 수 있습니다.

지난 세션에서 Kanister 데모를 위해 생성한 S3 버킷을 정리했습니다.

![](/2022/Days/Images/Day89_Data1.png)

포트 포워딩을 통해 K10 대시보드에 접속하고 새 터미널을 열어 아래 명령을 실행합니다:

`kubectl --namespace kasten-io port-forward service/gateway 8080:8000`

Kasten 대시보드는 `http://127.0.0.1:8080/k10/#/`에서 확인할 수 있습니다.

![](/2022/Days/Images/Day87_Data4.png)

대시보드로 인증하려면 이제 다음 명령으로 얻을 수 있는 토큰이 필요합니다.

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

이제 Kasten K10 대시보드로 돌아왔으므로 location profile을 추가하고 페이지 상단의 "Settings"와 "New Profile"을 선택합니다.

![](/2022/Days/Images/Day89_Data2.png)

아래 이미지에서 이 location profile의 위치를 선택할 수 있는 것을 볼 수 있으며, Amazon S3를 선택하고 민감한 액세스 자격 증명, 지역 및 버킷 이름을 추가할 것입니다.

![](/2022/Days/Images/Day89_Data3.png)

새 프로필 생성 창을 아래로 스크롤하면 S3 Object Lock API를 활용하는 변경 불가능한 백업도 활성화할 수 있다는 것을 알 수 있습니다. 이 데모에서는 이 기능을 사용하지 않겠습니다.

![](/2022/Days/Images/Day89_Data4.png)

"Save Profile"을 누르면 아래와 같이 새로 생성되거나 추가된 location profile을 볼 수 있습니다.

![](/2022/Days/Images/Day89_Data5.png)

### pacman 앱을 오브젝트 스토리지에 보호하는 정책 만들기

이전 세션에서는 pacman 애플리케이션의 임시 스냅샷만 만들었으므로 애플리케이션 백업을 새로 생성한 오브젝트 스토리지 위치로 전송하는 백업 정책을 만들어야 합니다.

대시보드로 돌아가서 정책 카드를 선택하면 아래와 같은 화면이 표시됩니다. "Create New Policy"를 선택합니다.

![](/2022/Days/Images/Day89_Data6.png)

먼저 정책에 유용한 이름과 설명을 지정할 수 있습니다. 또한 on-demand로 사용하는 데모 용도로 백업 빈도를 정의할 수도 있습니다.

![](/2022/Days/Images/Day89_Data7.png)

다음으로 스냅샷 내보내기를 통해 백업을 활성화하여 데이터를 location profile로 내보내려고 합니다. location profile이 여러 개 있는 경우 백업을 전송할 위치를 선택할 수 있습니다.

![](/2022/Days/Images/Day89_Data8.png)

다음으로 이름 또는 레이블로 애플리케이션을 선택하는데, 저는 이름과 모든 리소스를 기준으로 선택하겠습니다.

![](/2022/Days/Images/Day89_Data9.png)

고급 설정에서는 이 중 어떤 것도 사용하지 않을 것이지만, 어제 작성한 [Kanister 사용법](https://github.com/MichaelCade/90DaysOfDevOps/blob/main/Days/day88.md)에 따라 Kasten K10의 일부로 Kanister를 활용하여 데이터의 애플리케이션 일관된 사본을 가져올 수 있습니다.

![](/2022/Days/Images/Day89_Data10.png)

마지막으로 "Create Policy"을 선택하면 이제 정책 창에서 정책을 볼 수 있습니다.

![](/2022/Days/Images/Day89_Data11.png)

생성된 정책의 맨 아래에는 "Show import details"가 있는데, standby 클러스터로 import하기 위해서는 이 문자열이 필요합니다. 지금은 안전한 곳에 복사하세요.

![](/2022/Days/Images/Day89_Data12.png)

계속 진행하기 전에 "run once"을 선택하여 개체 스토리지 버킷으로 백업을 전송하기만 하면 됩니다.

![](/2022/Days/Images/Day89_Data13.png)

아래 스크린샷은 데이터 백업 및 내보내기가 성공적으로 완료된 것을 보여드리기 위한 것입니다.

![](/2022/Days/Images/Day89_Data14.png)

### 새 MiniKube 클러스터 생성 및 K10 배포

그런 다음 두 번째 Kubernetes 클러스터를 배포해야 하며, OpenShift를 포함하여 지원되는 모든 버전의 Kubernetes가 될 수 있지만 교육용으로는 다른 이름의 무료 버전의 MiniKube를 사용하겠습니다.

`minikube start --addons volumesnapshots,csi-hostpath-driver --apiserver-port=6443 --container-runtime=containerd -p standby --kubernetes-version=1.21.2`를 사용하여 새 클러스터를 생성할 수 있습니다.

![](/2022/Days/Images/Day89_Data15.png)

그리고 다음을 사용하여 이 클러스터에 Kasten K10을 배포할 수 있습니다:

`helm install k10 kasten/k10 --namespace=kasten-io --set auth.tokenAuth.enabled=true --set injectKanisterSidecar.enabled=true --set-string injectKanisterSidecar.namespaceSelector.matchLabels.k10/injectKanisterSidecar=true --create-namespace`

이 작업에는 시간이 걸리겠지만, 그동안 `kubectl get pods -n kasten-io -w`를 사용하여 pod가 실행 상태가 되는 진행 상황을 확인할 수 있습니다.

Minikube를 사용하고 있기 때문에 import 정책을 실행할 때 애플리케이션이 실행되며, 이 standby 클러스터에서 스토리지 클래스가 동일하다는 점에 주목할 필요가 있습니다. 그러나 마지막 세션에서 다룰 내용은 이동성과 변환입니다.

pod가 가동되고 실행되면 다른 클러스터에서 이전 단계에서 수행한 단계를 따를 수 있습니다.

포트 포워딩을 통해 K10 대시보드에 액세스하고 새 터미널을 열어 아래 명령을 실행합니다.

`kubectl --namespace kasten-io port-forward service/gateway 8080:8000`

Kasten 대시보드는 `http://127.0.0.1:8080/k10/#/`에서 확인할 수 있습니다.

![](/2022/Days/Images/Day87_Data4.png)

대시보드로 인증하려면 이제 다음 명령으로 얻을 수 있는 토큰이 필요합니다.

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

### pacman을 새로운 Minikube 클러스터로 import

이제 해당 standby 클러스터에서 import 정책을 생성하고 개체 스토리지 백업에 연결하여 어떤 모양과 방식으로 가져올지 결정할 수 있습니다.

먼저, 앞서 다른 클러스터에서 살펴본 location profile을 추가하고, 여기에 다크 모드를 표시하여 프로덕션 시스템과 DR standby location의 차이를 보여줍니다.

![](/2022/Days/Images/Day89_Data16.png)

이제 대시보드로 돌아가서 정책 탭으로 이동하여 새 정책을 생성합니다.

![](/2022/Days/Images/Day89_Data17.png)

아래 이미지에 따라 import 정책을 만듭니다. 완료되면 정책을 만들 수 있습니다. 여기에는 import 후 복원하는 옵션이 있으며, 일부 사용자는 이 옵션을 원할 수 있으며, 이 옵션은 완료되면 standby 클러스터로 복원됩니다. 또한 복원할 때 애플리케이션의 구성을 변경할 수 있으며, 이는 [day 90](day90.md)에 문서화한 내용입니다.

![](/2022/Days/Images/Day89_Data18.png)

on demeand import를 선택했지만 언제 import를 수행할지 일정을 설정할 수 있습니다. 이 때문에 한 번 실행하겠습니다.

![](/2022/Days/Images/Day89_Data19.png)

아래에서 성공적인 import 정책 작업을 확인할 수 있습니다.

![](/2022/Days/Images/Day89_Data20.png)

이제 대시보드로 돌아가서 애플리케이션 카드로 이동하면 아래에 "Removed" 아래에 표시되는 드롭다운을 선택하면 여기에 애플리케이션이 표시됩니다. 복원을 선택합니다.

![](/2022/Days/Images/Day89_Data21.png)

여기에서 사용 가능한 복원 지점을 확인할 수 있습니다. 이것은 기본 클러스터에서 pacman 애플리케이션에 대해 실행한 백업 작업입니다.

![](/2022/Days/Images/Day89_Data22.png)

다음 세션에서 더 자세히 다루고자 하므로 기본값은 변경하지 않겠습니다.

![](/2022/Days/Images/Day89_Data23.png)

"Restore"을 누르면 확인 메시지가 표시됩니다.

![](/2022/Days/Images/Day89_데이터24.png)

아래에서 standby 클러스터에 있는 것을 확인할 수 있으며, pod를 확인하면 실행 중인 애플리케이션이 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day89_Data25.png)

이제 포트 포워딩을 할 수 있습니다.(real-life/프로덕션 환경에서는 애플리케이션에 액세스하기 위해 이 단계가 필요하지 않으며, ingress를 사용할 것입니다).

![](/2022/Days/Images/Day89_Data26.png)

다음으로 애플리케이션 이동성 및 변환에 대해 살펴보겠습니다.

## 자료

- [Kubernetes Backup and Restore made easy!](https://www.youtube.com/watch?v=01qcYSck1c4&t=217s)
- [Kubernetes Backups, Upgrades, Migrations - with Velero](https://www.youtube.com/watch?v=zybLTQER0yY)
- [7 Database Paradigms](https://www.youtube.com/watch?v=W2Z7fbCLSTw&t=520s)
- [Disaster Recovery vs. Backup: What's the difference?](https://www.youtube.com/watch?v=07EHsPuKXc0)
- [Veeam Portability & Cloud Mobility](https://www.youtube.com/watch?v=hDBlTdzE6Us&t=3s)

[Day 90](day90.md)에서 봐요!
