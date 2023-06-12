Bugün, Değişkenler, Sabitler ve Veri Türleri konularına bir göz atacağız ve yeni bir program yazacağız. [Techworld con Nana](https://www.youtube.com/watch?v=yyUHQIec83I) 


## Go'da Değişkenler ve Sabitler

Uygulamamızı planlamaya başlayalım. #90DaysOfDevOps meydan okumasında kaç gün kaldığımızı bize söyleyen bir program üzerinde çalışmak iyi bir fikir gibi görünüyor.

Burada dikkate almanız gereken ilk şey, uygulamayı inşa ettikçe, katılımcılarımızı karşılayarak ve kullanıcıya tamamladığı gün sayısı hakkında geri bildirim vererek, program boyunca #90DaysOfDevOps terimini birçok kez kullanabileceğiz. Bu nedenle, programımız için #90DaysOfDevOps isimli bir değişken yapma fırsatımız var.

- Değişkenler değerleri depolamak için kullanılır.
- Bilgilerimizi veya saklanan değerlerimizi içeren küçük kutular gibidirler (Bu, başlayanlara anlatılan şekildedir 😊).
- Bu değişkeni program boyunca kullanabiliriz, bu da meydan okuma veya değişken değişirse, sadece bir yerde değeri değiştirmemiz gerektiği anlamına gelir. Bu, bu değeri değiştirerek topluluktaki diğer meydan okumalara taşıyabileceğimiz anlamına gelir.

Bu Go programımızda bunu deklare etmek için değişkenler için bir **anahtar kelime** kullanarak bir değer tanımlıyoruz. Bu, daha sonra göreceğiniz `func main` kod bloğu içinde kalacaktır. Anahtar kelimeler hakkında daha fazla bilgiyi [burada](https://go.dev/ref/spec#Keywords)bulabilirsiniz.

Değişken adlarının açıklayıcı olduğundan emin olun. Bir değişkeni tanımlarsanız, onu kullanmalısınız; aksi takdirde bir hata alırsınız. Bu, kullanılmayan ölü kodu önlemek içindir. Aynı durum kullanılmayan paketler için de geçerlidir.

```go
var challenge = "#90DaysOfDevOps"
```

Yukarıdakini belirledikten sonra, aşağıdaki kod parçasında bunu kullanacağız. Aşağıdaki çıktıda, değişkeni kullandığımızı görebilirsiniz.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    fmt.Println("Welcome to", challenge, "")
}
```

Yukarıdaki kodu [day11_example1.go](Go/day11_example1.go) dosyasında bulabilirsiniz.

Aşağıdaki çıktıyı oluşturmak için yukarıdaki örneği derleyip çalıştırırsanız:

![](Images/Day11_Go1.png)

Bu özel meydan okumanın 90 gün olduğunu biliyoruz, ancak bir sonraki meydan okuma belki 100 gün olacak, bu yüzden burada bize yardımcı olacak bir değişken tanımlayacağız. Bununla birlikte, programımız için bunu bir sabit olarak tanımlamak istiyoruz. Sabitler, değişkenler gibi, ancak değerleri kod içinde değiştirilemez (ancak daha sonra aynı kodla yeni bir uygulama oluşturabilir ve sabiti değiştirebiliriz. 90, bu uygulamayı çalıştırırken değişmeyecektir).


Kodumuza `const` ekleyerek bunu yapalım ve bunu yazdırmak için başka bir satır kod ekleyelim. 

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge")
}
```


Yukarıdaki kodu [day11_example2.go](Go/day11_example2.go) dosyasında bulabilirsiniz.

Ardından `go build` işlemini tekrar gerçekleştirip çalıştırdığımızda aşağıdaki sonucu göreceğiz.


![](Images/Day11_Go2.png)

Son olarak, tamamladığımız gün sayısı için başka bir değişken ekleyeceğiz. Ancak bu programın sonu olmayacak, daha fazla işlevsellik eklemek için Día 12 programımıza devam edeceğiz.

Aşağıda, tamamlanan gün sayısıyla birlikte `dayscomplete` değişkeninin eklenmiş halini göreceksiniz.

```go
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90
    var dayscomplete = 11

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge and you have completed", dayscomplete, "days")
    fmt.Println("Great work")
}
```

Önceki kodu [day11_example3.go](Go/day11_example3.go) dosyasında bulabilirsiniz.

Tekrar `go build` komutuyla derleyip oluşturulan dosyayı çalıştırarak veya doğrudan `go run` komutunu kullanarak çalıştıralım.


![](Images/Day11_Go3.png)

Kodu daha okunabilir ve düzenlenebilir hale getirmek için bazı örnekler bulunmaktadır. Şimdiye kadar `Println` kullanıyorduk, ancak `%v` ile değiştirerek kodumuzu basitleştirebiliriz, bu şekilde değişkenlerimizi kod satırının sonunda sırayla tanımlamış oluruz. Ayrıca, bir satır atlamak için `\n` kullanacağız.

`%v` kullanarak varsayılan bir değere sahip olacağız, ancak diğer seçenekleri fmt paketi belgelerinde bulabilirsiniz.

Daha da ilginç hale geliyor, örnek kodu [day11_example4.go](Go/day11_example4.go) dosyasında görebilirsiniz.

Değişkenler ayrıca daha basit bir formatta da tanımlanabilir. var ve "tip" yerine, aynı işlevselliği elde etmek için aşağıdaki şekilde kodlayabilirsiniz, bu da daha temiz ve basit bir görünüm sağlar. Bu sadece değişkenler için çalışır, sabitler için geçerli değildir.

```go
func main() {
    challenge := "#90DaysOfDevOps"
    const daystotal = 90
```

## Veri Tipleri

Önceki örneklerde değişkenlerin tipini belirtmedik. Bunun sebebi, bir değer atadığımızda Go'nun hangi tür olduğunu bilmek veya en azından depolanan değere dayanarak tahmin edebilmesidir. Ancak, bir kullanıcının bunu girmesini istiyorsak belirli bir tipe ihtiyacımız olacaktır.

Şu ana kadar kodumuzda String'ler ve Integer'lar kullandık. Integer'lar gün sayısı için kullanılırken, String'ler meydan okuma adını içeriyor.

Her veri tipinin farklı işlevleri ve davranışları olduğunu unutmak önemlidir. Örneğin, Integer'lar çarpılabilirken, String'ler çarpılamaz.

Dört kategori vardır:

- **Temel Veri Tipleri**: Sayılar, String'ler ve Boolean'lar bu kategoriye girer.
- **Toplu Veri Tipleri**: Diziler ve struct'lar bu kategoriye girer.
- **R**eferans Veri Tipleri**: İşaretçiler (pointers), dilimler (slices), haritalar (maps), fonksiyonlar (functions) ve kanallar (channels) bu kategoriye dahildir.
- **Arayüz Veri Tipi**

Veri tipi, programlamada önemli bir kavramdır çünkü değişkenlerin değerlerinin boyutunu ve türünü belirtir.

Go, statik olarak yazılmış bir dildir, yani bir değişkenin tipi bir kez tanımlandığında, yalnızca o türde veri depolayabilir.

Go'nun üç temel veri tipi vardır:


- **bool**:  Boolean değerini temsil eder, true veya false olabilir.
- **Numeric**:  Integer türleri, ondalık noktalı sayılar ve karmaşık türleri temsil eder.
- **string**: Bir string değerini temsil eder.

Bu kaynak, veri tipleri hakkında detaylı bilgi içerir: [Golang by example](https://golangbyexample.com/all-data-types-in-golang-with-examples/).

Bir değişkene tür atamamız gerekiyorsa aşağıdaki gibi yapabiliriz:

```go
var TwitterHandle string
var DaysCompleted uint
```

Go, değer atanan değişkenleri ima ettiği için bu değerleri şu şekilde yazdırabiliriz:

```go
fmt.Printf("challenge is %T, daystotal is %T, dayscomplete is %T\n", conference, daystotal, dayscomplete)
```

Farklı türlerde çok sayıda integer ve float bulunmaktadır.


- **int** =  Tam sayıları temsil eder.
- **unint** = Pozitif tam sayıları temsil eder.
- **floating point types** = Ondalık noktalı sayıları içeren türlerdir.

## Recursos

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Bir sonraki günde, kullanıcının programımıza veri girebilmesi için giriş işlevselliği eklemeye başlayacağız, örneğin tamamlanan günleri girebilmesi gibi.

Gorusmek Uzere [Gun 12](day12.md).
