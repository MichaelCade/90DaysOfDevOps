# IAST and DAST in conjunction - lab time

After learning what IAST and DAST are it's time to get our hands dirty and perform an exercise in which we use these processes to find vulnerabilities in real applications.

**NOTE:** There are no open-source IAST implementations, so we will have to use a commerical solution.
Don't worry, there is a free-tier, so you will be able to follow the lab without paying anything.

This lab is based on this [repo](https://github.com/rstatsinger/contrast-java-webgoat-docker).

It contains a vulnerable Java application to be tested and exploited, Docker and Docker Compose for easy setup and [Contrast Community Edition](https://www.contrastsecurity.com/contrast-community-edition?utm_campaign=ContrastCommunityEdition&utm_source=GitHub&utm_medium=WebGoatLab) for IAST solution.

## Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop/)
- [Docker Compose](https://docs.docker.com/compose/)
- Contrast CE account. Sign up for free [here](https://www.contrastsecurity.com/contrast-community-edition?utm_campaign=ContrastCommunityEdition&utm_source=GitHub&utm_medium=WebGoatLab).

**NOTE:** The authors of this article and of the 90 Days of DevOps program are in way associated or affilited with Contrast Security.
We are using this commercial solution, because there is not an open-source one, and because this one has a free-tier that does not require paying or providing a credit card.

1. As there are no open-source IAST implementation will use a commercial one with some free licenses. For this purpose, you will need 2 componenets:
   IAST solution from here - <https://github.com/rstatsinger/contrast-java-webgoat-docker>. You need docker and docker-compose installed in mac or linux enviroment (this lab is tested on Mint). Please follow the README to create account in Contrast.

## Getting started

To start, clone the [repository](https://github.com/rstatsinger/contrast-java-webgoat-docker).

Get your credentials from Contrast Security.
Click on your name in the top-right corner -> `Organization Settings` -> `Agent`.
Get the values for `Agent Username`, `Agent Service Key` and `API Key`.

Replace these values in the `.env.template` file in the newly cloned repository.

**NOTE:** These values are secret.
Do not commit them to Git.
It's best to put the `.env.template` under `.gitignore` so that you don't commit these values by mistake.

## Running the vulnerable application

To run the vulnerable application, run:

```sh
./run.sh
```

or

```sh
docker compose up
```

Once ready, the application UI will be accessible on <http://localhost:8080/WebGoat>.

## Do some damage

Now that we have a vulnerable application let's try to exploit it.

1. Install ZAP Proxy from [here](https://www.zaproxy.org/download/)

   An easy way to do that is via a DAST scanner.
   One such scanner is [ZAP Proxy](https://www.zaproxy.org/).
   It is a free and open-source web app scanner.

2. Install `zap-cli` from [here](https://github.com/Grunny/zap-cli)

   Next, install `zap-cli`.
   `zap-cli` is an open-source CLI for ZAP Proxy.

3. Run ZAP proxy

   Run ZAP Proxy from its installed location.
   In Linux Mint it is by default in `/opt/zaproxy`.
   In MacOS it is in `Applications`.

4. Set env variables for `ZAP_API_KEY` and `ZAP_PORT`

   Get these values from ZAP Proxy.
   Go to `Options...` -> `API` to get the API Key.

   Go to `Options...` -> `Network` -> `Local Servers/Proxies` to configure and obtain the port.

5. Run several commands with `zap-cli`

   For example:

   ```sh
   zap-cli quick-scan -s all --ajax-spider -r http://127.0.0.1:8080/WebGoat/login.mvc
   ```

   Alternatively, you can follow the instructions in the [repo](https://github.com/rstatsinger/contrast-java-webgoat-docker/blob/master/Lab-WebGoat.pdf)
   to cause some damage to the vulnerable application.

6. Observe findings in Constrast

   Either way, if you go to the **Vulnerabilities** tab for your application in Contrast you should be able to see that Contrast detected the vulnerabilities
   and is warning you to take some action.

## Bonus: Image Scanning

We saw how an IAST solution helped us detect attacks by observing the behaviour of the application.
Let's see whether we could have done something to prevent these attacks in the first place.

The vulnerable application we used for this demo was packages as a container.
Let's scan this container via the `grype` scanner we learned about in Days [14](day14.md) and [15](day15.md) and see the results.

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

As we can see this image is full with vulnerabilities.

If we dive into each one we will see we have vulnerabilities like RCE (Remote Code Execution), SQL Injection, XML External Entity Vulnerability, etc.

## Week Summary

IAST and DAST are important methods that can help us find vulnerabilities in our application via monitoring its behaviour.
This is done once the application is already deployed.

Container Image Scanning can help us find vulnerabilities in our application based on the library that are present inside the container.

Image Scanning and IAST/DAST are not mutually-exclusive.
They both have their place in a Secure SDLC and can help us find different problems before the attackers do.

See you on [Day 21](day21.md).
