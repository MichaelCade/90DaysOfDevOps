---
title: '#90DaysOfDevOps - ArgoCD Overview - Day 76'
published: false
description: 90DaysOfDevOps - ArgoCD Overview
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048809
---

## ArgoCD 개요

"Argo CD는 Kubernetes를 위한 선언적 GitOps 지속적 배포 도구입니다."

버전 제어가 핵심입니다. 환경을 즉석에서 변경했는데 그 변경 사항을 기억하지 못하고 불이 켜져 있고 모든 것이 초록색이기 때문에 계속 진행했던 적이 있으신가요? 변경을 시도했다가 모든 것 또는 일부가 망가진 적이 있나요? 변경한 사실을 알았다면 잘못된 스크립트나 맞춤법 오류를 빠르게 되돌릴 수 있을 것입니다. 하지만 대규모로 변경을 진행하다 보니 본인도 몰랐을 수도 있고, 바로 발견하지 못해 비즈니스가 어려움을 겪고 있을 수도 있습니다. 따라서 버전 관리가 중요합니다. 뿐만 아니라 "애플리케이션 정의, 구성 및 환경은 선언적이어야 하며 버전 관리가 이루어져야 합니다." 이 외에도 "애플리케이션 배포 및 수명 주기 관리는 자동화되고, 감시 가능하며, 이해하기 쉬워야 한다"고 언급합니다.

운영 배경을 가지고 있지만 IaC로 사용하는 데 많은 경험을 쌓은 저로서는 지속적인 배포/제공 워크플로우를 통해 이 모든 좋은 것들을 처리할 수 있는 다음 단계라고 생각합니다.

[ArgoCD란](https://argo-cd.readthedocs.io/en/stable/)

### ArgoCD 배포

이번 배포에서는 신뢰할 수 있는 Minikube Kubernetes 클러스터를 다시 로컬로 사용할 것입니다.

```Shell
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

![](/2022/Days/Images/Day76_CICD1.png)

`kubectl get pods -n argocd`로 모든 ArgoCD pod가 실행되고 있는지 확인합니다.

![](/2022/Days/Images/Day76_CICD2.png)

또한, `kubectl get all -n argocd`로 네임스페이스에 배포한 모든 것을 확인해봅시다.

![](/2022/Days/Images/Day76_CICD3.png)

위의 내용이 정상적으로 보이면 포트 포워드를 통해 접근하는 것을 고려해야 합니다. 새 터미널에서 `kubectl port-forward svc/argocd-server -n argocd 8080:443` 명령을 사용합니다.

그런 다음 새 웹 브라우저를 열고 `https://localhost:8080`로 이동합니다.

![](/2022/Days/Images/Day76_CICD4.png)

로그인하려면 관리자 사용자 이름이 필요하며, 생성한 암호를 비밀번호로 사용하려면 `kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d && echo`를 사용합니다.

![](/2022/Days/Images/Day76_CICD5.png)

로그인하면 비어있는 CD 캔버스가 표시됩니다.

![](/2022/Days/Images/Day76_CICD6.png)

### 애플리케이션 배포

이제 ArgoCD를 실행하고 실행 중이므로 이제 Helm뿐만 아니라 Git 리포지토리에서 애플리케이션을 배포하는 데 사용할 수 있습니다.

제가 배포하고자 하는 애플리케이션은 Pac-Man입니다. 네, Pac-Man은 유명한 게임이자 데이터 관리와 관련하여 많은 데모에서 사용하는 게임이며, 우리가 Pac-Man을 보는 것은 이번이 마지막이 아닐 것입니다.

[Pac-Man](https://github.com/MichaelCade/pacman-tanzu.git)의 리포지토리는 여기에서 찾을 수 있습니다.

스크린샷을 사용하여 각 단계를 설명하는 대신, 이 특정 애플리케이션 배포를 위해 수행한 단계를 다루는 연습 동영상을 만드는 것이 더 쉬울 것이라고 생각했습니다.

[ArgoCD 데모 - 90일간의 개발 운영](https://www.youtube.com/watch?v=w6J413_j0hA)

참고 - 영상 중 앱 상태가 정상인데도 만족스럽지 않은 서비스가 있는데, 이는 Pac-Man 서비스에 대해 설정된 로드밸런서 유형이 보류 중이고, Minikube에는 로드밸런서가 구성되지 않았기 때문입니다. 이 문제를 테스트하고 싶으시다면 서비스에 대한 YAML을 ClusterIP로 변경하고 포트 포워딩을 사용하여 게임을 플레이할 수 있습니다.

이것으로 CICD 파이프라인 섹션을 마칩니다. 현재 업계에서 이 영역에 많은 관심이 집중되고 있으며, 일반적으로 CICD 내에서 사용되는 방법론과 관련된 GitOps 관련 용어도 듣게 될 것입니다.

다음 섹션에서는 새로운 개념은 아니지만 환경을 다르게 바라보면서 점점 더 중요해지고 있는 또 다른 개념 또는 영역인 Observability에 대해 살펴봅니다.

## 자료

- [Jenkins is the way to build, test, deploy](https://youtu.be/_MXtbjwsz3A)
- [Jenkins.io](https://www.jenkins.io/)
- [ArgoCD](https://argo-cd.readthedocs.io/en/stable/)
- [ArgoCD Tutorial for Beginners](https://www.youtube.com/watch?v=MeU5_k9ssrs)
- [What is Jenkins?](https://www.youtube.com/watch?v=LFDrDnKPOTg)
- [Complete Jenkins Tutorial](https://www.youtube.com/watch?v=nCKxl7Q_20I&t=3s)
- [GitHub Actions](https://www.youtube.com/watch?v=R8_veQiYBjI)
- [GitHub Actions CI/CD](https://www.youtube.com/watch?v=mFFXuXjVgkU)

[Day 77](day77.md)에서 봐요!
