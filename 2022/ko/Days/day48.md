---
title: '#90DaysOfDevOps - Alternatives to Docker - Day 48'
published: false
description: 90DaysOfDevOps - Alternatives to Docker
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048807
---

## Docker의 대안

이 섹션의 맨 처음에 Docker를 사용할 것이라고 말씀드린 이유는 단순히 리소스 측면에서 매우 많고 커뮤니티가 매우 크기 때문이기도 하지만, 컨테이너를 대중화하기 위한 시도가 바로 여기에서 시작되었기 때문입니다. Docker의 역사와 그 탄생 과정을 살펴보는 것이 매우 유용하다고 생각합니다.

하지만 앞서 언급했듯이 Docker를 대체할 수 있는 다른 대안이 있습니다. Docker가 무엇이고 무엇을 다루었는지 생각해 보면 다음과 같습니다. 애플리케이션을 개발, 테스트, 배포 및 관리하기 위한 플랫폼입니다.

앞으로 실제로 사용할 수 있거나 앞으로 사용할 수 있는 Docker의 몇 가지 대안을 강조하고 싶습니다.

### Podman

Podman이란? Podman은 리눅스 시스템에서 OCI 컨테이너를 개발, 관리, 실행하기 위한 데몬이 없는 컨테이너 엔진입니다. 컨테이너는 루트로 실행하거나 루트리스 모드로 실행할 수 있습니다.

여기서는 Windows 관점에서 살펴볼 것이지만, Docker와 마찬가지로 Windows에서는 할 수 없는 기본 OS를 사용하기 때문에 가상화가 필요하지 않습니다.

Podman은 WSL2에서도 실행할 수 있지만, Docker Desktop의 경험만큼 매끄럽지는 않습니다. 컨테이너가 실행될 Linux VM에 연결할 수 있는 Windows 원격 클라이언트도 있습니다.

제가 사용하는 WSL2의 우분투는 20.04 릴리스입니다. 다음 단계에 따라 WSL 인스턴스에 Podman을 설치할 수 있습니다.

```Shell
echo "deb https://download.opensuse.org/repositories/devel:/kubic:/libcontainers:/stable/xUbuntu_20.04/ /" |
sudo tee /etc/apt/sources.list.d/devel:kubic:libcontainers:stable.list
```

GPG 키를 추가합니다.

```Shell
curl -L "https://download.opensuse.org/repositories/devel:/kubic:\
/libcontainers:/stable/xUbuntu_20.04/Release.key" | sudo apt-key add -
```

`sudo apt-get update && sudo apt-get upgrade`명령으로 시스템 업데이트 및 업그레이드를 실행합니다. 마지막으로`sudo apt install podman`을 사용하여 podman을 설치할 수 있습니다.

이제 docker에 사용하던 것과 동일한 명령을 많이 사용할 수 있습니다. 다만 멋진 docker 데스크톱 UI가 없다는 점에 유의하세요. 아래에서 볼 수 있듯이 `podman images`를 사용했고 설치 후 아무것도 없으므로 `podman pull ubuntu`를 사용하여 우분투 컨테이너 이미지를 끌어 내렸습니다.

![](/2022/Days/Images/Day48_Containers1.png)

그런 다음 `podman run -dit ubuntu`와 `podman ps`를 사용하여 우분투 이미지를 실행하여 실행 중인 이미지를 확인할 수 있습니다.

![](/2022/Days/Images/Day48_Containers2.png)

그런 다음 해당 컨테이너에 들어가기 위해 `podman attach dazzling_darwin`을 실행하면 컨테이너 이름이 달라질 수 있습니다.

![](/2022/Days/Images/Day48_Containers3.png)

docker에서 podman으로 이동하는 경우, 설정 파일에 `alias docker=podman`으로 변경하여 docker로 실행하는 모든 명령이 podman을 사용하도록 하는 것도 일반적입니다.

### LXC

