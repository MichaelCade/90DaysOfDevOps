---
title: '#90DaysOfDevOps - Mô hình Điện toán Microsoft Azure - Ngày 31'
published: false
description: 90DaysOfDevOps - Mô hình Điện toán Microsoft Azure
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049040
---

## Mô hình Điện toán Microsoft Azure

Tiếp theo phần trình bày về những kiếm thức cơ bản về các mô hình bảo mật trong Microsoft Azure ngày hôm qua, hôm nay chúng ta sẽ xem xét các dịch vụ điện toán khác nhau có sẵn trong Azure.

### Tùy chọn khả dụng của dịch vụ (Service Availability Options)

Phần này rất gần gũi với tôi dưới vai trò Data Management. Giốn như với on-premises, điều tối quan trọng là phải đảm bảo tính khả dụng của các dịch vụ của bạn.

- Tính sẵn sàng cao (high availability) (trong một region)
- Khắm phục thảm hoạ (disaster recovery) (giữa các regions)
- Sao lưu (backup) (phục hồi từ một thời điểm)

Microsoft triển khai nhiều regions trong một ranh giới địa chính trị.

Hai khái niệm với Azure cho tính khả dụng của dịch vụ: sets và zones.

Availability Sets - Cung cấp khả năng phục hồi trong trung tâm dữ liệu

Availability Zones - Cung cấp khả năng phục hồi giữa các trung tâm dữ liệu trong một khu vực.

### Máy ảo (Virtual Machines)

Gần như đây là điểm khởi đầu cho bất cứ ai làm việc với đám mây công cộng.

