---
title: '#90DaysOfDevOps - State and Ingress in Kubernetes - Day 55'
published: false
description: 90DaysOfDevOps - State and Ingress in Kubernetes
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048779
---

## Kubernetes의 State와 Ingress

이번 Kubernetes 마지막 섹션에서는 State와 Ingress에 대해 살펴보겠습니다.

지금까지 설명한 모든 것은 상태 비저장에 관한 것이며, 상태 비저장은 실제로 애플리케이션이 어떤 네트워크를 사용하든 상관하지 않고 영구적인 저장소가 필요하지 않은 경우입니다. 예를 들어 Stateful 애플리케이션과 데이터베이스가 제대로 작동하려면, 호스트 이름, IP 등 변경되지 않는 고유 ID를 통해 pod가 서로 연결될 수 있도록 해야 합니다. Stateful 애플리케이션의 예로는 MySQL 클러스터, Redis, Kafka, MongoDB 등이 있습니다. 기본적으로 데이터를 저장하는 모든 애플리케이션을 통해 가능합니다.

### Stateful 애플리케이션

StatefulSet은 Kubernetes가 스케줄링 위치에 관계없이 유지 관리하는 고유하고 영구적인 ID와 안정적인 호스트 이름을 가진 pod의 집합을 나타냅니다. 특정 StatefulSet pod의 상태 정보 및 기타 복원력 있는 데이터는 StatefulSet과 연결된 영속성 디스크 스토리지에 유지됩니다.

### Deployment vs StatefulSet

- Stateful 애플리케이션을 복제하는 것은 더 어렵다.
- 배포(Stateless 애플리케이션)에서 pod를 복제하는 것은 동일하며 상호 교환이 가능하다.
- 임의의 해시를 사용하여 임의의 순서로 pod를 생성한다.
- 모든 pod에 로드 밸런싱하는 하나의 서비스이다.

StatefulSet 또는 Stateful 애플리케이션의 경우 위의 내용이 더 어렵습니다.

- 동시에 생성하거나 삭제할 수 없다.
- 임의로 주소를 지정할 수 없다.
- 레플리카 pod는 동일하지 않다.

곧 데모에서 보게 될 것은 각 pod가 고유한 식별자를 가지고 있다는 것입니다. 상태 비저장 애플리케이션을 사용하면 임의의 이름을 볼 수 있습니다. 예를 들어 `app-7469bbb6d7-9mhxd`인 반면, 상태 저장 애플리케이션은 `mongo-0`에 더 가깝고 확장 시 `mongo-1`이라는 새 pod를 생성합니다.

이러한 pod는 동일한 사양으로 생성되지만 상호 교환할 수는 없습니다. 각 StatefulSet pod는 모든 스케줄링에 걸쳐 영구 식별자를 가지고 있습니다. 이는 데이터베이스에 쓰고 읽어야 하는 데이터베이스와 같은 Stateful 워크로드가 필요할 때, 데이터 불일치를 초래할 수 있기 때문에 두 개의 pod가 인식 없이 동시에 쓰게 할 수 없기 때문에 필요합니다. 주어진 시간에 데이터베이스에 하나의 pod만 쓰도록 해야 하지만 여러 개의 pod가 해당 데이터를 읽을 수 있습니다.

StatefulSet의 각 pod는 영구 볼륨과 데이터베이스의 복제본에 액세스하여 읽을 수 있으며, 이는 마스터로부터 지속적으로 업데이트됩니다. 또한 각 pod는 이 영속성 볼륨에 pod 상태도 저장하는데, 만약 `mongo-0`이 죽으면 새 pod가 프로비저닝될 때 스토리지에 저장된 pod 상태를 이어받게 된다는 점도 흥미롭습니다.

TLDR; StatefulSet vs Deployment

- 예측 가능한 pod 이름 = `mongo-0`
- 고정된 개별 DNS 이름
- pod 아이덴티티 - 상태 유지, 역할 유지
- Stateful 앱 복제는 복잡함
  - 해야 할 일이 많음:
    - 복제 및 데이터 동기화를 구성
    - 원격 공유 스토리지를 사용할 수 있도록 설정
    - 관리 및 백업

### 영속성 볼륨 | Claims | StorageClass

Kubernetes에서 데이터를 어떻게 지속하나요?

위에서 상태 저장 애플리케이션이 있을 때 상태를 어딘가에 저장해야 한다고 언급했는데, 바로 이 부분에서 볼륨의 필요성이 대두되는데, Kubernetes는 기본적으로 지속성을 제공하지 않습니다.

