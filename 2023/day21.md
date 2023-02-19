# Continuous Image Repository Scan

In [Day 14](day14.md), we learned what container image scanning is and why it's important.
We also learned about tools like Grype and Trivy that help us scan our container images.

However, in modern SDLCs, a DevSecOps engineer would rarely scan container images by hand, e.g., they would not be running Grype and Trivy locally and looking at every single vulnerability.
Instead, they would have the image scanning configured as part of the CI/CD pipeline.
This way, they would be sure that all the images that are being built by the pipelines are also scanned by the image scanner.
These results could then be sent by another system, where the DevSecOps engineers could look at them and take some action depending on the result.

A sample CI/CD pipeline could look like this:

0. _Developer pushes code_
1. Lint the code
2. Build the code
3. Test the code
4. Build the artifacts (container images, helm charts, etc.)
5. Scan the artifacts
6. (Optional) Send the scan results somewhere
7. (Optional) Verify the scan results and fail the pipeline if the verification fails
8. Push the artifacts to a repository

A failure in the scan or verify steps (steps 6 and 7) would mean that our container will not be pushed to our repository, and we cannot use the code we submitted.

Today, we are going to take a look at how we can set up such a pipeline and what would be a sensible configuration for one.

## Setting up a CI/CD pipeline with Grype

