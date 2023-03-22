---
title: '#90DaysOfDevOps - Bức tranh lớn: DevOps và Linux - Ngày 14'
published: false
description: 90DaysOfDevOps - Bức tranh lớn: DevOps và Linux
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049033
---

## Bức tranh lớn: DevOps và Linux

Linux và DevOps có văn hóa và nhiều quan điểm rất giống nhau; cả hai đều tập trung vào khả năng tùy biến và khả năng mở rộng. Cả hai khía cạnh này của Linux đều có tầm quan trọng đặc biệt đối với DevOps.

Rất nhiều công nghệ bắt đầu với Linux, đặc biệt nếu chúng liên quan đến phát triển phần mềm hoặc quản lý cơ sở hạ tầng.

Cũng như rất nhiều dự án mã nguồn mở, đặc biệt là các công cụ DevOps, được thiết kế để chạy trên Linux ngay từ đầu.

Từ góc độ DevOps hoặc bất kỳ góc độ nào từ việc vận hành, bạn sẽ tiếp xúc với Linux nếu không muốn nói là chủ yếu. Có một nơi dành cho WinOps nhưng phần lớn thời gian bạn sẽ quản lý và triển khai các máy chủ Linux.

Tôi đã sử dụng Linux hàng ngày trong vài năm nhưng máy tính để bàn của tôi luôn là macOS hoặc Windows. Tuy nhiên, khi tôi bắt đầu làm việc dưới role Cloud Native hiện tại, tôi đã cố gắng đảm bảo máy tính xách tay của tôi hoàn toàn chạy trên Linux, và mặc dù tôi vẫn cần Windows cho các ứng dụng trong công việc và rất nhiều thiết bị âm thanh và video gear của tôi không chạy trên Linux, tôi đã buộc mình phải dùng một máy tính chạy trên Linux để hiểu rõ hơn về những thứ mà chúng ta sẽ tiếp xúc trong 7 ngày tới.

## Bắt đầu

Tôi không khuyến nghị bạn làm giống như tôi vì có nhiều tùy chọn dễ dàng hơn nhưng tôi sẽ nói rằng việc thực hiện điều đó buộc bạn phải học nhanh hơn cách làm cho mọi thứ hoạt động trên Linux.

Trong phần lớn thời gian 7 ngày này, tôi sẽ triển khai máy ảo (Virtual Machine) sử dụng Virtual Box trên máy Windows của mình. Tôi cũng sẽ triển khai một bản phân phối Linux (Linux distro), trong khi phần lớn các máy chủ Linux mà bạn sẽ quản lý có thể sẽ là các máy chủ không có GUI và mọi thứ đều dựa trên shell. Tuy nhiên, như tôi đã nói ở phần đầu, rất nhiều các công cụ mà chúng ta đã đề cập trong suốt 90 ngày bắt đầu chạy trên Linux. Do đó, tôi cũng khuyến khích bạn tham gia vào việc sử dụng Linux với máy tính của bạn để có cùng một trải nghiệm học tập với tôi.

