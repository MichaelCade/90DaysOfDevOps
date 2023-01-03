# Day 7: Secure Coding Overview

Secure coding is the practice of writing software in a way that ensures the security of the system and the data it processes. It involves designing, coding, and testing software with security in mind to prevent vulnerabilities and protect against potential attacks.

There are several key principles of secure coding that developers should follow:

1. Input validation: It is important to validate all user input to ensure that it is in the expected format and does not contain any malicious code or unexpected characters. This can be achieved through the use of regular expressions, data type checks, and other validation techniques.
2. Output encoding: Output data should be properly encoded to prevent any potential injection attacks. For example, HTML output should be properly escaped to prevent cross-site scripting (XSS) attacks, and SQL queries should be parameterized to prevent SQL injection attacks.
3. Access control: Access control involves restricting access to resources or data to only those users who are authorized to access them. This can include implementing authentication and authorization protocols, as well as enforcing least privilege principles to ensure that users have only the access rights they need to perform their job duties.
4. Error handling: Error handling is the process of properly handling errors and exceptions that may occur during the execution of a program. This can include logging errors, displaying appropriate messages to users, and mitigating the impact of errors on system security.
5. Cryptography: Cryptography should be used to protect sensitive data and communications, such as passwords, financial transactions, and sensitive documents. This can be achieved through the use of encryption algorithms and secure key management practices.
6. Threat Modeling: Document, locate, address, and validate are the four steps to threat modeling. To securely code, you need to examine your software for areas susceptible to increased threats of attack. Threat modeling is a multi-stage process that should be integrated into the software lifecycle from development, testing, and production.
7. Secure storage: Secure storage involves properly storing and handling sensitive data, such as passwords and personal information, to prevent unauthorized access or tampering. This can include using encryption, hashing, and other security measures to protect data at rest and in transit.
8. Secure architecture: Secure architecture is the foundation of a secure system. This includes designing systems with security in mind, using secure frameworks and libraries, and following secure design patterns.

There are several tools and techniques that can be used to help ensure that code is secure, including Static Application Security Testing (SAST), Software Composition Analysis (SCA), and Secure Code Review.

### Static Application Security Testing (SAST)

SAST is a method of testing software code for security vulnerabilities during the development phase. It involves analyzing the source code of a program without executing it, looking for vulnerabilities such as injection attacks, cross-site scripting (XSS), and other common security issues. SAST tools can be integrated into the software development process to provide ongoing feedback and alerts about potential vulnerabilities as the code is being written.

### Software Composition Analysis (SCA)

SCA is a method of analyzing the third-party components and libraries that are used in a software application. It helps to identify any vulnerabilities or security risks that may be present in these components, and can alert developers to the need to update or replace them. SCA can be performed manually or with the use of automated tools.

### Secure Code Reviews

Secure Code Review is a process of reviewing software code with the goal of identifying and addressing potential security vulnerabilities. It is typically performed by a team of security experts who are familiar with common coding practices and security best practices. Secure Code Review can be done manually or with the use of automated tools, and may involve a combination of SAST and SCA techniques.

In summary, Overall, secure coding is a crucial practice that helps protect software and its users from security vulnerabilities and attacks. By following best practices and keeping software up to date, developers can help ensure that their software is as secure as possible.

### Resources

- [Secure Coding Best Practices | OWASP Top 10 Proactive Control](https://www.youtube.com/watch?v=8m1N2t-WANc)

- [Secure coding practices every developer should know](https://snyk.io/learn/secure-coding-practices/)

- [10 Secure Coding Practices You Can Implement Now](https://codesigningstore.com/secure-coding-practices-to-implement)

- [Secure Coding Guidelines And Best Practices For Developers](https://www.softwaretestinghelp.com/guidelines-for-secure-coding/)

In the next part [Day 8](day08.md), we will discuss Static Application Security Testing (SAST) in more detail.
