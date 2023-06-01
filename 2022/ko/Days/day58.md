---
title: '#90DaysOfDevOps - HashiCorp Configuration Language (HCL) - Day 58'
published: false
description: 90DaysOfDevOps - HashiCorp Configuration Language (HCL)
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048741
---

## HashiCorp 구성 언어(HCL)

Terraform으로 무언가를 만들기 시작하기 전에 HashiCorp 구성 언어(HCL)에 대해 조금 알아볼 필요가 있습니다. 지금까지 챌린지를 진행하면서 몇 가지 다른 스크립팅 및 프로그래밍 언어를 살펴봤는데, 여기에 또 다른 언어가 있습니다. [Go 프로그래밍 언어](day07.md)에 이어 [bash 스크립트](day19.md)를 다루었고, [네트워크 자동화](day27.md)와 관련해서는 파이썬도 조금 다뤄봤습니다.

이제 HashiCorp 구성 언어(HCL)를 다뤄야 하는데, 이 언어를 처음 접하는 분들에게는 다소 어렵게 느껴질 수 있지만 매우 간단하고 강력한 언어입니다.

이 섹션을 진행하면서 어떤 OS를 사용하든 시스템에서 로컬로 실행할 수 있는 예제를 사용할 것이며, 일반적으로 Terraform에서 사용하는 인프라 플랫폼은 아니지만 VirtualBox를 사용할 것입니다. 그러나 로컬에서 실행하는 것은 무료이며 이 게시물에서 원하는 것을 달성할 수 있습니다. 이 포스팅의 개념을 도커나 쿠버네티스로 확장할 수도 있습니다.

하지만 일반적으로는 퍼블릭 클라우드(AWS, Google, Microsoft Azure)뿐만 아니라 가상화 환경(VMware, Microsoft Hyper-V, Nutanix AHV)에도 인프라를 배포하는 데 Terraform을 사용하거나 사용해야 합니다. 퍼블릭 클라우드에서 Terraform을 사용하면 가상 머신 자동 배포뿐 아니라 PaaS 워크로드와 같은 모든 필수 인프라와 VPC 및 보안 그룹과 같은 모든 네트워킹 필수 자산을 생성할 수 있습니다.

Terraform에는 두 가지 중요한 측면이 있는데, 이 포스팅에서 다룰 code와 state입니다. 이 두 가지를 함께 Terraform의 핵심이라고 부를 수 있습니다. 그런 다음 우리가 대화하고 배포하고자 하는 환경이 있는데, 이는 지난 세션에서 간략히 언급했지만, AWS 공급자, Azure 공급자 등을 사용하여 실행되는 Terraform 공급자를 사용하여 실행됩니다. 수백 개가 있습니다.

### 기본 Terraform 사용법

Terraform `.tf` 파일이 어떻게 구성되는지 살펴보겠습니다. 첫 번째로 살펴볼 예제는 AWS에 리소스를 배포하는 코드이며, 이를 위해서는 시스템에 AWS CLI를 설치하고 계정에 맞게 구성해야 합니다.

### 공급자