LXC는 사용자가 다시 여러 개의 격리된 리눅스 컨테이너 환경을 생성할 수 있게 해주는 컨테이너화 엔진입니다. Docker와 달리, LXC는 별도의 시스템 파일과 네트워킹 기능을 갖춘 여러 리눅스 머신을 생성하기 위한 하이퍼바이저 역할을 합니다. docker보다 먼저 등장했다가 docker의 단점으로 인해 잠시 사용되었습니다.

LXC는 docker만큼 가볍고 쉽게 배포할 수 있습니다.

### Containerd

독립형 컨테이너 런타임. Containerd는 단순성과 견고성은 물론 이식성까지 제공합니다. 이전에는 Docker 컨테이너 서비스의 일부로 실행되는 도구로 사용되다가 Docker가 구성 요소를 독립형 구성 요소로 분리하기로 결정하기 전까지 Containerd가 사용되었습니다.

클라우드 네이티브 컴퓨팅 재단의 프로젝트로, Kubernetes, Prometheus, CoreDNS와 같은 인기 컨테이너 도구와 같은 부류에 속합니다.

### 기타 Docker 도구

Rancher와 VirtualBox와 관련된 도구와 옵션도 언급할 수 있지만 다음 기회에 더 자세히 다루겠습니다.

[**Gradle**](https://gradle.org/)

- 빌드 스캔을 통해 팀은 공동으로 스크립트를 디버깅하고 모든 빌드의 이력을 추적할 수 있습니다.
- 실행 옵션을 통해 팀은 변경 사항이 입력될 때마다 작업이 자동으로 실행되도록 지속적으로 빌드할 수 있습니다.
- 사용자 지정 리포지토리 레이아웃을 통해 팀은 모든 파일 디렉토리 구조를 아티팩트 리포지토리로 취급할 수 있습니다.

[**Packer**](https://packer.io/)

- 여러 머신 이미지를 병렬로 생성하여 개발자의 시간을 절약하고 효율성을 높일 수 있습니다.
- 패커의 디버거를 사용하여 빌드를 쉽게 디버깅할 수 있어 실패를 검사하고 빌드를 다시 시작하기 전에 솔루션을 시험해 볼 수 있습니다.
- 플러그인을 통해 다양한 플랫폼을 지원하므로 팀이 빌드를 커스터마이징할 수 있습니다.

[**Logspout**](https://github.com/gliderlabs/logspout)

- 로깅 도구 - 이 도구의 사용자 지정 기능을 통해 팀은 동일한 로그를 여러 대상에 전송할 수 있습니다.
- 이 도구는 Docker 소켓에 액세스하기만 하면 되기 때문에 팀에서 파일을 쉽게 관리할 수 있습니다.
- 완전히 오픈 소스이며 배포가 쉽습니다.

[**Logstash**](https://www.elastic.co/products/logstash)

- Logstash의 플러그형 프레임워크를 사용해 파이프라인을 사용자 정의하세요.
- 분석을 위해 데이터를 쉽게 구문 분석하고 변환하여 비즈니스 가치를 제공할 수 있습니다.
- Logstash의 다양한 출력을 통해 원하는 곳으로 데이터를 라우팅할 수 있습니다.

[**Portainer**](https://www.portainer.io/)

- 미리 만들어진 템플릿을 활용하거나 직접 템플릿을 만들어 애플리케이션을 배포하세요.
- 팀을 생성하고 팀원에게 역할과 권한을 할당하세요.
- 도구의 대시보드를 사용하여 각 환경에서 무엇이 실행되고 있는지 파악하세요.

## 자료

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)
- [Blog on getting started building a docker image](https://stackify.com/docker-build-a-beginners-guide-to-building-docker-images/)
- [Docker documentation for building an image](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- [YAML Tutorial: Everything You Need to Get Started in Minute](https://www.cloudbees.com/blog/yaml-tutorial-everything-you-need-get-started)
- [Podman | Daemonless Docker | Getting Started with Podman](https://www.youtube.com/watch?v=Za2BqzeZjBk)
- [LXC - Guide to building an LXC Lab](https://www.youtube.com/watch?v=cqOtksmsxfg)

[Day 49](day49.md)에서 봐요!
