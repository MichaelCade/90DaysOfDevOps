## [Planlama](#planlama) > [Kodlama](#kodlama) > [Insa Etme](#insaetme) > [Test Etne](#testetme) > [Versiyon Yayınlama](#versiyonyayinlama) > [Dagitim](#dagitim) > [Isletme](#isletme) > [Izleme](#izleme) > [Sürekli Tekrar](#sureklitekrar)


Planlama > Kodlama > Insa Etme > Test Etme > Surum Yayinlama > Dagitma > Isletme > Izleme >

Bugün DevOps bakış açısından bir Uygulamanın yasam döngüsünde adım adım ilerlemeye odaklanacağız.

![DevOps](Images/Day5_DevOps8.png)

### Planlama(Plan)

Her şey planlama süreciyle başlar, bu süreçte geliştirme ekibi bir araya gelir ve bir sonraki sprintlerinde hangi tür özellikleri ve hata düzeltmelerini yayımlayacaklarını belirler. Bir DevOps mühendisi olarak, bu noktada da dahil olmanız önemlidir, böylece ileride nelerin geleceğini öğrenir ve kararları etkileyerek hem onlar için en iyi uyarlanabilecek altyapıyla çalışmalarına yardımcı olabilirsiniz. Burada dikkat edilmesi gereken önemli bir nokta, hem geliştiricilerin hem de yazılım mühendisliği ekibinin DevOps mühendislerinin müşterileri olduğudur, bu nedenle yanlış bir yola gitmeden önce müşterilerle çalışmak önemlidir.

### Kodlama(Code)

Planlama oturumu tamamlandıktan sonra kod yazmaya başlayacaklar ve önemli olan, kod yazılırken DevOps felsefesinin, hangi hizmetlere ihtiyaç duyulduğu biliniyorsa altyapıyı daha iyi anlamada yardımcı olabilmesidir. Başlangıcı takip etmek, mümkün olan en kısa sürede bu kodu depolara birleştirmek için uygun hizmetleri seçmenin en iyi yoludur.

### Insa Etme (Build)

