---
title: '#90DaysOfDevOps - Viewing, unstaging, discarding & restoring - Day 39'
published: false
description: '90DaysOfDevOps - Viewing, unstaging, discarding & restoring'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048827
---

## 보기, 스테이징 해제, 삭제 및 복원

어제에 이어서 git에서 사용할 수 있는 몇 가지 명령어와 프로젝트에서 git을 활용하는 방법에 대해 알아보겠습니다. 아직 GitHub나 다른 git 기반 서비스에 대해서는 다루지 않았지만, 지금은 로컬에서 프로젝트를 제어하는 데 도움이 되는 내용이며, 향후 해당 도구에 통합되기 시작하면 모두 유용하게 사용할 수 있을 것입니다.

### 스테이징된 변경 사항과 스테이징되지 않은 변경 사항 보기

commit하기 전에 스테이징된 코드와 스테이징되지 않은 코드를 확인하는 것이 좋습니다. 다음 `git diff --staged` 명령을 실행하면 됩니다.

![](/2022/Days/Images/Day39_Git1.png)

그러면 우리가 수행한 모든 변경 사항과 추가하거나 삭제한 모든 새 파일이 표시됩니다.

수정한 파일의 변경 사항은 `---` 또는 `+++`로 표시되며, 아래에서 방금 + 추가한 텍스트는 새로운 줄임을 의미합니다.

![](/2022/Days/Images/Day39_Git2.png)

또한 `git diff`를 실행하여 스테이징 영역과 작업 디렉토리를 비교할 수 있습니다. 새로 추가한 code.txt 파일을 변경하고 몇 줄의 텍스트를 추가하면 다음과 같습니다.

![](/2022/Days/Images/Day39_Git3.png)

그런 다음 `git diff`를 실행하면 아래와 같은 출력을 비교하여 확인할 수 있습니다.

![](/2022/Days/Images/Day39_Git4.png)

### 시각적 Diff 도구

저는 위의 방법이 혼란스럽다고 생각하기 때문에 시각적 도구를 사용하는 편이 낫습니다,

몇 가지 시각적 Diff 도구를 소개합니다:

- KDiff3
- P4Merge
- WinMerge (Windows 전용)
- VSCode

git에서 설정하려면 다음 명령 `git config --global diff.tool vscode`를 실행합니다.

위의 명령어를 실행하고 VScode를 시작할 때 몇 가지 매개 변수를 설정하겠습니다.

![](/2022/Days/Images/Day39_Git5.png)

또한 `git config --global -e`로 설정을 확인할 수 있습니다.

![](/2022/Days/Images/Day39_Git6.png)

이제 `git difftool`을 사용하여 Diff 시각화 도구를 열 수 있습니다.

![](/2022/Days/Images/Day39_Git7.png)

이제 Diff 페이지에서 VScode 에디터를 열고 두 파일을 비교하면, 아무것도 없는 상태에서 오른쪽에 코드 한 줄을 추가한 파일 하나만 수정했습니다.

![](/2022/Days/Images/Day39_Git8.png)

이 방법은 변경 사항을 추적하기가 훨씬 쉬우며, GitHub와 같은 Git 기반 서비스를 살펴볼 때 보게 될 것과 비슷한 방식입니다.

또한 `git difftool --staged`를 사용하여 commit된 파일과 스테이지를 비교할 수 있습니다.

![](/2022/Days/Images/Day39_Git9.png)

그러면 commit하기 전에 변경된 파일들을 확인할 수 있게 됩니다.

![](/2022/Days/Images/Day39_Git10.png)

저는 IDE로 VScode를 사용하고 있으며 대부분의 IDE와 마찬가지로 이미 내장되어 있는 명령어이므로 터미널에서 직접 실행해야 하는 경우는 매우 드물지만, 어떤 이유로 IDE가 설치되어 있지 않은 경우 유용합니다.

### 히스토리 보기

앞서 리포지토리에서 수행한 모든 commit을 종합적으로 볼 수 있는 `git log`에 대해 살펴봤습니다.

![](/2022/Days/Images/Day39_Git11.png)

각 commit에는 리포지토리에 고유한 16진수 문자열이 있습니다. 여기에서 작업 중인 branch와 작성자, 날짜, commit 메시지를 확인할 수 있습니다.

또한 `git log --oneline`을 사용하면 다른 `diff` 명령에서 사용할 수 있는 훨씬 더 작은 버전의 16진수 문자열을 얻을 수 있습니다. 또한 한 줄 설명 또는 commit 메시지만 있습니다.

![](/2022/Days/Images/Day39_Git12.png)

