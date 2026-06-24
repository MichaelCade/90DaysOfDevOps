---
title: '#90DaysOfDevOps - Automate tasks with bash scripts - Day 19'
published: false
description: 90DaysOfDevOps - Automate tasks with bash scripts
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048774
---

## bash 스크립팅을 이용한 작업 자동화

오늘 사용할 shell은 bash이지만, 내일 ZSH에 대해 다룰 때 다른 shell도 다룰 예정입니다.

BASH - **B**ourne **A**gain **Sh**ell

프로그래밍 언어처럼 shell 스크립팅에도 거의 일주일 동안의 섹션을 쓸 수 있습니다. bash는 다른 자동화 도구와 함께 작업하는 기능을 제공합니다.

여전히 많은 사람들이 복잡한 shell 스크립트를 만들어 무언가를 실행하게 하고, 비즈니스에서 가장 중요한 부분을 이 스크립트에 의존하고 있다고 합니다. 하지만 이런 이유로 shell/bash 스크립팅을 이해해야 한다고 말하는 것은 아닙니다. 자동화 도구와 함께 사용하고 ad-hoc(특별한) 작업을 위해 shell/bash 스크립팅을 배워야 합니다.

이 섹션에서 사용한 예로, 매주 월요일 아침에 리눅스 VM을 새로 생성하고, 해당 리눅스 기계에 필요한 모든 소프트웨어 스택을 추가하는 등의 작업을 할 수 있는 간단한 bash 스크립트로 VAGRANTFILE을 감쌀 수 있습니다.

면접에서 스크립팅 경험 관련 질문이 점점 더 많아지고 있다는 소식도 종종 듣고 있습니다.

### 시작하기

이 90일 동안 다루는 많은 것들과 마찬가지로, 진정한 배움은 경험을 통해 알게 됩니다. 실제로 경험을 통해 배우면 기억에 오래 남습니다.

우선, 텍스트 에디터가 필요합니다. [Day 17](day17.md)에 가장 일반적인 텍스트 에디터 두 가지와 그것들을 어떻게 사용하는지에 대해 다뤘습니다.

첫 번째 shell 스크립트를 생성해봅시다.

`touch 90DaysOfDevOps.sh`

그다음 `nano 90DaysOfDevOps.sh`를 입력하면 nano에서 새로운 빈 shell 스크립트가 열립니다. 여기서 원하는 텍스트 에디터를 선택할 수 있습니다.

모든 bash 스크립트의 첫 줄은 다음과 같아야 합니다: `#!/usr/bin/bash`(bash 바이너리의 경로입니다.)

터미널에서 `which bash`를 실행하여 이를 확인해야 합니다. Ubuntu를 사용하지 않는 경우 터미널에서는 `whereis bash`도 시도해 볼 수 있습니다.

그러나 이미 생성된 shell 스크립트는 다른 경로들이 있을 수도 있습니다:

- `#!/bin/bash`
- `#!/usr/bin/env bash`

스크립트의 다음 줄에는 주석을 추가하여 스크립트의 목적이나 스크립트에 대한 정보를 제공합니다. `#`을 사용하면 코드의 특정 줄에 주석을 달고 다가올 명령어가 무엇을 할 것인지 설명할 수 있습니다. 주석을 많이 남길수록 사용자 경험이 향상된다고 생각합니다. 특히 공유할 때 그렇습니다.

리눅스 섹션에서 설치한 figlet 프로그램을 사용하여 스크립트에서 asci 아트를 만들 때도 있습니다.

![](/2022/Days/Images/Day19_Linux1.png)

이 리눅스 섹션에서 앞서 다룬 모든 명령어들([Day 15](day15.md))은 스크립트를 테스트하는 간단한 명령어로 사용할 수 있습니다.

스크립트에 간단한 코드 블록을 추가해 봅시다.

```bash
mkdir 90DaysOfDevOps
cd 90DaysOfDevOps
touch Day19
ls
```

텍스트 에디터를 저장하고 종료한 후, `./90DaysOfDevOps.sh`를 실행하면 권한이 거부되었다는 메시지가 표시됩니다. `ls -al` 명령을 사용하여 이 파일의 권한을 확인할 수 있는데 이 파일에 실행 권한이 없음을 확인할 수 있습니다.

![](/2022/Days/Images/Day19_Linux2.png)

