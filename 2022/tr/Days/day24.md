## Network Otomasyonu

### Network Otomasyonunun Temelleri

Network otomasyonunun başlıca itici güçleri:

- Çeviklik elde etmek
- Maliyetleri azaltmak
- Hataları ortadan kaldırmak
- Uyumluluğu sağlamak
- Merkezi yönetim

Otomasyonun benimsenme süreci her şirket için özeldir. Otomasyonun uygulanması söz konusu olduğunda herkes için tek bir standart yoktur; organizasyonunuz için en iyi çalışan yaklaşımı belirleme ve benimseme yeteneği, daha çevik bir ortam oluşturmak veya sürdürmek için önemlidir. Her zaman iş değeri ve son kullanıcı deneyimi üzerinde odaklanılmalıdır. (Birkaç gün önce, DevOps ve kültürel değişimle ilişkili olarak benzer bir konu ve otomatikleşme süreci hakkında konuşulmuştu).

Bunu ayrıntılandırmak için, otomatikleştirmeye çalıştığınız görevin veya sürecin, son kullanıcı deneyimini veya iş değerini nasıl geliştireceğini belirlemelisiniz ve sistematik bir adım adım yaklaşım izlemelisiniz.

"Nereye gittiğinizi bilmiyorsanız, herhangi bir yol sizi oraya götürecektir."

Ulaşmayı hedeflediğiniz bir tasarım çerçevesi veya yapıya sahip olun, nihai hedefinizi bilin ve iş sonuçlarına dayanarak otomasyonun başarılarını farklı aşamalarda ölçün.

Otomasyon kavramlarını bir kabarcıkta tasarlamaya gerek yok çünkü bunlar uygulamanıza, hizmetinize ve altyapınıza uygulanmalıdır. Bu nedenle, mevcut altyapınıza ve uygulamalarınıza odaklanarak kavramları oluşturmaya ve modellenmeye başlayın.

### Ağ Otomasyonuna Yaklaşım

Ağdaki görevleri ve ağ değişiklik taleplerini **belirlemeliyiz**, böylece en yaygın sorunları tespit edebilir ve sonrasında çözümün otomatikleştirilmesini sağlayabiliriz.

- Şu anda manuel olarak ele alınan tüm değişiklik talepleri ve iş akışlarının bir listesini yapın.
- En yaygın, zaman alıcı ve hata yapma olasılığı yüksek olan faaliyetleri belirleyin.
- İş odaklı bir yaklaşım benimseyerek talepleri önceliklendirin.
- Bu, otomasyon sürecini oluşturmanın ve nelerin otomatikleştirilmesi gerektiğinin çerçevesini oluşturur.

Daha sonra, **görevleri parçalara ayırın ve analiz edin.** Farklı ağ fonksiyonlarının nasıl çalıştığını ve birbirleriyle etkileşime girdiğini anlamak önemlidir:

- Altyapı/ağ ekibi, uygulamaları dağıtmak için birden çok katmanda değişiklik talepleri alır.
- Ağ hizmetlerine dayanarak, bunları farklı alanlara bölebilir ve nasıl etkileşime girdiklerini anlayabilirsiniz.
  - Uygulama optimizasyonu
  - ADC - Uygulama Dağıtım Denetleyicisi
  - Firewallar
  - DDI (DNS, DHCP, IPAM, vb.)
  - Yonlendirme
  - Digerleri
- Farklılıkları ele almak ve ekip arasında işbirliğini teşvik etmek için farklı bağımlılıkları tanımlayın.

**Yeniden kullanılabilir politikalar**: Yeniden kullanılabilir hizmetlerin, süreçlerin ve giriş/çıkışların tanımlanması ve basitleştirilmesi.

- Farklı hizmetler, süreçler ve giriş/çıkışlar için tekliflerin belirlenmesi.
- Dağıtım sürecinin basitleştirilmesi, yeni ve mevcut iş yüklerinin pazarlama süresini kısaltacaktır.
- Standart bir süreç oluşturulduktan sonra, bireysel taleplerle sıralanabilir ve çoklu iş parçacığı yaklaşımıyla uyumlu hale getirilebilir.

**Şirketin özel faaliyetleriyle politikaları birleştirin**. Bu politikaların uygulanması şirkete nasıl yardımcı olur? Zaman tasarrufu sağlar mı? Para tasarrufu sağlar mı? İş sonuçlarını iyileştirir mi?

- Hizmet görevlerinin birbiriyle uyumlu olduğundan emin olun.
- Artan hizmet görevlerini birleştirerek iş hizmetleri oluşturacak şekilde ilişkilendirin.
- Taleplere bağlı olarak hizmet görevlerini ilişkilendirme ve yeniden ilişkilendirme esnekliği sağlayın.
- Otomatik hizmet yeteneklerini dağıtın ve işletme verimliliğini artırmak için yol açın.
- Birden fazla teknoloji beceri setinin denetim ve uyum sağlamaya devam etmesine izin verin.

**Politika ve süreçleri iteratif olarak** değerlendirin ve geliştirin, aynı zamanda kullanılabilirliği ve hizmeti sürdürerek zamanla iyileştirmeler yapın:

- Var olan görevleri otomatikleştirmeye yavaşça başlayın.
- Otomasyon sürecine aşina olun, böylece otomasyondan fayda sağlayabilecek diğer alanları tespit edebilirsiniz.
- Otomasyon girişimlerinizi, kullanılabilirliği sağlarken aşamalı olarak esneklik ekleyerek aşama aşama gerçekleştirin.
- Başarıya giden yol için aşamalı bir yaklaşım benimseyin.