`git log --oneline --reverse`를 실행하면 페이지 상단에 첫 번째 commit이 표시됩니다.

![](/2022/Days/Images/Day39_Git13.png)

### commit 보기

commit 메시지를 볼 수 있는 것은 모범 사례를 따르고 의미 있는 commit 메시지를 추가한 경우 매우 유용하지만, commit을 검사하고 볼 수 있는 `git show` 명령도 있습니다.

우리는 `git log --oneline --reverse`를 사용하여 commit 목록을 가져올 수 있습니다. 그런 다음 이를 가져와서 `git show <commit ID>`를 실행할 수 있습니다.

![](/2022/Days/Images/Day39_Git14.png)

이 명령의 출력은 commit, 작성자 및 변경된 내용에 대한 세부 정보와 함께 아래와 같이 표시됩니다.

![](/2022/Days/Images/Day39_Git15.png)

`git show HEAD~1`을 사용할 수도 있습니다. 여기서 1은 현재 버전에서 몇 단계 뒤로 돌아가려는지 나타냅니다.

이 방법은 파일에 대한 세부 정보를 원하지만, 전체 스냅샷 디렉토리에 대한 트리의 모든 파일을 나열하려는 경우에 유용합니다. 마지막 commit에서 다시 한 스냅샷을 거슬러 올라가는 `git ls-tree HEAD~1` 명령을 사용하면 이 작업을 수행할 수 있습니다. 아래에서 두 개의 blob이 있는 것을 볼 수 있는데, blob은 파일을 나타내고 트리는 디렉토리를 나타냅니다. 이 정보에서 commit과 태그도 볼 수 있습니다.

![](/2022/Days/Images/Day39_Git16.png)

이제 위의 내용을 사용하여 `git show` 명령을 사용하여 파일(blob)의 내용을 자세히 확인할 수 있습니다.

![](/2022/Days/Images/Day39_Git17.png)

그러면 특정 버전의 파일 내용이 표시됩니다.

![](/2022/Days/Images/Day39_Git18.png)

### 파일 스테이징 해제

`git add .`를 사용했지만 아직 해당 스냅샷에 commit하고 싶지 않은 파일이 있는 경우가 있을 수 있습니다. 아래 예제에서는 스테이징 영역에 newfile.txt를 추가했지만, 이 파일을 commit할 준비가 되지 않았으므로 `git restore --staged newfile.txt`를 사용하여 `git add` 단계를 실행 취소하겠습니다.

![](/2022/Days/Images/Day39_Git19.png)

위 그림에서 수정된 파일(예: main.js)에 대해서도 동일한 작업을 수행하여 commit의 스테이징을 해제할 수 있습니다.

![](/2022/Days/Images/Day39_Git20.png)

저는 이 명령어가 90DaysOfDevOps 기간 동안 매우 유용하다는 것을 알았습니다. 다음 날을 위한 메모를 작성하고 싶지만, 공개 GitHub 리포지토리에 commit하고 push하고 싶지 않은 날을 앞두고 작업하는 경우가 종종 있기 때문입니다.

### 로컬 변경 사항 삭제하기

때때로 우리는 변경을 했지만, 그 변경이 마음에 들지 않아서 버리고 싶을 때가 있습니다. 다시 `git restore` 명령을 사용하여 스냅샷 또는 이전 버전에서 파일을 복원할 수 있습니다. 디렉토리에 대해 `git restore .`를 실행하면 스냅샷에서 모든 것을 복원할 수 있지만 추적되지 않은 파일이 여전히 존재한다는 것을 알 수 있습니다. 추적 중인 이전 파일은 newfile.txt라는 파일이 없습니다.

![](/2022/Days/Images/Day39_Git21.png)

이제 새로운 파일 txt 또는 추적되지 않은 파일을 제거합니다. 우리는 `git clean`을 사용하면 경고만 받을 수 있습니다.

![](/2022/Days/Images/Day39_Git22.png)

또는 결과를 알고 있다면 `git clean -fd`를 실행하여 모든 디렉토리를 강제로 제거할 수 있습니다.

![](/2022/Days/Images/Day39_Git23.png)

### 파일을 이전 버전으로 복원하기

앞서 언급했듯이 Git의 큰 장점 중 하나는 스냅샷에서 파일 사본을 복원할 수 있다는 점입니다(백업은 아니지만 매우 빠른 복원 방법입니다). 백업 솔루션을 사용하여 코드 사본을 다른 위치에도 저장하는 것을 추천드립니다.

예시로 디렉토리에서 가장 중요한 파일을 삭제해 보겠습니다. 디렉토리에서 이 파일을 제거하기 위해 git 명령이 아닌 Unix 기반 명령을 사용하고 있음을 알 수 있습니다.