`chmod +x 90DaysOfDevOps.sh`를 사용하여 이를 변경한 다음 `x`인 것을 봤을 때 스크립트를 실행할 수 있음을 확인할 수 있습니다.

![](/2022/Days/Images/Day19_Linux3.png)

`./90DaysOfDevOps.sh`를 사용하여 스크립트를 다시 실행할 수 있습니다. 스크립트를 실행한 후 새 디렉토리가 생성되고 해당 디렉토리로 이동한 다음 새 파일이 생성됩니다.

![](/2022/Days/Images/Day19_Linux4.png)

기본적인 것들이지만, 다른 도구를 호출하여 일상적인 작업을 쉽게 만들고 자동화하는 데 사용할 수 있다는 것을 알 수 있습니다.

### 변수, 조건문

이 섹션의 대부분은 Golang을 배울 때 다룬 것과 중복되지만, 여기에서 다시 살펴볼 가치가 있다고 생각합니다.

- ### 변수

변수를 사용하면 스크립트 내에서 반복되는 특정 용어를 한 번만 정의할 수 있습니다.

스크립트에 변수를 추가하려면 다음과 같이 새로운 줄에 입력하면 됩니다.

`challenge="90DaysOfDevOps"`

이렇게 하면 코드에서 `$challenge`를 사용할 때 변수가 변경되면 전체에 반영됩니다.

![](/2022/Days/Images/Day19_Linux5.png)

이제 `sh` 스크립트를 실행하면 스크립트에 추가된 출력이 표시됩니다.

![](/2022/Days/Images/Day19_Linux6.png)

또한 다음과 같이 사용자 입력을 요청하여 변수를 설정할 수도 있습니다:

```bash
echo "Enter your name"
read name
```

이렇게 하면 입력이 변수 `$name`으로 정의됩니다. 이후에 이를 사용할 수 있습니다.

- ### 조건문

우리는 도전 과제에 참여한 사람과 그들이 완료한 날짜 수를 알아보고 싶을 수도 있습니다. 이를 `if`, `if-else`, `else-if` 조건문을 사용해 정의할 수 있습니다. 아래 스크립트에 이를 정의한 것을 확인할 수 있습니다.

```bash
#!/bin/bash
#  ___   ___  ____                   ___   __ ____              ___
# / _ \ / _ \|  _ \  __ _ _   _ ___ / _ \ / _|  _ \  _____   __/ _ \ _ __  ___
#| (_) | | | | | | |/ _` | | | / __| | | | |_| | | |/ _ \ \ / / | | | '_ \/ __|
# \__, | |_| | |_| | (_| | |_| \__ \ |_| |  _| |_| |  __/\ V /| |_| | |_) \__ \
#   /_/ \___/|____/ \__,_|\__, |___/\___/|_| |____/ \___| \_/  \___/| .__/|___/
#                         |___/                                     |_|
#
# 이 스크립트는 bash 스크립팅을 시연하기 위한 것입니다!

# 정의할 변수

ChallengeName=#90DaysOfDevOps
TotalDays=90

# 사용자 입력

echo "Enter Your Name"
read name
echo "Welcome $name to $ChallengeName"
echo "How Many Days of the $ChallengeName challenge have you completed?"
read DaysCompleted

if [ $DaysCompleted -eq 90 ]
then
  echo "You have finished, well done"
elif [ $DaysCompleted -lt 90 ]
then
  echo "Keep going you are doing great"
else
  echo "You have entered the wrong amount of days"
fi
```

다음 단계로 이동하기 위해 값을 비교하거나 확인하는 과정을 위에서 볼 수 있습니다. 여기서 주목할 만한 다양한 옵션이 있습니다.

- `eq` - 두 값이 같으면 TRUE 반환
- `ne` - 두 값이 같지 않으면 TRUE 반환
- `gt` - 첫 번째 값이 두 번째 값보다 크면 TRUE 반환
- `ge` - 첫 번째 값이 두 번째 값보다 크거나 같으면 TRUE 반환
- `lt` - 첫 번째 값이 두 번째 값보다 작으면 TRUE 반환
- `le` - 첫 번째 값이 두 번째 값보다 작거나 같으면 TRUE 반환

파일 및 폴더에 대한 정보를 결정하기 위해 bash 스크립팅을 사용할 수도 있습니다. 이를 파일 조건이라고 합니다.

- `-d file` 파일이 디렉토리인 경우 True
- `-e file` 파일이 존재하는 경우 True
- `-f file` 제공된 문자열이 파일인 경우 True
- `-g file` 파일에 그룹 ID가 설정된 경우 True
- `-r file` 파일이 읽을 수 있는 경우 True
- `-s file` 파일의 크기가 0이 아닌 경우 True

```bash
FILE="90DaysOfDevOps.txt"
if [ -f "$FILE" ]
then
  echo "$FILE is a file"
