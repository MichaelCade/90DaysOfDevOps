---
title: '#90DaysOfDevOps - The Big Picture: Monitoring - Day 77'
published: false
description: 90DaysOfDevOps - The Big Picture Monitoring
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048715
---

## 큰 그림: 모니터링

이 섹션에서는 모니터링이란 무엇이며 왜 필요한지에 대해 설명합니다.

### 모니터링이란 무엇인가요?

모니터링은 전체 인프라를 면밀히 주시하는 프로세스입니다.

### 모니터링이란 무엇이며 왜 필요한가요?1

애플리케이션 서버, 데이터베이스 서버, 웹 서버와 같은 다양한 특수 서버를 포함하여 수천 대의 서버를 관리한다고 가정해 봅시다. 또한 퍼블릭 클라우드 서비스 및 Kubernetes를 비롯한 추가 서비스와 다양한 플랫폼으로 인해 이 문제는 더욱 복잡해질 수 있습니다.

![](/2022/Days/Images/Day77_Monitoring1.png)

우리는 서버의 모든 서비스, 애플리케이션, 리소스가 정상적으로 실행되고 있는지 확인할 책임이 있습니다.

![](/2022/Days/Images/Day77_Monitoring2.png)

어떻게 하나요? 세 가지 방법이 있습니다:

- 모든 서버에 수동으로 로그인하여 서비스 프로세스 및 리소스에 대한 모든 데이터를 확인합니다.
- 우리를 대신하여 서버에 로그인하고 데이터를 확인하는 스크립트를 작성합니다.

이 두 가지 옵션 모두 상당한 양의 작업이 필요합니다,

세 번째 옵션은 더 쉬운 방법으로, 시중에 나와 있는 모니터링 솔루션을 사용할 수 있습니다.

Nagios와 Zabbix는 쉽게 사용할 수 있는 솔루션으로, 원하는 만큼 서버를 포함하도록 모니터링 인프라를 확장할 수 있습니다.

### Nagios

Nagios는 같은 이름의 회사에서 만든 인프라 모니터링 도구입니다. 이 도구의 오픈 소스 버전은 Nagios core라고 하며 상용 버전은 Nagios XI라고 합니다. [Nagios 웹사이트](https://www.nagios.org/)

이 도구를 사용하면 서버를 모니터링하고 서버가 충분히 활용되고 있는지 또는 해결해야 할 장애 작업이 있는지 확인할 수 있습니다.

![](/2022/Days/Images/Day77_Monitoring3.png)

기본적으로 모니터링을 통해 서버와 서비스의 상태를 확인하고 인프라의 상태를 파악하는 두 가지 목표를 달성할 수 있으며, 전체 인프라에 대한 40,000피트 뷰를 제공하여 서버가 제대로 작동하고 있는지, 애플리케이션이 제대로 작동하는지, 웹 서버에 연결할 수 있는지 여부를 확인할 수 있습니다.

특정 서버에서 지난 10주 동안 디스크가 10%씩 증가하고 있으며, 향후 4~5일 이내에 완전히 소진될 것이고 곧 응답하지 못할 것임을 알려줍니다. 디스크 또는 서버가 위험한 상태에 있을 때 경고하여 서비스 중단을 피하기 위한 적절한 조치를 취할 수 있도록 알려줍니다.

이 경우 일부 디스크 공간을 확보하여 서버에 장애가 발생하지 않고 사용자에게 영향을 미치지 않도록 할 수 있습니다.

대부분의 모니터링 엔지니어에게 어려운 질문은 무엇을 모니터링할 것인가, 또는 무엇을 모니터링하지 않을 것인가입니다.

모든 시스템에는 여러 리소스가 있는데, 이 중 어떤 리소스를 주시해야 하고 어떤 리소스를 간과할 수 있는지, 예를 들어 CPU 사용량을 모니터링해야 하는지에 대한 대답은 분명히 '예'이지만 그럼에도 불구하고 여전히 결정해야 할 사항은 시스템의 열린 포트 수를 모니터링해야 하는지 여부입니다. 범용 서버인 경우 상황에 따라 필요하지 않을 수도 있지만 웹 서버인 경우 다시 모니터링해야 할 수도 있습니다.

### 지속적인 모니터링

모니터링은 새로운 항목이 아니며 지속적인 모니터링도 많은 기업이 수년 동안 채택해 온 이상입니다.

모니터링과 관련하여 세 가지 주요 영역에 중점을 둡니다.

- 인프라 모니터링
- 애플리케이션 모니터링
- 네트워크 모니터링

이 세션에서 두 가지 일반적인 시스템과 도구에 대해 언급했지만, 사용 가능한 도구는 매우 많다는 점에 유의해야 합니다. 모니터링 솔루션의 진정한 이점은 무엇을 모니터링해야 하고 무엇을 모니터링하지 않아야 하는지에 대한 질문에 답하는 데 시간을 할애했을 때 나타납니다.

어떤 플랫폼에서든 모니터링 솔루션을 켜면 정보를 수집하기 시작하지만, 그 정보가 너무 많으면 솔루션의 이점을 제대로 활용하기 어렵기 때문에 시간을 들여 구성해야 합니다.

다음 세션에서는 모니터링 도구를 직접 사용해 보고 무엇을 모니터링할 수 있는지 살펴보겠습니다.

## 자료

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)

[Day 78](day78.md)에서 봐요!