pod 라이프사이클에 의존하지 않는 스토리지 계층이 필요합니다. 이 스토리지는 모든 Kubernetes 노드에서 사용할 수 있고 액세스할 수 있어야 합니다. 또한 이 스토리지는 Kubernetes 클러스터가 충돌하더라도 생존할 수 있도록 Kubernetes 클러스터 외부에 있어야 합니다.

### 영속성 볼륨

- 데이터를 저장하기 위한 클러스터 리소스(예: CPU 및 RAM)
- YAML 파일을 통해 생성
- 실제 물리적 스토리지(NAS)가 필요
- Kubernetes 클러스터에 대한 외부 통합
- 스토리지에 다양한 유형의 스토리지를 사용
- PV는 네임스페이스가 없음
- 로컬 스토리지를 사용할 수 있지만 클러스터의 한 노드에 한정
- 데이터베이스 지속성은 원격 스토리지(NAS)를 사용

### 영구 볼륨 Claims

위의 영구 볼륨만 있어도 사용할 수 있지만 애플리케이션에서 Claims하지 않으면 사용되지 않습니다.

- YAML 파일을 통해 생성
- 영속성 볼륨 Claims은 pod 구성(볼륨 어트리뷰트)에서 사용
- PVC는 pod와 동일한 네임스페이스에 존재
- 볼륨이 pod에 마운트됨
- pod는 여러 가지 볼륨 유형(ConfigMaps, Secrets, PVC)을 사용

PV와 PVC를 생각하는 또 다른 방법은 다음과 같습니다.

PV는 Kubernetes 어드민에 의해 생성됩니다.
PVC는 사용자 또는 애플리케이션 개발자가 생성합니다.

또한 자세히 설명하지는 않겠지만 언급할 가치가 있는 두 가지 다른 유형의 볼륨이 있습니다:

### ConfigMaps | Secrets

- pod의 구성 파일
- pod의 인증서 파일

### StorageClass

- YAML 파일을 통해 생성
- PVC가 영속성 볼륨을 Claims할 때 동적으로 프로비저닝
- 각 스토리지 백엔드에는 프로비저너가 있음
- 스토리지 백엔드는 (프로비저너 속성을 통해) YAML에 정의됨
- 기본 스토리지 공급자 추상화
- 해당 스토리지에 대한 파라미터 정의

### 연습 시간

어제 세션에서는 상태 비저장 애플리케이션을 생성하는 방법을 살펴봤는데, 여기서는 동일한 작업을 수행하되 Minikube 클러스터를 사용하여 상태 저장 워크로드를 배포하고자 합니다.

지속성을 사용하는 기능과 애드온을 갖기 위해 사용하는 Minikube 명령에 대한 요약은 `minikube start --addons volumesnapshots,csi-hostpath-driver --apiserver-port=6443 --container-runtime=containerd -p mc-demo --kubernetes-version=1.21.2`입니다.

이 명령은 나중에 보여드리는 스토리지 클래스를 제공하는 CSI-hostpath-driver를 사용합니다.

애플리케이션의 빌드 아웃은 아래와 같습니다:

![](/2022/Days/Images/Day55_Kubernetes1.png)

이 애플리케이션의 YAML 구성 파일은 여기에서 찾을 수 있습니다. [pacman-stateful-demo.yaml](/2022/Days/Kubernetes)

### StorageClass 구성

애플리케이션 배포를 시작하기 전에 실행해야 하는 한 단계가 더 있는데, 그것은 StorageClass(CSI-hostpath-sc)가 기본 StorageClass인지 확인하는 것입니다. 먼저 `kubectl get storageclass` 명령을 실행하여 확인할 수 있지만, 기본적으로 Minikube 클러스터는 표준 스토리지 클래스를 기본값으로 표시하므로 다음 명령으로 변경해야 합니다.

이 첫 번째 명령은 CSI-hostpath-sc 스토리지 클래스를 기본값으로 설정합니다.

`kubectl patch storageclass csi-hostpath-sc -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'`

이 명령은 표준 StorageClass에서 기본 어노테이션을 제거합니다.

`kubectl patch storageclass standard -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"false"}}}'`

![](/2022/Days/Images/Day55_Kubernetes2.png)

클러스터에 Pacman 네임스페이스가 없는 상태에서 시작합니다. `kubectl get namespace`

