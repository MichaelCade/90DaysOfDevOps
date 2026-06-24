---
title: '#90DaysOfDevOps - Tự Động Hóa Thiết Lập Mạng - Ngày 24'
published: false
description: 90DaysOfDevOps - Tự Động Hóa Thiết Lập Mạng
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048805
---


## Tự động hóa mạng

### Khái niệm cơ bản về tự động hóa mạng

Mục đích của việc Tự động hóa mạng

- Đạt được sự linh hoạt
- Giảm chi phí
- Loại bỏ lỗi
- Tuân thủ các quy tắc, quy định (compliance)
- Quản lý tập trung

Quá trình áp dụng tự động hóa là riêng biệt cho từng doanh nghiệp. Không có một giải pháp nào phù hợp với tất cả các yêu cầu khi triển khai tự động hóa, khả năng xác định và nắm bắt phương pháp phù hợp nhất với tổ chức của bạn là rất quan trọng trong việc tiến tới duy trì hoặc tạo ra một môi trường linh hoạt hơn, trọng tâm luôn phải là giá trị kinh doanh và mục tiêu cuối cùng - trải nghiệm người dùng. (Chúng ta đã nói những điều tương tự ngay từ đầu về văn hóa DevOps và sự thay đổi văn hóa cũng như quy trình tự động mà điều này mang lại)

Để phân tích vấn đề này, bạn cần xác định bằng cách nào những nhiệm vụ hoặc quy trình mà bạn đang cố gắng tự động hóa sẽ giúp cải thiện trải nghiệm của người dùng cuối hoặc giá trị kinh doanh trong khi vẫn tuân theo phương pháp tiếp cận có hệ thống từng bước.

"Nếu bạn không biết mình đang đi đâu, thì bất kỳ con đường nào cũng sẽ đưa bạn đến đích."

Có một framework hoặc bản thiết kế mà bạn đang cố gắng để hoàn thành, biết rõ mục tiêu cuối cùng của mình là gì và sau đó làm việc từng bước để đạt được mục tiêu đó, đo lường mức độ thành công của việc tự động hóa ở các giai đoạn khác nhau dựa trên kết quả kinh doanh.

Xây dựng các khái niệm đã được mô hình hóa xung quanh các ứng dụng hiện có, không cần phải thiết kế các khái niệm xung quanh một mô hình giả tưởng vì chúng cần được áp dụng cho ứng dụng, dịch vụ và cơ sở hạ tầng của bạn. Vì vậy hãy bắt đầu xây dựng các khái niệm và mô hình hóa xung quanh cơ sở hạ tầng và ứng dụng hiện có của bạn.

### Cách tiếp cận việc Tự động hóa Mạng

Chúng ta nên xác định các tác vụ và thực hiện khám phá các yêu cầu thay đổi trong thiết lập mạng để bạn có danh sách các vấn đề và sự cố phổ biến nhất mà cần một giải pháp tự động hóa.

- Lập danh sách tất cả các yêu cầu thay đổi và quy trình công việc hiện đang được giải quyết theo cách thủ công.
- Xác định các hoạt động phổ biến, tốn thời gian và dễ mắc lỗi nhất.
- Ưu tiên các yêu cầu bằng dựa theo định hướng kinh doanh của doanh nghiệp.
- Nếu đây là bộ khung để xây dựng quy trình tự động hóa, thì cái gì phải tự động hóa, cái gì không.

Sau đó, chúng ta nên phân chia các nhiệm vụ và phân tích cách các chức năng mạng khác nhau hoạt động và tương tác với nhau.

- Nhóm Hạ tầng/Mạng nhận yêu cầu thay đổi ở nhiều lớp để triển khai ứng dụng.
- Dựa trên các dịch vụ mạng, hãy chia chúng thành các khu vực khác nhau và hiểu cách chúng tương tác với nhau.
   - Tối ưu hóa ứng dụng
   - ADC (Bộ điều khiển phân phối ứng dụng)
   - Tường lửa (Firewall)
   - DDI (DNS, DHCP, IPAM, v.v.)
   - Định tuyến
   - Các vấn đề khác