- Cung cấp máy ảo (VM) từ nhiều dòng và kích cỡ, cấu hình khác nhau với nhiều khả năng khác nhau (đôi khi hơi quá so với nhu cầu) [Kích cỡ, cấu hình cho máy ảo của Azure](https://docs.microsoft.com/en-us/azure/virtual-machines/sizes)
- Có nhiều tuỳ chọn và những điểm tập trung khác nhau cho máy ảo, từ hiệu năng cao, độ trễ thấp tới những máy ảo có bộ nhớ tối ưu.
- Chúng ta cũng có loại VM có thể burst (dành cho workloads vượt quá cấu hình định mức) trong dòng B. Điều này rất tốt cho các workloads mà bạn có thể yêu cầu CPU thấp hầu hết thời gian nhưng lại yêu cầu tăng đột biến với tần suất có thể là mỗi tháng một lần.
- Máy ảo được đặt trên một mạng ảo có thể cung cấp với bất cứ mạng nào khác.
- Hỗ trọ hệ điều hành Linux và Windows.
- Ngoài ra còn có các kernels được điều chỉnh cho Azure khi nói đến các bản phân phối Linux cụ thể. [Azure Tuned Kernals](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/endorsed-distros#azure-tuned-kernels)

### Tạo mẫu (Templating)

Như tôi đã đề cập trước đó, mọi thứ đằng sau hoặc bên dưới Microsoft Azure đều là JSON.

Có một số cổng quản lý và bảng điều khiển khác nhau mà chúng ta có thể sử dụng để tạo tài nguyên của mình, cách được ưu tiên là thông qua các JSON templates.


Idempotent deployments ở chế độ incremental hoặc complete - ví dụ: trạng thái mong muốn có thể được lặp lại nhiều lần (repeatable desired state)

Có nhiều lựa chọn templates có thể export các định nghĩa tài nguyên đã được triển khai. 
Tôi nghĩ về tính năng tạo template này là một thứ gì đó giống với AWS CloudFormation hoặc có thể là Terraform cho lựa chọn multi-cloud. Chúng ta sẽ đề cập nhiều hơn tới Terraform trong phần Cơ sở hạ tầng dưới dạng mã.

### Thay đổi quy mô (Scaling)

Tự động thay đổi quy mô (automatic scaling) là một tính năng lớn của đám mây công cộng, có thể rút bớt các tài nguyên mà bạn không sử dụng hoặc tăng thêm khi bạn cần chúng.

Với Azure, chúng ta có một thứ gọi là Virtual Machine Scale Sets (VMSS) cho IaaS. Điều này cho phép tạo và scale tự động một images tiêu chuẩn theo các lịch trình và số liệu (schedules & metrics)

Việc này rất lý tưởng cho việc thiết lập các khoảng cửa sổ cập nhật (updating windows) để bạn có thể cập nhật các images và triểu khai chúng với tác động ít nhất.

Các dịch vụ khác ví dụ như Azure App Service được tích hợp sẵn tính năng tự động thay đổi quy mô (auto-scaling).

### Containers

Chúng ta chưa đề cập chi tiết tới containers trong hành trình học DevOps này nhưng tôi nghĩ chúng ta cần đề cập tới việc Azure có một số dịch vụ tập trung vào containers. 

Azure Kubernetes Service (AKS) - Cung cấp giải pháp quản lý Kubernetes, không cần lo lắng về control plane hay quản lý cụm tài nguyên (cluster). Chúng ta sẽ tìm hiểu về Kubernetes thêm vào phần sau.

Azure Container Instances - Containers as a service với mô hình thanh toán theo giây (per-second billing). Chạy một image và tích hợp nó với mạng ảo của bạn, không cần container orchestration (điều phối container)

Service Fabric - Có rất nhiều khả năng nhưng bao gồi điều phối các container instances.

Azure cũng có Container Registry cung cấp một registry riêng cho Docker images, Helm charts, OCI Artifacts và images. Chúng ta sẽ tìm hiểu kỹ hơn trong phần về containers.

Chúng ta cũng nên đều cập rằng rất nhiều dịch vụ container cũng có thể sử dụng container ở phía dưới nhưng điều này cũng được trừu tượng hoá và bạn không cần phải quản lý chúng.

Chúng ta cũng có thể thấy các dịch vụ tập trung vào container này xuất hiện ở mọi đám mây công cộng khác.

### Application Services

- Azure Application Services cung cấp giải pháp hosting với một phương pháp đơn giản để triển khai, thiết lập dịch vụ.
- Tự động triểu khai và mở rộng quy mô.
- Hỗ trợ các giải pháp dựa trên Windows và Linux.
- Các dịnh vụ chạy trong một App Service Plan được chọn type và size.
- Thích hợp cho các dịch vụ khác nhau bao gồm wep apps, API apps và mobile apps.
- Hỗ trợ Deployment slots cho việc triển khai và kiểm thử.

### Serverless Computing

Serverless đối với tôi là một chủ đề thú vị mà tôi cực kỳ muốn tìm hiểu thêm.

Mục tiêu với serverless là chúng ta chỉ trả tiền cho thời gian chạy của một hàm và không cần phải chạy các máy ảo hoặc ứng dụng PaaS mọi lúc. Chúng ta đơn giản là chỉ chạy hàm của mình khi cần và sau đó nó biến mất.

Azure Functions - Cung cấp mã serverless. Nếu chúng ta nhớ lại những điều được đề cập trong bài đầu tiên về điện toán đám mấy, chúng ta sẽ nhớ về các lớp quản lý trừu tượng, với serverless funtion, chúng ta chỉ cần quản lý code.

Event-Driven với quy mô siêu lớn, tôi có kế hoạch làm một thứ gì đó khi tôi thực hành phần này, hy vọng sẽ thực hiện được ở phần sau.

Cung cấp input và output cho nhiều dịch vụ Azure và của bên thứ ba.

Hỗ trợ nhiều ngôn ngữ lập trình. (C#, NodeJS, Python, PHP, batch, bash, Golang and Rust. hoặc bất kỳ tệp thực thi nào)

Azure Event Grid cho phép có thể kích hoạt logic từ các dịch vụ khác hoặc events.

Azure Logic App cung cấp một quy trình làm việc và tích hợp dựa trên đồ hoạ.

Chúng ta cũng có thể nhắc tơi Azure Batch, dịch vụ giúp chạy các công việc có quy mô lớn trên các nodes Windows hoặc Linux với khả năng quản lý và lập lịch nhất quán.

## Tài liệu tham khảo

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

Hẹn gặp lại vào [ngày 32](day32.md)
