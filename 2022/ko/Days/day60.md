---
title: '#90DaysOfDevOps - Docker Containers, Provisioners & Modules - Day 60'
published: false
description: '90DaysOfDevOps - Docker Containers, Provisioners & Modules'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049052
---

## Docker 컨테이너, Provisioners 및 모듈

[Day 59](day59.md)에는 Terraform을 사용하여 로컬 무료 VirtualBox 환경에 가상 머신을 프로비저닝했습니다. 이 섹션에서는 몇 가지 구성이 포함된 Docker 컨테이너를 로컬 Docker 환경에 배포해 보겠습니다.

### Docker 데모

먼저 아래 코드 블록을 사용하여 간단한 웹 앱을 Docker에 배포하고 이를 게시하여 네트워크에서 사용할 수 있도록 하겠습니다. nginx를 사용할 것이며 로컬 호스트와 포트 8000을 통해 노트북에서 외부에서 사용할 수 있도록 할 것입니다. 커뮤니티에서 제공하는 Docker 공급자를 사용하고 있으며 구성에서도 사용 중인 Docker 이미지를 확인할 수 있습니다.

```
terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "2.16.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false
}

resource "docker_container" "nginx" {
  image = docker_image.nginx.latest
  name  = "tutorial"
  ports {
    internal = 80
    external = 8000
  }
}
```

첫 번째 작업은 `terraform init` 명령을 사용하여 로컬 머신에 프로바이더를 다운로드하는 것입니다.

![](/2022/Days/Images/Day60_IAC1.png)

그런 다음 `terraform apply`를 실행한 다음 `docker ps`를 실행하면 컨테이너가 실행되는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day60_IAC2.png)

이제 브라우저를 열어 `http://localhost:8000/`으로 이동하면 NGINX 컨테이너에 액세스할 수 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day60_IAC3.png)

