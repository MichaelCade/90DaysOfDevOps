---
title: '#90DaysOfDevOps - Docker Networking & Security - Day 47'
published: false
description: 90DaysOfDevOps - Docker Networking & Security
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049078
---

## Docker 네트워킹 및 보안

지금까지 컨테이너 세션에서 우리는 작업을 수행했지만, 네트워킹 관점에서 작업이 어떻게 작동했는지 살펴보지 않았으며 보안에 대해서도 다루지 않았습니다.

### Docker 네트워킹 기본 사항

터미널을 열고, 컨테이너 네트워크를 구성하고 관리하기 위한 기본 명령어인 `docker network` 명령을 입력합니다.

아래에서 이 명령어를 사용하는 방법과 사용 가능한 모든 하위 명령을 확인할 수 있습니다. 새로운 네트워크를 생성하고, 기존 네트워크를 나열하고, 네트워크를 검사 및 제거할 수 있습니다.

![](/2022/Days/Images/Day47_Containers1.png)

설치 이후 우리가 가지고 있는 기존 네트워크를 살펴보기 위해 `docker network list` 명령을 사용하는 기본 Docker 네트워킹의 모습을 살펴봅시다.

각 네트워크는 고유한 ID와 이름을 갖습니다. 각 네트워크는 또한 단일 드라이버와 연결됩니다. "bridge" 네트워크와 "host" 네트워크는 각각의 드라이버와 이름이 동일합니다.

![](/2022/Days/Images/Day47_Containers2.png)

다음으로 `docker network inspect` 명령으로 네트워크를 더 자세히 살펴볼 수 있습니다.

`docker network inspect bridge`를 실행하면 특정 네트워크 이름에 대한 모든 구성 세부 정보를 얻을 수 있습니다. 여기에는 이름, ID, 드라이버, 연결된 컨테이너 등이 포함되며 보시다시피 훨씬 더 많은 정보를 확인할 수 있습니다.

![](/2022/Days/Images/Day47_Containers3.png)

### Docker: bridge 네트워킹

위에서 보았듯이 Docker Desktop을 표준 설치하면 `bridge`라는 사전 구축된 네트워크가 제공됩니다. `docker network list` 명령을 다시 실행하면 bridge라는 네트워크가 `bridge` 드라이버와 연결되어 있는 것을 볼 수 있습니다. 이름이 같다고 해서 같은 것은 아닙니다. 연결되었지만 같은 것은 아닙니다.

위의 출력은 또한 bridge 네트워크가 로컬로 범위가 지정되었음을 보여줍니다. 이는 네트워크가 이 Docker host에만 존재한다는 것을 의미합니다. 이는 bridge 드라이버를 사용하는 모든 네트워크에 해당되며, bridge 드라이버는 단일 host 네트워킹을 제공합니다.

bridge 드라이버로 생성된 모든 네트워크는 Linux bridge(virtual switch라고도 함)를 기반으로 합니다.

### 컨테이너 연결

기본적으로 bridge 네트워크는 새 컨테이너에 할당되므로 네트워크를 지정하지 않으면 모든 컨테이너가 bridge 네트워크에 연결됩니다.

`docker run -dt ubuntu sleep infinity` 명령으로 새 컨테이너를 생성해 보겠습니다.

위의 sleep 명령은 컨테이너를 백그라운드에서 계속 실행시켜서 컨테이너를 마음대로 다룰 수 있도록 합니다.

![](/2022/Days/Images/Day47_Containers4.png)

이제 `docker network inspect bridge`로 bridge 네트워크를 확인하면 네트워크를 지정하지 않았기 때문에 방금 배포한 것과 일치하는 컨테이너가 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day47_Containers5.png)

`docker exec -it 3a99af449ca2 bash`를 사용하여 컨테이너를 자세히 살펴볼 수도 있습니다. 컨테이너 ID를 얻으려면 `docker ps`를 사용해야 합니다.

여기에서 이미지에는 핑할 항목이 없으므로, `apt-get update && apt-get install -y iputils-ping`을 실행한 다음 외부 인터페이스 주소를 핑합니다. `ping -c5 www.90daysofdevops.com`

![](/2022/Days/Images/Day47_Containers6.png)

이 문제를 해결하기 위해 `docker stop 3a99af449ca2`를 다시 실행하고 `docker ps`를 사용하여 컨테이너 ID를 찾을 수 있지만, 이렇게 하면 컨테이너가 제거됩니다.

### 외부 연결을 위한 NAT 구성하기

이 단계에서는 새 NGINX 컨테이너를 시작하고 Docker host의 포트 8080을 컨테이너 내부의 포트 80으로 매핑합니다. 즉, 포트 8080의 Docker host에 도달하는 트래픽은 컨테이너 내부의 포트 80으로 전달됩니다.

`docker run --name web1 -d -p 8080:80 nginx`를 실행하여 공식 NGINX 이미지를 기반으로 새 컨테이너를 시작합니다.

![](/2022/Days/Images/Day47_Containers7.png)

`docker ps`를 실행하여 컨테이너 상태와 포트 매핑을 검토합니다.

![](/2022/Days/Images/Day47_Containers8.png)

