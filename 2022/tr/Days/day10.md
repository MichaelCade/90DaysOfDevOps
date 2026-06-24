## Go Workspace(Calısma Alanı)

Go Çalışma Alanı hakkında daha fazla bilgi vermemiz gerekiyor. [Gun 8](day08.md) 'de, Go'yu başlatmak ve `Hello #90DaysOfDevOps` demosuna ulaşmak için Go çalışma alanı hakkında kısaca konuştuk.

Hatırlıyor musunuz, varsayılan değerleri seçtik ve daha sonra zaten tanımlanmış GOPATH'deki Go klasörümüzü oluşturduk, değil mi? Ancak aslında GOPATH'i istediğiniz yere değiştirebilirsiniz.

Eğer

```shell
echo $GOPATH
```

Çıktı benimkine benzer olmalı (kullanıcı adınızla birlikte):

```shell
/home/michael/projects/go
```
> **Linux'ta config yapmanız gerekiyor** 
> 
> Benim durumumda, zsh shell kullanıyorum ve `$HOME/.zshrc` dosyasında yapılır. Eğer bash kullanıyorsanız, `$HOME/.bashrc` dosyasında yapılır. 
> 
> Sadece şu satırı eklememiz gerekiyor:
> ```shell
> export GOPATH=$HOME/work
> export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
> ```
>
> ![](Images/Day10_Go0.png)
>
> Kabuğu yeniden başlatın ve önce bahsedilen `echo` komutunu çalıştırın.
> 
> ```shell
> exec zsh
> echo $GOPATH
> ```
> 
> ![](Images/Day10_Go0b.png)
> 
> Daha fazla bilgi icin [resmi dokuman](https://go.dev/doc/install)

İşte burada önceki günlerde oluşturduğumuz 3 directory'yi oluşturduğumuz yer **src**, **pkg** ve **bin** .

![](Images/Day10_Go1.png) 

- **src** Bu, tüm Go programlarınızın ve projelerinizin depolandığı yerdir. Bu, Go depolarınızın isim alanı paket yönetimini yönetir. İşte bizim Hello projesi için Hello klasörünü gördüğümüz yer." turkçe çevirisidir.

![](Images/Day10_Go2.png)

- **pkg** Bu, programlara yüklenmiş veya yüklenmiş olan paketlerin arşivlenmiş dosyalarının bulunduğu yerdir. Bu, kullanılan paketlerin değiştirilip değiştirilmediğine bağlı olarak derleme sürecini hızlandırmaya yardımcı olur.

![](Images/Day10_Go3.png)

- **bin** Bu, derlenmiş tüm ikili dosyaların depolandığı yerdir.

![](Images/Day10_Go4.png)

Hello #90DaysOfDevOps programımız karmaşık bir program olmadığından, burada başka bir harika kaynaktan alınmış daha karmaşık bir Go Programı örneği bulunmaktadır. Bu kaynağa göz atmanızı öneririm. [GoChronicles](https://gochronicles.com/)

![](Images/Day10_Go5.png)

Hello #90DaysOfDevOps" programımız karmaşık bir program değil. Biraz fikir sahibi olmanız için, başka bir önemli kaynaktan alınmış ve değerli bir incelemeyi hak eden daha karmaşık bir Go programı örneği şu şekildedir. [GoChronicles](https://gochronicles.com/)

## Derleme(Compiling) & Running Code

[Gun 9](day09.md)'da kod derlemesine kısa bir giriş yaptık, ancak biraz daha derinlemesine inebiliriz. Devam edelim.

Kodumuzu çalıştırmak için öncelikle **derleme**miz gerekiyor. Go içinde bunu yapmanın üç farklı yolu vardır.

- `go build`
- `go install`
- `go run`

Go kurulumuyla elde ettiğimiz şeylere, bahsettiğimiz derleme aşamasına geçmeden önce bir göz atalım.

Go'yu [Gun 8](day08.md) 'de yüklediğimizde, Go araçları olarak adlandırılan bir şeyi yüklemiş olduk. Bu araçlar, Go kaynak dosyalarımızı oluşturmamıza ve işlememize olanak sağlayan çeşitli programlardan oluşur. Bu araçlar arasında `Go` komutu da bulunur.

Ayrıca, Go'nun standart kurulumunda bulunmayan, projeniz için oldukça faydalı olabilecek ek araçlar da yüklenebilir.

Terminalinizi açıp `go` komutunu yazarsanız aşağıdaki resimdeki gibi bir çıktı almanız gerekiyor. Daha sonra "Additional Help Topics" (Ek Yardım Konuları) gibi ek yardım konularını göreceksiniz, ancak şu anda bunlara odaklanmamıza gerek yok.

![](Images/Day10_Go6.png)

> [Go Yardım Dokumanı](Go/Ayuda_go_traducida.md).

Eğer önceki günlerden buradaysan, en azından ikisini zaten kullandığımızı hatırlayacaksınız. [Gun 8](day08.md).

![](Images/Day10_Go7.png)

Öğreneceğimiz komutlar `build`, `install` ve `run`.

![](Images/Day10_Go8.png)

- `go run` - komut satırında belirtilen .go dosyalarından oluşan ana paketi derler ve çalıştırır. Komut geçici bir klasörde derlenir.
- `go build` - Go build komutu, paketleri ve bağımlılıkları derlerken, paketi geçerli dizinde derler. Bu şekilde `main` paketi, yürütülebilir dosyayı geçerli dizine yerleştirir. Aksi takdirde, yürütülebilir dosyayı `pkg` klasörüne yerleştirir. `go build` ayrıca, Go ile uyumlu herhangi bir işletim sistemi için yürütülebilir bir dosya oluşturmanıza olanak tanır.
- `go install` - `go install` komutu, `go build` ile aynı işlemi yapar, ancak yürütülebilir dosyayı `bin` klasörüne yerleştirir.

`go install` komutunu zaten gördük, ancak isterseniz tekrar yapabilirsiniz. Hatırlayacağınız gibi, bu komut yürütülebilir dosyayı `bin` klasörüne yerleştirir.

![](Images/Day10_Go9.png)

Eğer aşağıdaki çalma listeleri veya videoları takip ediyorsanız, bu notlarda çevirilen bölümlerin bir kısmını göreceksiniz. Bu, genel olarak ihtiyaç duyduğunuz birçok alanın çok daha iyi bir anlayışını sağlayacak kaynaklardır. Ancak, 7 günlük (veya 7 saatlik) süreyi en ilgili konular üzerinde belgelemeye çalışıyoruz.

## Kaynaklar

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Gorusmek Uzere [Gun 11](day11.md).
