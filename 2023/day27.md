# Network Vulnerability Scan

On [Day 25](day25.md) we learned that vulnerability scanning is the process of scanning a network or system to identify any existing security vulnerabilities.
We also learned that Network Vulnerability Scanning is a subset of Systems Vulnerability Scanning, e.g. we are only scanning the network part of our system.

Today we are going to dive deeper into what Network Vulnerability Scanning is and how we can do it.

## Network Vulnerability Scanning

**Network vulnerability scanning** is the process of identifying weaknesses on a network that is a potential target for exploitation by threat actors.

Once upon a time, before the cloud, network security was easy (sort of, good security is never easy).
You build a huge firewall around your data center, allow traffic only to the proper entrypoints and assume that everything that managed to get inside is legitimate.

This approach has one huge flaw - if an attacker managed to get through the wall, there are no more lines of defence to stop them.

Nowadays, such an approach would work even less.
With the cloud and microservices architecture, the actors in a network has grown exponentially.

This requires us to change our mindset and adopt new processes and tools in building secure systems.

One such process is **Network Vulnerability Scanning**.
The tool that does that is called **Network Vulnerability Scanner**.

## How does network vulnerability scanning work?

Vulnerability scanning software relies on a database of known vulnerabilities and automated tests for them.
A scanner would scan a wide range of devices and hosts on your networks, identifying the device type and operating system, and probing for relevant vulnerabilities.

A scan may be purely network-based, conducted from the wider internet (external scan) or from inside your local intranet (internal scan).
It may be a deep inspection that is possible when the scanner has been provided with credentials to authenticate itself as a legitimate user of the host or device.

## Vulnerability management

After a scan has been performed and has found vulnerabilities, the next step is to address them.
This is the vulnerability management phase.

A vulnerability could be marked as false positive, e.g. the scanner reported something that is not true.
It could be acknowledged and then assessed by the security team.

Many vulnerabilities can be addressed by patching, but not all.
A cost/benefit analysis should be part of the process because not all vulnerabilities are security risks in every environment, and there may be business reasons why you can’t install a given patch.
It would be useful if the scanner reports alternative means to remediate the vulnerability (e.g., disabling a service or blocking a port via firewall).

## Caveats

Similar to container image vulnerability scanning, network vulnerability scanning tests your system for _known_ vulnerabilities.
So it will not find anything that is not already reporter.

Also, it will not protect you from something like exposing your admin panel to the internet and using the default password.
(Although I would assume that some network scanner are smart enough to test for well-known endpoints that should not be exposed).

At the end of the day, it's up to you to know your system, and to know the way to test it, and protect it.
Tools only go so far.

## Network Scanners

Here is a list of network scanners that can be used for that purpose.

**NOTE:** The tools on this list are not free and open-source, but most of them have free trials, which you can use to evaluate them.

- [Intruder Network Vulnerability Scanner](https://www.intruder.io/network-vulnerability-scanner)
- [SecPod SanerNow Vulnerability Management](https://www.secpod.com/vulnerability-management/)
- [ManageEngine Vulnerability Manager Plus](https://www.manageengine.com/vulnerability-management/)
- [Domotz](https://www.domotz.com/features/network-security.php)
- [Microsoft Defender for Endpoint](https://www.microsoft.com/en-us/security/business/endpoint-security/microsoft-defender-endpoint)
- [Rapid7 InsightVM](https://www.rapid7.com/products/insightvm/)

## Summary

As with all the security processes we talked about in the previous day, network scanning is not a silver bullet.
Utilizing a network scanner would not make you secure if you are not taking care of the other aspects of systems security.

Also, using a tool like a network scanner does not mean that you don't need a security team.

Quite, the opposite, a good Secure SDLC starts with enabling the security team to run that kind of tool againts the system.
Then they would also be responsible for triaging the results and working with the revelant teams that need to fix the vulnerabilities.
That will be done by either patching up the system, closing a hole that is not necessary, or re-architecturing the system in a more secure manner.

## Resources

<https://www.comparitech.com/net-admin/free-network-vulnerability-scanners/>

<https://www.rapid7.com/solutions/network-vulnerability-scanner/>

See you on [Day 28](day28.md).
