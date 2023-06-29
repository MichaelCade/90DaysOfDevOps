## DevOps & The Cloud

Bulut bilişim ve sunulan hizmetler, etik ve DevOps süreçleriyle çok iyi uyum sağlar. Bulut bilişimi, teknoloji ve hizmetler sunarken, daha önce birçok kez bahsettiğimiz gibi, DevOps süreci ve sürecin iyileştirilmesiyle ilgilenir.

Ancak başlangıçta bulut öğrenme yolculuğu dik bir öğrenme eğrisi gerektirir ve tüm bileşenleri veya doğru fiyat noktası için en iyi hizmeti seçmeyi anlamak ve anlamak önemlidir, bu da kafa karıştırıcı olabilir.

![](Images/Day28_Cloud1.png)

Halka açık bulut, bir DevOps zihniyeti gerektirir mi? Buradaki cevabım hayır, ancak gerçekten bulut bilişimden faydalanmak ve insanların büyük bulut faturalarından kaçınmasını sağlamak için Bulut Bilişim ve DevOps'u birlikte düşünmek önemlidir.

40.000 fitlik bir bakış açısından halka açık bulutun ne demek olduğuna baktığımızda, bazı sorumlulukların yönetilen bir hizmete aktarılmasıyla, sizin ve ekibinizin daha önemli konulara odaklanmasına olanak tanır, bu da uygulama ve nihai kullanıcılar olmalıdır. Sonuçta, halka açık bulut sadece başka bir kişinin bilgisayarından ibarettir.

![](Images/Day28_Cloud2.png)

Bu ilk bölümde, Public Bulut'un ne olduğunu ve genel olarak Public Bulut'a ilişkin yapı taşlarının biraz daha açıklamasına girmek istiyorum.

### SaaS

İlk olarak ele alınacak alan, yerinde çalışan bir hizmetin yönetim yükünün neredeyse tamamını ortadan kaldıran Yazılım olarak Hizmet (SaaS) alanıdır. E-posta için Microsoft Exchange'i düşünelim: Eskiden merkezi bir veri merkezinde veya belki de merdiven altındaki bir dolapta bulunan fiziksel bir sunucuydu. Bu sunucuyu beslemek ve sulamak gerekiyordu. Yani sunucuyu güncel tutmanız, sunucu donanımını satın almanız, muhtemelen işletim sistemini yüklemeniz, gerekli uygulamaları yüklemeniz ve ardından güncellemeleri takip etmeniz gerekiyordu. Bir şeyler ters giderse sorunları çözmeniz ve işleri tekrar çalışır hale getirmeniz gerekebilirdi.

ve verilerin yedeklenmesini de sağlamanız gerekecekti, ancak SaaS ile bu çoğunlukla değişmez.

