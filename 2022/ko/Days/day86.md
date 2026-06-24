---
title: '#90DaysOfDevOps - Backup all the platforms - Day 86'
published: false
description: 90DaysOfDevOps - Backup all the platforms
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049058
---

## 모든 플랫폼 백업

이번 챌린지를 진행하는 동안 다양한 플랫폼과 환경에 대해 논의했습니다. 이 모든 플랫폼과 환경의 공통점은 모두 어느 정도의 데이터 보호가 필요하다는 사실입니다!

데이터 보호는 수년 전부터 존재해 왔지만, 오늘날의 풍부한 데이터와 이러한 데이터가 가져다주는 가치는 여러 노드와 애플리케이션 전반의 고가용성을 통해 인프라 장애에 대한 복원력을 확보해야 할 뿐만 아니라 장애 시나리오가 발생할 경우 안전하고 안전한 위치에 중요한 데이터의 사본이 필요하다는 점을 고려해야 한다는 것을 의미합니다.

요즘 사이버 범죄와 랜섬웨어에 대한 이야기를 많이 듣는데, 이는 엄청난 위협이며 랜섬웨어의 공격을 받을 수 있다는 사실에 대해 오해하지 마세요. 랜섬웨어는 시기의 문제가 아닙니다. 따라서 그 시기가 닥쳤을 때를 대비해 데이터를 안전하게 보호해야 하는 이유는 더욱 커집니다. 그러나 데이터 손실의 가장 흔한 원인은 랜섬웨어나 사이버 범죄가 아니라 실수로 인한 삭제입니다!

우리 모두는 실수로 삭제하지 말았어야 할 파일을 삭제하고 순간적으로 후회한 경험이 있을 것입니다.

챌린지 기간 동안 논의한 모든 기술과 자동화를 통해 상태 저장 데이터 또는 복잡한 stateless 구성을 보호해야 하는 요구 사항은 플랫폼에 관계없이 여전히 존재합니다.

![](/2022/Days/Images/Day86_Data1.png)

하지만 자동화를 염두에 두고 데이터 보호를 수행할 수 있어야 하며 워크플로우에 통합할 수 있어야 합니다.

백업이 무엇인지 살펴보면

_정보 기술에서 백업 또는 데이터 백업은 데이터 손실이 발생한 후 원본을 복원하는 데 사용할 수 있도록 컴퓨터 데이터의 복사본을 다른 곳에 저장하는 것입니다. 이러한 과정을 나타내는 동사 형태는 "back up"이고, 명사 및 형용사 형태는 "백업"입니다._

이를 가장 간단한 형태로 분류하면 백업은 데이터를 복사하여 새 위치에 붙여 넣는 것입니다. 간단히 말해, 지금 당장 C: 드라이브에 있는 파일을 D: 드라이브로 복사하여 백업을 만들면 C: 드라이브에 문제가 발생하거나 파일 내에서 잘못 편집된 부분이 있을 때를 대비해 복사본을 확보할 수 있습니다. D: 드라이브에 있는 사본으로 되돌릴 수 있습니다. 이제 C 드라이브와 D 드라이브가 모두 있는 곳에서 컴퓨터가 죽으면 보호되지 않으므로 시스템 외부의 솔루션이나 집에 있는 NAS 드라이브에 데이터 사본을 저장하는 방법을 고려해야 하나요? 하지만 집에 무슨 일이 생기면 다른 위치의 다른 시스템에 저장하는 것을 고려해야 할 수도 있고, 클라우드가 옵션이 될 수도 있습니다. 장애 위험을 줄이기 위해 중요한 파일의 사본을 여러 위치에 저장할 수 있을까요?

### 3-2-1 백업 방법론

