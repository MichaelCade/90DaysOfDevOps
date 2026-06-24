## DevOps & Linux

Linux ve DevOps, özelleştirme ve ölçeklenebilirlik üzerinde odaklanarak benzer kültür ve perspektiflere sahiptir. Bu iki Linux özelliği, özellikle DevOps için önemlidir.

Birçok teknoloji, özellikle yazılım geliştirme veya altyapı yönetimiyle ilgili olanlar, Linux'tan kaynaklanır.

Ayrıca, birçok açık kaynak proje, özellikle DevOps araçları, başlangıçtan itibaren Linux'ta çalışacak şekilde tasarlandı.

DevOps veya herhangi bir operasyon işlevi açısından bakıldığında, Linux ile karşılaşacaksınız. WinOps için bir yer var, ancak çoğu zaman Linux sunucularını yöneteceksiniz ve dağıtacaksınız. 🐧

Michael Cade, macOS veya Windows tabanlı bir masaüstüne sahip olmasına rağmen, Linux'u bir süredir günlük olarak kullanmaktadır. Ancak, şu anda bulunduğu Cloud Native rolüne başladığında, dizüstü bilgisayarını tamamen Linux tabanlı ve günlük olarak kullanılan bir işletim sistemine dönüştürdü. Hala çok platformlu olmayan iş uygulamaları ve Windows veya Mac'in kurumsal sahipliği nedeniyle Linux'ta çalışmayan ses ve video ekipmanı nedeniyle Windows'a ihtiyaç duyuyordu. Linux'u tam zamanlı olarak çalıştırarak, gelecek 7 gün boyunca göreceğimiz birçok konuyu daha iyi anlamak için tam anlamıyla bir Linux Masaüstü deneyimi elde etmek istedi.

Öte yandan, çevirmen, Linux işletim sistemlerini 10 yılı aşkın bir süredir kullanmaktadır ve yalnızca sosyal veya iş ortamı gereksinimleri nedeniyle zaman zaman Windows veya MacOS kullanmaktadır. Yalnızca birinci sınıf bir oyun oynamak veya yazılımı test etmek istediğimde Windows'u özledim. İlk durumda gerektiğinde kullanmak için hazır bir sabit diskim var; ikinci durumda ise sadece deneme amaçlı bir Windows 10 sanallaştırma ortamım bulunuyor. Şimdiye kadar, Windows 11 gerçekten başarısız oldu. Microsoft'un 10 sürümüyle gerçekten bir adım attığını ve operasyonel ve gizlilik ihtiyaçlarımı karşılamasa da önceki başarısızlıklarına kıyasla oldukça sağlam olduğunu söylemeliyim.


## Baslayalım

Buradaki amacım, okuyucunun yukarıda açıklanan adımları takip etmesini değil, önceki görüşleri bilmesinin ilginç olabileceğini göstermektir. Elbette daha kolay seçenekler mevcuttur; Microsoft veya Apple işletim sistemleri "çalıştır ve kullan" prensibine dayanır. Ancak, tam zamanlı olarak bir Linux dağıtımı kullanmaya karar vermek, işlerin nasıl çalıştığını daha hızlı öğrenmenizi, kabukla tanışmanızı ve en önemlisi, işletim sistemlerinin, ağın ve donanımın nasıl çalıştığını anlamanızı sağlar.

Bu 7 günün çoğunda, bir sanal makineyi Virtual Box'ta bir Linux dağıtımının masaüstü sürümünü dağıtacağız. Yöneteceğiniz Linux sunucularının çoğu muhtemelen kullanıcı arabirimi (GUI) olmayan sunuculardır ve her şey kabuk (CLI) temelindedir. Ancak, başlangıçta belirttiğim gibi, bu 90 gün boyunca ele alınan araçların çoğu Linux'ta başladı, bu yüzden Linux dağıtımlarının özellikle CLI olmak üzere çalıştırılmasına dalmanız ve öğrenme deneyimini de yaşamanız önerilir. Bu, Windows veya MacOS gibi diğer işletim sistemlerini daha iyi anlamanıza da yardımcı olacaktır.

Bu gönderinin geri kalanında, Virtual Box ortamında bir Ubuntu Masaüstü sanal makinesi elde etmeye odaklanacağız. İlerlemek için Virtual Box ve Ubuntu ISO sayfalarından indirme yapın.

Linux dağıtımlarını kullanmanın bir başka iyi nedeni de ücretsiz ve açık kaynak olmalarıdır ve hatta bazıları özgür yazılım olabilir. Şu anda en yaygın kullanılan dağıtım olduğu için Ubuntu'yu seçtik (akıllı telefonlar ve kurumsal sunucular hariç). Bu, öznel bir görüş olabilir ve yanlış olabilir, ancak CentOS ile son yaşananlar göz önüne alındığında Ubuntu'nun en üst sıraya yerleşmiş olması mümkündür.


## HashiCorp Vagrant Tanıtımı

Vagrant, sanal makinelerinizi yöneten bir CLI yardımcı programıdır. Vagrant'ı kullanarak vSphere, Hyper-V, VirtualBox ve Docker gibi farklı platformlarda sanal makineleri oluşturabilir ve sonlandırabilirsiniz. Diğer sağlayıcılar da mevcuttur, ancak biz VirtualBox ile devam edeceğiz, bu yüzden hazırsınız.

İlk olarak, Vagrant'ı bilgisayarınıza yüklemeniz gerekmektedir. İndirme sayfasına gittiğinizde, birçok işletim sistemi seçeneği görürsünüz. HashiCorp Vagrant sayfasından işletim sisteminize uygun sürümü indirebilirsiniz.

