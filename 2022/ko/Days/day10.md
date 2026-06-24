---
title: '#90DaysOfDevOps - The Go Workspace - Day 10'
published: false
description: 90DaysOfDevOps - The Go Workspace
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048701
---

### Go 워크스페이스

[Day 8](day08.md)에서 Go 워크스페이스를 간략히 살펴보고 'Hello #90DaysOfDevOps'의 데모를 실행했습니다. 그러나 Go 워크스페이스에 대해 좀 더 자세히 설명해야 합니다.

기본값을 선택한 다음 이미 정의된 GOPATH에 Go 폴더가 생성됐지만, 실제로는 이 GOPATH를 원하는 위치로 변경할 수 있습니다.

다음 명령어를 실행하면

```
echo $GOPATH
```

출력은 다음과 유사해야 합니다.(사용자 아이디가 다를 수 있습니다.)

```
/home/michael/projects/go
```

그런 다음 여기에 3개의 디렉토리를 생성했습니다. **src**, **pkg**, **bin**.

![](/2022/Days/Images/Day10_Go1.png)

**src**는 Go 프로그램과 프로젝트를 저장하는 곳으로, 모든 Go 리포지토리의 네임스페이스 패키지 관리를 담당합니다. 컴퓨터에서 Hello #90DaysOfDevOps 프로젝트를 위한 Hello 폴더를 확인할 수 있습니다.

![](/2022/Days/Images/Day10_Go2.png)

**pkg**는 프로그램에 설치되거나 설치되었던 패키지의 파일을 저장하는 곳입니다. 이것은 사용 중인 패키지가 수정되었는지 여부에 따라 컴파일 프로세스의 속도를 향상하는 데 도움이 됩니다.

![](/2022/Days/Images/Day10_Go3.png)

**bin**은 컴파일된 모든 바이너리가 저장되는 곳입니다.

![](/2022/Days/Images/Day10_Go4.png)

Hello #90DaysOfDevOps는 복잡한 프로그램이 아니기 때문에, 다른 훌륭한 리소스인 [GoChronicles](https://gochronicles.com/)에서 가져온 좀 더 복잡한 Go 프로그램의 예시를 소개해 드리겠습니다.

![](/2022/Days/Images/Day10_Go5.png)

이 페이지는 [GoChronicles](https://gochronicles.com/project-structure/)에 언급되지 않은 다른 폴더에 대한 자세한 설명과 함께, 레이아웃이 왜 이렇게 구성되었는지에 대해 자세히 설명합니다.

### 코드 컴파일 및 실행

[Day 9](day09.md)에서도 코드 컴파일에 대한 간략한 소개를 다루었지만, 이번에는 더 자세히 살펴보겠습니다.

코드를 실행하기 위해서는 먼저 코드를 **컴파일**해야 합니다. Go에서 컴파일하는 방법은 세 가지가 있습니다.

- go build
- go install
- go run

컴파일을 진행하기 전에 Go를 설치하면 무엇을 얻을 수 있는지 살펴봅시다.

8일 차에 Go를 설치하는 동안 Go 소스 파일의 빌드 및 처리를 용이하게 하는 다양한 프로그램으로 구성된 Go 도구를 얻었습니다. 이러한 도구 중 하나가 `Go`입니다.

표준 Go 설치에 포함되지 않은 추가 도구도 설치할 수 있다는 점에 유의하세요.

Go가 제대로 설치되었는지 확인하려면 명령 프롬프트를 열고 `go`를 입력합니다. 아래 이미지와 비슷한 내용이 표시되고 그 뒤에 "Additional Help Topics"가 표시됩니다. 현재로서는 이것에 대해 생각할 필요는 없습니다.

![](/2022/Days/Images/Day10_Go6.png)

8일째인 지금까지 이미 이 도구 중 최소 두 가지를 사용했다는 사실을 기억하실 겁니다.

![](/2022/Days/Images/Day10_Go7.png)

자세히 알고 싶은 것은 "build", "install", 그리고 "run"입니다.

![](/2022/Days/Images/Day10_Go8.png)

- `go run` - 이 명령은 command line에 지정한 .go 파일로 구성된 기본 패키지를 컴파일하고 실행합니다. 이때, 컴파일된 실행 파일은 임시 폴더에 저장됩니다.
- `go build` - 현재 디렉토리에서 패키지와 종속성을 컴파일합니다. 만약 프로젝트에 `main` 패키지가 포함되어 있다면, 실행 파일이 현재 디렉토리에 생성됩니다. 그렇지 않은 경우, 실행 파일은 `pkg` 폴더에 생성되며, 이후 다른 Go 프로그램에서 가져와서 사용할 수 있습니다. 또한 `go build`를 사용하면 Go가 지원하는 모든 OS 플랫폼에 대해 실행 파일을 빌드할 수 있습니다.
- `go install` - go build와 비슷하지만, 실행 파일을 `bin` 폴더에 저장합니다.

`go build`과 `go run`을 실행했지만, 원하는 경우 이곳에서 다시 실행할 수 있습니다. `go install`은 실행 파일을 bin 폴더에 넣는 것으로 설명한 대로 수행됩니다.

![](/2022/Days/Images/Day10_Go9.png)

이 글을 따라오면서 아래의 재생 목록이나 동영상 중 하나를 시청하시기 바랍니다. 제 목표는 여러분과 함께하는 7일 또는 7시간 동안의 여정에서 제가 발견한 흥미로운 정보들을 공유하는 것이기에 저는 이 모든 자료에서의 핵심 정보를 모아서 Go 언어의 기본 개념을 설명하고자 합니다. 아래 자료들은 여러분이 이해해야 할 필수적인 주제에 대한 이해에 도움을 줄 수 있습니다.

## 자료

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

[Day 11](day11.md)에서 봐요!
