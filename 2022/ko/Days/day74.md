---
title: '#90DaysOfDevOps - Hello World - Jenkinsfile App Pipeline - Day 74'
published: false
description: 90DaysOfDevOps - Hello World - Jenkinsfile App Pipeline
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048744
---

## Hello World - Jenkinsfile 앱 파이프라인

지난 섹션에서는 공개 GitHub 리포지토리에 있는 dockerfile에서 비공개 Dockerhub 리포지토리로 docker 이미지를 push하는 간단한 Jenkins 파이프라인을 구축했습니다.

이 섹션에서는 한 단계 더 나아가 간단한 애플리케이션을 통해 다음을 달성하고자 합니다.

### 목표

- Dockerfile (Hello World)
- Jenkinsfile
- GitHub 리포지토리가 업데이트될 때 트리거할 Jenkins 파이프라인
- GitHub 리포지토리를 소스로 사용
- Run - 리포지토리 복제/가져오기, 빌드, 테스트, 배포 단계
- 점진적인 버전 번호로 DockerHub에 배포
- Kubernetes 클러스터에 배포하는(여기에는 GitHub credentials를 사용하는 다른 job 및 매니페스트 리포지토리가 포함됨) Stretch Goal

### 1단계

[GitHub 리포지토리](https://github.com/MichaelCade/Jenkins-HelloWorld)가 있습니다. 여기에는 현재 Dockerfile과 index.html이 있습니다.

![](/2022/Days/Images/Day74_CICD1.png)

위에서 파이프라인에서 소스로 사용하던 것을 이제 해당 Jenkins 파이프라인 스크립트를 GitHub 리포지토리에도 추가하려고 합니다.

![](/2022/Days/Images/Day74_CICD2.png)

이제 Jenkins 대시보드로 돌아가서 새 파이프라인을 만들되, 스크립트를 붙여 넣는 대신 "Pipeline script from SCM"를 사용하고 아래의 구성 옵션을 사용하겠습니다.

참고로 리포지토리 URL은 `https://github.com/MichaelCade/Jenkins-HelloWorld.git`을 사용하겠습니다.

![](/2022/Days/Images/Day74_CICD3.png)

이 시점에서 저장 및 적용을 누르면 수동으로 파이프라인을 실행하여 DockerHub 리포지토리에 업로드된 새 Docker 이미지를 빌드할 수 있습니다.

하지만 리포지토리 또는 소스 코드가 변경될 때마다 빌드를 트리거하도록 일정을 설정하고 싶습니다. webhooks를 사용하거나 예약 pull을 사용할 수 있습니다.

파이프라인을 유지하기 위해 값비싼 클라우드 리소스를 사용하고 있고 코드 저장소에 변경 사항이 많다면 많은 비용이 발생할 수 있으므로 이 점을 크게 고려해야 합니다. 이것이 데모 환경이라는 것을 알고 있기 때문에 "poll scm" 옵션을 사용하고 있습니다. (또한 Minikube를 사용하면 webhooks를 사용할 수 있는 기능이 부족하다고 생각합니다.)

![](/2022/Days/Images/Day74_CICD4.png)

어제 세션 이후 변경한 한 가지는 이제 이미지를 공개 저장소(이 경우 michaelcade1\90DaysOfDevOps)에 업로드하고 싶다는 것인데, 제 Jenkinsfile에는 이미 이 변경 사항이 있습니다. 그리고 이전 섹션에서 기존 데모 컨테이너 이미지를 모두 제거했습니다.

![](/2022/Days/Images/Day74_CICD5.png)

여기서 거꾸로 돌아가서 파이프라인을 생성한 다음 이전에 표시된 대로 구성을 추가했습니다.

![](/2022/Days/Images/Day74_CICD6.png)

이 단계에서는 파이프라인이 실행되지 않았으며 스테이지 뷰는 다음과 같이 표시됩니다.

![](/2022/Days/Images/Day74_CICD7.png)

이제 "Build Now" 버튼을 트리거해 보겠습니다. 그러면 스테이지 뷰에 스테이지가 표시됩니다.

![](/2022/Days/Images/Day74_CICD8.png)

이제 DockerHub 리포지토리로 이동하면 2개의 새 Docker 이미지가 있어야 합니다. 빌드 ID는 1이고 최신 빌드가 있어야 하는데, 이는 "Upload to DockerHub"를 기반으로 생성하는 모든 빌드에 대해 Jenkins Build_ID 환경 변수를 사용하여 버전을 전송하고 최신 빌드도 발행하기 때문입니다.

![](/2022/Days/Images/Day74_CICD9.png)

이제 아래와 같이 깃허브 리포지토리에 index.html 파일을 업데이트해 보겠습니다.

![](/2022/Days/Images/Day74_CICD10.png)

Jenkins로 돌아가서 "Build Now"를 다시 선택하면 됩니다. 2번 빌드가 성공했는지 확인해 보겠습니다.

![](/2022/Days/Images/Day74_CICD11.png)

그런 다음 DockerHub를 간단히 살펴보면 태그가 버전 2와 최신 태그가 있는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day74_CICD12.png)

여기서 한 가지 주목할 점은, 제가 Kubernetes 클러스터에 액세스 및 인증을 통해 DockerHub로 도커 빌드를 push할 수 있는 시크릿을 추가했다는 것입니다. 이 과정을 따라하는 경우 계정에 대해 이 과정을 반복하고 내 리포지토리 및 계정과 연결된 Jenkins파일도 변경해야 합니다.

## 자료

- [Jenkins is the way to build, test, deploy](https://youtu.be/_MXtbjwsz3A)
- [Jenkins.io](https://www.jenkins.io/)
- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)
- [ArgoCD Tutorial for Beginners](https://www.youtube.com/watch?v=MeU5_k9ssrs)
- [What is Jenkins?](https://www.youtube.com/watch?v=LFDrDnKPOTg)
- [Complete Jenkins Tutorial](https://www.youtube.com/watch?v=nCKxl7Q_20I&t=3s)
- [GitHub Actions](https://www.youtube.com/watch?v=R8_veQiYBjI)
- [GitHub Actions CI/CD](https://www.youtube.com/watch?v=mFFXuXjVgkU)

[Day 75](day75.md)에서 봐요!
