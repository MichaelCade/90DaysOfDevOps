## Gitting to know Git (Git'i Ogrenmek)

Özür dilerim, başlıktaki ve içerik boyunca korkunç kelime oyunları için. Git'i bir baba şakasına dönüştüren ilk kişi kesinlikle ben değilim!

Son iki yazıda, sürüm kontrol sistemleri hakkında ve git'in bir sürüm kontrol sistemi olarak temel iş akışlarından bazılarından bahsettik [Day 35](day35.md). Sonra sistemimize git'i yükledik, güncelledik ve yapılandırdık. Ayrıca, Client-Server versiyon kontrol sistemi ile Git arasındaki teoriye biraz daha derinlemesine girdik [Gun 36](day36.md).

Şimdi hepimizin git ile sıkça karşılaşacağı bazı komutları ve kullanım durumlarını öğreneceğiz.

### Git'te git yardımı almak için nereye gideriz?

Git ile yapılacak işleri unuttuğunuz veya bilmediğiniz zamanlar olacak. Yardıma ihtiyacınız olacak.

Yardım aramak için ilk başvuracağınız yer, Google veya başka bir arama motoru olacaktır.

İkinci olarak, bir sonraki yer resmi git sitesi ve belgelendirmesi olacaktır. [git-scm.com/docs](http://git-scm.com/docs) Burada yalnızca tüm mevcut komutlara sağlam bir referans bulmakla kalmayacak, aynı zamanda birçok farklı kaynağa da ulaşacaksınız.

![](Images/Day37_Git1.png)

Ayrıca bu belgelere terminalden bağlantınız yoksa oldukça yararlı olan aynı belgelere erişebiliriz. Örneğin, `git add` komutunu seçtiğimizde, `git add --help` komutunu çalıştırabiliriz ve aşağıda gördüğümüz gibi kılavuzu görürüz.

![](Images/Day37_Git2.png)

Ayrıca kabukta `git add -h` komutunu kullanarak da bize mevcut seçeneklerin özetini verebiliriz.

![](Images/Day37_Git3.png)

### Git ile ilgili yanlış inanışlar

"Git'in erişim kontrolü yok" - Bir lideri kaynak kodunu yönetmek için yetkilendirebilirsiniz.

"Git çok ağır" - Git, büyük projelere sahipseniz tarih miktarını azaltan sığ depolar sağlayabilir.

### Gerçek Eksiklikler

Binary dosyalar için ideal değil. Kaynak kodu için harika, ancak örneğin yürütülebilir dosyalar veya videolar için uygun değil.

Git kullanıcı dostu değil, aracın komutları ve işlevleri hakkında zaman harcamamız gereken gerçeği muhtemelen bunun önemli bir göstergesidir.

Ancak genel olarak, Git öğrenmesi zor ama kullanması kolaydır.

### Git Ekosistemi 

Git çevresindeki ekosistemi kısaca ele almak istiyorum, ancak bu alanların derinlemesine inmeyeceğiz, ancak yine de yüksek düzeyde burada belirtmek önemli olduğunu düşünüyorum.

Hemen hemen tüm modern geliştirme araçları Git'i destekler.

- Geliştirici araçları - Zaten Visual Studio Code'u bahsettik, ancak Sublime Text ve diğer metin düzenleyicileri ve IDE'ler için de git eklentileri ve entegrasyonlar bulacaksınız.
- Takım araçları - CI/CD açısından Jenkins gibi araçlar, mesajlaşma çerçevesi olarak Slack ve proje yönetimi ve sorun takibi için Jira gibi araçlar da bahsedildi.
- Bulut Sağlayıcıları - Tüm büyük bulut sağlayıcıları Git'i destekler, Microsoft Azure, Amazon AWS ve Google Cloud Platform.
- Git Tabanlı hizmetler - Ardından GitHub, GitLab ve BitBucket gibi konuları daha sonra daha detaylı ele alacağız. Bu hizmetleri kod için sosyal ağ olarak duymuştum!

### Git Cheatsheet(Kısa Notlar)

Çoğu bu komutları henüz ele almadık, ancak çevrimiçi olarak bulunan bazı notlara(cheatsheet) baktıktan sonra bazı git komutlarını ve amaçlarını belgelemek istedim. Hepsini hatırlamamıza gerek yok ve daha fazla pratik ve kullanım ile en azından git temellerini öğreneceksiniz.

Bunları [atlassian](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet) sitesinden aldım, ancak yazıya dökmek ve açıklamaları okumak, komutların ne olduğunu tanımak ve günlük görevlerde pratiğe geçmek için iyi bir yoldur.

### Git Temelleri

| Command       | Example                     | Description                                                                                                                 |
| ------------- | --------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| git init      | `git init <directory>`      | Belirtilen dizinde boş bir git deposu oluşturur.                                                                  |
| git clone     | `git clone <repo>`          | \<repo> konumundaki depoyu yerel makineye klonlar.                                                                      |
| git config    | `git config user.name`      | Geçerli depo için tüm taahhütlerde kullanılacak yazar adını tanımlar. `system`, `global`, `local` bayrağıyla yapılandırma seçenekleri ayarlanabilir. |
| git add       | `git add <directory>`       | \<directory> içindeki tüm değişiklikleri bir sonraki taahhüt için sahne alır. Aynı zamanda <dosyalar> ve <.> ile her şeyi ekleyebiliriz.                       |
| git commit -m | `git commit -m "<message>"` | Sahnelenen anlık görüntüyü taahhüt eder, \<message> ile neyin taahhüt edildiğini açıklar.                                                |
| git status    | `git status`                | Sahnelenen, sahnelenmeyen ve izlenmeyen dosyaları listeler.                                                                         |
| git log       | `git log`                   | 	İndeksiniz ile çalışma dizininiz arasındaki sahnelenmeyen değişiklikleri gösterir.                        |
| git diff      | `git diff`                  |  İndeksiniz ile çalışma dizininiz arasındaki sahnelenmemiş değişiklikleri gösterir.                                                             |

### Git Degisiklikleri Geri Alma

| Command    | Example               | Description                                                                                                                           |
| ---------- | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| git revert | `git revert <commit>` | \<commit> içinde yapılan tüm değişiklikleri geri alan yeni bir taahhüt oluşturur ve bunu mevcut dal üzerine uygular.                              |
| git reset  | `git reset <file>`    | \<file> dosyasını hazırlama alanından kaldırır, ancak çalışma dizinini değiştirmez. Bu, değişiklikleri üzerine yazmadan bir dosyayı hazırlamadan kaldırır. |
| git clean  | `git clean -n`        | Çalışma dizininden kaldırılacak dosyaları gösterir. Temizlemeyi gerçekleştirmek için `-n` yerine `-f` kullanın.                        |

### Git History Yeniden Yazma

| Command    | Example              | Description                                                                                                                              |
| ---------- | -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| git commit | `git commit --amend` | Son taahhüdü sahnelenen değişiklikler ve son taahhüdün birleştirilmiş haliyle değiştirir. Son taahhüt mesajını düzenlemek için sahnelenmemiş olarak kullanın. |
| git rebase | `git rebase <base>`  | Mevcut dalı \<base> üzerine yeniden tabanlar. \<base>, bir taahhüt kimliği, dal adı, bir etiket veya HEAD'e göre göreceli bir referans olabilir.                   |
| git reflog | `git reflog`         | Yerel depo başına yapılan değişikliklerin bir günlüğünü gösterir. Tarih bilgisi görmek için --relative-date bayrağını, tüm başvuruları görmek için --all bayrağını ekleyin.              |

### Git Branches

| Command      | Example                    | Description                                                                                                   |
| ------------ | -------------------------- | ------------------------------------------------------------------------------------------------------------- |
| git branch   | `git branch`               | Depodaki tüm dalları listeler. Bir \<branch> argümanı ekleyerek <branch> adında yeni bir dal oluşturabilirsiniz. |
| git checkout | `git checkout -b <branch>` | \<branch> adında yeni bir dal oluşturup ona geçiş yapar. -b bayrağını bırakarak varolan bir dala geçebilirsiniz.            |
| git merge    | `git merge <branch>`       | \<branch>'ı mevcut dala birleştirir.                                                                       |

### Git Uzak Repoları

| Command        | Example                       | Description                                                                                                                         |
| -------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| git remote add | `git remote add <name> <url>` | Uzak depo için yeni bir bağlantı oluşturur. Bir uzak depo ekledikten sonra, diğer komutlarda \<name> adını \<url> için kısayol olarak kullanabilirsiniz.      |
| git fetch      | `git fetch <remote> <branch>` | Belirtilen depodan belirtilen \<branch>'ı alır. Tüm uzak başvuruları almak için \<branch> kısmını bırakın.                                 |
| git pull       | `git pull <remote>`           | Belirtilen uzaktan mevcut daldaki kopyayı alır ve hemen yerel kopyayla birleştirir.                                     |
| git push       | `git push <remote> <branch>`  | Dalları \<remote> depoya gönderir ve gerekli taahhütleri ve nesneleri oluşturur. Eğer oluşturulmamışsa uzak depoda adlandırılmış bir dal oluşturur. |

### Git Diff

| Command           | Example             | Description                                                            |
| ----------------- | ------------------- | ---------------------------------------------------------------------- |
| git diff HEAD     | `git diff HEAD`     | Çalışma dizini ile son taahhüt arasındaki farkı gösterir.              |
| git diff --cached | `git diff --cached` | Sahne alanındaki değişiklikler ile son taahhüt arasındaki farkı gösterir. |

### Git Config(Yapılandırma)

| Command                                              | Example                                                | Description                                                                                                                                   |
| ---------------------------------------------------- | ------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------- |
| git config --global user.name \<name>                 | `git config --global user.name <name>`                 | Geçerli kullanıcı tarafından yapılan tüm taahhütler için yazar adını tanımlar.                                                        |
| git config --global user.email \<email>               | `git config --global user.email <email>`               | Geçerli kullanıcı tarafından yapılan tüm taahhütler için yazar e-postasını tanımlar.                                                  |
| git config --global alias \<alias-name> \<git-command> | `git config --global alias <alias-name> <git-command>` | Bir git komutu için kısayol oluşturur.                                                                                                  |
| git config --system core.editor \<editor>             | `git config --system core.editor <editor>`             | Makinedeki tüm kullanıcılar için komutlar tarafından kullanılacak metin düzenleyiciyi tanımlar. \<editor>, istenilen düzenleyici komutu olmalıdır. |
| git config --global --edit                           | `git config --global --edit `                          | Tüm yapılandırmaları manuel olarak düzenlemek için genel yapılandırma dosyasını metin düzenleyicide açar.                                                                       |

### Git Rebase

| Command              | Example                | Description                                                                                                                                 |
| -------------------- | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| git rebase -i \<base> | `git rebase -i <base>` | 	Mevcut dalı interaktif olarak \<base> üzerine yeniden tabanlar. Her taahhütün yeni tabana nasıl aktarılacağı için komutları girmek için düzenleyiciyi açar. |

### Git Pull

| Command                    | Example                      | Description                                                                                                                                   |
| -------------------------- | ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| git pull --rebase \<remote> | `git pull --rebase <remote>` | Mevcut dalın uzak kopyasını alır ve yerel kopyasına yeniden tabanlar. Birleştirmek yerine dalları bütünleştirmek için git rebase kullanır. |

### Git Reset

| Command                   | Example                     | Description                                                                                                                                   |
| ------------------------- | --------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| git reset                 | `git reset `                | stage alanını en son sekliyle eşleştirir, ancak çalışma dizinini değiştirmez.                                              |
| git reset --hard          | `git reset --hard`          | stage alanını ve çalışma dizinini en son sekliyle eşleştirir, çalışma dizinindeki tüm değişiklikleri geçersiz kılar.             |
| git reset \<commit>        | `git reset <commit>`        | Mevcut branchi ucunu geriye taşır, stage alanını eşleştirir, ancak çalışma dizinini değiştirmez.                     |
| git reset --hard \<commit> | `git reset --hard <commit>` | Öncekine benzer, ancak hem sahne alanını hem de çalışma dizinini eşleştirir \<commit> sonrası yapılan tüm taahhütleri ve değişiklikleri siler. |

### Git Push

| Command                   | Example                     | Description                                                                                                                                     |
| ------------------------- | --------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| git push \<remote> --force | `git push <remote> --force` | Eğer sonuçta hızlı ileri birleştirme yapılmayan bir commit oluşursa, git push işlemini zorlar. --force bayrağını kullanmadan önce ne yaptığınızdan emin olun.  |
| git push \<remote> --all   | `git push <remote> --all`   | Tüm local branchleri belirtilen uzak depoya gönderir.                                                                        |
| git push \<remote> --tags  | `git push <remote> --tags`  | commitler otomatik olarak birleştirme yapılması veya --all flagi kullanıldığında gönderilmez. --tags flagi, tüm local tagleri uzak repoya gönderir. |

## Kaynaklar

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)
- [Git cheatsheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)

Gorusmek Uzere [Gun 38](day38.md)
