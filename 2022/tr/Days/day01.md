## Giris - Gun 1

90 günlük maceramıza hoş geldiniz! Bu süre zarfında DevOps'un temel anlayışını öğrenmek ve birçok DevOps aracını keşfetmek için çıkıyoruz.

Bu öğrenme yolculuğu yıllar önce yazar [Michael Cade](https://github.com/MichaelCade) için başladı. Başlangıçta, sanallaştırma platformları ve bulut tabanlı teknolojilere odaklansa da, özellikle Altyapıyı Kodlama (IaC) ve Terraform ve Chef ile uygulama yapılandırma yönetimi konularına ilgi duyuyordu.

2021 yılında, Kasten by Veeam'da Cloud Native stratejisine odaklanma fırsatını elde etti. Bu fırsat, Kubernetes ve DevOps ile ilgili büyük bir odaklanma ve bu teknolojileri çevreleyen toplulukla etkileşim kurma fırsatı sunuyordu. Öğrenme yolculuğuna başladı ve hızla Kubernetes ve konteynerleme temellerini öğrenmekten daha geniş bir dünyanın olduğunu fark etti. Bunun üzerine toplulukla iletişime geçmeye ve DevOps kültürü, araçları ve süreçleri hakkında daha fazla bilgi edinmeye başladı. Bu süreçte öğrendiklerini halka açık bir şekilde belgelemeye başladı.


[Öyleyse DevOps öğrenmek istiyor musun?](https://blog.kasten.io/devops-learning-curve)

## Yolculuk Başlasın

Onceki blog yazısını okuduğunuzda, bu içeriğin öğrenme yolculuğu için yüksek seviyeli bir içerik olduğunu ve yazarın hiçbir bölümde uzman olmadığını göreceksiniz. Ancak amacının, her iki durumu da dikkate alarak ücretsiz ve bazı ücretli kaynakları paylaşmak olduğunu belirtmek istedi. Çünkü hepimizin farklı koşulları olduğunu unutmamak önemlidir.

Gelecek 90 gün boyunca, bu kaynakların belgelerini takip edebilir ve bu temel alanları kapsayabilirsiniz. Toplumun da bu sürece dahil olması harika olurdu. Yolculuğunuzu ve kaynaklarınızı paylaşarak hep birlikte kamuya açık bir şekilde öğrenelim ve birbirimize yardımcı olalım: [Repo original](https://github.com/MichaelCade/90DaysOfDevOps).

Projenin depo başlangıç [README](../README.md) dosyasında, her alanın toplamda 12 hafta ve 6 gün boyunca bölümlere ayrıldığını göreceksiniz. İlk 6 gün boyunca, öncelikle genel olarak DevOps temellerini keşfedeceksiniz, daha sonra belirli bazı alanlara derinlemesine gireceğiz. Bu liste kesinlikle eksiksiz değildir ve tekrar söylemek gerekirse, topluluğun bu belgeyi daha da geliştirmek için yardımcı olması harika olurdu.

Bu noktada, DevOps felsefesine ilgi duyan herkes için ilginç bir kaynak daha [DevOps Roadmap](https://roadmap.sh/devops) olabilir. Şu anda bulunduğunuz konumu ve ulaşılması gereken hedefleri akılda canlandırmak için bir göz atmanızı öneririm.

![DevOps Roadmap 2022](https://gitea.vergaracarmona.es/manuelver/awesome-roadmaps/media/branch/main/img/devops.png)

Bu, başlamak için harika bir kaynak ve Michael Cade'in başlangıç listesini ve blog girişini oluşturmak için temel aldığı bir kaynaktır. Ayrıca, bu depoda listelenen 12 konunun dışında daha detaylı diğer alanları da görebilirsiniz.

## İlk Adımlar - DevOps Nedir?

Burada sıralanabilecek birçok blog yazısı ve YouTube videoları var, ancak 90 günlük bir meydan okuma başlatacağımız için her gün yaklaşık bir saatini yeni bir şeyler öğrenerek DevOps hakkında bilgi edinmeye odaklanacağız. Yine de, başlamak için "DevOps nedir" konusunda bazı yüksek düzey bilgilere sahip olmak uygun olacaktır.

Öncelikle, **DevOps bir araç değildir**. Satın alınamaz, bir yazılım SKU'su veya indirilebilen bir açık kaynak GitHub deposu değildir. Aynı zamanda bir programlama dili de değildir. Ve tabii ki, bir sihirli karanlık sanat da değildir, ancak bilgisayar bilimi alanında sık sık böyle bir algı oluşmaktadır.

**DevOps, Yazılım Geliştirme sürecinde işleri daha akıllıca yapmanın bir yoludur.** - Peki... Yazılım geliştiricisi değilseniz, hemen geri adım atıp bu projeye dalış yapmamalı mısınız? Kesinlikle hayır. Kalın... Çünkü DevOps, **yazılım geliştirme ve operasyonları birleştirir**. Daha önce sanal makinelerin tarafında olduğu ve genellikle operasyonların tarafına düştüğü belirtilmişti, ancak topluluk içinde farklı uzmanlık alanlarına sahip insanlar vardır. Kesin olan şey, DevOps'un teknoloji dünyasında herkesin faydalanabileceği, geliştiricilerin, operasyonların veya QA mühendislerinin olabileceği yönünde. Herkes, DevOps'u daha iyi anlamanın sunduğu iyi uygulamaları eşit olarak öğrenebilir.

DevOps, bir ürünün fikir aşamasından son kullanıcıya, kim olursa olsun (bir iç ekip, müşteriler, topluluk vb.) üretimdeki dağıtım süresini azaltmaya yardımcı olan bir dizi uygulamadır.

Bu ilk haftada bu yolculuğun üzerinde durduğu bir diğer alan, **Çevik metodoloji** hakkında. DevOps ve Çevik, **Uygulama**nın sürekli teslimini gerçekleştirmek için yaygın olarak birlikte kullanılır.

Yüksek düzeydeki fikir, bir DevOps zihniyetinin veya kültürünün, uzun ve karmaşık yazılım sürüm süreçlerini azaltarak daha küçük ve daha sık teslim edilebilir sürümler yapma amacıyla çalıştığıdır. Burada anlaşılması gereken başka önemli bir nokta da, önceden bahsedilen geliştirici, operasyon ve QA ekipleri arasındaki engelleri kırmak için bir DevOps mühendisinin sorumluluğudur.

DevOps perspektifinden, **Geliştirme, Test ve Dağıtım** DevOps felsefesinin yetkinlikleri içerisindedir.

Bu süreci en etkili ve verimli hale getirmek için ise **Otomasyon** son derece önemlidir.


## Kaynaklar

Michael Cade her zaman öğrenme aracı olarak bu README dosyalarına ek kaynaklar eklemeye açıktır.

Bu belgeleri tam olarak takip etmek için bugün gördüğümüz konular hakkında genel bir fikir edinmek ve daha fazla derinlemesine araştırma yapabilmek için aşağıdaki videoları izlemeniz en iyisidir:

- [DevOps in 5 Minutes (ENG)](https://www.youtube.com/watch?v=Xrgk023l4lI)
- [What is DevOps? Easy Way (ENG)](https://www.youtube.com/watch?v=_Gpe1Zn-1fE&t=43s)
- [DevOps roadmap 2022 | Success Roadmap 2022 (ENG)](https://www.youtube.com/watch?v=7l_n97Mt0ko)

[Gun 2](day02.md)'de görüşmek üzere..
