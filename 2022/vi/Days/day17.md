---
title: "#90DaysOfDevOps - Text Editors - nano vs vim - Ngày 17"
published: false
description: 90DaysOfDevOps - Text Editors - nano vs vim
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048703
---

## Text Editors - nano vs vim

Phần lớn các hệ thống Linux của bạn sẽ là máy chủ và chúng sẽ không có GUI. Tôi cũng đã đề cập trong phần trước rằng Linux chủ yếu được tạo thành từ các tệp cấu hình, để thực hiện các thay đổi, bạn cần có khả năng chỉnh sửa các tệp cấu hình đó để quản lý hệ thống.

Có rất nhiều tùy chọn ngoài kia nhưng tôi nghĩ có lẽ chúng ta nên đề cập đến hai trình soạn thảo văn bản trên terminal phổ biến nhất. Tôi đã sử dụng cả hai trình chỉnh sửa này và đối với tôi, tôi thấy `nano` dễ sử dụng hơn cho các thay đổi nhanh nhưng xét về các chức năng, `vim` là editor nổi trội hơn.

### nano

- Không có sẵn trên mọi hệ thống.
- Dễ để bắt đầu hơn

Nếu bạn chạy `nano 90DaysOfDevOps.txt`, chúng ta sẽ có một tệp mới không có gì trong đó, từ đây chúng ta có thể thêm văn bản của mình và hướng dẫn bên dưới là những gì chúng ta muốn thay đổi tệp đó.

![](../../Days/Images/Day17_Linux1.png)

Bây giờ chúng ta có thể sử dụng `control x + enter` và sau đó chạy `ls` để có thể thấy tệp văn bản mới của chúng ta.

![](../../Days/Images/Day17_Linux2.png)

Bây giờ chúng ta có thể chạy `cat` đối với tệp đó để đọc tệp của mình. Sau đó, chúng ta có thể sử dụng cùng `nano 90DaysOfDevOps.txt` đó để thêm văn bản bổ sung hoặc sửa đổi tệp của bạn.

Đối với tôi, sử dụng nano là cực kỳ dễ dàng khi thực hiện các thay đổi nhỏ trên các tệp cấu hình.

### vim

Có thể coi đây là trình soạn thảo văn bản phổ biến nhất? Là anh em của trình soạn thảo văn bản UNIX vi từ năm 1976, chúng ta có rất nhiều chức năng với vim.

- Được hỗ trợ trên hầu hết mọi bản phân phối Linux.
- Vô cùng đầy đủ! Bạn có thể tìm thấy một khóa học kéo dài 7 giờ chỉ bao gồm vim.

Chúng ta có thể chuyển sang vim bằng lệnh `vim` hoặc nếu muốn chỉnh sửa tệp txt của mình, chúng ta có thể chạy lệnh `vim 90DaysOfDevOps.txt` nhưng trước tiên bạn sẽ thấy thiếu các menu trợ giúp ở phía dưới so với `nano`

Câu hỏi đầu tiên có thể là "Làm cách nào để thoát vim?" đầu tiên là nhấn `escape` và nếu chúng ta không thực hiện bất kỳ thay đổi nào thì nhập `:q`

![](../../Days/Images/Day17_Linux3.png)

Bạn bắt đầu ở chế độ `normal`, có các chế độ `command, normal, visual, insert`, nếu muốn thêm văn bản, chúng ta sẽ cần chuyển từ `normal` sang `insert` bằng phím `i`, nếu bạn đã thêm/sửa văn bản và muốn lưu lại những thay đổi này, nhấm `escape` và sau đó là `:wq`

![](../../Days/Images/Day17_Linux4.png)

![](../../Days/Images/Day17_Linux5.png)

Bạn có thể xác nhận bằng lệnh `cat` để kiểm tra xem bạn đã lưu các thay đổi đó chưa.

Có một số chức năng thú vị của vim cho phép bạn thực hiện các tác vụ đơn giản rất nhanh nếu bạn biết các phím tắt, bản thân việc đó có thể được coi là một bài học. Giả sử chúng ta đã thêm một danh sách các từ được lặp lại và bây giờ chúng ta cần thay đổi danh sách đó, có thể đó là tệp cấu hình và tên mạng đã được lặp lại nhiều lần. Chúng ta muốn nhanh chóng thay đổi điều tên mạng cho tệp cấu hình. Tôi đang sử dụng từ "Day" cho ví dụ này.

![](../../Days/Images/Day17_Linux6.png)

Bây giờ chúng ta muốn thay thế từ đó bằng 90DaysOfDevOps, có thể thực hiện việc này bằng cách nhấn `ESC` và nhập `:%s/Day/90DaysOfDevOps`

![](../../Days/Images/Day17_Linux7.png)

Kết quả khi bạn nhấn enter là từ "Day" được thay thế bằng "90DaysOfDevOps".

![](../../Days/Images/Day17_Linux8.png)

Copy và Paste thực sự khiến tôi mở rộng tầm mắt. Copy không phải là copy mà là yank. Chúng ta có thể sao chép bằng `yy` trên bàn phím ở chế độ `normal`. `p` đẻ paste trên cùng một dòng, `P` paste trên một dòng mới.

Bạn cũng có thể xóa những dòng này bằng cách chọn số dòng bạn muốn xóa, sau đó là `dd`

Cũng có thể đôi khi bạn sẽ cần tìm kiếm trong một tệp, chúng ta có thể sử dụng `grep` như đã đề cập trong ngày hôm trước nhưng cũng có thể sử dụng vim. Sử dụng `/word` và thao tác này sẽ tìm kết quả khớp đầu tiên, để chuyển qua phần tiếp theo, bạn sẽ sử dụng phím `n`, v.v.

Đối với vim, những điều này thậm chí còn chưa phải là giớt thiệu qua về công cụ này. Lời khuyên mà tôi có thể đưa ra là hãy bắt tay vào sử dụng vim bất cứ khi nào có thể.

Một câu hỏi phỏng vấn phổ biến là trình soạn thảo văn bản yêu thích của bạn trong Linux là gì và tôi sẽ đảm bảo rằng bạn có ít nhất kiến ​​​​thức này về cả hai trình soạn thảo trên để có thể trả lời, nano là một câu trả lời không tệ vì nó đơn giản. Ít nhất thì bạn cũng thể hiện được khả năng hiểu trình soạn thảo văn bản là gì. Nhưng hãy thực hành với chúng để thành thạo hơn.

Một con trỏ khác để điều hướng trong vim, chúng ta có thể sử dụng `H,J,K,L` cũng như các phím mũi tên của mình.

## Resources

- [Vim in 100 Seconds](https://www.youtube.com/watch?v=-txKSRn0qeA)
- [Vim tutorial](https://www.youtube.com/watch?v=IiwGbcd8S7I)
- [Learn the Linux Fundamentals - Part 1](https://www.youtube.com/watch?v=kPylihJRG70)
- [Linux for hackers (don't worry you don't need to be a hacker!)](https://www.youtube.com/watch?v=VbEx7B_PTOE)

Hẹn gặp lại vào [ngày 18](day18.md)
