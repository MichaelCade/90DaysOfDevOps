---
title: '#90DaysOfDevOps - Vòng đời DevOps - Tập trung vào ứng dụng - Ngày 3'
published: false
description: 90DaysOfDevOps - Vòng đời DevOps - Tập trung vào ứng dụng
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048825
---

## Vòng đời DevOps - Tập trung vào ứng dụng

Trong vài tuần tới, chúng ta sẽ nói đến vòng đời của DevOps (Phát triển liên tục, Kiểm thử, Triển khai, Giám sát) rất nhiều lần. Nếu bạn muốn trở thành một kỹ sư DevOps thì bạn sẽ quen với việc lặp đi lặp lại những điều này. Tuy nhiên, việc liên tục cải thiện từng quá trình trong vòng đời này sẽ khiến công việc của chúng ta trở nên thú vị.

Trong lần này, chúng ta sẽ nhìn tổng thể các quá trình phát triển ứng dụng từ khi bắt đầu tới khi kết thúc và nhắc lại điều này nhiều lần như một vòng lặp liên tục.

## Phát triển

Hãy tưởng tượng chúng ta đang bắt đầu phát triển mội ứng dụng hoàn toàn mới. Nếu bạn là một kỹ sư phát triển, bạn cần thảo luận với khách hàng hoặc người dùng cuối về các yêu cầu của họ rồi đưa ra kế hoạch, những yêu cầu về mặt tính năng, thiết kế cho ứng dụng của bạn. Sau đó, chúng ta tạo ứng dụng mới này từ các yêu cầu đó.

Nói tới các công cụ ở quá trình này, không có yêu cầu thực sự nào khác ngoài việc chọn môi trường phát triển tích hợp (IDE) và ngôn ngữ lập trình mà bạn muốn sử dụng để viết ứng dụng.

Hãy nhớ rằng, là một kỹ sư DevOps, bạn có có thể không phải là người lên kế hoạch cũng như lập trình ứng dụng cho người dùng cuối. Việc này thường sẽ được một kỹ sư phát triển đảm nhiện. Tuy nhiên, sẽ rất tuyệt vời nếu bạn có thể đọc được một số đoạn mã và hiểu được các yêu cầu của ứng dụng, từ đó đưa ra các quyết định tốt nhất cho cơ sở hạ tầng cho hệ thống mới.

Ứng dụng này có thể được viết bằng bất cứ ngôn ngữ nào, nhưng điều quan trọng là mã ứng dụng nên được quản lý bởi một hệ thống kiểm soát phiên bản (version control system). Chúng ta sẽ tìm hiểu chi tiết hơn về mục này ở phần sau (cụ thể là sẽ tập trung vào git).

Dù có 1 hay nhiều thành viên tham gia vào dự án, phương pháp tốt nhất vẫn là sử dụng một kho lưu trữ(code repository) để lưu trữ mã giúp nhiều thanh viên có thể cộng tác khi làm việc với mã chương trình. Kho lưu trữ mã có thể được lưu trữ công khai hoặc riêng tư, bạn cũng có thể sẽ nghe nói đến việc sử dụng GitHub hoặc GitLab làm kho lưu trữ. Chúng ta sẽ đề cập đến những điều này trong phần nói về Git.

## Kiểm thử

Ở quá trình này, chúng ta đã có các yêu cầu và đang phát triển ứng dụng. Vấn đề tiếp theo là chúng ta cần đảm bảo rằng mã ứng dụng cần được kiểm thử ở các môi trường khác nhau mà chúng ta có hoặc cụ thể hơn là với ngôn ngữ lập trình chúng ta đã chọn.

Quá trình này cho phép nhóm quản lý chất lượng (QA) kiểm tra lỗi, chúng ta thường sử dụng các containers để mô phỏng môi trường kiểm thử để có thể giảm thiểu chi phí cho các máy chủ vật lý hoặc cơ sở hạ tầng trên cloud.

Quá trình này có khả năng cao là sẽ được tự động hoá như một phần của quá trình Tích hợp liên tục được nhắc tới sau đây.