İşte otomatizasyon süreçlerimizin ilki olan inşa sürecine başlayacağımız yer burasıdır. Uygulamanın kodunu alacak ve kullandıkları dilin bağımlılıklarını oluşturacağız. Bu, yorumlanabilir veya derlenebilir olabilir veya bu kodla bir Docker imajı oluşturulabilir. Her durumda, bu süreci bir pipeline kullanarak baştan itibaren gerçekleştireceğiz. [CI/CD](https://en.wikipedia.org/wiki/CI/CD).

## Test Etme (Testing)

Container'imizi oluşturduktan sonra ilk testleri gerçekleştireceğiz. Geliştirme ekibi genellikle yazılım için uygun testleri yazar, ancak biz üretim ortamında hataların en aza indirgenmesine yardımcı olma amacıyla testleri gerçekleştirmeliyiz. Bu, tam bir garanti sağlamaz, ancak olası sorunları en aza indirgemek ve daha önce çalışan şeylerin bozulmamasını sağlamak için elimizden geleni yapmalıyız. Production ortamında "Benim makinede çalışıyordu" demek kabul edilemez.

## Versiyon Yayınlama(Release)

Testlerden geçtikten sonra kodumuzu yayınlama sürecine devam ederiz ve yine, çalıştığımız uygulama türüne bağlı olarak doğru adımları izlememiz gerekmektedir. Örneğin, kod GitHub veya herhangi bir özel git deposunda olabilir veya derlenmiş bir kod veya zaten oluşturduğunuz bir Docker imajı olabilir ve bunu üretim sunucularının erişebileceği bir kayıt veya depoya yüklemeniz gerekebilir, böylece dağıtım ve sürüm kontrolü süreci mümkün olduğunca basit olur.

## Dagıtım (Deploy)

İşte yaşam döngüsünün son aşamasına geldik çünkü dağıtımlarla kodu üretim ortamına yerleştiriyoruz ve işletme burada yazılım mühendisliği ekiplerinin harcadığı tüm çaba ve zamanın değerini fark ediyor.

## Isletme(Operate)

Şimdi dağıtımımızla işletmeye başlayacağız. Bu, site veya uygulamanın yavaş çalıştığından dolayı müşterilerden gelen şikayetler almaya başlamak gibi bir şeyi içerebilir. Bu durumda, mümkün olan en kısa sürede neyin yanlış gittiğini metrikler aracılığıyla belirlememiz gerekecektir. Büyük trafik pikleri olan dönemlerde kaynak sayısındaki artışı yönetmek için otomatik ölçeklendirme yapmak, düşük yoğunluklu dönemlerde kaynakları azaltmak gibi işlemleri gerçekleştirmemiz gerekebilir. Bu sadece dikkate almanız gereken işlemlerden biridir.

Diğer bir işlem, üretim ekibine geri bildirim döngüsü olarak üretimle ilgili önemli olayları ilk elden bildirmelerine izin vermektir. Yanlış bir dağıtım veya bir sürüm geri alma ihtiyacı gibi. Bu otomatik olabilir veya olmayabilir, her zaman ortama bağlıdır. Amaç, mümkün olduğunca otomatikleştirmek olsa da bazı ortamlarda maliyetler açısından otomatik ölçeklendirme bir dezavantaj olabilir. Bazı ortamlarda otomatik ölçeklendirme için bazı adımların hazırlanması gerekebilir veya üretim zaten istikrarlıdır ve otomatik ölçeklendirmeye ihtiyaç duymaz.

İdeal olan, her tahmin edilebilir dağıtımın otomasyon süreci içinde olmasıdır. Ve bunu yaparken, otomatikleştirmenin yanı sıra tüm işlemleri içeren işlemler adımlarını da içermelidir, ayrıca operasyon ekibinin otomatik değişiklikleri ve tüm önemli olayları bildiren bir bildirim türü gibi. Tabii ki, dağıtımlar da bunun içinde yer alır.

## Izleme (Monitoring)

Tüm bu önceki adımlar sonucunda, son adıma geçiyoruz: izleme. Özellikle bahsettiğimiz otomatik ölçeklendirme sorunlarının çözümüne odaklanır. İzleme olmadan bir sorunun olduğunu nasıl bileceksiniz ki?

Bu nedenle, izleme aracılığıyla kontrol edilebilecek bazı şeyler şunlardır:

- Bellek kullanımı.
- CPU kullanımı.
- Disk alanı.
- API'nin son noktası.
- Yanıt süresi.
- Son noktanın ne kadar hızlı yanıt verdiği.
- Trafik genişliği.
- Bunların birçoğunda size loglar yardımcı olacak, geliştiricilerin üretim sistemlerine erişmeden neler olduğunu görmelerini sağlayan bir yetenek sunarlar.

## Surekli Tekrarlama

Bu noktada, planlama aşamasına geri dönülür ve süreç ikinci, üçüncü, dördüncü tekrarlamada yeniden başlatılır. Yazılım mühendisliği ve altyapı çalışması, her zaman en iyi uygulamayı müşterilere/son kullanıcılara sunmak için sürekli tekrarlamaya devam etmektir.

## Süreklilik

Sürekli işlemi gerçekleştirmemize yardımcı olan birçok araç vardır, tüm kodu ve son hedefi tamamen otomatik hale getirmek. Bulut altyapısı veya herhangi bir ortam sıklıkla **Sürekli Entegrasyon / Sürekli Dağıtım / Sürekli Dağıtım** olarak tanımlanır. Kısaltması "CI/CD"dır. Bu 90 günlük süre boyunca, temelleri anlamak ve içselleştirmek için CI/CD'yi anlamak için örnekler ve rehberlerle bir hafta ayrılacaktır.

### Sürekli Teslimat

**Sürekli Teslimat** = Planlama > Kodlama > Insa Etme(Build) > Test Etme

### Sürekli Entegrasyon

Bu, önceki **Sürekli Dağıtım** aşamalarının sonuçlarının ve Başarısızlık ve Başarı aşamalarının sonuçlarının birleşimidir. Ancak bu, **sürekli dağıtım**a geri bildirimde bulunur veya **Sürekli Dağıtım** olarak taşınır.

**Sürekli Entegrasyon** = Planlama > Kodlama > Oluşturma > Test Etme > Başlatma

### Sürekli Dagıtım

Başarılı bir **sürekli entegrasyon** ile **Sürekli Dağıtım**a geçebiliriz ve aşağıdaki aşamaları içerir:

Başarılı sürekli entegrasyon = **Sürekli Dağıtım** = Dağıtma > İşletme > İzleme

Bu üç sürekli kavram, DevOps yaşam döngüsündeki aşamaları kapsar.

Bu son bölüm, Ders 3 için bir özet oldu ve mümkün olduğunca açık olması için yapıldı.

### Kaynaklar

- [DevOps for Developers – Software or DevOps Engineer?](https://www.youtube.com/watch?v=a0-uE3rOyeU)
- [Techworld with Nana -DevOps Roadmap 2022 - How to become a DevOps Engineer? What is DevOps?](https://www.youtube.com/watch?v=9pZ2xmsSDdo&t=125s)
- [How to become a DevOps Engineer in 2021 - DevOps Roadmap](https://www.youtube.com/watch?v=5pxbp6FyTfk)

Gorusmek Uzere [Gun 6](day06.md).
