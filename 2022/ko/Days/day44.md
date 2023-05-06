---
title: '#90DaysOfDevOps - Docker Images & Hands-On with Docker Desktop - Day 44'
published: false
description: 90DaysOfDevOps - Docker Images & Hands-On with Docker Desktop
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048708
---

## Docker 이미지 및 Docker Desktop 실습하기

이제 시스템에 Docker Desktop을 설치했습니다. (Linux를 실행하는 경우 여전히 옵션은 있지만 GUI는 없지만 Docker는 Linux에서 작동합니다.)[우분투에 Docker Engine 설치](https://docs.docker.com/engine/install/ubuntu/) (다른 배포판도 사용할 수 있습니다.)

이 글에서는 몇 가지 이미지를 우리 환경에 배포하는 작업을 시작하겠습니다. Docker 이미지에 대한 요약 - Docker 이미지는 Docker 컨테이너에서 코드를 실행하는 데 사용되는 파일입니다. Docker 이미지는 템플릿처럼 Docker 컨테이너를 빌드하기 위한 일련의 지침 역할을 합니다. 또한 Docker 이미지는 Docker를 사용할 때 시작점 역할을 하기도 합니다.

지금 바로 [DockerHub](https://hub.docker.com/)에서 계정을 생성하세요.

![](/2022/Days/Images/Day44_Containers1.png)

DockerHub는 Docker 및 그 구성 요소로 작업하기 위한 중앙 집중식 리소스입니다. 가장 일반적으로는 docker 이미지를 호스팅하는 레지스트리로 알려져 있습니다. 그러나 여기에는 부분적으로 자동화와 함께 사용하거나 보안 검사뿐만 아니라 GitHub에 통합할 수 있는 많은 추가 서비스가 있습니다.

로그인한 후 아래로 스크롤하면 컨테이너 이미지 목록이 표시되며, MySQL, hello-world 등의 데이터베이스 이미지가 표시될 수 있습니다. 데이터베이스 이미지가 필요하거나 직접 만들 필요가 없는 경우 공식 이미지를 사용하는 것이 가장 좋습니다.

![](/2022/Days/Images/Day44_Containers2.png)

사용 가능한 이미지 보기를 더 자세히 살펴보고 카테고리, 운영 체제 및 아키텍처에 따라 검색할 수 있습니다. 아래에서 강조 표시한 것은 공식 이미지로, 이 컨테이너 이미지의 출처에 대해 안심할 수 있습니다.

![](/2022/Days/Images/Day44_Containers3.png)

특정 이미지를 검색할 수도 있습니다. 예를 들어 워드프레스가 좋은 기본 이미지가 될 수 있으므로 상단에 있는 이미지를 검색하여 워드프레스와 관련된 모든 컨테이너 이미지를 찾을 수 있습니다. 아래는 확인된 퍼블리셔에 대한 공지사항입니다.

- 공식 이미지 - Docker 공식 이미지는 엄선된 Docker 오픈 소스 및 "drop-in" 솔루션 리포지토리 집합입니다.

- 검증된 퍼블리셔 - 검증된 퍼블리셔의 고품질 Docker 콘텐츠입니다. 이러한 제품은 상업적 주체가 직접 게시하고 유지 관리합니다.

![](/2022/Days/Images/Day44_Containers4.png)

### Docker Desktop 살펴보기

저희 시스템에는 Docker Desktop이 설치되어 있으며, 이 파일을 열면 이미 설치되어 있지 않다면 아래 이미지와 비슷한 것을 볼 수 있을 것입니다. 보시다시피 컨테이너가 실행되고 있지 않고 Docker Engine이 실행되고 있습니다.

![](/2022/Days/Images/Day44_Containers5.png)

새로 설치한 것이 아니기 때문에 이미 다운로드하여 시스템에 사용할 수 있는 이미지가 몇 개 있습니다. 여기에는 아무것도 표시되지 않을 것입니다.

![](/2022/Days/Images/Day44_Containers6.png)

원격 리포지토리 아래에서 DockerHub에 저장한 컨테이너 이미지를 찾을 수 있습니다. 아래에서 볼 수 있듯이 이미지가 없습니다.

![](/2022/Days/Images/Day44_Containers7.png)

DockerHub 사이트에서도 리포지토리가 없음을 확인할 수 있습니다.

![](/2022/Days/Images/Day44_Containers8.png)

다음으로 Volumes 탭이 있는데, 지속성이 필요한 컨테이너가 있는 경우 여기에서 로컬 파일 시스템이나 공유 파일 시스템에 이러한 Volumes를 추가할 수 있습니다.

![](/2022/Days/Images/Day44_Containers9.png)

이 글을 쓰는 시점에 Environments 탭도 있는데, 이 탭은 다른 git branch 사이를 이동하지 않고 팀과 협업하는 데 도움이 될 것입니다. 여기서는 다루지 않겠습니다.

![](/2022/Days/Images/Day44_Containers10.png)

첫 번째 탭으로 돌아가면 시작 컨테이너를 실행할 수 있는 명령이 있음을 알 수 있습니다. 터미널에서 `docker run -d -p 80:80 docker/getting-started`를 실행해 보겠습니다.

![](/2022/Days/Images/Day44_Containers11.png)

Docker Desktop 창을 다시 확인하면 컨테이너가 실행 중인 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day44_Containers12.png)

WSL2를 사용하고 있다는 것을 눈치채셨을 텐데요, 이를 사용하려면 설정에서 이 기능이 활성화되어 있는지 확인해야 합니다.

![](/2022/Days/Images/Day44_Containers13.png)

이제 Images 탭으로 다시 이동하여 확인하면 이제 사용 중인 이미지인 docker/getting-started를 볼 수 있습니다.

![](/2022/Days/Images/Day44_Containers14.png)

Containers/Apps 탭으로 돌아가서 실행 중인 컨테이너를 클릭합니다. 기본적으로 로그가 표시되고 상단에 선택할 수 있는 몇 가지 옵션이 있는데, 이 컨테이너에서 실행 중인 웹 페이지가 될 것이 확실하므로 브라우저에서 열기를 선택하겠습니다.

![](/2022/Days/Images/Day44_Containers15.png)

위의 버튼을 누르면 로컬호스트로 연결되는 웹 페이지가 열리고 아래와 비슷한 내용이 표시됩니다.

이 컨테이너에는 컨테이너와 이미지에 대한 자세한 내용도 있습니다.

![](/2022/Days/Images/Day44_Containers16.png)
이제 첫 번째 컨테이너를 실행했습니다. 아직은 쉽습니다. 컨테이너 이미지 중 하나를 DockerHub에서 가져오고 싶다면 어떻게 해야 할까요? 아마도 우리가 사용할 수 있는 `hello world` docker 컨테이너가 있을 것입니다.

시작 컨테이너는 리소스를 많이 차지하기 때문이 아니라 몇 가지 단계를 더 진행하면서 깔끔하게 정리하기 위해 중단했습니다.

터미널로 돌아와서 `docker run hello-world`를 실행하고 어떤 일이 일어나는지 살펴봅시다.

로컬에 이미지가 없었기 때문에 이미지를 끌어내렸고, 컨테이너 이미지에 시작 및 실행에 대한 몇 가지 정보와 참조 지점에 대한 링크가 포함된 메시지가 표시되는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day44_Containers17.png)

