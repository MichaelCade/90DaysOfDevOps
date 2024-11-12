## Bức Tranh Toàn Cảnh: DevSecOps

Chào mừng đến với Ngày 2 của phiên bản 2023, trong module đầu tiên của 6 ngày tới, chúng ta sẽ xem xét tổng quan cơ bản về DevSecOps.

### DevSecOps là gì?

DevSecOps là một phương pháp phát triển phần mềm nhằm kết hợp các nhóm phát triển, bảo mật và vận hành để xây dựng và duy trì các ứng dụng phần mềm an toàn. Nó dựa trên các nguyên tắc của tích hợp liên tục, phân phối liên tục và triển khai liên tục, nhằm cung cấp các bản cập nhật và tính năng phần mềm nhanh chóng và thường xuyên hơn. Trong DevSecOps, bảo mật là một phần không thể thiếu của quy trình phát triển phần mềm, thay vì được xem xét cuối cùng. Điều này có nghĩa là kiểm tra bảo mật, giám sát và các biện pháp bảo mật khác được tích hợp vào vòng đời phát triển phần mềm (SDLC) từ đầu và trong mọi công đoạn thay vì được thêm vào sau. DevSecOps nhằm cải thiện sự hợp tác và giao tiếp giữa các nhóm phát triển, bảo mật và vận hành, để tạo ra một quy trình phát triển phần mềm hiệu quả và hiệu quả hơn.

### DevSecOps vs DevOps

Tôi sử dụng "vs" một cách nhẹ nhàng ở đây, nhưng nếu chúng ta nhìn lại phiên bản 2022 và mục tiêu của DevOps là cải thiện tốc độ, độ tin cậy và chất lượng của các bản phát hành phần mềm.

DevSecOps là một phần mở rộng của triết lý DevOps nhấn mạnh việc tích hợp các thực hành bảo mật vào quy trình phát triển phần mềm. Mục tiêu của DevSecOps là xây dựng các biện pháp bảo mật vào quy trình phát triển phần mềm để bảo mật là một phần không thể thiếu của phần mềm từ đầu, thay vì được xem xét cuối cùng. Điều này giúp giảm nguy cơ các lỗ hổng bảo mật trong phần mềm và làm cho việc xác định và sửa chữa bất kỳ vấn đề nào dễ dàng hơn.

DevOps tập trung vào việc cải thiện sự hợp tác và giao tiếp giữa các nhà phát triển và nhân viên vận hành để cải thiện tốc độ, độ tin cậy và chất lượng của các bản phát hành phần mềm, trong khi DevSecOps tập trung vào việc tích hợp các thực hành bảo mật vào quy trình phát triển phần mềm để giảm nguy cơ các lỗ hổng bảo mật và cải thiện tổng thể bảo mật của phần mềm.

### Bảo mật tự động

Bảo mật tự động đề cập đến việc sử dụng công nghệ để thực hiện các nhiệm vụ bảo mật mà không cần sự can thiệp của con người. Điều này có thể bao gồm các phần mềm bảo mật giám sát mạng để phát hiện các mối đe dọa và hành động để chặn chúng, hoặc các hệ thống sử dụng trí tuệ nhân tạo để phân tích cảnh quay bảo mật và xác định hoạt động bất thường. Các hệ thống bảo mật tự động được thiết kế để làm cho các quy trình bảo mật hiệu quả và hiệu quả hơn, và giúp giảm tải công việc cho nhân viên bảo mật.

Một thành phần chính của tất cả các thứ DevSecOps là khả năng tự động hóa nhiều nhiệm vụ khi phát triển và phân phối phần mềm, khi chúng ta thêm bảo mật từ đầu có nghĩa là chúng ta cũng cần xem xét khía cạnh tự động hóa của bảo mật.

### Bảo mật ở Quy mô Lớn (Containers và Microservices)

Chúng ta biết rằng quy mô và cơ sở hạ tầng động đã trở nên phổ biển nhờ container và microservices, chúng cũng thay đổi cách mà hầu hết các tổ chức hoạt động.

Đây cũng là lý do tại sao chúng ta phải đưa bảo mật tự động vào các nguyên tắc DevOps của mình để đảm bảo rằng các hướng dẫn bảo mật container cụ thể được đáp ứng.

Ý tôi là với các công nghệ cloud native, chúng ta không thể chỉ có các chính sách và tư thế bảo mật (security posture) tĩnh; mô hình bảo mật của chúng ta cũng phải động với khối lượng công việc trong tay và cách nó đang chạy.

Các nhóm DevOps sẽ cần bao gồm bảo mật tự động để bảo vệ môi trường và dữ liệu tổng thể, cũng như các quy trình tích hợp liên tục và phân phối liên tục.

