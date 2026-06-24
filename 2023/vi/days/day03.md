## Suy nghĩ như một kẻ tấn công

Hôm qua chúng ta đã tìm hiểu về DevSecOps, trong bài viết này chúng ta sẽ xem xét một số đặc điểm của kẻ tấn công. Khi suy nghĩ về kẻ tấn công, tốt nhất chúng ta nên suy nghĩ như một kẻ tấn công.

### Đặc điểm của kẻ tấn công

Trước hết, tất cả các doanh nghiệp và phần mềm đều là mục tiêu tấn công của kẻ tấn công, không có nơi nào an toàn, chúng ta chỉ có thể làm cho nơi đó an toàn hơn và ít hấp dẫn hơn để bị tấn công.

![](../../images/day03-2.jpg)
***[hình ảnh từ nguồn này](https://www.trainerize.me/articles/outrun-bear/)***

Khi suy nghĩ về điều đó, kẻ tấn công luôn là mối đe dọa liên tục!

Kẻ tấn công sẽ xác định các lỗ hổng bảo mật bằng cách thực hiện các cuộc tấn công theo một thứ tự cụ thể để truy cập, lấy dữ liệu và thành công trong nhiệm vụ của họ.

Kẻ tấn công có thể may mắn, nhưng họ chắc chắn sẽ làm việc trên các cuộc tấn công có mục tiêu.

Các cuộc tấn công có thể chậm và kiên trì hoặc nhanh chóng để đạt được sự xâm nhập. Không phải tất cả các cuộc tấn công đều giống nhau.

### Động cơ của kẻ tấn công

Là một nhóm DevOps, bạn sẽ cung cấp cơ sở hạ tầng, phần mềm và bảo vệ các môi trường này có thể trải dài trên nhiều đám mây, ảo hóa và container hóa trên các nền tảng.

Chúng ta phải xem xét các điều sau:

- **Làm thế nào** họ sẽ tấn công chúng ta?
- **Tại sao** họ sẽ tấn công chúng ta?
- **Chúng ta** có gì có giá trị đối với kẻ tấn công?

Động cơ của kẻ tấn công cũng sẽ khác nhau tùy thuộc vào kẻ tấn công. Ý tôi là nó có thể chỉ là để cho vui... Chúng ta có lẽ tất cả đã từng như vậy, trong trường học và chỉ đi quá sâu vào những thứ trên mạng để tìm thêm thông tin. Ai cũng sẽ có câu chuyện riêng để kể.

Nhưng như chúng ta đã thấy trên các phương tiện truyền thông, các cuộc tấn công thường liên quan đến tiền bạc, gian lận hoặc thậm chí là các cuộc tấn công chính trị vào các doanh nghiệp và tổ chức.

Trong không gian Kubernetes, chúng ta thậm chí đã thấy kẻ tấn công tận dụng và sử dụng sức mạnh tính toán của một môi trường để khai thác tiền điện tử.

Trọng tâm của cuộc tấn công này có thể sẽ là **DỮ LIỆU**

Dữ liệu của một công ty có thể sẽ cực kỳ có giá trị đối với công ty nhưng cũng có thể tiềm ẩn nguy cơ bị lộ ra ngoài. Đó là lý do tại sao chúng ta đặt rất nhiều sự nhấn mạnh vào việc bảo vệ dữ liệu này, đảm bảo rằng dữ liệu được bảo mật và mã hóa.

### Bản đồ tấn công

Bây giờ chúng ta đã có động cơ và một số đặc điểm của kẻ tấn công hoặc một nhóm kẻ tấn công, nếu đây là một cuộc tấn công có kế hoạch thì bạn sẽ cần một kế hoạch, bạn cần xác định những dịch vụ và dữ liệu nào bạn đang nhắm mục tiêu.

Một **bản đồ tấn công** là một biểu diễn trực quan của một cuộc tấn công vào mạng máy tính. Nó hiển thị các giai đoạn khác nhau của cuộc tấn công, các công cụ và kỹ thuật được sử dụng bởi kẻ tấn công, và các điểm vào và ra khỏi mạng. Bản đồ tấn công có thể được sử dụng để phân tích chi tiết các cuộc tấn công trước đây, xác định các lỗ hổng trong mạng, và lập kế hoạch phòng thủ chống lại các cuộc tấn công trong tương lai. Chúng cũng có thể được sử dụng để truyền đạt thông tin về một cuộc tấn công cho các bên liên quan không chuyên về kỹ thuật, chẳng hạn như các giám đốc điều hành hoặc đội ngũ pháp lý.

Bạn có thể thấy từ mô tả trên rằng một bản đồ tấn công nên được tạo ra từ cả hai phía hoặc cả hai đội (về phía đội ngũ, đây là điều mà tôi sẽ đề cập trong một bài viết sau).

Nếu bạn tạo một bản đồ tấn công mạng gia đình hoặc doanh nghiệp của bạn, một số điều bạn muốn hiểu rõ sẽ là:

- Biểu diễn đồ họa của ứng dụng của bạn bao gồm tất cả các luồng giao tiếp và công nghệ đang được sử dụng.

- Một danh sách các lỗ hổng tiềm năng và các khu vực có thể bị tấn công.

- Xem xét tính bảo mật, tính toàn vẹn và tính khả dụng cho mỗi kết nối/tương tác trong ứng dụng.

- Lập bản đồ các cuộc tấn công/lỗ hổng.

Một bản đồ tấn công có thể trông giống như thế này với giải thích các số đại diện cho điều gì.

![](../../images/day03-1.png)

Từ bản đồ này, chúng ta có thể xem xét có thể có một cuộc tấn công từ chối dịch vụ hoặc một cuộc tấn công từ nội bộ và truy cập vào S3 bucket để ngăn ứng dụng lưu trữ dữ liệu hoặc gây ra việc lưu trữ dữ liệu xấu.

Bản đồ này không bao giờ được hoàn thiện, giống như ứng dụng của bạn liên tục tiến lên thông qua phản hồi, bản đồ tấn công này cũng cần được kiểm tra, điều này cung cấp phản hồi và từ đó tăng cường tư thế bảo mật chống lại các cuộc tấn công này. Bạn có thể gọi đây là "Phản hồi Liên tục" trong vòng phản hồi bảo mật.

Ở mức tối thiểu, chúng ta nên tuân theo mô hình tốt, tốt hơn, tốt nhất để cải thiện tư thế bảo mật.

- **Tốt** - Xác định các ràng buộc thiết kế bảo mật và các biện pháp kiểm soát cần được xây dựng vào phần mềm để giảm thiểu một cuộc tấn công.

- **Tốt hơn** - Ưu tiên và xây dựng bảo mật cho các vấn đề được phát hiện sau này trong chu kỳ phần mềm.

- **Tốt nhất** - Tích hợp tự động hóa vào triển khai kịch bản để phát hiện các vấn đề, kiểm tra đơn vị, kiểm tra bảo mật, kiểm tra hộp đen.

Bảo mật là một ràng buộc thiết kế (design constraint) - mặc dù là một ràng buộc bất tiện.

## Tài liệu tham khảo 

- [devsecops.org](https://www.devsecops.org/)

- [TechWorld with Nana - What is DevSecOps? DevSecOps explained in 8 Mins](https://www.youtube.com/watch?v=nrhxNNH5lt0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=1&t=19s)

- [What is DevSecOps?](https://www.youtube.com/watch?v=J73MELGF6u0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=2&t=1s)

- [freeCodeCamp.org - Web App Vulnerabilities - DevSecOps Course for Beginners](https://www.youtube.com/watch?v=F5KJVuii0Yw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=3&t=67s)

- [The Importance of DevSecOps and 5 Steps to Doing it Properly (DevSecOps EXPLAINED)](https://www.youtube.com/watch?v=KaoPQLyWq_g&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=4&t=13s)

- [Continuous Delivery - What is DevSecOps?](https://www.youtube.com/watch?v=NdvMUcWNlFw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=5&t=6s)

- [Cloud Advocate - What is DevSecOps?](https://www.youtube.com/watch?v=a2y4Oj5wrZg&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=6)

- [Cloud Advocate - DevSecOps Pipeline CI Process - Real world example!](https://www.youtube.com/watch?v=ipe08lFQZU8&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=7&t=204s)

Hẹn gặp lại vào [ngày 4](day04.md) 
