## Microsoft Azure Senaryoları ve Uygulamaları 

Son 6 gün, Microsoft Azure ve genel olarak halka açık buluta odaklandı. Azure'nin temel yapı taşlarını anlamak için bu temelin biraz teorik bilgi içermesi gerekiyordu, ancak bu bilgi, diğer büyük bulut sağlayıcılarında da geçerli olacaktır.

Başlangıçta halka açık bulutun temel bilgisini edinmek ve en azından başlangıçta bir sağlayıcı seçmekten bahsettim. Farklı bulutlar arasında gidip gelirseniz, kolayca kaybolabileceğinizi düşünüyorum. Ancak bir tane seçtiğinizde temelleri anlarsınız ve bu temelleri anladığınızda diğer bulutlara geçmek ve öğrenmeyi hızlandırmak oldukça kolay olur.

Bu final oturumunda, Microsoft tarafından oluşturulan ve[AZ-104 Microsoft Azure Administrator](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/) sınavına hazırlık için kullanılan bir referans olan bu sayfadan kendi uygulamalı senaryolarımı seçeceğim.

Şu anda ayrıntılı olarak ele almadığımız, Konteynerler ve Kubernetes gibi bazı konular burada bulunmaktadır, bu yüzden henüz oraya atlamak istemiyorum.

Daha önceki yazılarda, Modül 1, 2 ve 3'ün çoğunu oluşturduk.

### Sanal Ag(Virtual Networking)

Takip edilerek [Modül 04](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_04-Implement_Virtual_Networking.html):

Yukarıdakileri geçtim ve bazı adları #90DaysOfDevOps için değiştirdim. Ayrıca, Cloud Shell'i kullanmak yerine Windows makinedeki Azure CLI ile önceki günlerde oluşturduğum yeni kullanıcımla giriş yaptım.

Bunu, tarayıcıyı açacak ve hesabınıza kimlik doğrulama yapmanıza izin verecek `az login` komutunu kullanarak yapabilirsiniz.

Daha sonra PowerShell betiği oluşturdum ve bazı referanslarını bu klasöre yerleştirerek aşağıdaki görevleri oluşturmak için modülden yararlandım.(Cloud\01VirtualNetworking)

Lütfen scriptte dosya konumunu kendi ortamınıza uyacak şekilde değiştirin.

Bu ilk aşamada, ortamımızda sanal ağ veya sanal makine oluşturulmadı. Yalnızca kaynak grubumdaki bir Cloud Shell depolama konumu yapılandırdım.

Öncelikle [PowerShell script](Cloud/01VirtualNetworking/Module4_90DaysOfDevOps.ps1) çalıştırıyorum.

![](Images/Day34_Cloud1.png)

- Task 1: Sanal ağ oluşturma ve yapılandırma

![](Images/Day34_Cloud2.png)

- Task 2: Sanal ağa sanal makinelerin dağıtılması

![](Images/Day34_Cloud3.png)

- Task 3: Azure VM'lerin özel ve genel IP adreslerini yapılandırma

![](Images/Day34_Cloud4.png)

- Task 4: Ağ güvenlik gruplarını yapılandırma

![](Images/Day34_Cloud5.png)
![](Images/Day34_Cloud6.png)

- Task 5: İç isim çözümü için Azure DNS'yi yapılandırma

![](Images/Day34_Cloud7.png)
![](Images/Day34_Cloud8.png)

### Network Trafik Yönetimi

Takip edilerek [Modül 06](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_06-Implement_Network_Traffic_Management.html):

Sonraki adımda, bir öncekinden başlayarak kaynak grubumuza gittik ve kaynaklarımızı sildik. Kullanıcı hesabınızı yalnızca o kaynak grubuna erişime sahip olacak şekilde ayarlamamışsanız, modülü takip ederek adı `90Days*` olarak değiştirerek tüm kaynakları ve kaynak grubunu silebilirsiniz. Bu, aşağıdaki lablarda her bir işlemim için sürecim olacak.

Bu lab için, ayrıca aşağıdaki görevleri oluşturmak için PowerShell betiği ve modülden bazı referansları kullanarak ilgili dosyaları bu klasörde buldum.
(Cloud\02TrafficManagement)

- Task 1: Laboratuvar ortamının sağlanması

Oncelikle [PowerShell script](Cloud/02TrafficManagement/Mod06_90DaysOfDevOps.ps1) calıstırıyorum.

![](Images/Day34_Cloud9.png)

- Task 2:  Hub ve spoke ağ topolojisinin yapılandırılması

![](Images/Day34_Cloud10.png)

- Task 3: Sanal ağ peering'inin geçirgenliğini test etme

90DaysOfDevOps grubumun Network Watcher'a erişimi olmadığı için bunun nedeni, Network Watcher'ların kaynak grubuna bağlı olmayan kaynaklardan biri olmasıdır. Bu nedenle, bu kullanıcı için RBAC'nin kapsandığı kaynak grubu değildir. 90DaysOfDevOps grubuna East US Network Watcher katkıda bulunan rolünü ekledim.

