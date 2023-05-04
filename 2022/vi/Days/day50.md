---
title: '#90DaysOfDevOps - Chọn nền tảng chạy Kubernetes - Ngày 50'
published: false
description: 90DaysOfDevOps - Chọn nền tảng chạy Kubernetes
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049046
---

## Chọn nền tảng chạy Kubernetes

Tôi muốn sử dụng phần này để chia nhỏ một số nền tảng, một thách thức với thế giới Kubernetes là loại bỏ, giảm bớt sự phức tạp.

Kubernetes đã trải qua một chặng đường khó khăn về cách xây dựng từ con số không thành một Kubernetes cluster đầy đủ chức năng, điều này cực quan trọng nhưng ngày càng nhiều, ít nhất những người mà tôi đang nói chuyện đang muốn loại bỏ sự phức tạp đó và chạy Kubernetes cluster được quản lý. Vấn đề là nó tốn nhiều tiền hơn nhưng lợi ích đem lại nếu bạn sử dụng dịch vụ được quản lý, bạn có cần biết kiến thức node nền tảng và điều gì đang xảy ra trong Control Plan node khi nhìn chung bạn không có quyền truy cập vào nó.

Sau đó, chúng ta có các bản phân phối phát triển cục bộ cho phép chúng ta sử dụng hệ thống của mình và chạy phiên bản Kubernetes cục bộ để các nhà phát triển có thể có môi trường làm việc đầy đủ để chạy ứng dụng của họ trong nên tảng mà họ dự định sử dụng.

Cơ sở chung của tất cả các khái niệm này là chúng đều là một loại của Kubernetes, nghĩa là chúng ta có thể tự do migrate và di chuyển workloads của mình đến nơi chúng ta cần để phù hợp với yêu cầu của mình.

Rất nhiều khi lựa chọn của chúng ta cũng sẽ phụ thuộc vào những khoản đầu tư đã được thực hiện. Tôi cũng đã đề cập đến trải nghiệm của nhà phát triển nhưng một số môi trường Kubernetes cục bộ chạy trên máy tính xách tay của chúng ta cũng đã rất tuyệt vời để nắm bắt công nghệ mà không tốn bất kỳ khoản tiền nào.

### Bare-Metal Clusters

Một lựa chọn cho nhiều người có thể là chạy hệ điều hành Linux của bạn trực tiếp trên một số máy chủ vật lý để tạo cluster của chúng ta, đó cũng có thể là Windows nhưng tôi chưa nghe nhiều tỷ lệ tương thích của Windows, Containers và Kubernetes. Nếu chúng ta là một doanh nghiệp và quyết định một khoản đầu tư vào các máy chủ vật lý thì đây có thể là lựa chọn của bạn cho việc chạy Kubernetes cluster, tuy nhiên bạn sẽ phải tự xây dựng và quản lý mọi thứ từ đầu.

### Ảo hoá (Virtualisation)

Bất kể môi trường ảo hoá bao gồm thử nghiệm hay học tập hay sẵn sàng cho các enterprise của Kubernetes cluster đều là một cách tốt để bạn bắt đầu, điển hình là việc tạo ra các máy ảo để hoạt động như các node của bạn sau đó nhóm chúng lại với nhau cũng như tận dụng khoản tiền đã chi. Ví dụ như VMware cung cấp một giải pháp tuyệt vời cho cả máy ảo và Kubernetes với nhiều tuỳ chọn khác nhau. 

Kubernetes cluster của tôi được xây dựng dựa trên ảo hoá bằng Microsoft Hyper-V trên một máy chủ cũ mà tôi có. Máy chủ này có khả năng chạy một số máy ảo làm node của tôi.

### Các tuỳ chọn trên máy tính để bàn

Có một số lựa chọn khi chạy Kubernetes cluster cục bộ trên máy tính để bàn hoặc máy tính xách tay của bạn. Như đã nói ở trên, điều này mang lại cho các nhà phát triển khả năng kiểm tra ứng dụng của họ mà không cần phải có nhiều cluster phức tạp hoặc tốn kém. Cá nhân tôi nghĩ rằng đây là một trong những thứ mà tôi đã sử dụng rất nhiều và cụ thể là tôi đã dùng minikube. Nó có một số chức năng và tiện ích bổ sung rất tốt giúp thay đổi cách bạn thiết lập và chạy một thứ gì đó.

### Dịch vụ Kubernetes được quản lý

