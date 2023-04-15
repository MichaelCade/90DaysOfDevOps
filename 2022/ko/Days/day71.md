---
title: '#90DaysOfDevOps - What is Jenkins? - Day 71'
published: false
description: 90DaysOfDevOps - What is Jenkins?
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048745
---

## Jenkins란 무엇인가요?

Jenkins는 새로 만든 코드를 지속적으로 개발, 테스트 및 배포할 수 있는 지속적 통합 도구입니다.

야간 빌드 또는 지속적 개발을 통해 이를 달성할 수 있는 방법에는 두 가지가 있습니다. 첫 번째 옵션은 개발자가 하루 종일 작업을 개발하다가 정해진 하루가 끝나면 변경 사항을 소스 코드 저장소에 push하는 것입니다. 그런 다음 밤에는 단위 테스트를 실행하고 소프트웨어를 빌드합니다. 이것이 모든 코드를 통합하는 오래된 방식이라고 할 수 있습니다.

![](/2022/Days/Images/Day71_CICD1.png)

다른 옵션이자 선호되는 방식은 개발자가 소스 코드에 변경 사항을 커밋하고 해당 코드 커밋이 완료되면 빌드 프로세스가 지속적으로 시작되는 방식입니다.

![](/2022/Days/Images/Day71_CICD2.png)

위의 방법을 사용하면 전 세계에 분산된 개발자가 있는 경우 매일 정해진 시간에 코드 변경 사항 커밋을 중단할 수 없습니다. 이러한 테스트와 빌드 프로세스를 제어하기 위해 CI 서버 역할을 하는 것이 바로 Jenkins입니다.

![](/2022/Days/Images/Day71_CICD3.png)

여기서는 Jenkins에 대해 이야기하고 있지만 나중에 몇 가지를 더 추가하여 Jenkins가 전반적으로 가장 인기 있는 것으로 보이는 이유와 그 이유, 그리고 다른 도구가 Jenkins보다 무엇을 할 수 있는지에 대해 알아보고자 합니다.

- TravisCI - GitHub에서 호스팅되는 소프트웨어 프로젝트를 빌드하고 테스트하는 데 사용되는 호스팅된 분산형 지속적 통합 서비스입니다.
- Bamboo - 빠른 컴파일을 위해 여러 빌드를 병렬로 실행할 수 있으며, 리포지토리와 연결할 수 있는 기능이 내장되어 있고 Ant 및 Maven용 빌드 작업이 있습니다.
- Buildbot - 소프트웨어 빌드, 테스트 및 릴리스 프로세스를 자동화하기 위한 오픈 소스 프레임워크입니다. Python으로 작성되었으며 여러 플랫폼에서 작업의 분산 병렬 실행을 지원합니다.
- Apache Gump - 매일 밤 Java 프로젝트를 빌드하고 테스트하도록 설계된 Java 프로젝트 전용으로, 모든 프로젝트가 API 및 기능 수준 모두에서 호환되도록 보장합니다.

이제부터는 Jenkins에 초점을 맞추겠습니다. Jenkins는 위의 모든 도구와 마찬가지로 오픈 소스이며 Java로 작성된 자동화 서버입니다. 지속적인 통합을 통해 소프트웨어 개발 프로세스를 자동화하고 지속적인 배포를 용이하게 하는 데 사용됩니다.

### Jenkins의 특징

예상할 수 있듯이 Jenkins는 다양한 영역에 걸쳐 많은 기능을 가지고 있습니다.

**간편한 설치** - Jenkins는 Windows, macOS 및 Linux 운영 체제용 패키지와 함께 실행할 준비가 된 독립형 자바 기반 프로그램입니다.

**간편한 구성** - 오류 확인 및 기본 제공 도움말이 포함된 웹 인터페이스를 통해 쉽게 설정 및 구성할 수 있습니다.

**플러그인** - 업데이트 센터에서 다양한 플러그인을 사용할 수 있으며 CI/CD 툴체인의 여러 도구와 통합됩니다.

**확장 가능** - 사용 가능한 플러그인 외에도 플러그인 아키텍처를 통해 Jenkins를 확장할 수 있어 거의 무한한 용도로 사용할 수 있는 옵션을 제공합니다.

**분산** - Jenkins는 여러 머신에 작업을 쉽게 분산하여 여러 플랫폼에서 빌드, 테스트 및 배포 속도를 높일 수 있도록 지원합니다.

### Jenkins 파이프라인