**Ağ hizmetlerini orkestre edin:**

- Uygulamaları hızla sunmak için dağıtım sürecini otomatikleştirmek önemlidir.
- Esnek bir hizmet ortamı oluşturmak için farklı bileşenleri teknoloji yığınları arasında yönetin.
- Otomasyonu ve dağıtım sırasını kontrol etmeyi sağlayan uçtan uca bir orkestrasyon için hazırlık yapın.

## Ağ Otomasyonu Araçları

İyi bir haber, ağ otomasyonunda kullandığımız araçların çoğunlukla diğer otomasyon alanlarında da kullanılan araçlar olduğudur. Şimdiye kadar ele aldığımız veya gelecekte ele alacağımız konuları kapsar.

**İşletim Sistemi** - Bu dersin çoğunu bir Linux işletim sistemiyle öğrenmek odaklanır. Linux bölümünde sebepleri açıklandı, ancak dokunacağımız araçların çoğu zaten çok platformlu olsa da çoğunluğu Linux tabanlı uygulamalar olarak başladı.

**Entegre Geliştirme Ortamı (IDE)** - Burada söyleyecek çok şey yok, ancak farklı diller için birçok eklentiye sahip olduğu için Visual Studio Code'yu bir IDE olarak öneririm.

**Konfigurasyon Yönetimi** - Henüz yapılandırma yönetimi bölümüne ulaşmadık, ancak yapılandırma yönetimi ve otomasyonu için bu alanda favorilerden biri olan Ansible'ı belirtmek önemlidir. Ansible Python ile yazılmış olmasına rağmen, Python bilmek gerekli değildir.

- Agent gerektirmez (Agentless)
- Sadece SSH gerektirir.
- Büyük Destek Topluluğu
- Birçok Network Modülü
- Sadece Itme Modeli
- YAML ile yapılandırılır
- Açık kaynaklıdır!

[Ansible Network Modulleri Linki](https://docs.ansible.com/ansible/2.9/modules/list_of_network_modules.html)

**Ansible Tower** Yapılandırma yönetimi bölümünde Ansible Tower'a değineceğiz, bu Ansible'ın grafik arayüzüdür(GUI).

**CI/CD** - Entegrasyon ve sürekli dağıtım kavramları ve araçlarının daha fazla yönünü ele alacağız, ancak en azından burada bahsetmek önemlidir, çünkü bu ağ oluşturma ve hizmet ve platformların tüm sağlanmasını kapsar.

Özellikle, Jenkins, network otomasyonu için popüler bir araçtır.

- Git reposundaki değişiklikler için izler ve ardından bunları başlatır.

**Sürüm Kontrolü** - Daha sonra daha detaylı olarak ele alacağımız bir başka konudur.

- Git, yerel cihazınızdaki ve uzaktaki kodun sürüm kontrolünü sağlar.
- Çok platformlu
- GitHub, GitLab, Gitea, BitBucket vb. depolama alanlarıdır. Kodlarınızı saklayabileceğiniz ve yükleyebileceğiniz çevrimiçi web siteleridir.

**Language | Scripting** - Kapsamayacağımız bir programlama dili Python'dır. Göreceli bir bakış açısıyla çeşitli ihtiyaçları karşılamak için Go dili seçildi, ancak Python her yerde yaygın olarak kullanılmaktadır ve ağ otomasyonunda kazanan gibi görünmektedir.

- Nornir, Python'da yazılmış bir otomasyon çerçevesidir. Bu, özellikle Ağ Otomasyonu etrafında Ansible'ın rolünü üstlenmektedir. Nornir Belgelerine göz atın. [Nornir Dokumanları](https://nornir.readthedocs.io/en/latest/)

**API'leri Analiz Etmek** - Postman, RESTful API'leri analiz etmek için harika bir araçtır. API'leri oluşturmanıza, test etmenize ve değiştirmenize yardımcı olur.

- POST >>> Kaynak objeleri olusturmak için
- GET >>> Bir kaynaktan veri almak için.
- PUT >>> Kaynakları oluşturmak veya değiştirmek için.
- PATCH >>> Bir kaynak nesnesi oluşturmak veya güncellemek için.
- Delete >>> Bir kaynağı silmek için.

[Postman Araçlarının İndirme Linki](https://www.postman.com/downloads/)

### Bahsedilmesi gereken diğer araçlar

- [Cisco NSO (Network Services Orchestrator - Ağ Hizmetleri Orkestratörü)](https://www.cisco.com/c/en/us/products/cloud-systems-management/network-services-orchestrator/index.html)
- [NetYCE - Ağ Otomasyonunu Basitleştirme](https://netyce.com/)
- [Ağ Test Otomasyonu](https://pubhub.devnetcloud.com/media/genie-feature-browser/docs/#/)

Önümüzdeki 3 gün boyunca ağ üzerine kapsadığımız bazı teorilerle çalışmaya başlayacağız, böylece Python ve ağ otomasyonu konularında biraz pratik yapabileceğiz.

Ağ konularının tamamını kapsamadık, sadece temel kavramlara sahip olmak için gerekli olan teorik bilgileri gözden geçirdik. Aşağıda verilen kaynaklarla bilgilerinizi genişletmeniz önerilir.

## Kaynaklar

- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Uygulamalı Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

Gorusmek Uzere [Gun 25](day25.md).
