---
title: '#90DaysOfDevOps - Bức tranh toàn cảnh: Quản lý log - Ngày 79'
published: false
description: 90DaysOfDevOps - Bức tranh toàn cảnh: Quản lý log
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049057
---

## Bức tranh toàn cảnh: Quản lý Log

Tiếp nối về giải pháp giám sát cơ sở hạ tầng, quản lý log là một mảnh ghép khác về khả năng quan sát tổng thể.

### Quản lý và tổng hợp Log

Hãy nói về hai khái niệm cốt lõi, trong đó khái niệm đầu tiên là tổng hợp log và cách thu thập và gắn tag các log ứng dụng từ nhiều dịch vụ khác nhau trong một bảng điều khiển duy nhất có thể tìm kiếm dễ dàng.

Một trong những hệ thống đầu tiên cần được xây dựng trong một hệ thống quản lý hiệu suất ứng dụng (APM) là tổng hợp log. Quản lý hiệu suất ứng dụng là phần của vòng đời DevOps khi mọi thứ đã được xây dựng và triển khai, bạn cần đảm bảo rằng chúng vẫn hoạt động liên tục, có đủ tài nguyên được phân bổ và lỗi không hiển thị cho người dùng. Trong hầu hết các triển khai trên môi trường sản xuất, nhiều sự kiện liên quan tạo ra log trên các dịch vụ. Ví dụ, với Google, một lần tìm kiếm duy nhất có thể phát sinh truy cập vào mười dịch vụ khác nhau trước khi kết quả được trả về cho người dùng. Nếu bạn nhận được kết quả tìm kiếm không mong đợi, điều đó có nghĩa là có vấn đề logic tại bất kỳ một trong mười dịch vụ đó. Tổng hợp log giúp các công ty như Google chẩn đoán, phát hiện các vấn đề trong môi trường sản xuất. Họ đã xây dựng một bảng điều khiển duy nhất, nơi họ có thể ánh xạ mọi request tới một ID duy nhất. Khi bạn tìm kiếm trên Google, request tìm kiếm của bạn sẽ nhận được một ID duy nhất và sau mỗi lần tìm kiếm đó đi qua một dịch vụ khác nhau, dịch vụ đó sẽ kết nối ID đó với những xử lý trong dịch vụ đó.

Đây là điều cơ bản của một nền tảng tổng hợp log tốt, giúp thu thập log từ mọi nơi tạo ra chúng và có thể dễ dàng tìm kiếm lại trong trường hợp có lỗi phát sinh.

### Ứng dụng ví dụ

Ứng dụng ví dụ của chúng ta là một ứng dụng web gồm có frontend và backend lưu trữ dữ liệu quan trọng trong cơ sở dữ liệu MongoDB.

Nếu một người dùng nói với chúng ta rằng trang web trắng toàn bộ và trả lại một thông báo lỗi, chúng ta sẽ gặp khó khăn trong việc chẩn đoán vấn đề với stack hiện tại. Người dùng sẽ cần phải gửi lỗi một cách thủ công và chúng ta cần phải kết hợp nó với log liên quan trong ba dịch vụ khác.

### ELK

Hãy cùng tìm hiểu về ELK, một stack tổng hợp log mã nguồn mở phổ biến được đặt theo tên ba thành phần của nó là Elasticsearch, Logstash và Kibana. Chúng ta sẽ cài đặt nó trong cùng một môi trường với ứng dụng ví dụ.

Ứng dụng web sẽ kết nối với giao diện frontend trước, sau đó kết nối với backend, backend sẽ gửi log đến Logstash và sau đó ba thành phần này sẽ hoạt động.

### Các thành phần của ELK

Elasticsearch, Logstash và Kibana: tất cả các dịch vụ gửi log đến Logstash, Logstash lấy những log này là văn bản được phát ra bởi ứng dụng. Ví dụ, trong ứng dụng web, khi bạn truy cập vào một trang web, trang web có thể ghi log truy cập của khách truy cập này vào trang này vào thời gian này, và đó là một ví dụ về một thông báo log. Những log đó sẽ được gửi đến Logstash.

