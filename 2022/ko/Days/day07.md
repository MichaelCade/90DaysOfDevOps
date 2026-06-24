---
title: '#90DaysOfDevOps - The Big Picture: Learning a Programming Language - Day 7'
published: false
description: 90DaysOfDevOps - The Big Picture DevOps & Learning a Programming Language
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048856
---

## 큰 그림: 데브옵스 및 프로그래밍 언어 학습

장기적으로 데브옵스 엔지니어로 성공하려면, 적어도 하나의 프로그래밍 언어를 기초 수준에서 이해하는 것이 매우 중요합니다. 이번 섹션에서는 이러한 중요성에 대해 자세히 알아보고, 이번 주나 섹션이 끝날 때까지 학습 계획을 수립하기 위해 왜, 어떻게, 무엇을 해야 하는지 더 잘 이해하도록 합니다.

소셜 미디어에서 데브옵스 관련 직무를 수행하기 위해 프로그래밍 기술이 필요한지 묻는다면 대부분 '그렇다'고 대답할 것입니다. 하지만 이러한 역할을 위해 어떤 프로그래밍 언어를 배워야 하는지에 대한 질문은 대답하기 어렵습니다. 일반적으로 Python을 추천하지만, Go(Golang)를 추천하는 경우도 점점 늘어나고 있습니다.

제 생각에는 데브옵스에서 성공하려면 프로그래밍 기술에 대한 탄탄한 기초를 갖추는 것이 중요합니다. 그러나 이러한 기술이 왜 필요한지 이해해야 목표를 달성하기 위한 적절한 길을 선택할 수 있습니다.

## 프로그래밍 언어를 배워야 하는 이유를 이해합니다.

데브옵스에서 사용되는 대부분의 도구가 Python 또는 Go로 작성되어 있기 때문에 데브옵스 엔지니어에게 Python과 Go를 자주 권장합니다. 이러한 도구를 만들거나 이러한 도구를 만드는 팀에 합류하는 데 관심이 있다면 데브옵스 도구를 만드는 데 사용되는 언어를 배우는 것이 좋습니다. Kubernetes나 컨테이너에 많이 관여하고 있다면 프로그래밍 언어로 Go를 배우는 것이 더 유리할 수 있습니다. 제 경우에는 클라우드 네이티브 에코시스템에 속해 있으며 Kubernetes의 데이터 관리에 중점을 두는 회사로, 모든 도구가 Go로 작성되어 있는 Kasten by Veeam에서 일하고 있습니다.

아직 명확한 결정이 내려지지 않은 학생이나 경력 전환기의 경우, 진로를 선택하는 것이 쉽지 않을 수 있습니다. 이런 상황에서는 자신이 만들고 싶은 애플리케이션과 가장 잘 맞고 가장 공감할 수 있는 것을 선택하는 것이 가장 좋습니다.

저는 소프트웨어 개발자가 되고자 하는 것이 아닙니다. 프로그래밍 언어에 대한 기본적인 이해를 바탕으로 다양한 도구의 기능을 이해하고 이를 개선할 수 있는 방법을 찾는 것이 목표입니다.

Kasten K10이나 Terraform, HCL과 같은 데브옵스 도구와 상호 작용하는 방법을 이해하는 것이 중요합니다. 이는 일반적으로 YAML로 작성되는 구성 파일을 통해 이루어집니다. 이러한 파일을 통해 DevOps 도구로 작업을 수행할 수 있습니다. (이 섹션의 마지막 날에는 YAML에 대해 더 자세히 살펴보겠습니다.)

## 제가 방금 프로그래밍 언어를 배우지 말라고 설득한 건가요?

많은 경우, 데브옵스 엔지니어의 주된 책임은 엔지니어링 팀과 협력하여 데브옵스 관행을 workflow에 통합하는 것입니다. 여기에는 애플리케이션에 대한 광범위한 테스트와 workflow가 앞서 설명한 데브옵스 원칙을 준수하는지 확인하는 작업인 경우가 많습니다. 그러나 작업의 상당 부분에는 애플리케이션 성능 또는 기타 기술적 결함과 관련된 문제 해결도 포함될 수 있습니다. 따라서 작업 중인 애플리케이션에서 사용하는 프로그래밍 언어에 대한 전문 지식을 갖추는 것이 중요합니다. 예를 들어, 애플리케이션이 Node.js로 작성된 경우 Go나 Python을 알아도 별 도움이 되지 않습니다.

