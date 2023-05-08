---
title: '#90DaysOfDevOps - Managing your Linux System, Filesystem & Storage - Day 16'
published: false
description: '90DaysOfDevOps - Managing your Linux System, Filesystem & Storage'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048702
---

## Linux 시스템, 파일 시스템 및 스토리지 관리

지금까지 우리는 리눅스와 데브옵스에 대한 간략한 개요를 살펴보았고, [Day 14](day14.md)에서는 실습 환경을 설정하였으며, [Day 15](day15.md)에서는 터미널에서 작업을 수행할 때 자주 사용되는 명령어들을 살펴보았습니다.

여기서는 업데이트, 소프트웨어 설치, 시스템 폴더의 용도와 스토리지에 대해 살펴보겠습니다. 이는 이해해야 할 핵심 영역 세 가지입니다.

## 우분투 및 소프트웨어 관리

가장 먼저 살펴볼 사항은 운영체제 업데이트 방법입니다. 대부분의 사용자는 Windows OS와 macOS에서 이 프로세스에 익숙할 것이며, Linux 데스크톱과 서버에서는 조금 다릅니다.

업데이트 및 소프트웨어 설치를 위해 Ubuntu VM에서 사용할 apt 패키지 관리자를 살펴볼 것입니다.

일반적으로, 개발 워크스테이션에서 소프트웨어를 설치하기 전에 중앙 리포지토리에서 사용 가능한 최신 업데이트가 있는지 확인하기 위해 이 명령을 실행하는 것이 좋습니다.

`sudo apt-get update`

![](/2022/Days/Images/Day16_Linux1.png)

이제 최신 OS 업데이트가 설치된 업데이트된 우분투 가상 머신이 생겼습니다. 이제 여기에 몇 가지 소프트웨어를 설치하려고 합니다.

텍스트 배너를 생성하는 프로그램인 `figlet`을 선택해 보겠습니다.

터미널에 `figlet`을 입력하면 시스템에 이 프로그램이 설치되어 있지 않다는 것을 알 수 있습니다.

![](/2022/Days/Images/Day16_Linux2.png)

위의 내용을 보면 시도해 볼 수 있는 몇 가지 `apt` 설치 옵션이 제공된다는 것을 알 수 있습니다. 기본 리포지토리에는 figlet이라는 프로그램이 있기 때문입니다. `sudo apt install figlet`을 시도해 보겠습니다.

![](/2022/Days/Images/Day16_Linux3.png)

이제 아래에서 볼 수 있듯이 `figlet` 앱을 사용할 수 있습니다.

![](/2022/Days/Images/Day16_Linux4.png)

`apt` 패키지 관리자를 사용하여 해당 소프트웨어나 다른 소프트웨어를 제거할 수도 있습니다.

`sudo apt remove figlet`

![](/2022/Days/Images/Day16_Linux5.png)

시스템에 추가할 수 있는 타사 리포지토리도 있는데, 바로 액세스할 수 있는 리포지토리는 Ubuntu 기본 리포지토리입니다.

Ubuntu 가상 머신에 Vagrant를 설치하려면 현재로서는 불가능합니다. 이를 확인하려면 아래의 첫 번째 명령어를 실행하면 됩니다. 그다음에는 HashiCorp 리포지토리를 신뢰하는 키를 추가하고, 리포지토리를 시스템에 추가해야 합니다.

![](/2022/Days/Images/Day16_Linux6.png)

HashiCorp 리포지토리를 추가했으면 이제 `sudo apt install vagrant`를 실행하여 시스템에 vagrant를 설치할 수 있습니다.

![](/2022/Days/Images/Day16_Linux7.png)

우분투에는 소프트웨어 설치와 관련하여 내장된 패키지 관리자를 사용할 수 있는 다양한 옵션이 있으며, 스냅을 이용한 소프트웨어 설치도 가능합니다.

이 글을 통해 Linux에서 OS 및 소프트웨어 설치를 관리하는 방법에 대한 감을 잡으셨기를 바랍니다.

## 파일 시스템 설명

Linux는 구성 파일들로 이루어져 있으며, 변경하고자 하는 내용이 있다면 해당 구성 파일을 수정하면 됩니다.

Windows 운영체제에서는 C 드라이브가 루트 디렉토리입니다. 반면 Linux 운영체제에서는 `/` 디렉토리가 중요한 위치로, 시스템 내의 여러 폴더를 찾을 수 있는 기본적인 디렉토리입니다.

![](/2022/Days/Images/Day16_Linux8.png)

- `/bin` - 바이너리의 줄임말로, 시스템에서 필요한 실행 파일, 도구 및 바이너리가 있는 폴더입니다.

![](/2022/Days/Images/Day16_Linux9.png)

- `/boot` - 시스템 부팅에 필요한 모든 파일이 위치합니다. 부팅 방법과 부팅할 드라이브를 찾을 수 있습니다.

![](/2022/Days/Images/Day16_Linux10.png)

- `/dev` - 이 폴더에서는 장치 정보를 찾을 수 있으며, 디스크 드라이브에 대한 포인터를 찾을 수 있습니다. 일반적으로 `sda`는 기본 OS 디스크입니다.

![](/2022/Days/Images/Day16_Linux11.png)

- `/etc` - 리눅스 시스템에서 가장 중요한 폴더로, 대부분의 구성 파일이 있는 곳입니다.

![](/2022/Days/Images/Day16_Linux12.png)

