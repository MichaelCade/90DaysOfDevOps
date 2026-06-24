## Containers(Konteynerlar)

Şimdi yeni bir bölüme başlıyoruz ve bu bölüm, özellikle Docker'a odaklanacak ve Konteynerler hakkında daha fazla anlayış kazanmak için bazı temel konulara odaklanacak.

Ayrıca, bu bölümde ve ilerleyen zorlukların sonraki bölümlerinde kullanabileceğimiz bir konteyner oluşturmak için burada pratik yapmaya çalışacağım.

Her zamanki gibi, bu ilk gönderi, nasıl buraya geldiğimiz ve bunun ne anlama geldiği hakkında büyük resme odaklanacaktır.

#Platformların ve uygulama geliştirmenin tarihi
#Sanallaştırma ve Konteynerleştirme hakkında konuşmak istiyor muyuz?

### Neden başka bir uygulama çalıştırma yöntemi?

Bakmamız gereken ilk şey, neden yazılım veya uygulamalarımızı çalıştırmak için başka bir yönteme ihtiyacımız var? İşte sadece seçeneğin harika olmasıyla ilgili, uygulamalarımızı birçok farklı şekilde çalıştırabiliriz. Fiziksel donanıma sahip bir işletim sistemi ve tek bir uygulamanın dağıtıldığı durumları görebiliriz veya uygulamamızı çalıştıran sanal makine veya bulut tabanlı IaaS örneklerini görebiliriz, bunlar daha sonra bir veritabanına entegre olabilir veya kamu bulutunda bir PaaS teklifi olarak sunulabilir. Veya uygulamalarımızı konteynerlerde çalıştırırken görebiliriz.

Yukarıdaki seçeneklerin hiçbiri yanlış veya doğru değil, ancak her birinin var olma nedenleri vardır ve ayrıca bu seçeneklerden hiçbirinin ortadan kalkacağına inanmıyorum. Konteynerler vs. Sanal Makineler gibi konuları ele alan birçok içerik gördüm ve gerçekten bir tartışma olmamalı çünkü bu daha çok elma ile armut karşılaştırması gibi bir şey, her ikisi de meyve (uygulamalarımızı çalıştırmanın yolları) olsa da aynı değil.

Ayrıca, eğer başlıyorsanız ve bir uygulama geliştiriyorsanız, konteynerlere doğru yönelmeniz gerektiğini söylerim, çünkü bu konulara daha sonra değineceğiz, ancak bu verimlilik, hız ve boyutla ilgilidir. Ancak bunun bir bedeli vardır, konteynerler hakkında hiçbir fikriniz yoksa, nedenini anlamak ve bu zihniyete girmek için bir öğrenme eğrisi olacaktır. Eğer uygulamalarınızı belirli bir şekilde geliştirdiyseniz veya yeşil saha bir ortamda değilseniz, konteynerleri düşünmeden önce ele almanız gereken daha fazla zorluk yaşayabilirsiniz.

Belirli bir yazılım parçasını indirme konusunda birçok farklı seçeneğimiz var, kullanabileceğimiz bir dizi farklı işletim sistemi var. Ve uygulamalarımızı nasıl kurmamız gerektiği için belirli talimatlar mevcut.

![](Images/Day42_Containers1.png)

Son zamanlarda giderek daha fazla gözlemliyorum ki bir zamanlar tam bir sunucu işletim sistemine, bir sanal makineye, fiziksel bir sunucuya veya bulut örneğine ihtiyaç duyduğumuz uygulamalar, şimdi yazılımlarının konteyner tabanlı sürümlerini yayınlıyorlar. Bu durumu ilginç buluyorum çünkü bu, konteynerlerin ve ardından herkese değil sadece uygulama geliştiricilerine odaklanmayı açan Kubernetes dünyasını herkese açıyor gibi görünüyor.

![](Images/Day42_Containers2.png)

Muhtemelen daha önce söylediğim gibi, cevabın ne olduğunu söyleyeceğimi savunmayacağım, soru nedir! Ama uygulamalarımızı dağıttığımızda farkında olmamız gereken başka bir seçenek olduğunu tartışmak istiyorum.

![](Images/Day42_Containers4.png)

Uzun bir süredir konteyner teknolojimiz var, o zaman neden son 10 yılda, son 5 yılda daha da popüler hale geldiğini söyleyebilirim. Onlarca yıldır konteynerlerimiz var. Bunun nedeni, sadece konteyner teknolojisine sahip olsak bile, yazılım yönetimi ile yaşadığımız aynı sorunlara sahip olmamızdır.

