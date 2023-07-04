## SSH & Web Server

Bu makale boyunca belirtildiği gibi, muhtemelen birçok uzaktaki Linux sunucusunu yönetiyorsunuz. Bu nedenle, uzak sunuculara olan bağlantınızın mümkün olduğunca güvenli olduğundan emin olmanız gerekecektir. Bu bölümde, herkesin bilmesi gereken SSH'nin temellerini ve uzak sistemlere güvenli bir tünel oluşturmanıza yardımcı olacak bazı konuları ele alacağız.

- SSH ile bağlantı kurma.
- Dosya transferi.
- Özel anahtarınızı oluşturma.

### SSH'a Giriş

- Shell.
- Bir Network protokolü.
- Güvenli iletişim sağlar.
- Herhangi bir Network hizmetini güvence altına alabilir.
- Genellikle komut satırına uzaktan erişim için kullanılır.

Önceki günleri takip ettiyseniz, SSH'yi zaten kullandığımızı hatırlarsınız. Ancak, bu Vagrant aracılığıyla yapılandırılmış ve otomatikleştirilmiş olduğundan, sadece `vagrant ssh` komutunu çalıştırdık ve uzak sanal makineye erişim elde ettik.

Eğer uzaktaki makine sistemimizde değil de, örneğin bulutta veya yalnızca İnternet üzerinden erişilebilen bir veri merkezinde çalışıyorsa, sistemini yönetmek için güvenli bir erişim yöntemine ihtiyacımız olur.

Bu nedenle, kötü niyetli kişiler tarafından herhangi bir şeyin ele geçirilemeyeceği bir güvenli bir tünel sağlayan SSH protokolünden yararlanacağız.

![](Images/Day18_Linux1.png)

Sunucu, her zaman çalışan ve belirli bir TCP portunda (22) dinleyen bir SSH sunucu servisine sahiptir.

Doğru kimlik bilgileri veya SSH anahtarıyla istemcimizi kullanarak sunucuya bağlanırsak, o sunuca erişim elde ederiz.

### Sistemimize bir Bridge Network Adaptoru ekliyoruz.

VirtualBox üzerindeki sanal makinemizle bunu simüle edebilmek için sanal makineye bir Bridge Network Adaptoru eklememiz gerekmektedir.

Sanal makinenin kapalı olduğundan emin olun, VirtualBox içindeki makineye sağ tıklayın ve "Settings" seçeneğini seçin. Açılan yeni pencerede, "Network" seçeneğini seçin.

![](Images/Day18_Linux2.png)


Şimdi makinenizi tekrar açın ve yerel makinenizde bir IP adresine sahip olacaksınız. Bununla ilgili olarak `IP addr` komutunu kullanarak doğrulama yapabilirsiniz.

### SSH sunucusunun çalıştığını doğrulama

SSH'nin zaten makinenize yapılandırıldığını biliyoruz çünkü vagrant ile kullanıyorduk, ancak bunu doğrulayabiliriz.

```shell
sudo systemctl status ssh
```

![](Images/Day18_Linux3.png)

Eğer sisteminizde SSH sunucusu yoksa, aşağıdaki komutu kullanarak yükleyebilirsiniz.

```shell
sudo apt install OpenSSH-server
```

Ardından, eğer bir güvenlik duvarı çalışıyorsa, SSH'nin güvenlik duvarı tarafından izin verildiğinden emin olmanız gerekmektedir. Örneğin, ufw adlı bir güvenlik duvarı kullanıyorsanız, aşağıdaki komutla protokolü etkinleştirmeniz gerekmektedir.

```shell
sudo ufw allow ssh
``` 
Sanal makinede, Vagrant ile otomatik olarak yapılandırdığımız için bu adıma gerek yok.

### Remote Access(Uzak Erisim) - SSH Parolası

Şimdi SSH sunucumuzun gelen bağlantı istekleri için 22 numaralı portu dinlemesini sağladık ve köprü ağını ekledik. Windows'ta, putty veya başka bir SSH istemcisi kullanarak ana makinede, SSH'ye bağlanmak için sistemimize bağlanabiliriz.

![](Images/Day18_Linux4.png)

Ardından, "Open" düğmesine tıklayın. Eğer bu IP adresi üzerinden ilk kez bu sisteme bağlanıyorsanız, aşağıdaki uyarıyı göreceksiniz. Bu sistemin bizim olduğunu biliyoruz, bu yüzden "Evet" seçeneğini seçebilirsiniz.

![](Images/Day18_Linux5.png)

Şimdi kullanıcı adımızı (vagrant) ve şifremizi (default şifre - vagrant) girmemiz istenir. Sonra SSH istemcimizin (Putty) sanal makineye bağlandığını göreceksiniz.

![](Images/Day18_Linux6.png)

