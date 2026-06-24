## Thực hành: Xây dựng một ứng dụng yếu

Không ai thực sự muốn xây dựng một ứng dụng yếu hoặc dễ bị tấn công... đúng không?

Câu trả lời đúng là không, không ai nên hoặc muốn xây dựng một ứng dụng yếu, và không ai có ý định sử dụng các gói hoặc phần mềm mã nguồn mở khác mang theo các lỗ hổng của riêng nó.

Trong phần giới thiệu cuối cùng này về DevSecOps, tôi muốn cố gắng xây dựng và nâng cao nhận thức về một số cấu hình sai và điểm yếu có thể bị bỏ qua. Sau đó, trong 84 ngày tới hoặc thậm chí sớm hơn, chúng ta sẽ nghe từ một số chuyên gia trong lĩnh vực bảo mật về cách ngăn chặn những điều xấu và ứng dụng yếu được tạo ra.

### Xây dựng ứng dụng yếu đầu tiên của chúng ta

<span style="color:red">**Thông điệp quan trọng: Bài tập này nhằm làm nổi bật các điểm yếu và sai sót trong một ứng dụng, Vui lòng thử điều này tại nhà nhưng hãy cẩn thận vì đây là thực hành không tốt**</span>

Ở giai đoạn này, tôi sẽ không đi qua môi trường phát triển phần mềm của mình một cách chi tiết. Tôi thường sử dụng VScode trên Windows với WSL2 được kích hoạt. Chúng ta có thể sử dụng Vagrant để cung cấp các phiên bản tính toán chuyên dụng cho VirtualBox, tất cả những điều này tôi đã đề cập trong suốt phần #90DaysOfDevOps phiên bản trước, chủ yếu trong phần Linux.

### Các thực hành code xấu

Rất dễ dàng để sao copy paste trong GitHub!

Có bao nhiêu người kiểm tra từ đầu đến cuối package mà họ bao include trong mã của bạn?

Chúng ta cũng phải xem xét:

- Chúng ta có tin tưởng người dùng/người duy trì không
- Không xác thực đầu vào trong mã của chúng ta
- Hardcode cứng các secrets thay vì sử dụng env hoặc quản lý secrets
- Tin tưởng mã mà không xác thực
- Thêm secret của bạn vào các kho lưu trữ công khai (Có bao nhiêu người đã làm điều này?)

Quay lại chủ đề tổng thể, DevSecOps, tất cả những gì chúng ta đang làm hoặc hướng tới là các vòng lặp nhanh hơn của ứng dụng hoặc phần mềm của chúng ta, nhưng điều này có nghĩa là chúng ta có thể có các lỗi và rủi ro sớm hơn.

Chúng ta cũng có thể triển khai cơ sở hạ tầng của mình bằng mã, một rủi ro khác là bao gồm mã xấu ở đây cho phép các tác nhân xấu xâm nhập qua các lỗi.

Các triển khai cũng sẽ bao gồm quản lý cấu hình ứng dụng, một cấp độ khác của các lỗi có thể xảy ra.

Tuy nhiên! Các vòng lặp nhanh hơn có thể và thực sự có nghĩa là các bản sửa lỗi sớm hơn.

### OWASP - Dự án Bảo mật Ứng dụng Web Mở

*"[OWASP](https://owasp.org/) là một tổ chức phi lợi nhuận hoạt động để cải thiện bảo mật phần mềm. Thông qua các dự án phần mềm mã nguồn mở do cộng đồng dẫn dắt, hàng trăm chương trình trên toàn thế giới, hàng chục nghìn thành viên và các hội nghị giáo dục và đào tạo hàng đầu, Quỹ OWASP là nguồn lực cho các nhà phát triển và công nghệ để bảo mật web."*