Docker'ı bir araç olarak düşünürsek, popüler hale gelmesinin nedeni, kullanımı kolay ve bulması kolay görüntüler ekosistemidir. Sistemlerinize almak ve çalıştırmak kolaydır. Bunun büyük bir kısmı, yazılım ile karşılaştığımız tüm bu farklı zorlukların tamamında tutarlılıktır. MongoDB veya nodeJS olsun, her ikisini de çalıştırmak için izlenen süreç aynı olacaktır. Her ikisini durdurma süreci de aynıdır. Tüm bu sorunlar hala var olacak, ancak iyi bir konteyner ve görüntü teknolojisini bir araya getirdiğimizde, artık tüm bu farklı sorunları ele almamıza yardımcı olacak tek bir araç setine sahibiz. Bu sorunların bazıları aşağıda listelenmiştir:

- İlk olarak, internet üzerinde yazılım bulmamız gerekiyor.
- Ardından bu yazılımı indirmemiz gerekiyor. 
- Kaynağa güveniyor muyuz?
- Sonra bir lisansa mı ihtiyacımız var? Hangi lisans? 
- Farklı platformlarla uyumlu mu?
- Paket nedir? İkili mi? Yürütülebilir mi? Paket yöneticisi mi? 
- Yazılımı nasıl yapılandırırız?
- Bağımlılıklar? Genel indirme işini örttü mü, yoksa bunları da mı gerektiriyor?
- Bağımlılıkların bağımlılıkları?
- Uygulamayı nasıl başlatırız? 
- Uygulamayı nasıl durdururuz?
- Otomatik yeniden başlar mı? 
- Başlangıçta otomatik başlar mı?? 
- Kaynak çatışmaları? 
- Çakışan kütüphaneler? 
- Port cakısmaları
- Yazılım için güvenlik?
- Yazılım güncellemeleri?
- Yazılımı nasıl kaldırabilirim?? 

Yukarıdakileri yazılımın karmaşıklığının üç alanına bölebiliriz ve konteynerler ve görüntüler bu konularda yardımcı olurlar

| Distribution | Installation | Operation          |
| ------------ | ------------ | -----------------  |
| Find         | Install      | Start              |
| Download     | Configuration| Security           |
| License      | Uninstall    | Ports              |
| Package      | Dependencies | Resource Conflicts |
| Trust        | Platform     | Auto-Restart       |
| Find         | Libraries    | Updates            |

Konteynerler ve görüntüler, muhtemelen diğer yazılım ve uygulamalarda yaşadığımız bazı zorlukları aşmamıza yardımcı olacak. 

Yüksek seviyede, kurulumu ve işletmeyi aynı liste taşıyabiliriz, Görüntüler bize bir dağıtım açısından yardımcı olurken, konteynerler kurulum ve işletme konularına yardımcı olur.

Tahminen harika ve heyecan verici gelebilir, ancak hala bir konteynerin ne olduğunu anlamamız gerekiyor ve şimdi görüntülerden bahsettim, bu yüzden bir sonraki bölümde bu konuları ele alalım.

Yazılım geliştirme için Konteynerler hakkında konuşurken sıkça gördüğünüz bir diğer şey, gemi konteynerleri ile birlikte kullanılan bir benzetme, denizlerin üzerinde büyük gemiler kullanılarak çeşitli malların taşınmasında gemi konteynerlerinin kullanıldığı bir gerçektir.

![](Images/Day42_Containers5.png)

Bu, konteynerler konusundaki konumuzla ne ilgisi var? Yazılım geliştiricilerinin yazdığı kodları düşünün, bu belirli kodu bir makineden başka bir makineye nasıl taşıyabiliriz?

Önce yazılım dağıtımı, kurulum ve işletmeye dokunduğumuz şeyi düşünürsek ve şimdi bunu bir çevresel görsel olarak inşa etmeye başlarsak, donanım ve bir işletim sistemi ile birden fazla uygulama çalıştıracağınız bir ortamınız var. Örneğin, nodejs belirli bağımlılıklara ve belirli kütüphanelere ihtiyaç duyar. Daha sonra MySQL'i yüklemek istiyorsanız, gerekli kütüphaneler ve bağımlılıklara ihtiyaç duyar. Her yazılım uygulamasının kendi kütüphanesi ve bağımlılığı olacaktır. Özel kütüphaneler ve bağımlılıkların çakıştığı ve sorunlara yol açtığı özel bir durumumuz olmadıkça, herhangi bir uygulama arasında çakışma olasılığının daha fazla olduğu daha fazla uygulama olacaktır. Bununla birlikte, bu yalnızca her şeyin yazılım uygulamalarınızı düzelten bir dağıtım olduğu bir durum değildir, yazılım uygulamalarınız güncellenecek ve bu çakışmaları da tanıtabiliriz.

![](Images/Day42_Containers6.png)

