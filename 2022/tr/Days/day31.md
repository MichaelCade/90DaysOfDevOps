## Microsoft Azure Hesaplama Modelleri

Microsoft Azure'daki güvenlik modellerine ilişkin temel kavramları incelemeye devam ederek, Azure'da mevcut olan farklı hesaplama hizmetlerine bakalım.

### Hizmet Yeterlilik Seçenekleri

Bu bölüm, verilerin yönetimi açısından yazar için önemlidir. On-premises ortamında olduğu gibi, hizmetlerinizin kullanılabilirliğini sağlamak kritik öneme sahiptir.

- Yüksek Kullanılabilirlik (Bir bölge içinde koruma)
- Felaket Kurtarma (Bir bölgeden diğerine koruma)
- Yedekleme (Zamana göre geri yükleme)

Microsoft, bir coğrafi sınırlar içinde birden fazla bölge dağıtmaktadır. Azure'da Hizmet Yeterlilik için iki kavram vardır: küme ve bölgeler:

- **Cluster** - Bir veri merkezi içinde dayanıklılık sağlar.
- **Zone'lar** - Bir bölge içindeki veri merkezleri arasında dayanıklılık sağlar.

### Sanal Makineler

Muhtemelen herkesin halka açık bir bulutta başlamak için tercih ettiği noktadır.

