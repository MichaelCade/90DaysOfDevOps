# Container Image Scanning Advanced

## SBOM

**SBOM** stands for **S**oftware **B**ill **O**f **M**aterials.

It is a list of all the components that make up a software application or system.
It includes information about the various third-party libraries, frameworks, and other open-source or proprietary components that are used to build the software.
An SBOM can also include details about the versions of these components, their licensing information, and any known vulnerabilities or security issues.

The objective of an SBOM is to list these components, providing software users visibility over what is included in a software product, and allowing them to avoid components that can be harmful for security or legal reasons.

Usage of SBOMs became more common the past years, after few big supply chain attacks [this](https://www.immuniweb.com/blog/5-biggest-supply-chain-attacks-in-2022-so-far.html) and [last year](https://cyolo.io/blog/top-5-supply-chain-attacks-of-2021/).

In the context of a container image, an SBOM for a container image will contain:

- the Linux packages and libraries installed in the containers
- the language-specific packages installed for the application running in the container (e.g. Python packages, Go packages, etc.)

There are tool that can help you extract the SBOM from a container images.

One such tool is [syft](https://github.com/anchore/syft).

For example, we can use syft to generate the SBOM for the `ubuntu:latest` container image:

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

We see that the SBOM not only contains the packages and libraries installed inside the container image,
but also list their types and versions.
We can use now cross-reference this list with a vulnerability database to see whether we have any vulnerabilities inside the container.

So what is a **Vulnerability Database**?

## Vulnerability database

A vulnerability database is a collection of information about known vulnerabilities in software, hardware, and other systems.
It typically includes details about the nature of the vulnerability, such as the type of vulnerability, the severity of the vulnerability, and the potential impact of the vulnerability.
A vulnerability database may also include information about how the vulnerability can be exploited, and about any available patches or fixes for the vulnerability.

Some vulnerability databases are [vuldb.com](https://vuldb.com/), [NIST](https://nvd.nist.gov/vuln), [cvedetails.com](https://www.cvedetails.com/) and [Snyk Vulnerability Database](https://security.snyk.io/).

They provide APIs or raw data that you can download, and cross-reference the packages in our SBOM with the vulnerability information about.
This way, we can find if any of our packages has vulnerabilities that we need to care about.

Usually we can also find information about the library version in which this vulnerability has been introduced and whether it has been fixed in a newer version.
Using this information, we can decide whether to update/downgrade our dependency to mitigate the vulnerability.
As we already established in [Day 14](./day14.md), updating a dependency is not always trivial, because sometimes this update comes with behaviour or API changes.

Another important piece of information about a vulnerability is its **CVSS Score**.

## CVSS

**CVSS** stands for **C**ommon **V**ulnerability **S**coring **S**ystem.

It provides a way to capture the principal characteristics of a vulnerability and produce a numerical score reflecting its severity.
The numerical score can then be translated into a qualitative representation (such as low, medium, high, and critical) to help organizations properly assess and prioritize their vulnerability management processes.

Basically, one vulnerability can be more severe than another.
We need a system that can objectively rank vulnerabilities based on how easy they are to exploit and how much damage they can cause.

This is where CVSS comes in.

CVSS v3 defines 8 criteria based on which the CVSS score is calculated.
These criteria are:

### Attack Vector

Reflects the context by which vulnerability exploitation is possible.

Possible values: **Network(N)**, **Adjacent(A)**, **Local(L)**, **Physical(P)**

### Attack Complexity

Describes the conditions beyond the attacker's control that must exist in order to exploit the vulnerability.

Possible values: **Low(L)**, **High(H)**

### Priviledges Required

Describes the level of privileges an attacker must possess before successfully exploiting the vulnerability.

Possible values: **None(N)**, **Low(L)**, **High(H)**

### User Interaction

The requirement for a user, other than the attacker, to participate in the successful compromise of the vulnerable component.

Possible values: **None(N)**, **Required(R)**

### Scope

The ability for a vulnerability in one software component to impact resources beyond its means, or privileges.

Possible values: **Unchanged(U)**, **Changed(C)**

### Confidentiality

The impact to the confidentiality of the information resources managed by a software component due to a successfully exploited vulnerability.

Possible values: **None(N)**, **Low(L)**, **High(H)**

### Integrity

The impact to integrity of a successfully exploited vulnerability.

Possible values: **None(N)**, **Low(L)**, **High(H)**

### Availability

The impact to the availability of the impacted component resulting from a successfully exploited vulnerability.

Possible values: **None(N)**, **Low(L)**, **High(H)**

The combination of these 8 vectors determines the CVSS score.
It is between 0 and 10.
0 being the lowest possible, and 10 being the highest (most critical).

[Here](https://www.first.org/cvss/calculator/3.0) you can find a CVSS calculator, wher you can calculate the score of each vulnerability.

## Resources

<https://www.nist.gov/itl/executive-order-improving-nations-cybersecurity>

<https://www.aquasec.com/cloud-native-academy/supply-chain-security/sbom/>


On [Day 16](day16.md) we will take a look into "Fuzzing" or Fuzz Testing. 