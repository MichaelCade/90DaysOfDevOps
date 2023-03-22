---
title: '#90DaysOfDevOps - Giới thiệu - Ngày 1'
published: true
description: 90DaysOfDevOps - Giới thiệu
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048731
date: '2022-04-17T10:12:40Z'
---

## Giới thiệu - Ngày 1

Ngày 1 - ngày đầu tiên của 90 ngày và cũng là ngày đầu tiên trong hành trình học những kiến thức nền tảng và các công cụ giúp bạn hình thành tư duy DevOps của mình.

Hành trình này bắt đầu với tôi từ vài năm trước, khi đó tôi chỉ tập trung vào các nền tảng ảo hóa và công nghệ liên quan tới cloud, chủ yếu làm việc với Infrastructure as Code (IaC) và quản lý cấu hình sử dụng Terraform và Chef.

Cho tới tháng 3/2021, tôi có cơ hội tuyệt vời khi được phát triển một chiến lược Cloud Native tại Kasten by Veeam và từ đó tập trung vào Kubernetes, DevOps và cộng đồng của các công nghệ này. Tôi bắt đầu hành trình DevOps của mình và nhanh chóng nhận ra rằng có rất nhiều thứ tuyệt vời để học ngoài kiến thức cơ bản về Kubernetes và Containerisation, đó cũng là lúc tôi bắt đầu trao đổi với cộng đồng, tìm hiểu sâu hơn về văn hoá, các công cụ và quy trình DevOps, đồng thời khi chép lại những điều mà tôi đã học và chia sẻ nó.

[So you want to learn DevOps?](https://blog.kasten.io/devops-learning-curve)

## Bắt đầu hành trình ngay bây giờ

Sau khi bạn đọc bài viết ở trên, bạn có thể thấy đó là một tóm tắt "high-level" cho hành trình học tập của tôi. Tại thời điểm này, dù cho tôi không phải là một chuyên gia trong bất cứ lĩnh vực nào nhưng tôi muốn chia sẻ một số tài nguyên gồm có miễn phí và trả phí để các bạn có thể lựa chọn phù hợp với điều kiện của bản thân.

Trong 90 ngày tới, tôi sẽ cố gắng ghi lại danh sách các tài nguyên này và đề cập tới những kiến thức nền tảng. Tôi cũng muốn nhận được sự giúp đỡ, ủng hộ của cộng đồng bằng cách chia sẻ hành trình của riêng bạn cũng như các tài nguyên bạn đã sử dụng để chúng ta có thể học hỏi, giúp đỡ lẫn nhau.

Như bạn có thể thấy trong file README ở thư mục gốc, hành trình sẽ được chia thành 12 tuần + 6 ngày. Trong 6 ngày đầu tiên, chúng ta sẽ tìm hiểu, khám phá các nguyên tắc cơ bản của DevOps trước khi đi sâu vào từng phần cụ thể. Danh sách này khó có thể chứa đầy đủ tất cả các thông tin và một lần nữa, tôi rất mong cộng đồng hỗ trợ trong việc xây dựng một bộ tài liệu hữu ích.

Dưới đây là một road map về DevOps mà tôi nghĩ rằng mọi người nên xem kỹ từ đó có thể tạo bản đồ tư duy sao cho phù hợp với sở thích và mục đích của riêng mình:

[DevOps Roadmap](https://roadmap.sh/devops)

Đây là một tài liệu tuyệt vời giúp tôi trong việc bắt đầu tạo danh sách này và chuẩn bị cho bài viết đầu tiên. Bạn cũng có thể thấy các chủ đề khác được đề cập chi tiết hơn 12 chủ đề được liệt kê trong danh sách của tôi.

## Bước chân đầu tiên - DevOps là gì?

Trước khi bắt đầu dành khoảng 1 tiếng mỗi ngày trong hành trình 90 ngày để học những kiến thức liên quan tới DevOps, chúng ta nên bắt đầu với câu hỏi "DevOps là gì?"

Đầu tiên, DevOps không phải một công cụ. Bạn không thể mua, nó cũng không phải một phần mềm hay một dự án mã nguồn mở trên Github chúng ta có thể download. Nó cũng không phải là một ngôn ngữ lập trình và càng không phải là một phép thuật hắc ám. 

DevOps là một cách để phát triển phần mềm một cách thông minh hơn. - Khoan đã, nếu bạn không phải một nhà phát triển phần mềm, bạn có nên dừng hành trình của mình ngay bây giờ và không tiếp tục đọc tài liệu này??? Không hề. Hãy ở lại với anh em... Lý do vì DevOps là sự kết hợp giữa việc phát triển (Dev) và vận hành (Ops). Như bạn có thể thấy, tôi chủ yếu làm việc với các VM và ở gần hơn với phần vận hành (Operations). Tuy nhiên trong cộng đồng, dù cho các bạn có các nền tảng khác nhau, làm việc ở các vị trí khác nhau, DevOps chắc chắn vẫn sẽ giúp bạn trong công việc hàng ngày. Dù là kỹ sư phát triển phần mềm, kỹ sư vận hành hệ thống tới các kỹ sư quản lý chất lượng đều có thể học và áp dụng các phương pháp tốt nhất trong công việc của mình bằng cách hiểu rõ hơn về DevOps.

DevOps là một tập hợp các phương pháp thực hành giúp chúng ta đạt được mục tiêu sau: Giảm thời gian từ khi lên ý tưởng cho một sản phẩm đến khi phát hành được sản phẩm nhằm phục vụ người dùng cuối, bất kể đó là một khách hàng bên ngoài hay là một người dùng nội bộ.

Một chủ đề khác mà chúng ta sẽ đề cập trong tuần đầu tiên là **Phương pháp Agile**

DevOps và Agile được áp dụng rộng rãi cùng với nhau giúp ứng dụng của bạn được phân phối liên tục.

Điểm mấu chốt của tư duy, văn hoá DevOps là thu hẹp, chia nhỏ quy trình phát triển phần mềm có thể kéo dài nhiều năm thành các bản phát hành nhỏ hơn và thường xuyên hơn. Một điểm quan trọng khác cần hiểu ở đây là trách nhiệm của một kỹ sư DevOps trong việc thu hẹp khoảng cách giữa các nhóm đã được đề cập: nhóm phát triển, nhóm vận hành, nhóm quản lý chất lượng.

Từ góc độ DevOps, **Phát triển, Kiểm thử và Triển khai** đều liên quan tới nhóm DevOps.

Cuối cùng, **Tự động hoá** cũng phải được triển khai để các công việc đạt hiệu quả cao nhất.

## Tài liệu tham khảo

Vì đây là một danh sách các tài liệu học tập nên rất mong các bạn có thể đóng góp, bổ sung thêm tư liệu hữu ích vào các file README.

Bạn nên xem tất cả những videos bên dưới và hi vọng bạn có thể liên kết chúng với những giải thích ở phía trên.

- [DevOps in 5 Minutes](https://www.youtube.com/watch?v=Xrgk023l4lI)
- [What is DevOps? Easy Way](https://www.youtube.com/watch?v=_Gpe1Zn-1fE&t=43s)
- [DevOps roadmap 2022 | Success Roadmap 2022](https://www.youtube.com/watch?v=7l_n97Mt0ko)

Nếu bạn đã đi được tới đây, bạn chắn hẳn đã biết rõ liệu đây có phải là hành trình bạn muốn theo đuổi hay không. Hẹn gặp lại vào [Ngày 2](day02.md).
