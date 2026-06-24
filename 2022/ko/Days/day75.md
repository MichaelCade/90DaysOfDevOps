---
title: '#90DaysOfDevOps - GitHub Actions Overview - Day 75'
published: false
description: 90DaysOfDevOps - GitHub Actions Overview
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049070
---

## GitHub Actions 개요

이 섹션에서는 방금 시간을 할애한 것과는 다른 접근 방식을 살펴보고자 합니다. 이 세션에서 집중적으로 다룰 내용은 GitHub Actions입니다.

GitHub Actions는 파이프라인의 다른 작업들 사이에서 빌드, 테스트 및 배포할 수 있는 CI/CD 플랫폼입니다. 이 플랫폼은 GitHub 리포지토리를 대상으로 빌드하고 테스트하는 Wolkflow의 개념을 가지고 있습니다. 또한 GitHub Actions를 사용하여 리포지토리 내에서 발생하는 Event를 기반으로 다른 Wolkflow를 구동할 수도 있습니다.

### Wolkflow

전반적으로 GitHub Actions에서는 작업을 **Wolkflow**라고 부릅니다.

- **Wolkflow**는 구성 가능한 자동화된 프로세스입니다.
- YAML 파일로 정의됩니다.
- 하나 이상의 **Job**을 포함하고 실행합니다.
- 리포지토리의 **Event**에 의해 트리거될 때 실행되거나 수동으로 실행할 수 있습니다.
- 리포지토리당 여러 Wolkflow를 사용할 수 있습니다.
- **Wolkflow**에는 **Job**과 해당 **Job**을 달성하기 위한 **Step**이 포함됩니다.
- **Wolkflow** 내에는 **Wolkflow**가 실행되는 **Runner**도 있습니다.

예를 들어 PR을 빌드하고 테스트하는 **Wolkflow**, 릴리스가 만들어질 때마다 애플리케이션을 배포하는 **Wolkflow**, 누군가 새 issue를 열 때마다 레이블을 추가하는 또 다른 **Wolkflow**가 있을 수 있습니다.

### Event

Event는 Wolkflow를 실행하도록 트리거하는 리포지토리의 특정 이벤트입니다.

### Job

Job은 Runner에서 실행되는 Wolkflow의 Step 집합입니다.

### Step

Job 내의 각 Ste은 실행되는 shell 스크립트 또는 Action이 될 수 있습니다. Step은 순서대로 실행되며 서로 종속됩니다.

### Action

자주 반복되는 작업에는 반복 가능한 사용자 지정 애플리케이션이 사용됩니다.

### Runner

Runner는 Wolkflow를 실행하는 서버로, 각 Runner는 한 번에 하나의 작업을 실행합니다. GitHub Actions는 Ubuntu Linux, Microsoft Windows 및 macOS Runner를 실행할 수 있는 기능을 제공합니다. 특정 OS 또는 하드웨어에서 직접 호스팅할 수도 있습니다.

아래에서 Wolkflow를 트리거하는 Event > Wolkflow가 두 개의 Job으로 구성 > 작업 내에 단계와 액션이 있는 모습을 볼 수 있습니다.

아래에서 Wolkflow를 트리거하는 Event가 있고 > Wolkflow가 두 개의 Job으로 구성되어 있으며 > Job 내에 Step가 있고 > Action이 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day75_CICD1.png)

### YAML

실제 사용 사례를 살펴보기 전에 위의 이미지를 예제 YAML 파일 형태로 간단히 살펴봅시다.

주석을 추가하여 YAML Wolkflow의 구성 요소를 찾을 수 있도록 했습니다.

```Yaml
#Workflow
name: 90DaysOfDevOps
#Event
on: [push]
#Jobs
jobs:
  check-bats-version:
    #Runner
    runs-on: ubuntu-latest
    #Steps
    steps:
        #Actions
      - uses: actions/checkout@v2
      - uses: actions/setup-node@v2
        with:
          node-version: '14'
      - run: npm install -g bats
      - run: bats -v
```

### GitHub Actions 실습하기

코드 빌드, 테스트, 배포 및 그 이후의 지속적인 단계와 관련하여 CI/CD 요구 사항을 충족할 수 있는 GitHub Actions에는 많은 옵션이 있다고 생각합니다.

GitHub Actions를 사용할 수 있는 많은 옵션과 기타 자동화된 작업을 볼 수 있습니다.

### 코드 lint에 GitHub 액션 사용하기

한 가지 옵션은 리포지토리 내에서 코드를 깨끗하고 깔끔하게 정리하는 것입니다. 이것이 첫 번째 예제 데모가 될 것입니다.

이 섹션의 리소스 중 하나에 링크된 몇 가지 예제 코드를 사용하여 `GitHub/super-linter`를 사용하여 코드를 확인하겠습니다.