Đối với phần còn lại của bài này, chúng ta sẽ tập trung vào việc thiết lập và chạy một máy ảo Ubuntu Desktop trong Virtual Box. Đầu tiên, hãy tải xuống [Virtual Box](https://www.virtualbox.org/) và download [Ubuntu ISO](https://ubuntu.com/download) bản mới nhất và tiếp tục hoàn thiện môi trường của của chúng ta. Nó không được coi là DevOps đúng không?

Một lý do tốt khác để sử dụng hầu hết các bản phân phối Linux là chúng miễn phí và là phần mềm mã nguồn mở. Chúng ta cũng đang chọn Ubuntu vì nó có lẽ là bản phân phối được sử dụng rộng rãi nhất được triển khai bỏ qua các thiết bị di động và máy chủ RedHat Enterprise dành cho doanh nghiệp. Tôi có thể sai nhưng với CentOS, tôi nghĩ rằng Ubuntu có một vị trí cao trong danh sách và nó cực kỳ đơn giản.

## Giới thiệu HashiCorp Vagrant

Vagrant là một tiện ích CLI quản lý vòng đời của các máy ảo của bạn. Chúng ta có thể sử dụng vagrant để khởi động và tắt các máy ảo trên nhiều nền tảng khác nhau bao gồm vSphere, Hyper-v, Virtual Box và cả Docker. Sẽ có nhiều các nhà cung cấp khác nhưng chúng ta sẽ tiếp tục với Virtual Box.

Điều đầu tiên chúng ta cần làm là cài đặt Vagrant trên máy của bạn, khi truy cập trang tải xuống, bạn sẽ thấy tất cả các hệ điều hành được liệt kê để bạn lựa chọn tại [trang download của HashiCorp Vagrant](https://www.vagrantup.com/downloads). Tôi đang sử dụng Windows nên đã download tệp nhị phân cho hệ thống của mình và cài đặt tệp này lên hệ thống của mình.

Tiếp theo, chúng ta cũng cần cài đặt [Virtual Box](https://www.virtualbox.org/wiki/Downloads). Một lần nữa, phần mềm này cũng có thể được cài đặt trên nhiều hệ điều hành khác nhau và một lý do chính đáng để chọn 2 công nghệ này là vì chúng ta có thể sử dụng trên hầu hết tất cả các hệ điều hành.

Cả hai đều khá dễ đề cài đặt và đều có những cộng đồng tuyệt vời xung quanh chúng, vì vậy hãy liên hệ nếu bạn gặp vấn đề.

## VAGRANTFILE đầu tiên

VAGRANTFILE mô tả loại máy chúng ta muốn triển khai. Nó cũng định nghĩa cấu hình và khởi tạo máy này.

Khi nói đến việc lưu các tệp này và sắp xếp các VAGRANTFILE của bạn, tôi có xu hướng đặt chúng vào các thư mục của chúng trong không gian làm việc của mình. Bạn có thể thấy bên dưới giao diện này trên hệ thống của tôi. Hy vọng rằng bạn cũng sẽ thử sử dụng Vagrant và thấy sự dễ dàng khi muốn thay đổi hệ thống, rồi cuối cùng bạn sẽ nghiện vọc vạch các bản distro khác nhau của Linux cho Máy tính để bàn Linux.

![](../../Days/Images/Day14_Linux1.png)

Hãy cùng nhìn lại VAGRANTFILE đó và xem những gì chúng ta đang xây dựng.

```ruby:Vagrantfile
Vagrant.configure("2") do |config|
    config.vm.box = "chenhan/ubuntu-desktop-20.04"
        config.vm.provider :virtualbox do |v|
        v.memory  = 8096
        v.cpus    = 4
        v.customize ["modifyvm", :id, "--vram", "128"]
    end
end
```

Đây là một VAGRANTFILE đơn giản. Chúng ta muốn nói rằng chúng ta muốn có một box cụ thể, một box có thể là hình ảnh công khai hoặc bản dựng riêng của hệ thống mà bạn đang tìm kiếm. Bạn có thể tìm thấy một danh sách dài các box công khai có sẵn tại đây trong [danh mục công khai của Vagrant box](https://app.vagrantup.com/boxes/search)

Dòng tiếp theo, chúng ta nói rằng chúng ta muốn sử dụng một nhà cung cấp cụ thể và trong trường hợp này là `VirtualBox`. Chúng ta cũng muốn bộ nhớ của máy là `8GB` và số lượng CPU là `4`. Theo kinh nghiệm của tôi, bạn có thể sẽ phải thêm dòng sau nếu bạn gặp sự cố hiển thị. Tôi sẽ tăng bộ nhớ này lên đến `128MB` nhưng nó hoàn toàn phụ thuộc vào cấu hình hệ thống của bạn.

```ruby
v.customize ["modifyvm", :id, "--vram", ""]
```

Tôi cũng đã đặt một bản sao của tệp vagrant cụ thể này trong [thư mục Linux](../../Days/Linux/VAGRANTFILE)

## Tạo Linux Desktop

Bây giờ chúng ta đã sẵn sàng để khởi động và chạy máy tính của mình trong terminal của máy trạm. Tôi đang sử dụng PowerShell trên máy Windows của mình. Điều hướng đến thư mục dự án và là nơi ta đặt VAGRANTFILE. Bạn có thể nhập lệnh `vagrant up` và nếu mọi thứ diễn ra bình thường, bạn sẽ thấy một cái gì đó như thế này.

![](../../Days/Images/Day14_Linux2.png)

Một điều nữa cần bổ sung ở đây là network sẽ được đặt thành `NAT` trên máy ảo của bạn. Ở giai đoạn này, chúng ta không cần biết về NAT và tôi dự định sẽ nói về nó sâu hơn khi trong phần Network. Hãy coi rằng đây là cài đặt đơn giản để đưa máy ảo vào mạng của bạn, đây cũng là chế độ kết nối mạng mặc định trên Virtual Box. Bạn có thể tìm hiểu thêm trong [Virtual Box documentation](https://www.virtualbox.org/manual/ch06.html#network_nat)

Sau khi hoàn tất quá trình `vagrant up`, chúng ta có thể sử dụng `vagrant ssh` để truy cập vào terminal của máy ảo mới của mình.

![](../../Days/Images/Day14_Linux3.png)

Đây là nơi chúng ta sẽ khám phá trong vài ngày tới. Tôi cũng muốn đi sâu vào một số tùy biến cho máy trạm dành cho nhà phát triển của bạn mà tôi cũng đã làm và nó giúp cuộc sống của bạn đơn giản hơn nhiều khi sử dụng nó hàng ngày và tất nhiên, bạn có thực sự thích DevOps khi bạn sử dụng một terminal tiêu chuẩn và không cool ngầu?

Để xác nhận trong Virtual Box, bạn sẽ thấy lời nhắc đăng nhập khi bạn chọn máy ảo của mình.

![](../../Days/Images/Day14_Linux4.png)

Nếu bạn đã đi được đến đây và đang tự hỏi "TÊN NGƯỜI DÙNG & MẬT KHẨU LÀ GÌ?"

- Username = vagrant
- Password = vagrant

Ngày mai chúng ta sẽ tìm hiểu về một số lệnh và tác dụng của chúng, Terminal sẽ là nơi để biến mọi thứ thành hiện thực.

## Tài liệu tham khảo

- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

Sẽ có rất nhiều tài nguyên mà và giống như các tài nguyên về Go, tôi sẽ cố gắng sử dụng nội dung MIỄN PHÍ để tất cả chúng ta có thể tham gia và học tập.

Như tôi đã đề cập, trong ngày tiếp theo, chúng ta sẽ tìm hiểu các lệnh có thể hàng ngày trong môi trường Linux.

Hẹn gặp lại các bạn vào [Ngày 15](day15.md)
