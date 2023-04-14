---
title: '#90DaysOfDevOps - SSH & Web Server - Day 18'
published: false
description: 90DaysOfDevOps - SSH & Web Server
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048733
---

## SSH 및 웹 서버

앞서 언급한 바와 같이, 여러분은 다수의 원격 Linux 서버를 관리할 가능성이 높습니다. 그러므로 이러한 원격 서버와의 연결이 안전한지 확인하는 것이 중요합니다. 본 섹션에서는 원격 시스템과의 보안 터널 구축을 돕기 위해, 모든 사용자가 알아야 할 SSH의 주요 기본 개념들을 살펴볼 것입니다.

- SSH로 연결 설정하기
- 파일 전송하기
- 개인 키 만들기

### SSH 소개

- 보안 shell
- 네트워킹 프로토콜
- 보안 통신 허용
- 모든 네트워크 서비스를 보호
- 일반적으로 원격 command line 액세스에 사용

사실 우린 이미 vagrant 구성과 자동화를 통해 SSH를 사용하고 있었습니다. `vagrant ssh`만 실행해도 원격 가상 머신에 접근할 수 있었습니다.

원격 머신이 워크스테이션과 동일한 시스템이 아니라 원격 위치, 클라우드 기반 시스템이거나 인터넷을 통해서만 접근할 수 있는 데이터 센터에서 실행될 경우, 시스템을 안전하게 관리하기 위한 접근 방법이 필요합니다.

SSH는 클라이언트와 서버 간의 보안 터널을 제공하여 공격자가 정보를 가로챌 수 없게 합니다.

![](/2022/Days/Images/Day18_Linux1.png)

서버에는 항상 특정한 TCP 포트(22)에서 실행되어 대기 중인 서버 측 SSH 서비스가 존재합니다.

올바른 자격증명이나 SSH 키를 가진 클라이언트를 사용하면 해당 서버에 접근할 수 있습니다.

### 시스템에 브리지 네트워크 어댑터 추가하기

현재의 virtual box VM과 함께 사용하려면, 우리 시스템에 브리지 네트워크 어댑터를 추가해야 합니다.

가상 머신을 종료한 다음, Virtual Box 내의 머신을 마우스 우클릭 후 Settings를 선택합니다. 새 창에서 Network를 선택하세요.

![](/2022/Days/Images/Day18_Linux2.png)

이제 가상 머신을 다시 시작하면 로컬 머신에서 IP 주소를 갖게 됩니다. `IP addr` 명령어로 확인할 수 있습니다.

### SSH 서버 실행 확인

우리는 vagrant를 사용해왔기 때문에 SSH가 우리 머신에 이미 설정되어 있다는 것을 알고 있지만, 다음 명령어를 실행하여 확인할 수도 있습니다.

`sudo systemctl status ssh`

![](/2022/Days/Images/Day18_Linux3.png)

시스템에 SSH 서버가 없으면 `sudo apt install OpenSSH-server` 명령어를 사용하여 설치할 수 있습니다.

방화벽이 실행 중인 경우 SSH가 허용되도록 하려면 `sudo ufw allow ssh`를 사용할 수 있습니다. 하지만 vagrant 프로비저닝을 사용하여 자동화했으므로 이 작업은 필요하지 않습니다.

### 원격 접속 - SSH password

이제 SSH 서버가 포트 22에서 들어오는 연결 요청을 수신하도록 설정했고 브리지 네트워킹을 추가했으므로, 로컬 머신의 putty 또는 SSH 클라이언트를 사용하여 SSH로 시스템에 연결할 수 있습니다.

