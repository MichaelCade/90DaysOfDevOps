## Cơ sở hạ tầng dưới dạng mã (Iac): Bức Tranh Toàn Cảnh! 

Ngày nay ở bất kì lĩnh vực nào, chúng ta đang ở trong thời kì phát triển và tận dụng tối đa máy tính, máy móc thay con người xử lý các việc lặp đi lặp lại để bảo đảm giảm thiểu sai sót cá nhân trong mọi công việc
Ví dụ: máy đóng hàng, phân loại hàng, dây chuyền sản xuất nước ngọt v..v. Hay nói riêng lĩnh vực công nghệ thông tin: xây dựng hạ tầng, tự động hóa triển khai ứng dụng 

Vậy? chúng ta - là những lập trình viên xây dựng ra các ứng dựng, hệ thống tự động hóa nói trên như thế nào
Chúng đã xây dựng những kịch bản trong trường hợp rủi ro mất toàn bộ hạ tầng vật lý, ảo hóa, điện toán đám mây?
Mất bao lâu để chúng ta có thể thay thế toàn bộ chúng trong trường hợp nêu trên?

Cơ sở hạ tầng dưới dạng mã ra đời nhằm giúp chúng ta có khả năng trả lời các câu hỏi trên, kiểm tra, thay thế trong trường hợp sự cố, tuy nhiên chúng ta không nên nhầm lẫn chúng với khái niệm sao lưu và phục hồi. Thay vào đó, xét theo khía cạnh về cơ sở hạ tầng và các môi trường phát triển phần mềm, chúng ta có khả năng triển khai, vận hành và thay đổi chúng một cách thân thiện, và dễ dàng nhất.

TLDR; Chúng ta có khả năng sử dụng mã để xây dựng lại toàn bộ môi trường từ đầu.

Nếu còn nhớ ngay từ những ngày đầu tiên chúng ta đã nhắc đến khái niệm DevOps nói chung là một cách để phá vỡ rào cản để cung cấp hệ thống cho môi trường (production) của một hệ thống công nghệ thông tin một cách an toàn và nhanh chóng.

Cơ sở hạ tầng dưới dạng mã giúp chúng ta cung cấp các hệ thống này, chúng ta cũng đã nói đến rất nhiều các quy trình, công cụ. IaC mang đến chúng ta những công cụ đơn giản để thực hiện quy trình triển khai hạ tầng như vậy.

Ở phần này, chúng ta sẽ tập trung vào cơ sở hạ tầng dưới dạng mã (IaC). Các bạn cũng có thể nghe nói đến cơ sở hạ tầng từ mã, cấu hình dạng mã (CaS). Tuy nhiên xét đến mức độ phổ biến chúng ta sẽ hiểu nó là Cơ Sở Hạ Tầng Dưới Dạng Mã (IaC)

Ở bài viết gốc, tác giả nhắc đến việc sử dụng IaC giống như chúng ta đang chăn nuôi gia súc và vật nuôi trong nhà với sự hài hòa, thân thiện với con người.

### Tính thân thiện

Cùng xem xét đến các yếu tố ban đầu của DevOps, chúng ta đã nghe nói đến yêu cầu của việc xây dựng một ứng dụng mới, hầu hết trước đây chúng ta sẽ triển khai, chuẩn bị các hệ thống máy chủ một cách thủ công. Các bước này bao gồm:

- Triển khai máy ảo | Trước khi có công nghệ ảo hóa thì sẽ là các máy chủ vật lý và cài đặt hệ điều hành
- Cài đặt, cấu hình mạng
- Cài đặt, cấu hình bảng định tuyến
- Cài đặt các phần mềm, thư viện cần thiết và cập nhật chúng
- Cấu hình phần mềm
- Cài đặt cơ sở dữ liệu

Đương nhiên là trước đây, tất cả các tác vụ này đều được thực hiện thủ công bởi chuyên viên quản trị hệ thống SysAdmin. Đối với ứng dụng có quy mô lớn hơn, cần nhiều tài nguyên và máy chủ hơn, đương nhiên sẽ cần nhiều công sức hơn để cài đặt chúng. Do vậy chúng sẽ lấy đi rất nhiều công sức của nhân lực lao động (là chính chúng ta) cũng như thời gian, điều mà các doanh nghiệp sẽ phải chi trả trong toàn bộ quá trình xây dựng môi trường công nghệ thông tin ứng dụng này. Đặc biệt là ngay từ đầu nội dung buổi hôm nay tác giả đã chia sẻ, việc gặp rủi ro từ cá nhân đều có thể xẩy ra khiến chi phí để bù lấp sai lầm này có thể sẽ rất lơn, vì vậy tự động hóa sẽ là sự lựa chọn hàng đầu.