Tự động hoá quá trình kiểm thử sẽ giải phóng 10, 100 thậm chí hàng nghìn kỹ sư quản lý chất lượng khỏi việc phải kiểm tra một cách thủ công. Từ đó họ có thể tập trung vào các phần khác trong hệ thống giúp tăng tốc độ phát triển, tạo ra nhiều chức năng mới hơn thay vì sa lầy vào việc kiểm tra lỗi và phần mềm và làm chậm quá trình phát hành các phiên bản mới theo mô hình thác nước truyền thống.

## Tích hợp

Ở giữa vòng đời DevOps là Tích hợp - một phần rất quan trọng. Đây là một phương thức khi mà khi các lập trình viên thay đổi mã nguồn một cách thường xuyên hơn. Điều này có thể diễn ra hàng ngày hoặc hàng tuần.

Với mỗi thay đổi, ứng dụng sẽ được kiểm tra bằng các giai đoạn kiểm thử tự động, điều này sẽ phát hiện sớm các vấn đề hoặc lỗi trước khi đi tới quá trình tiếp theo.

> Nhưng chúng tôi không tạo ra ứng dụng, chúng tôi mua nó từ một nhà cung cấp phần mềm 

Đừng lo, rất nhiều doanh nghiệp làm vậy và sẽ tiếp tục làm điều này. Nhà cung cấp phần mềm sẽ là người thực hiện 3 quá trình ở trên nhưng có thể bạn vẫn muốn áp dụng quá trình cuối cùng vào ứng dụng của bạn. Quá trình này cho phép việc triển khai ứng dụng nhanh chóng và hiệu quản hơn.

Việc hiểu về các quá trình cũng rất quan trọng vì có thể bạn sẽ mua các ứng dụng từ nhà cung cấp ngày hôm nay, nhưng không chắc bạn sẽ làm điều tương tự vào ngày mai, hoặc có thể nó sẽ có ích khi bạn chuyển sang một công việc mới?

## Triển khai

Cuối cùng thì chúng ta cũng đã xây dựng xong ứng dụng, kiểm thử với các yêu cầu của người dùng cuối, bây giờ chúng ta cần triển khai ứng dụng này để người dùng cuối có thể sử dụng.

Đây chính là quá trình mà các đoạn mã sẽ được triển khai lên các máy chủ của môi trường production và cũng chính là một quá trình cực kỳ thú vị và cũng chính là phần mà chúng ta sẽ đào sâu hơn trong 86 ngày còn lại. 

Các ứng dụng khác nhau đòi hỏi các yêu cầu khác nhau về phần cứng và cấu hình. Đó là khi quản lý cấu hình ứng dụng (Application Configuration Management) và cơ sở hạ tầng ứng dụng dưới dạng mã (Infrastructure as Code) đóng vai trò then chốt trong vòng đời DevOps. Các ứng dụng có thể được đóng gói và chạy trong các containers hoặc chạy trên các máy ảo (VM). Điều này khiến chúng ta cần sử dụng các nền tảng như Kubernetes để điều phối các containers và đảm bảo ứng dụng ở trong trạng thái mong muốn nhằm phục vụ người dùng cuối.

Chúng ta sẽ tìm hiểu chi tiết về các chủ đề quan trọng này trong vài tuần tới để có kiến thức nền tảng tốt hơn về chúng và khi nào thì nên sử dụng.

## Giám sát

Mọi thứ diễn ra rất nhanh chóng, chúng ta đã triển khai úng dụng mới và liên tục cập nhật những tính năng mới và luôn kiểm thử trước mỗi lần phát hành để đảm bảo không có bug nào trong ứng dụng. Ứng dụng cũng đang chạy trong môi trường mà cấu hình cũng như hiệu năng ổn định theo đúng yêu cầu.

