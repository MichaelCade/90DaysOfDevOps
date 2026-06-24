---
title: '#90DaysOfDevOps - Application Focused - Day 3'
published: false
description: 90DaysOfDevOps - Application Focused
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048825
---

## 데브옵스 수명 주기 - 애플리케이션 중심

앞으로 몇 주 동안 계속 진행하면서 지속적 개발, 테스트, 배포, 모니터에 대해 100% 반복해서 접하게 될 것입니다. DevOps 엔지니어 역할로 향하고 있다면 반복성이 익숙해질 것이지만 매번 지속적으로 향상시키는 것도 흥미를 유지하는 또 다른 요소입니다.

이번 시간에는 애플리케이션을 처음부터 끝까지 살펴본 다음 다시 반복하듯 되돌아보는 고차원의 시각으로 살펴보겠습니다.

### Development (개발)

애플리케이션의 새로운 예를 들어 보겠습니다. 먼저 아무것도 만들어지지 않은 상태에서 개발자는 고객 또는 최종 사용자와 요구 사항을 논의하고 애플리케이션에 대한 일종의 계획이나 요구 사항을 마련해야 합니다. 그런 다음 요구 사항을 바탕으로 새로운 애플리케이션을 만들어야 합니다.

이 단계의 도구와 관련해서는 애플리케이션을 작성하는 데 사용할 IDE와 프로그래밍 언어를 선택하는 것 외에는 실제 요구 사항이 없습니다.

데브옵스 엔지니어로서 이 계획을 만들거나 최종 사용자를 위해 애플리케이션을 코딩하는 것은 여러분이 아니라 숙련된 개발자가 할 일이라는 점을 기억하세요.

그러나 애플리케이션에 대한 최상의 인프라 결정을 내릴 수 있도록 일부 코드를 읽을 수 있는 것도 나쁘지 않을 것입니다.

앞서 이 애플리케이션은 어떤 언어로든 작성할 수 있다고 언급했습니다. 중요한 것은 버전 관리 시스템을 사용하여 유지 관리해야 한다는 것인데, 이 부분은 나중에 자세히 다룰 것이며 특히 **Git**에 대해 자세히 살펴볼 것입니다.

또한 이 프로젝트에서 한 명의 개발자가 작업하는 것이 아닐 수도 있지만, 이 경우에도 모범 사례에서는 코드를 저장하고 협업하기 위한 코드 저장소가 필요하며, 이는 비공개 또는 공개일 수 있고 호스팅되거나 비공개로 배포될 수 있으며 일반적으로 **GitHub 또는 GitLab**과 같은 코드 저장소가 코드 저장소로 사용되는 것을 듣게 될 것입니다. 이에 대해서는 나중에 **Git** 섹션에서 다시 다루겠습니다.

### Testing (테스팅)

이 단계에서는 요구 사항이 있고 애플리케이션이 개발되고 있습니다. 하지만 우리가 사용할 수 있는 모든 다양한 환경, 특히 선택한 프로그래밍 언어에서 코드를 테스트하고 있는지 확인해야 합니다.

이 단계에서 QA는 버그를 테스트할 수 있으며, 테스트 환경을 시뮬레이션하는 데 컨테이너를 사용하는 경우가 많아져 전반적으로 물리적 또는 클라우드 인프라의 비용 오버헤드를 개선할 수 있습니다.

이 단계는 또한 다음 영역인 지속적 통합의 일부로 자동화될 가능성이 높습니다.

이 테스트를 자동화할 수 있다는 것은 수십, 수백, 심지어 수천 명의 QA 엔지니어가 이 작업을 수동으로 수행해야 하는 것과 비교하면 그 자체로 의미가 있으며, 이러한 엔지니어는 스택 내에서 다른 작업에 집중하여 워터폴 방법론을 사용하는 대부분의 기존 소프트웨어 릴리스에서 지체되는 경향이 있는 버그 및 소프트웨어 테스트 대신 더 빠르게 움직이고 더 많은 기능을 개발할 수 있습니다.

