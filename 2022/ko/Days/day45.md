---
title: '#90DaysOfDevOps - The anatomy of a Docker Image - Day 45'
published: false
description: 90DaysOfDevOps - The anatomy of a Docker Image
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048777
---

## Docker 이미지의 구조

지난 세션에서는 DockerHub와 결합된 Docker Desktop을 사용하여 몇 가지 검증된 이미지를 배포하고 실행하는 방법에 대한 몇 가지 기본 사항을 다루었습니다. 이미지가 무엇인지에 대한 요약을 계속 언급하면 잊어버리지 않을 것입니다.

Docker 이미지는 Docker 플랫폼에서 실행할 수 있는 컨테이너를 만들기 위한 일련의 지침이 포함된 읽기 전용 템플릿입니다. 애플리케이션과 사전 구성된 서버 환경을 패키징하여 개인적으로 사용하거나 다른 Docker 사용자와 공개적으로 공유할 수 있는 편리한 방법을 제공합니다. 또한 Docker 이미지는 Docker를 처음 사용하는 모든 사용자에게 시작점이 됩니다.

자체 Docker 이미지를 만들려면 어떻게 해야 하나요? 이를 위해서는 Dockerfile을 생성해야 합니다. 우분투 컨테이너 이미지를 가져와서 소프트웨어를 추가하면 원하는 소프트웨어가 포함된 컨테이너 이미지를 갖게 되고 모든 것이 좋지만, 컨테이너가 종료되거나 폐기되면 모든 소프트웨어 업데이트와 설치가 사라져 반복할 수 있는 버전이 없습니다. 따라서 컨테이너를 실행할 때마다 동일한 소프트웨어 세트가 설치된 여러 환경에 걸쳐 이미지를 전송하는 데는 도움이 되지 않습니다.

### Dockerfile이란?

Dockerfile은 일반적으로 docker 이미지를 빌드하기 위해 수동으로 실행하는 명령이 포함된 텍스트 파일입니다. docker는 Dockerfile에 있는 지침을 읽어 이미지를 자동으로 빌드할 수 있습니다.

docker 이미지를 구성하는 각 파일을 레이어라고 하며, 이러한 레이어는 단계적으로 서로 위에 빌드되는 일련의 이미지를 형성합니다. 각 레이어는 바로 아래 레이어에 종속됩니다. 레이어의 순서는 docker 이미지의 라이프사이클 관리 효율성의 핵심입니다.

가장 자주 변경되는 레이어는 가능한 한 스택에서 가장 높은 곳에 구성해야 하는데, 이는 이미지의 레이어를 변경하면 Docker가 해당 레이어뿐만 아니라 해당 레이어에서 빌드된 모든 레이어를 다시 빌드하기 때문입니다. 따라서 맨 위에 있는 레이어를 변경하면 전체 이미지를 다시 빌드하는 데 필요한 작업량이 가장 적습니다.

docker가 이미지로부터 컨테이너를 실행할 때마다 (어제 실행한 것처럼), 컨테이너 레이어라고 하는 쓰기 가능한 레이어가 추가됩니다. 이 레이어는 컨테이너가 실행되는 동안 모든 변경 사항을 저장합니다. 이 레이어는 실제 운영 중인 컨테이너와 소스 이미지 자체 사이의 유일한 차이점입니다. 같은 이미지에 기반한 동일한 컨테이너들이 상태를 유지하면서 액세스를 공유할 수 있습니다.

예제로 돌아가서, 어제 우분투 이미지를 사용했습니다. 동일한 명령을 여러 번 실행하여 첫 번째 컨테이너에는 pinta를 설치하고 두 번째 컨테이너에는 두 개의 다른 애플리케이션, 다른 목적, 다른 크기 등을 가진 figlet을 설치할 수 있습니다. 배포한 각 컨테이너는 동일한 이미지를 공유하지만, 동일한 상태는 아니며, 컨테이너를 제거하면 해당 상태는 사라집니다.

![](/2022/Days/Images/Day45_Containers1.png)

위의 예제에서 우분투 이미지뿐만 아니라 DockerHub 및 기타 서드파티 리포지토리에서 사용할 수 있는 다른 많은 기성 컨테이너 이미지도 마찬가지입니다. 이러한 이미지를 일반적으로 부모 이미지라고 합니다. 이 이미지는 다른 모든 레이어가 구축되는 기반이 되며 컨테이너 환경의 기본 구성 요소를 제공합니다.

개별 레이어 파일 세트와 함께 Docker 이미지에는 manifest라는 추가 파일도 포함됩니다. 이 파일은 기본적으로 이미지에 대한 JSON 형식의 설명이며 이미지 태그, 전자 서명, 다양한 유형의 호스트 플랫폼에 맞게 컨테이너를 구성하는 방법에 대한 세부 정보와 같은 정보로 구성됩니다.

![](/2022/Days/Images/Day45_Containers2.png)

### Docker 이미지를 생성하는 방법

docker 이미지를 생성하는 방법에는 두 가지가 있습니다. 어제 시작한 프로세스로 즉석에서 만들 수 있습니다. 기본 이미지를 선택하여 해당 컨테이너를 스핀업하고 컨테이너에 원하는 모든 소프트웨어와 종속성을 설치합니다.

그런 다음 `docker commit container name`을 사용하면 이 이미지의 로컬 복사본이 docker 이미지와 Docker Desktop 이미지 탭 아래에 있습니다.

매우 간단하지만, 프로세스를 이해하고 싶은 게 아니라면, 이 방법을 권장하지 않습니다. 이 방법은 라이프사이클 관리를 관리하기가 매우 어렵고 수동 구성/재구성이 많이 필요합니다. 하지만 docker 이미지를 빌드하는 가장 빠르고 간단한 방법입니다. 테스트, 문제 해결, 종속성 검증 등에 적합합니다.

