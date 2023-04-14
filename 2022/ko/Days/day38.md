---
title: '#90DaysOfDevOps - Staging & Changing - Day 38'
published: false
description: 90DaysOfDevOps - Staging & Changing
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049042
---

## 스테이징 및 변경

이미 몇 가지 기본 사항을 다뤘지만, 이렇게 연습을 해보면 어떻게 그리고 왜 이런 방식으로 작업하는지 더 잘 이해하고 배울 수 있습니다. GitHub와 같은 git 기반 서비스에 들어가기 전에 로컬 워크스테이션에서 활용할 수 있는 git의 강력한 기능을 살펴봅시다.

git 세션을 시작할 때 만든 프로젝트 폴더를 가지고 git으로 할 수 있는 몇 가지 간단한 예제를 보여드리겠습니다. 로컬 머신에 폴더를 생성하고 `git init` 명령으로 초기화했습니다.

![](/2022/Days/Images/Day38_Git1.png)

이제 폴더를 초기화했으므로 디렉토리에 숨겨진 폴더가 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day38_Git2.png)

이 폴더에는 branch 및 커밋에 관한 정보뿐만 아니라 git 저장소의 세부 정보가 저장됩니다.

### 스테이징 파일

그런 다음 빈 폴더에서 작업을 시작하고 작업 첫날에 소스 코드를 추가할 수 있습니다. README.md 파일을 생성하면 디렉토리에서 해당 파일을 볼 수 있고, 다음으로 `git status`를 확인하면 새 README.md 파일이 보이지만, 아직 그 파일을 커밋하지는 않은 상태입니다.

![](/2022/Days/Images/Day38_Git3.png)

이제 `git add README.md` 명령으로 README.md 파일을 스테이징하면 이전에 없던 변경 사항과 초록색으로 표시되는 새 파일을 커밋할 수 있는 상태가 됩니다.

![](/2022/Days/Images/Day38_Git4.png)

다음으로 프로젝트의 첫 번째 커밋 또는 첫 번째 스냅샷인 이 파일을 커밋하고 싶습니다. 각 커밋에 대해 변경된 내용을 쉽게 확인할 수 있도록 `git commit -m "의미 있는 메시지"` 명령을 사용하여 이 작업을 수행할 수 있습니다. 또한 노란색 십자가가 이제 녹색 체크 표시로 바뀐 것을 확인할 수 있습니다. 이것은 제가 터미널에서 사용하는 테마로, 리눅스 섹션에서 다룬 내용입니다.

![](/2022/Days/Images/Day38_Git5.png)

### 변경 사항 커밋

파일을 더 추가하거나 디렉토리에 있는 파일을 변경하고 싶을 때가 많습니다. 위에서 이미 첫 번째 커밋을 했습니다. 하지만 이제 더 많은 세부 사항으로 더 많은 파일을 추가하겠습니다.

이전 프로세스를 반복하여 파일을 만들거나 수정한 다음 `git add .`를 실행하여 모든 파일을 스테이징 영역에 추가한 다음 `git commit -m "의미 있는 메시지"`를 실행하면 정상적으로 작동할 수 있습니다. 그러나 커밋에 변경된 내용에 대한 의미 있는 메시지를 제공하려면 `git commit -m "음, 일부 코드가 작동하지 않아서 변경했으며, 이를 수정할 때 모든 사람이 사용자 경험에 대해 알 수 있도록 README.md에 새로운 내용을 추가했습니다."`와 같이 설명적인 내용을 작성하는 것도 가능하지만 여기서는 텍스트 편집기를 사용하여 추가하는 것이 더 바람직할 수 있습니다.

`git add`를 실행한 후 `git commit`을 실행하면 기본 텍스트 편집기(여기서는 nano)가 열립니다. 다음은 파일에 몇 가지 변경 사항을 추가하기 위해 수행한 단계이며, `git status`를 실행하여 스테이징된 항목과 스테이징되지 않은 항목을 표시합니다. 그런 다음 `git add`를 사용하여 파일을 스테이징 영역에 추가한 다음 `git commit`을 실행하여 nano를 열었습니다.

![](/2022/Days/Images/Day38_Git6.png)

Nano가 열리면 설명을 추가한 다음 파일을 저장할 수 있습니다.

![](/2022/Days/Images/Day38_Git7.png)

### 커밋 모범 사례

커밋 시기와 커밋 횟수 사이에는 균형이 필요합니다. 프로젝트가 끝날 때까지 기다렸다가 커밋하는 것은 바람직하지 않으며, 각 커밋은 의미가 있어야 하고 서로 관련 없는 작업들은 함께 커밋해서는 안 됩니다. 버그 수정과 오타를 해결한 경우, 두 가지 커밋을 분리하여 커밋하는 것이 좋습니다.

