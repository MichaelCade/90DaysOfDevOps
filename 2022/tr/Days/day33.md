## Microsoft Azure Ag(Network) Modelleri + Azure Yonetimi

Bugün Microsoft Azure'nun kuruluşunun ve 12. yıl dönümünün kutlandığı bir gün! (1 Şubat 2022) Neyse ki, Microsoft Azure'daki ağ modellerini ve Azure için bazı yönetim seçeneklerini ele alacağız. Şu ana kadar yalnızca Azure portalını kullandık, ancak platformdaki kaynaklarımızı yönetmek ve oluşturmak için başka alanlar olduğunu da belirttik.

## Azure Network(Ag) Modelleri

### Sanal Aglar(Virtual Networks)

- Sanal Network, Azure'da oluşturulan bir yapıdır.
- Bir sanal ağın bir veya daha fazla IP aralığı vardır.
- Sanal ağlar bir bölge içinde bir abonelikte bulunur.
- Sanal ağlarda, ağ aralığını bölmek için sanal alt ağlar oluşturulur.
- Sanal makineler sanal alt ağlara yerleştirilir.
- Bir sanal ağın içindeki tüm sanal makineler birbirleriyle iletişim kurabilir.
- Her bir Sanal Ağda 65.536 özel IP adresi kullanılabilir.
- Bir bölgeden çıkan trafiğin bedeli ödenir. (Bölgeden çıkan veri trafiği)
- IPv4 ve IPv6 desteklenir.
   - IPv6, halka açık ve sanal ağlar içinde kullanılabilir.

Azure Virtual Networks, AWS VPC'lerine benzetilebilir. Ancak, dikkate alınması gereken bazı farklılıklar vardır:

- AWS'de varsayılan bir VNet (Virtual Network) otomatik olarak oluşturulurken, Azure'da ilk sanal ağınızı gereksinimlerinize göre oluşturmanız gerekmektedir.
- Azure'da, varsayılan olarak tüm sanal makinelerin internete NAT (Network Address Translation) erişimi vardır. Azure, AWS'deki NAT Gateway'leri gibi NAT ağ geçitleri kullanmaz.
- Azure'da, AWS'deki gibi özel veya genel alt ağ kavramı bulunmaz.
- Genel IP adresleri, Azure'da vNIC'ler veya Yük Dengeleyicilere atanabilen bir kaynaktır.
- Sanal Ağlar ve alt ağlar, kendi ACL'lerine (Erişim Kontrol Listelerine) sahiptir. Bu, alt ağ düzeyinde yetkilendirme yapmanızı sağlar.
- Azure'da alt ağlar, AWS'deki gibi her bir Kullanılabilirlik Bölgesi için değil, Kullanılabilirlik Bölgeleri arasında yayılabilir.

Ayrıca, Azure'da Sanal Ağ Eşlemesi (Virtual Network Peering) özelliği bulunur. Bu özellik, kiracılar ve bölgeler arasında sanal ağları Azure altyapısı kullanarak birbirine bağlama imkanı sağlar. Bu özellik varsayılan olarak geçişken değildir, ancak hub sanal ağında Azure Güvenlik Duvarı kullanarak etkinleştirilebilir. Sanal ağların bağlantılı ağa bağlanabilmesini sağlamak için geçiş noktası transit kullanılır ve bunun bir örneği olarak ExpressRoute ile Kurumsal Ağa bağlantı kurulabilir.

### Access Control(Erişim Kontrolü)

- Azure, Ağ Güvenlik Gruplarını (Network Security Groups - NSG) kullanır ve bunlar durumları korur.
- Kuralların oluşturulmasına ve daha sonra bir ağ güvenlik grubuna atanmasına izin verir.
- Ağ güvenlik grupları, alt ağlara veya sanal makineler (VM) üzerine uygulanır.
- Bir alt ağa uygulandığında, bunun bir "Kenar" cihazı olmadığı Virtual Machine NIC'inde hala uygulanan bir kuraldır.

![](Images/Day33_Cloud1.png)

- Kurallar, bir Ağ Güvenlik Grubunda birleştirilir.
- Önceliğe göre esnek yapılandırmalar mümkündür.
- Daha düşük öncelik numarası yüksek önceliği ifade eder.
- Çoğu mantık IP adresleri üzerinde kurulmuştur, ancak bazı etiketler ve etiketler de kullanılabilir.

| Description      | Priority | Source Address     | Source Port | Destination Address | Destination Port | Action |
| ---------------- | -------- | ------------------ | ----------- | ------------------- | ---------------- | ------ |
| Inbound 443      | 1005     | \*                 | \*          | \*                  | 443              | Allow  |
| ILB              | 1010     | Azure LoadBalancer | \*          | \*                  | 10000            | Allow  |
| Deny All Inbound | 4000     | \*                 | \*          | \*                  | \*               | DENY   |

Ayrıca, Uygulama Güvenlik Grupları (Application Security Groups - ASG) mevcuttur.