Nếu chúng ta nhìn vào tập dữ liệu gần đây nhất của họ và [top 10](https://owasp.org/www-project-top-ten/) của họ, chúng ta có thể thấy các mục lớn sau đây là lý do tại sao mọi thứ trở nên tồi tệ và sai lầm.

1. Broken Access Control
2. Cryptographic Failures
3. Injection (2020 #1)
4. Insecure Design (New for 2021)
5. Security Misconfiguration
6. Vulnerable and Outdated Components (2020 #9)
7. Identification and authentication failures (2020 #2)
8. Software and Data integrity failures (New for 2021)
9. Security logging and monitoring failures (2020 #10)
10. Server-side request forgery (SSRF)

### Quay lại ứng dụng

<span style="color:red">**Cảnh báo trên vẫn còn hiệu lực, tôi sẽ triển khai điều này lên một máy ảo VirtualBox cục bộ. NẾU bạn quyết định triển khai điều này lên một phiên bản đám mây thì trước tiên hãy cẩn thận và thứ hai là biết cách whitelist CSP của bạn chỉ cho phép truy cập bằng IP của riêng bạn!**</span>

Được rồi, tôi nghĩ rằng cảnh báo như vậy là đủ, tôi chắc chắn rằng chúng ta có thể thấy các cảnh báo màu đỏ trong vài tuần tới khi chúng ta đi sâu hơn vào thảo luận về chủ đề này.

Ứng dụng mà tôi sẽ sử dụng sẽ từ [DevSecOps.org](https://github.com/devsecops/bootcamp/blob/master/Week-2/README.md). Đây là một trong những bootcamp của họ từ nhiều năm trước nhưng vẫn cho phép chúng ta thấy một ứng dụng xấu trông như thế nào.

Có khả năng thấy một ứng dụng xấu hoặc yếu có nghĩa là chúng ta có thể bắt đầu hiểu cách bảo mật nó.

Một lần nữa, tôi sẽ sử dụng VirtualBox trên máy cục bộ của mình và tôi sẽ sử dụng vagrantfile sau (liên kết ở đây để giới thiệu về vagrant)

Cảnh báo đầu tiên là vagrant box này đã được tạo ra hơn 2 năm trước!
```
Vagrant.configure("2") do |config|
  config.vm.box = "centos/7"
  config.vm.provider :virtualbox do |v|
   v.memory  = 8096
   v.cpus    = 4
end
end
```
Nếu điều hướng đến thư mục này, bạn có thể sử dụng `vagrant up` để khởi động máy centos7 trong môi trường của bạn.

![](../../images/day06-1.png)

Sau đó, chúng ta sẽ cần truy cập vào máy của mình, bạn có thể làm điều này với `vagrant ssh`.

Chúng ta sẽ cài đặt MariaDB làm cơ sở dữ liệu cục bộ để sử dụng trong ứng dụng của mình.

`sudo yum -y install mariadb mariadb-server mariadb-devel`

khởi động dịch vụ với

`sudo systemctl start mariadb.service`

Chúng ta phải cài đặt một số dependencies, đây cũng là nơi tôi phải thay đổi những gì Bootcamp đề xuất vì NodeJS không có sẵn trong các kho lưu trữ hiện tại.

`sudo yum -y install links`
`sudo yum install --assumeyes epel-release`
`sudo yum install --assumeyes nodejs`

Bạn có thể xác nhận bạn đã cài đặt node với `node -v` và `npm -v` (npm sẽ được cài đặt như một dependencies)

Đối với ứng dụng này, chúng ta sẽ sử dụng ruby, một ngôn ngữ mà chúng ta chưa đề cập đến và chúng ta sẽ không thực sự đi vào chi tiết về nó, tôi sẽ cố gắng tìm một số tài liệu tham khảo tốt và thêm chúng bên dưới.

Cài đặt với

`curl -L https://get.rvm.io | bash -s stable`

Bạn có thể được yêu cầu thêm khóa với lệnh trên, hãy làm theo các bước đó.

Để sử dụng rvm, chúng ta cần làm như sau:

`source /home/vagrant/.rvm/scripts/rvm`

và cuối cùng, cài đặt nó với

`rvm install ruby-2.7`

Kiểm tra cài đặt và phiên bản với

`ruby --version`

Tiếp theo, chúng ta cần framework Ruby on Rails, có thể cài đặt bằng lệnh sau.

`gem install rails`

Tiếp theo, chúng ta cần git và có thể cài đặt với

`sudo yum install git`

Chỉ để ghi lại và không chắc chắn liệu nó có cần thiết hay không, tôi cũng đã cài đặt Redis trên máy của mình vì tôi đang làm một việc khác nhưng nó vẫn có thể cần thiết, vì vậy có thể thực hiện các bước sau.

```
sudo yum install epel-release
sudo yum install redis
```

Nó có thể liên quan đến turbo streams nhưng tôi không có thời gian để tìm hiểu thêm về ruby on rails.

Bây giờ hãy tạo ứng dụng của chúng ta (tôi đã trải qua rất nhiều để đảm bảo các bước này hoạt động trên hệ thống của tôi nên chúc bạn may mắn)

tạo ứng dụng với lệnh sau

`rails new myapp --skip-turbolinks --skip-spring --skip-test-unit -d mysql`

tiếp theo, chúng ta sẽ tạo cơ sở dữ liệu và schema:

```
cd myapp
bundle exec rake db:create
bundle exec rake db:migrate
```

Chúng ta có thể chạy ứng dụng của mình với `bundle exec rails server -b 0.0.0.0`

![](../../images/day06-2.png)

Sau đó, mở trình duyệt để truy cập vào box đó, tôi đã phải thay đổi mạng VM VirtualBox của mình thành bridged thay vì NAT để có thể điều hướng đến nó thay vì sử dụng vagrant ssh.

![](../../images/day06-3.png)

Bây giờ chúng ta cần **scaffold** một mô hình cơ bản

Scaffold là một tập hợp các tệp được tạo tự động tạo thành cấu trúc cơ bản của một dự án Rails.

Chúng ta làm điều này với các lệnh sau:

```
bundle exec rails generate scaffold Bootcamp name:string description:text dates:string
bundle exec rake db:migrate
```

![](../../images/day06-4.png)

Thêm một route mặc định vào config/routes.rb

`root bootcamps#index`

![](../../images/day06-5.png)

Bây giờ chỉnh sửa app/views/bootcamps/show.html.erb và làm cho trường mô tả là một trường raw. Thêm đoạn mã dưới đây.

```
<p>
  <strong>Description:</strong>
  <%=raw @bootcamp.description %>
</p>
```

Lý do tại sao điều này có liên quan là việc sử dụng raw trong trường mô tả có nghĩa là trường này bây giờ trở thành một mục tiêu tiềm năng cho XSS. Hoặc cross-site scripting.

Điều này có thể được giải thích rõ hơn với video [What is Cross-Site Scripting?](https://youtu.be/DxsmEXicXEE)

Phần còn lại của Bootcamp tiếp tục thêm chức năng tìm kiếm, điều này cũng làm tăng khả năng tấn công XSS và đây là một ví dụ tuyệt vời khác về một cuộc tấn công demo mà bạn có thể thử trên một [ứng dụng dễ bị tấn công](https://www.softwaretestinghelp.com/cross-site-scripting-xss-attack-test/).

### Tạo chức năng tìm kiếm

Trong app/controllers/bootcamps_controller.rb, chúng ta sẽ thêm logic sau vào phương thức index:

```
def index
  @bootcamps = Bootcamp.all
  if params[:search].to_s != ''
    @bootcamps = Bootcamp.where("name LIKE '%#{params[:search]}%'")
  else
    @bootcamps = Bootcamp.all
  end
end
```

Trong app/views/bootcamps/index.html.erb, chúng ta sẽ thêm trường tìm kiếm:

```
<h1>Search</h1>
<%= form_tag(bootcamps_path, method: "get", id: "search-form") do %>
  <%= text_field_tag :search, params[:search], placeholder: "Search Bootcamps" %>
  <%= submit_tag "Search Bootcamps"%>
<% end %>

<h1>Listing Bootcamps</h1>
```

Cảm ơn rất nhiều đến [DevSecOps.org](https://www.devsecops.org/), đây là nơi tôi tìm thấy hướng dẫn cũ nhưng tuyệt vời với một vài điều chỉnh ở trên, còn rất nhiều thông tin khác có thể tìm thấy ở đó.

Với hướng dẫn dài hơn dự kiến này, tôi sẽ chuyển sang các phần và tác giả tiếp theo để làm nổi bật cách không làm điều này và cách đảm bảo chúng ta không phát hành mã xấu hoặc các lỗ hổng ra ngoài môi trường.

## Tài liệu tham khảo 

- [devsecops.org](https://www.devsecops.org/)

- [TechWorld with Nana - What is DevSecOps? DevSecOps explained in 8 Mins](https://www.youtube.com/watch?v=nrhxNNH5lt0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=1&t=19s)

- [What is DevSecOps?](https://www.youtube.com/watch?v=J73MELGF6u0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=2&t=1s)

- [freeCodeCamp.org - Web App Vulnerabilities - DevSecOps Course for Beginners](https://www.youtube.com/watch?v=F5KJVuii0Yw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=3&t=67s)

- [The Importance of DevSecOps and 5 Steps to Doing it Properly (DevSecOps EXPLAINED)](https://www.youtube.com/watch?v=KaoPQLyWq_g&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=4&t=13s)

- [Continuous Delivery - What is DevSecOps?](https://www.youtube.com/watch?v=NdvMUcWNlFw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=5&t=6s)

- [Cloud Advocate - What is DevSecOps?](https://www.youtube.com/watch?v=a2y4Oj5wrZg&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=6)

- [Cloud Advocate - DevSecOps Pipeline CI Process - Real world example!](https://www.youtube.com/watch?v=ipe08lFQZU8&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=7&t=204s)

Hẹn gặp lại bạn vào [Ngày 7](day07.md) nơi chúng ta sẽ bắt đầu một phần mới về Lập trình An toàn.

