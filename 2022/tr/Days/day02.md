## Bir DevOps Mühendisinin Sorumlulukları

marım buraya 90DaysOfDevOps kaynakları ve yayınları aracılığıyla ulaşmışsınızdır. [Gun 1 #90GunDevOps](day01.md).

İlk gün DevOps kavramına kısaca değinildi, ancak şimdi daha derinlemesine bir şekilde incelemeli ve bir uygulama oluşturulurken iki ana bölüm olduğunu anlamalıyız:
- **Geliştirme** kısmı, yazılım geliştiricilerin uygulamayı programladığı ve test ettiği kısımdır.
- **Operasyonlar** kısmı, uygulamanın bir sunucuya dağıtıldığı ve bakımının yapıldığı kısımdır.

## DevOps, Geliştirme ve Operasyonlar Arasındaki Bağlantıdır.

DevOps'un işleyişini veya DevOps odaklı bir mühendisin gerçekleştirdiği görevleri anlamak için, geliştirme ve operasyonların araçlarını, sürecini ve genel bakışını anlamamız gerekiyor.

Her şey uygulama ile başlar, DevOps'ta her şey uygulama etrafında döner.

Geliştiriciler bir uygulama oluşturacak, bunu birçok farklı teknoloji yığınıyla yapabilirler ve şu anda bunu hayal gücümüze bırakalım, çünkü daha sonra bu konuya daha fazla değineceğiz. Bu aynı zamanda farklı programlama dilleri, yapı araçları, kod depoları vb. içerebilir.

DevOps mühendisi olarak uygulamayı kodlama göreviniz olmasa da, bir geliştiricinin nasıl çalıştığını ve kullandığı sistemleri, araçları ve süreçleri anlamanız başarı için önemlidir.

Çok genel bir düzeyde, uygulamanın tüm gereksinim duyduğu hizmetlerle, veri hizmetleriyle nasıl iletişim kurduğunu ve bunun nasıl test edilmesi gerektiğini bilmelisiniz.

Uygulamanın bir sunucuda (nerede olduğu önemli değil, ancak bir sunucuda) dağıtılması gerekecektir. Uygulamanın oluşturulduğu uygulamaya bağlı olarak, müşteri veya son kullanıcının buna erişebilmesi beklenir.

Bu sunucunun bir yerde çalışması gerekmektedir, yerel sunucularda, genel bulutlarda, serverless ortamlarda (Tamam, bu biraz uzak bir konu, bunu ele almayacağız, ancak bu bir seçenek ve giderek daha fazla şirket bu yöne yönelmektedir)... Bu sunucuların oluşturulması, yapılandırılması ve uygulamanın çalışması için hazırlanması gerekmektedir. Bu görevler, bir DevOps mühendisi olarak sizin sorumluluğunuza düşebilir.

Bu sunucular bir işletim sistemi üzerinde çalışır ve genellikle bir Linux dağıtımı olacaktır. Linux hakkında temel bilgilere değineceğimiz bir bölümümüz daha olacak (bir hafta boyunca).

Ayrıca, ağ ve ağ yapılandırması hakkında biraz bilgiye ihtiyaç duyabiliriz, bu da bir ölçüde DevOps mühendisinin sorumluluğuna girebilir. DNS, DHCP, yük dengeleme vb. ile ilgili ayrıntıları daha sonra ele alacağız.

## Herseyi bilen hiçbir şeyi bilemez.

Bu noktada, ağ veya altyapı konusunda uzman olmanız gerekmediğini söylemeliyim, ancak işlerin nasıl çalıştığını ve birbiriyle nasıl iletişim kurduğunu anlama ve temel bir programlama diline sahip olma bilgisine ihtiyaç duyulur. Geliştirici olmanız gerekmez, bunu bir alanda uzman olarak da benimseyebilirsiniz ve bu diğer alanlara uyum sağlamak için büyük bir temel olabilir.

Ayrıca, muhtemelen bu sunucuların veya uygulamanın günlük yönetiminden sorumlu olmayacaksınız.

Sunuculardan bahsettik, ancak uygulamanızın çoğunlukla çalıştığı bir sunucuda çalıştırılması için konteynerlere yönelik bir geliştirme yapılması muhtemeldir. Sadece sanallaştırma konusunda değil, hizmet olarak bulut altyapısı (IaaS), konteynerleştirme vb. konularında da bir anlayışa ihtiyaç duyulur. Bu 90 günlük süreçte odak noktası daha çok konteynerlere yönelik olacaktır.

## Genel Bakıs Acısıyla

Bir yandan geliştiricilerimiz uygulama için yeni özellikler ve işlevsellikler (hata düzeltmeleri de dahil) oluşturuyorlar.

Diğer yandan, bu uygulamayı çalıştırmak ve gerekli tüm hizmetlerle iletişim kurmak için yapılandırılmış ve yönetilen bir ortam, altyapı veya sunucularımız bulunuyor.

Büyük soru şudur: Bu özellikleri ve hata düzeltmelerini nasıl ürünlerimize entegre eder ve son kullanıcılara sunarız?

Uygulamanın yeni bir sürümünü nasıl piyasaya süreriz? Bu, bir DevOps mühendisinin başlıca görevlerinden biridir ve burada önemli olan yalnızca bunu bir kez yapmanın yollarını bulmak değil, bunu sürekli, otomatik ve verimli bir şekilde yapmaktır ve ayrıca uygun testleri içermelidir.

Burada öğrenme günümüzü tamamlayacağız. Devam etmek için güç toplamamız gerekiyor. Önümüzdeki günlerde, DevOps'un bazı alanlarını daha derinlemesine ele alacak ve araçların, süreçlerin ve bunların faydalarının özel bölümlerine giriş yapmaya hazırlanacağız.

## Kaynaklar

Her zaman readme dosyalarında ek kaynaklar bırakılacaktır, çünkü burada öğrenmek için bulunuyoruz.

Bugünkü konulara daha genel bir bakış elde etmek ve daha fazla derinlemesine bilgi edinmek için aşağıdaki videoları izlemenizi ve makaleleri okumanızı öneririm:

- [What is DevOps? - TechWorld with Nana (ENG)](https://www.youtube.com/watch?v=0yWAtQ6wYNM)
- [What is DevOps? - GitHub YouTube (ENG)](https://www.youtube.com/watch?v=kBV8gPVZNEE)
- [What is DevOps? - IBM YouTube (ENG)](https://www.youtube.com/watch?v=UbtB4sMaaNM)
- [What is DevOps? - AWS (ENG)](https://aws.amazon.com/devops/what-is-devops/)
- [What is DevOps? - Microsoft (ENG)](https://docs.microsoft.com/en-us/devops/what-is-devops)

Görüşmek üzere [Gun 3](./day03.md).'de buluşalım.
