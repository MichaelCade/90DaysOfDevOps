## DevOps Yaşam Döngüsü - Uygulama Odaklı

Devam ettikçe, Sürekli Geliştirme, test etme, Dağıtım, İzleme gibi kavramlarla sürekli olarak karşılaşacağız. Eğer bir DevOps Mühendisi olma yolunda ilerliyorsanız, tekrarlanabilirlik sizin alıştığınız bir şey olacak. Sürekli olarak gelişmek, her şeyi ilginç tutan şeydir.

Bu saatte, uygulamanın başından sonuna kadar ve ardından sürekli bir döngü olarak bir kez daha genel bir bakış açısına sahip olacağız.

### Gelistirme

Bir uygulama örneği üzerinden başlayalım. Başlangıçta henüz hiçbir şey oluşturulmamıştır, belki de bir geliştirici olarak müşteriyle veya son kullanıcıyla gereksinimleri tartışmamız ve uygulama için bir plan veya gereksinimlere ulaşmamız gerekmektedir. Ardından, bu gereksinimlere dayanarak yeni uygulamamızı oluşturmamız gerekmektedir.

Bu aşamada araçlar açısından, uygulamanın hangi IDE ve programlama diliyle oluşturulmak istendiğinden başka bir gerçek gereksinim yoktur.

Bir DevOps mühendisi olarak, bu planı oluşturmayacak veya uygulamayı son kullanıcı için kodlamayacaksınız, bunu bir uzman geliştirici yapacaktır. Ancak, uygulamaya yönelik altyapı kararlarını alabilmek için kodun bir kısmını okuyabiliyor olmanız faydalı olabilir.

Daha önce belirtildiği gibi, uygulama herhangi bir dilde yazılabilir. Sürüm kontrol sistemi kullanmanız önemlidir, bu konu daha sonra detaylı bir şekilde ele alınacak ve özellikle **Git** üzerine odaklanacağız.

Proje üzerinde yalnızca bir geliştiricinin çalışmadığı durumlar da olasıdır, ancak bu durumda bile en iyi uygulamalar, kodu depolamak ve işbirliği yapmak için bir kod deposu gerektirir. Bu kod deposu, özel veya genel olabilir ve özel veya genel olarak dağıtılabilir. Genellikle bu amaçla **GitHub** veya **GitLab** gibi bir kod deposu kullanılır. Bunlar hakkında daha fazla bilgiyi daha sonra **Git** bölümünde ele alacağız.

### Testing


Bu aşamada, gereksinimlerimiz ve geliştirme aşamasındaki uygulamamız var. Ancak, kodumuzu farklı mevcut ortamlarda veya özellikle seçtiğimiz programlama dilinde test ettiğimizden emin olmamız gerekiyor.

Bu aşamada, kalite kontrol departmanı olası hataları test eder. Gittikçe daha fazla şekilde test ortamını simüle etmek için konteynerlerin kullanıldığını görüyoruz, bu genel olarak fiziksel altyapı veya bulutun genel giderlerinde büyük tasarruf sağlayabilir.

Bu aşama aynı zamanda sürekli entegrasyonun bir parçası olarak otomatikleştirilebilir.

Bu testleri manuel olarak yapması gereken 10, 100 veya hatta 1000 kalite kontrol mühendisiyle karşılaştırıldığında bu testleri otomatikleştirebilme yeteneği kendisi için konuşur. Bu mühendisler diğer konulara odaklanabilir, örneğin, yığının hızlı ilerlemesini sağlamak veya hata ve yazılım testlerine karşı daha fazla işlevsellik geliştirmek gibi. Bu, çoğu geleneksel yazılım yayınında kullanılan aşamalı bir metodolojiyi kullanan yazılım yayınlarında gecikmeye neden olur.

### Entegrasyon

DevOps yaşam döngüsünün merkezinde bulunur. Bu uygulama geliştiricilerin kod değişikliklerini daha sık doğrulamalarını gerektiren bir uygulamadır. Bu günlük veya haftalık olabilir.

Her bir doğrulama ile uygulamanız otomatik test aşamalarından geçebilir ve bu, sonraki aşamaya geçmeden önce sorunların veya hataların erken tespitini sağlar.

Endişelenmeyin, birçok şirket bunu yapıyor ve yapmaya devam edecek ve yazılım sağlayıcısı önceki 3 aşamaya odaklanacak, ancak son aşamayı benimsemek isteyebilirsiniz çünkü bu hızlı ve verimli bir dağıtım sürecine olanak tanır.

Bu bilgiye sahip olmak önemlidir çünkü bugün platform yazılımını satın alabilirsiniz, ancak yarın veya gelecekteki işte ne olacağını bilemezsiniz.

### Deploy

