---
title: '#90DaysOfDevOps - Dev workstation setup - All the pretty things - Day 20'
published: false
description: 90DaysOfDevOps - Dev workstation setup - All the pretty things
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048734
---

## 개발 워크스테이션 설정 - 예쁜 모든 것들

이 내용은 Linux 서버를 설정하는 것과 혼동되지 않도록 주의해야 하며, Linux 데스크톱의 선택성 및 유연성을 보여주고자 합니다.

저는 거의 일 년 동안 Linux 데스크톱을 사용해 왔고, 외관과 느낌 측면에서 원하는 대로 구성했습니다. Ubuntu VM과 Virtual Box를 이용하여 사용자 지정 맞춤 구성을 해봅시다.

더 쉽게 따라 할 수 있는 분들을 위해 설명 외의 부분에 대한 YouTube 동영상을 준비했습니다.

[![Click to access YouTube Video](/2022/Days/Images/Day20_YouTube.png)](https://youtu.be/jeEslAtHfKc)

기본적으로 시스템은 아래와 같이 표시됩니다.

![](/2022/Days/Images/Day20_Linux1.png)

기본 bash shell은 아래에서 볼 수 있습니다.

![](/2022/Days/Images/Day20_Linux2.png)

이 중 많은 부분이 dotfiles와 관련이 있으며, 이 시리즈의 마지막 Linux 세션에서 다루게 됐습니다.

### dotfiles

먼저 dotfiles에 대해 알아보고 싶습니다. 리눅스가 구성 파일로 구성되어 있다고 이전에 언급했습니다. 이 dotfiles는 리눅스 시스템과 응용 프로그램의 구성 파일입니다.

또한 dotfiles는 데스크톱을 꾸미고 예쁘게 만드는 데 사용되는 것뿐만 아니라 생산성을 돕는 구성 변경 및 설정이 있다고 덧붙입니다.

앞서 언급했듯이 많은 소프트웨어 프로그램들이 이 dotfiles에 구성을 저장합니다. 이 dotfiles는 기능 관리를 돕습니다.

각 dotfile은 `.`으로 시작합니다. 이름이 어떻게 지어졌는지 짐작할 수 있을 것입니다.

지금까지 우리는 shell로서 bash를 사용해 왔으므로 홈 폴더에 .bashrc와 .bash_profile이 있을 것입니다. 아래에서 시스템에 있는 몇 가지 dotfiles를 볼 수 있습니다.

![](/2022/Days/Images/Day20_Linux3.png)

우리는 shell을 변경할 것이므로, 나중에 새로운 `.zshrc` dotfile을 볼 수 있습니다.

dotfiles은 구성 파일임을 알 수 있습니다. 우리는 이것들을 사용하여 명령 프롬프트에 별칭을 추가하거나 다른 위치로 경로를 추가할 수 있습니다. 일부 사람들은 자신의 dotfiles를 공개적으로 사용할 수 있도록 게시합니다. 제 dotfiles는 [MichaelCade/dotfiles](https://github.com/MichaelCade/dotfiles)에서 찾을 수 있습니다. 여기에는 저의 맞춤형 `.zshrc` 파일, 제가 선택한 터미널인 terminator의 구성 파일이 폴더에 있고, 그리고 일부 배경 옵션이 있습니다.

### ZSH

앞에서 언급했듯이, 지금까지의 상호 작용에서 우리는 기본 shell인 Ubuntu와 함께 제공되는 bash shell을 사용했습니다. ZSH는 매우 유사하지만, bash보다 몇 가지 이점이 있습니다.

Zsh에는 대화형 탭 완성, 자동 파일 검색, 정규식 통합, 명령 범위를 정의하는 고급 단축어 및 풍부한 테마 엔진과 같은 기능이 있습니다.

우리는 `apt` 패키지 관리자를 사용하여 시스템에 zsh를 설치할 수 있습니다. bash 터미널에서 `sudo apt install zsh`를 실행해 봅시다. 이 작업은 VM 콘솔 내에서 수행할 예정이며 SSH를 통해 연결하는 대신에 진행합니다.

설치 명령이 완료되면 터미널 내에서 `zsh`를 실행할 수 있으며, 이렇게 하면 shell 구성 스크립트가 시작됩니다.

![](/2022/Days/Images/Day20_Linux4.png)

위 질문에 대해 `1`을 선택했고, 이제 더 많은 옵션을 볼 수 있습니다.

![](/2022/Days/Images/Day20_Linux5.png)

이 메뉴에서 몇 가지 기본 제공 편집을 통해 필요에 따라 ZSH를 구성할 수 있음을 알 수 있습니다.

`0`으로 종료한 다음 `ls -al | grep .zshrc`를 사용하면 새로운 구성 파일이 있는 것을 확인할 수 있습니다.

이제 우리는 터미널을 열 때마다 zsh를 기본 shell로 사용하려고 합니다. 이를 위해 shell을 변경하는 `chsh -s $(which zsh)` 명령을 실행할 수 있습니다. 그런 다음 로그아웃하고 다시 로그인하여 변경 사항이 적용되도록 해야 합니다.

로그인 후 터미널을 열면 아래와 같이 보일 것입니다. 또한 `which $SHELL`을 실행하여 shell이 변경되었음을 확인할 수 있습니다.

![](/2022/Days/Images/Day20_Linux6.png)

일반적으로 저는 각 Ubuntu 데스크톱에서 이 단계를 수행하며, 이 외에도 zsh shell이 bash보다 약간 빠르다고 느낍니다.

### OhMyZSH

다음으로 터미널 내에서 조금 더 나은 모습과 추가 기능을 제공하고자 합니다.

OhMyZSH는 zsh 구성을 관리하기 위한 무료 오픈 소스 프레임워크입니다. 플러그인, 테마 등 많은 것들이 있어 zsh shell과 상호 작용하는 것이 훨씬 재밌습니다.

[ohmyzsh](https://ohmyz.sh/)에 대한 자세한 정보를 확인해보세요.

Oh My ZSH를 설치합시다. `curl`, `wget` 또는 `fetch` 중에서 선택할 수 있는데, 첫 번째 두 가지는 시스템에서 사용 가능하지만, `curl`로 시작하겠습니다.

`sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"`

위 명령을 실행하면 아래와 같은 출력을 볼 수 있습니다.

![](/2022/Days/Images/Day20_Linux7.png)

이제 우리의 경험을 위한 테마를 시작할 수 있습니다. Oh My ZSH에는 100개 이상의 테마가 포함되어 있지만, 저의 모든 애플리케이션과 모든 것에 대한 기본값은 Dracula 테마입니다.

또한 Oh My ZSH를 사용할 때 다음 두 가지 플러그인을 반드시 사용해야 한다고 강조하고 싶습니다.

`git clone https://github.com/zsh-users/zsh-autosuggestions.git $ZSH_CUSTOM/plugins/zsh-autosuggestions`

`git clone https://github.com/zsh-users/zsh-syntax-highlighting.git $ZSH_CUSTOM/plugins/zsh-syntax-highlighting`

`nano ~/.zshrc`

플러그인을 편집하여 이제 `plugins=(git zsh-autosuggestions zsh-syntax-highlighting)`를 포함하도록 합니다.

### Gnome Extensions

저는 Gnome Extensions도 사용하며, 특히 아래 목록을 사용합니다.

[Gnome extensions](https://extensions.gnome.org)

    - Caffeine
    - CPU Power Manager
    - Dash to Dock
    - Desktop Icons
    - User Themes

## 소프트웨어 설치

`apt`를 사용하여 제 컴퓨터에 설치한 프로그램 목록입니다.

    - VSCode
    - azure-cli
    - containerd.io
    - docker
    - docker-ce
    - google-cloud-sdk
    - insomnia
    - packer
    - terminator
    - terraform
    - vagrant

### Dracula 테마

[Dracula Theme](https://draculatheme.com/)는 제가 유일하게 사용하는 테마입니다. 깔끔하고 명확한 모습이며 모든 것이 훌륭해 보입니다. 이 사이트에서는 컴퓨터에서 사용하는 다양한 프로그램을 지원합니다.

위 링크에서 zsh를 검색하면 최소 두 가지 옵션을 찾을 수 있습니다.

수동 설치 또는 git을 사용하여 설치할 수 있는 지침을 따릅니다. 그런 다음 아래와 같이 `.zshrc` 구성 파일을 편집해야 합니다.

![](/2022/Days/Images/Day20_Linux8.png)

다음으로 [Gnome Terminal Dracula theme](https://draculatheme.com/gnome-terminal)를 원하게 될 텐데, 여기에 모든 지침이 나와 있습니다.

모든 단계를 문서화하는 데 오랜 시간이 걸릴 것이므로, 저는 과정을 동영상으로 제공했습니다.(아래 이미지를 클릭하세요)

[![](/2022/Days/Images/Day20_YouTube.png)](https://youtu.be/jeEslAtHfKc)

여기까지 왔다면, 이제 #90DaysOfDevOps의 Linux 섹션을 완료했습니다. 다시 한번, 피드백과 추가 자료에 대해서는 언제든지 환영입니다.

여기에 적는 것보다 동영상을 통해 많은 단계를 보여드리는 것이 더 쉬울거라고 생각했는데, 이에 대해 어떻게 생각하시나요? 저는 이 글을 다시 읽어보고 가능하면 동영상 연습을 만들어서 우리가 다룬 내용들을 더 잘 설명하고 보여줄 수 있도록 하고자 합니다. 어떻게 생각하시나요?

## 자료

- [Bash in 100 seconds](https://www.youtube.com/watch?v=I4EWvMFj37g)
- [Bash script with practical examples - Full Course](https://www.youtube.com/watch?v=TPRSJbtfK4M)
- [Client SSH GUI - Remmina](https://remmina.org/)
- [The Beginner's guide to SSH](https://www.youtube.com/watch?v=2QXkrLVsRmk)
- [Vim in 100 Seconds](https://www.youtube.com/watch?v=-txKSRn0qeA)
- [Vim tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I)
- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

내일부터 7일간 네트워킹에 대해 알아보는 시간을 가지며, 데브옵스와 관련된 네트워킹에 대한 기초적인 지식과 이해를 쌓을 것입니다.

[Day 21](day21.md)에서 봐요!
