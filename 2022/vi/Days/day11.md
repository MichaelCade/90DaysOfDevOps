---
title: "#90DaysOfDevOps - Biến, hằng số & kiểu dữ liệu - Ngày 11"
published: false
description: 90DaysOfDevOps - Biến và hằng số trong Go
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048862
---

Trước khi đi vào chủ đề ngày hôm nay, tôi muốn gửi lời cảm ơn tới [Techworld with Nana](https://www.youtube.com/watch?v=yyUHQIec83I) với video đi hết những kiến thức cơ bản về Go.

Vào [ngày 8](day08.md), chúng ta thiết lập môi trường, vào [ngày 9](day09.md), chúng ta đã xem qua mã Hello #90DaysOfDevOps và vào [ngày 10](day10.md)), chúng ta đã tìm hiểu về không gian làm việc Go và đi sâu hơn một chút vào biên dịch và chạy mã.

Hôm nay chúng ta sẽ tìm hiểu về Biến, Hằng số và Kiểu dữ liệu trong khi viết một chương trình mới.

## Biến và Hằng số trong Go

Hãy bắt đầu bằng cách lên kế hoạch cho ứng dụng của chúng ta. Tôi nghĩ một chương trình cho bạn biết số ngày còn lại trong thử thách #90DaysOfDevOps sẽ là một ý tưởng hay.

Đầu tiên là khi chúng ta xây dựng ứng dụng, nó chào mừng người tham gia và nhận phản hồi từ người dùng về số ngày đã hoàn thành. Chúng ta có thể sử dụng thuật ngữ #90DaysOfDevOps nhiều lần trong suốt chương trình và đây là trường hợp hoàn hảo để #90DaysOfDevOps trở thành một biến trong chương trình.

- Các biến được sử dụng để lưu trữ các giá trị.
- Giống như một hộp nhỏ chứa thông tin hoặc giá trị của chúng ta.
- Biến này có thể được sử dụng trong suốt chương trình và cũng có ưu điểm là nếu bạn muốn thay đổi tên thử thách hoặc biến này, bạn chỉ phải thay đổi nó ở một nơi. Nói cách khác, bằng cách thay đổi một giá trị của biến này, nó có thể được chuyển sang các tên các thử thách khác trong cộng đồng.

Để khai báo điều này trong Go, hãy sử dụng **từ khóa** cho các biến. Khai báo này sẽ được sử dụng trong khối mã `func main` mà chúng ta sẽ nhắc tới sau. Giải thích chi tết về **Từ khoá** tại [đây](https://go.dev/ref/spec#Keywords).

Hãy nhớ rằng tên biến có tính mô tả. Nếu bạn khai báo một biến, bạn phải sử dụng nó hoặc bạn sẽ gặp lỗi, điều này để tránh có thể có mã chết (mã không bao giờ được sử dụng). Điều này cũng tương tự cho các gói (packages) không được sử dụng.

```
var challenge = "#90DaysOfDevOps"
```

Với khai báo ở trên, bạn có thể thấy chúng ta đã sử dụng một biến khi in ra chuỗi ký tự ở đoạn mã dưới đây.

```
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    fmt.Println("Welcome to", challenge, "")
}
```

Bạn có thể tìm thấy đoạn mã trên tại [day11_example1.go](../../Days/Go/day11_example1.go)

Sau đó, chúng ta xây dựng mã của với ví dụ trên và nhận được kết quả hiển thị như dưới đây.

![](../../Days/Images/Day11_Go1.png)

Chúng ta cũng biết rằng thử thách này kéo dài ít nhất 90 ngày, nhưng với thử thách tiếp theo, nó có thể là 100 ngày, chính vì thế, chúng ta cũng cần một biến ở đây. Tuy nhiên, với chương trình này, chúng ta muốn khai báo nó như một hằng số. Các hằng số cũng giống như các biến, ngoại trừ việc giá trị của chúng không thể thay đổi trong đoạn mã (chúng ta vẫn có thể tạo một ứng dụng mới với mã được giữ nguyên và thay đổi hằng số này nhưng giá trị 90 sẽ không thay đổi khi chúng ta chạy ứng dụng của mình)

Thêm `const` vào mã và thêm một dòng mã khác để in ra kết quả.

```
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge")
}
```

Bạn có thể tìm thấy đoạn mã trên tại [day11_example2.go](../../Days/Go/day11_example2.go)

Nếu chúng ta thực hiện lại câu lệnh `go build` và chạy lại, bạn sẽ thấy kết quả bên dưới.

![](../../Days/Images/Day11_Go2.png)

Đây sẽ không phải là phần cuối của chương trình, chúng ta sẽ quay lại vào [ngày 12](day12.md) để thêm các chức năng khác. Bây giờ ta sẽ thêm một biến khác cho số ngày đã hoàn thành trong thử thách.

Bên dưới, tôi đã thêm biến `dayscomplete` với số ngày đã hoàn thành.

```
package main

import "fmt"

func main() {
    var challenge = "#90DaysOfDevOps"
    const daystotal = 90
    var dayscomplete = 11

    fmt.Println("Welcome to", challenge, "")
    fmt.Println("This is a", daystotal, "challenge and you have completed", dayscomplete, "days")
    fmt.Println("Great work")
}
```

Bạn có thể tìm thấy đoạn mã trên tại [day11_example3.go](../../Days/Go/day11_example3.go)

Hãy chạy lại lệnh `go build` hoặc có thể sử dụng `go run`

![](../../Days/Images/Day11_Go3.png)

```
package main

import "fmt"

func main() {
	var challenge = "#90DaysOfDevOps"
	const daystotal = 90
	var dayscomplete = 11

	fmt.Printf("Welcome to %v\n", challenge)
	fmt.Printf("This is a %v challenge and you have completed %v days\n", daystotal, dayscomplete)
	fmt.Println("Great work")
}
```

Đây là một số ví dụ khác để làm cho mã dễ đọc và chỉnh sửa hơn. Hiện giờ, chúng ta vẫn đang sử dụng hàm `Println` nhưng nó có thể được đơn giản hóa bằng cách sử dụng` Printf` với `%v`, có nghĩa là các biến theo sẽ được xác định ở cuối dòng mã. Chúng ta cũng có thể sử dụng `\n` để xuống dòng.

Tôi đang sử dụng `%v` vì nó là giá trị mặc định, các tùy chọn khác có ở đây trong [tài liệu của gói fmt](https://pkg.go.dev/fmt), bạn có thể đoạn mã ví dụ tại [day11_example4.go](../../Days/Go/day11_example4.go)

Các biến cũng có thể được khai bảo một cách đơn giản hơn trong mã của bạn. Thay vì xác định rằng đó là `var` và` type`, bạn có thể viết mã này như sau để có cùng kết quả nhưng đoạn mã sẽ gọn, đẹp và đơn giản hơn. Cách này chỉ áp dụng được với các biến, không sử dụng với hằng số.

```
func main() {
    challenge := "#90DaysOfDevOps"
    const daystotal = 90
```

## Kiểu dữ liệu

Trong các ví dụ trên, chúng ta chưa xác định kiểu dữ liệu của biến, điều này là do chúng ta đã gán cho nó một giá trị và Go đủ thông minh để biết kiểu dữ liệu đó là gì hoặc ít nhất có thể suy ra nó là gì dựa trên giá trị bạn đã gán. Tuy nhiên, nếu chúng ta muốn người dùng nhập dữ liệu, chúng ta phải sử dụng một kiểu dữ liệu cụ thể.

Cho đến giờ, chúng ta đã sử dụng Chuỗi và Số nguyên trong mã của mình. Số nguyên cho số ngày và chuỗi là tên của thử thách.

Điều quan trọng cần lưu ý là mỗi kiểu dữ liệu có thể làm những việc khác nhau và hoạt động khác nhau. Ví dụ: số nguyên có thể nhân lên trong khi chuỗi thì không.

Có bốn loại:

- **Loại cơ bản - Basic type**: Số, chuỗi và boolean (Numbers, strings, booleans)
- **Loại tổng hợp - Aggregate type**: Mảng và cấu trúc (Array, structs)
- **Loại tham chiếu -Reference type**: Con trỏ, lát cắt, bản đồ, chức năng và kênh (Pointers, slices, maps, functions, and channels)
- **Loại giao diện - Interface type**

Kiểu dữ liệu là một khái niệm quan trọng trong lập trình. Kiểu dữ liệu chỉ định kích thước và kiểu của các giá trị biến.

Go được nhập tĩnh, có nghĩa là khi một kiểu dữ liệu của biến được xác định, nó chỉ có thể lưu trữ dữ liệu của kiểu đó.

Go có ba kiểu dữ liệu cơ bản:

- **bool**: đại diện cho một giá trị boolean hoặc đúng hoặc sai
- **Numeric**: đại diện cho kiểu số nguyên, giá trị dấu phẩy động và kiểu phức tạp
- **string**: đại diện cho một giá trị chuỗi

Tôi thấy đây là nguồn tài liệu rất chi tết về các kiểu dữ liệu [Golang by example](https://golangbyexample.com/all-data-types-in-golang-with-examples/)

Tôi cũng sẽ đề xuất video [Techworld with Nana](https://www.youtube.com/watch?v=yyUHQIec83I&t=2023s). Tại thời điểm này, video đề cập chi tiết rất nhiều về các loại dữ liệu trong Go.

Chúng ta có thể làm như sau khi cần khai báo kiểu cho biến của mình:

```
var TwitterHandle string
var DaysCompleted uint
```

Bởi vì Go ngụ ý các biến trong đó một giá trị được đưa ra, chúng ta có thể in ra các giá trị đó như sau:

```
fmt.Printf("challenge is %T, daystotal is %T, dayscomplete is %T\n", conference, daystotal, dayscomplete)
```

Có nhiều kiểu số nguyên và kiểu float khác nhau, các liên kết ở trên sẽ trình bày chi tiết về những kiểu này.

- **int** = số nguyên
- **unint** = số nguyên dương
- **các loại dấu phẩy động** = các số có chứa thành phần thập phân

## Tài liệu tham khảo

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Ở phần sau, chúng ta sẽ thêm một số chức năng nhập liệu của người dùng vào chương trình để người dùng có thể trả lời câu hỏi đã hoàn thành bao nhiêu ngày.

Hẹn gặp lại tại [ngày 12](day12.md).
