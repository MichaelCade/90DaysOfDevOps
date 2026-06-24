## Genel Bakış: DevOps ve Bir Programlama Dilini Öğrenmek

Uzun vadeli başarı için, bir DevOps mühendisi olarak en azından temel düzeyde bir programlama dilini bilmek gerekmektedir. Yeni "Bir Programlama Dilini Öğrenmek" bölümünün bu ilk oturumunda, neden programlama dillerini bilmek bu kadar önemli bir beceri olduğunu keşfedeceğiz. Bu haftanın veya bölümün sonunda, bu öğrenme yolculuğunda ilerlemek için neden, nasıl ve ne yapmamız gerektiği konusunda daha iyi bir anlayışa sahip olacağız.

Sosyal medyada, DevOps ile ilgili pozisyonlarda programlama bilgisinin gerekip gerekmediğini sorsak, cevap kesin bir EVET olurdu. Ancak daha da önemli bir soru ortaya çıkıyor ve burada kesin bir cevap alamayız: **hangi programlama dili?** En yaygın cevap genellikle Python olur, çünkü öğrenme eğrisi daha düşük olabilir. Ayrıca, giderek daha fazla şekilde, başlangıç dilleri arasında Go veya Golang'in önem kazandığını görüyoruz.

DevOps'ta başarılı olmak için iyi bir programlama becerisine sahip olmak gerekmektedir. Şimdi, doğru yolunu seçebilmek için neden ihtiyacımız olduğunu anlamamız gerekiyor.

## Bir programlama dilini öğrenmeniz gerektiğini anlamak için nedenler şunlardır:

Python ve Go dillerinin DevOps mühendisleri için sık sık önerilmesinin nedeni, birçok DevOps aracının bu dillerle yazılmasıdır. Bu, DevOps araçları oluşturacaksanız veya mevcut araçları genişletecekseniz mantıklıdır. Bunun önemli olduğunu anlamak önemlidir, çünkü bu gerçekten neyi öğrenmeniz gerektiğini ve hangi dilin en faydalı olduğunu belirleyecektir.

- Eğer DevOps araçları oluşturacaksanız veya bu alanda çalışan bir ekibe katılıyorsanız, aynı dil öğrenmek mantıklı olacaktır.
Eğer Kubernetes veya Konteynerlere çok fazla dahil olacaksanız, Go'yu tercih etmek isteyebilirsiniz.
Yazar Michael Cade'in çalıştığı şirket (Kasten by Veeam), Kubernetes için veri yönetimine odaklanan Cloud-Native ekosistemde yer alıyor ve her şey Go dilinde yazılmış.

- Belki bu net bir seçim yapmak için yeterli sebepler değildir, belki hala eğitim görüyor veya bir kariyer geçişinde karar verme aşamasındasınızdır. Bu durumda, çalışmak istediğiniz uygulamalarla en iyi uyum sağlayan programlama dilini seçmek için emin olmanız gerekmektedir. Sonuçta, programlama dilinin temel kavramlarını iyi anlarsanız, bir dilin iyi öğrenildikten sonra diğerlerini daha hızlı öğrenebilirsiniz.

Unutmayın, DevOps felsefesinde bir yazılım geliştirici olmayı hedeflemiyorsunuz, sadece biraz daha fazla programlama dilini anlamak ve bu araçların ne yaptığını okuyup anlamak istiyorsunuz. Bu, özel ihtiyaçları olan projeler için işleri geliştirmemize yardımcı olacaktır.

