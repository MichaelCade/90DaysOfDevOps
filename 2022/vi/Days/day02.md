---
title: "#90DaysOfDevOps - Trách nhiệm của kỹ sư DevOps - Ngày 2"
published: false
description: 90DaysOfDevOps - Trách nhiệm của kỹ sư DevOps
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048699
date: "2022-04-17T21:15:34Z"
---

## Trách nhiệm của kỹ sư DevOps

Hy vọng bạn tới với bài viết này sau khi đọc xong [ngày 1 của #90daysdevops](day01.md)

Dù đã đề cập ngắn gọn trong bài viết đầu tiên, chúng ta cần phải đi sâu hơn vào việc có hai phần chính khi tạo một ứng dụng. Đầu tiên là phát triển (Develop), nơi kỹ sư phát triển phầm mềm lập trình ứng dụng và kiểm tra (test) nó. Sau đó là phần vận hành (Operations) nơi ứng dụng được triển khai (deploy) và duy trì (maintain) trên máy chủ.

## DevOps là liên kết giữa hai quá trình

Để hiểu rõ về DevOps và các nhiệm vụ mà một kỹ sư DevOps cần thực hiện, chúng ta cần hiểu về các công cụ, quy trình, tổng quan về chúng và cách chúng kết hợp với nhau.

Tất cả bắt đầu với ứng dụng! Bạn sẽ biết rõ rằng, khi nói tới DevOps, ứng dụng là trung tâm.

Các kỹ sư phát triển tạo ra các ứng dụng bằng nhiều công nghệ khác nhau, chúng ta sẽ đề cập tới điều này sau. Để phát triển ứng dụng, các kỹ sư phát triển cũng có thể sử dụng nhiều ngôn ngữ lập trình, công cụ xây dựng (build tools), kho mã (code repositories), v.v.

Là một kỹ sư DevOps, bạn không lập trình các ứng dụng nhưng việc hiểu rõ về cách làm việc của một kỹ sư phát triển cũng như các hệ thống, công cụ và quy trình mà họ sử dụng là chìa khoá thành công của bạn.

Bạn cần biết ứng dụng được cấu hình như thế nào để tương tác với tất các các dịch vụ hoặc dữ liệu cũng như đưa ra yêu cầu về cách mà chúng ta có thể và nên kiểm tra điều đó.

Ứng dụng cần được triển khai ở đâu đó, để cho đơn giản thì chúng ta sẽ coi đó là một máy chủ. Sau đó, người dùng cuối hoặc khách hàng sẽ truy cập tới ứng dụng thông qua máy chủ được triển khai.

Máy chủ này sẽ chạy ở đâu đó, có thể là on-premise, public cloud hoặc serverless (OK tôi đã đi quá xa, chúng ta sẽ không đề cập tới serverless nhưng đó là lựa chọn mà ngày càng nhiều doanh nghiệp đang sử dụng). Ai đó cần triển khai, cài đặt cấu hình các máy chủ này và chuẩn bị để chúng có thể chạy ứng dụng đã được phát triển. Công việc này có thể sẽ là nhiệm vụ của một kỹ sư DevOps.

Các máy chủ chạy một hệ điều hành và để đơn giản thì chúng ta sẽ coi đó là Linux, vì thế chúng ta sẽ có nguyên một phần hoặc một tuần để bạn có được kiến thức nền tảng về phần này.

Chúng ta cũng nên có kiến thức về mạng và cấu hình mạng vì có thể các máy chủ cần giao tiếp với nhau hoặc với các thành phần khác trong mạng và môi trường của nó. Đôi khi đây cũng sẽ là trách nhiệm của một kỹ sư DevOps. Chúng ta sẽ đi vào chi tiết trong các phần riêng về DNS, DHCP, cân bằng tải, v.v.

## Quý ông/quý bà biết tuốt

Bạn không cần phải là một chuyên gia về mạng hoặc cơ sở hạ tầng nhưng bạn cần có kiến thức nền tảng về cách triển khai các thành phần của hệ thống và giúp chúng có thể kết nối với nhau. Cũng giống như vậy, bạn không cần là một kỹ sư phát triển nhưng cần có kiến thức cơ bản về ngôn ngữ lập trình. Bạn hoàn toàn có thể bắt đầu với tư cách là một chuyên gia trong lĩnh vực nào đó và sử dụng nó như một bước đệm tốt để bạn có thể làm quen với các lĩnh vực còn lại.

Khả năng cao là bạn cũng sẽ không quản lý và tương tác với các máy chủ hoặc ứng dụng hàng ngày.

Chúng ta đã nhắc tới máy chủ khá nhiều nhưng có khả năng là ứng dụng của bạn sẽ được phát triển để chạy trong các containers. Do phần lớn các containers vẫn cần chạy trên các máy chủ nên ngoài kiến thức về chúng, bạn cũng cần hiểu về ảo hoá, dịch vụ cơ sở hạ tầng trên cloud (Cloud IaaS). Trong 90 ngày này, containers sẽ là một phần được chú trọng.

## Tóm tắt tổng quan

Ở một mặt, chúng ta có các kỹ sư phát triển để tạo ra các tính năng và sửa lỗi cho ứng dụng.

Mặt khác, chúng ta có một các môi trường, cơ sở hạ tầng hoặc máy chủ được cấu hình và quản lý để chạy ứng dụng này và giao tiếp với tất cả các dịch vụ mà ứng dụng cần.

Câu hỏi lớn là làm thế nào để chúng ta đưa những tính năng và bản sửa lỗi đó vào ứng dụng của mình và cung cấp cho những người dùng cuối?

Làm cách nào để chúng ta phát hành phiên bản mới của ứng dụng? Đây là một trong những nhiệm vụ chính của kỹ sư DevOps và điều quan trọng ở đây là không chỉ tìm ra cách thực hiện điều này một lần mà chúng ta cần thực hiện việc này liên tục theo một cách tự động, hiệu quả mà vẫn bao gồm quá trình kiểm thử!

Chúng ta sẽ kết thúc ngày thứ 2 tại đây. mong rằng đây là một bài viết hữu ích. Trong vài ngày tới, chúng ta sẽ đi sâu hơn vào một số lĩnh vực khác của DevOps. Sau đó, chúng ta sẽ tìm hiểu kỹ về các công cụ, quy trình cũng như lợi ích của chúng.

## Tài liệu tham khảo

Luôn sẵn lòng chào đón thêm các tài nguyên mới tại đây.

Lời khuyên của tôi là hãy xem tất cả những tài liệu bên dưới và hy vọng bạn sẽ kết nối được với những điều được viết ở phía trên.

- [What is DevOps? - TechWorld with Nana](https://www.youtube.com/watch?v=0yWAtQ6wYNM)
- [What is DevOps? - GitHub YouTube](https://www.youtube.com/watch?v=kBV8gPVZNEE)
- [What is DevOps? - IBM YouTube](https://www.youtube.com/watch?v=UbtB4sMaaNM)
- [What is DevOps? - AWS](https://aws.amazon.com/devops/what-is-devops/)
- [What is DevOps? - Microsoft](https://docs.microsoft.com/en-us/devops/what-is-devops)

Nếu bạn đã đi được tới đây, bạn chắn hẳn đã biết rõ liệu đây có phải là hành trình bạn muốn theo đuổi hay không. Hẹn gặp lại vào ngày 3.