- Xác định các yếu tố phụ thuộc khác nhau để giải quyết các khác biệt về kinh doanh và văn hóa, đồng thời mang lại sự hợp tác giữa các nhóm.

- Chính sách tái sử dụng, xác định và đơn giản hóa các tác vụ, quy trình và đầu vào/đầu ra của dịch vụ có thể tái sử dụng.

- Định nghĩa các dịch vụ, quy trình và đầu vào/đầu ra khác nhau.
- Đơn giản hóa quy trình triển khai sẽ giảm thời gian hoàn thành cho cả khối lượng công việc mới và hiện có.
- Sau khi bạn có một quy trình tiêu chuẩn, quy trình đó có thể được sắp xếp theo trình tự và căn chỉnh theo các yêu cầu riêng lẻ để có cách tiếp cận và phân phối đa luồng.

Kết hợp các chính sách với các hoạt động kinh doanh cụ thể. Việc thực hiện chính sách này giúp gì cho doanh nghiệp? Tiết kiệm thời gian? Tiết kiệm tiền? Cung cấp một kết quả kinh doanh tốt hơn?

- Đảm bảo rằng các tác vụ dịch vụ có thể tương tác với nhau.
- Liên kết các nhiệm vụ dịch vụ gia tăng sao cho chúng phối hợp để tạo ra các dịch vụ kinh doanh.
- Cho phép việc linh hoạt trong liên kết lại các nhiệm vụ dịch vụ theo yêu cầu.
- Triển khai các dịch vụ tự làm việc và mở đường cho việc cải thiện hiệu quả hoạt động.
- Cho phép nhiều bộ kỹ năng công nghệ tiếp tục đóng góp vào việc giám sát và tuân thủ.

**Lặp đi lặp lại** các chính sách và quy trình, bổ sung và cải thiện trong khi vẫn duy trì tính khả dụng của dịch vụ.

- Bắt đầu bằng cách tự động hóa các nhiệm vụ hiện có.
- Làm quen với quy trình tự động hóa để bạn có thể xác định các lĩnh vực khác có thể hưởng lợi từ tự động hóa.
- Lặp đi lặp lại các sáng kiến tự động hóa của bạn, tăng dần sự linh hoạt trong khi vẫn duy trì tính khả dụng cần thiết.
- Thực hiện một cách tiếp cận tăng dần sẽ mở đường cho thành công!

Điều phối các dịch vụ mạng!

- Tự động hóa quy trình triển khai là cần thiết để phân phối ứng dụng nhanh chóng.
- Việc tạo ra một môi trường dịch vụ linh hoạt đòi hỏi phải quản lý các yếu tố khác nhau thông qua nhiều kỹ năng kỹ thuật.
- Chuẩn bị cho sự phối hợp từ đầu đến cuối cung cấp khả năng kiểm soát tự động hóa và thứ tự trong việc triển khai.


## Công cụ tự động hóa mạng

Tin tốt ở đây là phần lớn các công cụ chúng ta sử dụng ở đây cho tự động hóa Mạng nói chung giống với những công cụ mà chúng ta sẽ sử dụng cho các lĩnh vực tự động hóa khác đối với những gì chúng ta đã đề cập cho đến nay hoặc những gì chúng ta sẽ đề cập trong các phần sau.

Hệ điều hành - Như tôi đã vượt qua thử thách này, tôi đang thực hiện hầu hết bài học của mình với HĐH Linux, lý do đó đã được đưa ra trong phần Linux nhưng hầu như tất cả các công cụ mà chúng ta sẽ sử dụng mặc dù hôm nay có thể là các nền tảng đa hệ điều hành, tuy nhiên tất cả đều bắt đầu dưới dạng các ứng dụng hoặc công cụ dựa trên Linux.

Môi trường phát triển tích hợp (IDE) - Một lần nữa, không có nhiều điều để nói ở đây ngoài việc tôi sẽ đề xuất Visual Studio Code làm IDE xuyên suốt của bạn, nó cung cấp các plugin mở rộng có sẵn cho rất nhiều ngôn ngữ khác nhau.

