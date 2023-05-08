---
title: '#90DaysOfDevOps - Kubernetes Application Deployment - Day 54'
published: false
description: 90DaysOfDevOps - Kubernetes Application Deployment
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048764
---

## Kubernetes 애플리케이션 배포

이제 드디어 일부 애플리케이션을 클러스터에 배포하게 되었는데, 일부 사람들은 이것이 바로 애플리케이션 배포를 위해 Kubernetes가 존재하는 이유라고 말할 수 있습니다.

여기서 아이디어는 컨테이너 이미지를 가져와서 이제 컨테이너 오케스트레이터로서 Kubernetes를 활용하기 위해 컨테이너 이미지를 Kubernetes 클러스터에 pod로 배포할 수 있다는 것입니다.

### Kubernetes에 앱 배포하기

애플리케이션을 Kubernetes 클러스터에 배포하는 방법에는 여러 가지가 있지만, 가장 일반적인 두 가지 접근 방식인 YAML 파일과 Helm 차트를 다뤄보겠습니다.

이러한 애플리케이션 배포에는 Minikube 클러스터를 사용할 것입니다. 앞서 언급한 Kubernetes의 구성 요소 또는 빌딩 블록 중 일부를 살펴볼 것입니다.

이 섹션과 컨테이너 섹션을 통해 이미지와 Kubernetes의 장점, 그리고 이 플랫폼에서 확장을 매우 쉽게 처리할 수 있는 방법에 대해 논의했습니다.

이 첫 번째 단계에서는 Minikube 클러스터 내에 상태 비저장 애플리케이션을 간단히 생성해 보겠습니다. 사실상의 표준 상태 비저장 애플리케이션인 `nginx`를 사용하여 첫 번째 데모에서 배포를 구성하여 pod를 제공한 다음 nginx pod에서 호스팅하는 간단한 웹 서버로 이동할 수 있는 서비스도 생성할 것입니다. 이 모든 것이 네임스페이스에 포함될 것입니다.

![](/2022/Days/Images/Day54_Kubernetes1.png)

### YAML 생성

첫 번째 데모에서는 YAML로 수행하는 모든 작업을 정의하고자 합니다. YAML에 대한 전체 섹션이 있을 수 있지만, 여기서는 간략히 살펴보고 마지막에 YAML을 더 자세히 다룰 몇 가지 리소스를 남겨두려고 합니다.

다음을 하나의 YAML 파일로 만들 수도 있고, 애플리케이션의 각 측면별로 나눌 수도 있습니다. 즉, 네임스페이스, 배포 및 서비스 생성을 위한 별도의 파일일 수 있지만 이 파일에서는 아래에서 `---`를 사용하여 하나의 파일로 구분했습니다. 이 파일은 [여기](/2022/Days/Kubernetes)에서 찾을 수 있습니다(파일명:- nginx-stateless-demo.YAML).