맨 위 줄은 NGINX를 실행하는 새 web1 컨테이너를 보여줍니다. 컨테이너가 실행 중인 명령과 포트 매핑에 주목하세요. - `0.0.0.0:8080->80/tcp`는 모든 host 인터페이스의 8080 포트를 web1 컨테이너 내부의 80 포트에 매핑합니다. 이 포트 매핑을 통해 외부 소스에서 컨테이너의 웹 서비스에 효과적으로 액세스할 수 있습니다(포트 8080의 Docker host IP 주소를 통해).

이제 실제 host에 대한 IP 주소가 필요하며, WSL 터미널로 이동하여 `IP addr` 명령을 사용하여 이를 수행할 수 있습니다.

![](/2022/Days/Images/Day47_Containers9.png)

그런 다음, 이 IP를 가지고 브라우저를 열어 `http://172.25.218.154:8080/`로 이동하면 IP가 다를 수 있습니다. 이렇게 하면 NGINX에 액세스할 수 있음을 확인할 수 있습니다.

![](/2022/Days/Images/Day47_Containers10.png)

이 사이트의 이 지침은 2017년 DockerCon에서 가져온 것이지만 오늘날에도 여전히 유효합니다. 그러나 나머지 연습은 Docker Swarm에 대한 것이므로 여기서는 다루지 않겠습니다. [Docker 네트워킹 - DockerCon 2017](https://github.com/docker/labs/tree/master/dockercon-us-2017/docker-networking)

### 컨테이너 보안

컨테이너는 전체 서버 구성에 비해 워크로드에 안전한 환경을 제공합니다. 컨테이너는 애플리케이션을 서로 격리된 훨씬 더 작고 느슨하게 결합된 구성 요소로 분할할 수 있는 기능을 제공하므로 전체적으로 공격 표면을 줄이는 데 도움이 됩니다.

하지만 시스템의 취약점을 공격하려는 해커로부터 자유롭지는 않습니다. 우리는 여전히 이 기술의 보안 결함을 이해하고 모범 사례를 유지해야 합니다.

### 루트 권한에서 벗어나기

지금까지 배포한 모든 컨테이너는 컨테이너 내 프로세스에 대한 루트 권한을 사용해 왔습니다. 즉, 컨테이너와 host 환경에 대한 모든 관리 액세스 권한이 있다는 뜻입니다. 이제 이러한 시스템이 오래 가동되지 않을 것이라는 것을 알고 있었습니다. 하지만 시작하고 실행하는 것이 얼마나 쉬운지 보셨을 것입니다.

프로세스에 몇 단계를 추가하여 루트 사용자가 아닌 사용자가 선호하는 모범 사례를 사용할 수 있도록 할 수 있습니다. dockerfile을 만들 때 사용자 계정을 만들 수 있습니다. 이 예제는 리포지토리의 컨테이너 폴더에서도 찾을 수 있습니다.

```dockerfile
# 공식 Ubuntu 18.04를 기본으로 사용하세요.
FROM ubuntu:18.04
RUN apt-get update && apt-get upgrade -y
RUN groupadd -g 1000 basicuser && useradd -r -u 1000 -g basicuser basicuser
USER basicuser
```

`docker run --user 1009 ubuntu` 명령은 dockerfile에 지정된 모든 사용자를 재정의합니다. 따라서 다음 예제에서는 컨테이너가 항상 권한이 가장 낮은 사용자 식별자 1009로 실행되며 권한 수준도 가장 낮습니다.

그러나 이 방법은 이미지 자체의 근본적인 보안 결함을 해결하지 못합니다. 따라서 컨테이너가 항상 안전하게 실행되도록 dockerfile에 루트 사용자가 아닌 사용자를 지정하는 것이 좋습니다.

### 비공개 레지스트리

조직에서 컨테이너 이미지의 비공개 레지스트리를 설정하면 원하는 곳에서 호스팅할 수 있거나 이를 위한 관리형 서비스도 있지만, 대체로 사용자와 팀이 사용할 수 있는 이미지를 완벽하게 제어할 수 있습니다.

DockerHub는 기준선을 제시하는 데는 훌륭하지만, 이미지 게시자를 많이 신뢰해야 하는 기본적인 서비스만 제공합니다.

### 린 & 클린

보안과 관련되지는 않았지만, 전체적으로 언급했습니다. 그러나 애플리케이션에서 사용하지 않는 리소스가 컨테이너에 필요하지 않은 경우 컨테이너의 크기는 공격 표면 측면에서 보안에 영향을 미칠 수 있습니다.

또한 `latest` 이미지를 가져올 때 이미지에 많은 부풀림을 가져올 수 있기 때문에 이것이 저의 주요 관심사입니다. DockerHub는 리포지토리에 있는 각 이미지의 압축 크기를 표시합니다.

`docker image`를 확인하면 이미지의 크기를 확인할 수 있는 좋은 명령어입니다.

![](/2022/Days/Images/Day47_Containers11.png)

## 자료

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)
- [Blog on getting started building a docker image](https://stackify.com/docker-build-a-beginners-guide-to-building-docker-/2022/Days/images/)
- [Docker documentation for building an image](https://docs.docker.com/develop/develop-/2022/Days/images/dockerfile_best-practices/)
- [YAML Tutorial: Everything You Need to Get Started in Minute](https://www.cloudbees.com/blog/yaml-tutorial-everything-you-need-get-started)

[Day 48](day48.md)에서 봐요!
