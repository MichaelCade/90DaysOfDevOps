---
title: '#90DaysOfDevOps - Testing, Tools & Alternatives - Day 62'
published: false
description: '90DaysOfDevOps - Testing, Tools & Alternatives'
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049053
---
## Kiểm thử, Công cụ và Các phương pháp thay thế

Khi kết thúc phần này về Cơ sở hạ tầng dưới dạng Mã, chúng ta phải đề cập đến việc kiểm thử mã nguồn, các công cụ khác nhau có sẵn và sau đó là một số phương pháp thay thế Terraform để đạt được điều này. Như tôi đã nói ở đầu phần này, tập trung của tôi là Terraform vì nó trước tiên là miễn phí và mã nguồn mở, thứ hai, nó là đa nền tảng và không phụ thuộc vào môi trường. Nhưng cũng có các phương pháp thay thế khác mà cần được xem xét, nhưng mục tiêu chung là làm cho mọi người nhận thức rằng đây là cách triển khai cơ sở hạ tầng của bạn.

### Mã lỗi

Phần đầu tiên mà tôi muốn đề cập trong phiên này là mã lỗi. Khác với mã ứng dụng, cơ sở hạ tầng dưới dạng mã có thể được sử dụng và sau đó không được sử dụng trong một thời gian rất dài. Hãy lấy ví dụ rằng chúng ta sẽ sử dụng Terraform để triển khai môi trường máy ảo trên AWS, hoàn hảo và nó hoạt động lần đầu tiên và chúng ta có môi trường của chúng ta. Nhưng môi trường này không thay đổi quá thường xuyên, vì vậy mã bị bỏ lại trong trạng thái hiện tại, có thể hoặc hy vọng được lưu trữ ở một vị trí trung tâm, nhưng mã không thay đổi.

Điều gì sẽ xảy ra nếu có thay đổi trong cơ sở hạ tầng? Nhưng nó được thực hiện bên ngoài quy trình chính hoặc các yếu tố khác thay đổi trong môi trường của chúng ta.

Các thay đổi bên ngoài quy trình chính
Phiên bản không được gắn kết
Các phụ thuộc đã bị loại bỏ
Các thay đổi chưa được áp dụng


### Kiểm thử

Một lĩnh vực quan trọng khác liên quan đến mã lỗi và nói chung là khả năng kiểm thử IaC và đảm bảo rằng tất cả các phần đều hoạt động theo đúng yêu cầu.

Đầu tiên, chúng ta có thể xem xét một số lệnh kiểm thử tích hợp sẵn:


| Command              | Description                                                                                |
| -------------------- | ------------------------------------------------------------------------------------------ |
| `terraform fmt`      | Định dạng lại tệp terraform (.tf), giúp cho việc đọc code dễ dàng hơn                      |
| `terraform validate` | Xác thực, phát hiện lỗi cú pháp các tệp cấu hình trong một thư mục                         |
| `terraform plan`     | Tạo một kế hoạch thực thi, cho phép xem trước các thay đổi mà Terraform dự định thực hiện. |
| Custom validation    | Xác thực lại các biến đầu vào, bảo đảm khớp với những gì mà chúng ta yêu cầu               |

Ngoài ra chúng ta cũng có một số công cụ kiểm thử khác không thuộc hệ sinh thái của Hashicorp nhưng hỗ trợ terraform:

