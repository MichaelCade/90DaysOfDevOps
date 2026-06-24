# DAST
DAST, or Dynamic Application Security Testing, is a technique that is used to evaluate the security of an application by simulating attacks from external sources.
Idea is to automate as much as possible black-box penetration testing. 
It can be used for acquiring the low-hanging fruits so a real human’s time will be spared and additionally for generating traffic to other security tools (e.g. IAST). 

Nevertheless, It is an essential component of the SSDLC, as it helps organizations uncover potential vulnerabilities early in the development process, before the application is deployed to production. By conducting DAST testing, organizations can prevent security incidents and protect their data and assets from being compromised by attackers.

## Tools

There are various open-source tools available for conducting DAST, such as ZAP, Burp Suite, and Arachni. These tools can simulate different types of attacks on the application, such as SQL injection, cross-site scripting, and other common vulnerabilities. For example, if an application is vulnerable to SQL injection, a DAST tool can send a malicious SQL query to the application, such as ' OR 1=1 --, and evaluate its response to determine if it is vulnerable. If the application is vulnerable, it may return all records from the database, indicating that the SQL injection attack was successful.
As some of the tests could be quite invasive (for example it may include ‘DROP TABLE’ or something similar) or at least put a good amount of test data into the databases or even DOS the app, 
__DAST tools should never run against a production environment!!!__ 
All tools have the possibility for authentication into the application and this could lead to production credentials compromise. Also when run authenticated scans against the testing environment, use suitable roles (if RBAC model exists, for the application, of course), e.g. DAST shouldn’t use role that have the possibility to delete or modify other users because this way the whole environment can became unusable.
As with other testing methodologies it is necessary to analyze the scope, so not unneeded targets are scanned. 

## Usage
Common error is scanning compensating security controls (e.g. WAF) instead of the real application. DAST is in its core an application security testing tool and should be used against actual applications, not against security mitigations. As it uses pretty standardized attacks, external controls can block the attacking traffic and this way to cover potentially exploitable flows (as per definition adversary would be able to eventually bypass such measures)
Actual scans are quite slow, so sometimes they should be run outside of the DevOps pipeline. Good example is running them nightly or during the weekend. Some of the simple tools (zap / arachny, …) could be used into pipelines but often, due to the nature of the scan can slow down the whole development process.
Once the DAST testing is complete, the results are analyzed to identify any vulnerabilities that were discovered. The organization can then take appropriate remediation steps to address the vulnerabilities and improve the overall security of the application. This may involve fixing the underlying code, implementing additional security controls, such as input validation and filtering, or both.
In conclusion, the use of DAST in the SSDLC is essential for ensuring the security of an application. By conducting DAST testing and identifying vulnerabilities early in the development process, organizations can prevent security incidents and protect their assets from potential threats. Open-source tools, such as ZAP, Burp Suite, and Arachni, can be used to conduct DAST testing and help organizations improve their overall security posture.
As with all other tools part of DevSecOps pipeline DAST should not be the only scanner in place and as with all others, it is not a substitute for penetration test and good development practices.

## Some useful links and open-source tools:
- https://github.com/zaproxy/zaproxy
- https://www.arachni-scanner.com/
- https://owasp.org/www-project-devsecops-guideline/latest/02b-Dynamic-Application-Security-Testing 

See you on [Day 19](day19.md).
