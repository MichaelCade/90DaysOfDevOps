---
title: '#90DaysOfDevOps - Setting up a multinode Kubernetes Cluster - Day 52'
published: false
description: 90DaysOfDevOps - Setting up a multinode Kubernetes Cluster
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049050
---

## 멀티노드 Kubernetes 클러스터 설정하기

이 제목을 "Vagrant로 멀티노드 Kubernetes 클러스터 설정하기"로 하고 싶었지만, 너무 길 것 같았습니다!

어제 세션에서는 멋진 프로젝트를 사용하여 첫 번째 Kubernetes 클러스터를 배포하고 Kubernetes를 사용할 때 접하게 될 가장 중요한 CLI 도구(kubectl)를 조금 실습해 보았습니다.

여기서는 VirtualBox를 기본으로 사용하지만, 지난번 리눅스 섹션에서 Vagrant에 대해 이야기할 때 언급했듯이 지원되는 모든 하이퍼바이저 또는 가상화 도구를 사용할 수 있습니다. 리눅스 섹션에서 우분투 머신을 배포한 것은 [Day 14](day14.md)였습니다.

### Vagrant에 대한 간략한 요약

Vagrant는 가상 머신의 라이프사이클을 관리하는 CLI 유틸리티입니다. vSphere, Hyper-v, Virtual Box, Docker 등 다양한 플랫폼에서 가상 머신을 스핀업 및 스핀다운하는 데 Vagrant를 사용할 수 있습니다. 다른 공급업체도 있지만 여기서는 Virtual Box를 사용하고 있으므로 계속 사용하겠습니다.