## 왜 Go인가요?

Go가 데브옵스를 위한 유망한 프로그래밍 언어인 이유. Go는 최근 몇 년 동안 상당한 인기를 얻었으며, 2021년 StackOverflow 설문조사에서 가장 선호하는 프로그래밍, 스크립팅 및 마크업 언어 중 하나로 4위에 올랐고, 1위는 Python이었습니다. 자세히 살펴보겠습니다. [StackOverflow 2021 Developer Survey – Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)

앞서 언급했듯이, Kubernetes, Docker, Grafana, Prometheus 등 가장 잘 알려진 데브옵스 도구 및 플랫폼 중 일부는 Go로 작성되었습니다.

데브옵스 목적에 특히 적합한 Go의 주요 기능이나 특성은 무엇일까요?

## Go 프로그램 구축 및 배포

데브옵스 역할에서 Python과 같은 언어를 사용하면 프로그램을 실행하기 전에 컴파일할 필요가 없다는 이점이 있습니다. 빌드 프로세스가 길어지면 불필요하게 프로세스 속도가 느려질 수 있으므로 소규모 자동화 작업에 특히 유용합니다. Go는 컴파일 언어이긴 하지만 **머신 코드로 직접 컴파일되며** 컴파일 시간이 빠른 것으로 알려져 있다는 점에 집중할 필요가 있습니다.

## 데브옵스를 위한 Go vs Python

Go 프로그램은 정적으로 링크됩니다. 즉, Go 프로그램을 컴파일할 때 모든 것이 단일 바이너리 실행 파일에 포함되며 원격 시스템에 외부 종속성을 설치할 필요가 없습니다. 따라서 외부 라이브러리에 의존하는 경우가 많은 Python 프로그램에 비해 Go 프로그램을 더 쉽게 배포할 수 있습니다. Python을 사용하면 프로그램을 실행하려는 원격 머신에 필요한 모든 라이브러리가 설치되어 있는지 확인해야 합니다.

Go는 특정 플랫폼에 종속되지 않는 언어이기 때문에 Linux, Windows, macOS 등 \*다양한 운영 체제용 바이너리 실행 파일을 쉽게 만들 수 있습니다. 반면에 Python으로 특정 운영 체제용 바이너리 실행 파일을 생성하는 것은 더 어려울 수 있습니다.

Go는 Python에 비해 컴파일 및 실행 시간이 빠르고 CPU와 메모리 측면에서 리소스 소비가 적은 매우 효율적인 프로그래밍 언어입니다. 이 언어에는 다양한 최적화가 구현되어 있어 뛰어난 성능을 제공합니다. (자세한 내용은 아래 리소스를 참조하세요.)

특정 프로그래밍 작업을 위해 종종 서드파티 라이브러리에 의존하는 Python에 비해 Go에는 데브옵스에 필요한 대부분의 기능을 바로 사용할 수 있는 표준 라이브러리가 있습니다. 여기에는 파일 처리, HTTP 웹 서비스, JSON 처리, 동시성 및 병렬 처리를 위한 기본 지원, 심지어 테스트 기능까지 포함됩니다.

제가 Go를 선택한 이유는 파이썬을 비판하기 위해서가 아니라는 점을 분명히 말씀드리고 싶습니다. 제가 근무하는 회사에서 Go로 소프트웨어를 개발하기 때문에 Go를 사용하는 것이 합리적입니다.

프로그래밍 언어 하나를 배우면 다른 언어를 배우기가 더 쉬워진다는 말을 들었습니다. 어느 회사에서든 자바스크립트 및 Node.js 애플리케이션을 관리, 설계, 오케스트레이션, 디버깅하게 될 가능성이 있고, 지금 Go를 배워두면 도움이 될 것입니다.

## 자료

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

앞으로 6일 동안 위에서 언급한 자료를 공부하고 매일 진행 상황을 메모할 계획입니다. 전체 과정은 약 3시간 정도 걸리지만, 저는 이 학습에 하루 한 시간씩 할애할 예정입니다. 시간이 여유가 있으신 분들도 함께 하며 깊이 있게 공부하실 수 있도록 전체 리소스 목록을 공유해 드립니다.

[Day 8](day08.md)에서 봐요!
