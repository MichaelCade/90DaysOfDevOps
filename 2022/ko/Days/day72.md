---
title: '#90DaysOfDevOps - Getting hands-on with Jenkins - Day 72'
published: false
description: 90DaysOfDevOps - Getting hands-on with Jenkins
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048829
---

## Jenkins 실습하기

오늘 계획은 Jenkins를 직접 사용해 보고 CI 파이프라인의 일부로 사용할 수 있는 몇 가지 예제 코드 베이스를 살펴보면서 무언가를 만들어 보는 것입니다.

### 파이프라인이란 무엇인가요?

시작하기 전에 CI와 관련하여 파이프라인이 무엇인지 알아야 하는데, 어제 세션에서 이미 다음 이미지와 함께 이 내용을 다루었습니다.

![](/2022/Days/Images/Day71_CICD4.png)

위의 프로세스 또는 단계를 자동화하여 최종적으로 배포된 애플리케이션을 고객, 최종 사용자 등에게 제공할 수 있는 결과물을 얻고자 합니다.

이 자동화된 프로세스를 통해 사용자와 고객에 대한 버전 관리를 수행할 수 있습니다. 모든 변경 사항, 기능 향상, 버그 수정 등은 이 자동화된 프로세스를 거쳐 코드가 양호한지 확인하기 위해 많은 수동 개입 없이 모든 것이 정상인지 확인합니다.

이 프로세스에는 안정적이고 반복 가능한 방식으로 소프트웨어를 빌드하고, 여러 단계의 테스트 및 배포를 통해 빌드된 소프트웨어('build'라고 함)를 발전시키는 작업이 포함됩니다.

Jenkins 파이프라인은 Jenkins파일이라는 텍스트 파일로 작성됩니다. 이 파일은 소스 제어 리포지토리에 커밋되어야 합니다. 이를 코드형 파이프라인이라고도 하며, 몇 주 전에 다룬 IaC에 매우 유사하게 비유할 수도 있습니다.