이 [블로그 및 리포지토리](https://devopscube.com/kubernetes-cluster-vagrant/)를 기준으로 하여 구성을 안내해 드리겠습니다. 하지만 Kubernetes 클러스터를 처음 배포하는 경우라면 수동으로 이 작업을 수행하는 방법도 살펴보고 최소한 어떤 모습인지 알고 계실 것을 권해드리고 싶습니다. Kubernetes는 릴리스될 때마다 더욱 효율적으로 개선되고 있다고 말씀드리고 싶습니다. 저는 이것을 VMware와 ESX 시절에 비유하자면, ESX 서버 3대를 배포하는 데 적어도 하루는 필요했지만, 지금은 한 시간 안에 이를 실행할 수 있습니다. 저희는 Kubernetes와 관련해서는 그 방향으로 나아가고 있습니다.

### Kubernetes 랩 환경

환경을 구축하는 데 사용할 vagrantfile을 [Kubernetes 폴더](/2022/Days/Kubernetes)에 업로드했습니다. 이 파일을 잡고 터미널에서 이 디렉토리로 이동합니다. 저는 다시 Windows를 사용하므로 PowerShell을 사용하여 vagrant로 워크스테이션 명령을 수행하겠습니다. vagrant가 없는 경우 어제 Minikube 및 기타 도구를 설치할 때 다룬 arkade를 사용할 수 있습니다. 간단한 명령어인 `arkade get vagrant`를 실행하면 최신 버전의 vagrant를 다운로드하여 설치할 수 있습니다.

디렉토리에 들어가면 `vagrant up`을 실행하고 모든 것이 올바르게 구성되었다면 터미널에 다음과 같은 킥오프가 표시됩니다.

![](/2022/Days/Images/Day52_Kubernetes1.png)

터미널에서 몇 가지 단계가 진행되는 것을 볼 수 있지만, 그동안 우리가 여기서 무엇을 빌드하고 있는지 살펴봅시다.

![](/2022/Days/Images/Day52_Kubernetes2.png)

위에서 보면 3개의 가상 머신을 빌드하고 컨트롤 플레인 노드와 두 개의 워커 노드가 있다는 것을 알 수 있습니다. [Day 49](day49.md)로 돌아가면 이미지에서 볼 수 있는 이러한 영역에 대한 설명이 더 있습니다.

또한 이미지에서는 클러스터 외부에서 kubectl 액세스가 발생하여 해당 kube apiserver에 도달하는 것으로 표시되어 있지만, 실제로는 vagrant 프로비저닝의 일부로 각 노드 내에서 클러스터에 액세스할 수 있도록 각 노드에 kubectl을 배포하고 있습니다.

이 실습을 구축하는 과정은 설정에 따라 5분에서 30분 정도 걸릴 수 있습니다.

곧 스크립트에 대해서도 다룰 예정이지만, 배포의 일부로 3개의 스크립트를 호출하는 vagrant 파일을 보면 클러스터가 실제로 생성되는 곳이라는 것을 알 수 있습니다. Vagrant boxes를 사용하여 가상 머신과 OS 설치를 배포하는 것이 얼마나 쉬운지 살펴보았지만, 배포 프로세스의 일부로 셸 스크립트를 실행할 수 있는 기능이 있다는 것은 이러한 실습 빌드 아웃 자동화와 관련하여 매우 흥미로운 부분입니다.

완료되면 터미널에서 노드 중 하나에 `vagrant ssh master`로 접속하면 액세스할 수 있으며, 기본 사용자 이름과 비밀번호는 `vagrant/vagrant`입니다.

원하는 경우 `vagrant ssh node01` 및 `vagrant ssh node02`를 사용하여 작업자 노드에 액세스할 수도 있습니다.

![](/2022/Days/Images/Day52_Kubernetes3.png)

이제 새 클러스터의 위 노드 중 하나에서 `kubectl get nodes`를 실행하여 3노드 클러스터와 그 상태를 확인할 수 있습니다.

![](/2022/Days/Images/Day52_Kubernetes4.png)

이 시점에서, 컨트롤 플레인 노드 1개와 워커 노드 2개로 구성된 3노드 클러스터가 실행되고 있습니다.

### Vagrant 파일 및 셸 스크립트 연습

vagrantfile을 살펴보면 여러 작업자 노드, VirtualBox 내의 브리지 네트워크에 대한 네트워킹 IP 주소, 그리고 일부 이름 지정을 정의하고 있음을 알 수 있습니다. 또한 특정 호스트에서 실행하려는 일부 스크립트를 호출하고 있음을 알 수 있습니다.

```shell
NUM_WORKER_NODES=2
IP_NW="10.0.0."
IP_START=10

Vagrant.configure("2") do |config|
    config.vm.provision "shell", inline: <<-SHELL
        apt-get update -y
        echo "$IP_NW$((IP_START))  master-node" >> /etc/hosts
        echo "$IP_NW$((IP_START+1))  worker-node01" >> /etc/hosts
        echo "$IP_NW$((IP_START+2))  worker-node02" >> /etc/hosts
    SHELL
    config.vm.box = "bento/ubuntu-21.10"
    config.vm.box_check_update = true

    config.vm.define "master" do |master|
      master.vm.hostname = "master-node"
      master.vm.network "private_network", ip: IP_NW + "#{IP_START}"
      master.vm.provider "virtualbox" do |vb|
          vb.memory = 4048
          vb.cpus = 2
          vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
      end
      master.vm.provision "shell", path: "scripts/common.sh"
      master.vm.provision "shell", path: "scripts/master.sh"
    end

    (1..NUM_WORKER_NODES).each do |i|
      config.vm.define "node0#{i}" do |node|
        node.vm.hostname = "worker-node0#{i}"
        node.vm.network "private_network", ip: IP_NW + "#{IP_START + i}"
        node.vm.provider "virtualbox" do |vb|
            vb.memory = 2048
            vb.cpus = 1
            vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
        end
        node.vm.provision "shell", path: "scripts/common.sh"
        node.vm.provision "shell", path: "scripts/node.sh"
      end
    end
  end
```

실행 중인 스크립트를 분석해 보겠습니다. 특정 노드에서 실행할 세 개의 스크립트가 위의 VAGRANTFILE에 나열되어 있습니다.

`master.vm.provision "shell", path: "scripts/common.sh"`

위의 스크립트는 노드를 준비하는 데 초점을 맞출 것이며, 3개의 노드 모두에서 실행될 것이며, 기존의 모든 Docker 구성 요소를 제거하고 Docker와 ContainerD는 물론 kubeadm, kubelet 및 kubectl을 다시 설치합니다. 이 스크립트는 또한 시스템의 기존 소프트웨어 패키지도 업데이트합니다.

`master.vm.provision "shell", path: "scripts/master.sh"`

master.sh 스크립트는 컨트롤 플레인 노드에서만 실행되며, 이 스크립트는 kubeadm 커맨드를 사용하여 Kubernetes 클러스터를 생성합니다. 또한 이 클러스터에 대한 액세스를 위한 구성 컨텍스트도 준비할 것이며, 이는 다음에 다룰 것입니다.

`node.vm.provision "shell", path: "scripts/node.sh"`

이것은 단순히 마스터가 생성한 구성을 가져와서 우리 노드를 Kubernetes 클러스터에 추가하는 것이며, 이 추가 프로세스는 다시 kubeadm과 config 폴더에서 찾을 수 있는 다른 스크립트를 사용합니다.

### Kubernetes 클러스터에 액세스하기

이제 두 개의 클러스터가 배포되었습니다. 이전 섹션에서 배포한 Minikube 클러스터와 방금 VirtualBox에 배포한 새로운 3노드 클러스터가 있습니다.

또한 vagrant를 실행한 머신에서도 액세스할 수 있는 구성 파일에는 워크스테이션에서 클러스터에 액세스하는 방법이 포함되어 있습니다.

이를 보여드리기 전에 컨텍스트에 대해 말씀드리겠습니다.

![](/2022/Days/Images/Day52_Kubernetes5.png)

컨텍스트가 중요하며, 데스크톱이나 노트북에서 Kubernetes 클러스터에 액세스할 수 있는 기능이 필요합니다. 다양한 옵션이 존재하며 사람들은 각기 다른 운영 체제를 일상적으로 사용합니다.

기본적으로, Kubernetes CLI 클라이언트(kubectl)는 엔드포인트 및 자격 증명과 같은 Kubernetes 클러스터 세부 정보를 저장하기 위해 C:\Users\username.kube\config를 사용합니다. 클러스터를 배포한 경우 해당 위치에서 이 파일을 볼 수 있습니다. 하지만 지금까지 마스터 노드에서 SSH 또는 다른 방법을 통해 모든 kubectl 명령을 실행했다면 이 포스팅이 워크스테이션과 연결할 수 있는 방법을 이해하는 데 도움이 되길 바랍니다.

그런 다음 클러스터에서 kubeconfig 파일을 가져오거나 배포된 구성 파일에서 가져올 수도 있고, SCP를 통해 이 파일의 내용을 가져오거나 마스터 노드에 콘솔 세션을 열고 로컬 윈도우 머신에 복사할 수도 있습니다.

![](/2022/Days/Images/Day52_Kubernetes6.png)

그런 다음 해당 구성 파일의 복사본을 가져와서 `$HOME/.kube/config` 위치로 이동합니다.

![](/2022/Days/Images/Day52_Kubernetes7.png)

이제 로컬 워크스테이션에서 `kubectl cluster-info`와 `kubectl get nodes`를 실행하여 클러스터에 액세스할 수 있는지 확인할 수 있습니다.

![](/2022/Days/Images/Day52_Kubernetes8.png)

이렇게 하면 윈도우 머신에서 연결 및 제어가 가능할 뿐만 아니라 윈도우 머신에서 특정 서비스에 액세스하기 위해 포트 포워딩을 수행할 수 있습니다.

워크스테이션에서 여러 클러스터를 관리하는 방법에 관심이 있으시다면 [여기](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-6)에 더 자세한 안내가 있습니다.

이 목록은 제가 배포 중인 다양한 Kubernetes 클러스터에 대해 수행한 워크스루 블로그입니다.

- [Kubernetes playground – How to choose your platform](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-1)
- [Kubernetes playground – Setting up your cluster](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-2)
- [Getting started with Amazon Elastic Kubernetes Service (Amazon EKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-amazon-elastic-kubernetes-service-amazon-eks)
- [Getting started with Microsoft Azure Kubernetes Service (AKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-azure-kubernetes-service-aks)
- [Getting Started with Microsoft AKS – Azure PowerShell Edition](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-aks-azure-powershell-edition)
- [Getting started with Google Kubernetes Service (GKE)](https://vzilla.co.uk/vzilla-blog/getting-started-with-google-kubernetes-service-gke)
- [Kubernetes, How to – AWS Bottlerocket + Amazon EKS](https://vzilla.co.uk/vzilla-blog/kubernetes-how-to-aws-bottlerocket-amazon-eks)
- [Getting started with CIVO Cloud](https://vzilla.co.uk/vzilla-blog/getting-started-with-civo-cloud)
- [Minikube - Kubernetes Demo Environment For Everyone](https://vzilla.co.uk/vzilla-blog/project_pace-kasten-k10-demo-environment-for-everyone)

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

[Day 53](day53.md)에서 봐요!
