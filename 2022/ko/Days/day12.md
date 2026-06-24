---
title: '#90DaysOfDevOps - Getting user input with Pointers and a finished program - Day 12'
published: false
description: 90DaysOfDevOps - Getting user input with Pointers and a finished program
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048864
---

## 포인터를 이용하여 사용자 입력을 받아 프로그램 완성하기

어제([Day 11](day11.md)) 독립적인 첫 번째 Go 프로그램을 만들 때, 코드 내에서 사용자 입력을 받기 위한 변수를 생성하고 값을 할당했습니다. 이제 사용자에게 입력을 요청하여 최종 메시지 값을 변수에 할당하려고 합니다.

## 사용자 입력 받기

사용자 입력을 받기 전에 애플리케이션을 다시 살펴보고 변수를 점검해보겠습니다.

어제 우리는 [day11_example4.go](/2022/Days/Go/day11_example4.go) 코드에서 `challenge, daystotal, dayscomplete`를 변수와 상수로 정의했습니다.

이제 새로운 변수 `TwitterName`을 추가하겠습니다. 이 변수는 [day12_example1.go](/2022/Days/Go/day12_example1.go) 코드에서 찾을 수 있습니다. 코드를 실행하면 아래와 같이 출력됩니다.

![](/2022/Days/Images/Day12_Go1.png)

현재 12일째인데, 코드가 하드코딩되어 있다면 매일 `dayscomplete`를 수정하고 매일 코드를 컴파일해야 하기 때문에 효율적이지 않습니다.

사용자로부터 이름과 수행한 일수를 입력받아야 합니다. 이를 위해서는 `fmt` 패키지 내의 다른 함수를 사용할 수 있습니다.

`fmt` 패키지는 포맷된 입력 및 출력(I/O)을 위한 다양한 기능을 제공합니다. 이 패키지에 대한 요약은 다음과 같습니다.

- 메시지 출력
- 사용자 입력 수집
- 파일에 쓰기

이 방법은 변수에 값을 할당하는 대신 사용자로부터 입력을 요청하는 방식입니다.

```go
fmt.Scan(&TwitterName)
```

변수 앞에 `&`를 사용하는 것도 주목해 주세요. 이것은 포인터라 불리며, 다음 섹션에서 다룰 예정입니다.

[day12_example2.go](/2022/Days/Go/day12_example2.go) 코드에서는 사용자로부터 `TwitterName`과 `DaysCompleted`라는 두 변수를 입력받고 있습니다.

프로그램을 실행하면 위의 두 변수 모두에 대한 입력을 받는 걸 볼 수 있습니다.

![](/2022/Days/Images/Day12_Go2.png)

사용자 의견을 반영하여 메시지를 출력하는 기능도 좋지만, 챌린지 종료까지 남은 일수를 알려주는 기능도 추가하는 것이 어떨까요?

이를 위해 `remainingDays`라는 변수를 만들고, `90`이라는 값을 할당하겠습니다. 그리고 `DaysCompleted`로 사용자 입력을 받으면 남은 날짜를 계산하여 `remainingDays`의 값을 변경하겠습니다. 다음과 같이 간단하게 변수를 변경하면 됩니다.

```go
remainingDays = remainingDays - DaysCompleted
```

[day12_example2.go](/2022/Days/Go/day12_example3.go) 코드에서 완성된 프로그램의 모습을 확인할 수 있습니다.

프로그램을 실행하면, 사용자 입력을 기반으로 `remainingDays` 값을 계산하는 것을 볼 수 있습니다.

![](/2022/Days/Images/Day12_Go3.png)

## 포인터가 무엇인가요?(특수 변수)

포인터는 다른 변수의 메모리 주소를 가리키는 (특수)변수입니다.

이에 대한 자세한 설명은 [geeksforgeeks](https://www.geeksforgeeks.org/pointers-in-golang/)에서 확인할 수 있습니다.

출력 명령에 `&`를 사용할 때와 사용하지 않은 코드 예시를 통해 포인터의 메모리 주소를 확인해 볼 수 있습니다. [day12_example4.go](/2022/Days/Go/day12_example4.go)에 코드 예제를 추가했습니다.

아래는 이 코드를 실행한 모습입니다.

![](/2022/Days/Images/Day12_Go4.png)

## 자료

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

[Day 13](day13.md)에서 봐요!