![](/2022/Days/Images/Day55_Kubernetes3.png)

그런 다음 YAML 파일을 배포합니다. `kubectl create -f pacman-stateful-demo.yaml` 이 명령에서 우리는 Kubernetes 클러스터 내에 여러 개의 오브젝트를 생성하고 있음을 볼 수 있습니다.

![](/2022/Days/Images/Day55_Kubernetes4.png)

이제 새로 생성된 네임스페이스가 생겼습니다.

![](/2022/Days/Images/Day55_Kubernetes5.png)

다음 이미지와 `kubectl get all -n pacman` 명령에서 네임스페이스 내부에서 여러 가지 일이 일어나고 있음을 확인할 수 있습니다. 우리는 NodeJS 웹 프론트엔드를 실행하는 pod를 가지고 있고, 백엔드 데이터베이스를 실행하는 mongo를 가지고 있습니다. Pacman과 mongo 모두 해당 pod에 액세스하기 위한 서비스가 있습니다. Pacman을 위한 배포와 mongo를 위한 StatefulSet이 있습니다.

![](/2022/Days/Images/Day55_Kubernetes6.png)

또한 영속성 볼륨과 영속성 볼륨 Claims도 가지고 있는데, `kubectl get pv`를 실행하면 네임스페이스가 없는 영속성 볼륨을 얻을 수 있고, `kubectl get pvc -n pacman`을 실행하면 네임스페이스가 있는 영속성 볼륨 Claims을 얻을 수 있습니다.

![](/2022/Days/Images/Day55_Kubernetes7.png)

### 게임 플레이하기 | 미션 크리티컬 애플리케이션에 액세스하기

앞서 언급한 바와 같이 상태 비저장 애플리케이션에서 Minikube를 사용하고 있기 때문에 애플리케이션에 액세스하는 데 있어 몇 가지 장애물이 있지만, 클러스터 내에서 Ingress 또는 로드 밸런서에 액세스하여 외부에서 액세스하기 위해 자동으로 IP를 받도록 설정되어 있습니다. (위의 Pacman 네임스페이스의 모든 구성 요소 이미지에서 이를 확인할 수 있습니다).

이 데모에서는 포트 포워드 방법을 사용하여 애플리케이션에 액세스하겠습니다. 새 터미널을 열고 다음 `kubectl port-forward svc/pacman 9090:80 -n pacman` 명령을 실행하고 브라우저를 열면 이제 애플리케이션에 액세스할 수 있습니다. AWS 또는 특정 위치에서 실행하는 경우, 위의 스크린샷에서 이 pod 이름을 다시 한번 확인하면 클라우드와 영역은 물론 Kubernetes 내의 pod와 동일한 호스트에 대해서도 보고됩니다.

![](/2022/Days/Images/Day55_Kubernetes8.png)

이제 데이터베이스에 저장할 높은 점수를 생성할 수 있습니다.

![](/2022/Days/Images/Day55_Kubernetes9.png)

좋아요, 이제 높은 점수를 얻었지만 `mongo-0` pod를 삭제하면 어떻게 되나요? `kubectl delete pod mongo-0 -n pacman`을 실행하면 삭제할 수 있으며, 아직 앱에 있는 경우 적어도 몇 초 동안은 높은 점수를 사용할 수 없는 것을 볼 수 있을 것입니다.

![](/2022/Days/Images/Day55_Kubernetes10.png)

이제 게임으로 돌아가서 새 게임을 만들면 제 높은 점수를 확인할 수 있습니다. 하지만 제 말을 진정으로 믿을 수 있는 유일한 방법은 직접 시도해보고 소셜 미디어에 최고 점수를 공유하는 것입니다!

![](/2022/Days/Images/Day55_Kubernetes11.png)

배포를 통해 이전 세션에서 다룬 커맨드를 사용하여 확장할 수 있지만, 특히 대규모 Pacman 파티를 주최하려는 경우 `kubectl scale deployment pacman --replicas=10 -n pacman`을 사용하여 확장할 수 있습니다.

![](/2022/Days/Images/Day55_Kubernetes12.png)

### Ingress 설명

Kubernetes에 대해 마무리하기 전에 Kubernetes의 중요한 측면인 Ingress에 대해서도 다루고 싶었습니다.

### Ingress란 무엇인가요?