이 파이프라인을 보셨겠지만, 훨씬 광범위하게 사용되며 특정 도구에 대해서는 언급하지 않았습니다.

Jenkins에 코드를 커밋하면 모든 자동화된 테스트를 통해 애플리케이션을 빌드한 다음 각 단계가 완료되면 해당 코드를 릴리스 및 배포합니다. 이 프로세스를 자동화할 수 있는 것이 바로 Jenkins입니다.

![](/2022/Days/Images/Day71_CICD4.png)

### Jenkins 아키텍처

먼저, 바퀴를 재발명하고 싶지 않기 때문에 항상 [Jenkins 문서](https://www.jenkins.io/doc/developer/architecture/)에서 시작하는 것이 가장 좋지만, 여기에도 제가 메모하고 배운 내용을 적어 두려고 합니다.

Jenkins는 윈도우, 리눅스, 맥OS 등 다양한 운영체제에 설치할 수 있을 뿐만 아니라 Docker 컨테이너로 배포하거나 Kubernetes 내에서 배포할 수 있는 기능도 있습니다. [Jenkins 설치하기](https://www.jenkins.io/doc/book/installing/)

이번 글에서는 Minikube 클러스터 내에 Jenkins를 설치하여 Kubernetes에 배포하는 과정을 시뮬레이션해 보겠습니다. 하지만 이는 이 섹션의 나머지 부분에서 구성한 시나리오에 따라 달라집니다.

이제 아래 이미지를 분석해 보겠습니다.

1단계 - 개발자가 소스 코드 리포지토리에 변경 사항을 커밋합니다.

2단계 - Jenkins가 정기적으로 리포지토리를 확인하여 새 코드를 가져옵니다.

3단계 - 빌드 서버가 코드를 실행 파일로 빌드하며, 이 예에서는 잘 알려진 빌드 서버로 maven을 사용하고 있습니다. 다루어야 할 또 다른 영역입니다.

4단계 - 빌드에 실패하면 개발자에게 피드백이 다시 전송됩니다.

5단계 - 그런 다음 Jenkins는 빌드된 앱을 테스트 서버에 배포합니다. 이 예에서는 잘 알려진 테스트 서버로 selenium을 사용하고 있습니다. 또 다른 영역입니다.

6단계 - 테스트가 실패하면 개발자에게 피드백이 전달됩니다.

7단계 - 테스트가 성공하면 프로덕션에 릴리스할 수 있습니다.

이 주기는 연속적이기 때문에 애플리케이션을 몇 시간, 며칠, 몇 달, 몇 년이 아닌 몇 분 만에 업데이트할 수 있습니다!

![](/2022/Days/Images/Day71_CICD5.png)

필요한 경우 master가 slave Jenkins 환경에 작업을 배포할 수 있는 master-slave 기능이 있는 등 Jenkins의 아키텍처에는 훨씬 더 많은 기능이 있습니다.

참고로 Jenkins는 오픈 소스이므로 지원을 필요로 하는 많은 기업이 있을 것이며, CloudBees는 유료 엔터프라이즈 고객을 위한 지원 및 기타 기능을 제공하는 엔터프라이즈 버전의 Jenkins입니다.

고객 중 한 예로 Bosch를 들 수 있는데, Bosch 사례 연구는 [여기](https://assets.ctfassets.net/vtn4rfaw6n2j/case-study-boschpdf/40a0b23c61992ed3ee414ae0a55b6777/case-study-bosch.pdf)에서 확인할 수 있습니다.

내일은 애플리케이션의 단계별 예시를 통해 Jenkins를 사용하는 방법을 살펴본 다음 다른 도구와 함께 사용할 수 있도록 하겠습니다.

## 자료

- [Jenkins is the way to build, test, deploy](https://youtu.be/_MXtbjwsz3A)
- [Jenkins.io](https://www.jenkins.io/)
- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)
- [ArgoCD Tutorial for Beginners](https://www.youtube.com/watch?v=MeU5_k9ssrs)
- [What is Jenkins?](https://www.youtube.com/watch?v=LFDrDnKPOTg)
- [Complete Jenkins Tutorial](https://www.youtube.com/watch?v=nCKxl7Q_20I&t=3s)
- [GitHub Actions](https://www.youtube.com/watch?v=R8_veQiYBjI)
- [GitHub Actions CI/CD](https://www.youtube.com/watch?v=mFFXuXjVgkU)

[Day 72](day72.md)에서 봐요!킨
