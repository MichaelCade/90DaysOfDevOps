# Container Image Scanning

A container image consists of an image manifest, a filesystem and an image configuration. [1](https://opencontainers.org/about/overview/)

For example, the filesystem of a container image for a Java application will have a Linux filesystem, the JVM, and the JAR/WAR file that represents our application.

If we are working with containers, an important part of our CI/CD pipeline should be the process of scanning these containers for known vulnerabilities.
This can give us valuable information about the number of vulnerabilities we have inside our containers, and can help us prevent deploying vulnerable applications to our production environment, and being hacked because of these vulnerabilities.

Let's take for example the [Log4Shell](https://en.wikipedia.org/wiki/Log4Shell) vulnerability that was discovered in late 2021.
Without going into too much detail, having this vulnerability in your application means that an attacker can execute arbitraty code on your servers.
It was made worse by the fact that this vulnerability is inside one of the most popular Java libraries - [Log4j](https://logging.apache.org/log4j/2.x/).
Pretty bad!

So how can we know if we are vulnerable?

The answer is **Image Scanning**.

The image scanning process consists of looking inside the container, getting the list of installed packages (that could be Linux packages, but also Java, Go, JavaScript packages, etc.), cross-referencing the package list against a database of known vulnerabilities for each package, and in the end producing a list of vulnerabilities for the given container image.

There are many open-source and proprietary image scanners, that you can install and start scanning your container images right away, either locally of your machine or in your CI/CD pipeline.
Two of the most popular ones are [Trivy](https://github.com/aquasecurity/trivy) and [Grype](https://github.com/anchore/grype).
Some proprietary ones are [Snyk](https://docs.snyk.io/products/snyk-container/snyk-cli-for-container-security) (requires an account, has a free tier) and [VMware Carbon Black](https://carbonblack.vmware.com/resource/carbon-black-cloud-container-security-faq#overview) (requires an account, no free tier).

Scanning a container image is as simple as installing one of these and running:

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

With this command we scanned the `ubuntu:latest` container image and found that it has 16 vulnerabilities.

Each vulnerability has a severity, based on its [CVSS](https://nvd.nist.gov/vuln-metrics/cvss) score.
Severities vary from `Low` to `Critical`.

16 vulnerabilities may seem a lot, but notice that none of them has a `Critical` severity.

We also see that the `FIXED-IN` column of the results table is empty.
This means that none of the vulnerabilities is fixed in newer versions of its package.

This is expected, because `ubuntu:latest` is the latest version of the official container image for Ubuntu.
Usually these images are updated regularly, so you should not expect to see many vulnerabilities in them (atleast not ones with available fixes).

This might not be case with older images, random images, not backed by big companies or your own images if you are not taking care of them.

For example, if we scan some random image 2 year old image from the `springio` organization on Docker Hub we see that there are a lot more vulnerabilities lurking:

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

Here we not only see a lot more `Critical` vulnerabilities, but we also see that a lot of them are fixed in newer version of the dependency they are coming from.
That means that a simple version bump of the given dependency will remove the vulnerability and make our image safer.

Of course, that is not always so simple.
Sometimes, a new version of a dependency may include breaking API changes that require change in your source code, it may include a behaviour change that leads to bugs in the way we interact with the dependency, or it might introduce bugs that we want to avoid until fixed.

Another thing worth mentioning is that this type of scanning only detects _known_ vulnerabilities.
That is, vulnerabilities that have been found by security researchers and that have assigned CVEs.
There might be still vulnerabilities that are not known and are just lurking in your code/dependencies (Log4Shell has been in the wild since 2013, only found in 2021).

In summary, image scanning is not a silver bullet.
If an image scanner tells you that you have 0 vulnerabilities in your image, that does not mean that you are 100% secure.

Also, mitigating vulnerabilities can be as simple as bumping a version of a dependency (or downgrading one), but sometimes it can be more tricky because that version bump might require a change in your code.

## CVEs

In the vulnerability table provided by our scanner we see something that starts with `CVE-`:

```text
bash  4.4.18-2ubuntu1.2  deb  CVE-2022-3715  Medium
```

[**CVE**](https://cve.mitre.org/) stands for **C**ommon **V**ulnerability and **E**xposures.

It is a system that allows us to track vulnerabilities and be able to easily search for them.

Each time a new vulnerability is found, it is assigned a CVE by the [CNA](https://www.cve.org/ProgramOrganization/CNAs) (CVE Numbering Authority) and associated with all components that contain that vulnerability.

Once this is done, this information is propagated to the vulnerabilities databases and can be leveraged by image scanners to warn about CVEs/vulnerabilities that are present in our container.

## Summary

Now we know why image scanning is important and how it can help us be more secure.
In [Day 15](day15.md) we are going to  dive deeper into the way the image scanners work under the hood, looking into things like SBOMs and vulnerability databases.
