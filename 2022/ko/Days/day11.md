---
title: '#90DaysOfDevOps - Variables & Constants in Go - Day 11'
published: false
description: 90DaysOfDevOps - Variables & Constants in Go
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048862
---

오늘의 주제에 들어가기 전에, [Techworld with Nana](https://www.youtube.com/watch?v=yyUHQIec83I)에서 다룬 Go의 기초에 대한 환상적이고 간결한 여정에 대해 큰 박수를 보내고 싶습니다.

[Day 8](day08.md)에는 환경을 설정하였고, [Day 9](day09.md)에는 Hello #90DaysOfDevOps 코드를 검토하였으며, [Day 10](day10.md)에는 Go 워크스페이스를 살펴보고 코드를 컴파일한 후 실행하는 과정까지 자세히 학습했습니다.

오늘은 새로운 프로그램을 만들어보면서 변수, 상수 그리고 데이터 타입에 대해 살펴볼 예정입니다.

## Go의 변수 및 상수

우선, 애플리케이션을 계획하는 것으로 시작하겠습니다. #90DaysOfDevOps 챌린지가 얼마나 남았는지 알려주는 프로그램을 만드는 것이 좋을 것 같습니다.

프로그램을 실행하기 전에 고려해야 할 사항은 앱을 빌드하고 참가자를 환영하며 완료한 일 수에 대한 피드백을 줄 때 "#90DaysOfDevOps"이란 용어를 여러 번 사용할 수 있다는 것입니다. 이것은 프로그램에서 #90DaysOfDevOps를 변수로 사용할 좋은 상황입니다.

- 변수는 값을 저장하기 위해 사용됩니다.
- 저장된 정보나 값이 담긴 작은 상자와 같은 것입니다.
- 변수는 프로그램 전체에서 사용할 수 있으며, 챌린지나 변수가 변경되더라도 한 곳에서만 변경하면 되는 장점이 있습니다. 즉, 변수 하나만 변경하면 커뮤니티의 다른 챌린지에도 적용할 수 있습니다.

Go 프로그램에서는 변수를 선언하고 값을 정의하기 위해 **키워드(var, const...)**를 사용합니다. 이러한 변수 선언은 `func main` 코드 블록 내에서 이루어집니다. `키워드`에 대한 자세한 내용은 [여기](https://go.dev/ref/spec#Keywords)에서 확인할 수 있습니다.

변수 이름이 명시적인지 항상 확인하세요. 변수를 선언하면 반드시 사용해야 하며, 그렇지 않으면 오류가 발생합니다. 이는 사용되지 않는 코드(죽은 코드)를 방지할 수 있습니다. 사용되지 않는 패키지도 마찬가지입니다.

```go
var challenge = "#90DaysOfDevOps"
```

다음 코드 스니펫에서 볼 수 있듯이 위와 같이 선언하면 변수를 사용할 수 있음을 알 수 있습니다.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    fmt.Println("Welcome to", challenge, "")
}
```

위 코드 스니펫은 [day11_example1.go](/2022/Days/Go/day11_example1.go)에서 볼 수 있습니다.

위 예제를 사용하여 코드를 빌드하면 아래와 같이 출력되는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day11_Go1.png)

이번 챌린지의 기간은 90일이지만, 다음 챌린지에서는 100일이 될 수도 있으므로, 이번 챌린지에서 사용되는 변수가 다음 챌린지에서도 유용하게 사용될 수 있도록 정의하고 싶습니다. 하지만 우리 프로그램에서는 이 변수를 상수로 정의하고자 합니다. 상수는 값을 변경할 수 없지만, 변수와 유사합니다.(이 코드를 사용하여 새로운 앱을 만들 경우 상수를 변경할 수 있지만, 애플리케이션을 실행하는 동안에는 이 90이 변경되지 않습니다.)

`const` 키워드를 코드에 추가하고, 출력을 위해 다른 코드 라인을 추가합니다.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge")
}
```

위 코드 스니펫은 [day11_example2.go](/2022/Days/Go/day11_example2.go)에서 볼 수 있습니다.

`go build`를 다시 실행하면 아래와 같은 결과를 확인할 수 있습니다.

![](/2022/Days/Images/Day11_Go2.png)

이것이 우리 프로그램의 끝은 아니며 [Day 12](day12.md)에서 더 많은 기능을 추가할 예정입니다. 현재 완료한 일 수 외에도 다른 변수를 추가하려고 합니다.

`dayscomplete` 변수를 추가하였고, 완료한 일 수도 함께 기록하였습니다.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90
    var dayscomplete = 11

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge and you have completed", dayscomplete, "days")
    fmt.Println("Great work")
}
```

위 코드 스니펫은 [day11_example3.go](/2022/Days/Go/day11_example3.go)에서 볼 수 있습니다.

`go build`를 다시 실행하거나 `go run`을 실행할 수 있습니다.

![](/2022/Days/Images/Day11_Go3.png)

코드를 더 쉽게 읽고 편집하기 위해 몇 가지 다른 예제입니다. 이전에는 `Println`을 사용했지만, `Printf`를 사용하여 코드 줄 끝에 변수를 순서대로 정의하는 `%v`를 사용하면 코드를 간단히 만들 수 있습니다. 또한 줄 바꿈은 `\n`을 사용합니다.

기본값을 사용하기 때문에 `%v`를 사용하고 있지만, 다른 옵션은 [fmt package documentation](https://pkg.go.dev/fmt)의 코드 예제 [day11_example4.go](/2022/Days/Go/day11_example4.go)에서 확인할 수 있습니다.

변수를 정의하는 방법 중에는 '변수'와 '타입'을 명시하는 대신 더 간단한 형태로 정의하는 것도 있습니다. 이렇게 코딩하면 코드의 기능은 동일하지만, 더 깔끔하고 간결해집니다. 하지만 이 방법은 변수에만 적용되며 상수에는 적용되지 않습니다.

```go
func main() {
    challenge := "#90DaysOfDevOps"
    const daystotal = 90
```

## 데이터 타입

예제에서는 변수의 타입을 정의하지 않았습니다. 이는 값을 지정하고 Go가 해당 타입을 유추할 수 있을 만큼 똑똑하거나 저장한 값에 기반하여 타입을 결정할 수 있기 때문입니다. 그러나 사용자가 입력하도록 하려면 특정 타입을 명시해야 합니다.

지금까지 코드에서는 문자열과 정수를 사용했습니다. 정수는 일 수를, 문자열은 챌린지 이름을 나타내는 데 사용했습니다.

각 데이터 타입은 서로 다른 작업을 수행하며, 이에 따라 다르게 작동한다는 것을 유의해야 합니다. 예를 들어, 정수는 문자열과는 곱할 수 없습니다.

총 네 가지 타입이 있습니다:

- **Basic type**: 숫자, 문자열, 불리언 값이 해당 범주에 속합니다.
- **Aggregate type**: 배열과 구조체가 해당 범주에 속합니다.
- **Reference type**: 포인터, 슬라이스, 맵, 함수, 채널이 해당 범주에 속합니다.
- **Interface type**

데이터 타입은 프로그래밍에서 매우 중요한 개념입니다. 이는 변수의 크기와 타입을 지정하는 데 사용됩니다.

Go는 정적으로 타입이 지정되어 있어서 변수의 타입이 정의된 후에는 해당 타입의 데이터만 저장할 수 있습니다.

Go 언어에는 다음과 같이 세 가지 기본 데이터 타입이 있습니다:

- **bool**: 참(True) 또는 거짓(False) 값을 나타냅니다.
- **Numeric**: 정수형(integer), 부동 소수점 값(floating-point), 복소수형(complex)을 나타냅니다.
- **string**: 문자열 값을 나타냅니다.

[Golang by example](https://golangbyexample.com/all-data-types-in-golang-with-examples/)에서 데이터 타입에 대해 매우 자세하게 설명되어 있습니다.

[Techworld with Nana](https://www.youtube.com/watch?v=yyUHQIec83I&t=2023s)에서도 데이터 타입에 대해 자세히 다루고 있으니 추천해드립니다.

변수의 타입을 정의해야 하는 경우 다음과 같이 정의할 수 있습니다:

```go
var TwitterHandle string
var DaysCompleted uint
```

Go 언어에서 변수는 값이 할당되는 것을 의미합니다. 따라서 해당 값을 출력하려면 다음과 같이 할 수 있습니다:

```go
fmt.Printf("challenge is %T, daystotal is %T, dayscomplete is %T\n", conference, daystotal, dayscomplete)
```

여러 종류의 정수 및 부동 소수점 타입이 있으며, 앞의 링크들에서 자세한 내용을 확인할 수 있습니다.

- **int** = 정수
- **unint** = 양의 정수
- **floating point types** = 소수점이 있는 실수

## 자료

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

이제부터는 프로그램에 몇 가지 사용자 입력 기능을 추가해서 완료된 일수를 입력받도록 하겠습니다.

[Day 12](day12.md)에서 봐요!