```Yaml
apiVersion: v1
kind: Namespace
metadata:
  name: nginx
  "labels": {
    "name": "nginx"
  }
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  namespace: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 1
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
  namespace: nginx
spec:
  selector:
    app: nginx-deployment
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

### 클러스터 확인

배포하기 전에 `nginx`라는 네임스페이스가 없는지 확인해야 하는데, `kubectl get namespace` 명령을 실행하여 확인할 수 있으며, 아래에서 볼 수 있듯이 `nginx`라는 네임스페이스가 없습니다.

![](/2022/Days/Images/Day54_Kubernetes2.png)

### 앱을 배포할 시간

이제 Minikube 클러스터에 애플리케이션을 배포할 준비가 되었으며, 이 프로세스는 다른 모든 Kubernetes 클러스터에서도 동일하게 작동합니다.

YAML 파일 위치로 이동한 다음 `kubectl create -f nginx-stateless-demo.yaml`을 실행하면 3개의 오브젝트가 생성되고 네임스페이스, 배포 및 서비스가 생성된 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day54_Kubernetes3.png)

클러스터에서 사용 가능한 네임스페이스를 확인하기 위해 `kubectl get namespace` 명령을 다시 실행하면 이제 새 네임스페이스가 있는 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day54_Kubernetes5.png)

이제 `kubectl get pods -n nginx`를 사용하여 네임스페이스에 pod가 있는지 확인하면 준비 및 실행 상태의 pod 1개가 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day54_Kubernetes4.png)

또한 `kubectl get service -n nginx`를 실행하여 서비스가 생성되었는지 확인할 수 있습니다.

![](/2022/Days/Images/Day54_Kubernetes6.png)

마지막으로, 배포를 확인하여 원하는 구성을 어디에 어떻게 유지하는지 확인할 수 있습니다.

![](/2022/Days/Images/Day54_Kubernetes7.png)

위는 몇 가지 알아두면 좋은 명령어를 사용했지만, `kubectl get all -n nginx`를 사용하여 하나의 YAML 파일로 배포한 모든 것을 볼 수도 있습니다.

![](/2022/Days/Images/Day54_Kubernetes8.png)

위 그림에서 replicaset가 있는 것을 볼 수 있는데, 배포에서 배포할 이미지의 레플리카 개수를 정의합니다. 처음에는 1로 설정되었지만, 애플리케이션을 빠르게 확장하려면 여러 가지 방법으로 확장할 수 있습니다.

터미널 내에서 텍스트 편집기를 열고 배포를 수정할 수 있는 `kubectl edit deployment  nginx-deployment -n nginx`를 사용하여 파일을 편집할 수 있습니다.

![](/2022/Days/Images/Day54_Kubernetes9.png)

터미널 내의 텍스트 편집기에서 위의 내용을 저장했을 때 문제가 없고 올바른 서식이 사용되었다면 네임스페이스에 추가로 배포된 것을 볼 수 있을 것입니다.

![](/2022/Days/Images/Day54_Kubernetes10.png)

또한 kubectl과 `kubectl scale deployment  nginx-deployment --replicas=10 -n nginx`를 사용하여 레플리카 수를 변경할 수 있습니다.

![](/2022/Days/Images/Day54_Kubernetes11.png)

두 방법 중 하나를 사용하려는 경우 이 방법을 사용하여 애플리케이션을 다시 1로 축소할 수 있습니다. 저는 편집 옵션을 사용했지만, 위의 스케일 명령을 사용할 수도 있습니다.

![](/2022/Days/Images/Day54_Kubernetes12.png)

여기서 사용 사례를 통해 매우 빠르게 스핀업 및 스핀다운할 수 있을 뿐만 아니라 애플리케이션을 빠르게 확장 및 축소할 수 있다는 것을 알 수 있기를 바랍니다. 이것이 웹 서버라면 부하가 많을 때는 확장하고 부하가 적을 때는 축소할 수 있습니다.

### 앱 노출하기

그렇다면 어떻게 웹 서버에 접속할 수 있을까요?

위에서 저희 서비스를 보면 사용 가능한 외부 IP가 없으므로 웹 브라우저를 열고 마술처럼 접속할 수는 없습니다. 접속을 위해 몇 가지 옵션이 있습니다.

**ClusterIP** - 표시되는 IP는 클러스터 내부 네트워크에 있는 클러스터IP입니다. 클러스터 내의 사물만 이 IP에 연결할 수 있습니다.

**NodePort** - NAT를 사용하여 클러스터에서 선택한 각 노드의 동일한 포트에 서비스를 노출합니다.

**로드 밸런서** - 현재 클라우드에 외부 로드 밸런서를 생성합니다. 저희는 Minikube를 사용하고 있지만, VirtualBox에서 했던 것과 같이 자체 Kubernetes 클러스터를 구축한 경우 이 기능을 제공하려면 metallb와 같은 로드밸런서를 클러스터에 배포해야 합니다.

**포트 포워드** - 로컬 호스트에서 내부 Kubernetes 클러스터 프로세스에 액세스하고 상호 작용할 수 있는 포트 포워드 기능도 있습니다. 이 옵션은 테스트 및 결함 발견에만 사용됩니다.

이제 선택할 수 있는 몇 가지 옵션이 생겼습니다. Minikube에는 본격적인 Kubernetes 클러스터와 비교했을 때 몇 가지 제한 사항이나 차이점이 있습니다.

다음 명령을 실행하여 로컬 워크스테이션을 사용하여 액세스를 포트 포워딩할 수 있습니다.

`kubectl port-forward deployment/nginx-deployment -n nginx 8090:80`

![](/2022/Days/Images/Day54_Kubernetes13.png)

위 명령을 실행하면 로컬 머신과 포트에 대한 포트 포워딩 역할을 하므로 이 터미널을 사용할 수 없게 됩니다.

![](/2022/Days/Images/Day54_Kubernetes14.png)

이제 Minikube를 통해 애플리케이션을 노출하는 방법을 구체적으로 살펴보겠습니다. Minikube를 사용하여 서비스에 연결하기 위한 URL을 생성할 수도 있습니다. [자세한 내용](https://minikube.sigs.k8s.io/docs/commands/service/)

먼저, `kubectl delete service nginx-service -n nginx`를 사용하여 서비스를 삭제합니다.

다음으로 `kubectl expose deployment nginx-deployment --name nginx-service --namespace nginx --port=80 --type=NodePort`를 사용하여 새 서비스를 생성합니다. 여기서 expose를 사용하고 유형을 NodePort로 변경한다는 점에 유의하세요.

![](/2022/Days/Images/Day54_Kubernetes15.png)

마지막으로 새 터미널에서 `minikube --profile='mc-demo' service nginx-service --url -n nginx`를 실행하여 서비스에 대한 터널을 생성합니다.

![](/2022/Days/Images/Day54_Kubernetes16.png)

브라우저 또는 제어를 열고 터미널에서 링크를 클릭합니다.

![](/2022/Days/Images/Day54_Kubernetes17.png)

### Helm

Helm은 애플리케이션을 배포할 수 있는 또 다른 방법입니다. "Kubernetes를 위한 패키지 관리자"로 알려진 Helm에 대한 자세한 내용은 [여기](https://helm.sh/)에서 확인할 수 있습니다.

Helm은 Kubernetes를 위한 패키지 매니저입니다. Helm은 Kubernetes에서 yum이나 apt에 해당하는 것으로 간주할 수 있습니다. Helm은 패키지 애플리케이션처럼 생각할 수 있는 차트를 배포하는데, 이는 미리 구성된 애플리케이션 리소스를 사용하기 쉬운 하나의 차트로 배포할 수 있는 청사진입니다. 그런 다음 다른 구성 세트로 차트의 다른 버전을 배포할 수 있습니다.

사용 가능한 모든 Helm 차트를 찾아볼 수 있는 사이트가 있으며 물론 직접 만들 수도 있습니다. 문서도 명확하고 간결하며 이 분야의 다른 모든 신조어들 사이에서 Helm이라는 용어를 처음 들었을 때처럼 어렵지 않습니다.

Helm을 시작하고 실행하거나 설치하는 것은 매우 간단합니다. 간단합니다. RaspberryPi arm64 장치를 포함한 거의 모든 배포판에 대한 바이너리와 다운로드 링크는 여기에서 찾을 수 있습니다.

또는 설치 스크립트를 사용할 수도 있는데, 이 경우 최신 버전의 Helm이 다운로드되어 설치된다는 이점이 있습니다.

```Shell
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3