![](/2022/Days/Images/Day39_Git24.png)

이제 작업 디렉토리에 readme.md가 없습니다. `git rm readme.md`를 사용하면 git 데이터베이스에 반영될 수 있습니다. 완전히 제거된 것을 시뮬레이션하기 위해 여기서도 삭제해 보겠습니다.

![](/2022/Days/Images/Day39_Git25.png)

이제 이 내용을 메시지와 함께 commit하고 작업 디렉토리 또는 스테이징 영역에 더 이상 아무것도 없음을 증명해 보겠습니다.

![](/2022/Days/Images/Day39_Git26.png)

실수가 있었으니 이제 이 파일을 되돌릴 필요가 있습니다!

`git undo` 명령을 사용하여 마지막 commit을 취소할 수 있지만, 오래된 commit이라면 어떻게 해야 할까요? `git log` 명령을 사용하여 commit 기록을 찾을 수 있습니다. 파일이 마지막 commit에 포함되어 있다는 것을 확인했지만, 모든 commit을 취소하고 싶지는 않습니다. 이 경우 `git restore --source=HEAD~1 README.md` 명령을 사용하여 특정 파일을 찾고 스냅샷에서 복원할 수 있습니다.

이 프로세스를 통해 파일을 작업 디렉토리에 다시 가져온 것을 볼 수 있습니다.

![](/2022/Days/Images/Day39_Git27.png)

이제 추적되지 않은 새 파일이 생겼으며 앞서 언급한 명령을 사용하여 파일과 변경 사항을 추적, 스테이징 및 commit할 수 있습니다.

### Rebase와 Merge

git 리포지토리에서 언제 rebase를 사용해야 하는지, 언제 merge를 사용해야 하는지가 가장 골치 아픈 문제인 것 같습니다.

가장 먼저 알아야 할 것은 `git rebase`와 `git merge`이 모두 같은 문제를 해결한다는 것입니다. 둘 다 한 branch의 변경 내용을 다른 branch에 통합하는 것입니다. 하지만 다른 방식으로 이 작업을 수행합니다.

새로운 feature branch로 새로운 기능을 만들어보겠습니다. main branch는 새로운 commit이 계속 추가되고 있습니다.

![](/2022/Days/Images/Day39_Git28.png)

여기서 쉬운 옵션은 `git merge feature main`을 사용하여 main branch를 feature branch에 merge하는 것입니다.

![](/2022/Days/Images/Day39_Git29.png)

merge는 비파괴적이기 때문에 간단합니다. 기존 branch는 어떤 방식으로도 변경되지 않습니다. 하지만 feature branch에 업스트림 변경 사항을 통합해야 할 때마다 관련 없는 merge commit이 발생하게 됩니다. main branch가 매우 바쁘거나 활성 상태인 경우 feature branch 히스토리를 오염시킬 수도 있습니다.

다른 옵션으로, feature branch를 main branch에 rebase하는 방법도 있습니다.

```
git checkout feature
git rebase main
```

이렇게 하면 feature branch(전체 feature branch)가 main branch에 모든 새 commit을 효과적으로 통합합니다. 그러나 merge commit을 사용하는 대신 rebase는 원래 branch에 있는 각 commit에 대해 새로운 commit을 생성하여 프로젝트 기록을 다시 작성하게 됩니다.

![](/2022/Days/Images/Day39_Git30.png)

rebase의 가장 큰 장점은 프로젝트 히스토리가 훨씬 깔끔해진다는 것입니다. 또한 불필요한 merge commit을 제거하며, 마지막 두 이미지를 비교하면 훨씬 더 깔끔한 선형 프로젝트 히스토리를 따라갈 수 있습니다.

아직 확실한 결론은 아니지만, 더 깔끔한 히스토리를 선택하는 것도 장단점이 있는데, [rebase의 황금률](https://www.atlassian.com/git/tutorials/merging-vs-rebasing#the-golden-rule-of-rebasing)을 따르지 않는다면 프로젝트 히스토리를 다시 작성하는 것은 협업 workflow에 치명적일 수 있습니다. 그리고 더 중요한 것은 rebase를 하면 merge commit이 제공하는 컨텍스트가 사라져 업스트림 변경 사항이 언제 feature에 통합되었는지 알 수 없다는 것입니다.

## 자료

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)
- [Git cheatsheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)
- [Exploring the Git command line – A getting started guide](https://veducate.co.uk/exploring-the-git-command-line/)

[Day 40](day40.md)에서 봐요!
