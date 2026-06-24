---
title: '#90DaysOfDevOps - DevOps -  本当の話 - 6日目'
published: false
description: 90DaysOfDevOps - DevOps - 本当の話
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048855
---
## DevOps - 本当の話

DevOpsは当初、NetflixやFortune 500のような環境や要件がないため、多くの人にとって手の届かないものと思われていましたが、今ではあらゆる種類のビジネスでDevOpsの実践を採用する際に、それが普通になってきているように思います。

文末のリソースの2番目のリンクから、さまざまな業界や業種がDevOpsを使用しており、ビジネス目標に大きなプラスの効果をもたらしていることがおわかりいただけるでしょう。

もちろん、ここでの包括的な利点は、DevOpsが正しく実行されれば、ソフトウェア開発のスピードと品質を向上させるのに役立つということです。

この日は、DevOpsの実践を採用した成功した企業を見て、そのリソースを共有したいと思います。あなたのビジネスでDevOps文化を採用しましたか？それは成功しましたか？

Netflixは非常に良いモデルであり、現在でも一般的に見られるものより進んでいるので、もう一度触れますが、成功している他の有名ブランドについても触れます。

## Amazon
2010年、Amazonは物理サーバーをAmazon Web Services（AWS）クラウドに移行し、非常に小さな単位で容量を増減させることでリソースを節約できるようになりました。また、このAWSクラウドは、Amazonの小売部門を運営しながら、莫大な収益を上げるようになったことも知っている。

Amazonは2011年に、開発者が好きな時に好きなサーバにコードをデプロイできる継続的なデプロイメントプロセスを採用した（下記のリソースによる。)これにより、アマゾンは平均11.6秒ごとに新しいソフトウェアを本番サーバにデプロイすることを実現しました。

## Netflix
Netflixを利用しない人はいないでしょう。明らかに巨大で高品質なストリーミングサービスであり、少なくとも個人的には素晴らしいユーザーエクスペリエンスを持っています。

なぜ、そのような素晴らしいユーザーエクスペリエンスが必要なのでしょうか？少なくとも私にとっては、不具合の記憶がないサービスを提供するためには、スピード、柔軟性、そして品質へのこだわりが必要です。

NetFlixの開発者は、ITオペレーションに依存することなく、コードの断片を自動的にデプロイ可能なWebイメージに構築することができます。イメージは更新されると、カスタムビルドされたウェブベースのプラットフォームを使って、Netflixのインフラに統合されます。

継続的なモニタリングにより、イメージのデプロイに失敗した場合、新しいイメージはロールバックされ、トラフィックは以前のバージョンにリルートされるようになっています。

Netflix がチーム内で実践している「やるべきこと」と「やってはいけないこと」については、次のような素晴らしい講演があります。

## Etsy
私たちの多くが、また多くの企業がそうであるように、遅くて辛いデプロイメントに本当に苦労していました。同じように、私たちも、サイロ化したチームや、連携がうまくいっていない会社で働くことを経験してきたかもしれません。

AmazonとNetflixの記事を読んだ限りでは、Etsyは2009年末頃に開発者が自分自身のコードをデプロイできるようにしたようですが、これは他の2社より早かったかもしれません。(面白い！）。

興味深いのは、開発者がデプロイメントに責任を感じると、アプリケーションのパフォーマンスやアップタイム、その他の目標にも責任を持つようになる、ということに気づいたということです。

学習文化はDevOpsの重要な部分であり、教訓が得られれば、失敗さえも成功になり得ます。(この引用が実際にどこから来たのかは分かりませんが、なんとなく納得です！）。

この他にも、DevOpsが大成功を収めた企業の中で、ゲームを変えたという話をいくつか追加しています。


## リソース

- [How Netflix Thinks of DevOps](https://www.youtube.com/watch?v=UTKIT6STSVM)
- [16 Popular DevOps Use Cases & Real Life Applications [2021]](https://www.upgrad.com/blog/devops-use-cases-applications/)
- [DevOps: The Amazon Story](https://www.youtube.com/watch?v=ZzLa0YEbGIY)
- [How Etsy makes DevOps work](https://www.networkworld.com/article/2886672/how-etsy-makes-devops-work.html)
- [Adopting DevOps @ Scale Lessons learned at Hertz, Kaiser Permanente and lBM](https://www.youtube.com/watch?v=gm18-gcgXRY)
- [Interplanetary DevOps at NASA JPL](https://www.usenix.org/conference/lisa16/technical-sessions/presentation/isla)
- [Target CIO explains how DevOps took root inside the retail giant](https://enterprisersproject.com/article/2017/1/target-cio-explains-how-devops-took-root-inside-retail-giant)

###  DevOpsに注目した最初の数日間を振り返る

- DevOpsは、開発と運用を組み合わせたもので、**開発**、*テスト**、*デプロイメント**、*運用**からなるアプリケーション開発ライフサイクル全体を1つのチームが管理できるようにするものです。

- DevOpsの主な焦点と目的は、開発ライフサイクルを短縮する一方で、ビジネス目標に密接に連携した特徴、修正、および機能を頻繁に提供することです。

- DevOpsは、ソフトウェアを確実かつ迅速に提供・開発するためのソフトウェア開発手法である。また、「Continuous Development, Testing, Deployment, Monitoring**（継続的開発、テスト、デプロイメント、モニタリング）」と表記されることもあります。

ここまでくれば、自分がやりたいことがここにあるのか、そうでないのかがわかるはずです。では、[7日目](day07.md)でお会いしましょう。

7日目はプログラミング言語に飛び込みます。私は開発者になることを目指しているわけではありませんが、開発者が何をしているのかを理解できるようになりたいです。

1週間でそれを達成できるでしょうか？しかし、7日間または7時間かけて何かを学べば、始めたときよりも多くのことを知ることができるようになります。