Harika! Artık uzaktan istemcimizden sanal makinemize bağlandık ve sistemimizde komutlarımızı çalıştırabiliriz.

Başka bir Linux'tan ise doğrudan "ssh" komutunu kullanarak shelle erişebilirsiniz. Komut satırı aşağıdaki gibidir: "

```shell
ssh <nombre-usuario>@<IP-equipo-remoto>
```

Eğer sanal makinenin IP adresi 192.168.169.135 ise, aşağıdaki komutu kullanmamız gerekecektir:

```shell
ssh vagrant@192.168.169.135
```

Bize şifre soracak ve ardından SSH'nin otomatik sertifikasını kabul etmemizi isteyecektir.

### Acceso remoto - SSH Anahtarı

Gördüğümüz yöntem, sistemlere hızlı ve kolay erişim sağlamanın bir yoludur, ancak hala kullanıcı adı ve şifre üzerine dayanmaktadır. Bu durumda, kötü niyetli bir aktör bu bilgilere ve sistemizin genel adresine veya IP'sine erişebilirse, güvenliğimiz tehlikeye girebilir. Bu zayıf noktayı azaltmak için SSH anahtarlarını kullanabiliriz.

SSH anahtarları, güvenilir bir cihaz olduğundan emin olmak için hem istemciye hem de sunucuya ait bir çift anahtar sağlar.

