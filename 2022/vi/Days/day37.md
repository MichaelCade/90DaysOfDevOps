---
title: '#90DaysOfDevOps - Giới thiệu về Git - Ngày 37'
published: false
description: 90DaysOfDevOps - Giới thiệu về Git
tags: 'DevOps, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048707
---

## Giới thiệu về git

Trong hai bàì đăng trước, chúng ta đã tìm hiệu về quản lý phiên bản và một số quy trình, workflow của git với tư cách là một hệ thống quản lý phiên bản vào [ngày 35](day35.md) Sau đó chúng ta đã cài đặt git trên hệ thống, cập nhật và cấu hình. Chúng ta cũng đã đi sâu hơn một chút và điểm khác nhau cơ bản trong hệ thống quản lý phiên bản Client-Server và hệ thống quản lý phiên bản phân tán ví dụ như Git vào [ngày 36](day36.md).

Bây giờ, chúng ta sẽ tìm hiểu các câu lệnh cơ bản và trường hợp sử dụng chúng với git.

### Trợ giúp với git

Sẽ có lúc bạn không thể nhớ hoặc không biết lệnh cần sử dụng để hoàn thành công việc với git. Bạn sẽ cần giúp đỡ.

Google hoặc bất kỳ công cụ tìm kiểm nào có thể là điểm đến đầu tiên khi bạn tìm kiếm sự trờ giúp.