else
  echo "$FILE is not a file"
fi
```

![](/2022/Days/Images/Day19_Linux7.png)

디렉토리에 파일이 아직 있다면 첫 번째 echo를 반환해야 합니다. 그러나 파일을 제거하면 두 번째 echo를 반환하게 됩니다.

![](/2022/Days/Images/Day19_Linux8.png)

특정 항목을 시스템에서 검색할 때 시간을 절약하는 데 사용할 수 있는 방법을 확인하실 수 있습니다.

GitHub에서 놀라운 저장소를 발견했습니다. 많은 스크립트를 가진 것 같습니다. [DevOps Bash Tools](https://github.com/HariSekhon/DevOps-Bash-tools/blob/master/README.md)

### 예제

**시나리오**: 우리 회사는 "90DaysOfDevOps"라는 이름을 가지고 있으며, 이제 확장할 때가 왔습니다. 1인 팀에서 다음 몇 주 동안 더 많은 사람으로 늘리려고 합니다. 지금까지 나만이 온보딩 과정을 알고 있으므로, 이러한 작업 중 일부를 자동화하여 병목 현상을 줄이고자 합니다.

**요구 사항**:

- 사용자 이름이 커맨드 라인 인수로 전달될 수 있습니다.
- 커맨드 라인 인수의 이름으로 사용자가 생성됩니다.
- password가 커맨드 라인 인수로 전달될 수 있습니다.
- 사용자의 password가 설정됩니다.
- 계정 생성이 성공적으로 이루어졌음을 나타내는 메시지가 표시됩니다.

먼저 `touch create_user.sh`를 사용하여 shell 스크립트를 생성합시다.

계속 진행하기 전에, `chmod +x create_user.sh`를 사용하여 이를 실행 가능하게 만들어줍시다.

그런 다음 `nano create_user.sh`를 사용하여 설정한 시나리오에 대한 스크립트 편집을 시작하겠습니다.

첫 번째 요구 사항인 "사용자 이름이 커맨드 라인 인수로 전달될 수 있습니다."는 다음과 같이 사용할 수 있습니다.

```bash
#! /usr/bin/bash

# 사용자 이름이 커맨드 라인 인수로 전달될 수 있습니다.
echo "$1"
```

![](/2022/Days/Images/Day19_Linux9.png)

이를 실행하려면 `./create_user.sh Michael`을 사용하십시오. Michael 대신에 당신의 이름을 사용하세요.

![](/2022/Days/Images/Day19_Linux10.png)

다음으로 두 번째 요구 사항인 "커맨드 라인 인수의 이름으로 사용자가 생성됩니다."를 처리할 수 있습니다. 이는 `useradd` 명령어를 사용하여 수행할 수 있습니다. `-m` 옵션은 사용자 홈 디렉토리를 /home/username으로 생성하기 위한 것입니다.

```bash
#! /usr/bin/bash

# 사용자 이름이 커맨드 라인 인수로 전달될 수 있습니다.
echo "$1 user account being created."

# 커맨드 라인 인수의 이름으로 사용자가 생성됩니다.
sudo useradd -m "$1"
```

경고: 사용자 계정 이름을 제공하지 않으면 `$1` 변수가 채워지지 않아 오류가 발생합니다.

계정이 생성되었는지 확인하려면 `awk -F: '{ print $1}' /etc/passwd` 명령을 사용합니다.

![](/2022/Days/Images/Day19_Linux11.png)

다음 요구 사항은 "password가 커맨드 라인 인수로 전달될 수 있습니다."입니다. 먼저, 이는 실제 환경에서는 절대 사용하지 않는 것이 좋습니다. 이는 단지 연습을 위한 것이며, 이를 통해 요구 사항을 이해할 수 있습니다.

```bash
#! /usr/bin/bash

# 사용자 이름이 커맨드 라인 인수로 전달될 수 있습니다.
echo "$1 user account being created."