- [tflint](https://github.com/terraform-linters/tflint)

  - ìm các lỗi có thể xảy ra.
  - ảnh báo về cú pháp đã bị loại bỏ và các khai báo không được sử dụng.
  - áp dụng các quy ước tốt nhất và quy ước đặt tên.

Công cụ dò quét

- [checkov](https://www.checkov.io/) - dò quét toàn bộ hệ thống hạ tầng, phát hiện và cảnh báo các lỗi cấu hình.
- [tfsec](https://aquasecurity.github.io/tfsec/v1.4.2/) - Phân tích tính bảo mật của hạ tầng dựa vào code.
- [terrascan](https://github.com/accurics/terrascan) - Phân tích hạ tầng.
- [terraform-compliance](https://terraform-compliance.com/) - một framework kiểm thử nhẹ, tập trung vào bảo mật và tuân thủ quy định để kiểm tra terraform, giúp bạn có khả năng kiểm thử phản tác dụng cho cơ sở hạ tầng dưới dạng mã.
- [snyk](https://docs.snyk.io/products/snyk-infrastructure-as-code/scan-terraform-files/scan-and-fix-security-issues-in-terraform-files) -dò quét code phát hiện và cảnh báo các lỗi cấu hình về bảo mật.

Managed Cloud offering

- [Terraform Sentinel](https://www.terraform.io/cloud-docs/sentinel) - embedded policy-as-code framework integrated with the HashiCorp Enterprise products. It enables fine-grained, logic-based policy decisions, and can be extended to use information from external sources.

Công cụ tương tự như unit test cho terraform

- [Terratest](https://terratest.gruntwork.io/) - Terratest là một thư viện Go cung cấp mẫu và hàm kiểm thử cơ sở hạ tầng.

Một số công cụ khác nổi bật - tự động hóa

- [Terraform Cloud](https://cloud.hashicorp.com/products/terraform) - Terraform Cloud là dịch vụ quản lý của HashiCorp. Nó loại bỏ sự cần thiết của các công cụ và tài liệu không cần thiết cho các nhà thực hành, nhóm và tổ chức sử dụng Terraform trong môi trường sản xuất.

- [Terragrunt](https://terragrunt.gruntwork.io/) - Terragrunt is a là lớp bọc bên ngoài, cung cấp tính năng giữ cấu hình và triển khai, tránh việc lặp lại code, làm việc với nhiều module khác nhau.

- [Atlantis](https://www.runatlantis.io/) - Tự động hóa triển khai dựa trên gitops

### Các lựa chọn thay thế

Như đã đề cập ở ngày 57, dưới đây là các công cụ thay thế cho terraform

| Cloud Specific                  | Cloud Agnostic |
| ------------------------------- | -------------- |
| AWS CloudFormation              | Terraform      |
| Azure Resource Manager          | Pulumi         |
| Google Cloud Deployment Manager |                |

Tôi đã sử dụng AWS CloudFormation nhiều nhất trong danh sách trên, đặc biệt là trong môi trường AWS, nhưng tôi chưa sử dụng các công cụ khác ngoại trừ Terraform. Như bạn có thể tưởng tượng, các phiên bản dành riêng cho một nhà cung cấp điện toán đám mây cụ thể rất tốt trong đám mây đó, nhưng nếu bạn có nhiều môi trường đám mây khác nhau, bạn sẽ gặp khó khăn trong việc di chuyển các cấu hình đó hoặc bạn sẽ phải sử dụng nhiều mặt phẳng quản lý khác nhau cho các nỗ lực IaC của mình.

Tôi nghĩ bước tiếp theo thú vị đối với tôi là dành thời gian để tìm hiểu thêm về [Pulumi](https://www.pulumi.com/)

Theo so sánh Pulumi trên trang web của họ:

"Cả Terraform và Pulumi đều cung cấp mô hình mã hóa cơ sở hạ tầng dựa trên trạng thái mong muốn, trong đó mã đại diện cho trạng thái hạ tầng mong muốn và bộ triển khai so sánh trạng thái mong muốn này với trạng thái hiện tại của stack và xác định những tài nguyên cần tạo, cập nhật hoặc xóa."

Sự khác biệt lớn nhất mà tôi nhìn thấy là, so với ngôn ngữ cấu hình HashiCorp (HCL), Pulumi cho phép sử dụng ngôn ngữ thông dụng như Python, TypeScript, JavaScript, Go và .NET.

Tôi đã xem qua một video giới thiệu Giới thiệu về Pulumi [Introduction to Pulumi: Modern Infrastructure as Code](https://www.youtube.com/watch?v=QfJTJs24-JM): Cơ sở hạ tầng như mã hiện đại và tôi thích sự dễ dàng và sự lựa chọn mà nó mang lại. Tôi muốn tìm hiểu thêm về Pulumi.

Điều này kết thúc phần về Cơ sở hạ tầng như mã (Infrastructure as Code), tiếp theo chúng ta sẽ di chuyển đến một phần chồng chéo nhỏ với quản lý cấu hình cụ thể, đặc biệt là khi chúng ta vượt qua cái nhìn tổng quan về quản lý cấu hình lớn, chúng ta sẽ sử dụng Ansible cho một số tác vụ và bài thực hành.

## Tài liệu tham khảo

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
- [Pulumi - IaC in your favorite programming language!](https://www.youtube.com/watch?v=vIjeiDcsR3Q&t=51s)

Hẹn gặp lại các bạn ở [Ngày 63](day63.md)