자세한 내용은 [Docker Provider](https://registry.terraform.io/providers/kreuzwerker/docker/latest/docs/resources/container)에서 확인할 수 있습니다.

위는 Terraform과 Docker로 무엇을 할 수 있는지, 그리고 Terraform 상태에서 어떻게 관리할 수 있는지에 대한 아주 간단한 데모입니다. 컨테이너 섹션에서 docker-compose에 대해 다뤘고, 이것과 IaC, 그리고 Kubernetes 사이에는 약간의 교차점이 있습니다.

이를 보여드리고 Terraform이 어떻게 좀 더 복잡한 것을 처리할 수 있는지 보여드리기 위해, 우리가 docker-compose로 만든 워드프레스와 MySQL용 docker-compose 파일을 가져와서 이것을 Terraform에 넣도록 하겠습니다. [docker-wordpress.tf](2022/Days/IaC/Docker-WordPress/docker-WordPress.tf)를 찾을 수 있습니다.

```
terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "2.16.0"
    }
  }
}

provider "docker" {}

variable wordpress_port {
  default = "8080"
}

resource "docker_volume" "db_data" {
  name = "db_data"
}

resource "docker_network" "wordpress_net" {
  name = "wordpress_net"
}

resource "docker_container" "db" {
  name  = "db"
  image = "mysql:5.7"
  restart = "always"
  network_mode = "wordpress_net"
  env = [
     "MYSQL_ROOT_PASSWORD=wordpress",
     "MYSQL_PASSWORD=wordpress",
     "MYSQL_USER=wordpress",
     "MYSQL_DATABASE=wordpress"
  ]
  mounts {
    type = "volume"
    target = "/var/lib/mysql"
    source = "db_data"
    }
}

resource "docker_container" "wordpress" {
  name  = "wordpress"
  image = "wordpress:latest"
  restart = "always"
  network_mode = "wordpress_net"
  env = [
    "WORDPRESS_DB_HOST=db:3306",
    "WORDPRESS_DB_USER=wordpress",
    "WORDPRESS_DB_NAME=wordpress",
    "WORDPRESS_DB_PASSWORD=wordpress"
  ]
  ports {
    internal = "80"
    external = "${var.wordpress_port}"
  }
}
```

이 파일을 다시 새 폴더에 넣은 다음 `terraform init` 명령을 실행하여 필요한 Provisioners를 내려놓습니다.

![](/2022/Days/Images/Day60_IAC4.png)

그런 다음 `terraform apply` 명령을 실행한 다음 `docker ps` 출력을 살펴보면 새로 생성된 컨테이너를 볼 수 있습니다.

![](/2022/Days/Images/Day60_IAC5.png)

이제 WordPress 프론트엔드로 이동할 수도 있습니다. 컨테이너 섹션에서 docker-compose로 이 과정을 거쳤을 때와 마찬가지로 이제 설정을 실행할 수 있으며 WordPress 게시물이 MySQL 데이터베이스에 저장됩니다.

![](/2022/Days/Images/Day60_IAC6.png)

이제 컨테이너와 Kubernetes에 대해 자세히 살펴봤는데요, 테스트용으로는 괜찮지만, 실제 웹사이트를 운영할 계획이라면 컨테이너만으로는 이 작업을 수행하지 않고 Kubernetes를 사용하여 이를 달성할 것입니다. 다음에는 Terraform과 Kubernetes를 사용하여 살펴보도록 하겠습니다.

### Provisioners

Provisioners는 어떤 것이 선언적일 수 없는 경우 이를 배포에 파싱할 수 있는 방법을 제공하기 위해 존재합니다.

다른 대안이 없고 코드에 이러한 복잡성을 추가해야 하는 경우 다음 코드 블록과 유사한 것을 실행하여 이를 수행할 수 있습니다.

```
resource "docker_container" "db" {
  # ...

  provisioner "local-exec" {
    command = "echo The server's IP address is ${self.private_ip}"
  }
}

```

원격 실행 Provisioners는 원격 리소스가 생성된 후 원격 리소스에서 스크립트를 호출합니다. 이 스크립트는 OS에 따라 다르거나 구성 관리 도구에서 래핑하는 데 사용될 수 있습니다. 이 중 일부는 Provisioners에서 다루고 있습니다.

[Provisioners에 대한 자세한 내용](https://www.terraform.io/language/resources/provisioners/syntax)

- file
- local-exec
- remote-exec
- vendor
  - ansible
  - chef
  - puppet

### 모듈

모듈은 함께 사용되는 여러 리소스를 위한 컨테이너입니다. 모듈은 동일한 디렉터리에 있는 .tf 파일 모음으로 구성됩니다.

모듈은 인프라 리소스를 분리할 수 있는 좋은 방법일 뿐만 아니라 이미 만들어진 타사 모듈을 가져올 수 있으므로 처음부터 다시 만들 필요가 없습니다.

예를 들어, 동일한 프로젝트를 사용하여 일부 VM, VPC, 보안 그룹을 구축한 다음 Kubernetes 클러스터도 구축하려는 경우 리소스를 모듈로 분할하여 리소스와 그룹화 위치를 더 잘 정의하고 싶을 것입니다.

모듈의 또 다른 장점은 이러한 모듈을 가져와 다른 프로젝트에 사용하거나 커뮤니티를 돕기 위해 공개적으로 공유할 수 있다는 것입니다.

저희는 인프라를 구성 요소로 나누고 있으며, 여기서 구성 요소는 모듈로 알려져 있습니다.

## 자료

아래에 많은 리소스를 나열했으며 이 주제는 이미 여러 번 다루어졌다고 생각합니다. 추가 리소스가 있는 경우 리소스와 함께 PR을 올리면 기꺼이 검토하여 목록에 추가해 드리겠습니다.

- [What is Infrastructure as Code? Difference of Infrastructure as Code Tools](https://www.youtube.com/watch?v=POPP2WTJ8es)
- [Terraform Tutorial | Terraform Course Overview 2021](https://www.youtube.com/watch?v=m3cKkYXl-8o)
- [Terraform explained in 15 mins | Terraform Tutorial for Beginners](https://www.youtube.com/watch?v=l5k1ai_GBDE)
- [Terraform Course - From BEGINNER to PRO!](https://www.youtube.com/watch?v=7xngnjfIlK4&list=WL&index=141&t=16s)
- [HashiCorp Terraform Associate Certification Course](https://www.youtube.com/watch?v=V4waklkBC38&list=WL&index=55&t=111s)
- [Terraform Full Course for Beginners](https://www.youtube.com/watch?v=EJ3N-hhiWv0&list=WL&index=39&t=27s)
- [KodeKloud - Terraform for DevOps Beginners + Labs: Complete Step by Step Guide!](https://www.youtube.com/watch?v=YcJ9IeukJL8&list=WL&index=16&t=11s)
- [Terraform Simple Projects](https://terraform.joshuajebaraj.com/)
- [Terraform Tutorial - The Best Project Ideas](https://www.youtube.com/watch?v=oA-pPa0vfks)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)

[Day 61](day61.md)에서 봐요!