# 커맨드 라인 인수의 이름으로 사용자가 생성됩니다.
sudo useradd -m "$1"

# password가 커맨드 라인 인수로 전달될 수 있습니다.
sudo chpasswd <<< "$1":"$2"
```

두 개의 파라미터로 이 스크립트를 실행하려면 `./create_user.sh 90DaysOfDevOps password`를 사용하세요.

아래 이미지에서 스크립트를 실행하여 사용자와 password를 생성한 다음, 수동으로 해당 사용자로 전환하고 `whoami` 명령으로 확인했다는 것을 알 수 있습니다.

![](/2022/Days/Images/Day19_Linux12.png)

마지막 요구 사항은 "계정 생성이 성공적으로 이루어졌음을 나타내는 메시지가 표시됩니다."입니다. 우리는 이미 코드의 첫 번째 줄에 이를 포함하고 있으며, 위의 스크린샷에서 `90DaysOfDevOps user account being created`라는 메시지가 표시되는 것을 볼 수 있습니다. 이는 `$1` 매개변수로 테스트할 때 남겨진 것입니다.

이제 이 스크립트를 사용하여 리눅스 시스템에 새 사용자를 빠르게 온보딩하고 설정할 수 있습니다. 하지만 기존 사용자 중 몇 명이 이 과정을 거쳐 다른 사람들에게 새 사용자 이름이나 password를 알려주는 대신 앞서 다룬 사용자 입력을 추가하여 변수를 수집할 수 있습니다.

```bash
#! /usr/bin/bash

echo "What is your intended username?"
read  username
echo "What is your password"
read  password

# 사용자 이름이 커맨드 라인 인수로 전달될 수 있습니다.
echo "$username user account being created."

# 커맨드 라인 인수의 이름으로 사용자가 생성됩니다.
sudo useradd -m $username

# password가 커맨드 라인 인수로 전달될 수 있습니다.
sudo chpasswd <<< $username:$password
```

단계를 더욱 인터랙티브하게 만들면

![](/2022/Days/Images/Day19_Linux14.png)

마무리하기 위해, 새 사용자 계정이 생성되었다는 성공적인 출력을 원할 수도 있습니다.

![](/2022/Days/Images/Day19_Linux15.png)

한 가지 눈에 띄는 점은 입력 시 비밀번호가 표시된다는 점인데, 코드 `read -s password`에서의 `-s` 플래그를 사용하여 이를 숨길 수 있습니다.

![](/2022/Days/Images/Day19_Linux16.png)

실험 목적으로 생성한 사용자를 삭제하려면 `sudo userdel test_user`를 사용할 수 있습니다.

[예제 스크립트](/2022/Days/Linux/create-user.sh)

한 번 더 강조하고 싶은 바는, 이것은 shell 스크립팅이 유연하게 활용될 수 있는 다양한 용도를 갖는다는 것을 강조하기 위해 만들어진 것이라는 것입니다.

매일, 매주, 매월 반복되는 작업이 어떤 것이 있는지 생각해보고, 그것을 어떻게 더 잘 자동화할 수 있을지 생각해보세요. 첫 번째 옵션은 주로 bash 스크립트를 사용하는 것이며, 그 이후에는 더 복잡한 영역으로 이동할 수 있습니다.

저는 제 로컬 머신에서 minikube를 사용하여 쿠버네티스 클러스터를 빠르게 생성하고 데이터 서비스 및 Kasten K10을 사용하여 데이터 관리에 대한 요구 사항과 필요성을 보여주기 위한 매우 간단한 [bash 파일](https://github.com/MichaelCade/project_pace/blob/main/singlecluster_demo.sh)을 생성했습니다만, 아직 쿠버네티스를 다루지 않았기 때문에 여기서 다루는 것이 적절하지 않다고 생각합니다.

## 자료

- [Bash in 100 seconds](https://www.youtube.com/watch?v=I4EWvMFj37g)
- [Bash script with practical examples - Full Course](https://www.youtube.com/watch?v=TPRSJbtfK4M)
- [Client SSH GUI - Remmina](https://remmina.org/)
- [The Beginner's guide to SSH](https://www.youtube.com/watch?v=2QXkrLVsRmk)
- [Vim in 100 Seconds](https://www.youtube.com/watch?v=-txKSRn0qeA)
- [Vim tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I)
- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

[Day 20](day20.md)에서 봐요!
