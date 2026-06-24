---
title: '#90DaysOfDevOps - Bức tranh toàn cảnh: Quản lý cấu hình - Ngày 63'
published: false
description: 90DaysOfDevOps - The Big Picture Configuration Management
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048711
---

## Bức tranh toàn cảnh: Quản lý cấu hình

Ngay sau phần nói về Cơ sở hạ tầng dưới dạng mã, có nhiều khả năng sẽ có một số điểm giao khi chúng ta nói về Quản lý cấu hình hoặc Quản lý cấu hình ứng dụng.

Quản lý cấu hình là quá trình duy trì các ứng dụng, hệ thống và máy chủ ở trạng thái mong muốn. Sự trùng lặp với Cơ sở hạ tầng dưới dạng mã là IaC sẽ đảm bảo cơ sở hạ tầng của bạn ở trạng thái mong muốn nhưng sau đó, đặc biệt là terraform sẽ không quản lý trạng thái mong muốn của cài đặt hệ điều hành hoặc ứng dụng của bạn, đó là nơi các công cụ quản lý cấu hình xuất hiện đảm bảo rằng hệ thống và ứng dụng hoạt động theo cách được mong đợi khi có thay đổi.

Quản lý cấu hình giúp bạn không thực hiện các thay đổi nhỏ hoặc lớn mà không có sự quản lý bằng tài liệu.

### Tình huống: Tại sao bạn muốn sử dụng Quản lý cấu hình

Kịch bản hoặc lý do bạn muốn sử dụng Quản lý cấu hình, hãy gặp quản trị viên hệ thống của chúng ta, người làm việc trên tất cả các hệ thống - anh ấy tên là Dũng.

Điều gì sẽ xảy ra nếu hệ thống của họ gặp sự cố, nếu có hỏa hoạn, một máy chủ hoạt động không tốt? Dũng biết chính xác phải làm gì anh ấy có thể khắc phục sự cố đó một cách dễ dàng. Tuy nhiên, các vấn đề sẽ trở nên khó khăn với Dũng nếu nhiều máy chủ bắt đầu gặp sự cố, đặc biệt khi bạn có môi trường lớn và ngày càng mở rộng. Đây là lý do tại sao Dũng cần phải có một công cụ quản lý cấu hình. Các công cụ Quản lý Cấu hình có thể giúp Dũng trông giống như một ngôi sao nhạc rock, tất cả những gì anh ấy phải làm là định cấu hình đúng mã cho phép anh ấy đưa ra hướng dẫn về cách thiết lập từng mã các máy chủ một cách hiệu quả và ở quy mô lớn.

### Công cụ quản lý cấu hình

Có sẵn nhiều công cụ quản lý cấu hình và mỗi công cụ đều có các tính năng cụ thể giúp công cụ này hoạt động tốt hơn trong một số tình huống so với các công cụ khác.

![](../../Days/Images/Day63_config1.png)

Ở giai đoạn này, chúng ta sẽ xem nhanh các tùy chọn trong hình trên trước khi đưa ra lựa chọn chúng ta sẽ sử dụng tùy chọn nào.

- **Chef**

   - Chef đảm bảo cấu hình được áp dụng nhất quán trong mọi môi trường, ở mọi quy mô lớn với cơ sở hạ tầng tự động hóa.
   - Chef là một công cụ mã nguồn mở được phát triển bởi OpsCode được viết bằng Ruby và Erlang.
   - Chef phù hợp nhất cho các tổ chức có cơ sở hạ tầng không đồng nhất và đang tìm kiếm các giải pháp hoàn thiện.
   - Recipes và Cookbooks xác định mã cấu hình cho hệ thống của bạn.
   - Pro - Có sẵn một bộ sưu tập lớn các recipes
   - Pro - Tích hợp tốt với Git, cung cấp khả năng kiểm soát phiên bản mạnh mẽ
   - Con - Phải học rất nhiều và cần một lượng thời gian đáng kể để làm quen.
   - Con - Máy chủ chính không có nhiều quyền kiểm soát.
   - Kiến trúc - Server/Client
   - Dễ dàng thiết lập - Trung bình
   - Language - Procedural - Chỉ định cách thực hiện tác vụ