Logstash sau đó sẽ trích xuất các thông tin từ chúng. Ví dụ, thông báo log ghi nhận người dùng đã **làm gì** vào **thời gian nào**. Bạn có thể tìm kiếm chúng một cách dễ dàng và tìm tất cả các yêu cầu được thực hiện bởi một người dùng cụ thể.

Logstash không lưu trữ thông tin mà lưu trữ trong Elasticsearch, một cơ sở dữ liệu hiệu quả để truy vấn văn bản. Elasticsearch hiển thị kết quả qua Kibana. Kibana là một máy chủ web kết nối với Elasticsearch và cho phép các quản trị viên hoặc các thành viên trong nhóm, như kỹ sư trực, xem các log trong môi trường sản xuất bất cứ khi nào có lỗi lớn. Bạn, với tư cách là quản trị viên, sẽ kết nối với Kibana, và Kibana sẽ truy vấn Elasticsearch để tìm các log khớp với yêu cầu của bạn.

Bạn có thể nhập vào thanh tìm kiếm của Kibana để tìm lỗi và Kibana sẽ yêu cầu Elasticsearch tìm các thông báo chứa chuỗi "error". Elasticsearch sẽ trả về các kết quả đã được Logstash lưu trữ. Logstash đã nhận những kết quả đó từ tất cả các dịch vụ khác.

### Cách sử dụng ELK để chẩn đoán vấn đề trong môi trường sản xuất

Một người dùng nói rằng họ thấy mã lỗi 1234567 khi họ cố gắng thực hiện một thao tác. Với cấu hình ELK, chúng ta sẽ vào Kibana, nhập 1234567 vào thanh tìm kiếm và nhấn Enter. Kết quả sẽ hiển thị các log tương ứng và một trong số đó có thể hiển thị "internal server error returning 1234567". Chúng ta sẽ thấy dịch vụ nào đã tạo ra log đó, và thời gian log đó được ghi lại. Chúng ta có thể xem các thông báo phía trên và dưới log đó trong backend để hiểu rõ hơn về những gì đã xảy ra với yêu cầu của người dùng. Chúng ta sẽ lặp lại quy trình này với các dịch vụ khác cho đến khi tìm ra nguyên nhân gây ra vấn đề.

### Bảo mật và truy cập vào log

Một phần quan trọng của việc quản lý log là đảm bảo rằng log chỉ hiển thị cho quản trị viên (hoặc những người và nhóm cần truy cập). Log có thể chứa thông tin nhạy cảm như token, chỉ những người dùng đã xác thực mới nên có quyền truy cập. Bạn không nên để Kibana lộ ra ngoài internet mà không có cách nào để xác thực.

### Ví dụ về các công cụ quản lý log

Một số nền tảng quản lý log bao gồm:

- Elasticsearch
- Logstash
- Kibana
- Fluentd - một lựa chọn mã nguồn mở phổ biến
- Datadog - hosted offering, thường được sử dụng tại các doanh nghiệp lớn
- LogDNA - hosted offering
- Splunk

Các nhà cung cấp dịch vụ đám mây cũng cung cấp các công cụ logging như AWS CloudWatch Logs, Microsoft Azure Monitor và Google Cloud Logging.

Quản lý log là một khía cạnh quan trọng của việc quan sát tổng thể các ứng dụng và môi trường cơ sở hạ tầng của bạn để chẩn đoán các vấn đề trong môi trường sản xuất. Việc cài đặt một giải pháp trọn gói như ELK hoặc CloudWatch khá đơn giản và nó giúp việc chẩn đoán và xử lý sự cố trong sản xuất dễ dàng hơn đáng kể.

## Tài liệu tham khảo

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)
- [Log Management for DevOps | Manage application, server, and cloud logs with Site24x7](https://www.youtube.com/watch?v=J0csO_Shsj0)
- [Log Management what DevOps need to know](https://devops.com/log-management-what-devops-teams-need-to-know/)
- [What is ELK Stack?](https://www.youtube.com/watch?v=4X0WLg05ASw)
- [Fluentd simply explained](https://www.youtube.com/watch?v=5ofsNyHZwWE&t=14s)

Hẹn gặp lại vào [ngày 80](day80.md).
