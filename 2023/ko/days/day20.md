# IAST와 DAST의 결합 - 실습 시간

IAST와 DAST가 무엇인지 배웠다면 이제 실제 애플리케이션에서 이 프로세스를 사용하여 취약점을 찾아보는 실습을 해보겠습니다.

**참고: 오픈 소스 IAST 구현이 없으므로 상용 솔루션을 사용해야 합니다.
무료 티어가 있으므로 비용을 지불하지 않고도 실습을 따라갈 수 있으니 걱정하지 마세요.



여기에는 테스트하고 활용할 취약한 Java 애플리케이션, 손쉬운 설정을 위한 Docker 및 Docker Compose, IAST 솔루션용 [Contrast Community Edition](https://www.contrastsecurity.com/contrast-community-edition?utm_campaign=ContrastCommunityEdition&utm_source=GitHub&utm_medium=WebGoatLab)이 포함되어 있습니다.

## 전제조건

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Docker Compose](https://docs.docker.com/compose/)
- Contrast CE 계정. [이 곳](https://www.contrastsecurity.com/contrast-community-edition?utm_campaign=ContrastCommunityEdition&utm_source=GitHub&utm_medium=WebGoatLab)에서 무료로 회원가입하세요.

**참고** : 이 글과 90일간의 데브옵스 프로그램의 작성자는 Contrast Security와 어떤 식으로든 관련이 있거나 관련이 있습니다.
오픈 소스 솔루션이 없고, 이 솔루션에는 결제나 신용카드 제공이 필요 없는 무료 티어가 있기 때문에 이 상용 솔루션을 사용하고 있습니다.

1. 오픈 소스 IAST 구현이 없기 때문에 일부 무료 라이선스가 있는 상용 소스를 사용합니다. 이를 위해 2 개의 구성 요소가 필요합니다:
   IAST 솔루션 - <https://github.com/rstatsinger/contrast-java-webgoat-docker>. Mac 또는 Linux 환경에 설치된 docker 및 docker-compose가 필요합니다(이 실습은 Mint에서 테스트되었습니다). README를 따라 Contrast에서 계정을 생성하세요.

## Getting started


시작하려면 [리포지토리](https://github.com/rstatsinger/contrast-java-webgoat-docker)를 복제합니다.

Contrast Security에서 자격 증명을 받습니다.
오른쪽 상단의 `Organization Settings` -> `Agent`에서 당신의 이름을 클릭합니다.
그리고 `Agent Username`, `Agent Service Key` and `API Key`를 가져옵니다.
새로 복제된 리포지토리의 `.env.template` 파일에서 이 값으로 바꿉니다.

**참고:** 이 값은 비밀입니다.
Git에 커밋하지 마세요.
실수로 이 값을 커밋하지 않도록 `.env.template`을 `.gitignore` 아래에 두는 것이 가장 좋습니다.


## 취약한 애플리케이션 실행하기

취약한 애플리케이션을 실행합니다.

```sh
./run.sh
```

또는

```sh
docker compose up
```

준비가 완료되면 <http://localhost:8080/WebGoat>에서 애플리케이션 UI에 액세스할 수 있습니다.

## 피해 입히기

이제 취약한 애플리케이션이 생겼으니 이를 익스플로잇해 보겠습니다.

1. [여기](https://www.zaproxy.org/download/)에서 ZAP 프록시를 설치합니다.

   가장 쉬운 방법은 DAST 스캐너를 이용하는 것입니다.
   이러한 스캐너 중 하나는 [ZAP Proxy](https://www.zaproxy.org/)입니다.
   무료 오픈소스 웹 앱 스캐너입니다.

2. [여기](https://github.com/Grunny/zap-cli)에서 `zap-cli`를 설치합니다.

   다음으로 `zap-cli`를 설치합니다.
   `zap-cli`는 ZAP Proxy를 위한 오픈소스 CLI입니다.

3. ZAP 프록시 실행

   설치된 위치에서 ZAP 프록시를 실행합니다.
   Linux Mint에서는 기본적으로 `/opt/zaproxy`에 있습니다.
   MacOS에서는 `Applications`에 있습니다.

4. `ZAP_API_KEY` 및 `ZAP_PORT`에 대한 환경 변수를 설정합니다.

   ZAP 프록시에서 이 값을 가져옵니다.
   `Options...` -> `API`로 이동하여 API 키를 가져옵니다.

   `Options...` -> `Network` -> `Local Servers/Proxies`로 이동하여 포트를 설정하고 가져옵니다.

5. `zap-cli`로 몇 가지 명령을 실행합니다.

   예시:

   ```sh
   zap-cli quick-scan -s all --ajax-spider -r http://127.0.0.1:8080/WebGoat/login.mvc
   ```

   또는 [리포지토리](https://github.com/rstatsinger/contrast-java-webgoat-docker/blob/master/Lab-WebGoat.pdf)에 있는 지침에 따라 취약한 애플리케이션을 손상시킬 수 있습니다.

6. Constrast에서 결과 관찰

   어느 쪽이든, Contrast에서 애플리케이션의 **Vulnerabilities** 탭으로 이동하면, Contrast가 취약점을 감지한 것을 확인할 수 있을 것입니다.
   조치를 취하라는 경고가 표시됩니다.

## 보너스: 이미지 스캐닝

애플리케이션의 동작을 관찰하여 공격을 탐지하는 데 IAST 솔루션이 어떻게 도움이 되었는지 살펴보았습니다.
이러한 공격을 애초에 막을 수 있었는지 살펴봅시다.

이 데모에 사용한 취약한 애플리케이션은 컨테이너로 된 패키지였습니다.
[14일차](day14.md)와 [15일차](day15.md)에서 배운 `grype` 스캐너를 통해 이 컨테이너를 스캔하고 결과를 확인해 보겠습니다.

```sh
$ grype contrast-java-webgoat-docker-webgoat
 ✔ Vulnerability DB        [no update available]
 ✔ Loaded image
 ✔ Parsed image
 ✔ Cataloged packages      [316 packages]
 ✔ Scanned image           [374 vulnerabilities]
NAME                 INSTALLED               FIXED-IN                TYPE          VULNERABILITY        SEVERITY
apt                  1.8.2.3                                         deb           CVE-2011-3374        Negligible
axis                 1.4                                             java-archive  GHSA-55w9-c3g2-4rrh  Medium
axis                 1.4                                             java-archive  GHSA-96jq-75wh-2658  Medium
bash                 5.0-4                                           deb           CVE-2019-18276       Negligible
bash                 5.0-4                   (won't fix)             deb           CVE-2022-3715        High
bsdutils             1:2.33.1-0.1                                    deb           CVE-2022-0563        Negligible
bsdutils             1:2.33.1-0.1            (won't fix)             deb           CVE-2021-37600       Low
commons-beanutils    1.8.3                                           java-archive  CVE-2014-0114        High
commons-beanutils    1.8.3                                           java-archive  CVE-2019-10086       High
commons-beanutils    1.8.3                   1.9.2                   java-archive  GHSA-p66x-2cv9-qq3v  High
commons-beanutils    1.8.3                   1.9.4                   java-archive  GHSA-6phf-73q6-gh87  High
commons-collections  3.2.1                                           java-archive  CVE-2015-6420        High
commons-collections  3.2.1                   3.2.2                   java-archive  GHSA-6hgm-866r-3cjv  High
commons-collections  3.2.1                   3.2.2                   java-archive  GHSA-fjq5-5j5f-mvxh  Critical
commons-fileupload   1.3.1                                           java-archive  CVE-2016-1000031     Critical
commons-fileupload   1.3.1                                           java-archive  CVE-2016-3092        High
commons-fileupload   1.3.1                   1.3.2                   java-archive  GHSA-fvm3-cfvj-gxqq  High
commons-fileupload   1.3.1                   1.3.3                   java-archive  GHSA-7x9j-7223-rg5m  Critical
commons-io           2.4                                             java-archive  CVE-2021-29425       Medium
commons-io           2.4                     2.7                     java-archive  GHSA-gwrp-pvrq-jmwv  Medium
coreutils            8.30-3                                          deb           CVE-2017-18018       Negligible
coreutils            8.30-3                  (won't fix)             deb           CVE-2016-2781        Low
curl                 7.64.0-4+deb10u3                                deb           CVE-2021-22922       Negligible
curl                 7.64.0-4+deb10u3                                deb           CVE-2021-22923       Negligible
<truncated>
```

보시다시피 이 이미지는 취약점으로 가득 차 있습니다.

각 취약점을 자세히 살펴보면 RCE(원격 코드 실행), SQL 인젝션, XML 외부 엔티티 취약점 등과 같은 취약점이 있음을 알 수 있습니다.

## 주간 요약

IAST와 DAST는 애플리케이션의 동작을 모니터링하여 애플리케이션의 취약점을 찾는 데 도움이 되는 중요한 방법입니다.
이 작업은 애플리케이션이 이미 배포된 후에 수행됩니다.

컨테이너 이미지 스캔은 컨테이너 내부에 있는 라이브러리를 기반으로 애플리케이션의 취약점을 찾는 데 도움이 될 수 있습니다.

이미지 스캔과 IAST/DAST는 상호 배타적이지 않습니다.
둘 다 보안 SDLC에서 각자의 역할을 하며 공격자보다 먼저 다양한 문제를 발견하는 데 도움이 될 수 있습니다.

[21일차](day21.md)에 뵙겠습니다.
