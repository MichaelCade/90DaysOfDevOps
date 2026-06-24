## Microsoft Azure Storage(Depolama) Modelleri

### Storage(Depolama) Hizmetleri

- Azure depolama hizmetleri, depolama hesapları aracılığıyla sağlanır.
- Depolama hesapları genellikle REST API üzerinden erişilir.
- Bir depolama hesabının benzersiz bir adı olmalıdır ve DNS adının bir parçası olmalıdır:  `<Depolama(Storage) Hesap Adı>.core.windows.net`
- Çeşitli kopyalama ve şifreleme seçenekleri bulunur.
- Bir kaynak grubu içinde yer alır.

Azure Portal'ın üstündeki arama çubuğunda Sosyal Grup araması yaparak depolama grubumuzu kolayca oluşturabiliriz.

![](Images/Day32_Cloud1.png)

Daha sonra depolama hesabımızı oluşturmak için adımları takip edebiliriz. Bu adın benzersiz olması ve tamamen küçük harf olması gerektiğini unutmamız gerekir. Rakamları içerebilir ancak boşluk bulunmamalıdır.

![](Images/Day32_Cloud2.png)

Depolama hesabımıza karşı tercih ettiğimiz kopyalama seviyesini ve burada depoladığımız herhangi bir şeyi seçebiliriz. Liste aşağı indikçe maliyeti daha yüksek seçenekler ve verilerinizin yayılımı sağlanır.

Varsayılan kopyalama seçeneği bile verilerimizin 3 kopyasını sunar.

