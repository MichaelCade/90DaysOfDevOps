# 컨테이너 이미지 스캐닝 고급

## SBOM

**SBOM**은 **S**oftware **B**ill **O**f **M**aterials의 약자입니다.

이것은 소프트웨어 애플리케이션 또는 시스템을 구성하는 모든 구성 요소의 목록입니다.
여기에는 소프트웨어를 구축하는 데 사용되는 다양한 타사 라이브러리, 프레임워크 및 기타 오픈 소스 또는 독점 구성 요소에 대한 정보가 포함됩니다.
또한 SBOM에는 이러한 구성 요소의 버전, 라이선스 정보, 알려진 취약점 또는 보안 문제에 대한 세부 정보도 포함될 수 있습니다.

SBOM의 목적은 이러한 구성 요소를 나열하여 소프트웨어 사용자에게 소프트웨어 제품에 포함된 구성 요소에 대한 가시성을 제공하고 보안 또는 법적 이유로 유해할 수 있는 구성 요소를 피할 수 있도록 하는 것입니다.

[올해](https://www.immuniweb.com/blog/5-biggest-supply-chain-attacks-in-2022-so-far.html) 와 [작년](https://cyolo.io/blog/top-5-supply-chain-attacks-of-2021/) 지난 몇 년간 대규모 공급망 공격이 발생한 후 SBOM의 사용이 더욱 보편화되었습니다.


컨테이너 이미지의 컨텍스트에서, 컨테이너 이미지에 대한 SBOM에는 다음이 포함됩니다:

- 컨테이너에 설치된 리눅스 패키지 및 라이브러리
- 컨테이너에서 실행 중인 애플리케이션을 위해 설치된 언어별 패키지(예: Python 패키지, Go 패키지 등)

컨테이너 이미지에서 SBOM을 추출하는 데 도움이 되는 도구가 있습니다.

그러한 도구 중 하나는 [syft](https://github.com/anchore/syft)입니다.

예를 들어, syft를 사용하여 `ubuntu:latest` 컨테이너 이미지에 대한 SBOM을 생성할 수 있습니다:

```console
$ syft ubuntu
 ✔ Parsed image
 ✔ Cataloged packages      [101 packages]
NAME                 VERSION                                  TYPE
adduser              3.118ubuntu5                             deb
apt                  2.4.8                                    deb
base-files           12ubuntu4.2                              deb
base-passwd          3.5.52build1                             deb
bash                 5.1-6ubuntu1                             deb
bsdutils             1:2.37.2-4ubuntu3                        deb
coreutils            8.32-4.1ubuntu1                          deb
dash                 0.5.11+git20210903+057cd650a4ed-3build1  deb
debconf              1.5.79ubuntu1                            deb
debianutils          5.5-1ubuntu2                             deb
diffutils            1:3.8-0ubuntu2                           deb
dpkg                 1.21.1ubuntu2.1                          deb
e2fsprogs            1.46.5-2ubuntu1.1                        deb
findutils            4.8.0-1ubuntu3                           deb
gcc-12-base          12.1.0-2ubuntu1~22.04                    deb
gpgv                 2.2.27-3ubuntu2.1                        deb
grep                 3.7-1build1                              deb
gzip                 1.10-4ubuntu4.1                          deb
hostname             3.23ubuntu2                              deb
init-system-helpers  1.62                                     deb
libacl1              2.3.1-1                                  deb
libapt-pkg6.0        2.4.8                                    deb
libattr1             1:2.5.1-1build1                          deb
libaudit-common      1:3.0.7-1build1                          deb
libaudit1            1:3.0.7-1build1                          deb
libblkid1            2.37.2-4ubuntu3                          deb
libbz2-1.0           1.0.8-5build1                            deb
libc-bin             2.35-0ubuntu3.1                          deb
libc6                2.35-0ubuntu3.1                          deb
libcap-ng0           0.7.9-2.2build3                          deb
libcap2              1:2.44-1build3                           deb
libcom-err2          1.46.5-2ubuntu1.1                        deb
libcrypt1            1:4.4.27-1                               deb
libdb5.3             5.3.28+dfsg1-0.8ubuntu3                  deb
libdebconfclient0    0.261ubuntu1                             deb
libext2fs2           1.46.5-2ubuntu1.1                        deb
libffi8              3.4.2-4                                  deb
libgcc-s1            12.1.0-2ubuntu1~22.04                    deb
libgcrypt20          1.9.4-3ubuntu3                           deb
libgmp10             2:6.2.1+dfsg-3ubuntu1                    deb
libgnutls30          3.7.3-4ubuntu1.1                         deb
libgpg-error0        1.43-3                                   deb
libgssapi-krb5-2     1.19.2-2                                 deb
libhogweed6          3.7.3-1build2                            deb
libidn2-0            2.3.2-2build1                            deb
libk5crypto3         1.19.2-2                                 deb
libkeyutils1         1.6.1-2ubuntu3                           deb
libkrb5-3            1.19.2-2                                 deb
libkrb5support0      1.19.2-2                                 deb
liblz4-1             1.9.3-2build2                            deb
liblzma5             5.2.5-2ubuntu1                           deb
libmount1            2.37.2-4ubuntu3                          deb
libncurses6          6.3-2                                    deb
libncursesw6         6.3-2                                    deb
libnettle8           3.7.3-1build2                            deb
libnsl2              1.3.0-2build2                            deb
libp11-kit0          0.24.0-6build1                           deb
libpam-modules       1.4.0-11ubuntu2                          deb
libpam-modules-bin   1.4.0-11ubuntu2                          deb
libpam-runtime       1.4.0-11ubuntu2                          deb
libpam0g             1.4.0-11ubuntu2                          deb
libpcre2-8-0         10.39-3ubuntu0.1                         deb
libpcre3             2:8.39-13ubuntu0.22.04.1                 deb
libprocps8           2:3.3.17-6ubuntu2                        deb
libseccomp2          2.5.3-2ubuntu2                           deb
libselinux1          3.3-1build2                              deb
libsemanage-common   3.3-1build2                              deb
libsemanage2         3.3-1build2                              deb
libsepol2            3.3-1build1                              deb
libsmartcols1        2.37.2-4ubuntu3                          deb
libss2               1.46.5-2ubuntu1.1                        deb
libssl3              3.0.2-0ubuntu1.7                         deb
libstdc++6           12.1.0-2ubuntu1~22.04                    deb
libsystemd0          249.11-0ubuntu3.6                        deb
libtasn1-6           4.18.0-4build1                           deb
libtinfo6            6.3-2                                    deb
libtirpc-common      1.3.2-2ubuntu0.1                         deb
libtirpc3            1.3.2-2ubuntu0.1                         deb
libudev1             249.11-0ubuntu3.6                        deb
libunistring2        1.0-1                                    deb
libuuid1             2.37.2-4ubuntu3                          deb
libxxhash0           0.8.1-1                                  deb
libzstd1             1.4.8+dfsg-3build1                       deb
login                1:4.8.1-2ubuntu2                         deb
logsave              1.46.5-2ubuntu1.1                        deb
lsb-base             11.1.0ubuntu4                            deb
mawk                 1.3.4.20200120-3                         deb
mount                2.37.2-4ubuntu3                          deb
ncurses-base         6.3-2                                    deb
ncurses-bin          6.3-2                                    deb
passwd               1:4.8.1-2ubuntu2                         deb
perl-base            5.34.0-3ubuntu1.1                        deb
procps               2:3.3.17-6ubuntu2                        deb
sed                  4.8-1ubuntu2                             deb
sensible-utils       0.0.17                                   deb
sysvinit-utils       3.01-1ubuntu1                            deb
tar                  1.34+dfsg-1build3                        deb
ubuntu-keyring       2021.03.26                               deb
usrmerge             25ubuntu2                                deb
util-linux           2.37.2-4ubuntu3                          deb
zlib1g               1:1.2.11.dfsg-2ubuntu9.2                 deb
```

SBOM에는 컨테이너 이미지 내부에 설치된 패키지와 라이브러리가 포함되어 있을 뿐만 아니라,
유형과 버전도 나열되어 있습니다.
이제 이 목록을 취약점 데이터베이스와 상호 참조하여 컨테이너 내부에 취약점이 있는지 확인할 수 있습니다.

그렇다면 **취약점 데이터베이스(Vulnerability Database)**란 무엇인가요?

## 취약점 데이터베이스

취약점 데이터베이스는 소프트웨어, 하드웨어 및 기타 시스템의 알려진 취약점에 대한 정보 모음입니다.
일반적으로 취약점의 유형, 취약점의 심각도, 취약점의 잠재적 영향과 같은 취약점의 특성에 대한 세부 정보가 포함됩니다.
취약점 데이터베이스에는 취약점이 악용될 수 있는 방법과 해당 취약점에 대해 사용 가능한 패치 또는 수정 사항에 대한 정보도 포함될 수 있습니다.

일부 취약점 데이터베이스는 [vuldb.com](https://vuldb.com/), [NIST](https://nvd.nist.gov/vuln), [cvedetails.com](https://www.cvedetails.com/) 및 [Snyk 취약점 데이터베이스](https://security.snyk.io/) 등이 있습니다.

이들 기관은 사용자가 다운로드할 수 있는 API 또는 원시 데이터를 제공하며, SBOM의 패키지와 취약점 정보를 상호 참조할 수 있습니다.
이렇게 하면 패키지에 주의해야 할 취약점이 있는지 확인할 수 있습니다.

일반적으로 이 취약점이 도입된 라이브러리 버전과 최신 버전에서 취약점이 수정되었는지 여부에 대한 정보도 찾을 수 있습니다.
이 정보를 사용하여 취약점을 완화하기 위해 종속성을 업데이트/다운그레이드할지 여부를 결정할 수 있습니다.
[14일차](./day14.md)에서 이미 설명했듯이, 종속성 업데이트에는 동작 또는 API 변경이 수반되는 경우가 있기 때문에 종속성을 업데이트하는 것이 항상 간단한 것은 아닙니다.

취약점에 대한 또 다른 중요한 정보는 취약점의 **CVSS 점수**입니다.

## CVSS

**CVSS**는 **C**ommon **V**ulnerability **S**coring **S**ystem의 약자입니다.

이 시스템은 취약점의 주요 특성을 파악하고 그 심각성을 반영하는 수치 점수를 생성하는 방법을 제공합니다.
그런 다음 이 수치 점수를 정성적 표현(낮음, 중간, 높음, 심각 등)으로 변환하여 조직이 취약점 관리 프로세스를 적절히 평가하고 우선순위를 정할 수 있도록 도와줍니다.

기본적으로 하나의 취약점이 다른 취약점보다 더 심각할 수 있습니다.
익스플로잇이 얼마나 쉬운지, 얼마나 큰 피해를 입힐 수 있는지에 따라 취약점의 순위를 객관적으로 매길 수 있는 시스템이 필요합니다.

이것이 바로 CVSS가 필요한 이유입니다.

CVSS v3는 CVSS 점수를 계산하는 8가지 기준을 정의합니다.
이러한 기준은 다음과 같습니다:

### 공격 벡터 Attack Vector

취약점 악용이 가능한 컨텍스트를 반영합니다.

가능한 값 : **Network(N)**, **Adjacent(A)**, **Local(L)**, **Physical(P)**

### 공격 복잡성 Attack Complexity

취약점을 익스플로잇하기 위해 공격자가 통제할 수 없는 조건에 대해 설명합니다.

가능한 값 : **Low(L)**, **High(H)**

### 필요한 권한 Priviledges Required

공격자가 취약점을 성공적으로 익스플로잇하기 전에 보유해야 하는 권한 수준을 설명합니다.

가능한 값 : **None(N)**, **Low(L)**, **High(H)**

### 사용자 상호작용 User Interaction

공격자 이외의 사용자가 취약한 구성 요소의 성공적인 침해에 참여하기 위한 요구 사항입니다.

가능한 값 : **None(N)**, **Required(R)**

#### 범위 Scope

한 소프트웨어 구성 요소의 취약점이 리소스 또는 권한에 영향을 미칠 수 있는 능력입니다.

가능한 값 : **Unchanged(U)**, **Changed(C)**

### 기밀성 Confidentiality

성공적으로 익스플로잇된 취약점으로 인해 소프트웨어 구성 요소가 관리하는 정보 리소스의 기밀성에 미치는 영향입니다.

가능한 값 : **None(N)**, **Low(L)**, **High(H)**

### 무결성 Integrity

성공적으로 익스플로잇된 취약점이 무결성에 미치는 영향입니다.

가능한 값 : **None(N)**, **Low(L)**, **High(H)**

### 가용성  Availability

성공적으로 익스플로잇된 취약점으로 인해 영향을 받은 구성 요소의 가용성에 미치는 영향입니다.

가능한 값 : **None(N)**, **Low(L)**, **High(H)**

이 8가지 벡터의 조합에 따라 CVSS 점수가 결정됩니다.
0에서 10 사이입니다.
0이 가장 낮고 10이 가장 높습니다.(가장 위험)
[이 곳](https://www.first.org/cvss/calculator/3.0)에서 각 취약점의 점수를 계산할 수 있는 CVSS 계산기를 찾을 수 있습니다.

## 리소스

<https://www.nist.gov/itl/executive-order-improving-nations-cybersecurity>

<https://www.aquasec.com/cloud-native-academy/supply-chain-security/sbom/>


[16일차](day16.md)에서는 "Fuzzing" 또는 Fuzz Test에 대해 살펴보겠습니다.