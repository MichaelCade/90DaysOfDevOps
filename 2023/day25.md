# Systems Vulnerability Scanning

## What is systems vulnerability scanning?

Vulnerability scanning is the process of scanning a network or system to identify any existing security vulnerabilities.

It is a proactive measure used to detect any weaknesses that an attacker may exploit to gain unauthorised access to a system or network.

Vulnerability scanning can be either manual or automated.
It can involve scanning for known vulnerabilities, analysing the configuration of a system or network, or using an automated tool to detect any possible vulnerabilities.

## How do you perform a vulnerability scan?

A vulnerability scan is typically performed with specialised software that searches for known weaknesses and security issues in the system.

The scan typically looks for missing patches, known malware, open ports, weak passwords, and other security risks.

Once the scan is complete, the results are analysed to determine which areas of the system need to be addressed to improve its overall security.

## What are the types of vulnerability scans?

There are two main types of vulnerability scan: unauthenticated and authenticated.

Unauthenticated scans are conducted without any credentials and, as such, can only provide limited information about potential vulnerabilities.
This type of scan helps identify low-hanging fruit, such as unpatched systems or open ports.

Authenticated scans, on the other hand, are conducted with administrative credentials.
This allows the scanning tool to provide much more comprehensive information about potential vulnerabilities, including those that may not be easily exploitable.

In the next two days we are going to take a look at containers and network vulnerability scan, which are more specific subsets os system vulnerability scanning.

## Why are vulnerability scans important?

Vulnerabilities are widespread across organisations of all sizes.
New ones are discovered constantly or can be introduced due to system changes.

Criminal hackers use automated tools to identify and exploit known vulnerabilities and access unsecured systems, networks or data.

Exploiting vulnerabilities with automated tools is simple: attacks are cheap, easy to run and indiscriminate, so every Internet-facing organisation is at risk.

All it takes is one vulnerability for an attacker to access your network.

This is why applying patches to fix these security vulnerabilities is essential.
Updating your software, firmware and operating systems to the newest versions will help protect your organisation from potential vulnerabilities.

Worse, most intrusions are not discovered until it is too late. According to the global median, dwell time between the start of a cyber intrusion and its identification is 24 days.

## What does a vulnerability scan test?

Automated vulnerability scanning tools scan for open ports and detect common services running on those ports.
They identify any configuration issues or other vulnerabilities on those services and look at whether best practice is being followed, such as using TLSv1.2 or higher and strong cipher suites.

A vulnerability scanning report is then generated to highlight the items that have been identified.
By acting on these findings, an organisation can improve its security posture.

## Who conducts vulnerability scans?

IT departments usually undertake vulnerability scanning if they have the expertise and software to do so, or they can call on a third-party security service provider.

Vulnerability scans are also performed by attackers who scour the Internet to find entry points into systems and networks.

Many companies have bug bountry programs, that allow enthical hackers to report vulnerabilities and gain money for that.
Usually the bug bountry programs have boundaries, e.g. they define what is allowed and what is not.

Participating in big bounty programs must be done resposibly.
Hacking is a crime, and if you are caugh you cannot just claim that you did it for good, or that you were not going to exploit your findings.

## How often should you conduct a vulnerability scan?

Vulnerability scans should be performed regularly so you can detect new vulnerabilities quickly and take appropriate action.

This will help identify your security weaknesses and the extent to which you are open to attack.

## Penetration testing

Penetration testing is the next step after vulnerability scanning.
In penetration testing professional ethical hackers combine the results of automated scans with their expertise to reveal vulnerabilities that may not be identified by scans alone.

Penetration testers will also consider your environment (a significant factor in determining vulnerabilities’ true severity) and upgrade or downgrade the score as appropriate.

A scan can detect something that is vulnerability, but it cannot be actively exploited, because of the way it is incorporated into our system.
This makes the vulnerability a low priority one, because why fix something that presents no danger to you.

If an issue comes up in penetration testing then that means that this issue is exploitable, and probably a high priority - in the penetation testers managed to exploit it, so will the hackers.
See you on [Day 26](day26.md).
