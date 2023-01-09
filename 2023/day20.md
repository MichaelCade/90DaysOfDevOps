# IAST and DAST in conjunction - lab time

1. As there are no open-source IAST implementation will use a commercial one with some free licenses. For this purpose, you will need 2 componenets:
IAST solution from here - https://github.com/rstatsinger/contrast-java-webgoat-docker . You need docker and docker-compose installed in mac or linux enviroment (this lab is tested on Mint). Please follow the README to create account in Contrast. 
2. For running the IAST there are few ways to do it- manually via a DAST scanner, ...
- Easiest way to do it is to use ZAP proxy. For this purpose install ZAP from here - https://www.zaproxy.org/download/
- Install zap-cli - https://github.com/Grunny/zap-cli
- Run ZAP proxy (from installed location, in Mint it is by default in /opt/zaproxy)
- Set env variables for ZAP_API_KEY and ZAP_PORT
- Run several commands with zap cli. For example: zap-cli quick-scan -s all --ajax-spider -r http://127.0.0.1:8080/WebGoat/login.mvc . You should see some results in contrast UI.