Let's take a look at the [Grype](https://github.com/anchore/grype) scanner.
Grype is an open-source scanner maintained by the company [Anchore](https://anchore.com/).

### Scanning an image with Grype

Scanning a container image with Grype is as simple as running:

```shell
grype <IMAGE>
```

For example, if we want to scan the `ubuntu:20.04` image, we can run:

```shell
$ grype ubuntu:20.04

 ✔ Vulnerability DB        [no update available]
 ✔ Pulled image
 ✔ Loaded image
 ✔ Parsed image
 ✔ Cataloged packages      [92 packages]
 ✔ Scanned image           [19 vulnerabilities]

NAME          INSTALLED                 FIXED-IN  TYPE  VULNERABILITY   SEVERITY
coreutils     8.30-3ubuntu2                       deb   CVE-2016-2781   Low
gpgv          2.2.19-3ubuntu2.2                   deb   CVE-2022-3219   Low
libc-bin      2.31-0ubuntu9.9                     deb   CVE-2016-20013  Negligible
libc6         2.31-0ubuntu9.9                     deb   CVE-2016-20013  Negligible
libncurses6   6.2-0ubuntu2                        deb   CVE-2021-39537  Negligible
libncurses6   6.2-0ubuntu2                        deb   CVE-2022-29458  Negligible
libncursesw6  6.2-0ubuntu2                        deb   CVE-2021-39537  Negligible
libncursesw6  6.2-0ubuntu2                        deb   CVE-2022-29458  Negligible
libpcre3      2:8.39-12ubuntu0.1                  deb   CVE-2017-11164  Negligible
libsystemd0   245.4-4ubuntu3.19                   deb   CVE-2022-3821   Medium
libtinfo6     6.2-0ubuntu2                        deb   CVE-2021-39537  Negligible
libtinfo6     6.2-0ubuntu2                        deb   CVE-2022-29458  Negligible
libudev1      245.4-4ubuntu3.19                   deb   CVE-2022-3821   Medium
login         1:4.8.1-1ubuntu5.20.04.4            deb   CVE-2013-4235   Low
ncurses-base  6.2-0ubuntu2                        deb   CVE-2021-39537  Negligible
ncurses-base  6.2-0ubuntu2                        deb   CVE-2022-29458  Negligible
ncurses-bin   6.2-0ubuntu2                        deb   CVE-2021-39537  Negligible
ncurses-bin   6.2-0ubuntu2                        deb   CVE-2022-29458  Negligible
passwd        1:4.8.1-1ubuntu5.20.04.4            deb   CVE-2013-4235   Low
```

Of course, you already know that because we did it on [Day 14](day14.md).

However, this command will only output the vulnerabilities and exit with a success code.
So if this were in a CI/CD pipeline, the pipeline would be successful even if we have many vulnerabilities.

The person running the pipeline would have to open it, see the logs and manually determine whether the results are OK.
This is tedious and error prone.

Let's see how we can enforce some rules for the results that come out of the scan.

### Enforcing rules for the scanned images

As we already established, just scanning the image does not do much except for giving us visibility into the number of vulnerabilities we have inside the image.
But what if we want to enforce a set of rules for our container images?

For example, a good rule would be "an image should not have critical vulnerabilities" or "an image should not have vulnerabilities with available fixes."

Fortunately for us, this is also something that Grype supports out of the box.
We can use the `--fail-on <SEVERITY>` flag to tell Grype to exit with a non-zero exit code if, during the scan, it found vulnerabilities with a severity higher or equal to the one we specified.
This will fail our pipeline, and the engineer would have to look at the results and fix something in order to make it pass.

Let's tried it out.
We are going to use the `springio/petclinic:latest` image, which we already found has many vulnerabilities.
You can go back to [Day 14](day14.md) or scan it yourself to see how much exactly.

We want to fail the pipeline if the image has `CRITICAL` vulnerabilities.
We are going to run the can like this:

```shell
$ grype springio/petclinic:latest --fail-on critical
 ✔ Vulnerability DB        [no update available]
 ✔ Loaded image
 ✔ Parsed image
 ✔ Cataloged packages      [212 packages]
 ✔ Scanned image           [168 vulnerabilities]

NAME        INSTALLED FIXED-IN TYPE         VULNERABILITY    SEVERITY
spring-core 5.3.6              java-archive CVE-2016-1000027 Critical
spring-core 5.3.6              java-archive CVE-2022-22965   Critical
...
1 error occurred:
    * discovered vulnerabilities at or above the severity threshold

$ echo $?
1
```

We see two things here:

- apart from the results, Grype also outputted an error that is telling us that this scan violated the rule we had defined (no CRITICAL vulnerabilities)
- Grype exited with exit code 1, which indicates failure.
  If this were a CI pipeline, it would have failed.

When this happens, we will be blocked from merging our code and pushing our container to the registry.
This means that we need to take some action to fix the failure so that we can finish our task and push our change.

Let's see what our options are.

### Fixing the pipeline

Once we encounter a vulnerability that is preventing us from publishing our container, we have a few ways we can go depending on the vulnerability.

#### 1. The vulnerability has a fix

The best-case scenario is when this vulnerability is already fixed in a newer version of the library we depend on.

One such vulnerability is this one:

```text
NAME      INSTALLED FIXED-IN TYPE         VULNERABILITY       SEVERITY
snakeyaml 1.27      1.31     java-archive GHSA-3mc7-4q67-w48m High
```

This is a `High` severity vulnerability.
It's coming from the Java package `snakeyaml`, version `1.27`.
Grype is telling us that this vulnerability is fixed in version `1.31` of the same library.

In this case, we can just upgrade the version of this library in our `pom.xml` or `build.gradle` file,
test our code to make sure nothing breaks with the new version,
and submit the code again.

This will build a new version of our container, re-scan it, and hopefully, this time, the vulnerability will not come up, and our scan will be successful.

### 2. The vulnerability does not have a fix, but it's not dangerous

Sometimes a vulnerability we encounter will not have a fix available.
These are so-called zero-day vulnerabilities that are disclosed before a fix is available.

We can see two of those in the initial scan results:

```text
NAME        INSTALLED FIXED-IN TYPE         VULNERABILITY    SEVERITY
spring-core 5.3.6              java-archive CVE-2016-1000027 Critical
spring-core 5.3.6              java-archive CVE-2022-22965   Critical
```

When we encounter such a vulnerability, we need to evaluate how severe it is and calculate the risk of releasing our software with that vulnerability in it.

We can determine that the vulnerability does not constitute any danger to our software and its consumers.
One such case might be when a vulnerability requires physical access to the servers to be exploited.
If we are sure that our physical servers are secure enough and an attacker cannot get access to them, we can safely ignore this vulnerability.

In this case, we can tell Grype to ignore this vulnerability and not fail the scan because of it.

We can do this via the `grype.yaml` configuration file, where we can list vulnerabilities we want to ignore:

```yaml
ignore:
  # This is the full set of supported rule fields:
  - vulnerability: CVE-2016-1000027
    fix-state: unknown
    package:
      name: spring-core
      version: 5.3.6
      type: java-archive
  # We can list as many of these as we want
  - vulnerability: CVE-2022-22965
  # Or list whole packages which we want to ignore
  - package:
      type: gem
```

Putting this in our configuration file and re-running the scan will make our pipeline green.

However, it is crucial that we keep track of this file and not ignore vulnerabilities that have a fix.
For example, when a fix for this vulnerability is released, it's best we upgrade our dependency and remove this vulnerability from our application.

That way, we will ensure that our application is as secure as possible and there are no vulnerabilities that can turn out to be more severe than we initially thought.

### 3. Vulnerability does not have a fix, and IT IS dangerous

The worst-case scenario is if we encounter a vulnerability that does not have a fix, and it is indeed dangerous, and there is a possibility to be exploited.

In that case, there is no right move.
The best thing we can do is sit down with our security team and come up with an action plan.

We might decide it's best to do nothing while the vulnerability is fixed.
We might decide to manually patch some stuff so that we remove at least some part of the danger.
It really depends on the situation.

Sometimes, a zero-day vulnerability is already in your application that is deployed.
In that case, freezing deploys won't help because your app is already vulnerable.

That was the case with the Log4Shell vulnerability that was discovered in late 2021 but has been present in Log4j since 2013.
Luckily, there was a fix available within hours, but next time we might not be this lucky.

## Summary

As we already learned in [Day 14](day14.md), scanning your container images for vulnerabilities is important as it can give you valuable insights about
the security posture of your images.

Today we learned that it's even better to have it as part of your CI/CD pipeline and to enforce some basic rules about what vulnerabilities you have inside your images.

Finally, we discussed the steps we can take when we find a vulnerability.

Tomorrow we are going to take a look at container registries that enable this scanning out of the box and also at scanning other types of artifacts.
See you on [Day 22](day22.md).