우리가 이미지를 생성하려는 방법은 Dockerfile을 통한 방법입니다. 이는 이미지를 만드는 깔끔하고 간결하며 반복 가능한 방법을 제공합니다. 더 쉬운 라이프사이클 관리와 지속적인 통합 및 지속적인 배포 프로세스에 쉽게 통합할 수 있습니다. 하지만 처음 언급된 방법보다 조금 더 어려울 수 있습니다.

Dockerfile 방법을 사용하는 것이 실제 엔터프라이즈급 컨테이너 배포에 훨씬 더 잘 맞습니다.

Dockerfile은 3단계 프로세스를 통해 Dockerfile을 만들고 이미지를 조립하는 데 필요한 명령을 추가합니다.

다음 표는 사용자가 가장 많이 사용할 가능성이 높은 몇 가지 Dockerfile 설정을 보여줍니다.

| Command    | Purpose                                                                                                                              |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| FROM       | 상위 이미지를 지정합니다.                                                                                                            |
| WORKDIR    | Dockerfile에서 다음에 나오는 명령의 작업 디렉터리를 설정합니다.                                                                      |
| RUN        | 컨테이너에 필요한 애플리케이션과 패키지를 설치합니다.                                                                                |
| COPY       | 특정 위치에서 파일 또는 디렉터리를 복사합니다.                                                                                       |
| ADD        | 복사뿐만 아니라 원격 URL을 처리하고 압축 파일의 압축을 풀 수도 있습니다.                                                             |
| ENTRYPOINT | 컨테이너가 시작될 때 항상 실행되는 명령입니다. 지정하지 않으면 기본값은 "/bin/sh -c"입니다.                                          |
| CMD        | 엔트리포인트로 전달된 인자입니다. 엔트리포인트가 설정되지 않은 경우(기본값은 "/bin/sh -c"), 컨테이너가 실행하는 명령은 CMD가 됩니다. |
| EXPOSE     | 컨테이너 애플리케이션에 액세스할 포트를 정의합니다.                                                                                  |
| LABEL      | 이미지에 메타데이터를 추가합니다.                                                                                                    |

이제 첫 번째 Dockerfile을 빌드하는 방법에 대한 세부 정보를 얻었으므로 작업 디렉터리를 생성하고 Dockerfile을 만들 수 있습니다. 이 리포지토리 내에 작업 디렉터리를 만들었는데, [여기서](/2022/Days/Containers/) 제가 살펴봐야 할 파일과 폴더를 볼 수 있습니다.

이 디렉터리에는 지난 섹션에서 사용한 .gitignore와 유사한 .dockerignore 파일을 만들겠습니다. 이 파일에는 최종 빌드에서 제외하려는 Docker 빌드 프로세스 중에 생성되는 모든 파일이 나열됩니다.

컨테이너는 부피가 커지지 않고 가능한 한 빠르게 컴팩트하게 만들어야 함을 기억하세요.

위에 링크된 폴더에서 아래 레이아웃으로 매우 간단한 Dockerfile을 만들고 싶습니다.

```dockerfile
# 공식 우분투 18.04를 기본으로 사용하세요.
FROM ubuntu:18.04
# nginx 및 curl 설치
RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y nginx curl
RUN rm -rf /var/lib/apt/lists/*
```

터미널에서 이 디렉토리로 이동한 다음 `docker build -t 90daysofdevops:0.1 .`를 실행합니다. `-t`를 사용한 다에는 이미지 이름과 태그를 설정할 수 있습니다.

![](/2022/Days/Images/Day45_Containers3.png)

이제 이미지를 만들었으므로 Docker Desktop을 사용하여 이미지를 실행하거나 docker 커맨드라인을 사용할 수 있습니다. 저의 경우에는 Docker Desktop을 사용하여 컨테이너를 실행하고, 컨테이너의 cli에서 `curl`을 사용한 것을 볼 수 있습니다.

![](/2022/Days/Images/Day45_Containers4.png)

Docker Desktop에서는 UI를 활용하여 이 새로운 이미지로 몇 가지 작업을 더 수행할 수 있습니다.

![](/2022/Days/Images/Day45_Containers5.png)

이미지를 검사하면 컨테이너 내에서 실행하려고 하는 Dockerfile과 코드들을 매우 많이 볼 수 있습니다.

![](/2022/Days/Images/Day45_Containers6.png)

pull 옵션이 있지만, 이 이미지는 어디에도 호스팅되지 않기 때문에 이 옵션은 실패할 것이고 오류로 표시됩니다. 하지만 Push to hub 옵션이 있어 이미지를 DockerHub로 push할 수 있습니다.

앞서 실행한 것과 동일한 `docker build`를 사용하는 경우에도 이 방법이 작동하지 않으므로 빌드 명령을 `docker build -t {{username}}/{{imagename}}:{{version}}`로 해야 합니다.

![](/2022/Days/Images/Day45_Containers7.png)

이제 DockerHub 리포지토리로 이동하여 살펴보면 방금 새 이미지를 push한 것을 확인할 수 있습니다. 이제 Docker Desktop에서 해당 pull 탭을 사용할 수 있습니다.

![](/2022/Days/Images/Day45_Containers8.png)

## 자료

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)
- [Blog on gettng started building a docker image](https://stackify.com/docker-build-a-beginners-guide-to-building-docker-/2022/Days/images/)
- [Docker documentation for building an image](https://docs.docker.com/develop/develop-/2022/Days/images/dockerfile_best-practices/)

[Day 46](day46.md)에서 봐요!