Tiếp sau đó có thể là trang chính thức của git và tài liệu. [git-scm.com/docs](http://git-scm.com/docs) Tại đây, bạn sẽ không những chỉ tìm thấy tài liệu tham khảo tốt cho tất cả các câu lệnh, mà còn có rất nhiều các tài nguyên khác.

![](../../Days/Images/Day37_Git1.png)

Chúng ta cũng có thể truy cập tài liệu tương tự sau, điều này cực kỳ hữu ích nếu bạn không có kết nối nào từ terminal. Ví dụ: nếu chúng ta sử dụng lệnh `git add`, chúng ta có thể chạy `git add --help`  và đọc hướng dẫn dưới đây.

![](../../Days/Images/Day37_Git2.png)

Chúng ta cũng có thể dụng `git add -h` để cung cấp tống hợp các tuỳ chọn có sẵn mà chúng ta có thể sử dụng.

![](../../Days/Images/Day37_Git3.png)

### Những câu chuyện xung quanh Git

"Git không có quyền kiểm soát truy cập" - bạn có thể trao quyền cho một leader để duy trì mã nguồn.

"Git quá nặng" - Git có thể cung cấp các kho lưu trữ nông (shallow repositories), có nghĩa là lịch sử git sẽ được cắt giảm nếu bạn có một project lớn.

### Những thiếu sót thực tế

Không phải là phương pháp lý tưởng cho các tệp nhị phân. Tuyệt vời cho mã nguồn nhưng không tuyệt vời cho các tệp thực thi hoặc video chẳng hạn.

Git không thân thiện với người dùng, việc chúng ta phải dành thời gian để nói và tìm hiểu các lệnh và chứng năng của công cụ có lẽ là minh chứng rõ ràng nhất cho việc đó.

Nhìn chung, git khó học nhưng dễ sử dụng.

### Hệ sinh thái git

Tôi muốn trình bày ngắn gọn về hệ sinh thái xung quanh git nhưng không đi quá sâu và việc đó. Điều quan trọng là phải nhắc đến ở góc độ high-level.

Hầu như tất cả các công cụ phát triển hiện đại đều hỗ trợ git.

- Công cụ phát triển - chúng ta đã đề cập tới Visual Studio Code, bạn sẽ tìm thấy các plugin và các tích hợp của git trong Sublime text và các trình soạn thảo văn bản và IDE khác. 
- Công cụ nhóm - Cũng được đề cập trong các công cụ như Jenkins theo góc nhìn CI/CD, Slack từ góc độ dịch vụ nhắn tin và Jira từ góc độ quản lý dự án và theo dõi vấn đề.

- Cloud providers - Tất cả các nhà cung cấp điện toán đám mây lớn đều hỗ trợ git, Microsoft Azure, Amazon AWS và Google Cloud Platform.
- Các dịch vụ dựa trên Git - Sau đó, chúng ta có GitHub, GitLab và BitBucket mà chúng ta sẽ đề cập rõ hơn ở phần sau. Các dịch vụ này có thể coi như là một mạng xã hội dành cho mã nguồn.

### Git Cheatsheet

Chúng ta chưa đề cập đến hầu hết các lệnh này nhưng sau khi xem xét một số cheatsheet có sẵ, tôi cũng muốn ghi lại một số câu lệnh git và mục đích sử dụng của chúng. Chúng ta không cần phải nhớ tất cả những câu lệnh này, với việc thực hành và sử dụng nhiều hơn, bạn sẽ học được một số điều cơ bản về git.

Tôi đã lấy những câu lệnh từ [atlassian](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet) nhưng viết lại chúng và đọc mô tả là một cách hay để biết các lệnh làm gì cũng như bắt đầu thực hành trong công việc hàng ngày.

### Cơ bản về Git

| Lệnh       | Ví dụ                     | Mô tả                                                                                                                 |
| ------------- | --------------------------- | --------------------------------------------------------------------------------------------------------------------------- |
| git init      | `git init <directory>`      | Tạo một git repository tron thư mục được chỉ định                                                                  |
| git clone     | `git clone <repo>`          | Sao chép repository tại <repo> vào máy local                                                                      |
| git config    | `git config user.name`      | Khai báo tên người sử dụng cho tất cả các commit ở repository hiện tại, có các flag `system`, `global`, `local` để tuỳ chọn. |
| git add       | `git add <directory>`       | Stage tất cả các thay đổi trong <directory> để chuẩn bị cho commit tiếp theo. Chúng ta có thể thêm <files> và <.> cho tất cả mọi thứ.                       |
| git commit -m | `git commit -m "<message>"` | Commit các file đã được staged, sử dụng <message> để làm rõ thay đổi được commit là gì.                                           |
| git status    | `git status`                | Liệt kê các tệp được  staged, unstaged và untracked.                                                                         |
| git log       | `git log`                   | Hiển thị tất cả lịch sử commit bằng định dạng mặc định. Có các tuỳ chọn bổ sung với lệnh này.
| git diff      | `git diff`                  | Hiển thị các thay đổi của các thay đổi chưa được stage.                                                             |

### Git Hoàn tác thay đổi

| Lệnh       | Ví dụ                     | Mô tả                                                                                                                 |
| ---------- | --------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| git revert | `git revert <commit>` | Tạo một commit mới hoàn tác lại tất cả các thay đổi trong <commit> trên branch hiện tại.                              |
| git reset  | `git reset <file>`    | Xoá <file> khỏi stage, nhưng không thay đổi thư mục làm việc. Điều này unstaged một tệp nhưng không ghi đè bất cứ thay đổi nào. |
| git clean  | `git clean -n`        | Hiển thị tệp nào sẽ bị xoá khỏi thư mục làm việc. Sử dụng `-f` thay cho `-n` để thực hiện việc xoá.                        |

### Git viết lại lịch sử

| Lệnh       | Ví dụ                     | Mô tả                                                                                                                 |
| ---------- | -------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| git commit | `git commit --amend` | Thay thế commit gần nhất bằng các thay đỏi đã được staged và ghép nó với commit cuối cùng. Sử dụng trong trường hợp chưa có gì được stage để thay đổi nội dung commit. |
| git rebase | `git rebase <base>`  | Rebase nhanh hiện tại và trong  <base>. <base> có thể là một commit ID, tên branch, tag hoặc một tham chiếu tương đối đến HEAD.                   |
| git reflog | `git reflog`         | Hiện thị log của các thay đổi với HEAD của kho lưu trữ cục bộ (local repository). Thêm --relative-date hiển thị ngày tháng và --all để hiện thị tất cả các refs.              |

### Git Phân nhánh (Branch)

| Lệnh       | Ví dụ                     | Mô tả                                                                                                                 |
| ------------ | -------------------------- | ------------------------------------------------------------------------------------------------------------- |
| git branch   | `git branch`               | Liệt kê tất cả các branch trong repo của bạn. Thêm đối số <branch> để tạo một branch với với tên <branch>. |
| git checkout | `git checkout -b <branch>` | Tạo và checkout sang một branch mới tên <branch>. Bỏ -b flag để checkout sang một branch đã có sẵn.            |
| git merge    | `git merge <branch>`       | Merge <branch> và branch hiện tại.                                                                       |

### Git Remote Repositories

| Lệnh       | Ví dụ                     | Mô tả                                                                                                                 |
| -------------- | ----------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| git remote add | `git remote add <name> <url>` | Tạo một kết nối mới đến một remote repo. Sau khi thêm một remote, bạn có thể dùng <name> thay cho <url> trong các lệnh của bạn.      |
| git fetch      | `git fetch <remote> <branch>` | Kéo một nhánh cụ thể có tên <branch>, từ repo. Bỏ <branch> để kéo tất cả các remote refs.                                            |
| git pull       | `git pull <remote>`           | Kéo một bản sau của nhánh hiện tại từ remote repo và ngay lập tức merge nó vào bản sao local.                                   |
| git push       | `git push <remote> <branch>`  | Đẩy nhánh tới <remote>, cùng với các commits và objects. Tạo nhánh có thên trong remote repo nếu nó không tồn tại. |

### Git Diff

| Lệnh       | Ví dụ                     | Mô tả                                                                                                                 |
| ----------------- | ------------------- | ---------------------------------------------------------------------- |
| git diff HEAD     | `git diff HEAD`     | Hiển thị sự khác biệt giữa thư mục làm việc và commit cuối cùng. |
| git diff --cached | `git diff --cached` | Hiện thị sự khác biệt giữa các thay đổi đã được staged và commit cuối cùng                 |

### Git Config

| Lệnh       | Ví dụ                     | Mô tả                                                                                                                 |
| ---------------------------------------------------- | ------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------- |
| git config --global user.name <name>                 | `git config --global user.name <name>`                 | Khai báo tên người dùng sẽ được sử dụng cho tất cả các commit ủa người dùng hiện tại.                                                                        |
| git config --global user.email <email>               | `git config --global user.email <email>`               | Khai báo email sẽ được sử dụng cho tất cả các commit của người dùng hiện tại.                                                                           |
| git config --global alias <alias-name> <git-command> | `git config --global alias <alias-name> <git-command>` | Tạo một phím tắt cho lệnh git .                                                                                                           |
| git config --system core.editor <editor>             | `git config --system core.editor <editor>`             | Cài đặt trình soạn thảo văn bản sẽ được sử dụng cho tất cả người dùng trên máy. <editor> phải là lệnh khởi chạy trình chỉnh sửa mong muốn |
| git config --global --edit                           | `git config --global --edit `                          | Mở tệp cấu hình chung trong trình soạn thảo văn bản để chỉnh sửa thủ công.                                                                       |

### Git Rebase

| Command              | Example                | Description                                                                                                                                 |
| -------------------- | ---------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| git rebase -i <base> | `git rebase -i <base>` | Rebase nhánh hiện tại vào <base>. Khởi chạy trình chỉnh sửa cho phép chỉnh sửa các commit khi chuyển qua base mới. |

### Git Pull

| Lệnh | Ví dụ | Mô tả |
| -------------------------- | ---------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| git pull --rebase <remote> | `git pull --rebase <remote>` | Kéo các bản sao của nhánh hiện tại trên remote và rebasse vào bản sao tại local. Sử dụng git rebase thay cho merge để gộp các nhánh. |

### Git Reset

| Lệnh | Ví dụ | Mô tả |
| ------------------------- | --------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| git reset                 | `git reset `                | Reset khu vực staging về commit gần nhất nhưng giữ nguyên thư mục làm việc.                                             |
| git reset --hard          | `git reset --hard`          | Reset khu vực staging về commit gần nhất và ghi đè lên tất cả các thay đổi trong thư mục làm việc.                      |
| git reset <commit>        | `git reset <commit>`        | Di chuyển branch hiện tại về <commit>, reset khu vực staging reset về commit đó nhưng giữ nguyên khu vực làm việc.                     |
| git reset --hard <commit> | `git reset --hard <commit>` | Giống với lệnh trước đó nhưng reset cả khu vực staging, thư mục làm việc. Xoá bỏ các thay đổi chưa được commit và tất cả các commit phía sau <commit>. |

### Git Push

| Lệnh | Ví dụ | Mô tả |
| ------------------------- | --------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| git push <remote> --force | `git push <remote> --force` | Ép buộc git đảy lên ngay cả khi kết quả không phải là một non-fast-forward merge. Không sử dụng --force trừ khi bạn biết chắc chắn mình đang làm gì.  |
| git push <remote> --all   | `git push <remote> --all`   | Đẩy tất cả các nhánh ở local đến một remote xác định.                                                                                        |
| git push <remote> --tags  | `git push <remote> --tags`  | Tage không được tự động đẩy lên khi bạn đẩy một nhánh hay sử dụng --all. --tags sẽ gửi tất cả những local tags lên remote repo. |

## Tài liệu tham khảo

- [What is Version Control?](https://www.youtube.com/watch?v=Yc8sCSeMhi4)
- [Types of Version Control System](https://www.youtube.com/watch?v=kr62e_n6QuQ)
- [Git Tutorial for Beginners](https://www.youtube.com/watch?v=8JJ101D3knE&t=52s)
- [Git for Professionals Tutorial](https://www.youtube.com/watch?v=Uszj_k0DGsg)
- [Git and GitHub for Beginners - Crash Course](https://www.youtube.com/watch?v=RGOj5yH7evk&t=8s)
- [Complete Git and GitHub Tutorial](https://www.youtube.com/watch?v=apGV9Kg7ics)
- [Git cheatsheet](https://www.atlassian.com/git/tutorials/atlassian-git-cheatsheet)

Hẹn gặp lại vào [ngày 38](day38.md)
