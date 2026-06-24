---
title: '#90DaysOfDevOps - ラボの構築 - 26日目'
published: false
description: 90DaysOfDevOps - ラボの構築
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048762
---
## ラボの構築

EVE-NGを使用してエミュレートしたネットワークのセットアップを継続し、いくつかのデバイスをデプロイして、これらのデバイスの設定を自動化する方法について考え始めたいと考えています。[Day 25](day25.md) では、VMware Workstation を使用してマシンに EVE-NG をインストールする方法について説明しました。

### EVE-NG クライアントのインストール

デバイスにSSHするときにどのアプリケーションを使うかを選択できるクライアントパックもあります。また、リンク間のパケットキャプチャのためにWiresharkをセットアップします。お使いのOS（Windows、macOS、Linux）に対応したクライアントパックを入手することができます。

[EVE-NG Client Download](https://www.eve-ng.net/index.php/download/)

![](Images/Day26_Networking1.png)

クイックヒント: クライアントとしてLinuxを使用している場合、この[クライアントパック](https://github.com/SmartFinn/eve-ng-integration)があります。

インストールは簡単で、デフォルトのままにしておくことをお勧めします。

### ネットワークイメージの取得

このステップは挑戦でした、私はいくつかのリソースへのリンクとルータとスイッチの画像をダウンロードしながら、それらをアップロードする方法と場所を教えてくれる最後にリンクするいくつかのビデオに従いました。

これは、私は教育目的のためにすべてを使用していることに注意することが重要です。私は、ネットワークベンダから公式の画像をダウンロードすることをお勧めします。

[Blog & Links to YouTube videos](https://loopedback.com/2019/11/15/setting-up-eve-ng-for-ccna-ccnp-ccie-level-studies-includes-multiple-vendor-node-support-an-absolutely-amazing-study-tool-to-check-out-asap/) 

[How To Add Cisco VIRL vIOS image to Eve-ng](https://networkhunt.com/how-to-add-cisco-virl-vios-image-to-eve-ng/)

全体的にこのステップは少し複雑で、もっと簡単にできるかもしれませんが、上記のブログとビデオはEVE-NGボックスにイメージを追加するプロセスを説明しています。

私はFileZillaを使用して、SFTP経由でqcow2をVMに転送しました。

このラボでは、Cisco vIOS L2（スイッチ）とCisco vIOS（ルーター）が必要です。

### ラボの作成

EVE-NGのWebインターフェイス内で、新しいネットワークトポロジーを作成します。4台のスイッチと、外部ネットワークへのゲートウェイとして機能する1台のルータを用意する予定です。

| Node        | IP Address  |
| ----------- | ----------- |
| Router      | 10.10.88.110|
| Switch1     | 10.10.88.111|
| Switch2     | 10.10.88.112|
| Switch3     | 10.10.88.113|
| Switch4     | 10.10.88.114|

#### EVE-NGにノードを追加する

EVE-NGに初めてログインすると、以下のような画面が表示されますので、まずは最初のラボを作成してみたいと思います。

![](Images/Day26_Networking2.png)

研究室の名前を入力します。その他の項目は任意です。

![](Images/Day26_Networking3.png)

すると、ネットワークの作成を開始するための空白のキャンバスが表示されます。キャンバス上で右クリックし、ノードの追加を選択します。

ここで、ノードのオプションの長いリストが表示されます。上記で説明した通りなら、下図の青い2つのノードがあり、他のノードはグレーで選択できないようになっています。

![](Images/Day26_Networking4.png)

ラボに以下を追加したい:

- 1 x Cisco vIOS Router
- 4 x Cisco vIOS Switch

簡単なウィザードを実行し、ラボに追加すると次のようになります。

![](Images/Day26_Networking5.png)

#### ノードの接続

次に、ルータとスイッチの間の接続を追加する必要があります。デバイスの上にカーソルを置くと、以下のような接続アイコンが表示されるので、それを接続したいデバイスに接続すれば、簡単に接続することができます。

![](Images/Day26_Networking6.png)

環境の接続が完了したら、右クリックメニューにあるボックスや円を使用して、物理的な境界や位置を定義する方法を追加することもできます。また、テキストを追加することも可能で、ラボのネーミングやIPアドレスを定義するのに便利です。

私は、以下のようなラボを作りました。

![](Images/Day26_Networking7.png)

また、上のラボはすべて電源が切れていることに気づきます。すべてを選択して右クリックし、選択した状態で起動することで、ラボを開始することができます。

![](Images/Day26_Networking8.png)

このラボを立ち上げると、各デバイスにコンソールできるようになります。この段階では、何も設定されておらず、かなり間抜けな状態であることに気づくでしょう。各端末の設定をコピーするか、自分で作成することで、各ノードに設定を追加することができます。

私は参考のため、リポジトリの Networking フォルダに設定を残しておきます。

| Node        | Configuration         |
| ----------- | -----------           |
| Router      | [R1](Networking/R1)   |
| Switch1     | [SW1](Networking/SW1) |
| Switch2     | [SW2](Networking/SW2) |
| Switch3     | [SW3](Networking/SW3) |
| Switch4     | [SW4](Networking/SW4) |

## リソース

- [Free Course: Introduction to EVE-NG](https://www.youtube.com/watch?v=g6B0f_E0NMg)
- [EVE-NG - Creating your first lab](https://www.youtube.com/watch?v=9dPWARirtK8)
- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

私はネットワークエンジニアではないので、ここで使用している例のほとんどは、無料ではありませんが、Network Automationを理解するのに役立つシナリオのいくつかをこの本から引用しています。

- [Hands-On Enterprise Automation with Python (Book)](https://www.packtpub.com/product/hands-on-enterprise-automation-with-python/9781788998512)

[27日目](day27.md)でお会いしましょう。
