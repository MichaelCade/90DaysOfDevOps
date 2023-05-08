---
title: '#90DaysOfDevOps - Building a Jenkins Pipeline - Day 73'
published: false
description: 90DaysOfDevOps - Building a Jenkins Pipeline
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048766
---

## Jenkins 파이프라인 구축하기

지난 섹션에서는 Minikube 클러스터에 Jenkins를 배포하고 파이프라인의 단계를 echo하는 것 외에는 별다른 기능을 하지 않는 매우 기본적인 Jenkins 파이프라인을 설정했습니다.

또한 Jenkins Pipeline 생성에서 실행할 수 있는 몇 가지 예제 스크립트가 있다는 것을 보셨을 것입니다.

![](/2022/Days/Images/Day73_CICD1.png)

첫 번째 데모 스크립트는 "Declarative (Kubernetes)"이며 아래 단계를 볼 수 있습니다.

```Yaml
// Uses Declarative syntax to run commands inside a container.
pipeline {
    agent {
        kubernetes {
            // Rather than inline YAML, in a multibranch Pipeline you could use: yamlFile 'jenkins-pod.yaml'
            // Or, to avoid YAML:
            // containerTemplate {
            //     name 'shell'
            //     image 'ubuntu'
            //     command 'sleep'
            //     args 'infinity'
            // }
            yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: shell
    image: ubuntu
    command:
    - sleep
    args:
    - infinity
'''
            // Can also wrap individual steps:
            // container('shell') {
            //     sh 'hostname'
            // }
            defaultContainer 'shell'
        }
    }
    stages {
        stage('Main') {
            steps {
                sh 'hostname'
            }
        }
    }
}
```

이 파이프라인을 실행하면 어떤 결과가 발생하는지 아래에서 확인할 수 있습니다.

![](/2022/Days/Images/Day73_CICD2.png)

### Job 만들기

#### 목표

