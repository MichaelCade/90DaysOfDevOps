---
title: '#90DaysOfDevOps - All other things Ansible - Automation Controller (Tower), AWX, Vault - Day 69'
published: false
description: '90DaysOfDevOps - All other things Ansible - Automation Controller (Tower), AWX, Vault'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048714
---

## 기타 모든 Ansible - 자동화 컨트롤러(타워), AWX, Vault

구성 관리에 대한 섹션을 마무리하면서 Ansible을 다룰 때 접할 수 있는 다른 영역에 대해 살펴보고 싶었습니다.

Ansible 자동화 플랫폼을 구성하는 많은 제품이 있습니다.

Red Hat Ansible Automation Platform은 조직 전반에서 자동화를 구축하고 운영하기 위한 기반입니다. 이 플랫폼에는 전사적인 자동화를 구현하는 데 필요한 모든 도구가 포함되어 있습니다.

![](/2022/Days/Images/Day69_config1.png)

이 글에서는 이 중 일부를 다루려고 합니다. 하지만 더 자세한 내용은 공식 Red Hat Ansible 사이트에서 더 많은 정보를 확인할 수 있습니다. [Ansible.com](https://www.ansible.com/?hsLang=en-us)

### Ansible 자동화 컨트롤러 | AWX

자동화 컨트롤러와 AWX는 제공하는 기능이 매우 유사하기 때문에 이 두 가지를 함께 묶었습니다.

AWX 프로젝트 또는 줄여서 AWX는 Red Hat이 후원하는 오픈 소스 커뮤니티 프로젝트로, 사용자 환경 내에서 Ansible 프로젝트를 더 잘 제어할 수 있도록 지원합니다. AWX는 자동화 컨트롤러 구성 요소가 파생된 업스트림 프로젝트입니다.

엔터프라이즈 솔루션을 찾고 계신다면 자동화 컨트롤러를 찾으시거나 이전에 Ansible Tower라고 들어보셨을 것입니다. Ansible 자동화 컨트롤러는 Ansible 자동화 플랫폼의 컨트롤 플레인입니다.

AWX와 자동화 컨트롤러는 지금까지 이 섹션에서 다룬 다른 모든 기능 위에 다음과 같은 기능을 제공합니다.

- 사용자 인터페이스
- 역할 기반 액세스 제어
- workflow
- CI/CD 통합

자동화 컨트롤러는 지원 비용을 지불하는 엔터프라이즈 제품입니다.

Minikube kubernetes 환경 내에서 AWX를 배포하는 방법을 살펴보겠습니다.

### Ansible AWX 배포하기

AWX를 kubernetes 클러스터에 배포할 필요는 없으며, [github](https://github.com/ansible/awx)에서 AWX에 대한 자세한 내용을 확인할 수 있습니다. 그러나 버전 18.0부터는 AWX 오퍼레이터를 사용하여 AWX를 설치하는 것이 선호됩니다.

우선 Minikube 클러스터가 필요합니다. kubernetes 섹션에서 `minikube start --cpus=4 --memory=6g --addons=ingress` 명령으로 새 Minikube 클러스터를 생성하는 방법을 따라 했다면 이 작업을 수행할 수 있습니다.

![](/2022/Days/Images/Day69_config2.png)

공식 [Ansible AWX 오퍼레이터](https://github.com/ansible/awx-operator)는 여기에서 확인할 수 있습니다. 설치 지침에 명시된 대로 이 리포지토리를 복제한 다음 배포를 실행해야 합니다.

위의 리포지토리를 folk한 다음 `git clone https://github.com/MichaelCade/awx-operator.git`을 실행했는데, 여러분도 똑같이 하시고 제가 만든 리포지토리는 변경되거나 존재하지 않을 수 있으므로 사용하지 않는 것이 좋습니다.

복제된 리포지토리에서 아래와 같이 `NodePort`를 `ClusterIP`로 변경해야 하는 awx-demo.yml 파일을 찾을 수 있습니다:

```Yaml
---
apiVersion: awx.ansible.com/v1beta1
kind: AWX
metadata:
  name: awx-demo
spec:
  service_type: ClusterIP
```

다음 단계는 awx 연산자를 배포할 네임스페이스를 정의하는 것으로, `export NAMESPACE=awx` 명령을 사용한 다음 `make deploy`로 배포를 시작합니다.

![](/2022/Days/Images/Day69_config3.png)

확인을 통해 새 네임스페이스가 있고 네임스페이스에서 awx-operator-controller pod가 실행되고 있음을 확인합니다. `kubectl get pods -n awx`

![](/2022/Days/Images/Day69_config4.png)

복제된 리포지토리 내에서 awx-demo.yml이라는 파일을 찾을 수 있으며, 이제 이 파일을 Kubernetes 클러스터와 awx 네임스페이스에 배포하려고 합니다. `kubectl create -f awx-demo.yml -n awx`

![](/2022/Days/Images/Day69_config5.png)

`kubectl get pods -n awx -w`로 진행 상황을 계속 주시할 수 있습니다.

모든 것이 실행되면 아래와 같은 이미지와 비슷한 것이 보일 것입니다.

![](/2022/Days/Images/Day69_config6.png)

이제 새 터미널에서 `minikube service awx-demo-service --url -n $NAMESPACE`를 실행하여 Minikube 인그레스를 통해 이를 노출한 후 awx 배포에 액세스할 수 있어야 합니다.

![](/2022/Days/Images/Day69_config7.png)

그런 다음 해당 주소 []로 브라우저를 열면 사용자 이름과 비밀번호를 입력하라는 메시지가 표시되는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day69_config8.png)

기본적으로 사용자 이름은 admin이며, 비밀번호를 얻으려면 다음 명령어를 실행하여 `kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n awx| base64 --decode`를 얻을 수 있습니다.

![](/2022/Days/Images/Day69_config9.png)

이렇게 하면 중앙 집중화된 위치에서 playbook 및 구성 관리 작업을 관리할 수 있는 UI가 제공되며, 지금까지 하나의 Ansible 제어 스테이션에서 실행하던 것과 달리 팀원들이 함께 작업할 수 있습니다.

이 영역은 이 도구의 기능을 살펴보는 데 더 많은 시간을 할애할 수 있는 또 다른 영역 중 하나입니다.

Ansible AWX 사용에 대해 더 자세히 설명하는 Jeff Geerling의 훌륭한 리소스를 소개해 드리겠습니다. [Ansible 101 - 에피소드 10 - Ansible 타워 및 AWX](https://www.youtube.com/watch?v=iKmY4jEiy_A&t=752s)

이 비디오에서 그는 자동화 컨트롤러(이전 Ansible Tower)와 Ansible AWX(무료 및 오픈 소스)의 차이점에 대해서도 자세히 설명합니다.

### Ansible Vault

`ansible-vault`를 사용하면 Ansible 데이터 파일을 암호화하고 해독할 수 있습니다. 이 섹션에서는 일부 민감한 정보를 건너뛰고 일반 텍스트로 표시했습니다.

Ansible 바이너리에는 이 민감한 정보를 마스킹할 수 있는 `ansible-vault`가 내장되어 있습니다.

![](/2022/Days/Images/Day69_config10.png)

비밀 관리는 점차 HashiCorp Vault나 AWS 키 관리 서비스와 같은 도구와 함께 더 많은 시간을 할애해야 하는 또 다른 영역이 되었습니다. 이 부분은 더 자세히 살펴볼 영역으로 표시하겠습니다.

Jeff Geerling의 훌륭한 리소스와 데모를 다시 한번 링크해드리겠습니다 [Ansible 101 - 에피소드 6 - Ansible Vault와 역할](https://www.youtube.com/watch?v=JFweg2dUvqM).

### Ansible Galaxy(문서)

이제 데모 프로젝트의 일부 역할과 파일 구조를 생성하기 위해 이미 `ansible-galaxy`를 사용했습니다. 하지만 [Ansible Galaxy 문서](https://galaxy.ansible.com/docs/)도 있습니다.

"Galaxy는 Ansible 콘텐츠를 찾고 공유하기 위한 허브입니다."

### Ansible 테스트

- [Ansible Molecule](https://molecule.readthedocs.io/en/latest/) - Molecule 프로젝트는 Ansible 역할의 개발 및 테스트를 지원하도록 설계되었습니다.

- [Ansible Lint](https://ansible-lint.readthedocs.io/en/latest/) - playbook, 역할 및 컬렉션을 Lint하기 위한 CLI 도구입니다.

### 기타 리소스

- [Ansible 문서](https://docs.ansible.com/ansible/latest/index.html)

## 자료

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)
- [Your complete guide to Ansible](https://www.youtube.com/playlist?list=PLnFWJCugpwfzTlIJ-JtuATD2MBBD7_m3u)

위에 나열된 마지막 재생 목록은 이 섹션의 많은 코드와 아이디어가 나온 곳이며, 동영상 형식의 훌륭한 리소스이자 연습입니다.

이 포스팅에서는 구성 관리에 대해 살펴본 내용을 마무리하고, 다음에는 CI/CD 파이프라인과 애플리케이션 개발 및 릴리스에 이러한 workflow를 달성하기 위해 사용할 수 있는 몇 가지 도구 및 프로세스에 대해 살펴봅니다.

[Day 70](day70.md)에서 봐요!