[Azure Depolama Kopyalama](https://docs.microsoft.com/en-us/azure/storage/common/storage-redundancy)

Yukarıdaki bağlantının özeti aşağıda verilmiştir:

- **Locally-redundant storage** -  verilerinizi birincil bölgedeki tek bir veri merkezinde üç kez replike eder.
- **Geo-redundant storage** - verilerinizi birincil bölgedeki tek bir fiziksel konumda LRS kullanarak senkron olarak üç kez kopyalar.
- **Zone-redundant storage** - Azure Depolama verilerinizi birincil bölgedeki üç Azure kullanılabilirlik bölgesi arasında senkron olarak replike eder.
- **Geo-zone-redundant storage** - coğrafi replikasyon tarafından sağlanan bölgesel kesintilere karşı korumayla birlikte kullanılabilirlik bölgeleri arasında sağladığı yüksek kullanılabilirlik kombinasyonunu sunar. Bir GZRS depolama hesabında verileriniz birincil bölgedeki üç Azure kullanılabilirlik bölgesine kopyalanır ve aynı zamanda bölgesel felaketlere karşı koruma için ikinci bir coğrafi bölgeye replike edilir.

![](Images/Day32_Cloud3.png)

Performans seçeneklerine geri dönersek, Standard ve Premium seçeneklerimiz var. Yürüyüşümüzde Standard'ı seçtik, ancak premium size bazı özel seçenekler sunar.

![](Images/Day32_Cloud4.png)

Ardından açılır menüde, aşağıdaki üç seçeneği görebilirsiniz.

![](Images/Day32_Cloud5.png)

Depolama hesabınız için daha fazla gelişmiş seçenek mevcuttur, ancak şu anda bu alanlara girmemize gerek yok. Bu seçenekler şifreleme ve veri korumasıyla ilgilidir.

### Yönetilen Diskler

Depolama erişimi farklı yöntemlerle sağlanabilir.

Kimlik doğrulamalı erişim:

- Tam kontrol için paylaşılan anahtar.
- Delege edilmiş, ayrıntılı erişim için Paylaşılan Erişim İmzası.
- Azure Active Directory (Mevcut olduğunda)

Public Erisim:

- Genel erişim, HTTP üzerinden de dahil olmak üzere anonim erişime izin vermek için sağlanabilir.
- Bir örnek olarak, temel içeriği ve dosyaları bir blok blobunda barındırarak tarayıcının bu verileri görüntülemesine ve indirmesine izin verebilirsiniz.

Depolamanıza başka bir Azure hizmetinden erişiyorsanız, trafiğiniz Azure içinde kalır.

Depolama performansı konusunda iki farklı türümüz bulunmaktadır:

- **Standard** - Maksimum IOPS sayısı
- **Premium** - Garantili IOPS sayısı

IOPS =>  Saniyedeki Giriş/Çıkış işlemi

Ayrıca, doğru depolamayı seçerken dikkate almanız gereken yönetilmeyen ve yönetilen diskler arasında da bir fark vardır.

### Sanal Makine Storage(Depolama)

- Sanal Makine İşletim Sistemi diskleri genellikle kalıcı depolamada saklanır.
- Bazı durumsuz iş yükleri kalıcı depolamaya ihtiyaç duymaz ve düşük gecikme süresi daha büyük bir avantajdır.
- Ephemeral OS yönetilen diskleri destekleyen bazı sanal makineler bulunmaktadır ve bunlar düğeme yerel depolamada oluşturulur.
  - Bunlar aynı zamanda VM Scale Set'lerle de kullanılabilir.

Yönetilen Diskler, Azure Sanal Makinelerle kullanılabilen dayanıklı blok depolamadır. Ultra Disk Depolama, Premium SSD, Standart SSD veya Standart HDD şeklinde kullanılabilir. Ayrıca bazı özelliklere sahiptirler.

- Anlık görüntü (snapshot) ve imaj (image) desteği
- SKU'lar arasında kolay hareket ettirme
- Kullanılması durumunda daha iyi kullanılabilirlik sağlar.
- Faturalandırma, tüketilen depolama yerine disk boyutuna göre yapılır.

## Arşiv Depolama

- **Cool Tier** - Blok ve ekleme blobları için bir soğuk katman depolama mevcuttur.
  - Daha düşük depolama maliyeti
  - Daha yüksek işlem maliyeti.
- **Archive Tier** - Blok BLOB'ları için arşiv depolama mevcuttur.
  - Bu, BLOB başına yapılandırılır.
  - Daha ucuz maliyet, daha uzun veri alım gecikmesi.
  - Düzenli Azure Depolaması ile aynı Veri Dayanıklılığı.
  - İhtiyaç duyulduğunda özel Veri katmanlaması etkinleştirilebilir.

### Dosya Paylasımı

Yukarıdaki depolama hesabı oluşturma adımından sonra artık dosya paylaşımları oluşturabiliriz.

![](Images/Day32_Cloud6.png)

Bu, Azure'da SMB2.1 ve 3.0 dosya paylaşımları sağlar.

Azure içinde kullanılabilir ve SMB3 aracılığıyla internet üzerinden erişilebilir (port 445 açık).

Azure'da paylaşılan dosya depolaması sağlar.

REST API'ye ek olarak standart SMB istemcileri kullanılarak eşlenebilir.


Ayrıca [Azure NetApp Dosyaları](https://vzilla.co.uk/vzilla-blog/azure-netapp-files-how) (SMB and NFS) da fark edebilirsiniz. 

### Caching & Media Servisleri

Azure İçerik Dağıtım Ağı, dünya çapında konumlandırılmış statik web içeriği için bir önbellek sağlar.

Azure Medya Hizmetleri, oynatma hizmetlerine ek olarak medya dönüştürme teknolojileri sunar.

## Microsoft Azure Veritabanı Modelleri

[Gun 28](day28.md)'de çeşitli hizmet seçeneklerini ele aldık. Bunlardan biri PaaS (Platform olarak Hizmet) idi, burada büyük bir kısmı altyapı ve işletim sisteminden soyutlanır ve uygulama veya bu durumda veritabanı modellerinin kontrolü size bırakılır.

### İlişkisel Veritabanları

Azure SQL Veritabanı, Microsoft SQL Server'a dayalı olarak sunulan bir ilişkisel veritabanı hizmeti sağlar.

Bu, belirli bir işlevsellik sürümü gerektiğinde mevcut olan en son SQL dalında çalışan SQL'dir.

Bu yapılandırılabilir birkaç seçenek bulunur. Tek bir veritabanı sağlayabiliriz, bu durumda örnekte bir veritabanını temsil ederken, esnek bir havuz birden çok veritabanını paylaşan ve birlikte ölçeklenebilen bir kapasite havuzunu sağlar.

Bu veritabanı örneklerine, normal SQL örnekleri gibi erişilebilir.

MySQL, PostgreSQL ve MariaDB için ek yönetilen sunumlar mevcuttur.

![](Images/Day32_Cloud7.png)

### NoSQL Çözümleri

Azure Cosmos DB, şema bağımsız bir NoSQL uygulamasıdır.

99.99% hizmet düzeyi anlaşması (SLA)

Otomatik yönlendirme ile dünya genelinde herhangi bir yerde %99'luk dilimlerde tek basamaklı gecikmelerle dağıtılan bir veritabanı.

Verilerin bölümlenmesi/parçalanması/distribütörü için bölüm anahtarı kullanılır.

Farklı veri modellerini destekler (belgeler, anahtar-değer, grafik, sütun dostu).

Farklı API'leri destekler (DocumentDB SQL, MongoDB, Azure Tablo Depolama ve Gremlin).


![](Images/Day32_Cloud9.png)

CAP teoremine dayanan farklı tutarlılık modelleri mevcuttur.  [CAP teorem](https://en.wikipedia.org/wiki/CAP_theorem) temel alınır.

![](Images/Day32_Cloud8.png)

### Caching (Onbellekleme)

Redis gibi önbellekleme sistemlerine detaylı olarak girmeden, Microsoft Azure'ın Azure Cache for Redis adında bir hizmete sahip olduğunu belirtmek istedim.

Azure Cache for Redis, Redis yazılımına dayalı bir bellek içi veri deposu sağlar.

- Bu, açık kaynak Redis Cache'in bir uygulamasıdır.
  - Barındırılan, güvenli bir Redis Cache örneği.
  - Farklı seviyelerde bulunur.
  - Cache kullanmak için uygulama güncellenmelidir.
  - Yazma işlemlerine kıyasla yüksek okuma gereksinimine sahip bir uygulama için amaçlanmıştır.
  - Key-value tabanlı bir depodur.

![](Images/Day32_Cloud10.png)

Son birkaç gündeki Microsoft Azure hakkındaki not alma ve teorik bilgilerin oldukça yoğun olduğunun farkındayım, ancak bu bileşenlerin nasıl bir araya geldiği ve çalıştığına dair pratik yönlerine geçmeden önce temel kavramları ele almak istedim.

Hizmetlerin senaryo tabanlı dağıtımlarını çalışır duruma getirmeden önce, ağlarla ilgili biraz daha teorik bilgiye ihtiyacımız var. Ayrıca, şu ana kadar kullandığımız portal dışında Microsoft Azure ile etkileşim kurmanın farklı yollarına da bir göz atmak istiyoruz.

## Kaynaklar

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Gorusuruz [Gun 33](day33.md)