- NSG'ler büyüyen ortamlar için sürdürmesi zor olabilecek IP adresi aralıklarına odaklanırken,
- ASG'ler, farklı uygulama rolleri için gerçek isimlerin (Monikers) tanımlanmasına olanak sağlar (Web Sunucuları, Veritabanı sunucuları, WebApp1 vb.).
- Sanal Makine NIC'i bir veya daha fazla ASG üyesi yapılır.

ASG'ler daha sonra Ağ Güvenlik Gruplarındaki kurallarda kullanılabilir ve iletişim akışını kontrol etmek için NSG hizmet etiketleri gibi NSG özelliklerini kullanabilir.

| Action | Name               | Source     | Destination | Port         |
| ------ | ------------------ | ---------- | ----------- | ------------ |
| Allow  | AllowInternettoWeb | Internet   | WebServers  | 443(HTTPS)   |
| Allow  | AllowWebToApp      | WebServers | AppServers  | 443(HTTPS)   |
| Allow  | AllowAppToDB       | AppServers | DbServers   | 1443 (MSSQL) |
| Deny   | DenyAllinbound     | Any        | Any         | Any          |

### Load Balancing(Yük Dengeleme)

Microsoft Azure, iki ayrı yük dengeleme çözümüne sahiptir (birinci taraf, Azure Marketplace'de üçüncü taraf seçenekler de mevcuttur). Her ikisi de dışarıya yönelik veya içerideki uç noktalarla çalışabilir.

- Yük Dengeleyici (Katman 4), karma-tabanlı dağıtım ve port yönlendirmeyi destekler.
- App Gateway (Katman 7), SSL dışa aktarımı, çerez tabanlı oturum tutma ve URL tabanlı içerik yönlendirmesi gibi özellikleri destekler.

Ayrıca, App Gateway ile isteğe bağlı olarak Web Uygulama Güvenlik Duvarı bileşenini kullanabilirsiniz.

## Azure Management Tools(Azure Yonetim Aracları)

Teorik zamanımızın çoğunu Azure Portal üzerinde dolaşarak geçirdik. Bir DevOps kültürünü ve sürecini takip etmek söz konusu olduğunda, özellikle dağıtım ile ilgili birçok görev API veya komut satırı aracılığıyla gerçekleştirilecektir. Azure ortamlarımızı otomatikleştirdiğimizde kullanılabilecek diğer yönetim araçlarından bazılarına da değinmek istedim.

### Azure Portal

Microsoft Azure Portal, komut satırı araçlarına alternatif olarak sunulan web tabanlı bir konsoldur. Azure Portal içinde aboneliklerinizi yönetebilirsiniz. Basit bir web uygulamasından karmaşık bulut dağıtımlarına kadar her şeyi oluşturabilir, yönetebilir ve izleyebilirsiniz. Portalda bulabileceğiniz diğer bir özellik de bu kırıntı izleri (breadcrumbs)dir. Daha önce belirtildiği gibi, JSON, tüm Azure Kaynaklarının temelini oluşturur. Özellikleri, hizmetleri ve işlevselliği anlamak için Portal'da başlayabilir, ardından otomatik iş akışlarınıza dahil etmek için altında yatan JSON'ı daha sonra anlayabilirsiniz.

![](Images/Day33_Cloud2.png)

Ayrıca Azure Önizleme portalı da bulunmaktadır. Bu portal, yeni ve yakında sunulacak hizmetleri ve geliştirmeleri görüntülemek ve test etmek için kullanılabilir.

![](Images/Day33_Cloud3.png)

### PowerShell

Azure PowerShell hakkında konuşmadan önce önce PowerShell'i tanıtmakta fayda var. PowerShell, görev otomasyonu ve yapılandırma yönetimi çerçevesi, bir komut satırı kabuğu ve bir betikleme dilidir. Bu, Linux bölümünde ele aldığımız kabuk betiklemeye benzetebiliriz. PowerShell ilk olarak Windows işletim sisteminde bulunuyordu, ancak şimdi çapraz platform destekli hale gelmiştir.

Azure PowerShell, Azure kaynaklarını doğrudan PowerShell komut satırından yönetmek için bir cmdlet kümesidir.

Aşağıdaki gibi görebiliriz ki, PowerShell komutu `Connect-AzAccount` kullanarak aboneliğinize bağlanabilirsiniz. 

![](Images/Day33_Cloud4.png)

Ardından, Azure sanal makinelerle ilişkili belirli komutları bulmak isteseydik, aşağıdaki komutu çalıştırabilirdik. Bu PowerShell programlama dilini daha fazla öğrenmek ve anlamak için saatler harcayabilirsiniz.

![](Images/Day33_Cloud5.png)

Microsoft tarafından PowerShell ile başlamanız ve hizmetleri yapılandırmanız için harika başlangıç kılavuzları bulunmaktadır [burada](https://docs.microsoft.com/en-us/powershell/azure/get-started-azureps?view=azps-7.1.0)

### Visual Studio Code

Birçok kişi gibi, ve hepinizin gördüğü gibi benim tercih ettiğim IDE Visual Studio Code'dur.

Visual Studio Code, Windows, Linux ve macOS için Microsoft tarafından geliştirilen ücretsiz bir kaynak kodu düzenleyicisidir.

Aşağıda, Microsoft Azure ve içindeki hizmetlerle etkileşimde bulunmak için kullanabileceğiniz birçok entegrasyon ve aracın Visual Studio Code'a entegre edildiğini göreceksiniz.

![](Images/Day33_Cloud6.png)

### Cloud Shell

Azure Cloud Shell, Azure kaynaklarını yönetmek için etkileşimli, kimlik doğrulaması yapılmış, tarayıcı üzerinden erişilebilen bir kabuktur. Çalışma şeklinize en uygun kabuk deneyimini seçme esnekliği sağlar.

![](Images/Day33_Cloud7.png)

Aşağıdaki resimden, portal içinde Cloud Shell'i ilk başlattığımızda Bash ve PowerShell arasında seçim yapabileceğimizi görebilirsiniz.

![](Images/Day33_Cloud8.png)

Cloud Shell'i kullanmak için aboneliğinizde biraz depolama alanı sağlamanız gerekecektir.

Cloud Shell'i seçtiğinizde, geçici olarak bir makine oluşturulur. Bu makineler geçici olmasına rağmen dosyalarınız iki şekilde saklanır: bir disk görüntüsü ve bağlanmış bir dosya paylaşımı aracılığıyla.

![](Images/Day33_Cloud9.png)

- Cloud Shell runs on a temporary host provided on a per-session, per-user basis
- Cloud Shell times out after 20 minutes without interactive activity
- Cloud Shell requires an Azure file share to be mounted
- Cloud Shell uses the same Azure file share for both Bash and PowerShell
- Cloud Shell is assigned one machine per user account
- Cloud Shell persists $HOME using a 5-GB image held in your file share
- Permissions are set as a regular Linux user in Bash

The above was copied from [Cloud Shell Overview](https://docs.microsoft.com/en-us/azure/cloud-shell/overview)

### Azure CLI

Son olarak, Azure CLI hakkında konuşmak istiyorum. Azure CLI, Windows, Linux ve macOS üzerine kurulabilir. Kurulum yapıldıktan sonra `az` komutunu kullanarak diğer komutlarla Azure kaynaklarını oluşturabilir, güncelleyebilir, silebilir ve görüntüleyebilirsiniz.

Azure öğrenmeye başladığımda Azure PowerShell ve Azure CLI arasındaki farkı anlamakta biraz karışıklık yaşamıştım.

Bu konuda topluluktan bazı geri bildirimler almak isterim. Ancak benim gördüğüm şekilde Azure PowerShell, Windows PowerShell veya PowerShell Core'a (Diğer bazı işletim sistemlerinde de mevcut, ancak tümü değil) eklenen bir modüldür. Öte yandan, Azure CLI, Azure'a bağlanan ve bu komutları yürüten çapraz platform destekli bir komut satırı programıdır.

Her iki seçenek de farklı bir sözdizimine sahiptir, ancak gördüğüm ve yaptığım kadarıyla çok benzer görevleri yerine getirebilirler.

Örneğin, PowerShell'den bir sanal makine oluştururken `New-AzVM` cmdlet'ini kullanırken, Azure CLI'da `az VM create` komutunu kullanırsınız.

Daha önce sistemime Azure PowerShell modülünü yüklediğimi, ancak aynı zamanda PowerShell üzerinden çağrılabilen Azure CLI'nin de yüklendiğini görmüştünüz.

![](Images/Day33_Cloud10.png)

Buradaki önemli nokta, doğru aracı seçmektir. Azure otomasyon üzerine kuruludur. Portal içinde yaptığınız her eylem, kaynakları okumak, oluşturmak, değiştirmek veya silmek için bir yerlerde yürütülen kodlara dönüşür.

Azure CLI

- Çapraz platformda çalışan komut satırı arabirimidir ve Windows, macOS, Linux'e yüklenebilir.
- Windows PowerShell, Cmd, Bash ve diğer Unix kabuklarında çalışır.

Azure PowerShell

- Cross Platform'da çalışan bir PowerShell modülüdür ve Windows, macOS, Linux'ta çalışır.
- Windows PowerShell veya PowerShell gerektirir.

Eğer ortamınızda PowerShell kullanamıyorsanız ancak bash veya başka bir kabuk kullanabiliyorsanız, Azure CLI tercih edeceğiniz seçenek olacaktır.

Şimdiye kadar geçtiğimiz teorileri alıp bazı senaryolar oluşturacak ve Azure'da uygulamalar gerçekleştireceğiz.

## Kaynaklar

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Gorusmek Uzere [Gun 34](day34.md)