Konteynerler bu sorunu çözmeye yardımcı olabilir. Konteynerler uygulamanızı **inşa etme**nize, uygulamayı **paketleme**nize, **dağıtma**nıza ve **ölçeklendirme**nize yardımcı olur. Mimarisine bir göz atalım, donanım ve işletim sisteminizin üzerinde daha sonra ele alacağımız bir konteyner motoruna sahip olacaksınız. Konteyner motoru yazılımları, uygulamanın ihtiyaç duyduğu kütüphaneleri ve bağımlılıkları paketleyen konteynerler oluşturmanıza yardımcı olur, bu nedenle bu konteyneri bir makineden diğerine sorunsuzca taşıyabilirsiniz ve uygulamanın çalışması için gereken kütüphaneler ve bağımlılıklar hakkında endişelenmeniz gerekmez, çünkü bunlar bir paketin parçası olarak gelir, ki bu da konteynerden başka bir şey değildir, bu nedenle farklı konteynerlere sahip olabilirsiniz ve bu konteyneri sistemler arasında taşıyabilirsinizken uygulamanın çalışması için gerekli olan altta yatan bağımlılıklardan endişelenmeden.

![](Images/Day42_Containers7.png)

### Bu Containerların Avantajları

- Konteynerler, tüm bağımlılıkları içinde paketlemeye ve izole etmeye yardımcı olur.

- Konteynerleri yönetmek kolaydır.

- Bir sistemden diğerine taşımak kolaydır.

- Konteynerler, yazılımı paketlemeye yardımcı olur ve herhangi bir tekrar çaba gerektirmeden kolayca taşıyabilirsiniz.

- Konteynerler kolayca ölçeklenebilir.

Konteynerleri kullanarak bağımsız konteynerleri ölçeklendirebilir ve bir yük dengeleyici veya hizmet kullanabilirsiniz, bu da trafiği bölmeye yardımcı olur ve uygulamaları yatay olarak ölçeklendirebilirsiniz. Konteynerler, uygulamalarınızı nasıl yönettiğinizde büyük bir esneklik ve kolaylık sunar.

### Container Nedir? 

Bilgisayarlarımızda uygulamaları çalıştırdığımızda, bu web tarayıcısı veya bu yazıyı okumak için kullandığınız VScode gibi bir şey olabilir. Bu uygulama, bir işlem olarak çalışır veya işlem olarak bilinen bir şeydir. Dizüstü bilgisayarlarımız veya sistemlerimizde genellikle birden fazla uygulama veya işlem çalıştırmaya çalışırız. Yeni bir uygulama açtığımızda veya uygulama simgesine tıkladığımızda, çalıştırmak istediğimiz bir uygulamadır; bazen bu uygulama sadece arka planda çalışmasını istediğimiz bir hizmet olabilir. İşletim sistemimiz, arka planda çalışan ve sistemle ilgili kullanıcı deneyimini sağlayan birçok hizmetle doludur.

Bu uygulama simgesi, dosya sisteminizdeki bir yürütülebilir dosyaya bir bağlantıyı temsil eder; işletim sistemi daha sonra bu yürütülebilir dosyayı belleğe yükler. İlginç bir şekilde, bu yürütülebilir dosya bazen bir işlem hakkında konuştuğumuzda bir görüntü olarak adlandırılır.

Konteynerler işlemlerdir. Bir konteyner, kodu ve tüm bağımlılıklarını paketleyen standart bir yazılım birimidir, böylece uygulama bir hesaplama ortamından başka birine hızlı ve güvenilir bir şekilde çalışır.

Konteynerleştirilmiş yazılım her zaman altyapıdan bağımsız olarak aynı şekilde çalışacaktır. Konteynerler yazılımı ortamından izole eder ve geliştirme ve sahneleme gibi farklılıklar olsa bile aynı şekilde çalışmasını sağlar.

Önceki bölümde, konteynerlerin neden ve nasıl popüler hale geldiği konusunda konteynerler ve görüntülerin nasıl birleştiğinden bahsettim.

### Image Nedir? 

Bir konteyner görüntüsü, bir uygulamayı çalıştırmak için gereken her şeyi içeren hafif, bağımsız, yürütülebilir bir yazılım paketidir: kod, çalıştırma zamanı (runtime), sistem araçları, sistem kütüphaneleri ve ayarlar. Konteyner görüntüleri, çalışma zamanında konteynerlere dönüşürler.

## Kaynaklar 

- [TechWorld with Nana - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=3c-iBn73dDE)
- [Programming with Mosh - Docker Tutorial for Beginners](https://www.youtube.com/watch?v=pTFZFxd4hOI)
- [Docker Tutorial for Beginners - What is Docker? Introduction to Containers](https://www.youtube.com/watch?v=17Bl31rlnRM&list=WL&index=128&t=61s)
- [Introduction to Container By Red Hat](https://www.redhat.com/en/topics/containers)

Gorusmek Uzere [Gun 43](day43.md) 