Quản lý cấu hình - chúng ta chưa đến phần Quản lý cấu hình, nhưng rõ ràng là Ansible được yêu thích trong lĩnh vực này để quản lý và tự động hóa cấu hình. Ansible được viết bằng Python nhưng bạn không cần phải biết Python để sử dụng nó.

- Agentless
- Chỉ yêu cầu SSH
- Cộng đồng hỗ trợ lớn
- Rất nhiều mô-đun mạng
- Mô hình Push only
- Cấu hình với YAML
- Mã nguồn mở!

[Link to Ansible Network Modules](https://docs.ansible.com/ansible/2.9/modules/list_of_network_modules.html)

Chúng ta cũng sẽ tìm hiểu **Ansible Tower** trong phần quản lý cấu hình, nó được xem như là giao diện người dùng (GUI) cho Ansible.

CI/CD - Một lần nữa, chúng ta sẽ đề cập nhiều hơn về các khái niệm và công cụ xung quanh vấn đề này nhưng điều quan trọng là ít nhất phải đề cập ở đây vì khái niệm này không chỉ xuất hiện trong phần mạng mà còn bao gồm trong tất cả quá trình cung cấp dịch vụ và nền tảng.

Đặc biệt, Jenkins dường như là một công cụ phổ biến cho Tự động hóa mạng.

- Theo dõi kho lưu trữ git để biết các thay đổi và sau đó khởi tạo chúng.

Kiểm soát phiên bản - Một lần nữa chúng ta sẽ tìm hiểu sâu hơn về công nghệ này ở phần sau.

- Git cho phép kiểm soát các phiên bản code của bạn trên máy tính cục bộ - Hỗ trợ đa nền tảng
- GitHub, GitLab, BitBucket, v.v. là các trang web trực tuyến nơi bạn tạo ra các kho lưu trữ và tải code của mình lên.

Ngôn ngữ Lập trình | Scripting - Thứ mà chúng ta chưa đề cập ở đây là Python với tư cách là một ngôn ngữ, tôi quyết định đi sâu vào Go dựa trên hoàn cảnh của tôi. Tôi cho rằng có một cuộc so sánh giữa Golang và Python và Python có vẻ như là người chiến thắng cho ngôn ngữ lập trình để tự động hóa mạng.

- Nornir là thứ cần đề cập ở đây, một framework tự động hóa được viết bằng Python. Nó tương tự như Ansible nhưng cụ thể là xung quanh việc tự động hóa mạng. [Nornir documentation](https://nornir.readthedocs.io/en/latest/)

Phân tích API - Postman là một công cụ tuyệt vời để phân tích API RESTful. Giúp xây dựng, kiểm tra và sửa đổi API.

- POST >>> Để tạo các đối tượng tài nguyên.
- GET >>> Để truy xuất tài nguyên.
- PUT >>> Để tạo hoặc thay thế tài nguyên.
- PATCH >>> Để tạo hoặc cập nhật đối tượng tài nguyên.
- Delete >>> Để xóa tài nguyên

[Postman tool Download](https://www.postman.com/downloads/)

### Các công cụ khác cần đề cập

[Cisco NSO (Network Services Orchestrator)](https://www.cisco.com/c/en/us/products/cloud-systems-management/network-services-orchestrator/index.html)

[NetYCE - Simplify Network Automation](https://netyce.com/)

[Network Test Automation](https://pubhub.devnetcloud.com/media/genie-feature-browser/docs/#/)

Trong 3 ngày tới, tôi sẽ cung cấp nhiều hơn các bài thực hành với một số nội dung chúng ta đã đề cập và thực hiện một số công việc xung quanh Python và Tự động hóa mạng.

Cho đến nay, chúng ta vẫn chưa đề cập đến tất cả các chủ đề của mạng máy tính nhưng tôi cũng muốn làm cho chủ đề này đủ rộng để theo dõi và các bạn có thể tiếp tục học hỏi từ các tài nguyên mà tôi bổ sung bên dưới.

## Tài liệu tham khảo

- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

Hẹn gặp lại các bạn vào [Ngày 25](day25.md)
