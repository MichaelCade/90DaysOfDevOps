---
title: '#90DaysOfDevOps - Thực hành với công cụ giám sát - Ngày 78'
published: false
description: 90DaysOfDevOps - Thực hành với công cụ giám sát
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049056
---

## Thực hành với công cụ giám sát

Hôm nay tôi sẽ nói về Prometheus, tôi đã thấy ngày càng nhiều hệ thống sử dụng Prometheus cho Cloud-Native nhưng nó cũng có thể được sử dụng để quản lý các tài nguyên vật lý cũng như Kubernetes và những thứ tương tự.

### Prometheus - Giám sát gần như mọi thứ

Trước hết, Prometheus là Mã nguồn mở có thể giúp bạn giám sát các containers và hệ thống micro-services cũng như các máy chủ vật lý, ảo hoá và các dịch vụ khác. Có một cộng đồng lớn đằng sau Prometheus.

Prometheus có rất nhiều [integrations and exporters](https://prometheus.io/docs/instrumenting/exporters/) Điều quan trọng là có thể export các metrics hiện có dưới dạng Prometheus metrics. Trên hết, nó còn hỗ trợ nhiều ngôn ngữ lập trình.

Phương thức pull - Nếu bạn đang nói chuyện với hàng nghìn micro-services hoặc hệ thống và dịch vụ thì các bạn sẽ thấy dịch vụ đẩy tới hệ thống giám sát bằng phương thức push. Điều này mang đến một số thách thức xung quanh việc làm quá tải hệ thống mạng, CPU cao và cũng xuất hiện single point of failure. Trong khi đó, phương thức pull mang đến cho chúng ta trải nghiệm tốt hơn nhiều, khi đó Prometheus sẽ lấy dữ liệu từ metrics endpoint trên các dịch vụ.

Một lần nữa chúng ta thấy YAML để cấu hình cho Prometheus.

![](Images/Day78_Monitoring7.png)

Sau này, bạn sẽ thấy điều này trông như thế nào khi được triển khai vào Kubernetes, cụ thể là chúng ta có **PushGateway** lấy metrics từ các jobs/và exporters.

Chúng ta có **AlertManager** giúp đưa ra cảnh báo và đây là nơi chúng ta có thể tích hợp vào các dịch vụ bên ngoài như email, Slack và các công cụ khác.

Sau đó, chúng ta có máy chủ Prometheus quản lý việc truy xuất các pull metrics đó từ PushGateway và gửi các push alerts đó đến AlertManager. Máy chủ Prometheus cũng lưu trữ dữ liệu trên đĩa cục bộ. Mặc dù có thể tận dụng các giải pháp lưu trữ từ xa.


Chúng ta cũng có PromQL, ngôn ngữ được sử dụng để tương tác với các metrics, bạn có thể thấy điều này trong Giao diện người dùng web của Prometheus nhưng ở phần sau trong bài viết này, bạn cũng sẽ thấy cách nó cũng được sử dụng trong các công cụ trực quan hóa dữ liệu như Grafana.

### Các cách triển khai Prometheus

Có nhiều cách cài đặt Prometheus khác nhau, [Trong phần tải xuống](https://prometheus.io/download/) cũng có sẵn docker image.

`docker run --name prometheus -d -p 127.0.0.1:9090:9090 prom/prometheus`

Nhưng chúng ta sẽ tập trung vào việc triển khai trên Kubernetes. Bản thân việc này cũng có một vài lựa chọn.

- Cấu hình bằng các tệp YAML 
- Sử dụng Operator (quản lý tất cả các thành phần của Kubernetes)
- Sử dụng helm chart để triển khai operator

### Triển khai trên Kubernetes

Chúng ta sẽ sử dụng lại minikube cluster cục bộ của mình để cài đặt một cách nhanh chóng và đơn giản. Giống như các tương tác trước đây với minikube, chúng ta sẽ sử dụng heml để triển khai Prometheus heml chart.

`helm repo add prometheus-community https://prometheus-community.github.io/helm-charts`

![](Images/Day78_Monitoring1.png)

Như bạn có thể thấy ở trên, chúng ta cũng đã chạy helm repo update, hiện tại chúng ta đã có thể triển khai Prometheus vào môi trường minikube của mình bằng cách sử dụng lệnh `helm install stable prometheus-community/prometheus`.

![](Images/Day78_Monitoring2.png)

Sau vài phút, bạn sẽ thấy một số pods mới xuất hiện, trong demo này, tôi đã triển khai vào namespace mặc định, nhưng thông thường tôi sẽ triển khai trong namespace của nó.

![](Images/Day78_Monitoring3.png)

Sau khi tất cả các pods đã chạy, chúng ta cũng có thể xem xét tất cả các thành phần đã được triển khai của Prometheus.

![](Images/Day78_Monitoring4.png)

Bây giờ để truy cập vào Giao diện người của dùng máy chủ Prometheus, chúng ta có thể sử dụng lệnh sau để chuyển tiếp cổng.

```Shell
export POD_NAME=$(kubectl get pods --namespace default -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace default port-forward $POD_NAME 9090
```

Khi chúng ta lần đầu tiên mở trình duyệt của mình và truy cập `http://localhost:9090`, chúng ta thấy màn hình trống sau đây.

![](Images/Day78_Monitoring5.png)

Vì chúng ta đã triển khai trên Kubernetes cluster nên sẽ được tự động chọn các metrics từ Kubernetes API để sử dụng PromQL nhằm đảm bảo ít nhất thì chúng ta đang thu tập các metrics `container_cpu_usage_seconds_total`

![](Images/Day78_Monitoring6.png)

Nói ngắn gọn về việc học PromQL và áp dụng nó trong thực tế, điều này khá giống như tôi đã đề cập trước đây trong việc có được các metrics là rất tốt và việc giám sát cũng vậy, nhưng bạn phải biết bạn đang giám sát những gì và lý do cho việc đó, cũng như những gì bạn không theo dõi và lý do cho việc đó!

Tôi muốn nhắc lại về Prometheus nhưng bây giờ, tôi nghĩ chúng ta cần xem xét tới việc Quản lý log và Trực quan hóa dữ liệu để có thể quay lại Prometheus ở các phần sau.

## Tài liệu tham khảo

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b)
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg)
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)

Hẹn gặp lại vào [ngày 79](day79.md)