Danh sách dưới đây được lấy từ một [bài viết blog của RedHat](https://www.redhat.com/en/topics/devops/what-is-devsecops)

- Chuẩn hóa và tự động hóa môi trường: Mỗi dịch vụ nên có quyền tối thiểu để giảm thiểu các kết nối và truy cập trái phép.

- Tập trung hóa khả năng nhận dạng người dùng và kiểm soát truy cập: Kiểm soát truy cập chặt chẽ và các cơ chế xác thực tập trung là cần thiết để bảo mật microservices vì xác thực được khởi tạo tại nhiều điểm.

- Cách ly (isolate) các container chạy microservices với nhau và mạng: Điều này bao gồm cả dữ liệu đang truyền và dữ liệu lưu trữ vì cả hai đều có thể là mục tiêu có giá trị cao cho các kẻ tấn công.

- Mã hóa dữ liệu giữa các ứng dụng và dịch vụ: Một nền tảng điều phối container với các tính năng bảo mật tích hợp giúp giảm thiểu khả năng truy cập trái phép.

- Giới thiệu các cổng API bảo mật: Các API bảo mật tăng cường khả năng ủy quyền và định tuyến. Bằng cách giảm thiểu các API bị lộ (expose), các tổ chức có thể giảm bề mặt tấn công (attack surface).

### Bảo mật đang là chủ đề HOT hiện nay

Một điều bạn sẽ thấy bất kể nền tảng của bạn là gì là bảo mật đang là keyword rất hot khắp ngành công nghiệp, điều này một phần là do các vi phạm bảo mật xuất hiện trên tin tức toàn cầu và các thương hiệu lớn bị ảnh hưởng bởi các lỗ hổng bảo mật hoặc theo các thực hành xấu tiềm năng cho phép các kẻ xấu vào mạng của các công ty này. Có thể nói hoặc ít nhất từ quan điểm của tôi, việc tạo phần mềm hiện nay dễ dàng và có thể đạt được hơn bao giờ hết. Nhưng trong việc tạo phần mềm, nó ngày càng bị phơi bày với các lỗ hổng và tương tự, điều này cho phép các kẻ xấu gây ra sự hỗn loạn và đôi khi giữ dữ liệu để đòi tiền chuộc hoặc đóng cửa các doanh nghiệp, gây ra sự hỗn loạn. Chúng ta đã thảo luận về DevSecOps là gì nhưng tôi nghĩ cũng đáng để khám phá khía cạnh an ninh mạng của vector tấn công (attack vector) và lý do tại sao chúng ta bảo vệ chuỗi cung ứng phần mềm của mình để giúp tránh các cuộc tấn công mạng này.

### An ninh mạng vs DevSecOps

Như tiêu đề, nó không thực sự là một cuộc đối đầu mà là sự khác biệt giữa hai chủ đề. Nhưng tôi nghĩ điều quan trọng là phải nêu ra điều này vì thực sự điều này sẽ giải thích tại sao bảo mật phải là một phần của quy trình, nguyên tắc và phương pháp DevOps.

An ninh mạng là thực hành bảo vệ các hệ thống máy tính và mạng khỏi các cuộc tấn công kỹ thuật số, trộm cắp và thiệt hại. Nó bao gồm việc xác định và giải quyết các lỗ hổng, triển khai các biện pháp bảo mật và giám sát các hệ thống để phát hiện các mối đe dọa.

DevSecOps, mặt khác, là sự kết hợp của các thực hành phát triển, bảo mật và vận hành. Đó là một triết lý nhằm tích hợp bảo mật vào quy trình phát triển, thay vì coi nó là một bước riêng biệt. Điều này bao gồm sự hợp tác giữa các nhóm phát triển, bảo mật và vận hành trong suốt vòng đời phát triển phần mềm (SDLC).

Một số điểm khác biệt chính giữa an ninh mạng và DevSecOps bao gồm:

**Tập trung**: An ninh mạng chủ yếu tập trung vào việc bảo vệ các hệ thống khỏi các mối đe dọa bên ngoài, trong khi DevSecOps tập trung vào việc tích hợp bảo mật vào quy trình phát triển.

**Phạm vi**: An ninh mạng bao gồm một loạt các chủ đề, bao gồm bảo mật mạng, bảo mật dữ liệu, bảo mật ứng dụng và nhiều hơn nữa. DevSecOps, mặt khác, tập trung cụ thể vào việc cải thiện bảo mật của phát triển và triển khai phần mềm.

**Cách tiếp cận**: An ninh mạng thường bao gồm việc triển khai các biện pháp bảo mật sau khi quy trình phát triển hoàn tất, trong khi DevSecOps bao gồm việc tích hợp bảo mật vào quy trình phát triển từ đầu.

**Hợp tác**: An ninh mạng thường bao gồm sự hợp tác giữa các nhóm IT và bảo mật, trong khi DevSecOps bao gồm sự hợp tác giữa các nhóm phát triển, bảo mật và vận hành.

## Tài liệu tham khảo

Trong suốt 90 ngày, chúng ta sẽ có một danh sách tài nguyên hàng ngày sẽ mang lại nội dung liên quan giúp tiếp tục các chủ đề và nơi bạn có thể tìm hiểu thêm.

- [TechWorld with Nana - What is DevSecOps? DevSecOps explained in 8 Mins](https://www.youtube.com/watch?v=nrhxNNH5lt0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=1&t=19s)

- [What is DevSecOps?](https://www.youtube.com/watch?v=J73MELGF6u0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=2&t=1s)

- [freeCodeCamp.org - Web App Vulnerabilities - DevSecOps Course for Beginners](https://www.youtube.com/watch?v=F5KJVuii0Yw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=3&t=67s)

- [The Importance of DevSecOps and 5 Steps to Doing it Properly (DevSecOps EXPLAINED)](https://www.youtube.com/watch?v=KaoPQLyWq_g&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=4&t=13s)

- [Continuous Delivery - What is DevSecOps?](https://www.youtube.com/watch?v=NdvMUcWNlFw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=5&t=6s)

- [Cloud Advocate - What is DevSecOps?](https://www.youtube.com/watch?v=a2y4Oj5wrZg&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=6)

- [Cloud Advocate - DevSecOps Pipeline CI Process - Real world example!](https://www.youtube.com/watch?v=ipe08lFQZU8&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=7&t=204s)

Hy vọng bài viết này đã mang lại cho bạn một cái nhìn tổng quan về những gì bạn có thể mong đợi từ module này và một số tài nguyên trên sẽ giúp cung cấp thêm kiến thức có chiều sâu cho chủ đề này. Trong bài đăng vào [Ngày 3](day03.md), chúng ta sẽ xem xét cách suy nghĩ của kẻ tấn công, đó cũng là lý do tại sao chúng ta phải bảo vệ ngay từ đầu.