이제 3-2-1 규칙 또는 백업 방법론에 대해 이야기하기에 좋은 시기인 것 같습니다. 이 주제를 다룬 [라이트닝 토크](https://www.youtube.com/watch?v=5wRt1bJfKBw)를 진행한 적이 있습니다.

데이터를 보호해야 하는 이유에 대해서는 이미 몇 가지 극단적인 예를 들었지만, 아래에 몇 가지를 더 말씀드리겠습니다:

![](/2022/Days/Images/Day86_Data2.png)

이제 3-2-1 방법론에 대해 말씀드리겠습니다. 데이터의 첫 번째 복사본 또는 백업은 가능한 한 프로덕션 시스템에 가깝게 저장해야 하는데, 그 이유는 복구 속도를 고려한 것이며, 다시 실수로 삭제하는 경우가 가장 흔한 복구 이유가 될 것이라는 원래의 이야기로 돌아가겠습니다. 하지만 원본 또는 프로덕션 시스템 외부의 적절한 보조 미디어에 저장하고 싶습니다.

그런 다음 다른 집, 건물, 데이터 센터 또는 퍼블릭 클라우드 등 두 번째 위치로 데이터 사본을 외부 또는 오프사이트로 전송하고 싶습니다.

![](/2022/Days/Images/Day86_Data3.png)

### 백업 책임

백업을 하지 않아도 된다는 속설에 대해 "모든 것이 stateless 상태"라는 말을 들어보셨을 것입니다. 모든 것이 stateless 상태라면 비즈니스는 무엇일까요? 데이터베이스도 없고, 워드 문서도 없나요? 비즈니스 내의 모든 개인이 데이터를 보호해야 할 책임이 있지만, 미션 크리티컬 애플리케이션과 데이터에 대한 백업 프로세스를 제공하는 것은 대부분 운영 팀에서 담당하게 될 것입니다.

데이터베이스에 실수를 해서 클러스터의 모든 노드에 복제되거나 화재, 홍수, 피의 시나리오로 인해 클러스터를 더 이상 사용할 수 없게 되어 중요한 데이터가 손실되는 경우를 제외하고는 "고가용성은 내 백업이고, 클러스터에 여러 노드를 구축했으니 절대 다운될 일은 없습니다!"라고 생각하는 것도 좋은 생각일 수 있습니다. 고집을 부리는 것이 아니라 데이터와 서비스를 인식하는 것이 중요하며, 모든 사람이 아키텍처에 고가용성 및 내결함성을 고려해야 하지만 이것이 백업의 필요성을 대체할 수는 없습니다!

복제는 데이터의 오프사이트 복사본을 제공하는 것처럼 보일 수 있으며, 위에서 언급한 클러스터가 여러 위치에 분산되어 있을 수도 있지만 첫 번째 실수는 여전히 그곳에 복제될 것입니다. 하지만 백업 요구사항은 환경 내에서 애플리케이션 복제 또는 시스템 복제와 함께 고려해야 합니다.

이제 이 모든 것을 말했듯이 다른 쪽에서도 극단적으로 나아가 너무 많은 위치에 데이터 사본을 전송할 수 있으며, 이는 비용뿐만 아니라 표면 영역이 크게 확장되어 공격받을 위험이 높아집니다.

어쨌든 백업은 누가 관리하나요? 각 비즈니스마다 다르겠지만 누군가는 백업 요구 사항을 이해해야 합니다. 그리고 복구 계획도 이해해야 합니다!

### 모두가 신경 쓸 때까지 아무도 신경 쓰지 않습니다.

백업은 대표적인 예로, 무언가를 복원해야 할 때까지 아무도 백업에 대해 신경 쓰지 않습니다. 데이터를 백업해야 한다는 요구 사항과 함께 복원 방법도 고려해야 합니다!

텍스트 문서의 예에서는 아주 작은 파일에 대해 이야기하고 있으므로 앞뒤로 복사하는 기능이 쉽고 빠릅니다. 하지만 100GB 이상의 파일에 대해 이야기하는 경우에는 시간이 오래 걸립니다. 또한 가상 머신을 예로 들면 복구해야 하는 수준도 고려해야 합니다.

전체 가상 머신이 있고 운영 체제, 애플리케이션이 설치되어 있으며 데이터베이스 서버인 경우 일부 데이터베이스 파일도 있습니다. 실수로 데이터베이스에 잘못된 코드 줄을 삽입한 경우 전체 가상 머신을 복원할 필요는 없으며, 복구 대상만 세분화하여 복구하고 싶습니다.

### 백업 시나리오

이제 일부 데이터를 보호하는 시나리오를 구축해 보겠습니다. 특히, 로컬 컴퓨터(이 경우에는 Windows이지만 제가 사용할 도구는 무료 오픈 소스일 뿐만 아니라 크로스 플랫폼도 지원합니다)에서 일부 파일을 보호하고 싶습니다. 집에 로컬로 있는 NAS 장치뿐만 아니라 클라우드의 Object Storage 버킷에도 보호되도록 하고 싶습니다.

이 중요한 데이터를 백업하고 싶은데, 마침 90DaysOfDevOps의 저장소이기도 하고, 지금 이 글을 읽고 있는 GitHub로도 전송되고 있는데, 만약 내 컴퓨터가 죽고 GitHub가 다운되면 어떻게 될까요? 다른 사람이 어떻게 콘텐츠를 읽을 수 있을 뿐만 아니라 해당 데이터를 다른 서비스로 어떻게 복원할 수 있을까요?

![](/2022/Days/Images/Day86_Data5.png)

이를 달성하는 데 도움이 되는 많은 도구가 있지만, 저는 백업을 암호화, 중복 제거 및 압축하는 동시에 여러 위치로 백업을 보낼 수 있는 오픈 소스 백업 도구인 [Kopia](https://kopia.io/)라는 도구를 사용하려고 합니다.

[여기](https://github.com/kopia/kopia/releases)에서 릴리스를 다운로드할 수 있으며, 이 글을 쓰는 시점에서는 v0.10.6을 사용할 것입니다.

### Kopia 설치하기

Kopia CLI와 GUI가 있는데, 여기서는 GUI를 사용하겠지만 GUI를 제공하지 않는 리눅스 서버의 경우 CLI 버전도 사용할 수 있다는 것을 알고 있습니다.

저는 `KopiaUI-Setup-0.10.6.exe`를 사용하겠습니다.

다음번 설치는 매우 빠르게 진행되며, 애플리케이션을 열면 백업 저장소로 사용할 스토리지 유형을 선택할 수 있는 선택 항목이 표시됩니다.

![](/2022/Days/Images/Day86_Data6.png)

### 리포지토리 설정하기

먼저 로컬 NAS 장치를 사용하여 리포지토리를 설정하고 SMB를 사용하여 이 작업을 수행하려고 하지만 NFS를 사용할 수도 있습니다.

![](/2022/Days/Images/Day86_Data7.png)

다음 화면에서는 비밀번호를 정의할 것이며, 이 비밀번호는 리포지토리 콘텐츠를 암호화하는 데 사용됩니다.

![](/2022/Days/Images/Day86_Data8.png)

이제 리포지토리가 구성되었으므로 임시 스냅샷을 트리거하여 리포지토리에 데이터 쓰기를 시작할 수 있습니다.

![](/2022/Days/Images/Day86_Data9.png)

먼저 스냅샷할 대상의 경로를 입력해야 하며, 여기서는 `90DaysOfDevOps` 폴더의 복사본을 만들고자 합니다. 곧 스케줄링 측면으로 돌아가겠습니다.

![](/2022/Days/Images/Day86_Data10.png)

스냅샷 보존을 정의할 수 있습니다.

![](/2022/Days/Images/Day86_Data11.png)

제외하려는 파일이나 파일 형식이 있을 수 있습니다.

![](/2022/Days/Images/Day86_Data12.png)

일정을 정의하려면 이다음 화면에서 할 수 있으며, 이 스냅샷을 처음 생성할 때 이 페이지가 정의할 첫 페이지입니다.

![](/2022/Days/Images/Day86_Data13.png)

그리고 여기에서 처리할 수 있는 몇 가지 다른 설정이 표시됩니다.

![](/2022/Days/Images/Day86_Data14.png)

지금 스냅샷을 선택하면 데이터가 리포지토리에 기록됩니다.

![](/2022/Days/Images/Day86_Data15.png)

### S3로 오프사이트 백업

Kopia에서는 UI를 통해 한 번에 하나의 리포지토리만 구성할 수 있는 것처럼 보입니다. 그러나 UI를 통해 창의력을 발휘하여 여러 리포지토리 구성 파일을 선택하여 로컬 및 오프사이트에 복사본을 오브젝트 스토리지에 저장하려는 목표를 달성할 수 있습니다.

제가 데이터를 전송하기 위해 선택한 오브젝트 스토리지는 Google 클라우드 스토리지입니다. 먼저 Google 클라우드 플랫폼 계정에 로그인하고 스토리지 버킷을 생성했습니다. 시스템에 이미 구글 클라우드 SDK가 설치되어 있었지만 `gcloud auth application-default login`을 실행하면 계정으로 인증되었습니다.

![](/2022/Days/Images/Day86_Data16.png)

그런 다음 Kopia의 CLI를 사용하여 이전 단계에서 SMB 리포지토리를 추가한 후 리포지토리의 현재 상태를 표시했습니다. `C:\Program Files\KopiaUI\resources\server\kopia.exe" --config-file=C:\Users\micha\AppData\Roaming\kopia\repository.config repository status` 명령을 사용하여 이 작업을 수행했습니다.

![](/2022/Days/Images/Day86_Data17.png)

이제 데모를 위해 리포지토리에 대한 구성을 교체할 준비가 되었습니다. 이 두 리포지토리에 모두 적용하는 장기적인 솔루션을 원한다면 `smb.config` 파일과 `object.config` 파일을 생성하고 이 두 명령을 실행하여 데이터 사본을 각 위치로 전송할 수 있어야 할 것입니다. 리포지토리를 추가하기 위해 `"C:\Program Files\KopiaUI\resources\server\kopia.exe" --config-file=C:\Users\micha\AppData\Roaming\kopia\repository.config repository create gcs --bucket 90daysofdevops`를 실행합니다.

위의 명령은 우리가 생성한 구글 클라우드 스토리지 버킷의 이름이 `90daysofdevops`라는 것을 고려합니다.

![](/2022/Days/Images/Day86_Data18.png)

이제 새 리포지토리를 생성했으므로 `"C:\Program Files\KopiaUI\resources\server\kopia.exe" --config-file=C:\Users\micha\AppData\Roaming\kopia\repository.config repository status` 명령을 다시 실행하면 이제 GCS 리포지토리 구성이 표시됩니다.

![](/2022/Days/Images/Day86_Data19.png)

다음으로 해야 할 일은 스냅샷을 생성하여 새로 생성한 리포지토리로 전송하는 것입니다. `"C:\Program Files\KopiaUI\resources\server\kopia.exe" --config-file=C:\Users\micha\AppData\Roaming\kopia\repository.config kopia snapshot create "C:\Users\micha\demo\90DaysOfDevOps"` 명령을 사용하여 이 프로세스를 시작할 수 있습니다. 아래 브라우저에서 이제 Google 클라우드 스토리지 버킷에 백업에 기반한 kopia 파일이 있는 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day86_Data20.png)

위의 프로세스를 통해 중요한 데이터를 두 개의 다른 위치로 전송해야 하는 요구 사항을 해결할 수 있으며, 그중 하나는 Google Cloud Storage의 오프사이트에 있고 물론 다른 미디어 유형에 데이터의 프로덕션 사본이 여전히 있습니다.

### 복원

복원은 또 다른 고려 사항이며 매우 중요한 기능으로, Kopia는 기존 위치뿐만 아니라 새로운 위치로도 복원할 수 있는 기능을 제공합니다.

`"C:\Program Files\KopiaUI\resources\server\kopia.exe" --config-file=C:\Users\micha\AppData\Roaming\kopia\repository.config snapshot list` 명령을 실행하면 현재 구성된 리포지토리(GCS)에 있는 스냅샷이 나열됩니다.

![](/2022/Days/Images/Day86_Data21.png)

그런 다음 `"C:\Program Files\KopiaUI\resources\server\kopia.exe" --config-file=C:\Users\micha\AppData\Roaming\kopia\repository.config mount all Z:` 명령을 사용하여 GCS에서 직접 해당 스냅샷을 마운트할 수 있습니다.

![](/2022/Days/Images/Day86_Data22.png)

또한 `kopia snapshot restore kdbd9dff738996cfe7bcf99b45314e193`을 사용하여 스냅샷 내용을 복원할 수도 있습니다.

위의 명령어가 매우 긴데, 이는 실습 상단에 설명한 대로 KopiaUI 버전의 kopia.exe를 사용했기 때문이며, kopia.exe를 다운로드하여 경로에 넣으면 `kopia` 명령어만 사용하면 됩니다.

다음 세션에서는 Kubernetes 내에서 워크로드를 보호하는 데 중점을 두겠습니다.

## 자료

- [Kubernetes Backup and Restore made easy!](https://www.youtube.com/watch?v=01qcYSck1c4&t=217s)
- [Kubernetes Backups, Upgrades, Migrations - with Velero](https://www.youtube.com/watch?v=zybLTQER0yY)
- [7 Database Paradigms](https://www.youtube.com/watch?v=W2Z7fbCLSTw&t=520s)
- [Disaster Recovery vs. Backup: What's the difference?](https://www.youtube.com/watch?v=07EHsPuKXc0)
- [Veeam Portability & Cloud Mobility](https://www.youtube.com/watch?v=hDBlTdzE6Us&t=3s)

[Day 87](day87.md)에서 봐요!
