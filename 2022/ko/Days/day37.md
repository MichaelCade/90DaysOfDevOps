---
title: '#90DaysOfDevOps - Gitting to know Git - Day 37'
published: false
description: 90DaysOfDevOps - Gitting to know Git
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048707
---

## Gitting to know Git(Git에 대해 알아보기)

제목과 글 전체에 끔찍한 말장난이 들어가서 죄송합니다. Git을 아재 개그로 바꾼 사람이 제가 처음은 아닐 겁니다!

지난 두 번의 포스팅에서 버전 관리 시스템과 버전 관리 시스템으로서의 Git의 기본적인 workflow에 대해 배웠고, [Day 35](day35.md)에서는 시스템에 Git을 설치하고 업데이트와 구성을 했습니다. 또한 클라이언트-서버 버전 관리 시스템과 분산 버전 관리 시스템인 Git [Day 36](day36.md) 사이의 이론에 대해 조금 더 자세히 알아보았습니다.

이제 Git에서 흔히 볼 수 있는 몇 가지 명령어와 사용 사례를 살펴보겠습니다.

### git은 어디에서 도움을 받을 수 있나요?

git으로 작업을 완료하는 데 필요한 명령이 기억나지 않거나 모를 때가 있을 것입니다. 도움이 필요할 것입니다.

Google이나 검색 엔진이 도움을 찾을 때 가장 먼저 찾을 수 있는 곳일 것입니다.