- 간단한 앱을 만들어 GitHub 공용 리포지토리에 저장 [https://github.com/scriptcamp/kubernetes-kaniko.git](https://github.com/scriptcamp/kubernetes-kaniko.git)

- Jenkins를 사용하여 Docker 컨테이너 이미지를 빌드하고 DockerHub에 push합니다. (이를 위해 비공개 리포지토리를 사용합니다.)

Minikube에서 실행 중이거나 Minikube를 사용하는 Kubernetes 클러스터에서 이 작업을 수행하려면 [Kaniko](https://github.com/GoogleContainerTools/kaniko#running-kaniko-in-a-kubernetes-cluster)라는 것을 사용해야 하지만, 실제 Kubernetes 클러스터에서 Jenkins를 사용하거나 서버에서 실행하는 경우 에이전트를 지정하여 Docker 빌드 명령을 수행하고 이를 DockerHub에 업로드할 수 있는 기능을 제공하는 것이 일반적입니다.

위의 내용을 염두에 두고 GitHub credentials로 Kubernetes에 시크릿을 배포해 보겠습니다.

```Shell
kubectl create secret docker-registry dockercred \
    --docker-server=https://index.docker.io/v1/ \
    --docker-username=<dockerhub-username> \
    --docker-password=<dockerhub-password>\
    --docker-email=<dockerhub-email>
```

여기서 다룰 내용 대부분을 관통하는 [DevOpsCube.com](https://devopscube.com/build-docker-image-kubernetes-pod/)의 또 다른 훌륭한 리소스를 공유하고자 합니다.

### Jenkins에 credentials 추가하기

하지만 저희와 다른 Jenkins 시스템을 사용 중이라면 Jenkins 내에서 credentials를 정의한 다음 파이프라인 및 구성 내에서 여러 번 사용하고 싶을 것입니다. 생성 시 결정한 ID를 사용하여 파이프라인에서 이러한 credentials를 참조할 수 있습니다. 계속해서 단계를 진행하여 DockerHub 및 GitHub에 대한 사용자 항목을 만들었습니다.

먼저 "Manage Jenkins"를 선택한 다음 "Manage Credentials"를 선택합니다.

![](/2022/Days/Images/Day73_CICD3.png)

페이지 중앙에 Jenkins로 범위가 설정된 Stores가 표시되면 여기에서 Jenkins를 클릭합니다.

![](/2022/Days/Images/Day73_CICD4.png)

이제 Global Credentials (Unrestricted)를 선택합니다.

![](/2022/Days/Images/Day73_CICD5.png)

그런 다음 왼쪽 상단에 Add Credentials가 있습니다.

![](/2022/Days/Images/Day73_CICD6.png)

계정에 대한 세부 정보를 입력한 다음 확인을 선택하고, 이 credentials를 호출할 때 참조할 ID를 기억하세요. 여기서도 비밀번호 대신 특정 토큰 액세스를 사용하는 것이 좋습니다.

![](/2022/Days/Images/Day73_CICD7.png)

GitHub의 경우 [Personal Access Token](https://vzilla.co.uk/vzilla-blog/creating-updating-your-github-personal-access-token)을 사용해야 합니다.

저는 이 계정을 생성하는 과정이 직관적이지 않아서 사용하지는 않지만, UI에서 명확하지 않아서 그 과정을 공유하고 싶었습니다.

### 파이프라인 구축

Kubernetes 클러스터에 비밀로 배포된 DockerHub credentials를 가지고 있으며, 이 credentials를 파이프라인의 DockerHub 단계에 배포하기 위해 호출할 것입니다.

파이프라인 스크립트는 아래에서 볼 수 있는 것과 같으며, 파이프라인의 프로젝트 가져오기 단계에서도 볼 수 있는 GitHub 리포지토리에 있는 Jenkinsfile이 될 수 있습니다.

```Yaml
podTemplate(yaml: '''
    apiVersion: v1
    kind: Pod
    spec:
      containers:
      - name: maven
        image: maven
        command:
        - sleep
        args:
        - 99d
      - name: kaniko
        image: gcr.io/kaniko-project/executor:debug
        command:
        - sleep
        args:
        - 9999999
        volumeMounts:
        - name: kaniko-secret
          mountPath: /kaniko/.docker
      restartPolicy: Never
      volumes:
      - name: kaniko-secret
        secret:
            secretName: dockercred
            items:
            - key: .dockerconfigjson
              path: config.json
''') {
  node(POD_LABEL) {
    stage('Get the project') {
      git url: 'https://github.com/scriptcamp/kubernetes-kaniko.git', branch: 'main'
      container('maven') {
        stage('Test the project') {
          sh '''
          echo pwd
          '''
        }
      }
    }

    stage('Build & Test the Docker Image') {
      container('kaniko') {
        stage('Deploy to DockerHub') {
          sh '''
            /kaniko/executor --context `pwd` --destination me1e/helloword
          '''
        }
      }
    }

  }
}
```

Jenkins 대시보드에서 항목을 시작하려면 "New Item"을 선택해야 합니다.

![](/2022/Days/Images/Day73_CICD8.png)

그런 다음 항목에 이름을 지정하고 파이프라인을 선택한 다음 확인을 누릅니다.

![](/2022/Days/Images/Day73_CICD9.png)

일반 트리거나 빌드 트리거는 선택하지 않겠지만 흥미로운 일정과 기타 구성이 있을 수 있으므로 이 트리거를 사용해 보겠습니다.

![](/2022/Days/Images/Day73_CICD10.png)

마지막에 있는 Pipeline 탭에만 관심이 있습니다.

![](/2022/Days/Images/Day73_CICD11.png)

파이프라인 정의에서 위에 있는 파이프라인 스크립트를 스크립트 섹션에 복사하여 붙여 넣고 저장을 누릅니다.

![](/2022/Days/Images/Day73_CICD12.png)

다음으로 페이지 왼쪽에서 "Build Now" 옵션을 선택합니다.

![](/2022/Days/Images/Day73_CICD13.png)

이제 1분 미만의 짧은 시간을 기다려야 합니다. Status 아래에 위에서 스크립트에서 정의한 단계가 표시됩니다.

![](/2022/Days/Images/Day73_CICD14.png)

더 중요한 것은 이제 DockerHub로 이동하여 새 빌드가 있는지 확인하는 것입니다.

![](/2022/Days/Images/Day73_CICD15.png)

전체적으로 파악하는 데 시간이 좀 걸렸지만, Minikube와 GitHub 및 DockerHub에 대한 액세스를 사용하여 누구나 실행할 수 있는 시나리오를 실습하고 작업하기 위해 계속 진행하려고 했습니다.

이 데모에 사용한 DockerHub 리포지토리는 비공개 리포지토리였습니다. 하지만 다음 섹션에서는 이러한 단계 중 일부를 발전시켜 단순히 `pwd`를 출력하고 몇 가지 테스트 및 빌드 단계를 실행하는 대신 무언가를 수행하도록 하고 싶습니다.

## 자료

- [Jenkins is the way to build, test, deploy](https://youtu.be/_MXtbjwsz3A)
- [Jenkins.io](https://www.jenkins.io/)
- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)
- [ArgoCD Tutorial for Beginners](https://www.youtube.com/watch?v=MeU5_k9ssrs)
- [What is Jenkins?](https://www.youtube.com/watch?v=LFDrDnKPOTg)
- [Complete Jenkins Tutorial](https://www.youtube.com/watch?v=nCKxl7Q_20I&t=3s)
- [GitHub Actions](https://www.youtube.com/watch?v=R8_veQiYBjI)
- [GitHub Actions CI/CD](https://www.youtube.com/watch?v=mFFXuXjVgkU)

[Day 74](day74.md)에서 봐요!
