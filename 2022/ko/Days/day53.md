---
title: '#90DaysOfDevOps - Rancher Overview - Hands On - Day 53'
published: false
description: 90DaysOfDevOps - Rancher Overview - Hands On
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048742
---

## Rancher 개요 - 핸즈온

이 섹션에서는 Rancher에 대해 살펴볼 것인데, 지금까지는 클러스터 관리에 대한 좋은 가시성을 운영팀에 제공하는 몇 가지 좋은 UI와 멀티클러스터 관리 도구가 있지만, 지금까지 한 모든 작업은 cli와 kubectl을 사용했습니다.

Rancher는 [사이트](https://rancher.com/)에 따르면

> Rancher는 컨테이너를 도입하는 팀을 위한 완벽한 소프트웨어 스택입니다. 이 스택은 모든 인프라에서 여러 개의 Kubernetes 클러스터를 관리할 때 발생하는 운영 및 보안 문제를 해결하는 동시에 데브옵스 팀에 컨테이너화된 워크로드를 실행하기 위한 통합 도구를 제공합니다.

Rancher를 사용하면 거의 모든 위치에서 프로덕션급 Kubernetes 클러스터를 배포할 수 있으며 중앙 집중식 인증, 액세스 제어 및 통합 가시성을 제공합니다. 이전 섹션에서 Kubernetes와 관련하여 거의 압도적인 선택의 폭이 있으며, 어디에서 실행해야 하는지 또는 실행할 수 있는지에 대해 언급했지만, Rancher를 사용하면 어디에 있든 상관없습니다.

### Rancher 배포

가장 먼저 해야 할 일은 로컬 워크스테이션에 Rancher를 배포하는 것입니다. 이 단계를 진행하기 위해 선택할 수 있는 몇 가지 방법과 위치가 있는데, 저는 로컬 워크스테이션을 사용하고 Rancher를 docker 컨테이너로 실행하고 싶습니다. 아래 명령을 실행하면 컨테이너 이미지를 가져온 다음 Rancher UI에 액세스할 수 있습니다.

다른 Rancher 배포 방법은 [Rancher 빠른 시작 가이드](https://rancher.com/docs/rancher/v2.6/en/quick-start-guide/deployment/)에서 확인할 수 있습니다.

`sudo docker run -d --restart=unless-stopped -p 80:80 -p 443:443 --privileged rancher/rancher`

docker 데스크톱에서 볼 수 있듯이 실행 중인 Rancher 컨테이너가 있습니다.

![](/2022/Days/Images/Day53_Kubernetes1.png)

### Rancher UI 액세스

위의 컨테이너가 실행 중이면 웹 페이지를 통해 컨테이너로 이동할 수 있어야 합니다. `https://localhost`를 입력하면 아래와 같이 로그인 페이지가 나타납니다.

![](/2022/Days/Images/Day53_Kubernetes2.png)

아래 안내에 따라 필요한 비밀번호를 입력합니다. 저는 Windows를 사용하고, grep 명령이 필요하기 때문에 Windows용 bash를 사용하기로 했습니다.

![](/2022/Days/Images/Day53_Kubernetes3.png)

이제 위의 비밀번호를 사용하여 로그인하면 다음 페이지에서 새 비밀번호를 정의할 수 있습니다.

![](/2022/Days/Images/Day53_Kubernetes4.png)

위의 작업을 완료하면 로그인이 완료되고 시작 화면을 볼 수 있습니다. Rancher 배포의 일부로 로컬 K3 클러스터가 프로비저닝된 것도 볼 수 있습니다.

![](/2022/Days/Images/Day53_Kubernetes5.png)

### Rancher에 대한 간략한 둘러보기

가장 먼저 살펴볼 것은 로컬로 배포된 K3S 클러스터입니다. 아래에서 클러스터 내부에서 어떤 일이 일어나고 있는지 잘 볼 수 있습니다. 이것은 기본 배포이며 아직 이 클러스터에 아무것도 배포하지 않았습니다. 1개의 노드로 구성되어 있고 5개의 배포가 있는 것을 볼 수 있습니다. 그리고 pod, 코어, 메모리에 대한 몇 가지 통계가 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day53_Kubernetes6.png)

왼쪽 메뉴에는 앱 및 마켓플레이스 탭도 있는데, 이 탭을 통해 클러스터에서 실행할 애플리케이션을 선택할 수 있습니다. 앞서 언급했듯이 Rancher는 여러 개의 다른 클러스터를 실행하거나 관리할 수 있는 기능을 제공합니다. 마켓플레이스를 사용하면 애플리케이션을 매우 쉽게 배포할 수 있습니다.

![](/2022/Days/Images/Day53_Kubernetes7.png)

또 한 가지 언급할 것은 오른쪽 상단에 있는 Rancher에서 관리 중인 클러스터에 액세스해야 하는 경우 선택한 클러스터에 대한 kubectl 셸을 열 수 있다는 것입니다.

![](/2022/Days/Images/Day53_Kubernetes8.png)

### 새 클러스터 생성

지난 두 세션에 걸쳐 로컬에서 Minikube 클러스터를 생성하고 가상박스와 함께 Vagrant를 사용하여 3노드 Kubernetes 클러스터를 생성했으며, Rancher를 사용하여 클러스터를 생성할 수도 있습니다. [Rancher 폴더](/2022/Days/Kubernetes/Rancher)에는 동일한 3개의 노드를 구축할 수 있는 추가 vagrant 파일이 있지만 Kubernetes 클러스터를 생성하는 단계가 없습니다(Rancher가 이 작업을 대신 수행하기를 원합니다).

그러나 각 노드에서 `common.sh` 스크립트가 계속 실행되는 것을 볼 수 있도록 docker가 설치되고 OS가 업데이트되기를 원합니다. 이것은 또한 Kubeadm, Kubectl 등을 설치합니다. 그러나 노드를 클러스터로 생성하고 조인하기 위한 Kubeadm 명령은 실행되지 않습니다.

vagrant 폴더 위치로 이동하여 `vagrant up`을 실행하기만 하면 가상 박스에서 3개의 가상 머신을 생성하는 프로세스가 시작됩니다.

![](/2022/Days/Images/Day53_Kubernetes9.png)

이제 노드 또는 VM이 제자리에 배치되고 준비되었으므로 Rancher를 사용하여 새로운 Kubernetes 클러스터를 생성할 수 있습니다. 클러스터를 생성하는 첫 번째 화면에서는 클러스터가 어디에 있는지, 즉 퍼블릭 클라우드 관리형 Kubernetes 서비스를 사용 중인지, vSphere 또는 다른 것을 사용 중인지에 대한 몇 가지 옵션을 제공합니다.

![](/2022/Days/Images/Day53_Kubernetes10.png)

통합 플랫폼 중 하나를 사용하지 않으므로 "custom"을 선택하겠습니다. 시작 페이지는 클러스터 이름을 정의하는 곳입니다(아래에 로컬이라고 되어 있지만 로컬을 사용할 수 없습니다. 저희 클러스터는 vagrant라고 합니다.) 여기에서 Kubernetes 버전, 네트워크 공급자 및 기타 구성 옵션을 정의하여 Kubernetes 클러스터를 시작하고 실행할 수 있습니다.

![](/2022/Days/Images/Day53_Kubernetes11.png)

다음 페이지에서는 활성화할 적절한 서비스와 함께 각 노드에서 실행해야 하는 등록 코드(etcd, 컨트롤 플레인 및 워커)를 제공합니다. 마스터 노드의 경우, etcd와 컨트롤 플레인이 필요하므로 명령은 아래와 같습니다.

![](/2022/Days/Images/Day53_Kubernetes12.png)

```
sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run rancher/rancher-agent:v2.6.3 --server https://10. 0.0.1 --token mpq8cbjjwrj88z4xmf7blqxcfmwdsmq92bmwjpphdkklfckk5hfwc2 --ca-checksum a81944423cbfeeb92be0784edebba1af799735ebc30ba8cbe5cc5f996094f30b --etcd --controlplane
```

네트워킹이 올바르게 구성되었다면, 이제 첫 번째 마스터 노드가 등록되고 클러스터가 생성되고 있음을 나타내는 Rancher 대시보드에 다음과 같이 빠르게 표시되어야 합니다.

![](/2022/Days/Images/Day53_Kubernetes13.png)

그런 다음 다음 명령으로 각 워커 노드에 대한 등록 프로세스를 반복하면 얼마 후 마켓플레이스를 활용하여 애플리케이션을 배포할 수 있는 클러스터를 실행할 수 있게 됩니다.

```
sudo docker run -d --privileged --restart=unless-stopped --net=host -v /etc/kubernetes:/etc/kubernetes -v /var/run:/var/run rancher/rancher-agent:v2.6.3 --server https://10. 0.0.1 --token mpq8cbjjwrj88z4xmf7blqxcfmwdsmq92bmwjpphdkklfckk5hfwc2 --ca-checksum a81944423cbfeeb92be0784edebba1af799735ebc30ba8cbe5cc5f996094f30b --worker
```

![](/2022/Days/Images/Day53_Kubernetes14.png)

지난 세 세션 동안, 우리는 몇 가지 다른 방법으로 Kubernetes 클러스터를 시작하고 실행하는 방법을 사용했으며, 남은 날에는 플랫폼에서 가장 중요한 애플리케이션 측면을 살펴볼 것입니다. 서비스 프로비저닝과 Kubernetes에서 서비스를 프로비저닝하고 사용할 수 있는 방법에 대해 살펴보겠습니다.

부트스트랩 Rancher 노드에 대한 요구사항에 따라 해당 VM에 4GB 램이 있어야 하며 그렇지 않으면 크래시 루프가 발생한다고 들었는데, 이후 워커 노드에 2GB가 있는 것으로 업데이트했습니다.

### Kubernetes 시리즈에서 다룰 내용

아래에 언급된 내용 중 일부를 다루기 시작했지만, 내일 두 번째 클러스터 배포를 통해 더 많은 실습을 한 다음 클러스터에 애플리케이션 배포를 시작할 수 있습니다.

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

[Day 54](day54.md)에서 봐요!