Bir anahtar oluşturmak kolaydır. Yerel makinenizde (Windows veya Linux'ta) aşağıdaki komutu kullanarak sertifikaları oluşturabilirsiniz.

```shell
ssh-keygen -t ed25519
```

Daha fazla `ed25519` hakkında bilgi edinmek isterseniz arastırabilirsiniz [kriptografi](https://en.wikipedia.org/wiki/EdDSA#Ed25519). Heyecan verici bir konu.

![](Images/Day18_Linux7.png)

Bu noktada, SSH anahtarımızı Windows için `C:\sers\micha/.ssh/` veya Linux için `$HOME/.ssh` dizininde oluşturduk ve sakladık.

Ancak bu anahtarı Linux sanal makineyle ilişkilendirmek için kopyalamamız gerekiyor. Bunun için kabukta `ssh-copy-id vagrant@192.168.169.135` komutunu kullanabiliriz.

Windows'da Powershell'i kullanarak anahtarları oluşturabilirsiniz, ancak burada ssh-copy-id mevcut değildir. Windows'ta bunu yapmanın yolları vardır ve çevrimiçi küçük bir arama size alternatifleri gösterecektir, ancak biz Windows için kopyalamayı yapmak için [git bash](https://gitforwindows.org/) kullanacağız.

![](Images/Day18_Linux8.png)

Şimdi, SSH anahtarlarımızla şifre gerektirmeyen bir bağlantı kurup kurmadığımızı test etmek için Powershell'e dönebiliriz. Bu, daha önce bahsettiğimiz komutu Linux istemcisinde kullanarak yapabilirsiniz.

```shell
ssh vagrant@192.168.169.135
```

![](Images/Day18_Linux9.png)

Gerektiğinde bunu daha da güvenli hale getirebiliriz. Bir parola ifadesi kullanarak bunu sağlayabiliriz. Ayrıca, hiçbir şifre kullanılmamasını ve yalnızca SSH üzerinden anahtar çiftlerinin kabul edilmesini sağlayabiliriz. Bu, sanal makinenin aşağıdaki yapılandırma dosyasında yapılabilir.

```shell
sudo nano /etc/ssh/sshd_config
```

Yapılandırma dosyasında  `PasswordAuthentication yes` şeklinde bir satır bulacaksınız. Bu satırın başında `#` işareti olacak şekilde yorumlanmış olacak. Yorum işaretini kaldırın ve yes ifadesini no ile değiştirin. Ardından SSH hizmetini yeniden yüklemek için aşağıdaki komutu kullanmanız gerekecek:


```shell
sudo systemctl reload sshd
```

Ve güvenliği artırdık. KonsKonfıgurasyon seviyelerine kadar daha da iyileştirilebilir, ancak başlamak için yeterli olanları yaptık.

## Web Sunucu Yapılandırması

Linux oyun alanımız için Linux sanal makineye daha fazla özellik ekleyelim. Bir Apache web sunucusu ekleyerek yerel ağda görüntülenebilecek basit bir web sayfası barındırabiliriz. Unutmayın, bu web sayfası İnternet üzerinden erişilebilir olmayacak. Bu önerilmez, ancak yapılabilir. Ancak, burada bunu görmeyeceğiz. [Ipucu](https://www.noip.com/). 

Ayrıca bunu LAMP stack olarak da düşünebilirsiniz.

- **L**inux Ilsetim Sistemi
- **A**pache Web Server
- **M**ySQL Veritabanı
- **P**HP

### Apache2

Apache2, açık kaynaklı bir HTTP sunucusudur. Apache2'yi aşağıdaki komutu kullanarak kurabiliriz.

```shell
sudo apt-get install apache2
```

Apache2'nin başarıyla yüklendiğini doğrulamak için şu komutu çalıştırabilirsiniz:

```shell
sudo service apache2 restart
```

Ardından, SSH bölümünde gördüğümüz uzak makinenin ağ adresini kullanacağız. Bir tarayıcı açın ve adresi kontrol edin. Örnek olarak, adresimiz `http://192.168.169.135/`.

![](Images/Day18_Linux10.png)

### mySQL

MySQL, basit web sitemiz için verilerimizi depolayacağımız bir veritabanıdır. MySQL'i yüklemek için aşağıdaki komutu kullanmalıyız.

```shell
sudo apt-get install mysql-server
```

### PHP

PHP, sunucu tarafında çalışan bir programlama dilidir. MySQL veritabanıyla etkileşimde bulunmak için PHP'yi kullanacağız. Son kurulum adımı, PHP ve bağımlılıklarını aşağıdaki komutu kullanarak yüklemektir.

```shell
sudo apt-get install php libapache2-mod-php php-mysql
```

Yapmak istediğimiz ilk yapılandırma değişikliği, Apache'nin `index.html` yerine `index.php` kullanmasını sağlamaktır.

Aşağıdaki komutu kullanacağız.

```shell
sudo nano /etc/apache2/mods-enabled/dir.conf
``` 
Ve listelemede `index.php`'yi listenin ilk öğesine taşıyacağız.

![](Images/Day18_Linux11.png)

Apache2 hizmetini yeniden başlatıyoruz.

```shell
sudo systemctl restart apache2
```

Şimdi sistemimizin PHP için doğru şekilde yapılandırıldığını doğrulayalım. Aşağıdaki komutu kullanarak aşağıdaki dosyayı oluşturun, bu nano ile boş bir dosya açacaktır.

```shell
sudo nano /var/www/html/90Days.php
```

Daha sonra aşağıdaki metni belgeye kopyalayın ve öğrendiğimiz gibi, ctrl + x kullanarak çıkın ve dosyayı kaydedin.

```php
<?php
phpinfo();
?>
```

Şimdi tekrar Linux MV'nin IP adresine, URL'nin sonuna eklenen 90Days.php ile gezin: `http://192.168.169.135/90Days.php`. PHP doğru şekilde yapılandırılmışsa aşağıdakine benzer bir şey görmelisiniz. 

![](Images/Day18_Linux12.png)

### WordPress'in Kurulumu

Ardından, LAMP stack'imize WordPress kurulumu için bu rehberi takip ettim.

```shell
sudo mysql -u root -p
```

```mysql
CREATE DATABASE wordpressdb;

CREATE USER 'admin-user'@'localhost' IDENTIFIED BY 'password';

GRANT ALL PRIVILEGES ON wordpressdb.* TO 'admin-user'@'localhost';

FLUSH PRIVILEGES;

EXIT;
```

```shell
sudo apt install php-curl php-gd php-mbstring php-xml php-xmlrpc php-soap php-intl php-zip
```
```shell
sudo systemctl restart apache2
```

```shell
cd /var/www
```

```shell
sudo curl -O https://wordpress.org/latest.tar.gz
```

```shell
sudo tar -xvf latest.tar.gz
```

```shell
sudo rm latest.tar.gz
```
Bu noktada, "["Ubuntu üzerinde LAMP ile WordPress nasıl kurulur"](https://blog.ssdnodes.com/blog/how-to-install-wordpress-on-ubuntu-18-04-with-lamp-tutorial/)" adlı makalenin 4. adımındasınız ve WordPress dizini için doğru izinlerin yerinde olduğundan emin olmak için adımları izlemeniz gerekecektir.

Bu içsel bir kurulum olduğu için bu adımda "güvenlik anahtarları oluşturmak" gerekli değildir. Adım 5'e gidin, bu adımda Apache yapılandırmasını WordPress'e değiştirmeniz gerekmektedir.

Eğer her şey düzgün şekilde yapılandırılmışsa, dahili ağ adresiniz üzerinden WordPress kurulumunu çalıştırabilirsiniz.

## Kaynaklar

- [SSH GUI Client - Remmina](https://remmina.org/)
- [The Beginner's guide to SSH](https://www.youtube.com/watch?v=2QXkrLVsRmk)
- [Vim in 100 Seconds](https://www.youtube.com/watch?v=-txKSRn0qeA)
- [Vim tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I)
- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)
- [Webminal](https://www.webminal.org/) 
- [Vim Temel Rehberi](https://gitea.vergaracarmona.es/man-linux/Guia-VIM)

Gorusmek Uzere [Gun 19](day19.md).