chmod 700 get_helm.sh

./get_helm.sh
```

마지막으로, 애플리케이션 관리자를 위한 패키지 관리자, 맥용 homebrew, 윈도우용 chocolatey, Ubuntu/Debian용 apt, snap 및 pkg도 사용할 수 있습니다.

지금까지는 Helm이 클러스터에 다양한 테스트 애플리케이션을 다운로드하고 설치하는 데 가장 적합한 방법인 것 같습니다.

여기에 링크할 수 있는 좋은 리소스로는 Kubernetes 패키지를 찾고, 설치하고, 게시할 수 있는 리소스인 [ArtifactHUB](https://artifacthub.io/)를 들 수 있습니다. 또한 Helm 차트를 표시하는 UI인 [KubeApps](https://kubeapps.com/)에 대해서도 언급하겠습니다.

### Kubernetes에서 다룰 내용

아래에 언급된 내용 중 일부를 다루기 시작했지만, 내일 두 번째 클러스터 배포를 통해 더 많은 실습을 한 후 클러스터에 애플리케이션 배포를 시작할 수 있습니다.

- Kubernetes 아키텍처
- Kubectl 커맨드
- Kubernetes YAML
- Kubernetes Ingress
- Kubernetes Services
- Helm 패키지 관리자
- 영속성 스토리지
- stateful 앱

## 자료

사용하신 무료 리소스가 있으시면 리포지토리에 PR을 통해 여기에 추가해 주시면 기꺼이 포함시켜 드리겠습니다.

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

[Day 55](day55.md)에서 봐요!
