---
title: '#90DaysOfDevOps - Bức tranh toàn cảnh: Cơ sở hạ tầng dưới dạng mã (IaC) - Ngày 56'
published: false
description: 90DaysOfDevOps - Bức tranh toàn cảnh: Cơ sở hạ tầng dưới dạng mã (IaC)
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048709
---

## Bức tranh toàn cảnh: Cơ sở hạ tầng dưới dạng mã (IaC)

Con người phạm sai lầm và tự động hoá là đường đi đúng đắn!

Bạn đang làm thế nào để xây dựng hế thống của bạn?

Kế hoạch của bạn là gì nếu bạn mất tất cả mọi thứ hôm nay, các máy chủ vật lý, máy ảo, VMs trên điện toán đám mây, Cloud PaaS...?

Bạn sẽ mất bao lâu để thay đổi mọi thứ?

Cơ sở hạ tầng dưới dạng mã (Infrastructure as Code - IaC) cung cấp giải pháp để có thể thực hiện được việc này, đồng thời có thể kiểm thử chúng, chúng ta không nên nhầm lẫn điều này với sao lưu và phục hồi, xét về cơ sở hạ tầng và môi trường, các nền tảng của bạn sẽ được xử lý và đối xử như cách bạn nuôi thú cưng và gia xúc. 

TLDR là chúng ta có thể sử dụng mã để xây dựng lại toàn bộ môi trường của mình.

Nếu chúng ta còn nhớ ngay từ đầu, chúng ta đã nói về DevOps nói chung là một cách phá vỡ các rào cản để triển khai các hệ thống đến production một cách nhanh chóng và an toàn.

Cơ sở hạ tầng dưới dạng mã (IaC) giúp chúng ta triển khai các hạ tầng hệ thống, chúng ta đã nói rất nhiều về quy trình và các công cụ. IaC mang đến cho chúng ta nhiều công cụ để có thể có thể quen thuộc với việc làm điều này trong quy trình triển khai hệ thống.

Chúng ta sẽ tập trung vào Cơ sở hạ tầng dưới dạng mã trong tuần này. Bạn cũng có thể nghe điều này được đề cập với các khái niệm như Cơ sở hạ tầng từ mã hoặc cấu hình dưới dạng mã. Tôi nghĩ thuật ngữ phổ biến nhất vẫn sẽ là "Cơ sở hạ tầng dưới dạng mã - IaC"
### Tính thân thiện

Nếu chúng ta xem xét về thời trước khi có DevOps, nếu chúng ta có yêu cầu xây dựng ứng dụng mới, chúng ta sẽ chuẩn bị phần lớn các máy chủ của mình theo cách thủ công, việc này cũng giống như bạn chăm sóc cho thú cưng của mình.

- Triển khai máy ảo | Trước khi có công nghệ ảo hóa thì sẽ là các máy chủ vật lý và cài đặt hệ điều hành
- Cài đặt, cấu hình mạng
- Cài đặt, cấu hình bảng định tuyến
- Cài đặt các phần mềm, thư viện cần thiết và cập nhật chúng
- Cấu hình phần mềm
- Cài đặt cơ sở dữ liệu

Trước đây, tất cả các tác vụ này đều được thực hiện thủ công bởi chuyên viên quản trị hệ thống SysAdmin. Đối với ứng dụng có quy mô lớn hơn, cần nhiều tài nguyên và máy chủ hơn, đương nhiên sẽ cần nhiều công sức hơn để cài đặt chúng. Do vậy chúng sẽ lấy đi rất nhiều công sức của nhân lực lao động (là chính chúng ta) cũng như thời gian, điều mà các doanh nghiệp sẽ phải chi trả trong toàn bộ quá trình xây dựng môi trường công nghệ thông tin ứng dụng này. Đặc biệt là ngay từ đầu nội dung bài viết tôi đã chia sẻ, việc gặp rủi ro từ cá nhân đều có thể xảy ra khiến chi phí để bù lấp sai lầm này có thể sẽ rất lơn, vì vậy tự động hóa sẽ là sự lựa chọn hàng đầu.

Quay trở lại nội dung chính, tiếp theo sau khi các bước cài đặt ban đầu hoàn thiện, chúng ta vẫn còn cần tiến hành bảo trì, quản lý và vận hành các hệ thống máy chủ nói trên, bao gồm:

- Cập nhật các phiên bản mới
- Triển khai các phiên bản mới này
- Quản lý dữ liệu
- Khôi phục ứng dụng nếu có sự cố
- Loại bỏ, thêm bớt và mở rộng các máy chủ, tài nguyên phần cứng trong trường hợp cần thiết
- Cấu hình mạng

Các bạn hãy tưởng tượng từng đó công việc sẽ được lặp đi lặp lại cho các môi trường phát triển như: dev, test, production ... độ phức tạp vì vậy cũng tăng lên.

Đây là lúc cơ sở hạ tầng dưới dạng mã xuất hiện, ở trên đã có nhiều việc chúng ta chăm sóc những máy chủ này như những thú cưng, mọi người thậm chí còn còn gọi các máy chủ của mình với tên gọi như thú cưng hoặc ít nhất là đặt tên cho chúng vì chúng sẽ ở đó một thời gian và họ hi vọng chúng sẽ trở thành một phần của "gia đình" trong một thời gian.

Với cơ sở hạ tầng dưới dạng mã IaC, chúng ta có thể tự động hoá tất các tác vụ này từ đầu đến cuối.  Cơ sở hạ tầng dưới dạng mã là một khái niệm và có một số cung cụ thực hiện việc cung cấp cơ sở hạ tầng một cách tự động, và nếu có điều gì đó xảy ra với máy chủ, bạn sẽ vứt bỏ nó và khởi động lại một cái mới. Quá trình này được tự động hoá và máy chủ sẽ được tại chính xác như những gì được cấu hình trong mã. Tại thời điểm này, chúng ta không quan tâm chúng được gọi là gì, chúng ở đó phục vụ mục đích của chúng ta cho đến khi không còn ở đó nữa và chúng ta sẽ có cái khác để thay thế, có thể do bị lỗi hoặc vì chúng ta đã cập nhật một phần hoặc toàn bộ ứng dụng của mình.

IaC được sử dụng ở hầu hết mọi nền tảng, công nghệ ảo hóa, công nghệ điện toán đám mấy và công nghệ ứng dụng khai thác lợi ích của điện toán đám mây (Cloud Native) như Kubernetes hay containers.

## Khởi tạo cơ sở hạ tầng

Không phải các công cụ IaC sẽ bao gồm tất cả những điều bên dưới, bạn sẽ thấy rằng công cụ chúng ta tìm hiểu ở phần này chỉ thực sự bao gồm 2 lĩnh vực đầu tiên trong danh sách dưới đây. Terraform là công cụ mà chúng ta sẽ cho phép chúng ta bắt đầu từ số không, định nghĩa trong code xem cơ sở hạ tầng của chúng ta sẽ trông như thế nào và sau đó triển khai nó, nó cũng cho phép chúng ta quản lý cơ sở hạ tầng đó và ban đầu cũng có thể triển khai một ứng dụng nhưng sau đó, nó sẽ không quản lý ứng dụng nữa, đây là nơi mà phần tiếp theo sẽ xuất hiện và một công cụ quản lý cấu hình như Ansible sẽ đáp ứng được nhu cầu này một cách tốt hơn.

Các công cụ như `chef`, `puppet` và `ansible` ngay từ đầu là những công cụ vô cùng phù hợp để giải quyết yêu cầu khởi tạo, cài đặt ứng dụng và quản lý chúng.

Vậy các công việc khởi tạo, cài đặt cấu hình một phần mềm sẽ gồm những gì?

- Khởi tạo máy chủ mới 
- Cấu hình mạng máy tính
- Cấu hình bộ cân bằng tải ứng dụng
- Cấu hình ở mức hạ tầng

### Cấu hình cơ sở hạ tầng

- Cài đặt các ứng dụng, phần mềm theo yêu cầu lên máy chủ (các phần mềm để chạy được ứng dụng ví dụ: python, go...)
- Chuẩn bị hàng loạt máy chủ để triển khai ứng dụng (lặp lại các bước trên trên rất nhiều máy chủ)

### Triển khai phần mềm

Sau khi hạ tầng máy chủ đã sẵn sàng chúng ta sẽ triển khai ứng dụng lên các máy chủ này, việc này bao gồm các bước sau:

