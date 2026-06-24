---
title: '#90DaysOfDevOps - Text Editors - nano vs vim - Day 17'
published: false
description: 90DaysOfDevOps - Text Editors - nano vs vim
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048703
---

## 텍스트 편집기 - nano vs vim

대부분의 Linux 시스템은 서버로 사용되며, 이러한 시스템에는 GUI가 없습니다. 또한 이전 세션에서 언급했듯이, Linux 시스템은 대부분 구성 파일로 구성되어 있어 변경하려면 이러한 구성 파일을 편집하여 시스템의 모든 것을 변경해야 합니다.

다양한 옵션이 있지만, 가장 일반적으로 사용되는 두 가지 터미널 텍스트 편집기를 소개하고자 합니다. 제 경험상, 빠른 수정 작업에는 `nano`가 더 편리하지만, `vim`은 다양한 기능을 제공하여 더 복잡한 편집 작업을 수행할 수 있습니다.

### nano

- 모든 시스템에서 사용할 수 있는 것은 아닙니다.
- 처음 시작할 때 좋습니다.

`nano 90DaysOfDevOps.txt`를 실행하면 아무 내용도 없는 새 파일이 생성되며, 이 파일에 텍스트를 추가할 수 있습니다. 하단에 도움말 메뉴가 있습니다.

![](/2022/Days/Images/Day17_Linux1.png)

이제 `Ctrl X + enter`를 입력한 다음 `ls`를 실행하면 새 텍스트 파일을 볼 수 있습니다.

![](/2022/Days/Images/Day17_Linux2.png)

이제 해당 파일을 읽기 위해 `cat` 명령어를 실행할 수 있습니다. 그리고 `nano 90DaysOfDevOps.txt` 명령어를 사용하여 파일에 추가 텍스트를 입력하거나 수정할 수 있습니다.

제게 있어서 nano는 구성 파일에 작은 변경 사항을 적용하기 매우 편리한 편집기입니다.

### vim

1976년에 출시된 UNIX 텍스트 편집기인 vi의 형제 버전인 vim은 다양한 기능을 제공하여 가장 일반적인 텍스트 편집기 중 하나일까요?

- 대부분의 Linux 배포판에서 지원됩니다.
- vim만 다루는 7시간짜리 강좌를 찾을 수 있다는 것은 매우 놀랍습니다!

`vim` 명령어를 사용하여 vim 에디터로 이동하거나 새 텍스트 파일을 편집하려면 `vim 90DaysOfDevOps.txt`와 같은 명령어를 실행할 수 있습니다. 그러나 이 경우 하단에 도움말 메뉴가 없음을 확인할 수 있습니다.

첫 번째 질문은 "vim을 어떻게 종료할까요?"일 수 있습니다. 종료 방법으로는 `ESC` 키가 있으며, 변경 사항이 없다면 `:q`를 입력하면 됩니다.

![](/2022/Days/Images/Day17_Linux3.png)

기본적으로 `normal` 모드에서 시작하고 `command, normal, visual, insert` 모드도 있습니다. 텍스트를 추가하려면 `i`를 눌러 `normal` 모드에서 `insert` 모드로 전환해야 합니다. 변경 사항을 저장하고 텍스트를 추가한 후에는 `ESC` 키를 누르고 `:wq`를 입력하면 됩니다.

![](/2022/Days/Images/Day17_Linux4.png)

![](/2022/Days/Images/Day17_Linux5.png)

`cat` 명령으로 변경 사항을 저장했는지 확인할 수 있습니다.

vim은 단축키만 알고 있다면 작은 작업도 매우 빠르게 수행할 수 있는 멋진 기능이 있습니다. 예를 들어, 반복되는 단어 목록을 추가한 후 이를 변경해야 한다고 가정해 봅시다. 구성 파일에서 네트워크 이름이 반복되는 경우, 이를 빠르게 변경하고 싶을 수 있습니다. 이 예시에서는 'day'라는 단어를 사용하고 있습니다.

![](/2022/Days/Images/Day17_Linux6.png)

'day'를 '90DaysOfDevOps'로 변경하려면, `ESC` 키를 누르고 `:%s/Day/90DaysOfDevOps`를 입력하면 됩니다.

![](/2022/Days/Images/Day17_Linux7.png)

이제 Enter 키를 누르면 'day'라는 단어가 '90DaysOfDevOps로' 바뀝니다.

![](/2022/Days/Images/Day17_Linux8.png)

복사 및 붙여넣기는 저에게 정말 큰 도움이 됐습니다. normal 모드에서는 키보드의 `yy`를 사용하여 복사할 수 있으며, 같은 줄에 `p`를 사용하여 붙여넣기하거나, 새 줄에 `P`를 사용하여 붙여넣기할 수 있습니다.

특정 줄을 삭제하려면 삭제할 줄 수를 선택한 후 `dd`를 입력하면 됩니다.

때로는 파일을 검색해야 할 필요가 있습니다. 이전 세션에서 언급한 것처럼 `grep`을 사용할 수 있지만, `/word`를 사용하면 첫 번째 일치 항목을 찾을 수 있으며, `n` 키를 사용하여 다음 항목으로 이동할 수 있습니다.

제가 드릴 수 있는 가장 큰 조언은 가능한 한 직접 사용해 보라는 것입니다.

일반적으로 면접에서 물어보는 질문 중 하나는, "Linux에서 가장 좋아하는 텍스트 편집기는 무엇인가요?" 입니다. 이 질문에 대답하려면 적어도 두 가지 이상의 텍스트 편집기를 알고 있어야 하며, nano에 대한 대답도 괜찮습니다. 이 질문에 대답하는 것은 적어도 텍스트 편집기가 무엇인지 이해했음을 보여주는 것입니다. 하지만 더욱 능숙하게 사용하려면 직접 사용해 보는 것이 좋습니다.

vim에서 방향키로 사용할 수 있는 키는 화살표 키뿐만 아니라 `H,J,K,L`도 사용할 수 있습니다.

## 자료

- [Vim in 100 Seconds](https://www.youtube.com/watch?v=-txKSRn0qeA)
- [Vim tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I)
- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

[Day 18](day18.md)에서 봐요!