![](Images/Day34_Cloud11.png)
![](Images/Day34_Cloud12.png)
![](Images/Day34_Cloud13.png)

^  Bu, iki spoke sanal ağı birbirine bağlamadığımız için (sanal ağ peering geçişken değil) beklenen bir durumdur.

- Task 4: Hub ve spoke topolojisinde yönlendirme yapılandırması

Burada, hesabımın 90DaysOfDevOps grubu içindeki kullanıcı olarak komut çalıştıramamasıyla ilgili başka bir sorun yaşadım. Bunun nedenini anlamak için emin değilim, bu yüzden ana yönetici hesabıma geri döndüm. 90DaysOfDevOps grubu, 90DaysOfDevOps Kaynak Grubu'ndaki her şeyin sahibidir, bu yüzden VM içinde bir komut çalıştıramıyor olmamı anlamak isterim.

![](Images/Day34_Cloud14.png)
![](Images/Day34_Cloud15.png)

Daha sonra michael.cade@90DaysOfDevOps.com hesabıma geri döndüm ve bu bölümü tamamladım. Burada, aynı testi tekrar çalıştırıyoruz ancak bu kez sonuç ulaşılabilir olacak şekilde.

![](Images/Day34_Cloud16.png)

- Task 5: Azure Load Balancer Uygulama 

![](Images/Day34_Cloud17.png)
![](Images/Day34_Cloud18.png)

- Task 6: Azure Uygulama Gateway Uygulama

![](Images/Day34_Cloud19.png)
![](Images/Day34_Cloud20.png)

### Azure Storage(Depolama)

Takip edilerek [Modül 07](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_07-Manage_Azure_Storage.html):

Bu lab için, ayrıca aşağıdaki görevleri oluşturmak için PowerShell betiği ve modülden bazı referansları kullanarak ilgili dosyaları bu klasörde buldum.

(Cloud\03Storage)

- Task 1: Laboratuvar ortamının sağlanması

Öncelikle [PowerShell script](Cloud/03Storage/Mod07_90DaysOfDeveOps.ps1)  scriptini çalıştırıyorum. 

![](Images/Day34_Cloud21.png)

- Task 2: Azure Depolama hesaplarının oluşturulması ve yapılandırılması

![](Images/Day34_Cloud22.png)

- Task 3: Blob depolamasını yönetme

![](Images/Day34_Cloud23.png)

- Task 4: Azure Depolama için kimlik doğrulama ve yetkilendirme yönetimi

![](Images/Day34_Cloud24.png)
![](Images/Day34_Cloud25.png)

Buna izin verilmesini beklerken biraz sabırsızdım, ancak sonunda çalıştı.

![](Images/Day34_Cloud26.png)

- Task 5: Azure Dosyalar paylaşımlarının oluşturulması ve yapılandırılması

Komut çalıştırma işleminde, michael.cade@90DaysOfDevOps.com ile çalışmıyordu, bu yüzden yönetici hesabımı kullandım.

![](Images/Day34_Cloud27.png)
![](Images/Day34_Cloud28.png)
![](Images/Day34_Cloud29.png)

- Task 6: Azure Depolama için network erişimini yönetme

![](Images/Day34_Cloud30.png)

### Serverless (Web Uygulamaları Uygulama)

Takip ederek [Modül 09a](https://microsoftlearning.github.io/AZ-104-MicrosoftAzureAdministrator/Instructions/Labs/LAB_09a-Implement_Web_Apps.html):

- Task 1: Bir Azure web uygulaması oluşturma

![](Images/Day34_Cloud31.png)

- Task 2: Bir aşama deployment slots oluşturma

![](Images/Day34_Cloud34.png)

- Task 3: Web uygulaması deployment ayarlarını yapılandırma

![](Images/Day34_Cloud33.png)

- Task 4: Kodu aşama deployment slots dağıtma

![](Images/Day34_Cloud32.png)

- Task 5: Aşama slots değiştirme

![](Images/Day34_Cloud35.png)

- Task 6: Azure web uygulamasının otomatik ölçeklendirmesini yapılandırma ve test etme

Bu senaryoda kullandığım scripti (Cloud/04Serverless) dizininde bulunabilir.

![](Images/Day34_Cloud36.png)

Bu, Microsoft Azure ve genel olarak halka açık bulut hakkındaki bölümü sonlandırıyor. Bu senaryolara saldırmak ve üzerlerinde çalışmak gerçekten eğlenceliydi, diyebilirim.

## Kaynaklar

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Sonraki olarak, sürüm kontrol sistemlerine, özellikle de git'e ve kod deposu genel bakışlarına odaklanacağız ve tercih ettiğim seçenek olarak GitHub'ı seçeceğiz.

Gorusmek Uzere [Gun 35](day35.md)