İyi, uygulamamızı inşa ettik ve gereksinimlerimize uygun bir şekilde test ettik. Şimdi ilerlemeli ve kullanıcılarımızın uygulamayı kullanabileceği üretim ortamına dağıtmamız gerekiyor.

Bu, kodun üretim sunucularına dağıtıldığı aşamadır. İşler şimdi çok ilginç hale geliyor ve geriye kalan 86 günümüzü burada geçireceğiz. Farklı uygulamalar farklı donanımlar ve/veya yapılandırmalar gerektirebilir. Bu noktada, **Uygulama Yapılandırma Yönetimi** ve **Kod Üzerinde Altyapı** önemli bir rol oynayabilir. Uygulama **konteynerlerde** olabilir, ancak bir sanal makinede çalıştırılabilir. Bu bizi **Kubernetes** gibi platformlara yönlendiriyor, bu platform konteynerlerimizi orkestrasyon yapar ve istenen durumu sağlar.

Bu cesur konulara daha sonra daha detaylı olarak gireceğiz ve ne olduklarını ve ne zaman kullanılacaklarını daha iyi anlamak için temel bir bilgi edineceğiz.

### Izleme

Burada her şey hızlı ilerliyor ve sürekli olarak yeni özellikler ve işlevlerle uygulamamızı güncelliyoruz. Testlerin uygulamada herhangi bir hata olmadığından emin olması gerekiyor. Üretim ortamında uygulama çalışıyor ve gereken yapılandırma ve performans sürekli olarak korunuyor olabilir.

Ancak şimdi, son kullanıcılarımızın ihtiyaç duydukları deneyimi aldığından emin olmamız gerekiyor. Uygulamamızın performansı sürekli olarak izlenmelidir. Bu aşama, geliştiricilerin gelecekteki uygulama sürümlerinde iyileştirmeler yapmaları için en iyi kararları vermelerine olanak sağlar.

Bu bölüm aynı zamanda uygulamada uygulanmış özellikler hakkındaki geri bildirimi ve son kullanıcıların uygulamayı nasıl geliştirmek istediğini yakalamamız gereken yerdir.

Güvenilirlik burada önemli bir faktördür, sonuçta uygulamamızın ihtiyaç duyulan süre boyunca her zaman kullanılabilir olmasını istiyoruz, yüksek erişilebilirliğe sahip olmak istiyoruz. Bu, sürekli olarak **gözlem, güvenlik ve veri yönetimi** gibi diğer alanların izlenmesini gerektirir. Geri bildirim her zaman iyileştirmek, güncellemek ve uygulamayı sürekli olarak yayınlamak için kullanılabilir.

Bu sürekli süreçte, [FinOps](https://www.ibm.com/es-es/topics/finops) ekipleri de dahil olmalıdır. Uygulamalar ve veriler çalıştırıldığı ve depolandığı bir yerde finansal takip gerektirir, böylece ekonomik kaynaklar açısından değişiklikler önemli mali sorunlara yol açmaz.

Toplulukla konuşurken, "DevOps Mühendisi" unvanının kimse için hedef olmaması gerektiği çünkü aslında herhangi bir pozisyonun DevOps süreçlerini ve burada açıklanan kültürü benimsemesi gerektiği belirtilmektedir. DevOps, bulut mühendisi/mimar, sanallaştırma yöneticisi, bulut mimarı/mühendisi ve altyapı yöneticisi gibi farklı pozisyonlarda kullanılmalıdır. Bunlar sadece bazı örneklerdir, ancak yukarıda "DevOps Mühendisi" terimini kullanma nedeni, yukarıdaki pozisyonlardan herhangi biri tarafından kullanılan sürecin kapsamını vurgulamaktır.

## Kaynaklar

Her zaman readme dosyalarında ek kaynaklar bırakılacak, çünkü burada öğrenmek için bulunuyoruz.

Bu belgeleri tam olarak takip etmek için bugün gördüklerimizin genel bir fikrini edinmeniz ve daha fazla derinlemesine araştırma yapmanız için aşağıdaki videoları izlemeniz ve makaleleri okumanız en iyisidir.
- [Continuous Development](https://www.youtube.com/watch?v=UnjwVYAN7Ns) 
- [Continuous Testing - IBM YouTube](https://www.youtube.com/watch?v=RYQbmjLgubM)
- [Continuous Integration - IBM YouTube](https://www.youtube.com/watch?v=1er2cjUq1UI)
- [Continuous Monitoring](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [The Remote Flow](https://www.notion.so/The-Remote-Flow-d90982e77a144f4f990c135f115f41c6)
- [FinOps Foundation - What is FinOps](https://www.finops.org/introduction/what-is-finops/)
- [**NOT FREE** The Phoenix Project: A Novel About IT, DevOps, and Helping Your Business Win](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/1942788290/)

Görüşmek üzere, [Gun 4](day04.md)'te buluşalım.
