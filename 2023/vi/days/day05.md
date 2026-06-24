## Bảo mật mã nguồn mở

Phần mềm mã nguồn mở đã trở nên phổ biến trong vài năm qua nhờ vào tính cộng tác và cộng đồng/công khai của nó.

Thuật ngữ Mã nguồn mở đề cập đến phần mềm trong phạm vi công cộng mà mọi người có thể tự do sử dụng, sửa đổi và chia sẻ.

Lý do chính cho sự gia tăng này là tốc độ tăng cường mã độc quyền được phát triển nội bộ và điều này có thể rút ngắn thời gian ra mắt sản phẩm. Điều này có nghĩa là việc tận dụng OSS có thể tăng tốc phát triển ứng dụng và giúp đưa sản phẩm thương mại của bạn ra thị trường nhanh hơn.

### Bảo mật mã nguồn mở là gì?

Bảo mật mã nguồn mở đề cập đến việc đảm bảo an toàn và bảo mật cho các hệ thống máy tính và mạng sử dụng phần mềm mã nguồn mở. Như đã đề cập ở trên, phần mềm mã nguồn mở là phần mềm có sẵn miễn phí để sử dụng, sửa đổi và phân phối, và thường được phát triển bởi một cộng đồng tình nguyện viên. Tuy nhiên, có sự tham gia lớn từ các nhà cung cấp phần mềm lớn cũng đóng góp trở lại cho mã nguồn mở, bạn chỉ cần nhìn vào kho lưu trữ Kubernetes để thấy những nhà cung cấp nào đang đầu tư vào đó.

Vì phần mềm mã nguồn mở có sẵn miễn phí, nó có thể được sử dụng và nghiên cứu rộng rãi, điều này có thể giúp cải thiện bảo mật của nó. Tuy nhiên, điều quan trọng là phải đảm bảo rằng phần mềm mã nguồn mở được sử dụng một cách có trách nhiệm và bất kỳ lỗ hổng nào cũng được giải quyết kịp thời để duy trì bảo mật của nó.

### Hiểu về bảo mật chuỗi cung ứng OSS

Thông thường tôi sẽ ghi lại những phát hiện của mình dựa trên một video dài hơn thành một đoạn văn ở đây nhưng vì video này dài 10 phút nên tôi nghĩ rằng hợp lý khi liên kết tài nguyên ở đây [Hiểu về bảo mật chuỗi cung ứng mã nguồn mở](https://www.youtube.com/watch?v=pARGj6j0-ZY)

Dù là sản phẩm thương mại tận dụng OSS hay dự án OSS sử dụng các gói hoặc mã OSS khác, chúng ta phải có nhận thức từ trên xuống dưới và cung cấp khả năng hiển thị tốt hơn giữa các dự án.

### 3 A của bảo mật OSS

Một tài nguyên khác mà tôi thấy hữu ích ở đây từ IBM, sẽ được liên kết bên dưới trong phần tài nguyên.

- **Assess - Đánh giá** - Xem xét sức khỏe của dự án, kho lưu trữ hoạt động như thế nào, các nhà bảo trì phản hồi ra sao? Nếu những điều này cho thấy dấu hiệu xấu, thì bạn sẽ không hài lòng về bảo mật của dự án.

Ở giai đoạn này, chúng ta cũng có thể kiểm tra mô hình bảo mật, đánh giá mã, xác thực dữ liệu và phạm vi kiểm tra bảo mật. Dự án xử lý CVE như thế nào?

Dự án này có những dependencies nào? Khám phá sức khỏe của những dependencies này vì bạn cần đảm bảo toàn bộ stack là tốt.

- **Adopt - Áp dụng** - Nếu bạn sẽ sử dụng nó trong phần mềm của mình hoặc như một ứng dụng độc lập trong stack của riêng bạn, ai sẽ quản lý và duy trì nó? Đặt một số chính sách về người sẽ giám sát dự án và hỗ trợ cộng đồng.

- **Act - Hành động** - Bảo mật là trách nhiệm của mọi người, không chỉ của các nhà bảo trì, với tư cách là người dùng, bạn cũng nên hành động và hỗ trợ dự án.

### Lỗ hổng Log4j

Đầu năm 2022, chúng ta đã có một lỗ hổng gây chú ý lớn (Log4j (CVE-2021-44228) RCE Vulnerability)

Log4j là một thư viện rất phổ biến để ghi nhật ký trong Java. Lỗ hổng này sẽ ảnh hưởng đến hàng triệu ứng dụng dựa trên Java.

Một kẻ tấn công có thể sử dụng lỗ hổng này trong ứng dụng để truy cập vào hệ thống.

Hai điều lớn tôi đã đề cập,

- **hàng triệu** ứng dụng sẽ sử dụng gói này.
- **kẻ tấn công** có thể tận dụng điều này để truy cập hoặc cài đặt phần mềm độc hại vào môi trường.

Lý do tôi nêu ra điều này là bảo mật không bao giờ dừng lại, sự gia tăng của việc áp dụng mã nguồn mở đã tăng cường attack vector này trên các ứng dụng, và đây là lý do tại sao cần có một nỗ lực tổng thể về bảo mật từ ngày đầu tiên.

## Tài liệu tham khảo

- [Open Source Security Foundation](https://openssf.org/)
- [Snyk - State of open source security 2022](https://snyk.io/reports/open-source-security/)
- [IBM - The 3 A's of Open Source Security](https://www.youtube.com/watch?v=baZH6CX6Zno)
- [Log4j (CVE-2021-44228) RCE Vulnerability Explained](https://www.youtube.com/watch?v=0-abhd-CLwQ)

Hẹn gặp lại vào [ngày 6](day06.md).