Ayrıca,[Kasten K10](https://www.kasten.io/), [Terraform](https://www.terraform.io/) ve [HCL](https://github.com/hashicorp/hcl) gibi DevOps araçlarıyla nasıl etkileşimde bulunulacağını bilmek önemlidir. Bu araçlar, işleri gerçekleştirmek için bu araçlarla etkileşime geçmek için kullanılan yapılandırma dosyalarına ihtiyaç duyar. Bu yapılandırma dosyaları genellikle YAML formatında olacaktır. (Daha sonra YAML dosyalarının daha fazla detayını göreceğiz)

## Neden bir Programlama Dilini oğrenmem gerektiğine ikna oldum?

Genellikle veya rolünüze bağlı olarak, çoğu zaman mühendislik ekiplerine DevOps'u iş akışlarına entegre etmelerinde yardımcı olacak, uygulama etrafında birçok test gerçekleştirecek ve inşa edilen iş akışının önceden bahsettiğimiz DevOps prensipleriyle uyumlu olduğundan emin olacaksınız. Aslında, bu, uygulamanın performans sorunlarını çözmenin büyük bir kısmını veya benzer şeyleri yapmanın büyük bir kısmını oluşturacaktır. Bu, başlangıçtaki mantığa geri döner, ihtiyaç duyduğunuz programlama dili, üzerinde çalıştığınız veya ilgilendiğiniz projelerde veya işlerde yazılan kodlardır. Eğer uygulama NodeJS'te yazılmışsa, Go veya Python gibi yirmi tane sertifikanız olsa da çok yardımcı olmaz.

## Neden Go Dilini Tercih Etmeliyim?

Go (Golang), son yıllarda oldukça popüler hale gelmiş bir programlama dilidir. 2021 StackOverflow anketine göre, Go, en çok aranan programlama, betikleme ve işaretleme dilleri arasında dördüncü sırada yer aldı ve Python birinci sıradaydı. Go'nun bu popülerlik kazanmasının birkaç nedeni vardır:

- Basitlik ve Verimlilik: Go, basit bir sözdizimine ve az sayıda anahtar kelimeye sahip bir dil olarak tasarlanmıştır. Bu, kodun daha hızlı yazılmasını ve anlaşılmasını sağlar. Aynı zamanda derlenen bir dil olduğu için yüksek performanslı ve verimli çalışır.

- Paralellik ve Eşzamanlılık: Go, paralel ve eşzamanlı programlamaya doğal destek sağlar. "Go rutinleri" olarak adlandırılan hafif iş parçacıkları sayesinde, çoklu işlemleri kolayca yönetebilir ve performansı artırabilirsiniz.

- Hata Kontrolü: Go, hata kontrolü için dahili bir mekanizma olan hata döndürme (error handling) yöntemini kullanır. Bu, hataları daha iyi yönetmenize ve kodunuzu daha güvenilir hale getirmenize olanak tanır.

- Geniş Kütüphane Desteği: Go, zengin bir standart kütüphaneyle birlikte gelir ve birçok farklı alanda kullanılabilen çeşitli harici kütüphanelerle desteklenir. Bu, geliştirme sürecini hızlandırır ve daha karmaşık problemleri çözmek için kullanışlı araçlar sunar.

- Web Geliştirme: Go, web uygulamaları geliştirmek için kullanışlı bir dil olarak kabul edilir. Dahili olarak HTTP sunucu ve istemci kütüphaneleri sunar ve hızlı ve ölçeklenebilir web uygulamalarının geliştirilmesini kolaylaştırır.

Tabii ki, her proje veya senaryo için en iyi programlama dili farklı olabilir. Ancak, Go, DevOps için birçok avantaja sahip olan güçlü bir programlama dilidir ve bu nedenle birçok profesyonel tarafından tercih edilmektedir.

- [StackOverflow 2021 Developer Survey – Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)


Aynı şekilde, Kubernetes, Docker, Grafana ve Prometheus gibi en popüler DevOps araçları ve platformlarının birçoğu Go dilinde yazılmıştır.

DevOps için Go'yu tercih edilir kılan bazı özellikler nelerdir?

## Go dilinde programlar oluşturma ve dağıtma(deployment)

Python gibi yorumlanan bir dilin DevOps rolünde kullanılmasının bir avantajı, programı çalıştırmadan önce Python'u derlemeniz gerekmemesidir. Özellikle daha küçük otomasyon görevleri için, derleme gerektiren yapı sürecini yavaşlatmaz. Bununla birlikte, **Go bir derlenmiş bir dil olmasına rağmen**, Go doğrudan makine koduna derler. Go, hızlı derleme süreleriyle de tanınır.

## DevOps için Go vs Python

Go programları statik olarak bağlanır. Bu, bir Go programı derlendiğinde, her şeyin tek bir yürütülebilir ikili dosyada dahil edildiği anlamına gelir ve çalıştırmak istediğiniz uzak makinede kurulması gereken harici bağımlılıklara ihtiyaç duyulmaz. Bu, Go programlarının dağıtımını Python programına kıyasla daha kolay hale getirir, çünkü Python harici kütüphaneler kullanır ve bu kütüphanelerin tümünün çalıştırmak istediğiniz uzak makinede yüklü olduğundan emin olmanız gerekir.

Go, platformdan bağımsız bir dildir, yani Linux, Windows, macOS gibi tüm işletim sistemleri için yürütülebilir ikili dosyalar üretebilir ve bunu yapmak çok kolaydır. Python'da belirli işletim sistemleri için bu tür ikili dosyalar oluşturmak o kadar kolay değildir.

Go, çok verimli bir dil olup hızlı derleme süresi ve hızlı çalışma süresi ile CPU ve bellek gibi kaynakların daha az kullanımı sağlar, özellikle Python'a kıyasla. Go dili için birçok optimizasyon uygulanmıştır, bu da onu çok verimli hale getirir. (Aşağıda kaynaklar bölümünde daha fazla ayrıntı bulabilirsiniz)

Python'a karşılık olarak, genellikle belirli bir Python programını uygulamak için üçüncü taraf kütüphanelerinin kullanılmasını gerektiren durumlar olabilirken, Go dahili bir kütüphane içerir ve bu kütüphanenin çoğu DevOps için gereken işlevsellik doğrudan içerisine dahil edilmiştir. Bu, dosya işleme işlevselliği, HTTP web hizmetleri, JSON işleme, doğal destekli eşzamanlılık ve paralelizm, ayrıca dahili testler gibi işlevselliği içerir.

Bu, Python'ı hiçbir şekilde hafife almak anlamına gelmez, sadece yazarın tercih ettiği nedenleri belirtmektedir. Go'nun tercih edilmesinin sebebi, çalışılan şirkette ve Go ile yazılım geliştirdikleri projede işin mantığına uymasıdır.

Daha önce de belirtildiği gibi, ilk programlama dilini öğrendikten sonra diğer dilleri anlamak ve öğrenmek daha kolay olur. Muhtemelen hiçbir şirkette veya projede JavaScript ve Node JS uygulamalarının yönetimi, mimarisi, orkestrasyonu ve hata ayıklamasından kaçınmayacaksınız.

## Recursos

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Bu bölümde, önümüzdeki 6 gün boyunca yukarıda bahsedilen kaynaklar üzerinde çalışmak ve her gün notlar almak amaçlanmaktadır. Genellikle tam bir kurs için yaklaşık 3 saatlik bir süre ayrılmaktadır. Yazarın ilerlemek ve her bir kaynak üzerinde çalışmak için tam bir liste paylaşıldı, ancak zaman müsaade ederse.

Görüşmek üzere [Gun 8](day08.md).
