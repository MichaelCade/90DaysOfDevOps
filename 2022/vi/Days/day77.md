---
title: '#90DaysOfDevOps - Bức tranh toàn cảnh: Giám sát - Ngày 77'
published: false
description: 90DaysOfDevOps - Bức tranh toàn cảnh: Giám sát
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048715
---

## Bức tranh toàn cảnh: Giám sát 

Trong tuần này chúng ta sẽ nói về giám sát, nó là gì và tại sao chúng ta cần nó?

### Giám sát là gì?

Giám sát là quá trình theo dõi chặt chẽ toàn bộ cơ sở hạ tầng

### và tại sao chúng ta cần nó?

Giả sử chúng ta đang quản lý một nghìn máy chủ, bao gồm nhiều loại máy chủ chuyên dụng như máy chủ ứng dụng, máy chủ cơ sở dữ liệu và máy chủ web. Chúng ta cũng có thể có các dịch vụ bổ sung và các nền tảng khác, bao gồm cả dịch vụ đám mây công cộng và Kubernetes.

![](Images/Day77_Monitoring1.png)

Chúng ta chịu trách nhiệm đảm bảo rằng tất cả các dịch vụ, ứng dụng và tài nguyên trên máy chủ đều hoạt động bình thường.

![](Images/Day77_Monitoring2.png)

Chúng ta làm điều đó như thế nào? có ba cách:

- Đăng nhập thủ công vào tất cả các máy chủ và kiểm tra tất cả dữ liệu về processes và tài nguyên, dịch vụ.
- Viết script đăng nhập vào máy chủ và kiểm tra dữ liệu.

Cả hai lựa chọn này đều tạo ra một khối lượng công việc đáng cho chúng ta

Tùy chọn thứ ba dễ dàng hơn, chúng ta có thể sử dụng giải pháp giám sát có sẵn trên thị trường.

Nagios và Zabbix là những giải pháp sẵn có, chúng cho phép chúng ta nâng cấp cơ sở hạ tầng giám sát của mình để quản lý nhiều máy chủ như chúng ta muốn.

### Nagios

Nagios là một công cụ giám sát cơ sở hạ tầng được phát triển bởi một công ty cùng tên. Phiên bản mã nguồn mở của công cụ này được gọi là Nagios core trong khi phiên bản thương mại có tên là Nagios XI. [Trang web của Nagios](https://www.nagios.org/)

Công cụ này cho phép chúng ta giám sát các máy chủ của mình và xem liệu chúng có đang được sử dụng hiệu quả hoặc có bất kỳ tasks lỗi nào cần giải quyết hay không.

![](Images/Day77_Monitoring3.png)

Về cơ bản, việc giám sát cho phép chúng ta đạt được hai mục tiêu này, kiểm tra trạng thái máy chủ và dịch vụ cũng như xác định tình trạng cơ sở hạ tầng. Nó cũng cung cấp cho chúng ta cái nhìn high-level về cơ sở hạ tầng hoàn chỉnh để xem liệu máy chủ có hoạt động hay không nếu các ứng dụng hoạt động hoạt động bình thường và các máy chủ web có thể truy cập được hay không.

Nó sẽ cho chúng ta biết rằng ổ đĩa của chúng ta đã tăng 10 phần trăm trong 10 tuần qua trên một máy chủ cụ thể, rằng nó sẽ cạn kiệt hoàn toàn trong vòng bốn hoặc năm ngày tới và nếu chúng ta không phản hồi sớm, nó sẽ cảnh báo khi đĩa hoặc máy chủ đang ở trạng thái nguy hiểm để chúng ta có thể thực hiện các hành động thích hợp nhằm tránh những sự cố ngừng hoạt động có thể xảy ra.

Trong trường hợp này, chúng ta có thể giải phóng một số dung lượng ổ đĩa và đảm bảo rằng máy chủ của chúng ta không bị lỗi và người dùng không bị ảnh hưởng.

Một câu hỏi khó đối với hầu hết các kỹ sư giám sát là chúng ta sẽ giám sát những gì? và có thể là chúng ta không giám sát những gì?

Mỗi hệ thống đều có một số tài nguyên, trong đó chúng ta nên theo dõi chặt chẽ tài nguyên nào và tài nguyên nào và chúng ta có thể nhắm mắt làm ngơ. Chẳng hạn như có cần thiết phải giám sát việc sử dụng CPU không, câu trả lời là rõ ràng là có, tuy nhiên đó vẫn là một quyết định cần phải được đưa ra. Có cần thiết phải theo dõi số lượng cổng mở trong hệ thống hay không, điều đó tùy thuộc vào tình huống. Nếu đó là một máy chủ đa năng thì chúng ta có thể sẽ không phải làm vậy, nhưng nếu đó là máy chủ web thì chúng ta nên làm điều đó.

### Giám sát liên tục

Giám sát không phải là một khái niệm mới và thậm chí giám sát liên tục đã được nhiều doanh nghiệp áp dụng trong nhiều năm qua.

Có ba lĩnh vực trọng tâm chính khi nói đến giám sát.

- Giám sát cơ sở hạ tầng
- Giám sát ứng dụng
- Giám sát mạng

Điều cần lưu ý là có rất nhiều công cụ có sẵn, chúng ta đã đề cập đến hai hệ thống và công cụ chung trong phần này nhưng còn rất nhiều công cụ khác. Lợi ích thực sự của một giải pháp giám sát thể hiện bạn dành thời gian để trả lời câu hỏi chúng ta nên giám sát những gì và không nên giám sát những gì?

Chúng ta có thể bật giải pháp giám sát trên bất kỳ nền tảng nào của chúng ta và nó sẽ bắt đầu lấy thông tin nhưng nếu thông tin đó quá nhiều thì bạn sẽ gặp khó khăn trong việc thấy các lợi ích từ giải pháp đó, bạn phải dành thời gian để định cấu hình nó.

Trong phần tiếp theo, chúng ta sẽ thực hành với một công cụ giám sát và xem chúng ta có thể bắt đầu giám sát những gì.

## Tài liệu tham khảo

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)

Hẹn gặp lại vào [ngày 78](day78.md)