### Integration (통합)

매우 중요한 것은 통합이 데브옵스 라이프사이클의 중간에 있다는 것입니다. 개발자가 소스 코드에 변경 사항을 더 자주 커밋해야 하는 practice(관행)입니다. 이는 매일 또는 매주 단위로 이루어질 수 있습니다.

커밋할 때마다 애플리케이션은 자동화된 테스트 단계를 거치게 되며, 이를 통해 다음 단계로 넘어가기 전에 문제나 버그를 조기에 발견할 수 있습니다.

이제 이 단계에서 "하지만 우리는 애플리케이션을 만들지 않고 소프트웨어 공급업체에서 기성품을 구입합니다."라고 말할 수 있습니다. 많은 회사가 이렇게 하고 있고 앞으로도 계속 그렇게 할 것이며 위의 3단계에 집중하는 것은 소프트웨어 공급업체가 될 것이므로 걱정하지 마세요. 하지만 마지막 단계를 채택하면 기성품 배포를 더 빠르고 효율적으로 배포할 수 있으므로 여전히 채택하고 싶을 수도 있습니다.

오늘 당장 상용 소프트웨어를 구매할 수도 있지만 내일이나 또는... 다음 직장에서 사용할 수도 있기 때문에 위의 지식을 갖추는 것만으로도 매우 중요하다고 말씀드리고 싶습니다.

### Deployment (배포)

Ok so we have our application built and tested against the requirements of our end user and we now need to go ahead and deploy this application into production for our end users to consume.

This is the stage where the code is deployed to the production servers, now this is where things get extremely interesting and it is where the rest of our 86 days dives deeper into these areas. Because different applications require different possibly hardware or configurations. This is where **Application Configuration Management** and **Infrastructure as Code** could play a key part in your DevOps lifecycle. It might be that your application is **Containerised** but also available to run on a virtual machine. This then also leads us onto platforms like **Kubernetes** which would be orchestrating those containers and making sure you have the desired state available to your end users.

Of these bold topics, we will go into more detail over the next few weeks to get a better foundational knowledge of what they are and when to use them.

이제 최종 사용자의 요구 사항에 따라 애플리케이션을 빌드하고 테스트를 마쳤으므로 이제 최종 사용자가 사용할 수 있도록 이 애플리케이션을 프로덕션에 배포해야 합니다.

이 단계는 코드를 프로덕션 서버에 배포하는 단계로, 이제부터 매우 흥미로운 일이 벌어지며 나머지 86일 동안 이러한 영역에 대해 더 자세히 알아볼 것입니다. 애플리케이션마다 필요한 하드웨어나 구성이 다르기 때문입니다. 바로 이 부분에서 **Application Configuration Management(애플리케이션 구성 관리)**와 **Infrastructure as Code(코드형 인프라)**가 데브옵스 라이프사이클에서 중요한 역할을 할 수 있습니다. 애플리케이션이 **Containerised(컨테이너화)**되어 있지만 가상 머신에서 실행할 수 있는 경우도 있을 수 있습니다. 그런 다음 이러한 컨테이너를 오케스트레이션하고 최종 사용자가 원하는 상태를 사용할 수 있도록 하는 **Kubernetes**와 같은 플랫폼으로 이어집니다.

이 대담한 주제 중 앞으로 몇 주에 걸쳐 더 자세히 살펴보면서 컨테이너가 무엇이고 언제 사용하는지에 대한 기초 지식을 쌓을 것입니다.

### Monitoring (관제)

새로운 기능으로 지속적으로 업데이트하고 있는 애플리케이션이 있으며, 테스트 과정에서 문제점이 발견되지 않는지 확인하고 있습니다. 필요한 구성과 성능을 지속적으로 유지할 수 있는 애플리케이션이 우리 환경에서 실행되고 있습니다.