SaaS ve özellikle Microsoft 365 (Exchange'i bahsettiğim için) yaptığı şey, bu yönetim yükünü ortadan kaldırmak ve e-posta aracılığıyla Exchange'in işlevselliğini sağlayan, aynı zamanda kullanıcıya genel olarak büyük bir deneyim sunan üretkenlik (Office 365) ve depolama (OneDrive) gibi diğer birçok seçeneği sunan bir hizmet sağlamaktır.

Salesforce, SAP, Oracle, Google ve Apple gibi diğer SaaS uygulamaları da geniş çapta kullanılmaktadır. Bunlar, yığını daha fazla yönetme yükünden kurtulmayı sağlar.

SaaS tabanlı uygulamalarla ilgili DevOps ve bir hikaye olduğundan eminim, ancak bunun ne olabileceğini anlamakta zorlanıyorum. Azure DevOps'un Microsoft 365 ile bazı büyük entegrasyonları olduğunu biliyorum, bu konuda bir göz atabilir ve bilgi edinebilirsiniz.

![](Images/Day28_Cloud3.png)

### Public Cloud

Ardından, public bulut hizmetiyle ilgili olarak, insanların bunu çeşitli farklı şekillerde düşünebileceklerini söyleyebiliriz. Bazıları bunu yalnızca Microsoft Azure, Google Cloud Platform ve AWS gibi büyük bulut sağlayıcıları olarak görebilir.

![](Images/Day28_Cloud4.png)

Bazıları, public bulutu, hiper ölçeklendiricileri yanı sıra dünya çapındaki binlerce HSP'yi de içeren daha geniş bir teklif olarak görür. Bu yazıda, hiper ölçeklendiricileri ve HSP'leri içeren genel bulutu ele alacağız, ancak daha sonra, temel bilgileri edinmek için bir veya daha fazla hiper ölçeklendiriciye daha spesifik olarak odaklanacağız.

![](Images/Day28_Cloud5.png)
_Bu listede daha fazla şirket olabilir, sadece çalıştığım ve bildiğim yerel, bölgesel, telekomünikasyon ve küresel markaları seçiyorum._

SaaS bölümünde, bulutun bir sistemin bazı parçalarını yönetme sorumluluğunu veya yükünü ortadan kaldırdığını söyledik. SaaS hakkında konuşursak, birçok soyutlama katmanını ortadan kaldırdığını görürüz, yani fiziksel sistemleri, ağı, depolamayı, işletim sistemini ve hatta belirli bir noktaya kadar uygulamayı ortadan kaldırır. Bulutla ilgili olarak, gereksinimlerimize bağlı olarak ortadan kaldırabileceğimiz veya koruyabileceğimiz çeşitli soyutlama seviyeleri bulunmaktadır.

SaaS'ı zaten bahsetmiştik, ancak genel bulutla ilgili olarak bahsetmemiz gereken en az iki katman daha bulunmaktadır.

Altyapı olarak hizmet: Bu katmanı bir sanal makine gibi düşünebilirsiniz, ancak yerinde sistemlerde fiziksel katmanla ilgilenmeniz gerekirken, bulutta bu sorumluluk bulut sağlayıcısının sorumluluğundadır ve siz istediğiniz işletim sistemi, veri ve uygulamaları yönetir ve yönetirsiniz.

Platform olarak hizmet: Bu katman, sorumlulukları ortadan kaldırmaya devam eder ve aslında sizin verileri ve uygulamayı kontrol etmenizi sağlar, ancak altta yatan donanım veya işletim sisteminden endişelenmenize gerek kalmaz.

aaS tekliflerinin yanı sıra birçok başka seçenek daha bulunmaktadır. StaaS (Depolama olarak Hizmet) gibi altta yatan donanımla ilgilenmenize gerek olmayan depolama hizmetleri sunan teklifler olabilir. Containers as a Service (CaaS) olarak adlandırılan bir hizmet hakkında duymuş olabilirsiniz, bu konuya daha sonra değineceğiz. Bir sonraki 7 gün boyunca ele alacağımız bir diğer aaS ise Functions as a Service (FaaS) olacak. FaaS, sürekli çalışan bir sistem yerine bir işlevin istediğiniz zaman ve şekilde çalışmasını istediğiniz bir hizmet sunabilir.

Genel bulut, geçmek istediğiniz kontrol soyutlama katmanlarını sağlayarak istediğiniz şekilde kullanmanıza olanak tanır.

![](Images/Day28_Cloud6.png)

### Private Cloud

Kendi veri merkezinize sahip olmak geçmişte kalmış bir şey değildir. Birçok şirket, OPEX modelini yönetmekte zorlandığı ve sadece genel bulutu kullanmayı öğrenmek konusunda yetkinlik setlerine sahip olmadığı için özel bulut modellerine yeniden ilgi göstermektedir.

Burada dikkate alınması gereken önemli bir nokta, genel bulutun artık sizin sorumluluğunuzda olacağı ve kendi tesislerinizde bulunacağıdır.

Bu alanda ilginç gelişmeler yaşanmaktadır. Sadece sanallaştırma dönemini ve yerel altyapı ortamlarını domine eden VMware ile değil, aynı zamanda hiper ölçeklendiricilerin genel bulutlarının yerel bir sürümünü sunanlarla da ilgilidir.

![](Images/Day28_Cloud7.png)

### Hibrit Cloud

Genel ve özel buluta yapılan atıflarla, her iki ortamı da kapsayarak ikisi arasında esneklik sağlayabiliriz. Belki genel bulutta sunulan hizmetlerden yararlanırken aynı zamanda yerel tesislerde bulunmanın özelliklerinden ve işlevselliğinden de yararlanabiliriz veya belirli bir düzenleme, verilerinizi yerel olarak depolamanızı gerektirebilir.

![](Images/Day28_Cloud8.png)

Tüm bunları bir araya getirdiğimizde, iş yüklerimizi depolamak ve çalıştırmak için birçok seçeneğimiz vardır.

![](Images/Day28_Cloud9.png)

Belirli bir hiper ölçeklendiriciye girmeden önce, Twitter'ın gücüne sordum: Nereye gitmeliyiz?

![](Images/Day28_Cloud10.png)

[Twitter anketine link](https://twitter.com/MichaelCade1/status/1486814904510259208?s=20&t=x2n6QhyOXSUs7Pq0itdIIQ)

Hangi hiper ölçeklendiricinin en yüksek yüzdeye sahip olduğuna bağlı olarak, o hiper ölçeklendiricinin sunduğu tekliflere daha derinlemesine bir dalış yapacağız. Ancak önemli bir nokta, tüm bu hizmetlerin oldukça benzer olduğudur. Bu yüzden biriyle başlamayı öneririm çünkü birinin temel bilgisini edinerek sanal makine oluşturma, ağ yapılandırması vb. konularında diğer hiper ölçeklendiricilere hızla adapte olabileceğimi gördüm.

Her durumda, üç hiper ölçeklendiriciyi kapsayan **ÜCRETSİZ** kaynaklar paylaşacağım.

Ayrıca, diğer bölümlerde olduğu gibi ilerledikçe bir senaryo oluşturacağımızı da belirtmek isterim.

## Kaynaklar

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Gorusmek Uzere [Gun 29](day29.md).
