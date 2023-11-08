# 컨테이너 이미지 스캐닝

컨테이너 이미지는 이미지 매니페스트, 파일 시스템 및 이미지 구성으로 구성됩니다. [1](https://opencontainers.org/about/overview/)

예를 들어, Java 애플리케이션용 컨테이너 이미지의 파일 시스템에는 Linux 파일 시스템, JVM 및 애플리케이션을 나타내는 JAR/WAR 파일이 있습니다.

컨테이너를 사용하는 경우 CI/CD 파이프라인의 중요한 부분은 이러한 컨테이너에 알려진 취약성이 있는지 스캐닝하는 프로세스여야 합니다.
이를 통해 컨테이너 내부에 있는 취약성의 수에 대한 귀중한 정보를 얻을 수 있으며, 취약한 애플리케이션을 운영 환경에 배포하고 이러한 취약성으로 인해 해킹되는 것을 방지할 수 있습니다.

2021년 말 발견된 [Log4Shell](https://en.wikipedia.org/wiki/Log4Shell) 취약성을 예로 들어보겠습니다.
너무 자세히 설명하지 않아도 응용프로그램에 이 취약성이 있다는 것은 공격자가 서버에서 임의 코드를 실행할 수 있다는 것을 의미합니다.
이 취약성은 가장 인기 있는 Java 라이브러리 중 하나인 [Log4j](https://logging.apache.org/log4j/2.x/) 에 있기 때문에 더욱 악화되었습니다.
꽤나 나쁘네요!

그렇다면 우리가 취약한지 어떻게 알 수 있을까요?

정답은 **Image Scanning**입니다.

이미지 스캐닝 프로세스는 컨테이너 내부를 살펴보고, 설치된 패키지 목록(리눅스 패키지일 수도 있지만 자바, Go, 자바스크립트 패키지 등일 수도 있음)을 가져오고, 패키지 목록을 각 패키지에 대해 알려진 취약성 데이터베이스와 교차 참조하고, 최종적으로 지정된 컨테이너 이미지에 대한 취약성 목록을 생성하는 것으로 구성됩니다.

컴퓨터 로컬이나 CI/CD 파이프라인에서 컨테이너 이미지를 바로 설치하고 스캔을 시작할 수 있는 오픈 소스 및 독점 이미지 스캐너가 많이 있습니다.
가장 인기 있는 두 가지는 [Triby](https://github.com/aquasecurity/trivy) 와 [Grype](https://github.com/anchore/grype) 입니다.
일부 독점 제품은 [Snyk](https://docs.snyk.io/products/snyk-container/snyk-cli-for-container-security) (계정 필요, Free tier 있음) 및 [VMware Carbon Black](https://carbonblack.vmware.com/resource/carbon-black-cloud-container-security-faq#overview) (계정 필요, Free tier 없음)입니다.

컨테이너 이미지를 스캔하는 것은 다음 중 하나를 설치하고 실행하는 것처럼 간단합니다:

```console
$ grype ubuntu:latest
 ✔ Vulnerability DB        [updated]
 ✔ Pulled image
 ✔ Loaded image
 ✔ Parsed image
 ✔ Cataloged packages      [101 packages]
 ✔ Scanned image           [16 vulnerabilities]
NAME          INSTALLED                 FIXED-IN  TYPE  VULNERABILITY   SEVERITY
bash          5.1-6ubuntu1                        deb   CVE-2022-3715   Medium
coreutils     8.32-4.1ubuntu1                     deb   CVE-2016-2781   Low
gpgv          2.2.27-3ubuntu2.1                   deb   CVE-2022-3219   Low
libc-bin      2.35-0ubuntu3.1                     deb   CVE-2016-20013  Negligible
libc6         2.35-0ubuntu3.1                     deb   CVE-2016-20013  Negligible
libncurses6   6.3-2                               deb   CVE-2022-29458  Negligible
libncursesw6  6.3-2                               deb   CVE-2022-29458  Negligible
libpcre3      2:8.39-13ubuntu0.22.04.1            deb   CVE-2017-11164  Negligible
libsystemd0   249.11-0ubuntu3.6                   deb   CVE-2022-3821   Medium
libtinfo6     6.3-2                               deb   CVE-2022-29458  Negligible
libudev1      249.11-0ubuntu3.6                   deb   CVE-2022-3821   Medium
login         1:4.8.1-2ubuntu2                    deb   CVE-2013-4235   Low
ncurses-base  6.3-2                               deb   CVE-2022-29458  Negligible
ncurses-bin   6.3-2                               deb   CVE-2022-29458  Negligible
passwd        1:4.8.1-2ubuntu2                    deb   CVE-2013-4235   Low
zlib1g        1:1.2.11.dfsg-2ubuntu9.2            deb   CVE-2022-42800  Medium
```

이 명령을 사용하여 `ubuntu:latest` 컨테이너 이미지를 스캔한 결과 16개의 취약점이 있는 것으로 나타났습니다.

각 취약성에는 [CVSS](https://nvd.nist.gov/vuln-metrics/cvss) 점수에 따라 심각도가 있습니다.
심각도는 '낮음'부터 '심각함'까지 다양합니다.

16개의 취약점은 많아 보이지만, 어느 것도 '크리티컬(Critical)'한 심각성을 가지고 있지 않습니다.

또한 결과 테이블의 'FIXED-IN' 열이 비어 있습니다.
이는 해당 패키지의 최신 버전에서 이 취약성이 수정되지 않음을 의미합니다.

`ubuntu:latest`는 Ubuntu의 공식 컨테이너 이미지의 최신 버전이기 때문에 이는 예상됩니다.
일반적으로 이러한 이미지는 정기적으로 업데이트되므로 이러한 이미지에 많은 취약점이 있을 것으로 예상해서는 안 됩니다.(적어도 사용 가능한 수정 사항이 있는 이미지는 아님).

오래된 이미지, 임의의 이미지, 대기업에서 지원하지 않는 이미지 또는 자신이 관리하지 않는 이미지의 경우에는 그렇지 않을 수 있습니다.

예를 들어, Docker Hub의 `springio` 조직에서 임의 이미지 2년 된 이미지를 스캔하면 훨씬 더 많은 취약점이 잠재해 있음을 알 수 있습니다:

```console
$ grype springio/petclinic:latest
 ✔ Vulnerability DB        [no update available]
 ✔ Pulled image
 ✔ Loaded image
 ✔ Parsed image
 ✔ Cataloged packages      [213 packages]
 ✔ Scanned image           [167 vulnerabilities]
NAME              INSTALLED                 FIXED-IN                   TYPE          VULNERABILITY        SEVERITY
bash              4.4.18-2ubuntu1.2                                    deb           CVE-2022-3715        Medium
bash              4.4.18-2ubuntu1.2         4.4.18-2ubuntu1.3          deb           CVE-2019-18276       Low
coreutils         8.28-1ubuntu1                                        deb           CVE-2016-2781        Low
dpkg              1.19.0.5ubuntu2.3         1.19.0.5ubuntu2.4          deb           CVE-2022-1664        Medium
e2fsprogs         1.44.1-1ubuntu1.3         1.44.1-1ubuntu1.4          deb           CVE-2022-1304        Medium
gcc-8-base        8.4.0-1ubuntu1~18.04                                 deb           CVE-2020-13844       Medium
gpgv              2.2.4-1ubuntu1.4          2.2.4-1ubuntu1.6           deb           CVE-2022-34903       Medium
gpgv              2.2.4-1ubuntu1.4          2.2.4-1ubuntu1.5           deb           CVE-2019-13050       Low
gpgv              2.2.4-1ubuntu1.4                                     deb           CVE-2022-3219        Low
gzip              1.6-5ubuntu1              1.6-5ubuntu1.2             deb           CVE-2022-1271        Medium
h2                1.4.200                   2.0.202                    java-archive  GHSA-7rpj-hg47-cx62  High
h2                1.4.200                   2.0.206                    java-archive  GHSA-h376-j262-vhq6  Critical
h2                1.4.200                                              java-archive  CVE-2021-23463       Critical
h2                1.4.200                                              java-archive  CVE-2021-42392       Critical
h2                1.4.200                                              java-archive  CVE-2022-23221       Critical
h2                1.4.200                   2.1.210                    java-archive  GHSA-45hx-wfhj-473x  Critical
jackson-databind  2.11.4                    2.12.7.1                   java-archive  GHSA-jjjh-jjxp-wpff  High
jackson-databind  2.11.4                    2.12.7.1                   java-archive  GHSA-rgv9-q543-rqg4  High
jackson-databind  2.11.4                                               java-archive  CVE-2022-42004       High
jackson-databind  2.11.4                                               java-archive  CVE-2020-36518       High
jackson-databind  2.11.4                                               java-archive  CVE-2022-42003       High
jackson-databind  2.11.4                    2.12.6.1                   java-archive  GHSA-57j2-w4cx-62h2  High
jquery            2.2.4                                                java-archive  CVE-2019-11358       Medium
jquery            2.2.4                                                java-archive  CVE-2020-11022       Medium
jquery            2.2.4                                                java-archive  CVE-2015-9251        Medium
jquery            2.2.4                                                java-archive  CVE-2020-11023       Medium
jquery            2.2.4                                                java-archive  CVE-2007-2379        Medium
jquery-ui         1.11.4                                               java-archive  CVE-2021-41184       Medium
jquery-ui         1.11.4                                               java-archive  CVE-2016-7103        Medium
jquery-ui         1.11.4                                               java-archive  CVE-2021-41182       Medium
jquery-ui         1.11.4                                               java-archive  CVE-2021-41183       Medium
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-29562       Low
libc-bin          2.27-3ubuntu1.4                                      deb           CVE-2016-20013       Negligible
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-6096        Low
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-27618       Low
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2022-23218       Low
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2016-10228       Negligible
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2019-25013       Low
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-3326        Low
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-3999        Medium
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2022-23219       Low
libc-bin          2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-35942       Low
libc-bin          2.27-3ubuntu1.4                                      deb           CVE-2009-5155        Negligible
libc-bin          2.27-3ubuntu1.4                                      deb           CVE-2015-8985        Negligible
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-3999        Medium
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2016-10228       Negligible
libc6             2.27-3ubuntu1.4                                      deb           CVE-2009-5155        Negligible
libc6             2.27-3ubuntu1.4                                      deb           CVE-2016-20013       Negligible
libc6             2.27-3ubuntu1.4                                      deb           CVE-2015-8985        Negligible
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-3326        Low
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-35942       Low
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-27618       Low
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-6096        Low
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-29562       Low
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2022-23218       Low
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2022-23219       Low
libc6             2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2019-25013       Low
libcom-err2       1.44.1-1ubuntu1.3         1.44.1-1ubuntu1.4          deb           CVE-2022-1304        Medium
libext2fs2        1.44.1-1ubuntu1.3         1.44.1-1ubuntu1.4          deb           CVE-2022-1304        Medium
libgcc1           1:8.4.0-1ubuntu1~18.04                               deb           CVE-2020-13844       Medium
libgcrypt20       1.8.1-4ubuntu1.2          1.8.1-4ubuntu1.3           deb           CVE-2021-40528       Medium
libgcrypt20       1.8.1-4ubuntu1.2          1.8.1-4ubuntu1.3           deb           CVE-2021-33560       Low
libgmp10          2:6.1.2+dfsg-2            2:6.1.2+dfsg-2ubuntu0.1    deb           CVE-2021-43618       Low
libgnutls30       3.5.18-1ubuntu1.4         3.5.18-1ubuntu1.6          deb           CVE-2021-4209        Low
libgnutls30       3.5.18-1ubuntu1.4                                    deb           CVE-2018-16868       Low
libgnutls30       3.5.18-1ubuntu1.4         3.5.18-1ubuntu1.6          deb           CVE-2022-2509        Medium
libhogweed4       3.4-1ubuntu0.1            3.4.1-0ubuntu0.18.04.1     deb           CVE-2021-3580        Medium
libhogweed4       3.4-1ubuntu0.1            3.4.1-0ubuntu0.18.04.1     deb           CVE-2018-16869       Low
liblz4-1          0.0~r131-2ubuntu3         0.0~r131-2ubuntu3.1        deb           CVE-2021-3520        Medium
liblzma5          5.2.2-1.3                 5.2.2-1.3ubuntu0.1         deb           CVE-2022-1271        Medium
libncurses5       6.1-1ubuntu1.18.04                                   deb           CVE-2019-17594       Negligible
libncurses5       6.1-1ubuntu1.18.04                                   deb           CVE-2021-39537       Negligible
libncurses5       6.1-1ubuntu1.18.04                                   deb           CVE-2022-29458       Negligible
libncurses5       6.1-1ubuntu1.18.04                                   deb           CVE-2019-17595       Negligible
libncursesw5      6.1-1ubuntu1.18.04                                   deb           CVE-2019-17595       Negligible
libncursesw5      6.1-1ubuntu1.18.04                                   deb           CVE-2021-39537       Negligible
libncursesw5      6.1-1ubuntu1.18.04                                   deb           CVE-2022-29458       Negligible
libncursesw5      6.1-1ubuntu1.18.04                                   deb           CVE-2019-17594       Negligible
libnettle6        3.4-1ubuntu0.1            3.4.1-0ubuntu0.18.04.1     deb           CVE-2021-3580        Medium
libnettle6        3.4-1ubuntu0.1            3.4.1-0ubuntu0.18.04.1     deb           CVE-2018-16869       Low
libpcre3          2:8.39-9                                             deb           CVE-2017-11164       Negligible
libpcre3          2:8.39-9                  2:8.39-9ubuntu0.1          deb           CVE-2020-14155       Negligible
libpcre3          2:8.39-9                  2:8.39-9ubuntu0.1          deb           CVE-2019-20838       Low
libsepol1         2.7-1                     2.7-1ubuntu0.1             deb           CVE-2021-36086       Low
libsepol1         2.7-1                     2.7-1ubuntu0.1             deb           CVE-2021-36085       Low
libsepol1         2.7-1                     2.7-1ubuntu0.1             deb           CVE-2021-36087       Low
libsepol1         2.7-1                     2.7-1ubuntu0.1             deb           CVE-2021-36084       Low
libss2            1.44.1-1ubuntu1.3         1.44.1-1ubuntu1.4          deb           CVE-2022-1304        Medium
libssl1.1         1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.15  deb           CVE-2022-0778        High
libssl1.1         1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.20  deb           CVE-2022-2097        Medium
libssl1.1         1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.13  deb           CVE-2021-3712        Medium
libssl1.1         1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.13  deb           CVE-2021-3711        High
libssl1.1         1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.17  deb           CVE-2022-1292        Medium
libssl1.1         1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.19  deb           CVE-2022-2068        Medium
libstdc++6        8.4.0-1ubuntu1~18.04                                 deb           CVE-2020-13844       Medium
libsystemd0       237-3ubuntu10.47          237-3ubuntu10.56           deb           CVE-2022-2526        Medium
libsystemd0       237-3ubuntu10.47                                     deb           CVE-2022-3821        Medium
libsystemd0       237-3ubuntu10.47          237-3ubuntu10.49           deb           CVE-2020-13529       Low
libsystemd0       237-3ubuntu10.47          237-3ubuntu10.49           deb           CVE-2021-33910       High
libtinfo5         6.1-1ubuntu1.18.04                                   deb           CVE-2019-17595       Negligible
libtinfo5         6.1-1ubuntu1.18.04                                   deb           CVE-2021-39537       Negligible
libtinfo5         6.1-1ubuntu1.18.04                                   deb           CVE-2019-17594       Negligible
libtinfo5         6.1-1ubuntu1.18.04                                   deb           CVE-2022-29458       Negligible
libudev1          237-3ubuntu10.47          237-3ubuntu10.49           deb           CVE-2020-13529       Low
libudev1          237-3ubuntu10.47          237-3ubuntu10.49           deb           CVE-2021-33910       High
libudev1          237-3ubuntu10.47                                     deb           CVE-2022-3821        Medium
libudev1          237-3ubuntu10.47          237-3ubuntu10.56           deb           CVE-2022-2526        Medium
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2022-23219       Low
locales           2.27-3ubuntu1.4                                      deb           CVE-2016-20013       Negligible
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-3999        Medium
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2016-10228       Negligible
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2019-25013       Low
locales           2.27-3ubuntu1.4                                      deb           CVE-2009-5155        Negligible
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-35942       Low
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2022-23218       Low
locales           2.27-3ubuntu1.4                                      deb           CVE-2015-8985        Negligible
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-27618       Low
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-29562       Low
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2021-3326        Low
locales           2.27-3ubuntu1.4           2.27-3ubuntu1.5            deb           CVE-2020-6096        Low
log4j-api         2.13.3                                               java-archive  CVE-2021-45105       Medium
log4j-api         2.13.3                                               java-archive  CVE-2021-44832       Medium
log4j-to-slf4j    2.13.3                                               java-archive  CVE-2021-44832       Medium
log4j-to-slf4j    2.13.3                                               java-archive  CVE-2021-45105       Medium
logback-core      1.2.3                     1.2.9                      java-archive  GHSA-668q-qrv7-99fm  Medium
login             1:4.5-1ubuntu2                                       deb           CVE-2013-4235        Low
login             1:4.5-1ubuntu2            1:4.5-1ubuntu2.2           deb           CVE-2018-7169        Low
ncurses-base      6.1-1ubuntu1.18.04                                   deb           CVE-2022-29458       Negligible
ncurses-base      6.1-1ubuntu1.18.04                                   deb           CVE-2019-17595       Negligible
ncurses-base      6.1-1ubuntu1.18.04                                   deb           CVE-2019-17594       Negligible
ncurses-base      6.1-1ubuntu1.18.04                                   deb           CVE-2021-39537       Negligible
ncurses-bin       6.1-1ubuntu1.18.04                                   deb           CVE-2021-39537       Negligible
ncurses-bin       6.1-1ubuntu1.18.04                                   deb           CVE-2022-29458       Negligible
ncurses-bin       6.1-1ubuntu1.18.04                                   deb           CVE-2019-17595       Negligible
ncurses-bin       6.1-1ubuntu1.18.04                                   deb           CVE-2019-17594       Negligible
openssl           1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.13  deb           CVE-2021-3712        Medium
openssl           1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.19  deb           CVE-2022-2068        Medium
openssl           1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.20  deb           CVE-2022-2097        Medium
openssl           1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.13  deb           CVE-2021-3711        High
openssl           1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.15  deb           CVE-2022-0778        High
openssl           1.1.1-1ubuntu2.1~18.04.9  1.1.1-1ubuntu2.1~18.04.17  deb           CVE-2022-1292        Medium
passwd            1:4.5-1ubuntu2            1:4.5-1ubuntu2.2           deb           CVE-2018-7169        Low
passwd            1:4.5-1ubuntu2                                       deb           CVE-2013-4235        Low
perl-base         5.26.1-6ubuntu0.5         5.26.1-6ubuntu0.6          deb           CVE-2020-16156       Medium
snakeyaml         1.27                                                 java-archive  GHSA-w37g-rhq8-7m4j  Medium
snakeyaml         1.27                      1.32                       java-archive  GHSA-9w3m-gqgf-c4p9  Medium
snakeyaml         1.27                      1.31                       java-archive  GHSA-3mc7-4q67-w48m  High
snakeyaml         1.27                      1.31                       java-archive  GHSA-c4r9-r8fh-9vj2  Medium
snakeyaml         1.27                      1.31                       java-archive  GHSA-hhhw-99gj-p3c3  Medium
snakeyaml         1.27                      1.31                       java-archive  GHSA-98wm-3w3q-mw94  Medium
spring-core       5.3.6                                                java-archive  CVE-2022-22950       Medium
spring-core       5.3.6                                                java-archive  CVE-2022-22965       Critical
spring-core       5.3.6                                                java-archive  CVE-2021-22096       Medium
spring-core       5.3.6                                                java-archive  CVE-2022-22968       Medium
spring-core       5.3.6                                                java-archive  CVE-2022-22970       Medium
spring-core       5.3.6                                                java-archive  CVE-2022-22971       Medium
spring-core       5.3.6                                                java-archive  CVE-2021-22118       High
spring-core       5.3.6                                                java-archive  CVE-2016-1000027     Critical
spring-core       5.3.6                                                java-archive  CVE-2021-22060       Medium
tar               1.29b-2ubuntu0.2          1.29b-2ubuntu0.3           deb           CVE-2021-20193       Low
zlib1g            1:1.2.11.dfsg-0ubuntu2    1:1.2.11.dfsg-0ubuntu2.2   deb           CVE-2022-37434       Medium
zlib1g            1:1.2.11.dfsg-0ubuntu2    1:1.2.11.dfsg-0ubuntu2.1   deb           CVE-2018-25032       Medium
zlib1g            1:1.2.11.dfsg-0ubuntu2                               deb           CVE-2022-42800       Medium
```

여기서 우리는 훨씬 더 많은 `중요한` 취약점을 볼 수 있을 뿐만 아니라 많은 취약점이 해당 종속성의 최신 버전에서 수정되었음을 알 수 있습니다.
즉, 해당 종속성의 간단한 버전 업데이트만으로도 취약점을 제거하고 이미지를 더 안전하게 만들 수 있습니다.

물론 항상 그렇게 간단한 것은 아닙니다.
때로는 종속성의 새 버전에 소스 코드를 변경해야 하는 API 변경 사항이 포함되거나, 종속성과 상호 작용하는 방식에 버그를 유발하는 동작 변경이 포함되거나, 수정될 때까지 피하고 싶은 버그가 발생할 수 있습니다.

또 한 가지 언급할 만한 점은 이러한 유형의 스캔은 _알려진_ 취약점만 탐지한다는 것입니다.
즉, 보안 연구원이 발견하고 CVE를 지정한 취약점만 탐지합니다.
아직 알려지지 않은 취약점이 코드/종속성에 숨어 있을 수 있습니다. (Log4Shell은 2013년부터 사용되었지만 2021년에야 발견됨)

요약하자면, 이미지 스캔은 만병통치약이 아닙니다.
이미지 스캐너가 이미지에 취약점이 0개라고 알려준다고 해서 100% 안전하다는 의미는 아닙니다.

또한 취약점을 완화하는 것은 종속성 버전을 변경하거나 다운그레이드하는 것만큼 간단할 수도 있지만, 때로는 버전 변경을 위해 코드를 변경해야 할 수도 있기 때문에 더 까다로울 수 있습니다.

## CVEs

스캐너가 제공하는 취약점 표에서 'CVE-'로 시작하는 것을 볼 수 있습니다:

```text
bash  4.4.18-2ubuntu1.2  deb  CVE-2022-3715  Medium
```

[CVE**](https://cve.mitre.org/)는 **C**ommon **V**ulnerability and **E**xposures.의 약자입니다.

취약점을 추적하고 쉽게 검색할 수 있도록 하는 시스템입니다.

새로운 취약점이 발견될 때마다 [CNA](https://www.cve.org/ProgramOrganization/CNAs)(CVE 번호 부여 기관)에서 CVE를 할당하고 해당 취약점이 포함된 모든 구성 요소와 연관시킵니다.

이 작업이 완료되면 이 정보는 취약점 데이터베이스로 전파되며 이미지 스캐너를 통해 컨테이너에 존재하는 CVE/취약점에 대해 경고하는 데 활용될 수 있습니다.

## 요약

이제 이미지 스캔이 중요한 이유와 보안을 강화하는 데 어떻게 도움이 되는지 알게 되었습니다.
[15일차](day15.md)에서는 이미지 스캐너가 내부에서 작동하는 방식에 대해 자세히 살펴보고 SBOM과 취약성 데이터베이스 등을 살펴보겠습니다.