더 복잡하게 만들 때까지는 일반적으로 `main.tf`라고 부르는 `.tf` 파일 구조의 맨 위에 있습니다. 여기서는 앞서 언급했던 공급자를 정의합니다. 보시다시피 AWS 공급자의 소스는 `hashicorp/aws`이며, 이는 공급자가 HashiCorp에서 직접 유지 관리하거나 게시했음을 의미합니다. 기본적으로 [Terraform 레지스트리](https://registry.terraform.io/)에서 사용할 수 있는 제공자를 참조하게 되며, 제공자를 작성하여 로컬에서 사용하거나 Terraform 레지스트리에 자체 게시할 수도 있습니다.

```
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}
```

프로비저닝할 AWS 리전을 결정하기 위해 여기에 리전을 추가할 수도 있는데, 이를 위해 다음을 추가할 수 있습니다:

```
provider "aws" {
  region = "ap-southeast-1" //리소스를 배포해야 하는 지역
}
```

### Terraform 리소스

- EC2, 로드 밸런서, VPC 등과 같은 하나 이상의 인프라 개체를 설명하는 Terraform 구성 파일의 또 다른 중요한 구성 요소입니다.

- 리소스 블록은 지정된 유형("aws_instance")의 리소스를 지정된 로컬 이름("90daysofdevops")으로 선언합니다.

- 리소스 유형과 이름은 함께 지정된 리소스의 식별자 역할을 합니다.

```
resource "aws_instance" "90daysofdevops" {
  ami               = data.aws_ami.instance_id.id
  instance_type     = "t2.micro"
  availability_zone = "us-west-2a"
  security_groups   = [aws_security_group.allow_web.name]
  user_data         = <<-EOF
                #! /bin/bash
                sudo yum update
                sudo yum install -y httpd
                sudo systemctl start httpd
                sudo systemctl enable httpd
                echo "
<h1>Deployed via Terraform</h1>

" | sudo tee /var/www/html/index.html
        EOF
  tags = {
    Name = "Created by Terraform"
  }
}
```

위에서 `yum` 업데이트를 실행하고 ec2 인스턴스에 `httpd`를 설치하는 것을 볼 수 있습니다.

이제 전체 main.tf 파일을 보면 다음과 같이 보일 수 있습니다.

```
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.27"
    }
  }

  required_version = ">= 0.14.9"
}

provider "aws" {
  profile = "default"
  region  = "us-west-2"
}

resource "aws_instance" "90daysofdevops" {
  ami           = "ami-830c94e3"
  instance_type = "t2.micro"
  availability_zone = "us-west-2a"
    user_data         = <<-EOF
                #! /bin/bash
                sudo yum update
                sudo yum install -y httpd
                sudo systemctl start httpd
                sudo systemctl enable httpd
                echo "
<h1>Deployed via Terraform</h1>

" | sudo tee /var/www/html/index.html
        EOF
  tags = {
    Name = "Created by Terraform"

  tags = {
    Name = "ExampleAppServerInstance"
  }
}
```

위의 코드는 AWS에서 매우 간단한 웹 서버를 ec2 인스턴스로 배포합니다. 이 코드와 이와 같은 다른 구성의 가장 큰 장점은 이 작업을 반복할 수 있고 매번 동일한 출력을 얻을 수 있다는 것입니다. 제가 코드를 엉망으로 만들었을 가능성을 제외하고는 위와 같이 사람이 개입할 여지가 없습니다.

한 번도 사용하지 않을 것 같은 아주 간단한 예제를 살펴볼 수 있지만, 어쨌든 유머러스하게 만들어 보겠습니다. 모든 훌륭한 스크립팅 및 프로그래밍 언어가 그렇듯, Hello World부터 시작해야 합니다.

```
terraform {
  # 이 모듈은 현재 Terraform 0.13.x에서만 테스트 중입니다. 그러나 더 쉽게 업그레이드할 수 있도록 다음과 같이 설정하고 있습니다.
  # 0.12.26을 최소 버전으로 설정했는데, 이 버전은 소스 URL이 있는 required_providers에 대한 지원이 추가되었기 때문입니다.
  # 0.13.x 코드와 호환됩니다.
  required_version = ">= 0.12.26"
}

# website::tag::1:: 가장 간단한 Terraform 모듈: "Hello, World!"를 출력하기만 하면 됩니다.
output "hello_world" {
  value = "Hello, 90DaysOfDevOps from Terraform"
}
```

이 파일은 IaC 폴더 안의 Hello-world 폴더에서 찾을 수 있지만, Terraform 코드를 사용하기 위해 실행해야 하는 몇 가지 명령이 있기 때문에 바로 작동하지는 않습니다.

터미널에서 main.tf가 생성된 폴더로 이동합니다. 이 저장소에서 가져올 수도 있고 위의 코드를 사용하여 새 저장소를 생성할 수도 있습니다.

해당 폴더에서 `terraform init`을 실행합니다.

Terraform 코드가 있는 모든 디렉토리에서 또는 Terraform 코드를 실행하기 전에 이 작업을 수행해야 합니다. 구성 디렉터리를 초기화하면 구성에 정의된 공급자를 다운로드하여 설치합니다. 이 경우에는 공급자가 없지만 위의 예제에서는 이 구성에 대한 AWS 공급자를 다운로드합니다.

![](/2022/Days/Images/Day58_IAC1.png)

다음 명령은 `terraform plan`입니다.

`terraform plan` 명령은 실행 계획을 생성하여 Terraform이 인프라에 적용하려는 변경 사항을 미리 볼 수 있게 해줍니다.

hello-world 예제를 통해 아래에서 간단히 볼 수 있듯이, 이것이 AWS ec2 인스턴스였다면 생성할 모든 단계가 출력되는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day58_IAC2.png)