- `/home` - 사용자 폴더와 파일이 위치하는 곳입니다. 이 폴더에는 vagrant 사용자 폴더뿐만 아니라 명령어 섹션에서 작업한 `Document` 및 `Desktop` 폴더 있습니다.

![](/2022/Days/Images/Day16_Linux13.png)

- `/lib` - `/bin` 폴더가 바이너리 및 실행 파일을 가지고 있다면, 이 폴더는 이러한 파일들에 대한 공유 라이브러리를 가지고 있는 곳입니다.

![](/2022/Days/Images/Day16_Linux14.png)

- `/media` - 이동식 디바이스를 찾을 수 있는 곳입니다.

![](/2022/Days/Images/Day16_Linux15.png)

- `/mnt` - 임시 마운트 지점입니다. 스토리지 섹션에서 자세히 다루겠습니다.

![](/2022/Days/Images/Day16_Linux16.png)

- `/opt` - 옵션 소프트웨어 패키지가 위치하는 폴더입니다. 이 폴더에는 몇 가지 vagrant 및 가상 머신 패키지가 저장되어 있습니다.

![](/2022/Days/Images/Day16_Linux17.png)

- `/proc` - 커널 및 프로세스 정보를 찾을 수 있는 곳으로, `/dev`와 유사합니다.

![](/2022/Days/Images/Day16_Linux18.png)

- `/root` - 액세스 권한을 얻으려면 sudo를 사용해야 하는 루트의 홈 폴더입니다.

![](/2022/Days/Images/Day16_Linux19.png)

- `/run` - 애플리케이션 상태를 위한 자리 표시자입니다.

![](/2022/Days/Images/Day16_Linux20.png)

- `/sbin` - `bin` 폴더와 유사하지만 시스템에서 슈퍼유저 권한이 필요한 도구들이 위치합니다. 즉, `sudo bin`입니다.

![](/2022/Days/Images/Day16_Linux21.png)

- `/tmp` - 임시 파일이 위치하는 폴더입니다.

![](/2022/Days/Images/Day16_Linux22.png)

- `/usr` - 일반 사용자가 소프트웨어 패키지를 설치한 경우, 일반적으로 `/usr/bin` 위치에 설치됩니다.

![](/2022/Days/Images/Day16_Linux23.png)

- `/var` - 애플리케이션은 `bin` 폴더에 설치됩니다. 모든 로그 파일을 저장할 위치가 필요한데, 이 위치가 바로 `/var`입니다.

![](/2022/Days/Images/Day16_Linux24.png)

## 스토리지

Linux 시스템이나 다른 시스템에서도 사용 가능한 디스크와 해당 디스크의 여유 공간을 확인하는 것은 유용합니다. 아래의 명령어 몇 가지는 스토리지를 식별하고 사용 및 관리하는 데 도움이 됩니다.

- `lblk`는 블록 장치를 나열합니다. `sda`는 물리적 디스크이고 `sda1, sda2, sda3`는 해당 디스크의 파티션입니다.

![](/2022/Days/Images/Day16_Linux25.png)

- `df` 명령어는 파티션 정보, 전체 사용량 및 사용 가능한 용량에 대한 더 자세한 정보를 제공합니다. 다른 플래그를 사용하여 구문 분석할 수 있습니다. 일반적으로 `df -h` 명령어를 사용하여 인간이 읽기 쉬운 출력을 생성합니다.

![](/2022/Days/Images/Day16_Linux26.png)

새로운 디스크를 시스템에 추가할 경우, Windows에서와 마찬가지로 디스크 관리에서 디스크를 포맷해야 합니다. 그러나 Linux 터미널에서는 추가된 새 디스크와 관련된 sdb와 함께 `sudo mkfs -t ext4 /dev/sdb` 명령어를 사용하여 포맷할 수 있습니다.

새로 포맷한 디스크를 사용하기 위해서는 `/mnt` 폴더에서 `sudo mkdir NewDisk` 명령어로 디렉토리를 생성한 후, `sudo mount /dev/sdb NewDisk`를 사용하여 해당 위치에 디스크를 마운트해야 합니다.

시스템에서 스토리지를 안전하게 마운트 해제해야 할 때도 있고, 구성에서 그냥 가져와야 할 때도 있습니다. 이 작업은 `sudo umount /dev/sdb` 명령어로 수행할 수 있습니다.

만약에 당신이 해당 디스크를 마운트 해제하고 싶지 않고, 이 디스크를 데이터베이스나 지속적으로 사용할 다른 용도로 사용하려는 경우, 시스템을 재부팅 할 때 그 디스크가 있어야 합니다. 이를 위해, 이 디스크를 `/etc/fstab` 구성 파일에 추가하여 지속성을 유지해야 합니다. 그렇지 않으면 기계가 재부팅될 때 사용할 수 없으며 수동으로 위의 프로세스를 진행해야 합니다. 데이터는 여전히 디스크에 남아 있지만, 이 파일에 구성을 추가하지 않으면 자동으로 마운트되지 않습니다.

`fstab` 구성 파일을 편집한 후 `sudo mount -a` 명령을 실행하여 작동을 확인할 수 있습니다. 오류가 없다면, 이제 변경 사항이 재시작 시에도 지속됩니다.

텍스트 편집기를 사용하여 파일을 편집하는 방법은 다음 세션에서 다루겠습니다.

## 자료

- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

[Day 17](day17.md)에서 봐요!
