---
title: "#90DaysOfDevOps - Let's explain the Hello World code - Day 9"
published: false
description: 90DaysOfDevOps - Let's explain the Hello World code
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1099682
---

## Hello World 코드를 설명해 보겠습니다.

### Go 작동 방식

[Day 8](day08.md)에서는 컴퓨터에 Go를 설치하는 방법을 안내하고, 그 후 첫 번째 Go 애플리케이션을 만들었습니다.

이 섹션에서는 코드를 더 자세히 살펴보고 Go 언어를 더 잘 이해해보겠습니다.

### 컴파일이 무엇인가요?

[Hello World 코드](/2022/Days/Go/hello.go)를 이해하기 전에 컴파일에 대한 이해가 필요합니다.

파이썬, 자바, Go, C++ 등은 우리가 일반적으로 사용하는 고수준 프로그래밍 언어입니다. 이는 사람이 이해하기 쉽지만, 기계가 실행하기 위해서는 기계가 이해할 수 있는 형태로 변환되어야 합니다. 이러한 변환 작업을 컴파일이라고 합니다. 컴파일은 사람이 작성한 코드를 기계가 이해할 수 있는 코드로 변환하는 과정입니다.

![](/2022/Days/Images/Day9_Go1.png)

이전에 [Day 8](day08.md)에서 수행한 작업을 위 그림으로 확인할 수 있습니다. 우선 간단한 Hello World를 출력하는 main.go를 작성하고, 이를 `go build main.go` 명령을 사용하여 실행 파일로 컴파일했습니다.

### 패키지가 무엇인가요?

패키지는 같은 디렉토리에 있는 여러 개의 .go 확장자 파일의 모음으로 구성됩니다. 더 간단히 말하면, 패키지는 같은 디렉토리 내에서 함께 컴파일되는 소스 파일들입니다. 또한, 더 복잡한 Go 프로그램에서는 여러 패키지가 폴더1, 폴더2, 폴더3 등에 분산되어 있을 수 있으며, 각 패키지는 서로 다른 .go 파일들로 이루어질 수 있습니다.

패키지를 사용하면 다른 사람들의 코드를 재사용하여 처음부터 모든 것을 새로 작성할 필요가 없어집니다. 만약 프로그램에서 계산기가 필요하다면, 기존에 Go 언어로 작성된 패키지에서 수학 함수를 가져와 코드에 적용시키면 장기적으로 많은 시간과 노력을 절약할 수 있습니다.

Go 언어는 코드를 패키지로 구성하여 소스 코드의 재사용성과 유지 보수성을 쉽게 확보할 수 있도록 권장합니다.

### Hello #90DaysOfDevOps 한 줄 한 줄

이제 Hello #90DaysOfDevOps main.go 파일을 살펴보겠습니다.

![](/2022/Days/Images/Day9_Go2.png)

Go 언어에서는 모든 .go 파일은 패키지에 속해야 하며, `package something`과 같이 첫 줄에 패키지 이름을 명시해야 합니다. 이때, `package main`은 이 파일이 main 함수를 포함하는 패키지임을 나타냅니다.

패키지 이름은 원하는 대로 지정할 수 있습니다. 하지만 이 프로그램 시작점에서의 패키지는 `main`으로 지정해야 하며, 이는 규칙입니다. (이 규칙에 대해 추가적인 이해가 필요한가요?)

![](/2022/Days/Images/Day9_Go3.png)

코드를 컴파일하고 실행할 때, 시작해야 하는 위치를 컴퓨터에게 알려주어야 합니다. 이를 위해 'main'이라는 함수를 작성합니다. 컴퓨터는 프로그램의 시작점을 찾기 위해 'main' 함수를 찾습니다.

함수는 특정 작업을 수행하는 코드 블록으로, 프로그램 전체에서 사용할 수 있습니다.

`func`를 사용하여 함수를 어떤 이름으로든 선언할 수 있지만, 이 경우 코드가 시작되는 곳이므로 `main`이라는 이름을 지정해야 합니다.

![](/2022/Days/Images/Day9_Go4.png)

다음으로, 코드의 세 번째 줄인 'import'를 살펴보겠습니다. 이 줄은 메인 프로그램에 다른 패키지를 가져오고자 하는 것을 의미합니다. 'fmt'는 Go에서 제공하는 표준 패키지 중 하나이며, `Println()` 함수를 포함하고 있기 때문에, 이를 'import' 했으므로 여섯 번째 줄에서 이 함수를 사용할 수 있습니다. 프로그램에서 사용하거나 재사용할 수 있는 여러 표준 패키지가 있으므로, 처음부터 작성해야 하는 번거로움을 줄일 수 있습니다. [Go Standard Library](https://pkg.go.dev/std)

![](/2022/Days/Images/Day9_Go5.png)

여기에 있는 `Println()`은 실행 파일이 성공적으로 실행됐을 때 터미널에 STDOUT(standard output)을 출력하는 함수입니다. 괄호 안에는 원하는 메시지를 자유롭게 입력할 수 있습니다.

![](/2022/Days/Images/Day9_Go6.png)

### TLDR

- **1행** = 이 파일은 `main` 패키지에 있으며, 프로그램의 진입점을 포함하고 있으므로 `main`으로 불러야 합니다.
- **3행** = `Println()`을 사용하려면 6행에서 'fmt' 패키지를 'import'하여야 합니다.
- **5행** = 프로그램의 실제 시작점은 `main` 함수입니다.
- **6행** = 이렇게 하면 시스템에 "Hello #90DaysOfDevOps"을 출력할 수 있습니다.

## 자료

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

[Day 10](day10.md)에서 봐요!