하지만 이제 최종 사용자가 필요한 경험을 얻고 있는지 확인해야 합니다. 이 단계에서는 애플리케이션 성능을 지속적으로 모니터링하여 개발자가 향후 릴리스에서 애플리케이션을 개선하여 최종 사용자에게 더 나은 서비스를 제공할 수 있도록 더 나은 결정을 내릴 수 있도록 해야 합니다.

또한 이 섹션에서는 구현된 기능에 대한 피드백을 수집하고 최종 사용자가 어떻게 개선하기를 원하는지에 대한 피드백을 수집할 것입니다.

안정성은 여기서도 핵심 요소이며, 결국에는 애플리케이션이 필요할 때 항상 사용할 수 있기를 원합니다. 이는 지속적으로 모니터링해야 하는 다른 **observability, security and data management(관찰 가능성, 보안 및 데이터 관리)** 영역으로 이어지며, 피드백을 통해 애플리케이션을 지속적으로 개선, 업데이트 및 릴리스하는 데 항상 사용할 수 있습니다.

특히 [@\_ediri](https://twitter.com/_ediri) 커뮤니티의 일부 의견은 이러한 지속적인 프로세스의 일부로 FinOps 팀도 참여해야 한다고 언급했습니다. 앱과 데이터는 어딘가에서 실행되고 저장되므로 리소스 관점에서 상황이 변경될 경우 비용이 클라우드 요금에 큰 재정적 고통을 주지 않도록 지속적으로 모니터링해야 합니다.

위에서 언급한 "DevOps 엔지니어"에 대해서도 언급할 때가 되었다고 생각합니다. 많은 사람들이 DevOps 엔지니어라는 직책을 가지고 있지만, 이것은 DevOps 프로세스를 포지셔닝하는 이상적인 방법은 아닙니다. 제 말은 커뮤니티의 다른 사람들과 이야기할 때 데브옵스 엔지니어라는 직책이 누구의 목표가 되어서는 안 된다는 것입니다. 실제로 어떤 직책이든 여기에서 설명한 데브옵스 프로세스와 문화를 채택해야 하기 때문입니다. 데브옵스는 클라우드 네이티브 엔지니어/아키텍트, 가상화 관리자, 클라우드 아키텍트/엔지니어, 인프라 관리자와 같은 다양한 직책에서 사용되어야 합니다. 몇 가지 예를 들었지만, 위에서 DevOps 엔지니어를 사용한 이유는 위의 모든 직책에서 사용하는 프로세스의 범위 등을 강조하기 위해서입니다.

## Resources

I am always open to adding additional resources to these readme files as it is here as a learning tool.

My advice is to watch all of the below and hopefully you also picked something up from the text and explanations above.

학습 도구로서 이 Readme 파일에 추가 리소스를 추가하는 것은 언제나 열려 있습니다.

아래의 내용을 모두 보시고 위의 텍스트와 설명에서 많은 것들을 얻으셨으면 좋겠습니다.

- [Continuous Development](https://www.youtube.com/watch?v=UnjwVYAN7Ns) I will also add that this is focused on manufacturing but the lean culture can be closely followed with DevOps.
- [Continuous Testing - IBM YouTube](https://www.youtube.com/watch?v=RYQbmjLgubM)
- [Continuous Integration - IBM YouTube](https://www.youtube.com/watch?v=1er2cjUq1UI)
- [Continuous Monitoring](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [The Remote Flow](https://www.notion.so/The-Remote-Flow-d90982e77a144f4f990c135f115f41c6)
- [FinOps Foundation - What is FinOps](https://www.finops.org/introduction/what-is-finops/)
- [**NOT FREE** The Phoenix Project: A Novel About IT, DevOps, and Helping Your Business Win](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/1942788290/)

여기까지 왔다면 이곳이 자신이 원하는 곳인지 아닌지 알 수 있을 것입니다. 다음에 뵙겠습니다. [Day 4](day04.md).