그러나 이제 Docker Desktop을 살펴보면 실행 중인 컨테이너는 없지만, hello-world 메시지를 사용한 종료된 컨테이너, 즉 시작되어 메시지를 전달한 후 종료된 컨테이너는 있습니다.

![](/2022/Days/Images/Day44_Containers18.png)

마지막으로 Images 탭을 확인해보면 시스템에 로컬로 새로운 hello-world 이미지가 있는 것을 확인할 수 있습니다. 즉, 터미널에서 `docker run hello-world` 명령을 다시 실행해도 버전이 변경되지 않는 한 아무것도 가져올 필요가 없습니다.

![](/2022/Days/Images/Day44_Containers19.png)

'hello-world' 컨테이너에서 온 메시지로 인해 조금 더 야심찬 도전을 하고 싶어졌습니다.

도전!

![](/2022/Days/Images/Day44_Containers20.png)

터미널에서 `docker run -it ubuntu bash`를 실행할 때 우리는 운영 체제의 전체 복사본이 아닌 우분투의 컨테이너화된 버전을 실행할 것입니다. 이 특정 이미지에 대한 자세한 내용은 [DockerHub](https://hub.docker.com/_/ubuntu)에서 확인할 수 있습니다.

아래 명령어를 실행하면 대화형 프롬프트(`-it`)가 나타나고 컨테이너에 bash 셸이 생성되는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day44_Containers21.png)

bash 셸이 있지만 그 이상은 없기 때문에 이 컨테이너 이미지는 30MB 미만입니다.

![](/2022/Days/Images/Day44_Containers22.png)

하지만 여전히 이 이미지를 사용할 수 있고, apt package manager를 사용하여 소프트웨어를 설치할 수 있으며, 컨테이너 이미지를 업데이트하고 업그레이드할 수도 있습니다.

![](/2022/Days/Images/Day44_Containers23.png)

또는 컨테이너에 일부 소프트웨어를 설치하고 싶을 수도 있습니다. pinta는 이미지 편집기이고 200MB가 넘기 때문에 여기서는 정말 나쁜 예를 선택했지만 제가 이걸로 무엇을 하려는지 이해하시길 바랍니다. 이렇게 하면 컨테이너의 크기가 상당히 커지지만, 여전히 GB가 아닌 MB 단위로 사용할 것입니다.

![](/2022/Days/Images/Day44_Containers24.png)

이 글에서는 사용 사례를 통해 Docker Desktop과 컨테이너의 세계에 대한 개요를 간단하고 접근하기 쉬운 방식으로 제공하고자 합니다. 그러나 컨테이너 이미지를 다운로드하고 사용하는 것 외에도 네트워킹, 보안 및 기타 사용 가능한 옵션에 대해서도 다뤄야 합니다. 이 섹션의 궁극적인 목표는 무언가를 만들어서 DockerHub 리포지토리에 업로드하고 배포하는 것입니다.

## 자료

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)

[Day 45](day45.md)에서 봐요!
