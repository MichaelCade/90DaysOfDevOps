---
title: '#90DaysOfDevOps - Setting up your DevOps environment for Go & Hello World - Day 8'
published: false
description: 90DaysOfDevOps - Setting up your DevOps environment for Go & Hello World
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048857
---

## Go를 위한 데브옵스 환경 설정 및 Hello World

Go 언어에 대한 몇 가지 기본 사항을 살펴보기 전에, 먼저 컴퓨터에 Go를 설치하고 모든 "learning programming 101" 모듈에서 가르치는 Hello World 앱을 만들어야 합니다. 이 문서에서는 컴퓨터에 Go를 설치하는 단계를 안내하며, 그림으로 문서화하여 사람들이 쉽게 따라 할 수 있도록 하겠습니다.

먼저, [go.dev/dl](https://go.dev/dl/)로 이동하면 다운로드할 수 있는 몇 가지 옵션이 나타납니다.

![](/2022/Days/Images/Day8_Go1.png)

여기까지 오셨다면, 현재 사용 중인 컴퓨터의 운영 체제를 확인하고 해당하는 다운로드를 선택하여 설치를 시작하시면 됩니다. 제가 사용하는 운영 체제는 Windows입니다. 기본적으로 다음 화면부터는 모든 기본값을 그대로 두어도 됩니다. **(최신 버전이 나오면 스크린샷과 다르게 나올 수도 있으니 참고해주세요.)**

![](/2022/Days/Images/Day8_Go2.png)

만약 이전 버전의 Go가 설치되어 있다면, 먼저 해당 버전을 제거해야 합니다. Windows에서는 이를 위한 내장 제거 기능을 제공하며, 이를 통해 이전 버전을 제거하고 하나의 설치 프로그램으로 설치할 수 있습니다.

작업을 완료한 후 명령 프롬프트/터미널을 열어서 Go가 설치되어 있는지 확인해야 합니다. 만약 아래와 같은 출력이 표시되지 않았다면, Go가 설치되지 않은 것이므로 단계를 되짚어봐야 합니다.

`go version`

![](/2022/Days/Images/Day8_Go3.png)

다음으로 Go 환경변수를 확인하겠습니다. 작업 디렉토리가 올바르게 구성되었는지 항상 확인하는 것이 좋으므로, 아래 디렉토리가 시스템에 있는지 확인해야 합니다.

![](/2022/Days/Images/Day8_Go4.png)

확인하셨나요? 지금까지 잘 따라오고 있나요? 해당 디렉토리로 이동하면 아래와 같은 메시지가 나타납니다.

![](/2022/Days/Images/Day8_Go5.png)

그럼 PowerShell 터미널에서 mkdir 명령어를 사용하여 디렉토리를 만들어 보겠습니다. 또한, 아래에서 보이듯이 Go 폴더 내에 3개의 폴더를 생성해야 합니다.

![](/2022/Days/Images/Day8_Go6.png)

이제 Go를 설치하고 작업 디렉토리를 준비했습니다. 이제 통합 개발 환경(IDE)이 필요합니다. 사용 가능한 IDE는 여러 가지가 있지만, 가장 일반적이고 제가 사용하는 것은 VSCode(Visual Studio Code)입니다. IDE에 대한 자세한 내용은 [여기](https://www.youtube.com/watch?v=vUn5akOlFXQ)에서 확인할 수 있습니다. _(엄밀히 말하면 VSCode는 IDE가 아닌 코드 에디터입니다. - 옮긴이)_

만약 컴퓨터에 아직 VSCode를 다운로드하고 설치하지 않았다면, [여기](https://code.visualstudio.com/download)로 이동하여 설치할 수 있습니다. 아래에서 확인할 수 있는 것처럼, 다양한 운영체제를 제공됩니다.

![](/2022/Days/Images/Day8_Go7.png)

Go 설치와 마찬가지로, 기본값으로 다운로드하여 설치합니다. 설치가 완료되면, 파일 열기를 선택하고 Go 디렉토리를 만들었던 곳으로 이동하여 VSCode를 열 수 있습니다.

![](/2022/Days/Images/Day8_Go8.png)

신뢰에 관한 팝업이 표시될 수 있는데, 원한다면 이를 읽고 '예, 작성자를 신뢰합니다'를 누르세요. (나중에 신뢰하지 않는 파일을 열기 시작해도 저는 책임지지 않습니다!)

이전에 만든 세 개의 폴더를 모두 볼 수 있습니다. 이제 `src` 폴더를 마우스 오른쪽 버튼으로 클릭하고 `Hello`라는 새 폴더를 생성하겠습니다.

![](/2022/Days/Images/Day8_Go9.png)

여기까지는 꽤 쉬운 내용이었죠? 이제 다음 단계에서는 아무것도 이해하지 못한 상태에서 첫 번째 Go 프로그램을 만들 것입니다.

다음으로, `Hello` 폴더에 `main.go` 파일을 생성합니다. `main.go` 파일을 열고 Enter 키를 누르면 Go 확장 프로그램과 패키지를 설치할 것인지 묻는 메시지가 나타납니다. 이전에 만든 빈 `pkg` 폴더를 확인하면 이제 새로운 패키지가 추가된 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day8_Go10.png)

이제 Hello World 앱을 실행하고 다음 코드를 새 `main.go` 파일에 복사하여 저장해 보겠습니다.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello #90DaysOfDevOps")
}
```

이해가 어려울 수도 있지만, 함수, 패키지 등에 대해 더 자세하게 다룰 예정입니다. 현재는 앱을 실행해보겠습니다. 터미널에서 `Hello` 폴더로 돌아가서 제대로 작동하는지 확인할 수 있습니다. 다음 명령을 사용하여 프로그램이 제대로 작동하는지 확인할 수 있습니다.

```
go run main.go
```

![](/2022/Days/Images/Day8_Go11.png)

여기서 끝이 아닙니다. 다른 Windows 컴퓨터에서 이 프로그램을 실행하려면 어떻게 해야 할까요? 다음 명령을 사용하여 바이너리로 빌드하면 됩니다.

```
go build main.go
```

![](/2022/Days/Images/Day8_Go12.png)
_(Mac 운영체제의 경우 `main` 파일이 생성됩니다. - 옮긴이)_

이를 실행한다면 동일한 결과를 볼 수 있습니다:

```bash
# Windows
$ ./main.exe
Hello #90DaysOfDevOps

# Mac - 올긴이
$ ./main
Hello #90DaysOfDevOps
```

## 자료

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

[Day 9](day09.md)에서 봐요!

![](/2022/Days/Images/Day8_Go13.png)
