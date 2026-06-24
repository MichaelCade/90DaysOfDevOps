---
title: '#90DaysOfDevOps - Installing & Configuring Git - Day 36'
published: false
description: 90DaysOfDevOps - Installing & Configuring Git
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048738
---

## Git 설치 및 구성

Git은 버전 관리를 위한 오픈 소스 크로스 플랫폼 도구입니다. 저처럼 우분투 또는 대부분의 Linux 환경을 사용하는 경우 이미 Git이 설치되어 있을 수 있지만 설치 및 구성 과정을 살펴보겠습니다.

시스템에 이미 git이 설치되어 있더라도 최신 버전인지 확인하는 것도 좋은 방법입니다.

### Git 설치하기

이미 언급했듯이 Git은 크로스 플랫폼이므로 Windows와 Linux를 통해 실행할 것이지만 [여기](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)에서 macOS도 찾을 수 있습니다.

[Windows](https://git-scm.com/download/win)의 경우 공식 사이트에서 설치할 수 있습니다.

Windows 컴퓨터에서 `winget`을 사용할 수도 있는데, 이것을 Windows 애플리케이션 패키지 관리자라고 생각하면 됩니다.

설치하기 전에 Windows 머신에 어떤 버전이 있는지 확인해 보겠습니다. PowerShell 창을 열고 `git --version`을 실행합니다.

![](/2022/Days/Images/Day36_Git1.png)

WSL 우분투 버전의 Git도 확인할 수 있습니다.

![](/2022/Days/Images/Day36_Git2.png)

이 글을 쓰는 시점에 최신 Windows 릴리스는 `2.35.1`이므로 업데이트해야 할 사항이 몇 가지 있는데 이를 살펴보겠습니다. 리눅스도 마찬가지일 것으로 예상합니다.

저는 최신 설치 프로그램을 다운로드하고 마법사를 실행했으며 여기에 이를 문서화할 것입니다. 주의해야 할 중요한 점은 최신 버전을 설치하기 전에 git이 이전 버전을 제거한다는 것입니다.

즉, 아래에 표시된 프로세스도 대부분 git을 사용하지 않고 설치하는 것과 동일한 프로세스입니다.

매우 간단한 설치입니다. 다운로드가 완료되면 두 번 클릭하고 시작하세요. GNU 라이선스 계약을 읽어보세요. 하지만 이 소프트웨어는 무료 오픈소스 소프트웨어라는 점을 기억하세요.

![](/2022/Days/Images/Day36_Git3.png)

이제 추가로 설치할 컴포넌트를 선택해 git과 연결할 수 있습니다. 저는 Windows에서 bash 스크립트를 실행할 수 있는 Git Bash를 항상 설치합니다.

![](/2022/Days/Images/Day36_Git4.png)

그런 다음 사용할 SSH 실행 파일을 선택할 수 있습니다. 리눅스 섹션에서 보셨을 번들 OpenSSH를 그대로 두겠습니다.

![](/2022/Days/Images/Day36_Git5.png)

그다음에는 실험적인 기능을 활성화할 수 있는데, 저는 필요하지 않으므로 활성화하지 않고 나중에 설치를 통해 돌아와서 언제든지 활성화할 수 있습니다.

![](/2022/Days/Images/Day36_Git6.png)

설치가 완료되었으므로 이제 Git Bash 및 최신 릴리스 노트를 열도록 선택할 수 있습니다.

![](/2022/Days/Images/Day36_Git7.png)

마지막 확인은 PowerShell 창에서 현재 어떤 버전의 git이 있는지 살펴보는 것입니다.

![](/2022/Days/Images/Day36_Git8.png)

매우 간단하며 이제 최신 버전을 사용하고 있습니다. 리눅스 머신에서는 업데이트가 조금 늦은 것 같아서 업데이트 과정을 안내해 드리겠습니다.

`sudo apt-get install git` 명령을 실행하기만 하면 됩니다.

![](/2022/Days/Images/Day36_Git9.png)

다음을 실행하여 소프트웨어 설치를 위한 git 저장소를 추가할 수도 있습니다.

```
sudo add-apt-repository ppa:git-core/ppa -y
sudo apt-get update
sudo apt-get install git -y
git --version
```

### Git 구성하기

처음 git을 사용할 때는 몇 가지 설정을 정의해야 합니다,

- Name
- Email
- 기본 에디터
- Line Ending

이 설정은 세 가지 수준에서 수행할 수 있습니다.

- System = 모든 사용자
- Global = 현재 사용자의 모든 리포지토리
- Local = 현재 리포지토리

예제
`git config --global user.name "Michael Cade"`
`git config --global user.email Michael.Cade@90DaysOfDevOPs.com"`
운영 체제에 따라 기본 텍스트 편집기가 결정됩니다. 다음 명령을 설정하지 않은 제 우분투 머신에서는 Nano를 사용하고 있습니다. 아래 명령은 이를 비주얼 스튜디오 코드로 변경합니다.

`git config --global core.editor "code --wait"`

이제 모든 git 구성을 보려면 다음 명령을 사용하면 됩니다.

`git config --global -e`

![](/2022/Days/Images/Day36_Git10.png)

어떤 머신에서든 이 파일의 이름은 `.gitconfig`입니다. 제 Windows 머신에서는 사용자 계정 디렉토리에서 이 파일을 찾을 수 있습니다.

![](/2022/Days/Images/Day36_Git11.png)

### Git 이론

어제 포스트에서 다른 버전 관리 유형이 있으며 이를 두 가지 유형으로 나눌 수 있다고 언급했습니다. 하나는 "클라이언트-서버"이고 다른 하나는 "분산"입니다.

### 클라이언트-서버 버전 제어

git이 등장하기 전에는 클라이언트-서버가 버전 관리를 위한 사실상 유일한 방법이었다. 그 예로 2000년에 설립된 오픈소스 버전 관리 시스템인 [Apache Subversion](https://subversion.apache.org/)을 들 수 있습니다.

이 클라이언트-서버 버전 제어 모델에서는 개발자가 서버에서 소스 코드와 실제 파일을 다운로드하는 것이 첫 번째 단계입니다. 이렇게 한다고 해서 충돌이 예방되는 것은 아니지만 충돌의 복잡성과 해결 방법이 예방됩니다.

![](/2022/Days/Images/Day36_Git12.png)

예를 들어 두 명의 개발자가 같은 파일을 작업하고 있는데 한 개발자가 새로운 변경 사항을 먼저 커밋하거나 서버에 파일을 다시 업로드한다고 가정해 보겠습니다. 두 번째 개발자가 업데이트를 하려고 할 때 충돌이 발생합니다.

![](/2022/Days/Images/Day36_Git13.png)

이때, 첫 번째 개발자의 코드 변경 사항을 가져와서 자신의 작업과 충돌하는 부분을 모두 처리한 다음, 커밋해야 합니다.

![](/2022/Days/Images/Day36_Git15.png)

### 분산 버전 제어

Git이 유일한 분산 버전 제어 시스템은 아닙니다. 하지만 사실상 거의 유일한 시스템입니다.

Git의 주요 이점은 다음과 같습니다:

- 빠름
- 스마트
- 유연함
- 안전 및 보안

클라이언트-서버 버전 제어 모델과 달리 각 개발자는 커밋 히스토리, 모든 branch 등을 의미하는 소스 리포지토리를 다운로드합니다.

![](/2022/Days/Images/Day36_Git16.png)

## 자료

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)

[Day 37](day37.md)에서 봐요!
