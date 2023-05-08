# IAST (Interactive Application Security Testing)

IAST is a type of security testing tool that is designed to identify vulnerabilities in web applications and help developers fix them. It works by injecting a small agent into the application's runtime environment and monitoring its behaviour in real-time. This allows IAST tools to identify vulnerabilities as they occur, rather than relying on static analysis or simulated attacks. 

IAST works through software instrumentation, or the use of instruments to monitor an application as it runs and gather information about what it does and how it performs. IAST solutions instrument applications by deploying agents and sensors in running applications and continuously analyzing all application interactions initiated by manual tests, automated tests, or a combination of both to identify vulnerabilities in real time Instrumentation.
IAST agent is running inside the application and monitoring for known attack patterns. As it is part of the application, it can monitor traffic between different components (either as classic MVC deployments and in microservices deployment).

## For IAST to be used, there are few prerequisites.

- Application should be instrumented (inject the agent).
- Traffic should be generated - via manual or automated tests. Another possible approach is via DAST tools (OWASP ZAP can be used for example).

## Advantages

One of the main advantages of IAST tools is that they can provide detailed and accurate information about vulnerabilities and how to fix them. This can save developers a lot of time and effort, as they don't have to manually search for vulnerabilities or try to reproduce them in a testing environment. IAST tools can also identify vulnerabilities that might be missed by other testing methods, such as those that require user interaction or are triggered under certain conditions. Testing time depends on the tests used (as IAST is not a standalone system) and with faster tests (automated tests) can be included into CI/CD pipelines. It can be used to detect different kind of vulnerabilities and due to the nature of the tools (it looks for “real traffic only) false positives/negatives findings are relatively rear compared to other testing types.
IAST can be used in two flavours - as a typical testing tool and as real-time protection (it is called RAST in this case). Both work at the same principles and can be used together.

## There are several disadvantages of the technology as well:

- It is relatively new technology so there is not a lot of knowledge and experience both for the security teams and for the tools builders (open-source or commercial).
- The solution cannot be used alone - something (or someone) should generate traffic patterns. It is important that all possible endpoints are queried during the tests.
- Findings are based on traffic. This is especially true if used for testing alone - if there is no traffic to a portion of the app / site it would not be tested so no findings are going to be generated.
- Due to need of instrumentation of the app, it can be fairly complex, especially compared to the source scanning tools (SAST or SCA).

There are several different IAST tools available, each with its own features and capabilities.

## Some common features of IAST tools include:

- Real-time monitoring: IAST tools monitor the application's behaviour in real-time, allowing them to identify vulnerabilities as they occur.
- Vulnerability identification: IAST tools can identify a wide range of vulnerabilities, including injection attacks, cross-site scripting (XSS), and cross-site request forgery (CSRF).
- Remediation guidance: IAST tools often provide detailed information about how to fix identified vulnerabilities, including code snippets and recommendations for secure coding practices.
- Integration with other tools: IAST tools can often be integrated with other security testing tools, such as static code analysis or penetration testing tools, to provide a more comprehensive view of an application's security.

IAST tools can be a valuable addition to a developer's toolkit, as they can help identify and fix vulnerabilities in real-time, saving time and effort. If you are a developer and are interested in using an IAST tool, there are many options available, so it is important to research and compare different tools to find the one that best fits your needs.

## Tool example

There are almost no open-source tools on the market. Example is the commercial tool: Contrast Community Edition (CE) - Fully featured version for 1 app and up to 5 users (some Enterprise features disabled). Contrast CE supports Java and .NET only. 
Can be found here - https://www.contrastsecurity.com/contrast-community-edition

See you on [Day 20](day20.md).