Tôi đã đề cập đến ảo hoá và điều này có thể thực hiện được với các trình ảo hoá cục bộ, chúng ta đã biết từ các phần trước rằng chúng ta cũng có thể tận dụng các máy ảo trong dịch vụ điện toán đám mây công cộng để hoạt động như các nodes của chúng ta. Điều tôi đang muốn nhắc tới ở đây khi nói tới các dịch vụ được quản lý của Kubernetes là các dịch vụ mà chúng ta thấy từ các công ty điện toán đám mây cũng như từ các MSP (Managed service provider - Bên cung cấp dịch vụ quản lý) loại bỏ các tác vụ quản lý và kiểm soát khỏi người dùng cuối, điều này có thể loại bỏ quyền kiểm soát control plane khỏi người dùng cuối, đây là điều sẽ xảy ra với Amazon EKS, Microsoft AKS và Google Kubernetes Engine (GKE).

### Quá nhiều chọn lựa

Ý tôi là được lựa chọn là điều tuyệt vời nhưng có một điểm khiến chúng ta bị ngập trong các chọn lựa và bài viết này không đề cập một cách sâu sắc về tất cả các tuỳ chọn trong mỗi danh mục được liệt kê ở trên. Ngoài những điều được nhắc tới ở trên, chúng ta cũng có OpenShift của RedHat và lựa chọn này có thể được chạy trên các lựa chọn ở trên với tất cả các nhà cung cấp điện toán đám mây lớn và có lẽ sẽ mang lại khả năng sử dụng tổng thể tốt nhất cho các quản trị viên bất kể cluster được triển khai ở đâu.

Vậy bạn bắt đầu từ đâu trong hành trình học tập của mình, như tôi đã nói, tôi bắt đầu với lộ trình ảo hoá nhưng đó là vì tôi có quyền truy cập vào một máy chủ vậy lý mà tôi có thể sử dụng cho mục đích này, tôi rất biết ơn vì điều đó và trên thực tế, kể từ đó tôi không còn được truy cập vào lựa chọn này.

Lời khuyên thực tế của tôi bây giờ là sử dụng Minikube làm tuỳ chọn đầu tiên hoặc Kind (Kubernetes in Docker) nhưng Minikube mang lại cho chúng ta một số tiện ích bổ sung và gần như trừu tượng hoá sự phức tạp vì chúng ta có thể sử dụng các add-ón và xây dựng mọi thứ một cách nhanh chóng, sau đó xoá tất cả khi chúng ta hoàn thành, chúng ta cũng có thể chạy nhiều clusters, chạy nó ở hầu hết tất cả mọi nơi, đa nền tảng và không phụ thuộc và phần cứng. 

Tôi đã trài qua hành trình tìm hiểu về Kubernetes, vì vậy tôi sẽ để lại lựa chọn nền tảng và các chi tiết cụ thể ở đây để liệt kê các tuỳ chọn mà tôi đã thử để giúp tôi hiểu rõ hơn về nền tàng Kubernetes và nơi nó có thể chạy. Bạn có thể tham khảo những bài viết dưới đây để đưa ra lựa chọn của mình.

- [Kubernetes playground – How to choose your platform](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-1)
- [Kubernetes playground – Setting up your cluster](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-2)
- [Getting started with Amazon Elastic Kubernetes Service (Amazon EKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-amazon-elastic-kubernetes-service-amazon-eks)
- [Getting started with Microsoft Azure Kubernetes Service (AKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-azure-kubernetes-service-aks)
- [Getting Started with Microsoft AKS – Azure PowerShell Edition](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-aks-azure-powershell-edition)
- [Getting started with Google Kubernetes Service (GKE)](https://vzilla.co.uk/vzilla-blog/getting-started-with-google-kubernetes-service-gke)
- [Kubernetes, How to – AWS Bottlerocket + Amazon EKS](https://vzilla.co.uk/vzilla-blog/kubernetes-how-to-aws-bottlerocket-amazon-eks)
- [Getting started with CIVO Cloud](https://vzilla.co.uk/vzilla-blog/getting-started-with-civo-cloud)
- [Minikube - Kubernetes Demo Environment For Everyone](https://vzilla.co.uk/vzilla-blog/project_pace-kasten-k10-demo-environment-for-everyone)

### Những gì chúng ta sẽ đề cập trong loạt bài về Kubernetes

- Kiến trúc Kubernetes
- Các câu lệnh kubectl 
- Kubernetes YAML
- Kubernetes Ingress
- Kubernetes Services
- Helm Package Manager
- Lưu trữ liên tục - Persistent Storage
- Ứng dụng có trạng thái - Stateful Apps

## Tài liệu tham khảo

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

Hẹn gặp lại vào [ngày 51](day51.md)
