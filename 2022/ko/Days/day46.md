---
title: '#90DaysOfDevOps - Docker Compose - Day 46'
published: false
description: 90DaysOfDevOps - Docker Compose
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048740
---

## Docker Compose

하나의 컨테이너를 실행할 수 있는 기능은 단일 사용 사례에 필요한 모든 것을 갖춘 독립적인 이미지가 있는 경우 유용할 수 있지만, 서로 다른 컨테이너 이미지 간에 여러 애플리케이션을 빌드하려는 경우 흥미로운 문제가 발생할 수 있습니다. 예를 들어, 웹사이트 프론트엔드가 있지만 백엔드 데이터베이스가 필요한 경우 모든 것을 하나의 컨테이너에 넣을 수 있지만 데이터베이스용 컨테이너를 사용하는 것이 더 효율적이고 좋을 것입니다.

이때 여러 컨테이너에서 더 복잡한 앱을 실행할 수 있는 도구인 Docker Compose가 등장합니다. 단일 파일과 명령을 사용하여 애플리케이션을 스핀업할 수 있다는 이점이 있습니다. 이 글에서 안내하는 예제는 [Docker 빠른 시작 샘플 앱(빠른 시작: 작성 및 워드프레스)](https://docs.docker.com/samples/wordpress/)에서 가져온 것입니다.

이 첫 번째 예제에서는

- Docker Compose를 사용하여 WordPress와 별도의 MySQL 인스턴스를 불러옵니다.
- 'docker-compose.yml'이라는 YAML 파일을 사용합니다.
- 프로젝트 빌드
- 브라우저를 통해 워드프레스 구성
- 종료 및 정리

### Docker Compose 설치

앞서 언급했듯이 docker compose는 도구이며, 맥OS나 윈도우를 사용하는 경우 Docker Desktop 설치에 compose가 포함되어 있습니다. 하지만, 윈도우 서버 호스트나 리눅스 서버에서 컨테이너를 실행하고 싶을 수 있으며, 이 경우 다음 [docker compose 설치](https://docs.docker.com/compose/install/) 지침을 사용하여 설치할 수 있습니다.

터미널을 열고 위의 명령어를 입력하면 시스템에 `docker-compose`가 설치되었는지 확인할 수 있습니다.

![](/2022/Days/Images/Day46_Containers1.png)

### Docker-Compose.yml(YAML)

다음에 이야기할 것은 리포지토리의 컨테이너 폴더에서 찾을 수 있는 docker-compose.yml입니다. 그 전에, YAML에 대해 조금 설명하겠습니다.

거의 모든 프로그래밍 언어에서 사용할 수 있는 YAML은 다양한 곳에서 사용될 수 있습니다.

"YAML은 모든 프로그래밍 언어에 대한 인간 친화적인 데이터 직렬화 언어입니다."

일반적으로 구성 파일과 데이터를 저장하거나 전송하는 일부 애플리케이션에서 사용됩니다. 동일한 구성 파일을 제공하는 경향이 있는 XML 파일을 접한 적이 있을 것입니다. YAML은 최소한의 구문을 제공하지만, 동일한 사용 사례를 목표로 합니다.

YAML(YAML Ain't Markup Language)은 지난 몇 년 동안 꾸준히 인기가 높아진 직렬화 언어입니다. 객체 직렬화 기능 덕분에 JSON과 같은 언어를 대체할 수 있습니다.

YAML의 약어는 Yet Another Markup Language의 약자였습니다. 그러나 유지 관리자는 데이터 지향적 기능을 더 강조하기 위해 YAML Ain't Markup Language로 이름을 변경했습니다.

어쨌든, 다시 docker-compose.yml 파일로 돌아가겠습니다. 이 파일은 단일 시스템에 여러 개의 컨테이너를 배포할 때 수행하고자 하는 작업의 구성 파일입니다.

위에 링크된 튜토리얼에서 바로 파일의 내용을 보면 다음과 같이 보입니다:

```yaml
version: '3.9'

services:
  DB:
    image: mysql:5.7
    volumes:
      - db_data:/var/lib/mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: somewordpress
      MYSQL_DATABASE: wordpress
      MYSQL_USER: wordpress
      MYSQL_PASSWORD: wordpress

  wordpress:
    depends_on:
      - db
    image: wordpress:latest
    volumes:
      - wordpress_data:/var/www/html
    ports:
      - '8000:80'
    restart: always
    environment:
      WORDPRESS_DB_HOST: db
      WORDPRESS_DB_USER: wordpress
      WORDPRESS_DB_PASSWORD: wordpress
      WORDPRESS_DB_NAME: wordpress
volumes:
  db_data: {}
  wordpress_data: {}
```

버전을 선언한 다음, 이 docker-compose.yml 파일의 대부분은 서비스로 구성되어 있으며, DB 서비스와 워드프레스 서비스가 있습니다. 각 서비스에는 버전 태그가 연결된 이미지가 정의되어 있는 것을 볼 수 있습니다. 이제 첫 번째 연습과 달리 구성에 상태도 도입하고 있는데, 데이터베이스를 저장할 수 있도록 볼륨을 생성하겠습니다.

그런 다음 비밀번호 및 사용자 이름과 같은 몇 가지 환경 변수가 있습니다. 이러한 파일은 매우 복잡해질 수 있지만 YAML 구성 파일을 사용하면 전체적으로 단순화할 수 있습니다.

### 프로젝트 빌드

이제 터미널로 돌아가서 docker-compose 도구로 몇 가지 명령을 사용할 수 있습니다. 디렉토리로 이동하여 docker-compose.yml 파일이 있는 디렉터리로 이동합니다.

터미널에서 `docker-compose up -d`를 실행하면 해당 이미지를 가져와 멀티 컨테이너 애플리케이션을 세우는 프로세스가 시작됩니다.

이 명령의 `-d`는 분리 모드를 의미하며, 이는 실행 명령이 백그라운드에서 실행 중이거나 실행될 것임을 의미합니다.

![](/2022/Days/Images/Day46_Containers2.png)

이제 `docker ps` 명령을 실행하면 2개의 컨테이너가 실행 중이며 하나는 WordPress이고 다른 하나는 MySQL인 것을 볼 수 있습니다.

![](/2022/Days/Images/Day46_Containers3.png)

다음으로 브라우저를 열고 `http://localhost:8000`으로 이동하면 WordPress가 실행 중인지 확인할 수 있으며, WordPress 설정 페이지가 표시됩니다.

![](/2022/Days/Images/Day46_Containers4.png)

워드프레스 설정을 완료한 다음 아래 콘솔에서 원하는 대로 웹사이트 구축을 시작할 수 있습니다.

![](/2022/Days/Images/Day46_Containers5.png)

이제 새 탭을 열고 `http://localhost:8000` 이전과 동일한 주소로 이동하면 사이트 제목이 "90DaysOfDevOps"인 간단한 기본 테마가 표시되고 샘플 게시물이 표시됩니다.

![](/2022/Days/Images/Day46_Containers6.png)

변경하기 전에 Docker Desktop을 열고 볼륨 탭으로 이동하면 컨테이너와 연결된 두 개의 볼륨(하나는 워드프레스용, 다른 하나는 DB용)을 볼 수 있습니다.

![](/2022/Days/Images/Day46_Containers7.png)

현재 워드프레스 테마는 "Twenty Twenty-Two"이며 이를 "Twenty Twenty"로 변경하고 싶습니다. 대시보드로 돌아와서 변경할 수 있습니다.

![](/2022/Days/Images/Day46_Containers8.png)

또한 사이트에 새 게시물을 추가하려고 하는데, 아래에서 새 사이트의 최신 버전을 볼 수 있습니다.

![](/2022/Days/Images/Day46_Containers9.png)

### 정리 여부

이제 `docker-compose down` 명령을 사용하면 컨테이너가 삭제됩니다. 하지만 볼륨은 그대로 유지됩니다.

![](/2022/Days/Images/Day46_Containers10.png)

Docker Desktop에서 볼륨이 여전히 존재한다는 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day46_Containers11.png)

그런 다음 다시 백업하려면 동일한 디렉토리 내에서 `docker up -d` 명령을 실행하면 애플리케이션을 다시 백업하고 실행할 수 있습니다.

![](/2022/Days/Images/Day46_Containers12.png)

그런 다음 브라우저에서 동일한 주소인 `http://localhost:8000`으로 이동하면 새 글과 테마 변경 사항이 모두 그대로 유지되는 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day46_Containers13.png)

컨테이너와 해당 볼륨을 제거하려면 `docker-compose down --volumes`를 실행하면 볼륨도 제거됩니다.

![](/2022/Days/Images/Day46_Containers14.png)

이제 `docker-compose up -d`를 다시 사용하면 시작하지만, 이미지는 여전히 우리 시스템의 로컬에 있으므로 DockerHub 리포지토리에서 다시 가져올 필요가 없습니다.

docker-compose와 그 기능에 대해 알아보기 시작했을 때 이것이 Kubernetes와 같은 컨테이너 오케스트레이션 도구와 어디에 위치하는지 혼란스러웠는데, 이 짧은 데모에서 우리가 한 모든 작업은 로컬 데스크톱 머신에서 실행 중인 WordPress와 DB가 있는 호스트 하나에 초점을 맞추고 있습니다. 여러 가상 머신이나 여러 물리적 머신이 없으며 애플리케이션의 요구 사항을 쉽게 확장 및 축소할 수도 없습니다.

다음 섹션에서는 Kubernetes를 다룰 예정이지만, 먼저 컨테이너에 대한 전반적인 내용을 며칠 더 살펴보겠습니다.

[Awesome-Compose](https://github.com/docker/awesome-compose)는 여러 통합이 있는 docker-compose 애플리케이션의 샘플을 위한 훌륭한 리소스입니다.

위의 리포지토리에는 단일 노드에 Elasticsearch, Logstash, Kibana(ELK)를 배포하는 훌륭한 예제가 있습니다.

[Containers 폴더](/2022/Days/Containers/elasticsearch-logstash-kibana/)에 파일을 업로드했습니다. 이 폴더가 로컬에 있으면 해당 폴더로 이동하여 `docker-compose up -d`를 사용하면 간단하게 설치할 수 있습니다.

![](/2022/Days/Images/Day46_Containers15.png)

그런 다음 `docker ps`로 실행 중인 컨테이너가 있는지 확인할 수 있습니다.

![](/2022/Days/Images/Day46_Containers16.png)

이제 각 컨테이너에 대한 브라우저를 열 수 있습니다:

![](/2022/Days/Images/Day46_Containers17.png)

모든 것을 제거하려면 `docker-compose down` 명령을 사용할 수 있습니다.

## 자료

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [WSL 2 with Docker getting started](https://www.youtube.com/watch?v=5RQbdMn04Oc)
- [Blog on getting started building a docker image](https://stackify.com/docker-build-a-beginners-guide-to-building-docker-/2022/Days/images/)
- [Docker documentation for building an image](https://docs.docker.com/develop/develop-/2022/Days/images/dockerfile_best-practices/)
- [YAML Tutorial: Everything You Need to Get Started in Minute](https://www.cloudbees.com/blog/yaml-tutorial-everything-you-need-get-started)

[Day 47](day47.md)에서 봐요!