- Triển khai, quản lý ứng dụng (bao gồm ứng dụng, thư viện hỗ trợ)
- Bảo trì
- Cập nhật phần mềm (cũng có thể là các thư viện phụ thuộc)
- Định lại cấu hình trong trường hợp cần thiết

### Sự khác nhau giữa các công cụ IaC

Khai báo và trình tự thực hiện (Declarative vs procedural)

Trình tự thực hiện

- Các bước trong IaC được làm tuần tự, theo từng bước
- Khởi tạo máy chủ, thêm máy chủ vào hệ thống và thay đổi cấu hình

Khai báo

- Khai báo kết quả mong muốn (VD: tạo 1 hoặc nhiều máy chủ cùng lúc)
- Ví dụ: Khởi tạo 2 máy chủ, hoặc 2 buckets

Tính bất biến (gia súc) và tính khả biến (thú cưng)

Khả biến
- Có thể thay đổi cấu hình thay vì ghi đè hoặc thay thế (ví dụ: thay đổi tên của máy chủ Windows, đổi tagging của s3 bucket)
- Vì có tính khả biến nên vòng đời sẽ lâu hơn.

Bất biến
- Khi muốn thay đổi chúng ta thay thế mới
- Vòng đời ngắn hơn.

Mỗi tài nguyên trong hệ thống đều có thể có một hoặc nhiều tính khả biến và bất biến.

Lấy ví dụ như sau:
- 01 AWS S3 bucket khi đã đặt tên (bucket name) sẽ bắt buộc là duy nhất và không thể thay đổi được tuy nhiên ta hoàn toàn có thể thay đổi nhãn (tag) của chúng mà không cần phải tạo mới bucket
- Ví dụ thứ 2, với một container image, chúng ta sẽ cần nó là bất biến, nghĩa là khi muốn cập nhật mã nguồn chúng ta bắt buộc phải tạo container image mới

Với các ví dụ trên, có rất nhiều lựa chọn cho IaC, tuy nhiên không có công cụ IaC nào có thể định nghĩa cũng như giải quyết được hết các tính chất này, thay vào đó chúng ta phải hiểu được tính chất của từng tài nguyên (resource, infra - hạ tầng)

Cũng trong series này, chúng ta sẽ bắt đầu thực hành với Terraform, công cụ được coi là thích hợp nhất ở thời điểm hiện tại để giúp chúng ta thấy được lợi ích mà IaC mang lại. Việc thực hành cũng là cách tốt nhất để nâng cao khả năng, kỹ năng lập trình.

Tiếp theo, bắt đầu với lý thuyết về Terraform ở mức độ cơ bản (101) sau đó chúng ta sẽ bắt đầu thực hành

## Tài liệu Tham khảo

Tôi đã liệt kê ra rất nhiều nội dung bên dưới đây để chúng ta có thể bắt đầu học và tham khảo, với các nội dung này phần nào có thể giúp chúng ta nhanh chóng nắm được khái niệm và các lý thuyết xung quanh IaC

- [What is Infrastructure as Code? Difference of Infrastructure as Code Tools](https://www.youtube.com/watch?v=POPP2WTJ8es)
- [Terraform Tutorial | Terraform Course Overview 2021](https://www.youtube.com/watch?v=m3cKkYXl-8o)
- [Terraform explained in 15 mins | Terraform Tutorial for Beginners](https://www.youtube.com/watch?v=l5k1ai_GBDE)
- [Terraform Course - From BEGINNER to PRO!](https://www.youtube.com/watch?v=7xngnjfIlK4&list=WL&index=141&t=16s)
- [HashiCorp Terraform Associate Certification Course](https://www.youtube.com/watch?v=V4waklkBC38&list=WL&index=55&t=111s)
- [Terraform Full Course for Beginners](https://www.youtube.com/watch?v=EJ3N-hhiWv0&list=WL&index=39&t=27s)
- [KodeKloud - Terraform for DevOps Beginners + Labs: Complete Step by Step Guide!](https://www.youtube.com/watch?v=YcJ9IeukJL8&list=WL&index=16&t=11s)
- [Terraform Simple Projects](https://terraform.joshuajebaraj.com/)
- [Terraform Tutorial - The Best Project Ideas](https://www.youtube.com/watch?v=oA-pPa0vfks)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)

Hẹn gặp lại các bạn ở ngày [ngày 57](day57.md)