Ardından, VirtualBox'ı yüklemeniz gerekmektedir. Bu, birçok farklı işletim sisteminde kullanılabilir ve Windows, macOS veya Linux gibi farklı platformlarda kullanabilme avantajına sahip olmanızı sağlar.

Her iki kurulum da oldukça basittir ve her ikisi de büyük bir topluluğa sahiptir, bu yüzden herhangi bir sorun yaşarsanız, yardım almak için bu topluluklara başvurabilirsiniz. İlgili belgelendirme sayfalarını ziyaret ederek daha fazla bilgi edinebilirsiniz.

- [Virtual Box](https://www.virtualbox.org/wiki/Documentation)
- [Ubuntu](https://help.ubuntu.com/)
- [Vagrant](https://developer.hashicorp.com/vagrant/docs)

## VAGRANTFILE Olusturma

İlk makinemizi çalıştırmak için hazırız, çalışma istasyonunuzun terminalinde. Aşağıda Windows PowerShell örneğini görebilirsiniz. Projelerinizin klasörüne gidin ve orada VAGRANTFILE'ınızı bulacaksınız. Orada vagrant up komutunu yazabilir ve her şey yolundaysa aşağıdakine benzer bir çıktı göreceksiniz.

![](Images/Day14_Linux1.png)

Echemos un vistazo a ese VAGRANTFILE y veamos lo que estamos construyendo.

```
Vagrant.configure("2") do |config|

  config.vm.box = "chenhan/ubuntu-desktop-20.04"

  config.vm.provider :virtualbox do |v|

   v.memory  = 8096

   v.cpus    = 4

   v.customize ["modifyvm", :id, "--vram", "128"]

end

end
```
Bu çok basit bir VAGRANTFILE'dır. Belirli bir "box" kullanmak istediğimizi söylüyoruz, bu bir halka açık imaj veya aradığınız özel bir yapı olabilir. Halka açık olarak kullanılabilen birçok "box"un listesini Vagrant halka açık "box" kataloğunda bulabilirsiniz.

Bir sonraki satırda, belirli bir sağlayıcı kullanmak istediğimizi belirtiyoruz ve bu durumda `VirtualBox`'ı kullanıyoruz. Makinemizin belleğini `8GB` ve CPU sayısını 4 olarak tanımlıyoruz. Görsel sorunlarla karşılaşırsanız aşağıdaki satırı da eklemeniz gerekebilir. Bu, video belleğini istediğiniz şekilde ayarlar. Ben bunu `128MB`'a kadar artırırdım, ancak sisteminize bağlı olarak değişebilir.

```
v.customize ["modifyvm", :id, "--vram", ""]
```

Ayrıca bu Vagrant dosyasının bir kopyasını [Linux Klasörü](Linux/VAGRANTFILE) içine yerleştirmeniz gerekmektedir.

## Linux Desktopumuzu Oluşturmak

İlk makinemizi çalıştırmak için hazırız. İş istasyonumuzdaki terminalde, aşağıda Windows PowerShell'de bir örnek görebilirsiniz. Proje klasörünüze gidin ve orada VAGRANTFILE dosyasını bulacaksınız. Ardından `vagrant up` komutunu yazabilir ve her şey yolunda ise aşağıdaki gibi bir çıktı alırsınız.

![](Images/Day14_Linux2.png)


Burada eklememiz gereken bir şey daha var: sanal makinede ağ yapılandırması "NAT" olarak ayarlanacak. Bu aşamada NAT hakkında bilgi sahibi olmanıza gerek yok, çünkü Ağ oturumunda bununla ilgili tam bir oturum yapacağız. Ev ağınızda bir makine elde etmek için kolay bir yol olması ve ayrıca VirtualBox'ta varsayılan ağ modu olması nedeniyle tercih edilen ağ modudur. Daha fazla bilgiye VirtualBox belgelerinde bulabilirsiniz. [Dokuman Virtual Box](https://www.virtualbox.org/manual/ch06.html#network_nat)

`vagrant up` tamamlandıktan sonra, yeni sanal makinenin terminaline doğrudan erişmek için `vagrant ssh` komutunu kullanabiliriz.

![](Images/Day14_Linux3.png)

Burada, önümüzdeki günlerde çoğunlukla keşfedeceğimiz yerdir. Ayrıca, günlük olarak bunu çalıştırırken geliştirici çalışma istasyonunuz için bazı özelleştirmelerde de derinleşeceğiz ve hayatınızı çok daha kolay hale getireceğiz. Ve tabii ki, gerçekten DevOps'ta olmak için iyi bir terminala sahip olduğunuzda mısınız?

VirtualBox'ta, VM'nizi seçtiğinizde oturum açma ekranını görmelisiniz.

![](Images/Day14_Linux4.png)

Evet, bu noktaya kadar geldiyseniz ve "KULLANICI ADI VE ŞİFRE NE?" diye sorduysanız 🤔 

- Username = vagrant
- Password = vagrant

Yarın bazı komutları ve ne işe yaradıklarını göreceğiz. Terminal, her şeyi gerçekleştirmek için yer olacak ve zamanla dostunuz olacak.

## Kaynaklar

- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)
- [Webminal](https://www.webminal.org/) 

İlerledikçe birçok kaynak bulacaksınız ve Go kaynakları gibi, genellikle ÜCRETSİZ içerik içeren kaynakları tutmaya çalışacağım.

Gorusmek uzere [Gun 15](day15.md)