```Yaml
name: Super-Linter

on: push

jobs:
  super-lint:
    name: Lint code base
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v2

      - name: Run Super-Linter
        uses: github/super-linter@v3
        env:
          DEFAULT_BRANCH: main
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

**github/super-linter**
위에서 단계 중 하나에 GitHub/super-linter라는 Action이 있으며 이는 커뮤니티에서 이미 작성된 Step을 참조하고 있음을 알 수 있습니다. 이에 대한 자세한 내용은 [Super-Linter](https://github.com/github/super-linter)에서 확인할 수 있습니다.

"이 리포지토리는 super-linter를 실행하기 위한 Github Actions를 위한 저장소입니다. 소스 코드의 유효성을 검사하는 데 도움이 되도록 bash로 작성된 다양한 linter의 간단한 조합입니다."

또한 위의 코드 스니펫에 GITHUB_TOKEN이 언급되어 있어서 이것이 왜 필요한지, 어떤 용도로 사용되는지 궁금했습니다.

"참고: Wolkflow에서 환경 변수 `GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}`을 전달하면 GitHub Super-Linter가 PR의 Checks 섹션에 각 linter 실행 상태를 표시합니다. 이 기능이 없으면 전체 실행의 전체 상태만 볼 수 있습니다. **GitHub Secret은 GitHub에서 자동으로 설정하므로 설정할 필요가 없으며, Action으로 전달하기만 하면 됩니다.**"

굵은 텍스트는 이 단계에서 주목해야 할 중요한 부분입니다. 우리는 이 기능을 사용하지만, 리포지토리 내에서 환경 변수를 설정할 필요는 없습니다.

테스트할 리포지토리는 Jenkins 데모에서 사용한 리포지토리를 사용하겠습니다. [Jenkins-HelloWorld](https://github.com/MichaelCade/Jenkins-HelloWorld).

다음은 Jenkins 세션에 남겨둔 저장소입니다.

![](/2022/Days/Images/Day75_CICD2.png)

이 기능을 활용하려면 위의 Actions 탭을 사용하여 곧 다룰 마켓플레이스에서 선택하거나 위의 super-linter 코드를 사용하여 파일을 생성할 수 있는데, 직접 생성하려면 리포지토리에 이 위치에 새 파일을 만들어야 합니다. `.github/workflows/workflow_name`은 분명히 여러분이 알아볼 수 있는 유용한 이름이어야 하며, 여기에는 리포지토리에 대해 다양한 작업과 작업을 수행하는 다양한 Wolkflow가 있을 수 있습니다.

`.github/workflows/super-linter.yml`을 생성하겠습니다.

![](/2022/Days/Images/Day75_CICD3.png)

이제 코드를 붙여 넣고 리포지토리에 코드를 커밋한 다음 Actions 탭으로 이동하면 아래와 같은 super-linter Wolkflow가 표시됩니다.

![](/2022/Days/Images/Day75_CICD4.png)

코드에서 리포지토리에 무언가를 push할 때 이 Wolkflow가 실행되도록 정의했기 때문에, super-linter.yml을 리포지토리에 push할 때 Wolkflow가 트리거되었습니다.

![](/2022/Days/Images/Day75_CICD5.png)

위에서 보시다시피 해킹 능력과 코딩 능력에 따라 몇 가지 오류가 발생했습니다.

적어도 아직은 제 코드는 아니었지만, 이 코드를 실행하고 오류가 발생했을 때 이 [문제](https://github.com/github/super-linter/issues/2255)를 발견했습니다.

2번 super-linter의 버전을 버전 3에서 4로 변경하고 작업을 다시 실행했습니다.

![](/2022/Days/Images/Day75_CICD6.png)

예상대로 해커 코딩에서 몇 가지 문제가 발생했으며 여기 [Wolkflow](https://github.com/MichaelCade/Jenkins-HelloWorld/runs/5600278515?check_suite_focus=true)에서 확인할 수 있습니다.

Wolkflow 내에서 무언가가 실패하거나 오류를 보고했을 때 리포지토리에 표시되는 모습을 지금 보여주고 싶었습니다.

![](/2022/Days/Images/Day75_CICD7.png)

이제 제 코드에서 문제를 해결하고 변경 사항을 적용하면 Wolkflow가 다시 실행됩니다(이미지에서 'bugs'를 해결하는 데 시간이 걸린 것을 볼 수 있습니다.) 파일을 삭제하는 것은 권장되지 않지만, 문제가 해결되었음을 보여주는 매우 빠른 방법입니다.

![](/2022/Days/Images/Day75_CICD8.png)

위에 강조 표시된 new Wolkflow 버튼을 누르면 수많은 작업의 문이 열립니다. 이번 챌린지를 진행하면서 눈치채셨겠지만, 저희는 거인의 어깨 위에 서서 코드, 자동화 및 기술을 널리 공유하여 삶을 더 편하게 만들고자 하는 것이 아닙니다.

![](/2022/Days/Images/Day75_CICD9.png)

Wolkflow가 성공했을 때 리포지토리의 녹색 체크 표시를 보여드리지 못했네요.

![](/2022/Days/Images/Day75_CICD10.png)

지금까지는 GitHub Actions의 기본적인 관점에 대해 설명했지만, 저와 같은 분이라면 GitHub Actions를 사용하여 많은 작업을 자동화할 수 있는 다른 방법을 알고 계실 것입니다.

다음에는 CD의 또 다른 영역인 애플리케이션을 환경에 배포하기 위해 ArgoCD를 살펴볼 것입니다.

## 자료

- [Jenkins is the way to build, test, deploy](https://youtu.be/_MXtbjwsz3A)
- [Jenkins.io](https://www.jenkins.io/)
- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)
- [ArgoCD Tutorial for Beginners](https://www.youtube.com/watch?v=MeU5_k9ssrs)
- [What is Jenkins?](https://www.youtube.com/watch?v=LFDrDnKPOTg)
- [Complete Jenkins Tutorial](https://www.youtube.com/watch?v=nCKxl7Q_20I&t=3s)
- [GitHub Actions](https://www.youtube.com/watch?v=R8_veQiYBjI)
- [GitHub Actions CI/CD](https://www.youtube.com/watch?v=mFFXuXjVgkU)

[Day 76](day76.md)에서 봐요!
