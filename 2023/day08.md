# Day 8: SAST Overview

Static Application Security Testing (SAST) is a method of evaluating the security of an application by analyzing the source code of the application without executing the code. SAST is also known as white-box testing as it involves testing the internal structure and workings of an application.

SAST is performed early in the software development lifecycle (SDLC) as it allows developers to identify and fix vulnerabilities before the application is deployed. This helps prevent security breaches and minimizes the risk of costly security incidents.

One of the primary benefits of SAST is that it can identify vulnerabilities that may not be detected by other testing methods such as dynamic testing or manual testing. This is because SAST analyzes the entire codebase and can identify vulnerabilities that may not be detectable by other testing methods.

There are several types of vulnerabilities that SAST can identify, including:

- **Input validation vulnerabilities**: These vulnerabilities occur when an application does not adequately validate user input, allowing attackers to input malicious code or data that can compromise the security of the application.
- **Cross-site scripting (XSS) vulnerabilities**: These vulnerabilities allow attackers to inject malicious scripts into web applications, allowing them to steal sensitive information or manipulate the application for their own gain.
- **Injection vulnerabilities**: These vulnerabilities allow attackers to inject malicious code or data into the application, allowing them to gain unauthorized access to sensitive information or execute unauthorized actions.
- **Unsafe functions and libraries**: These vulnerabilities occur when an application uses unsafe functions or libraries that can be exploited by attackers.
- **Security misconfigurations**: These vulnerabilities occur when an application is not properly configured, allowing attackers to gain access to sensitive information or execute unauthorized actions.

### SAST Tools (with free tier plan)

- **[SonarCloud](https://www.sonarsource.com/products/sonarcloud/)**: SonarCloud is a cloud-based code analysis service designed to detect code quality issues in 25+ different programming languages, continuously ensuring the maintainability, reliability and security of your code.
- **[Snyk](https://snyk.io/)**: Snyk is a platform allowing you to scan, prioritize, and fix security vulnerabilities in your own code, open source dependencies, container images, and Infrastructure as Code (IaC) configurations.
- **[Semgrep](https://semgrep.dev/)**: Semgrep is a fast, open source, static analysis engine for finding bugs, detecting dependency vulnerabilities, and enforcing code standards.

## How SAST Works?

SAST tools typically use a variety of techniques to analyze the sourced code, including pattern matching, rule-based analysis, and data flow analysis. 

Pattern matching involves looking for specific patterns in the code that may indicate a vulnerability, such as the use of a known vulnerable library or the execution of user input without proper sanitization. 

Rule-based analysis involves the use of a set of predefined rules to identify potential vulnerabilities, such as the use of weak cryptography or the lack of input validation. 

Data flow analysis involves tracking the flow of data through the application and identifying potential vulnerabilities that may arise as a result, such as the handling of sensitive data in an insecure manner.

## Consideration while using SAST Tools

1. It is important to ensure that the tool is properly configured and that it is being used in a way that is consistent with best practices. This may include setting the tool's sensitivity level to ensure that it is properly identifying vulnerabilities, as well as configuring the tool to ignore certain types of vulnerabilities that are known to be benign.
2. SAST tools are not a replacement for manual code review. While these tools can identify many potential vulnerabilities, they may not be able to identify all of them, and it is important for developers to manually review the code to ensure that it is secure.
3. SAST is just one aspect of a comprehensive application security program. While it can be an important tool for identifying potential vulnerabilities, it is not a replacement for other security measures, such as secure coding practices, testing in the production environment, and ongoing monitoring and maintenance.

### Challenges associated with SAST

- **False positives**: Automated SAST tools can sometimes identify potential vulnerabilities that are not actually vulnerabilities. This can lead to a large number of false positives that need to be manually reviewed, increasing the time and cost of the testing process.
- **Limited coverage**: SAST can only identify vulnerabilities in the source code that is analyzed. If an application uses external libraries or APIs, these may not be covered by the SAST process.
- **Code complexity**: SAST can be more challenging for larger codebases or codebases that are written in languages that are difficult to analyze.
- **Limited testing**: SAST does not execute the code and therefore cannot identify vulnerabilities that may only occur when the code is executed.

Despite these challenges, SAST is a valuable method of evaluating the security of an application and can help organizations prevent security breaches and minimize the risk of costly security incidents. By identifying and fixing vulnerabilities early in the SDLC, organizations can build more secure applications and improve the overall security of their systems.

### Resources

- [SAST- Static Analysis with lab by Practical DevSecOps](https://www.youtube.com/watch?v=h37zp5g5tO4)
- [SAST – All About Static Application Security Testing](https://www.mend.io/resources/blog/sast-static-application-security-testing/)
- [SAST Tools : 15 Top Free and Paid Tools](https://www.appsecsanta.com/sast-tools)

In the next part [Day 9](day09.md), we will discuss SonarCloud and integrate it with different CI/CD tools.