Nhưng bây giờ chúng ta phải đảm bảo rằng người dùng cuối có được trải nghiệm theo đúng những gì họ mong đợi. Trong quá trình này, chúng ta cần phải đảm bảo rằng hiệu suất của ứng dụng được theo dõi liên tục. Việc này cũng giúp các kỹ sư phát triển có thể đưa ra quyết định tốt hơn để cải tiến ứng dụng trong các bản phát hành tiếp theo nhằm đem lại trải nghiệm tốt hơn cho người dùng cuối.

Giai đoạn này cũng là khi chúng ta nhận các phản hồi về các tính năng và nhu cầu của người dùng cuối với ứng dụng.

Độ tin cậy (reliability) cũng là một yếu tố quan trọng. Suy cho cùng thì chúng ta muốn ứng dụng của mình luôn sẵn sàng bất cứ khi nào người dùng cần. Chính vì vậy, các yếu tố khác như bảo mật, tính quan sát, quản lý dữ liệu cần được giám sát liên tục và kết quả được sử dụng để liên tục cải tiến, cập nhật ứng dụng bằng việc phát hành các phiên bản mới.

Một vài ý kiến đóng góp từ cộng đồng và cụ thể là [@\_ediri](https://twitter.com/_ediri) cũng đề cập rằng các **FinOps** cũng nên là một phần của quá trình liên tục này. Ứng dụng và Dữ liệu đang chạy và được lưu trữ cũng nên được theo dõi liên tục để đảm bảo mọi thay đổi của tài nguyên không gây ra thiệt hại đáng kể về mặt tài chính, đặc biệt với hoá đơn điện toán đám mây của bạn.

> FinOps, viết tắt của Financial Operations, là một phương thức quản lý nhằm nâng cao trách nhiệm chung đối với cơ sở hạ tầng và chi phí điện toán đám mây của một tổ chức.

Đây cũng là thời điểm thích hợp để nói về "Kỹ sư DevOps" được nhắc tới ở trên. Mặc dù có rất nhiều người đang nắm giữ vị trí Kỹ sư DevOps nhưng đó không phải là vị trí duy nhất để áp dụng quy trình DevOps. Quan điểm của tôi khi nói với những người trong cộng đồng là, danh hiệu "kỹ sư DevOps" không nên là mục tiêu cho bất kỳ ai vì bất cứ vị trí nào cũng nên áp dụng các quy trình, văn hoá DevOps được nhắc tới ở đây. DevOps nên được sử dụng ở các vị trí khác ví dụ như kỹ sư/kiến trúc sư Cloud, kỹ sư quản lý ảo hoá, kỹ sư quản lý cơ sở hạ tầng,... Việc sử dụng vị trí "kỹ sư DevOps" ở trên chỉ để làm nổi bật phạm vi mà quy trình được sử dụng bởi bất kỳ vị trí nào ở trên và các vị trí khác.

## Tài liệu tham khảo

Luôn sẵn lòng chào đón thêm các tài nguyên mới tại đây.

Lời khuyên của tôi là hãy xem tất cả những tài liệu bên dưới và hy vọng bạn sẽ kết nối được với những điều được viết ở phía trên.

- [Continuous Development](https://www.youtube.com/watch?v=UnjwVYAN7Ns) Video này tập trung vào quá trình sản xuất nhưng văn hóa tinh gọn có thể được áp dụng một cách chặt chẽ với DevOps.
- [Continuous Testing - IBM YouTube](https://www.youtube.com/watch?v=RYQbmjLgubM)
- [Continuous Integration - IBM YouTube](https://www.youtube.com/watch?v=1er2cjUq1UI)
- [Continuous Monitoring](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [The Remote Flow](https://www.notion.so/The-Remote-Flow-d90982e77a144f4f990c135f115f41c6)
- [FinOps Foundation - What is FinOps](https://www.finops.org/introduction/what-is-finops/)
- [**CÓ PHÍ** The Phoenix Project: A Novel About IT, DevOps, and Helping Your Business Win](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/1942788290/)

Nếu bạn đã đi được tới đây, bạn chắn hẳn đã biết rõ liệu đây có phải là hành trình bạn muốn theo đuổi hay không. Hẹn gặp lại vào [ngày 4](day04.md).
