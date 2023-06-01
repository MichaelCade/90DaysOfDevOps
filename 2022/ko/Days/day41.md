---
title: '#90DaysOfDevOps - The Open Source Workflow - Day 41'
published: false
description: 90DaysOfDevOps - The Open Source Workflow
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048806
---

## 오픈 소스 workflow

지난 7회에 걸친 Git 섹션을 통해 Git이 무엇인지, 그리고 GitHub와 같은 Git 기반 서비스가 어떻게 소스 코드 저장소를 제공할 뿐만 아니라 더 많은 커뮤니티가 함께 코드와 프로젝트를 협업할 수 있는 방법을 제공하는지 더 잘 이해하셨기를 바랍니다.

GitHub의 기본 사항을 살펴볼 때 임의의 프로젝트를 fork하고 로컬 리포지토리를 변경하는 과정을 거쳤습니다. 여기서는 한 단계 더 나아가 오픈소스 프로젝트에 기여하고자 합니다. 기여는 반드시 버그 수정이나 코딩 기능일 필요는 없으며 문서화일 수도 있다는 점을 기억하세요. 작은 도움이라도 모두 의미 있으며, 지금까지 다룬 몇 가지 git 기능을 직접 사용해 볼 수 있습니다.

## 프로젝트 fork하기

가장 먼저 해야 할 일은 우리가 기여할 수 있는 프로젝트를 찾는 것입니다. 저는 최근 [Kanister 프로젝트](https://github.com/kanisterio/kanister)에서 발표를 하고 있는데, 현재 유튜브에 올라와 있는 프레젠테이션을 프로젝트의 메인 README.md 파일에 공유하고자 합니다.

우선 프로젝트를 fork해야 합니다. 그 과정을 살펴보겠습니다. 위에서 공유한 링크로 이동하여 리포지토리를 fork하겠습니다.

![](/2022/Days/Images/Day41_Git1.png)

이제 전체 리포지토리의 복사본을 얻었습니다.

![](/2022/Days/Images/Day41_Git2.png)

참고로 README.md 파일에 나열된 원본 프레젠테이션은 이 두 개뿐이므로 프로세스를 통해 이 문제를 해결해야 합니다.

![](/2022/Days/Images/Day41_Git3.png)

## 로컬 머신에 복제

이제 fork를 로컬로 가져와서 파일에 대한 편집을 시작할 수 있습니다. 리포지토리의 코드 버튼을 사용하여 URL을 가져온 다음 리포지토리를 배치하려는 디렉토리에 `git clone url`을 사용하면 됩니다.

![](/2022/Days/Images/Day41_Git4.png)

## 변경하기

프로젝트가 로컬에 있으므로 VSCode 또는 원하는 IDE 또는 텍스트 편집기를 열어 수정 사항을 추가할 수 있습니다.

![](/2022/Days/Images/Day41_Git5.png)

READEME.md 파일은 마크다운 언어로 작성되었으며, 다른 사람의 프로젝트를 수정하고 있으므로 기존 프로젝트 형식을 따라 콘텐츠를 추가하겠습니다.

![](/2022/Days/Images/Day41_Git6.png)

## 변경 사항 테스트하기

코드 변경 후에도 애플리케이션이 계속 작동하는지 확인하려면 변경 사항을 테스트하는 것이 가장 좋은 방법이며, 문서 형식이 올바르게 지정되어 있는지 확인해야 합니다.

VSCode에서는 많은 플러그인을 추가할 수 있는데, 그중 하나가 마크다운 페이지를 미리 볼 수 있는 기능입니다.

![](/2022/Days/Images/Day41_Git7.png)

## 변경 사항을 fork된 리포지토리로 push하기

변경 사항을 Kanister 리포지토리로 직접 push할 수 있는 권한이 없으므로 이 경로를 사용해야 합니다. 이제 변경 사항에 만족하므로 이제 잘 알려진 몇 가지 git 명령을 실행할 수 있습니다.

![](/2022/Days/Images/Day41_Git8.png)

이제 GitHub로 돌아가서 변경 사항을 다시 한번 확인한 다음 마스터 프로젝트에 다시 기여합니다.

좋아 보이네요.

![](/2022/Days/Images/Day41_Git9.png)

이제 Kanister의 fork된 리포지토리 상단으로 돌아가서 kanisterio:master branch보다 commit이 1개 앞선 것을 볼 수 있습니다.

![](/2022/Days/Images/Day41_Git10.png)

다음으로 위에 강조 표시된 기여 버튼을 누릅니다. "Open Pull Request" 옵션이 표시됩니다.

![](/2022/Days/Images/Day41_Git11.png)

## Open Pull Request

다음 이미지에서 꽤 많은 일이 진행되고 있습니다. 왼쪽 위에는 이제 원본 또는 마스터 리포지토리에 있는 것을 볼 수 있습니다. 그리고 비교 대상인 원본 마스터 리포지토리와 fork된 리포지토리를 볼 수 있습니다. 이제 Create Pull Request 버튼이 있는데, 곧 다시 설명하겠습니다. 단일 commit이 있지만 변경 사항이 더 많으면 여기에 여러 개의 commit이 있을 수 있습니다. 그러면 README.md 파일에 변경 사항이 있습니다.

![](/2022/Days/Images/Day41_Git12.png)

위의 변경 사항을 검토했으며 녹색 버튼을 눌러 pull requests를 생성할 준비가 되었습니다.

그런 다음 프로젝트 관리자가 리포지토리에 pull requests 기능을 어떻게 설정했는지에 따라 관리자가 보고자 하는 내용을 알려주는 템플릿이 있을 수도 있고 없을 수도 있습니다.

여기서도 자신이 한 일에 대해 명확하고 간결하면서도 충분히 상세하게 의미 있는 설명을 작성해야 합니다. 간단한 변경 개요를 작성하고 문서화를 체크한 것을 볼 수 있습니다.

![](/2022/Days/Images/Day41_Git13.png)

## Create Pull Request

이제 pull requests를 만들 준비가 되었습니다. 페이지 상단의 "Create Pull Request"를 누르면 pull requests에 대한 요약이 표시됩니다.

![](/2022/Days/Images/Day41_Git14.png)

아래로 스크롤하면 일부 자동화가 진행 중인 것을 볼 수 있는데, 이 경우 검토가 필요하며 몇 가지 확인이 진행 중입니다. Travis CI가 진행 중이고 빌드가 시작되었음을 알 수 있으며, 업데이트를 확인하여 merge하기 전에 추가한 내용이 손상되지 않았는지 확인합니다.

![](/2022/Days/Images/Day41_Git15.png)

위 스크린샷에서 보이는 빨간색은 약간 위협적으로 보일 수 있으며, 마치 실수를 저질렀다고 생각할 수 있습니다. 걱정하지 마세요, 아무것도 망가뜨린 것이 없습니다. 여기서 가장 중요한 팁은 이 과정이 당신과 프로젝트 관리자를 돕기 위한 것이라는 것입니다. 혹시 실수를 저질렀다면 제 경험에 따르면 관리자가 연락하여 다음에 무엇을 해야 할지 조언해줄 것입니다.

이 pull requests는 이제 모든 사람이 볼 수 있도록 공개되었습니다 [added Kanister presentation/resource #1237](https://github.com/kanisterio/kanister/pull/1237).

merge 및 pull requests가 수락되기 전에 이 글을 게시하려고 합니다. 따라서 여전히 관심을 가지고 있는 사람들에게 성공적인 PR의 사진을 추가할 수 있는 작은 보상을 제공할 수 있을 것입니다.

1. 이 저장소를 귀하의 GitHub 계정으로 folk하세요.
2. 사진과 함께 텍스트를 추가해 주세요.
3. 변경 사항을 folk된 저장소에 push하세요.
4. 제가 볼 수 있고 승인할 수 있는 PR을 생성하세요.
5. 일종의 보상을 생각해 볼게요.

이것으로 Git과 GitHub에 대해 살펴본 내용을 마무리하고, 다음에는 컨테이너를 어떻게, 왜 사용하는지에 대한 큰 그림부터 시작하여 가상화에 대해 살펴보고 여기까지 오게 된 과정을 살펴볼 것입니다.

## 자료

- [Learn GitLab in 3 Hours | GitLab Complete Tutorial For Beginners](https://www.youtube.com/watch?v=8aV5AxJrHDg)
- [BitBucket Tutorials Playlist](https://www.youtube.com/watch?v=OMLh-5O6Ub8&list=PLaD4FvsFdarSyyGl3ooAm-ZyAllgw_AM5)
- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)
- [Git cheatsheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)

[Day 42](day42.md)에서 봐요!