- Farklı kapasitelere (bazıları ezici olabilir) sahip çeşitli sanal makine serileri ve boyutları sağlar. [Azure'daki Sanal Makine Boyutları](https://docs.microsoft.com/en-us/azure/virtual-machines/sizes)
- Yüksek performanslı ve düşük gecikmeli iş yüklerinden, yüksek bellekli seçeneklere kadar farklı MV seçenekleri ve yaklaşımları bulunmaktadır.
- Ayrıca, B Serisi altında bulunan patlayıcı bir MV türü vardır. Bu, genellikle çoğunlukla düşük CPU gereksinimine sahip olabilecek iş yükleri için idealdir, ancak aylık olarak bir performans pik gereksinimi olabilir.
- Sanal makineler, herhangi bir ağa bağlantı sağlayabilen bir sanal ağa yerleştirilir.
- Windows ve Linux gibi misafir işletim sistemlerini destekler.
- Ayrıca, belirli Linux dağıtımları söz konusu olduğunda Azure için uyarlanmış çekirdekler bulunmaktadır. [Azure Tuned Kernels](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/endorsed-distros#azure-tuned-kernels)

### Şablonlar

Daha önce belirtildiği gibi, Microsoft Azure'nun altında JSON tabanlı bir yapı bulunmaktadır.

Kaynaklarımızı oluşturmak için kullanabileceğimiz çeşitli yönetim portal ve konsolları bulunmaktadır, ancak tercih edilen yöntem JSON şablonları aracılığıyla oluşturmaktır.

İstenen durumu tekrarlayabilir hale getiren, artan veya tamamlanmış durumda idempotent dağıtımlar sağlar.

Dağıtılmış kaynak tanımlarını dışa aktarabilen geniş bir şablon seçeneği mevcuttur. Bu şablon özelliğini, AWS CloudFormation'a veya çoklu bulut seçeneği için Terraform'a benzetmekten hoşlanırım. Altyapı olarak kodlama bölümünde Terraform'ın gücü hakkında daha fazla bilgi vereceğiz.

### Ölçeklendirme

Otomatik ölçeklendirme, halka açık bulutun önemli bir özelliğidir, kullanılmayan kaynakları azaltma veya ihtiyaç duyulduğunda artırma yeteneğine sahip olmak.

Azure'da IaaS için Virtual Machine Scale Sets (VMSS) adı verilen bir özellik bulunmaktadır. Bu, bir standart altın imajını otomatik olarak oluşturmanıza ve ölçeklendirmenize olanak tanır, bu işlemi saatlere ve metriklere dayalı olarak gerçekleştirir.

Bu, güncelleme pencereleri için idealdir, böylece görüntülerinizi güncelleyebilir ve en az etkiyle dağıtabilirsiniz.

Azure App Services gibi diğer hizmetlerde yerleşik otomatik ölçeklendirme özelliğine sahiptir.

### Konteynerler

Henüz konteynerleri bir kullanım senaryosu olarak ele almadık ve DevOps öğrenme yolculuğumuzda neden ve nasıl gerektiklerini tartışmadık, ancak Azure'nun değerli birkaç konteyner odaklı hizmeti olduğunu belirtmeliyiz.

- **Azure Kubernetes Service (AKS)** - Kontrol düzlemi veya alt küme yönetimiyle uğraşmadan yönetilen bir Kubernetes çözümü sağlar. Kubernetes hakkında daha fazlasını ileride göreceğiz.

- **Azure Container Instances** -  Saniyelik faturalandırma ile bir Konteyner Hizmeti sunar. Bir görüntüyü çalıştırabilir ve sanal ağınızla entegre edebilirsiniz, Konteyner Yönetimi gerektirmez.

- **Service Fabric** - Birçok özelliği içerir, ancak konteyner örneklerini yönetmek için orkestrasyon sağlar.

Azure ayrıca Docker Görüntüleri, Helm grafikleri, OCI Sanatçıları ve görüntüler için özel bir Kayıt Defteri sağlar. Bu konuda daha fazlasını ilgili konteyner bölümüne geldiğimizde göreceğiz.

Ayrıca, birçok konteyner hizmetinin altında konteynerlerin kullanılabileceğini de belirtmeliyiz, ancak bunu yönetme ihtiyacınızdan soyutlarlar.

Bu konteyner odaklı hizmetlerin yanı sıra, diğer tüm halka açık bulutlarda da benzer hizmetler bulunmaktadır.

### Uygulama Hizmetleri

- Azure Application Services, hizmetlerinizi kurmanın kolay bir yolunu sağlayan bir uygulama barındırma çözümü sunar.
- Otomatik dağıtım ve ölçeklendirme.
- Windows ve Linux tabanlı çözümleri destekler.
- Hizmetler, bir Uygulama Hizmeti Planı içinde çalışır ve bir tür ve boyuta sahiptir.
- Web uygulamaları, API uygulamaları ve mobil uygulamalar dahil olmak üzere birçok farklı hizmet.
- Güvenilir test ve dağıtım için Dağıtım Yuvaları desteği.

### Sunucusuz Hesaplama

Sunucusuz hesaplama ile amacımız, yalnızca işlevin çalışma süresi için ödeme yapmak ve sürekli olarak çalışan sanal makineler veya PaaS uygulamalarına sahip olmamaktır. İşlevimizi sadece ihtiyaç duyduğumuzda çalıştırır ve ardından ortadan kalkar.

**Azure Fonksiyonları** - Sunucusuz bir kod sağlar. Genel bulut ortamına ilk bakışımızı hatırlarsak, yönetim katmanının soyutlama seviyesini hatırlayacağız. Sunucusuz işlevlerle, yalnızca kodu yönetmeniz gerekecektir.

**Event Driven** ,ölçeklenebilir. Azure ve üçüncü taraf hizmetlerine giriş ve çıkış bağlantısı sağlar.

Çok çeşitli programlama dillerini destekler. (C#, NodeJS, Python, PHP, batch, bash, Golang ve Rust. Herhangi bir yürütülebilir dosya)

**Azure Event Grid** hizmetlerden ve olaylardan mantığı tetiklemenizi sağlar.

**Azure Logic App** grafik tabanlı iş akışları ve entegrasyon sağlar.

Ayrıca, tutarlı yönetim ve programlama ile Windows ve Linux düğümlerinde büyük ölçekli işleri çalıştırabilen Azure Batch'e de göz atabiliriz.

## Kaynaklar

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Gorusmek Uzere [Gun 32](day32.md).
