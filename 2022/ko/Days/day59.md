---
title: '#90DaysOfDevOps - Create a VM with Terraform & Variables - Day 59'
published: false
description: 90DaysOfDevOps - Create a VM with Terraform & Variables
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049051
---

## Terraform 및 변수를 사용하여 VM 생성하기

이 세션에서는 VirtualBox 내에서 Terraform을 사용하여 VM을 한두 개 생성해 보겠습니다. VirtualBox는 워크스테이션 가상화 옵션이므로 Terraform의 사용 사례는 아니지만, 저는 현재 36,000피트 상공에 있으며 이 정도 높이의 클라우드에 퍼블릭 클라우드 리소스를 배포한 만큼 노트북에서 로컬로 이 작업을 수행하는 것이 훨씬 더 빠릅니다.

순전히 데모 목적이지만 개념은 동일합니다. 원하는 상태 구성 코드를 만든 다음 VirtualBox 공급자에 대해 실행할 것입니다. 과거에는 여기서는 vagrant를 사용했으며 이 섹션의 시작 부분에서 vagrant와 Terraform의 차이점에 대해 다루었습니다.

### VirtualBox에서 가상 머신 생성하기

가장 먼저 할 일은 VirtualBox라는 새 폴더를 생성한 다음 VirtualBox.tf 파일을 생성하고 여기에서 리소스를 정의하는 것입니다. VirtualBox 폴더에서 VirtualBox.tf로 찾을 수 있는 아래 코드는 Virtualbox에 2개의 VM을 생성합니다.

커뮤니티 VirtualBox 제공업체에 대한 자세한 내용은 [여기](https://registry.terraform.io/providers/terra-farm/virtualbox/latest/docs/resources/vm)에서 확인할 수 있습니다.

```
terraform {
  required_providers {
    virtualbox = {
      source = "terra-farm/virtualbox"
      version = "0.2.2-alpha.1"
    }
  }
}

# 현재 공급자 자체에 대한 구성 옵션이 없습니다.

resource "virtualbox_vm" "node" {
  count     = 2
  name      = format("node-%02d", count.index + 1)
  image     = "https://app.vagrantup.com/ubuntu/boxes/bionic64/versions/20180903.0.0/providers/virtualbox.box"
  cpus      = 2
  memory    = "512 mib"

  network_adapter {
    type           = "hostonly"
    host_interface = "vboxnet1"
  }
}

output "IPAddr" {
  value = element(virtualbox_vm.node.*.network_adapter.0.ipv4_address, 1)
}

output "IPAddr_2" {
  value = element(virtualbox_vm.node.*.network_adapter.0.ipv4_address, 2)
}

```

이제 코드가 정의되었으므로 이제 폴더에서 `terraform init`을 수행하여 Virtualbox용 공급자를 다운로드할 수 있습니다.

![](/2022/Days/Images/Day59_IAC1.png)

또한 시스템에도 VirtualBox가 설치되어 있어야 합니다. 그런 다음 `terraform plan`을 실행하여 코드가 무엇을 생성하는지 확인할 수 있습니다. 이어서 `terraform apply`를 실행하면 아래 이미지에 완성된 프로세스가 표시됩니다.

![](/2022/Days/Images/Day59_IAC2.png)

이제 Virtualbox에서 2개의 가상 머신을 볼 수 있습니다.

![](/2022/Days/Images/Day59_IAC3.png)

### 구성 변경

배포에 다른 노드를 추가해 보겠습니다. 카운트 라인을 변경하여 원하는 새로운 노드 수를 표시하면 됩니다. `terraform apply`를 실행하면 아래와 같이 표시됩니다.

![](/2022/Days/Images/Day59_IAC4.png)

VirtualBox에서 완료되면 이제 3개의 노드가 실행되고 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day59_IAC5.png)

완료되면 `terraform destroy`를 사용하여 이를 지우면 머신이 제거됩니다.

![](/2022/Days/Images/Day59_IAC6.png)

### 변수 및 출력

지난 세션에서 hello-world 예제를 실행할 때 출력을 언급했습니다. 여기서 더 자세히 살펴볼 수 있습니다.

하지만 여기에도 사용할 수 있는 다른 많은 변수가 있으며, 변수를 정의하는 몇 가지 다른 방법도 있습니다.

- `terraform plan` 또는 `terraform apply` 명령을 사용하여 변수를 수동으로 입력할 수 있습니다.

- 블록 내의 .tf 파일에 변수를 정의할 수 있습니다.

- `TF_VAR_NAME` 형식을 사용하여 시스템 내에서 환경 변수를 사용할 수 있습니다.

- 저는 프로젝트 폴더에 terraform.tfvars 파일을 사용하는 것을 선호합니다.

- \*auto.tfvars 파일 옵션이 있습니다.

- 또는 `-var` 또는 `-var-file`을 사용하여 `terraform plan` 또는 `terraform apply`을 실행할 때를 정의할 수 있습니다.

아래에서 위로 올라가는 것이 변수를 정의하는 순서입니다.

또한 상태 파일에 민감한 정보가 포함될 것이라고 언급했습니다. 민감한 정보를 변수로 정의하고 이를 민감한 정보로 정의할 수 있습니다.

```
variable "some resource"  {
    description = "something important"
    type= string
    sensitive = true

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

[Day 60](day60.md)에서 봐요!
