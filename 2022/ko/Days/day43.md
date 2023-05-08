---
title: '#90DaysOfDevOps - What is Docker & Getting installed - Day 43'
published: false
description: 90DaysOfDevOps - What is Docker & Getting installed
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048739
---

## Docker를 알아보고 설치하기

이전 포스팅에서 docker에 대해 한 번쯤은 언급했는데, 그 이유는 docker가 컨테이너가 오래전부터 사용되어 왔음에도 불구하고 대중화될 수 있었던 혁신적인 기술이기 때문입니다.

여기서는 docker를 사용하고 설명할 예정이지만, 벤더 종속의 위험을 피하면서 혁신을 장려하는 산업 표준 조직인 [Open Container Initiative(OCI)](https://www.opencontainers.org/)에 대해서도 언급해야 합니다. OCI 덕분에 컨테이너 툴체인을 선택할 때 Docker, [CRI-O](https://cri-o.io/), [Podman](http://podman.io/), [LXC](https://linuxcontainers.org/) 등을 선택할 수 있게 되었습니다.

docker는 컨테이너를 빌드, 실행 및 관리하기 위한 소프트웨어 프레임워크입니다. "docker"라는 용어는 도구(명령어 및 데몬) 또는 Dockerfile 파일 형식을 지칭할 수 있습니다.

여기서는 (교육 및 학습용으로) 무료인 Docker Personal을 사용하겠습니다. 여기에는 컨테이너와 도구에 대한 지식의 기초를 다지기 위해 다루어야 할 모든 필수 사항이 포함됩니다.

우리가 사용할 몇 가지 "docker" 도구와 그 용도를 세분화해 볼 가치가 있을 것입니다. docker라는 용어는 개발자와 관리자가 애플리케이션을 개발, 배포 및 실행하기 위한 플랫폼인 docker 프로젝트 전반을 지칭할 수 있습니다. 또한 이미지와 컨테이너를 관리하는 호스트에서 실행되는 docker 데몬 프로세스를 지칭할 수도 있으며, Docker Engine이라고도 합니다.

### Docker Engine

Docker Engine은 애플리케이션을 빌드하고 컨테이너화하기 위한 오픈소스 컨테이너화 기술입니다. Docker Engine은 클라이언트-서버 애플리케이션으로 작동합니다:

- 장기 실행 데몬 프로세스인 dockerd가 있는 서버.
- API는 프로그램이 Docker 데몬과 대화하고 지시하는 데 사용할 수 있는 인터페이스를 지정합니다.
- 커맨드라인 인터페이스(CLI) 클라이언트 docker입니다.

위의 내용은 공식 docker 문서와 특정 [Docker Engine 개요](https://docs.docker.com/engine/)에서 발췌한 것입니다.

### Docker Desktop

저희는 Windows와 macOS 시스템 모두를 위한 Docker Desktop을 제공합니다. 설치하기 쉬운 경량 docker 개발 환경입니다. 호스트 운영 체제의 가상화 기능을 활용하는 네이티브 OS 애플리케이션입니다.

Windows 또는 macOS에서 docker화된 애플리케이션을 빌드, 디버그, 테스트, 패키징 및 출시하려는 경우 가장 적합한 솔루션입니다.

Windows에서는 WSL2와 Microsoft Hyper-V도 활용할 수 있습니다. WSL2의 몇 가지 이점은 차근차근 살펴보도록 하겠습니다.

호스트 운영 체제의 하이퍼바이저 기능과 통합되어 있기 때문에 docker는 리눅스 운영 체제에서 컨테이너를 실행할 수 있는 기능을 제공합니다.

### Docker Compose

Docker Compose는 여러 컨테이너에서 더 복잡한 앱을 실행할 수 있는 도구입니다. 단일 파일과 명령을 사용하여 애플리케이션을 스핀업할 수 있다는 이점이 있습니다.

### DockerHub

docker와 그 구성 요소로 작업하기 위한 중앙 집중식 리소스입니다. 가장 일반적으로는 docker 이미지를 호스팅하는 레지스트리로 알려져 있습니다. 그러나 여기에는 부분적으로 자동화와 함께 사용하거나 보안 검사뿐만 아니라 GitHub에 통합할 수 있는 많은 추가 서비스가 있습니다.

### Dockerfile

Dockerfile은 일반적으로 docker 이미지를 빌드하기 위해 수동으로 실행하는 명령이 포함된 텍스트 파일입니다. docker는 Dockerfile에 있는 지침을 읽어 이미지를 자동으로 빌드할 수 있습니다.

## Docker Desktop 설치

[Docker 문서](https://docs.docker.com/engine/install/)는 매우 훌륭하며, 이제 막 docker에 입문하는 분이라면 꼭 한 번 읽어보시기 바랍니다. 여기서는 WSL2가 설치된 Windows에서 Docker Desktop을 사용하겠습니다. 여기서 사용하는 머신에 이미 설치를 완료했습니다.

![](/2022/Days/Images/Day43_Containers1.png)

설치를 진행하기 전에 시스템 요구 사항을 참고하시기 바라며, [윈도우에 Docker Desktop 설치하기](https://docs.docker.com/desktop/windows/install/)를 참고하시고, M1 기반 CPU 아키텍처를 포함한 맥OS를 사용하시는 경우 [맥OS에 docker Desktop 설치하기](https://docs.docker.com/desktop/mac/install/)도 참고하시면 됩니다.

다른 윈도우 머신에서 윈도우용 docker Desktop 설치를 실행하고 그 과정을 아래에 기록해 보겠습니다.

### Windows

- 장치의 운영 체제로 Windows를 선택합니다.
  <img src = ../../Days/Images/Day43_operatingSystem.png>

- 인스톨러를 저장할 폴더로 이동하여 저장합니다.

- 인스톨러를 실행하고 몇 초간 기다린 후 WSL에 대한 액세스 권한을 부여합니다.
  <img src = ../../Days/Images/Day43_EnableWSL.png>

- 확인을 클릭하면 설치가 시작됩니다.
  <img src = ../../Days/Images/Day43_install.png>

- Docker Desktop이 장치에 성공적으로 설치되었습니다. 이제 터미널에서 "docker" 명령을 실행하여 설치가 성공적으로 완료되었는지 확인할 수 있습니다.
  <img src = ../../Days/Images/Day43_check.png>

## 자료

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)

[Day 44](day44.md)에서 봐요!