지금까지 예제에서는 포트 포워드를 사용하거나 Minikube 내에서 특정 명령을 사용하여 애플리케이션에 액세스했지만, 프로덕션 환경에서는 이 방법이 작동하지 않습니다. 여러 사용자가 대규모로 애플리케이션에 액세스할 수 있는 더 나은 방법이 필요할 것입니다.

또한 NodePort가 옵션이라고 말씀드렸지만, 이 역시 테스트 목적으로만 사용해야 합니다.

Ingress는 애플리케이션을 노출하는 더 나은 방법을 제공하며, 이를 통해 Kubernetes 클러스터 내에서 라우팅 규칙을 정의할 수 있습니다.

Ingress의 경우, 애플리케이션의 내부 서비스에 대한 포워드 요청을 생성합니다.

### 언제 Ingress가 필요한가요?

클라우드 제공자를 사용하는 경우, 관리형 Kubernetes 제품에는 클러스터에 대한 Ingress 옵션이 있거나 로드 밸런서 옵션이 제공될 가능성이 높습니다. 이를 직접 구현할 필요가 없다는 것이 관리형 Kubernetes의 장점 중 하나입니다.

클러스터를 실행하는 경우 엔트리포인트를 구성해야 합니다.

### Minikube에서 Ingress 구성하기

제가 실행 중인 mc-demo라는 특정 클러스터에서 다음 명령을 실행하여 클러스터에서 Ingress를 활성화할 수 있습니다.

`minikube --profile='mc-demo' addons enable ingress`

![](/2022/Days/Images/Day55_Kubernetes13.png)

이제 네임스페이스를 확인하면 새로운 ingress-nginx 네임스페이스가 있는 것을 볼 수 있습니다. `kubectl get ns`

![](/2022/Days/Images/Day55_Kubernetes14.png)

이제 Pacman 서비스를 실행하기 위해 Ingress YAML 구성을 생성해야 합니다. 이 파일을 리포지토리 [pacman-ingress.yaml](/2022/Days/Kubernetes)에 추가했습니다.

그런 다음 `kubectl create -f pacman-ingress.yaml`을 사용하여 Ingress 네임스페이스에 이 파일을 생성할 수 있습니다.

![](/2022/Days/Images/Day55_Kubernetes15.png)

그런 다음 `kubectl get ingress -n pacman`을 실행하면 다음과 같이 출력됩니다.

![](/2022/Days/Images/Day55_Kubernetes16.png)

그러면 윈도우에서 WSL2에서 실행되는 Minikube를 사용하고 있기 때문에 `minikube tunnel --profile=mc-demo`를 사용하여 Minikube 터널을 생성해야 한다는 메시지가 표시됩니다.

하지만 여전히 192.168.49.2에 액세스하여 Pacman 게임을 플레이할 수 없습니다.

누구든지 Windows 및 WSL에서 이 기능을 사용할 수 있거나 사용할 수 있다면 피드백을 보내 주시면 감사하겠습니다. 리포지토리에 이 문제를 제기하고 시간과 수정 사항이 생기면 다시 돌아오겠습니다.

업데이트: 이 블로그가 WSL에서 작동하지 않는 원인을 파악하는 데 도움이 될 것 같습니다 [Docker 런타임을 사용하여 WSL2에서 Minikube를 실행하도록 Ingress 구성하기](https://hellokube.dev/posts/configure-minikube-ingress-on-wsl2/).

## 자료

사용하신 무료 리소스가 있으시면 리포지토리에 PR을 통해 여기에 추가해 주시면 기꺼이 포함시켜 드리겠습니다.

- [Kubernetes StatefulSet simply explained](https://www.youtube.com/watch?v=pPQKAR1pA9U)
- [Kubernetes Volumes explained](https://www.youtube.com/watch?v=0swOh5C3OVM)
- [Kubernetes Ingress Tutorial for Beginners](https://www.youtube.com/watch?v=80Ew_fsV4rM)
- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

이것으로 Kubernetes 섹션을 마무리합니다. Kubernetes에 대해 다룰 수 있는 추가 콘텐츠는 매우 많으며 7일 동안 기초적인 지식을 얻을 수 있지만, 사람들은 [100DaysOfKubernetes](https://100daysofkubernetes.io/overview.html)를 통해 심도 있게 살펴볼 수 있습니다.

다음 시간에는 IaC(Infrastructure as Code)와 이것이 데브옵스 관점에서 수행하는 중요한 역할에 대해 살펴보겠습니다.

[Day 56](day56.md)에서 봐요!