두 번째는 공식 git 사이트와 문서입니다. [git-scm.com/docs](http://git-scm.com/docs) 여기에서 사용 가능한 모든 명령어에 대한 확실한 정보뿐만 아니라 다양한 리소스를 찾을 수 있습니다.

![](/2022/Days/Images/Day37_Git1.png)

터미널에 연결할 수 없는 경우에도 동일한 문서에 액세스할 수 있어 매우 유용합니다. 예를 들어 `git add` 명령어를 선택했다면 `git add --help`를 실행하면 아래에 설명서가 표시됩니다.

![](/2022/Days/Images/Day37_Git2.png)

셸에서 `git add -h`를 사용하면 사용 가능한 옵션에 대한 요약을 볼 수 있습니다.

![](/2022/Days/Images/Day37_Git3.png)

### Git을 둘러싼 오해

"Git에는 액세스 제어 기능이 없다." - 리더에게 소스 코드를 유지 관리할 수 있는 권한을 부여할 수 있습니다.

"Git은 너무 무겁다." - Git은 대규모 프로젝트의 경우 기록을 줄여주는 얕은(shallow) 저장소를 제공할 수 있습니다.

### 실제 단점

바이너리 파일에는 적합하지 않습니다. 소스 코드에는 적합하지만 실행 파일이나 동영상 등에는 적합하지 않습니다.

도구의 명령어와 기능에 대해 설명하는 데 시간을 할애해야 한다는 점이 사용자 친화적이지 않다는 것을 보여주는 대표적인 예입니다.

하지만 전반적으로 Git은 배우기는 어렵지만 사용하기는 쉽습니다.

### git 생태계

여기서는 git과 관련된 생태계를 간략하게 다루되 일부 영역에 대해 깊이 있게 다루지는 않겠지만 개략적인 수준에서 알아두는 것이 중요하다고 생각합니다.

거의 모든 최신 개발 도구는 Git을 지원합니다.

- 개발자 도구 - 이미 비주얼 스튜디오 코드에 대해 언급했지만, 서브블룸 텍스트 및 기타 텍스트 편집기 및 IDE에 대한 git 플러그인 및 통합을 찾을 수 있습니다.
- 팀 도구 - CI/CD 관점에서의 Jenkins, 메시징 프레임워크의 Slack, 프로젝트 관리 및 이슈 추적을 위한 Jira와 같은 도구에 대해서도 언급했습니다.
- 클라우드 제공업체 - 모든 대형 클라우드 제공업체는 git, Microsoft Azure, Amazon AWS 및 Google Cloud Platform을 지원합니다.
- Git 기반 서비스 - 나중에 더 자세히 다룰 GitHub, GitLab 및 BitBucket이 있습니다. 이러한 서비스는 코드를 위한 소셜 네트워크입니다!

### Git 치트시트

대부분의 명령어를 다루지는 않았지만, 온라인에서 제공되는 몇 가지 치트 시트를 살펴본 후 몇 가지 git 명령어와 그 용도를 문서화하고 싶었습니다. 이 명령어들을 모두 기억할 필요는 없으며, 더 많은 실습과 사용을 통해 최소한 git 기본 명령어 정도는 익힐 수 있을 것입니다.

[Atlassian](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)에서 가져온 것이지만, 적어두고 설명을 읽어보면 명령이 무엇인지 알 수 있을 뿐만 아니라 일상적인 작업에서 실습을 해볼 수 있는 좋은 방법입니다.

### Git 기본 사항

| Command       | Example                     | Description                                                                                                                 |
| ------------- | --------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| git init      | `git init <directory>`      | 지정한 디렉토리에 빈 git 리포지토리를 만듭니다.                                                                             |
| git clone     | `git clone <repo>`          | \<repo>에 있는 리포지토리를 로컬 머신에 복제합니다.                                                                         |
| git config    | `git config user.name`      | 현재 리포지토리 `system`, `global`, `local` 플래그의 모든 commit에 사용할 작성자 이름을 정의하여 구성 옵션을 설정합니다.    |
| git add       | `git add <directory>`       | 다음 commit을 위해 \<directory>의 모든 변경 사항을 스테이징합니다. 모든 항목에 대해 \<files>와 \<.>을 추가할 수도 있습니다. |
| git commit -m | `git commit -m "<message>"` | 스테이징된 스냅샷을 commit하고, \<message>를 사용하여 commit되는 내용을 자세히 설명합니다.                                  |
| git status    | `git status`                | 스테이징된 파일, 스테이징되지 않은 파일 및 추적되지 않은 파일을 나열합니다.                                                 |
| git log       | `git log`                   | 기본 형식을 사용하여 모든 commit 기록을 표시합니다. 이 명령에는 추가 옵션이 있습니다.                                       |
| git diff      | `git diff`                  | 인덱스와 작업 디렉토리 사이의 스테이징되지 않은 변경 내용을 표시합니다.                                                     |

### Git 변경사항 되돌리기

| Command    | Example               | Description                                                                                                                                                   |
| ---------- | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| git revert | `git revert <commit>` | \<commit>의 모든 변경 사항을 취소하는 새 commit을 만든 다음 현재 branch에 적용합니다.                                                                         |
| git reset  | `git reset <file>`    | 스테이징 영역에서 \<file>을 제거하지만 작업 디렉토리는 변경하지 않고 그대로 둡니다. 이렇게 하면 변경 내용을 덮어쓰지 않고 파일을 스테이징 해제할 수 있습니다. |
| git clean  | `git clean -n`        | 작업 디렉토리에서 어떤 파일을 제거할지 표시합니다. clean을 실행하려면 `-n` 대신 `-f`를 사용해야 합니다.                                                       |

### Git 수정 기록

| Command    | Example              | Description                                                                                                                                                              |
| ---------- | -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| git commit | `git commit --amend` | 마지막 commit을 단계적 변경 사항과 마지막 commit을 결합한 것으로 바꿉니다. 아무것도 준비하지 않은 상태에서 사용하면 마지막 commit의 메시지를 편집할 수 있습니다.         |
| git rebase | `git rebase <base>`  | 현재 branch를 \<base>로 rebase합니다. \<base>는 commit ID, branch 이름, 태그 또는 HEAD에 대한 레퍼런스가 될 수 있습니다.                                                 |
| git reflog | `git reflog`         | 로컬 리포지토리의 HEAD에 대한 변경 로그를 표시합니다. 날짜 정보를 표시하려면 --relative-date 플래그를 추가해야 하고, 모든 레퍼런스를 표시하려면 --all을 추가해야 합니다. |

### Git Branchs

| Command      | Example                    | Description                                                                                                      |
| ------------ | -------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| git branch   | `git branch`               | 리포지토리에 있는 모든 branch를 나열합니다. \<branch> 인수를 추가하여 \<branch>라는 이름의 새 branch를 만듭니다. |
| git checkout | `git checkout -b <branch>` | \<branch>라는 이름의 새 branch를 생성하고 checkout합니다. 기존 branch를 checkout하려면 -b 플래그를 지웁니다.     |
| git merge    | `git merge <branch>`       | \<branch>를 현재 branch에 merge합니다.                                                                           |

### Git 원격 리포지토리

| Command        | Example                       | Description                                                                                                                               |
| -------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| git remote add | `git remote add <name> <url>` | 원격 리포지토리에 대한 새 연결을 생성합니다. 리모트를 추가한 후 \<name>을 다른 명령에서 \<url>에 대한 바로 가기로 사용할 수 있습니다.     |
| git fetch      | `git fetch <remote> <branch>` | 리포지토리에서 특정 \<branch>를 가져옵니다. 모든 원격 레퍼런스를 가져오려면 \<branch>를 생략하세요.                                       |
| git pull       | `git pull <remote>`           | 지정된 리모트의 현재 branch 복사본을 가져와서 로컬 복사본에 즉시 merge합니다.                                                             |
| git push       | `git push <remote> <branch>`  | branch를 필요한 commit 및 오브젝트와 함께 \<remote>로 push합니다. 원격 리포지토리에 이름이 지정된 branch가 없는 경우 branch를 생성합니다. |

### Git Diff

| Command           | Example             | Description                                           |
| ----------------- | ------------------- | ----------------------------------------------------- |
| git diff HEAD     | `git diff HEAD`     | 작업 디렉토리와 마지막 commit의 diff를 표시합니다.    |
| git diff --cached | `git diff --cached` | 단계적 변경 사항과 마지막 commit의 diff를 표시합니다. |

### Git Config

| Command                                                | Example                                               | Description                                                                                                                          |
| ------------------------------------------------------ | ----------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| git config --global user.name \<name>                  | `git config --global user.name <name>`                | 현재 사용자가 모든 commit에 사용할 작성자 이름을 정의합니다.                                                                         |
| git config --global user.email \<email>                | `git config --global user.email <email>`              | 현재 사용자가 모든 commit에 사용할 작성자 이메일을 정의합니다.                                                                       |
| git config --global alias \<alias-name> \<git-command> | `git config --global 별칭 <alias-name> <git-command>` | git 명령에 대한 단축어를 만듭니다.                                                                                                   |
| git config --system core.editor \<editor>              | `git config --system core.editor <editor>`            | 컴퓨터의 모든 사용자에 대한 명령에서 사용할 텍스트 편집기를 설정합니다. \<editor> 인수는 원하는 편집기를 실행하는 명령이어야 합니다. |
| git config --global --edit                             | `git config --global --edit `                         | 수동 편집을 위해 텍스트 편집기에서 전역 구성 파일을 엽니다.                                                                          |

### Git Rebase

| Command               | Example                | Description                                                                                                                |
| --------------------- | ---------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| git rebase -i \<base> | `git rebase -i <base>` | 현재 branch를 \<base>로 rebase합니다. 편집기를 사용하여 각 commit을 새 저장소로 옮기는 방법에 대한 명령을 입력해야 합니다. |

### Git Pull

| Command                     | Example                      | Description                                                                                                                     |
| --------------------------- | ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| git pull --rebase \<remote> | `git pull --rebase <remote>` | 현재 branch의 리모트 복사본을 가져와서 로컬 복사본으로 rebase합니다. branch를 통합하기 위해 merge 대신 git rebase를 사용합니다. |

### Git Reset

| Command                    | Example                     | Description                                                                                                                                           |
| -------------------------- | --------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| git reset                  | `git reset`                 | 스테이징 영역을 가장 최근 commit과 일치하도록 reset하되 작업 디렉토리는 변경하지 않습니다.                                                            |
| git reset --hard           | `git reset --hard`          | 스테이징 영역과 작업 디렉토리를 가장 최근 commit과 일치하도록 reset하고 작업 디렉토리의 모든 변경 내용을 덮어씁니다.                                  |
| git reset \<commit>        | `git reset <commit>`        | 현재 branch 끝을 \<commit> 뒤로 이동하고 스테이징 영역을 일치하도록 reset하되 작업 디렉토리는 그대로 둡니다                                           |
| git reset --hard \<commit> | `git reset --hard <commit>` | 이전과 동일하지만 스테이징 영역과 작업 디렉토리를 모두 일치하도록 reset합니다. commit되지 않은 변경 사항과 \<commit> 이후의 모든 commit을 삭제합니다. |

### Git Push

| Command                    | Example                     | Description                                                                                                                                             |
| -------------------------- | --------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| git push \<remote> --force | `git push <remote> --force` | non-fast-forward merge가 아닌 경우에도 git push를 강제로 수행합니다. 자신이 무엇을 하고 있는지 확실하지 않으면 --force 플래그를 사용하지 말아야 합니다. |
| git push \<remote> --all   | `git push <remote> --all`   | 모든 로컬 branch를 지정한 리모트로 push 합니다.                                                                                                         |
| git push \<remote> --tags  | `git push <remote> --tags`  | branch를 push하거나 --all 플래그를 사용하면 태그가 자동으로 push되지 않습니다. tags 플래그는 모든 로컬 태그를 원격 리포지토리에 보냅니다.               |

## 자료

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)
- [Git cheatsheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)

[Day 38](day38.md)에서 봐요!