커밋 메시지에 의미를 부여하세요.

문구 측면에서, 팀이나 본인이 각 커밋에 대해 동일한 문구를 사용해야 합니다.

### 스테이징 영역 건너뛰기

변경 사항을 커밋하기 전에 항상 스테이징해야 할까요?

정답은 '예'이지만, 롤백할 스냅샷이 필요하지 않다는 것을 100% 확신해야 하며, 이는 위험할 수 있습니다.

![](/2022/Days/Images/Day38_Git8.png)

### 파일 제거

프로젝트에서 파일을 제거하는 것은 어떨까요? 디렉토리에 커밋했지만 이제 프로젝트에 더 이상 필요하지 않거나 사용하지 않는 다른 파일이 있는 경우, 파일을 제거해야 합니다.

디렉토리에서 파일을 제거해도 git은 여전히 이 파일을 인식하므로 리포지토리에서도 파일을 제거해야 합니다. 아래에서 이를 위한 workflow를 확인할 수 있습니다.

![](/2022/Days/Images/Day38_Git9.png)

이동하는 파일과 폴더가 많은 대규모 프로젝트의 경우 기억하거나 처리하는 것이 다소 번거로울 수 있습니다. 이 작업은 `git rm oldcode.ps1` 명령 하나로 수행할 수 있습니다.

![](/2022/Days/Images/Day38_Git10.png)

### 파일 이름 바꾸기 또는 이동

운영 체제 내에서 파일 이름을 바꾸고 이동할 수 있습니다. 프로젝트에서 때때로 이 작업을 해야 할 때가 있을 것입니다. 2단계 프로세스가 있지만 제거와 유사하게, OS에서 파일을 변경한 다음 스테이징 영역이나 파일이 올바르게 추가되었는지 수정하고 확인해야 합니다. 단계는 다음과 같습니다:

![](/2022/Days/Images/Day38_Git11.png)

그러나 운영 체제에서 파일을 제거한 다음 git 리포지토리에서 파일을 제거하는 것과 마찬가지로 git 명령을 사용하여 이름 변경을 할 수 있습니다.

![](/2022/Days/Images/Day38_Git12.png)

### 파일 무시하기

프로젝트 내 로컬에서만 사용 중이거나 전체 프로젝트에 공유되면 공간 낭비일 수 있는 파일이나 폴더를 무시해야 하는 경우가 있는데, 그 좋은 예로 로그를 들 수 있습니다. 또한 공개적으로 또는 팀 간에 공유하고 싶지 않은 비밀에 대해서도 이 기능을 사용할 수 있다고 생각합니다.

프로젝트 디렉토리의 `.gitignore` 파일에 폴더나 파일을 추가하여 파일을 무시할 수 있습니다.

![](/2022/Days/Images/Day38_Git13.png)

`.gitignore` 파일을 열면 logs/ 디렉토리가 있는 것을 확인할 수 있습니다. 여기에 무시할 파일과 폴더를 추가할 수도 있습니다.

![](/2022/Days/Images/Day38_Git14.png)

이제 `git status`를 보고 어떤 일이 일어났는지 확인할 수 있습니다.

![](/2022/Days/Images/Day38_Git15.png)

로그 폴더를 이미 공유했지만, 나중에 공유하면 안 된다는 것을 깨달았을 때 파일과 폴더를 무시해야 하는 경우도 있습니다. 이전에 추적한 폴더가 있지만 이제 무시하고 싶은 경우 `git rm --cached`를 사용하여 스테이징 영역에서 파일과 폴더를 제거해야 합니다.

### 간략한 status

준비 영역에 무엇이 있고 무엇이 없는지 파악하기 위해 `git status`를 많이 사용해왔는데, 이 명령은 매우 포괄적이고 세부적인 내용을 담고 있습니다. 대부분의 경우 무엇이 수정되었는지 또는 무엇이 새로운 것인지 알고 싶을 것입니다. 이 세부 사항에 대한 간단한 상태를 확인하려면 `git status -s`를 사용할 수 있습니다. 저는 보통 시스템에서 단축어로 설정하여 더 자세한 명령 대신 `git status -s`만 사용하도록 합니다.

![](/2022/Days/Images/Day38_Git16.png)

내일 포스트에서는 이러한 일반적인 git 명령에 대한 간략한 예제를 계속 살펴보겠습니다.

## 자료

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)
- [Git cheatsheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)

[Day 39](day39.md)에서 봐요!
