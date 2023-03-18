---
title: '#90DaysOfDevOps - The Big Picture: DevOps and Linux - Day 14'
published: false
description: 90DaysOfDevOps - The Big Picture DevOps and Linux
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049033
---

## 큰 그림: 데브옵스와 Linux

Linux와 데브옵스는 유사한 문화와 관점을 공유하며, 둘 다 커스터마이징과 확장성에 중점을 둡니다. 이 두 가지 측면은 특히 데브옵스에서 중요한 역할을 합니다.

특히 소프트웨어 개발이나 인프라 관리와 관련 기술들이 Linux에서 개발되었습니다.

또한 많은 오픈 소스 프로젝트, 특히 DevOps 도구는 Linux에서 바로 실행되도록 설계되었습니다.

DevOps나 운영 역할 관점에서는 대부분 Linux를 다룰 가능성이 높습니다. WinOps를 다루는 일도 있지만, 보통은 Linux 서버를 관리하고 배포하는 일이 많을 것입니다.

저는 몇 년 동안 Linux를 사용해왔지만, 메인 데스크탑 컴퓨터는 주로 macOS나 Windows였습니다. 하지만 클라우드 네이티브 역할로 전환한 후에는 노트북에서 Linux를 기본 운영 체제로 사용하기로 결정했습니다. 일부 오디오 및 비디오 장비는 Linux와 호환되지 않기 때문에 업무 관련 애플리케이션은 여전히 Windows가 필요하지만, 앞으로 7일간은 개념을 더 잘 이해하기 위해 데스크탑에서 Linux를 사용하겠습니다.

## 시작하기

더 쉬운 방법이 있기 때문에 여러분도 저와 똑같이 하라고 권하고 싶지는 않지만, 저처럼 단계를 밟으면 Linux를 더 빨리 배울 수 있다는 걸 알려드리고 싶습니다.

이번 7일 동안, 저는 Windows에서 VirtualBox를 사용해 가상 머신을 띄울 계획입니다. 데스크탑 버전의 Linux 배포판도 함께 사용할 예정이지만, 여러분이 관리하게 될 대부분의 Linux 서버는 GUI가 없이 shell 기반으로 운영될 것입니다. 하지만, 앞서 90일 동안 학습한 많은 도구들이 Linux에서 시작되었다는 것을 언급했듯이, Linux 데스크탑을 실행해 보는 것도 강력히 추천합니다.