- **Puppet**

   - Puppet là công cụ quản lý cấu hình hỗ trợ triển khai tự động.
   - Puppet được xây dựng bằng Ruby và sử dụng DSL để viết bảng kê khai.
   - Puppet cũng hoạt động tốt với cơ sở hạ tầng không đồng nhất, nơi tập trung vào khả năng mở rộng.
   - Pro - Cộng đồng hỗ trợ lớn.
   - Pro - Cơ chế báo cáo được viết rất tốt.
   - Con - Các nhiệm vụ nâng cao yêu cầu kiến ​​thức về ngôn ngữ Ruby.
   - Con - Máy chủ chính không có nhiều quyền kiểm soát.
   - Kiến trúc - Server/Client
   - Dễ thiết lập - Trung bình
   - Language - Declarative - Chỉ định những việc cần làm

- **Ansible**

   - Ansible là một công cụ tự động hóa CNTT giúp tự động hóa việc quản lý cấu hình, triển khai trên đám mây, triển khai và điều phối.
   - Cốt lõi của Ansible playbooks được viết bằng YAML (Nên làm một phần về YAML như chúng ta đã thấy điều này một vài lần)
   - Ansible hoạt động tốt khi các môi trường tập trung vào việc khởi động và chạy mọi thứ một cách nhanh chóng.
   - Hoạt động trên playbook cung cấp hướng dẫn cho máy chủ của bạn.
   - Pro - Không agents trên trên các node từ xa.
   - Pro - YAML rất dễ học.
   - Con - Tốc độ thực hiện thường kém hơn các công cụ khác (Nhanh hơn Dũng tự làm thủ công)
   - Con - YAML không mạnh bằng Ruby nhưng học nó dễ dàng hơn Ruby.
   - Kiến trúc - client only
   - Dễ thiết lập - Rất dễ
   - Language - Procedural - Chỉ định cách thực hiện tác vụ

- **SaltStack**

   - SaltStack là một công cụ dựa trên CLI giúp tự động hóa việc quản lý cấu hình và thực thi từ xa.
   - SaltStack dựa trên Python trong khi các hướng dẫn được viết bằng YAML hoặc DSL của nó.
   - Hoàn hảo cho các môi trường ưu tiên khả năng mở rộng và khả năng phục hồi.
   - Pro - Dễ sử dụng khi thiết lập và chạy
   - Pro - Cơ chế báo cáo tốt
   - Con - Giai đoạn thiết lập khó khăn
   - Con - Giao diện người dùng web mới kém phát triển hơn nhiều so với các giao diện người dùng khác.
   - Kiến trúc - Server/Client
   - Dễ thiết lập - Trung bình
   - Language - Declarative - Chỉ định những việc cần làm

### Ansible vs Terraform

Công cụ mà chúng ta sẽ sử dụng cho phần này sẽ là Ansible (Dễ sử dụng và yêu cầu cơ bản về ngôn ngữ dễ dàng hơn).

Tôi nghĩ rằng điều quan trọng là phải đề cập đến một số khác biệt giữa Ansible và Terraform trước khi chúng ta tìm hiểu sâu hơn về công cụ này.

|                | Ansible                                                      | Terraform                                                        |
| -------------- | ------------------------------------------------------------ | ---------------------------------------------------------------- |
| Loại | Ansible là công cụ quản lý cấu hình | Terraform là công cụ điều phối |
| Cơ sở hạ tầng | Ansible cung cấp hỗ trợ cho cơ sở hạ tầng có thể thay đổi (mutable) | Terraform cung cấp hỗ trợ cho cơ sở hạ tầng không thể thay đổi (immutable) |
| Ngôn ngữ | Ansible tuân theo ngôn ngữ thủ tục | Terraform tuân theo ngôn ngữ khai báo |
| Cung cấp | Ansible cung cấp cung cấp một phần (VM, Mạng, Lưu trữ) | Terraform cung cấp cung cấp rộng rãi (VM, Mạng, Lưu trữ) |
| Đóng gói | Ansible cung cấp hỗ trợ đầy đủ cho việc đóng gói & tạo template | Terraform cung cấp hỗ trợ một phần cho đóng gói & tạo template |
| Quản lý vòng đời | Ansible không có quản lý vòng đời | Terraform phụ thuộc nhiều vào vòng đời và quản lý trạng thái |

## Tài liệu tham khảo

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)

Hẹn gặp lại vào [ngày 64](day64.md)