[# PuTTy installation Guide](https://www.cuit.columbia.edu/putty)

![](/2022/Days/Images/Day18_Linux4.png)

Open을 클릭하세요. 이 IP 주소를 통해 처음으로 이 시스템에 연결하는 경우 경고가 표시됩니다. 이 시스템이 우리 것임을 알고 있으므로 Yes를 선택할 수 있습니다.

![](/2022/Days/Images/Day18_Linux5.png)

username(vagrant)과 password(기본 - vagrant)를 입력하라는 메시지가 표시됩니다. 아래에서 SSH 클라이언트(PuTTY)를 사용하여 username과 password로 머신에 연결하는 것을 확인할 수 있습니다.

![](/2022/Days/Images/Day18_Linux6.png)

이제부터 우리의 원격 클라이언트는 VM에 연결되어 있으며, 시스템에서 명령어를 실행할 수 있습니다.

### 원격 접속 - SSH 키

앞서 설명한 방법으로 시스템에 쉽게 액세스할 수 있지만, username과 password에 의존하고 있습니다. 악의적인 사용자가 이 정보와 시스템의 공용 주소 또는 IP를 알게 되면 쉽게 침입할 수 있습니다. 이 때문에 SSH 키가 선호됩니다.

SSH 키를 사용하면, 클라이언트와 서버 모두 이 장치가 신뢰할 수 있는 것임을 알 수 있는 키 쌍을 제공합니다.

키를 생성하는 것은 쉽습니다. 로컬 머신(Windows)에서 다음 명령어를 실행하면 됩니다. 시스템에 SSH 클라이언트가 설치되어 있다면 이 명령어가 작동할 것입니다.

`ssh-keygen -t ed25519`

여기에서 `ed25519`가 무엇인지에 대해 자세히 설명하지는 않겠지만, 암호학에 관심이 있다면 [cryptography](https://en.wikipedia.org/wiki/EdDSA#Ed25519)를 참고해보세요.

![](/2022/Days/Images/Day18_Linux7.png)

생성된 SSH 키는 `C:\Users\micha/.ssh/`에 저장되어 있습니다.

하지만 이 키로 Linux VM에 연결하려면 키를 복사해야 합니다. `ssh-copy-id vagrant@192.168.169.135`를 사용하여 이 작업을 수행할 수 있습니다.

Windows 클라이언트에서 Powershell을 사용하여 키를 생성했지만, Windows에는 `ssh-copy-id`를 실행할 수 없습니다. Windows에서 이 작업을 수행하는 방법이 있으며, 온라인에서 간단한 검색을 통해 대체 방법을 찾을 수 있습니다. 여기서는 Windows 머신에서 git bash를 사용하여 복사하는 것을 보여 드리겠습니다.

![](/2022/Days/Images/Day18_Linux8.png)

이제 Powershell로 돌아와 password 없이 SSH 키로 연결이 작동하는지 테스트할 수 있습니다.

`ssh vagrant@192.168.169.135`

![](/2022/Days/Images/Day18_Linux9.png)

필요한 경우 passphrase를 사용하여 더욱 안전하게 만들 수 있습니다. 또한 password 없이 키 쌍만을 사용하여 SSH를 허용하도록 설정할 수 있습니다. 이렇게 하려면 다음 설정 파일에서 수정하면 됩니다.

`sudo nano /etc/ssh/sshd_config`

여기에 `#` 주석 처리가 되어있는 `PasswordAuthentication yes`라는 줄이 있고, 이를 주석 처리 해제하고 yes를 no로 변경해야 합니다. 그런 다음 `sudo systemctl reload sshd`를 실행해서 SSH 서비스를 다시 로드해야 합니다.

## 웹 서버 설정하기

위에서 설명한 SSH와는 별개로, 이번에는 웹 서버를 구축하는 방법을 다루고자 합니다. 처음에는 조금 어려워 보일 수 있지만, 실제로는 그렇지 않습니다.

리눅스 가상 머신을 이미 구축해 두었다고 가정하고, 이 가상 머신에 apache 웹 서버를 추가하여 내부 네트워크에서 접근할 수 있는 간단한 웹 사이트를 호스팅하려고 합니다. 인터넷에서 접근할 수 있는 웹 사이트는 아니며, 그에 대한 내용은 여기서 다루지 않습니다.

이 과정은 LAMP 스택이라고도 불립니다.

- **L**inux 운영 체제
- **A**pache 웹 서버
- **M**ySQL 데이터베이스
- **P**HP

### Apache2

Apache2는 오픈 소스 HTTP 서버입니다. 아래 명령어를 사용하여 Apache2를 설치할 수 있습니다.

`sudo apt-get install apache2`

Apache2가 정상적으로 설치되었는지 확인하려면 `sudo service apache2 restart`를 실행합니다.

SSH 통합 네트워크 주소를 사용하여 브라우저에서 해당 주소로 이동합니다. 예를 들면 `http://192.168.169.135/`와 같습니다.

![](/2022/Days/Images/Day18_Linux10.png)

### MySQL

MySQL은 웹 사이트의 데이터를 저장하는 데이터베이스입니다. `sudo apt-get install mysql-server` 명령어를 사용하여 MySQL을 설치합니다.

### PHP

PHP는 서버 측 스크립팅 언어로, MySQL 데이터베이스와 상호 작용하기 위해 사용합니다. PHP 및 종속성을 설치하려면 `sudo apt-get install php libapache2-mod-php php-mysql`을 사용합니다.

기본적으로 Apache는 index.html을 사용하지만, 대신 index.php를 사용하도록 설정하려고 합니다.

`sudo nano /etc/apache2/mods-enabled/dir.conf`를 사용하여 index.php를 목록의 첫 번째 항목으로 이동시킵니다.

![](/2022/Days/Images/Day18_Linux11.png)

Apache2 서비스를 다시 시작합니다. `sudo systemctl restart apache2`

이제 PHP가 올바르게 구성되었는지 확인해 보겠습니다. 다음 명령어를 사용하여 새 파일을 생성합니다.

`sudo nano /var/www/html/90Days.php`

아래 내용을 복사한 후 Ctrl + X를 눌러 종료하고 파일을 저장합니다.

```php
<?php
phpinfo();
?>
```

다음 URL로 이동하여 PHP가 올바르게 구성되었는지 확인합니다. `http://192.168.169.135/90Days.php` 이 주소로 이동하면 PHP가 올바르게 설정되었음을 확인할 수 있는 화면이 표시됩니다.

![](/2022/Days/Images/Day18_Linux12.png)

### 워드프레스 설치

LAMP 스택에 워드프레스를 설치하기 위해 [How to install WordPress on Ubuntu with LAMP](https://blog.ssdnodes.com/blog/how-to-install-wordpress-on-ubuntu-18-04-with-lamp-tutorial/)을 참고하였습니다. 이 글에 나와 있는 명령어들은 아래와 같습니다.

`sudo mysql -u root -p`

`CREATE DATABASE wordpressdb;`

`CREATE USER 'admin-user'@'localhost' IDENTIFIED BY 'password';`

`GRANT ALL PRIVILEGES ON wordpressdb.* TO 'admin-user'@'localhost';`

`FLUSH PRIVILEGES;`

`EXIT;`

`sudo apt install php-curl php-gd php-mbstring php-xml php-xmlrpc php-soap php-intl php-zip`

`sudo systemctl restart apache2`

`cd /var/www`

`sudo curl -O https://wordpress.org/latest.tar.gz`

`sudo tar -xvf latest.tar.gz`

`sudo rm latest.tar.gz`

위 링크의 Step 4에 도달했습니다. 워드프레스 디렉토리에 대한 모든 권한이 올바르게 설정되어 있는지 확인해야 합니다.

내부 네트워크 전용이므로 이 단계에서 "generate security keys"를 할 필요가 없습니다. apache 구성을 워드프레스로 변경하는 Step 5로 이동합니다.

모든 설정이 올바르게 구성되어 있다면, 내부 네트워크 주소를 통해 워드프레스 설치를 진행할 수 있습니다.

## 자료

- [Client SSH GUI - Remmina](https://remmina.org/)
- [The Beginner's guide to SSH](https://www.youtube.com/watch?v=2QXkrLVsRmk)
- [Vim in 100 Seconds](https://www.youtube.com/watch?v=-txKSRn0qeA)
- [Vim tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I)
- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

[Day 19](day19.md)에서 봐요!