[Jenkins 파이프라인 정의](https://www.jenkins.io/doc/book/pipeline/#ji-toolbar)

### Jenkins 배포하기

Jenkins 배포를 재미있게 해봤는데요, [문서](https://www.jenkins.io/doc/book/installing/)를 보면 Jenkins를 설치할 수 있는 위치에 대한 옵션이 많다는 것을 알 수 있습니다.

저는 Minikube를 가지고 있고 이를 여러 번 사용했기 때문에 이 작업에도 Minikube를 사용하고 싶었습니다. (또한 무료입니다!) [Kubernetes 설치](https://www.jenkins.io/doc/book/installing/kubernetes/)에 나와 있는 단계는 벽에 부딪혀서 실행하지 못했지만, 여기에 단계를 문서화하면 두 가지를 비교할 수 있습니다.

첫 번째 단계는 Minikube 클러스터를 시작하고 실행하는 것으로, `minikube start` 명령으로 간단히 할 수 있습니다.

![](/2022/Days/Images/Day72_CICD1.png)

여기에서 찾을 수 있는 모든 YAML 구성과 값이 있는 폴더를 [여기](/2022/Days/CICD/Jenkins) 추가했습니다.이제 클러스터가 생겼으므로 다음을 실행하여 jenkins 네임스페이스를 생성할 수 있습니다. `kubectl create -f jenkins-namespace.yml`

![](/2022/Days/Images/Day72_CICD2.png)

클러스터에 Jenkins를 배포하기 위해 Helm을 사용할 것이며, Helm은 Kubernetes 섹션에서 다루었습니다. 먼저 jenkinsci Helm 리포지토리 `helm repo add jenkinsci https://charts.jenkins.io`를 추가한 다음 `helm repo update` chart를 업데이트해야 합니다.

![](/2022/Days/Images/Day72_CICD3.png)

Jenkins의 기본 개념은 파이프라인의 상태를 저장하는 것으로, 위의 Helm 설치를 지속성 없이 실행할 수 있지만 해당 pod가 재부팅, 변경 또는 수정되면 생성한 파이프라인이나 구성이 손실됩니다. `kubectl apply -f jenkins-volume.yml` 명령으로 jenkins-volume.yml 파일을 사용하여 지속성을 위한 volume을 생성합니다.

![](/2022/Days/Images/Day72_CICD4.png)

또한 이 YAML 파일과 명령어를 사용하여 생성할 수 있는 서비스 계정이 필요합니다. `kubectl apply -f jenkins-sa.yml`

![](/2022/Days/Images/Day72_CICD5.png)
이 단계에서는 Helm chart를 사용하여 배포하는 것이 좋으며, 먼저 `chart=jenkinsci/jenkins`를 사용하여 chart를 정의한 다음, 이 명령을 사용하여 배포할 것이며, 여기에는 이전에 클러스터에 배포한 지속성 및 서비스 계정이 포함된 jenkins-values.yml이 포함됩니다. `helm install jenkins -n jenkins -f jenkins-values.yml $chart`

![](/2022/Days/Images/Day72_CICD6.png)

이 단계에서는 pod가 이미지를 가져올 것이지만 pod가 스토리지에 액세스할 수 없으므로 Jenkins를 시작하고 실행하기 위한 구성을 시작할 수 없습니다.

이 단계에서는 문서가 무슨 일이 일어나야 하는지 이해하는 데 큰 도움이 되지 않았습니다. 하지만 Jenkins 설치를 시작할 수 있는 권한이 없다는 것을 알 수 있습니다.

![](/2022/Days/Images/Day72_CICD7.png)

위의 문제를 수정하거나 해결하려면, 우리가 제안한 이 위치에 Jenkins pod가 쓸 수 있도록 액세스 권한 또는 올바른 권한을 제공해야 합니다. 이를 위해 `minikube ssh`를 사용하여 실행 중인 Minikube Docker 컨테이너에 들어간 다음 `sudo chown -R 1000:1000 /data/jenkins-volume`을 사용하여 데이터 volume에 대한 권한이 설정되어 있는지 확인할 수 있습니다.

![](/2022/Days/Images/Day72_CICD8.png)

위의 프로세스로 pod가 수정되어야 하지만, 그렇지 않은 경우 `kubectl delete pod jenkins-0 -n jenkins` 명령으로 pod를 강제로 새로 고칠 수 있습니다. 이 시점에서, 2/2의 실행 중인 pod인 jenkins-0이 있어야 합니다.

![](/2022/Days/Images/Day72_CICD9.png)

이제 관리자 비밀번호가 필요하며 다음 명령어를 사용하여 비밀번호를 입력할 수 있습니다. `kubectl exec --namespace jenkins -it svc/jenkins -c jenkins -- /bin/cat /run/secrets/chart-admin-password && echo`

![](/2022/Days/Images/Day72_CICD10.png)

이제 워크스테이션에서 접근할 수 있도록 `port-forward` 명령을 사용할 것이므로 새 터미널을 엽니다. `kubectl --namespace jenkins port-forward svc/jenkins 8080:8080`

![](/2022/Days/Images/Day72_CICD11.png)

이제 브라우저를 열고 `http://localhost:8080`에 로그인하여 이전 단계에서 수집한 username: admin과 password로 인증할 수 있어야 합니다.

![](/2022/Days/Images/Day72_CICD12.png)

인증이 완료되면 Jenkins 시작 페이지는 다음과 같이 표시되어야 합니다:

![](/2022/Days/Images/Day72_CICD13.png)

여기에서 "Manage Jenkins"로 이동하면 사용 가능한 몇 가지 업데이트가 있는 "Manage Plugins"가 표시됩니다. 해당 플러그인을 모두 선택하고 "Download now and install after restart"를 선택합니다.

![](/2022/Days/Images/Day72_CICD14.png)

더 나아가 shell 스크립트를 사용하여 Jenkins 배포를 자동화하고 싶다면 트위터에서 저와 공유한 이 훌륭한 리포지토리를 참고하세요 [mehyedes/nodejs-k8s](https://github.com/mehyedes/nodejs-k8s/blob/main/docs/automated-setup.md).

### Jenkinsfile

이제 Kubernetes 클러스터에 Jenkins가 배포되었으므로 이제 돌아가서 이 Jenkinsfile에 대해 생각해 볼 수 있습니다.

모든 Jenkinsfile은 이렇게 시작될 가능성이 높습니다. 먼저 파이프라인의 단계를 정의하는 곳이며, 이 경우에는 빌드 > 테스트 > 배포가 있습니다. 하지만 여기서는 특정 단계를 호출하기 위해 `echo` 명령을 사용하는 것 외에는 아무것도 하지 않습니다.

```
Jenkinsfile (Declarative Pipeline)

pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}

```

Jenkins 대시보드에서 "New Item"을 선택하고 항목에 이름을 지정합니다. 저는 "echo1"이라고 하고 이것이 파이프라인이라고 해보겠습니다.

![](/2022/Days/Images/Day72_CICD15.png)

확인을 누르면 파이프라인에만 관심이 있는 간단한 테스트를 위한 탭(일반, 빌드 트리거, 고급 프로젝트 옵션 및 파이프라인)이 표시됩니다. 파이프라인에서 스크립트를 추가할 수 있으며, 위의 스크립트를 복사하여 상자에 붙여 넣을 수 있습니다.

위에서 말했듯이 이것은 많은 작업을 수행하지는 않지만 빌드 > 테스트 > 배포의 단계를 보여줍니다.

![](/2022/Days/Images/Day72_CICD16.png)

저장을 클릭하면 이제 아래에 강조 표시된 "build now"를 사용하여 빌드를 실행할 수 있습니다.

![](/2022/Days/Images/Day72_CICD17.png)

또한 터미널을 열고 `kubectl get pods -n jenkins`를 실행하여 어떤 일이 발생하는지 확인해야 합니다.

![](/2022/Days/Images/Day72_CICD18.png)

자, 아주 간단하지만 이제 Jenkins 배포와 설치가 올바르게 작동하고 있으며 여기에서 CI 파이프라인의 구성 요소를 볼 수 있습니다.

다음 섹션에서는 Jenkins 파이프라인을 구축하겠습니다.

## 자료

- [Jenkins is the way to build, test, deploy](https://youtu.be/_MXtbjwsz3A)
- [Jenkins.io](https://www.jenkins.io/)
- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)
- [ArgoCD Tutorial for Beginners](https://www.youtube.com/watch?v=MeU5_k9ssrs)
- [What is Jenkins?](https://www.youtube.com/watch?v=LFDrDnKPOTg)
- [Complete Jenkins Tutorial](https://www.youtube.com/watch?v=nCKxl7Q_20I&t=3s)
- [GitHub Actions](https://www.youtube.com/watch?v=R8_veQiYBjI)
- [GitHub Actions CI/CD](https://www.youtube.com/watch?v=mFFXuXjVgkU)

[Day 73](day73.md)에서 봐요!