이 시점에서 리포지토리를 초기화했고 필요한 경우 제공자를 다운로드했으며, 테스트 워크스루를 실행하여 원하는 대로 표시되는지 확인했으므로 이제 코드를 실행하고 배포할 수 있습니다.

`terraform apply`를 사용하면 이 작업을 수행할 수 있으며, 이 명령에는 안전 조치가 내장되어 있어 앞으로 일어날 일에 대한 계획 보기가 다시 제공되므로 계속할 것인지에 대한 응답을 보장합니다.

![](/2022/Days/Images/Day58_IAC3.png)

값을 입력하기 위해 yes를 입력하면 코드가 배포됩니다. 그다지 흥미롭지는 않지만, 코드에서 정의한 출력이 나오는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day58_IAC4.png)

이제 우리는 아무것도 배포하지 않았고, 아무것도 추가, 변경 또는 삭제하지 않았지만, 만약 배포했다면 위와 같이 표시된 것을 볼 수 있을 것입니다. 그러나 무언가를 배포한 후 배포한 모든 것을 제거하려면 `terraform destroy` 명령을 사용할 수 있습니다. 이 경우에도 `apply` 및 `delete` 명령 끝에 `--auto-approve`을 사용하여 수동 개입을 우회할 수 있지만 예라고 입력해야 하는 안전성이 있습니다. 하지만 이 단축키는 학습 및 테스트 시에만 사용하는 것이 좋으며, 때로는 모든 것이 빌드된 것보다 빨리 사라질 수 있습니다.

지금까지 Terraform CLI에서 다룬 명령어는 총 4가지입니다.

- `terraform init` = 프로바이더로 프로젝트 폴더 준비하기
- `terraform plan` = 코드를 기반으로 다음 명령 중에 생성 및 변경될 내용을 표시합니다.
- `terraform apply`= 코드에 정의된 리소스를 배포합니다.
- `terraform destroy` = 프로젝트에서 생성한 리소스를 파괴합니다.

코드 파일에서 두 가지 중요한 측면도 다루었습니다.

- providers = API를 통해 Terraform이 최종 플랫폼과 대화하는 방법
- resources = 우리가 코드로 배포하고자 하는 것

또 한 가지 주의해야 할 점은 `terraform init`을 실행할 때 폴더의 트리를 전후로 살펴보고 어떤 일이 발생하고 제공자와 모듈을 어디에 저장하는지 확인하는 것입니다.

### Terraform state

또한 디렉터리 내부에 생성되는 state 파일도 알아야 하는데, 이 hello-world 예제에서 state 파일은 간단합니다. 이것은 Terraform에 따라 세계를 표현하는 JSON 파일입니다. 상태는 민감한 데이터를 기꺼이 보여줄 수 있으므로 주의해야 하며, 모범 사례로 GitHub에 업로드하기 전에 `.tfstate` 파일을 `.gitignore` 폴더에 넣는 것이 좋습니다.

기본적으로 상태 파일은 프로젝트 코드와 같은 디렉터리에 있지만 옵션으로 원격으로 저장할 수도 있습니다. 프로덕션 환경에서는 S3 버킷과 같은 공유 위치가 될 가능성이 높습니다.

또 다른 옵션으로는 유료 관리형 서비스인 Terraform Cloud가 있습니다. (최대 5명의 사용자까지 무료)

원격 위치에 상태를 저장할 때 얻을 수 있는 장점은 다음과 같습니다:

- 민감한 데이터 암호화
- 협업
- 자동화
- 그러나 복잡성이 증가할 수 있습니다.

```JSON
{
  "version": 4,
  "terraform_version": "1.1.6",
  "serial": 1,
  "lineage": "a74296e7-670d-0cbb-a048-f332696ca850",
  "outputs": {
    "hello_world": {
      "value": "Hello, 90DaysOfDevOps from Terraform",
      "type": "string"
    }
  },
  "resources": []
}
```

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

[Day 59](day59.md)에서 봐요!
