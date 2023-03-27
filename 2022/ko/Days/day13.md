---
title: '#90DaysOfDevOps - Tweet your progress with our new App - Day 13'
published: false
description: 90DaysOfDevOps - Tweet your progress with our new App
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048865
---

## 새로운 앱으로 진행 상황을 트윗하세요

프로그래밍 언어를 살펴본 마지막 날입니다. 우리는 프로그래밍 언어의 겉 부분만 살짝 살펴봤을 뿐이고 앞으로는 더 큰 흥미와 관심을 가지고 더욱 깊이 파고들어야 합니다.

최근 며칠간 애플리케이션에 작은 아이디어를 추가하면서 기능을 개선했습니다. 이번 세션에서는 앞서 언급한 패키지를 활용하여 화면에 진행 상황을 업데이트하는 것뿐만 아니라, 챌린지의 세부 정보와 상태를 트윗할 수 있는 기능을 만들어보려고 합니다.

## 진행 상황을 트윗하는 기능 추가

이 기능을 사용하려면 먼저 트위터에서 개발자 API 접근을 설정해야 합니다.

[Twitter Developer Platform](https://developer.twitter.com)으로 이동하여 내 트위터로 계정으로 로그인하면, 이미 만든 앱이 없는 경우 아래와 같은 화면이 표시됩니다.

![](/2022/Days/Images/Day13_Go1.png)

여기에서 Elevated 등급 계정을 요청할 수 있습니다. 시간이 조금 걸릴 수 있지만, 제 경우에는 빠르게 승인이 됐습니다.

다음으로, 프로젝트 및 앱을 선택하여 앱을 생성합니다. 보유한 계정 액세스 권한에 따라 제한이 있으며, Essential 계정은 하나의 앱과 하나의 프로젝트만, Elevated 계정은 3개의 앱만 만들 수 있습니다.

![](/2022/Days/Images/Day13_Go2.png)

애플리케이션의 이름을 지정합니다.

![](/2022/Days/Images/Day13_Go3.png)

API token이 제공됩니다. 이를 안전한 장소에 저장해야 합니다.(저는 이후 앱을 삭제했습니다.) 나중에 Go 애플리케이션에서 이 토큰이 필요할 것입니다.

![](/2022/Days/Images/Day13_Go4.png)

이제 앱이 생성되었습니다.(스크린샷에 보이는 앱 이름은 이미 생성된 것이기 때문에, 고유해야 하므로 앱 이름을 변경해야 했습니다.)

![](/2022/Days/Images/Day13_Go5.png)

"consumer keys"는 이전에 생성한 key를 의미하며, access token과 비밀번호도 필요합니다. 이 정보는 "Keys & Tokens" 탭에서 확인할 수 있습니다.

![](/2022/Days/Images/Day13_Go6.png)

트위터 개발자 포털에서 필요한 모든 작업을 마쳤습니다. 나중에 필요하실 수 있으니 key를 안전한 곳에 보관해주세요.

## Go 트위터 봇

애플리케이션을 시작했던 코드인 [day13_example1](/2022/Days/Go/day13_example1.go)를 기억해주세요. 하지만 먼저 올바른 코드로 트윗을 생성할 수 있는지 확인해야 합니다.

트위터에 메시지나 출력을 트윗 형태로 전달하기 위한 코드를 생각해봐야 합니다. 이를 위해 [go-twitter](https://github.com/dghubble/go-twitter) 라이브러리를 사용할 것입니다. 이 라이브러리는 Go 언어로 작성된 트위터 API 클라이언트 라이브러리입니다.

메인 애플리케이션에 적용하기 전에 `src` 폴더에 'go-twitter-bot'이라는 새 디렉토리를 만들고, 해당 폴더에서 `go mod init github.com/michaelcade/go-Twitter-bot` 명령을 실행하여 `go.mod` 파일을 생성한 후, 새로운 'main.go' 파일을 작성하여 테스트해 보았습니다.

트위터 개발자 포털에서 생성한 key, token 및 비밀번호가 필요하며, 이를 환경 변수로 설정해야 합니다. 하지만 이는 실행 중인 운영 체제에 따라 다를 수 있습니다.

환경 변수와 관련하여 몇 가지 질문이 있어, 블로그 게시물을 소개해드립니다. 해당 게시물은 환경 변수 설정 방법을 더 자세히 설명하고 있습니다. [How To Set Environment Variables](https://www.twilio.com/blog/2017/01/how-to-set-environment-variables.html)를 통해 확인해보세요.

Windows

```
set CONSUMER_KEY
set CONSUMER_SECRET
set ACCESS_TOKEN
set ACCESS_TOKEN_SECRET
```

Linux / MacOS

```
export CONSUMER_KEY
export CONSUMER_SECRET
export ACCESS_TOKEN
export ACCESS_TOKEN_SECRET
```

이 단계에서는 [day13_example2](/2022/Days/Go/day13_example2.go) 코드를 살펴볼 수 있습니다. 이 코드에서는 구조체를 사용하여 key, secret, token을 정의합니다.

Credentials를 분석하고 트위터 API에 연결하는 `func`가 있습니다.

이후 성공 여부에 따라 트윗을 전송합니다.

```go
package main

import (
    // ...
    "fmt"
    "log"
    "os"

    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

// Credentials는 트위터 REST API에 대한 인증에 필요한
// 모든 access/consumer token과 secret key를 저장합니다.
type Credentials struct {
    ConsumerKey       string
    ConsumerSecret    string
    AccessToken       string
    AccessTokenSecret string
}

// getClient는 인증에 필요한 모든 것을 포함하며,
// 트위터 클라이언트 또는 오류에 대한 포인터를 반환하는
// Credentials 구조체에 대한 포인터를 받아 나중에 트윗을 보내거나
// 새 트윗을 스트리밍하는 데 사용할 수 있는 헬퍼 함수입니다.
func getClient(creds *Credentials) (*twitter.Client, error) {
    // consumer key(API key)와 consumer secret(API secret)을 전달합니다.
    config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
    // access token과 access token secret을 전달합니다.
    token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)

    // Credentials 확인
    verifyParams := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }

    // 사용자를 검색하고 Credentials가 올바른지 확인할 수 있고,
    // 성공적으로 로그인할 수 있는지 확인할 수 있습니다!
    user, _, err := client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        return nil, err
    }

    log.Printf("User's ACCOUNT:\n%+v\n", user)
    return client, nil
}
func main() {
    fmt.Println("Go-Twitter Bot v0.01")
    creds := Credentials{
        AccessToken:       os.Getenv("ACCESS_TOKEN"),
        AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
        ConsumerKey:       os.Getenv("CONSUMER_KEY"),
        ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
    }

    client, err := getClient(&creds)
    if err != nil {
        log.Println("Error getting Twitter Client")
        log.Println(err)
    }

    tweet, resp, err := client.Statuses.Update("A Test Tweet from the future, testing a #90DaysOfDevOps Program that tweets, tweet tweet", nil)
    if err != nil {
        log.Println(err)
    }
    log.Printf("%+v\n", resp)
    log.Printf("%+v\n", tweet)
}

```

위와 같이 작성하면 상황에 따라 오류가 발생하거나 성공하여 코드에 적힌 메시지가 포함된 트윗이 전송됩니다.

## 두 가지의 결합 - Go-Twitter-Bot + 우리의 앱

이제 이 두 가지를 `main.go`에서 병합해야 합니다. 이 작업을 수행하는 더 좋은 방법이 있을 것이며, 프로젝트에 둘 이상의 `.go` 파일을 가질 수 있으므로 이에 대해 의견을 제시해 주시기 바랍니다.

[day13_example3](/2022/Days/Go/day13_example3.go)에서 병합된 코드를 볼 수 있지만 아래에서도 보여드리겠습니다.

```go
package main

import (
    // ...
    "fmt"
    "log"
    "os"

    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

// Credentials는 트위터 REST API에 대한 인증에 필요한
// 모든 access/consumer token과 secret key를 저장합니다.
Credentials REST API에 대한 인증에 필요한 모든 액세스/소비자 토큰과 비밀 키를 저장합니다.
type Credentials struct {
    ConsumerKey       string
    ConsumerSecret    string
    AccessToken       string
    AccessTokenSecret string
}

// getClient는 인증에 필요한 모든 것을 포함하며,
// 트위터 클라이언트 또는 오류에 대한 포인터를 반환하는
// Credentials 구조체에 대한 포인터를 받아 나중에 트윗을 보내거나
// 새 트윗을 스트리밍하는 데 사용할 수 있는 헬퍼 함수입니다.
func getClient(creds *Credentials) (*twitter.Client, error) {
    // consumer key(API key)와 consumer secret(API secret)을 전달합니다.
    config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
    // access token과 access token secret을 전달합니다.
    token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)

    // Verify Credentials
    verifyParams := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }

    // we can retrieve the user and verify if the credentials
    // we have used successfully allow us to log in!
    user, _, err := client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        return nil, err
    }

    log.Printf("User's ACCOUNT:\n%+v\n", user)
    return client, nil
}

func main() {
    creds := Credentials{
        AccessToken:       os.Getenv("ACCESS_TOKEN"),
        AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
        ConsumerKey:       os.Getenv("CONSUMER_KEY"),
        ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
    }
    {
        const DaysTotal int = 90
        var remainingDays uint = 90
        challenge := "#90DaysOfDevOps"

        fmt.Printf("Welcome to the %v challenge.\nThis challenge consists of %v days\n", challenge, DaysTotal)

        var TwitterName string
        var DaysCompleted uint

        // asking for user input
        fmt.Println("Enter Your Twitter Handle: ")
        fmt.Scanln(&TwitterName)

        fmt.Println("How many days have you completed?: ")
        fmt.Scanln(&DaysCompleted)

        // Credentials 확인
        remainingDays = remainingDays - DaysCompleted

        //fmt.Printf("Thank you %v for taking part and completing %v days.\n", TwitterName, DaysCompleted)
        //fmt.Printf("You have %v days remaining for the %v challenge\n", remainingDays, challenge)
        //fmt.Println("Good luck")

        client, err := getClient(&creds)
        if err != nil {
            log.Println("Error getting Twitter Client, this is expected if you did not supply your Twitter API tokens")
            log.Println(err)
        }

        message := fmt.Sprintf("Hey I am %v I have been doing the %v for %v days and I have %v Days left", TwitterName, challenge, DaysCompleted, remainingDays)
        tweet, resp, err := client.Statuses.Update(message, nil)
        if err != nil {
            log.Println(err)
        }
        log.Printf("%+v\n", resp)
        log.Printf("%+v\n", tweet)
    }
}
```

결과는 트윗으로 표시되어야 하지만, 환경 변수가 제공되지 않은 경우 아래와 같은 오류가 발생해야 합니다.

![](/2022/Days/Images/Day13_Go7.png)

만약 이 문제를 해결하거나 트위터 인증을 사용하지 않기로 선택했다면, 어제 작성한 코드를 사용할 수 있습니다. 성공한 경우 터미널 출력은 다음과 유사하게 표시됩니다:

![](/2022/Days/Images/Day13_Go8.png)

결과 트윗은 아래와 같이 표시되어야 합니다:

![](/2022/Days/Images/Day13_Go9.png)

## 여러 OS에 맞게 컴파일하는 방법

"여러 운영체제에서 컴파일하려면 어떻게 해야 할까요?"라는 질문에 대해 다루고자 합니다. Go의 가장 장점 중 하나는 다양한 운영체제에 대해 쉽게 컴파일할 수 있다는 점입니다. 아래 명령어를 실행하면 모든 운영체제 목록을 확인할 수 있습니다:

```
go tool dist list
```

지금까지 `go build` 명령을 사용하여 HM(host machine)과 빌드 대상을 환경 변수 `GOOS`와 `GOARCH`를 이용하여 결정할 수 있었습니다. 그러나 아래 예제와 같이 다른 바이너리를 생성할 수도 있습니다.

```
GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}_0.1_darwin main.go
GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}_0.1_linux main.go
GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}_0.1_windows main.go
GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}_0.1_linux_arm64 main.go
GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}_0.1_darwin_arm64 main.go
```

위의 모든 플랫폼에 대한 바이너리가 디렉토리에 생성됩니다. 이후 코드에 새로운 기능을 추가할 때마다, 바이너리로 빌드하기 위해 [makefile](/2022/Days/Go/makefile)을 사용할 수도 있습니다.

지금 리포지토리에서 볼 수 있는 릴리스를 만드는 데 사용한 [repository](https://github.com/MichaelCade/90DaysOfDevOps/releases)입니다.

## 자료

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)
- [A great repo full of all things DevOps & exercises](https://github.com/bregman-arie/devops-exercises)
- [GoByExample - Example based learning](https://gobyexample.com/)
- [go.dev/tour/list](https://go.dev/tour/list)
- [go.dev/learn](https://go.dev/learn/)

7일간의 프로그래밍 언어 학습을 마무리합니다. 앞으로도 더 많은 내용을 다룰 예정이며, 이번 학습을 통해 Go 프로그래밍 언어의 다른 측면도 이해할 수 있었기를 바랍니다.

다음으로, Linux와 Linux에서 알아야 할 몇 가지 기본 사항에 대해 살펴보겠습니다.

[Day 14](day14.md)에서 봐요!