이 글의 나머지 부분에서는 Virtual Box 환경에서 우분투 데스크탑 가상 머신 실행에 초점을 맞출 것입니다. [Virtual Box](https://www.virtualbox.org/)를 다운로드하고 링크된 사이트에서 최신 [Ubuntu ISO](https://ubuntu.com/download)를 가져와서 데스크탑 환경을 구축할 수도 있지만, 그렇게 하는 것은 데브옵스답지 않겠죠?

Linux 배포판을 선택한 또 다른 이유는 대부분 무료이며 오픈 소스이기 때문입니다. 또한, 모바일 디바이스와 엔터프라이즈 RedHat Enterprise 서버를 제외하고 가장 널리 배포된 배포판인 우분투를 선택했습니다. 제가 틀렸을 수도 있지만 CentOS의 역사를 살펴보면 우분투가 목록에서 높은 순위에 있으며 매우 간단한 편이라고 생각합니다.

## HashiCorp Vagrant 소개

Vagrant는 CLI 유틸리티로 가상 머신의 수명 주기를 관리합니다. vSphere, Hyper-v, Virtual Box, Docker 등 다양한 플랫폼에서 Vagrant를 사용하여 가상 머신을 생성하고 제거할 수 있습니다. 다른 대안도 있지만, 저는 Virtual Box를 사용할 것입니다.

먼저 해야 할 일은 Vagrant를 컴퓨터에 설치하는 것입니다. [HashiCorp Vagrant](https://www.vagrantup.com/downloads) 다운로드 페이지에서는 사용 가능한 모든 운영 체제를 볼 수 있습니다. 저는 Windows를 사용 중이기 때문에 시스템용 바이너리를 다운로드하여 설치하였습니다.

다음으로 [Virtual Box](https://www.virtualbox.org/wiki/Downloads)를 설치해야 합니다. 이 프로그램은 다양한 운영체제에서 실행할 수 있으며, 윈도우, 맥 OS, 또는 리눅스를 실행할 경우 필요합니다.

두 설치 모두 매우 간단하며, 훌륭한 커뮤니티도 있어서 문제가 발생하면 언제든지 연락하시면 도움을 드릴 수 있습니다.

## 첫 번째 VAGRANTFILE

VAGRANTFILE은 배포하려는 가상 머신의 유형을 정의하고, 해당 머신의 구성과 프로비저닝을 설정합니다.

VAGRANTFILE을 저장하고 정리할 때, 저는 워크스페이스의 폴더에 파일을 저장합니다. 제 시스템에서의 예시를 아래에서 확인하실 수 있습니다. 이 방법을 따라 하면 다른 시스템으로 쉽게 전환할 수 있으며, Linux 데스크탑을 "distro hopping"_(리눅스 사용자들이 사용하는 운영체제를 자주 바꾸는 행위 - 옮긴이)_ 하는 "rabbit hole"_(위험을 대비한 대안을 여러 개 준비하는 행위 - 옮긴이)_ 에도 적합합니다.

![](/2022/Days/Images/Day14_Linux1.png)

VAGRANTFILE을 살펴보고 우리가 현재 만들고 있는 것을 확인해 봅시다.

```
Vagrant.configure("2") do |config|

  config.vm.box = "chenhan/ubuntu-desktop-20.04"

  config.vm.provider :virtualbox do |v|

   v.memory  = 8096

   v.cpus    = 4

   v.customize ["modifyvm", :id, "--vram", "128"]

end

end
```

이건 매우 간단한 VAGRANTFILE입니다. 특정 "box"를 원한다는 뜻이며, 이 box는 여러분이 찾고 있는 시스템의 공개 이미지 또는 비공개 빌드일 가능성이 높습니다. 공개적으로 사용 가능한 "box"들의 목록은 [public catalogue of Vagrant boxes](https://app.vagrantup.com/boxes/search)에서 찾을 수 있습니다.

다음 줄에서는 `VirtualBox`라는 특정한 공급자를 사용하고자 합니다. 또한 컴퓨터의 메모리를 `8GB`로, CPU 수를 `4`로 정의합니다. 제 경험상, 디스플레이 문제가 발생하면 비디오 메모리를 추가하는 것이 좋습니다. 이렇게 하면 비디오 메모리를 `128MB`까지 늘릴 수 있지만 시스템에 따라 다를 수 있습니다.

```
v.customize ["modifyvm", :id, "--vram", ""]
```

이 vagrant 파일은 [Linux Folder](/2022/Days/Linux/VAGRANTFILE)에서 확인할 수 있습니다.

## Linux 데스크탑 프로비저닝

제 첫 번째 Linux 데스크탑을 시작하고 실행할 준비가 된 상태입니다. 저는 Windows 운영체제에서 PowerShell을 사용하고 있습니다. 프로젝트 폴더로 이동하여 VAGRANTFILE이 위치한 경로로 이동해주세요. 그곳에서 `vagrant up` 명령어를 입력하면 아래와 같은 내용이 표시됩니다.

![](/2022/Days/Images/Day14_Linux2.png)

추가해야 할 사항은 이 가상 머신에서 네트워크가 `NAT`로 설정된다는 것입니다. 현재 단계에서는 NAT에 대해 알 필요가 없으며, 네트워킹 세션에서 이에 대해 자세히 설명할 계획입니다. 가상 머신을 홈 네트워크에 연결하는 가장 간단한 방법이기도 하며, Virtual Box의 기본 네트워킹 모드입니다. 추가적인 정보는 [Virtual Box documentation](https://www.virtualbox.org/manual/ch06.html#network_nat)에서 확인할 수 있습니다.

`vagrant up`이 완료되면 `vagrant ssh`를 사용하여 새 VM의 터미널로 바로 이동할 수 있습니다.

![](/2022/Days/Images/Day14_Linux3.png)

앞으로 며칠 동안 이곳에서 대부분의 작업을 진행할 예정입니다. 또한 개발자 워크스테이션에 몇 가지 사용자 지정 기능을 추가하여 일상적인 작업을 훨씬 더 간편하게 만들 계획입니다. 비표준 터미널이 없다면, 스스로를 DevOps 전문가로 여길 수 있을까요?

확인을 위해 가상 머신을 선택하면 로그인 프롬프트가 표시되어야 합니다.

![](/2022/Days/Images/Day14_Linux4.png)

여기까지 읽으셨는데 "Username과 Password가 무엇인가요?"라는 질문이 계속 나오신다면?

- Username = vagrant

- Password = vagrant

내일은 터미널에서의 몇 가지 명령어와 그 명령어들이 하는 일에 대해 알아보도록 하겠습니다.

## 자료

- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

앞으로 더 많은 자원을 확보할 예정이며, Go 언어 자료와 마찬가지로 일반적으로 무료 콘텐츠를 제공하여 모두가 참여하고 배울 수 있도록 할 것입니다.

다음 글에서는 자주 사용할 수 있는 Linux 환경의 명령어에 대해 살펴볼 것입니다.

[Day 15](day15.md)에서 봐요!
