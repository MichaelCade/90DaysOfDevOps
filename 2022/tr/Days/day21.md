## DevOps & Networking

Tüm bölümlerde olduğu gibi, burada da açık ve ücretsiz eğitim materyalleri kullanılıyor. Gösterilen içeriğin büyük bir kısmı [Networking Uygulamalı](https://www.practicalnetworking.net/) ve [Networking Temelleri Serisi](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi) gibi kaynaklara atfedilebilir. Kaynaklar ve bağlantılarla birlikte belirtilmiştir, ancak topluluk perspektifinden bunu vurgulamak uygun olur. Bu kurs, benim belirli teknoloji alanları hakkında daha fazla anlamamı sağlamak için kullanılmıştır. Bu depo, notlarımı almak ve topluluğun bu notlar ve listelenen kaynaklardan faydalanmasını sağlamak için bir yer olarak hizmet eder.

21. güne hoş geldiniz! Önümüzdeki 7 gün boyunca ağları keşfedeceğiz. DevOps'ta, platform mühendisi pozisyonunda kendimizi savunabilmek için network(ag) oluşturma hakkında temel bilgilere sahip olmak önemlidir.

Sonuç olarak, daha önce de söylediğimiz gibi, DevOps, organizasyonlardaki bir kültür ve süreç değişikliğidir; sanal makineler, konteynerler veya Kubernetes ile olabilir. Ağ da bu değişime dahil olabilir. Altyapımız için DevOps prensiplerini kullanıyorsak, ağı da DevOps perspektifiyle ele almalıyız. Ayrıca, farklı topolojiler, ağ araçları ve mevcut olabilecek yığınlar gibi ağ hakkında bilgi sahibi olmanız gerekmektedir.

Tavsiye edilen yaklaşım, ağ cihazlarımızı altyapıyı kod olarak yapılandırarak ([IaC](https://es.wikipedia.org/wiki/Infraestructura_como_c%C3%B3digo)) sahip olmaktır. Bu konuyu daha önce de bahsetmiştik. Ayrıca, sanal makinelerimizi nasıl otomatikleştiriyorsak ağımızı da aynı şekilde otomatikleştirmeliyiz, ancak bunu yapmak için otomatikleştirdiğimiz şeyin iyi bir anlayışını edinmemiz gerekmektedir.

Hadi Başlayalım!

### NetDevOps | Network DevOps Nedir?

Bulut kültürü içinde Network DevOps veya NetDevOps gibi terimleri duyacaksınız. Belki de zaten bir ağ mühendisisiniz ve altyapıdaki ağ bileşenlerini çok iyi biliyorsunuz, bu durumda DHCP, DNS, NAT vb. gibi network ile ilgili unsurları ve bu unsurların önemini anlayacaksınız. Ayrıca, donanım veya yazılım tarafından tanımlanan ağ seçeneklerini, switch'leri, yönlendiricileri vb. iyi bir şekilde bileceksiniz.

Eğer bir Network mühendisi değilseniz, muhtemelen Network DevOps'un hedefini anlamak için bu alanda temel bir bilgiye ihtiyaç duyabilirsiniz.

NetDevOps veya Network DevOps terimlerini DevOps prensiplerinin ve uygulamalarının ağa uygulanması olarak düşünebiliriz. Bu, ağın oluşturulması, test edilmesi, izlenmesi ve dağıtılması süreçlerine sürüm kontrolü ve otomasyon araçlarının uygulanması anlamına gelir.

Network'un otomasyon gerektirmesi nedeniyle, DevOps'un takım arasındaki sınırları kaldıran bir yapı olduğunu söylediğimizde geriye dönmeliyiz. Eğer network ekipleri benzer bir model ve sürece geçmezse, genel başarısızlık noktasına veya hatta engelleyici bir faktöre dönüşebilirler.

Yeniden aprovision, yapılandırma, test, sürüm kontrolü ve dağıtım gibi alanlarda otomasyon prensiplerini kullanmak iyi bir başlangıçtır. Genel olarak, otomasyon, dağıtımlarda hız ve esnekliği artırmaya yardımcı olurken, ağ altyapısının istikrarını ve buna bağlı olarak paylaşılan bir süreci birden fazla ortamda test edildikten sonra elde etmeyi sağlar. Örneğin, bir ortamda tamamen test edilmiş bir ağ politikası, bir kod olduğu için hızla başka bir konumda kullanılabilir.

Bu düşüncenin iyi bir bakış açısı ve özeti şurada bulunabilir:  [Network DevOps](https://www.thousandeyes.com/learning/techtorials/network-devops).

## Network'un Temelleri

Bir süreliğine DevOps perspektifini bir kenara bırakıp, ağın bazı temel kavramlarına kısaca göz atalım.

*Bu içeriği video formatında tercih ederseniz, Practical Networking'in bu videolarına göz atabilirsiniz:*

* [*Network Cihazları - Hosts, IP Addresses, Networks - Networking Temelleri - Ders 1a*](https://www.youtube.com/watch?v=bj-Yfakjllc&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=1)
* [*Network Cihazları - Hub, Bridge, Switch, Router - Networking Temelleri - Ders 1b*](https://www.youtube.com/watch?v=H7-NR3Q3BeI&list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi&index=2)

### Network Cihazları

**Host** - Trafik gönderen veya alan tüm cihazlardır.

![](Images/Day21_Networking1.png)

**IP Adresi** - Her bir ana bilgisayarın kimliğidir.

![](Images/Day21_Networking2.png)

**Network** - Hostlar arasında trafiği taşıyan yapıdır. Ağ olmasaydı, verilerin manuel olarak taşınması gerekecekti!

Benzer bir bağlantı gerektiren mantıksal bir host grubu.

![](Images/Day21_Networking3.png)

**Switches** - Bir ağ içinde iletişimi kolaylaştırır. Veri paketlerini ana bilgisayarlar arasında ileterek ve doğrudan ana bilgisayarlara göndererek işlev görür.

- Network: Benzer bir bağlantı gerektiren ana hostların gruplandığı yapıdır.
- Bir ağdaki hostlar aynı IP adresi alanını paylaşırlar.

> *Daha gelişmiş olanlar, bir yönlendirici gibi davranabilir. Benzer özelliklere sahip cihazlar arasında yönlendirme yapabilir, ancak bir yönlendiricinin sağlayabileceği tüm yönlendirme protokollerine sahip olmayabilirler.*

![](Images/Day21_Networking4.png)
1
**Router(Yonlendirici)** - Ağlar arasındaki iletişimi kolaylaştırır. Bir switch ağ içinde iletişimi sağlarken, router bu ağları birleştirmemize veya izin verildiği durumlarda aralarında erişim sağlamamıza olanak tanır.

> *Evlerimizde bildiğimiz yönlendiriciden bahsetmiyoruz. Sadece yönlendirme işlevini yerine getiren cihazdan bahsediyoruz. Şu anda [ISP](https://es.wikipedia.org/wiki/Proveedor_de_servicios_de_internet)'lerden aldığımız, benim görüşüme göre hatalı olan şeyler olarak adlandırdığımız cihazlar aslında 4 bileşenden oluşur: Router, switch, erişim noktası ve modem. Bu cihazlar arasındaki farkı anlamak için bu cihaz farkını göz önünde bulundurmalıyız.*

Bir yönlendirici trafiği kontrol etmek için bir nokta sağlayabilir (güvenlik, filtreleme, yönlendirme). Günümüzde birçok switch de bu işlevlerin bazılarını sağlamaktadır.

Yönlendiriciler hangi ağlara bağlı olduklarını öğrenir. Bir yönlendirme tablosu, bir yönlendiricinin bildiği ağların kümesini içerir.

Bir yönlendiricinin bağlı olduğu ağlarda bir IP adresi bulunur. Bu IP aynı zamanda her bir ana bilgisayarın yerel ağından çıkış noktası olarak da kullanılır, bu aynı zamanda bir ağ geçidi (gateway) olarak da bilinir.

Yönlendiriciler ayrıca yukarıda bahsettiğimiz ağ hiyerarşisini oluştururlar.

> *Daha gelişmiş olanlar, bir switch ile aynı özelliklere sahip olabilir.*

![](Images/Day21_Networking5.png)

## Switches vs. Routers

**Routing**  verileri ağlar arasında taşıma işlemidir.

- Bir router, temel amacı yönlendirme olan bir cihazdır.

**Switching** verileri bir ağ içinde taşıma işlemidir.

- Bir switch, ana işlevi anahtarlama olan bir cihazdır.

Bazı cihazlara genel bir bakış verdik, ancak Network'te başka işlevlere sahip birçok daha fazla cihazın olduğunu bilmeliyiz, Örneğin:

- Access pointler
- Firewall'lar (fiziksel, donanım tabanlı).
- Load Balancerlar (fiziksel, donanım tabanlı).
- 3. Katman Switchler.
- IDS / IPS.
- Proxyler.
- Sanal Switchler(mantıksal, yazılım tabanlı).
- Sanal Routerlar(mantıksal, yazılım tabanlı).

Önceki cihazların hepsi, yönlendirme (routing) ve/veya anahtarlamayı (switching) bir şekilde gerçekleştirecektir.

Gelecek günlerde aşağıdaki konuları daha detaylı bir şekilde öğreneceğiz:

- OSI Modeli
- Network Protokolleri
- DNS (Domain Name System)
- NAT
- DHCP
- Subnetler(Alt aglar)

Bu kavramlar, ağları anlamak için temel oluşturacak çok önemli konulardır. Eğer bu terimlere aşina değilseniz, ağın ileriki günlerine devam etmeniz en iyisi olacaktır. Eğer bu terimleri biliyorsanız, gözden geçirebilir veya doğrudan bir sonraki konuya geçebilirsiniz: [Cloud Provider(Bulut Saglayıcı)](day28.md)

## Kaynaklar

* [Networking Temelleri](https://www.youtube.com/playlist?list=PLIFyRwBY_4bRLmKfP1KnZA6rZbRHtxmXi)
* [Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)

Gorusmek Uzere [Gun 22](day22.md).