Quay trở lại nội dung chính, tiếp theo sau khi các bước cài đặt ban đầu hoàn thiện, chúng ta vẫn còn cần tiến hành bảo trì, quản lý và vận hành các hệ thống máy chủ nói trên, bao gồm:

- Cập nhật các phiên bản mới
- Triển khai các phiên bản mới này
- Quản lý dữ liệu
- Khôi phục ứng dụng nếu có sự cố
- Loại bỏ, thêm bớt và mở rộng các máy chủ, tài nguyên phần cứng trong trường hợp cần thiết
- Cấu hình mạng

Các bạn hãy tưởng tượng từng đó công việc sẽ được lặp đi lặp lại cho các môi trường phát triển như: dev, test, production ... độ phức tạp vì vậy cũng tăng lên.

Đó cũng chính là lý do mà IaC giải quyết được các vấn đề này, việc triển khai lúc này cũng giống như việc chúng ta nuôi dưỡng một chú mèo, chúng ta có thể đặt tên cho chúng, nuôi dưỡng chúng và sống cùng chúng ta, gọi là vòng đời của một phần mềm (SDLC)

Với IaC, chúng ta tự động hóa tất cả các bước nêu trên từ đầu đến cuối. IaC chính là khái niệm và một số các công cụ chúng ta quan tâm để khởi tạo tự động hạ tầng, trong trường hợp xấu, rủi ro xẩy ra và ảnh hưởng đến hệ thống máy chủ, chúng ta hoàn toàn có thể loại bỏ nó và khởi tạo một máy chủ khác để thay thế nó. Cũng chính vì nó được tự động hóa hoàn toàn nên khi máy chủ mới được khởi tạo nó sẽ giống chính xác hoàn toàn với các đoạn mã mà chúng ta viết. **Câu này khó quá** `At this point we don't care what they are called they are there in the field serving their purpose until they are no longer in the field and we have another to replace it either because of a failure or because we updated part or all of our application.`

IaC được sử dụng ở hầu hết mọi nền tảng, công nghệ ảo hóa, công nghệ điện toán đám mấy và công nghệ ứng dụng khai thác lợi ích của điện toán đám mây (Cloud Native) như Kubernetes hay containers.

## Khởi tạo cơ sở hạ tầng
Thay vì việc nhắc đến `chef`, `puppet` và `ansible` ngay từ đầu là những công cụ vô cùng phù hợp để giải quyết yêu cầu khởi tạo, cài đặt ứng dụng và quản lý chúng.

Ở các tác vụ về việc khởi tạo cơ sở hạ tầng dưới đây, IaC không hoàn toàn hỗ trợ được hết các tác vụ này. Tuy nhiên với Terraform (quan điểm cá nhân mình đánh giá đây không phải ngôn ngữ lập trình mà là dạng khai báo), đây là công cụ cho phép chúng ta thực hiện tốt và đầy đủ 2 yếu tố được nêu ở phía dưới, nó cho phép chúng ta khởi tạo hạ tầng từ số 0, chúng ta định nghĩa hạ tầng bằng các đoạn mã với các tham số mà chúng ta đã lên kịch bản rồi triển khai nó (deploy). IaC cũng cho phép chúng ta quản lý hạ tầng và cũng có thể triển khai ứng dụng, tuy nhiên hạn chế của nó là không theo dõi được trạng thái hiện tại của ứng dụng sau khi chúng ta triển khai, về việc này ở nội dung bài sau chúng ta sẽ đề cập đến một công cụ khác là Ansile - một công cụ quản lý cấu hình.

Vậy các công việc khởi tạo, cài đặt cấu hình một phần mềm sẽ gồm những gì?

- Khởi tạo máy chủ đồng loạt
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

### Difference between IaC tools

Declarative vs procedural

Procedural

- Step-by-step instruction
- Create a server > Add a server > Make this change

Declarative

- declare the result
- 2 Servers

Mutable (pets) vs Immutable (cattle)

Mutable

- Change instead of replacing
- Generally long lived

Immutable

- Replace instead of change
- Possibly short-lived

This is really why we have lots of different options for Infrastructure as Code because there is no one tool to rule them all.

We are going to be mostly using terraform and getting hands-on as this is the best way to start seeing the benefits of Infrastructure as Code when it is in action. Getting hands-on is also the best way to pick up the skills you are going to be writing code.

Next up we will start looking into Terraform with a 101 before we get some hands-on getting used.

## Resources

I have listed a lot of resources down below and I think this topic has been covered so many times out there, If you have additional resources be sure to raise a PR with your resources and I will be happy to review and add them to the list.

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

See you on [Day 57](day57.md)
