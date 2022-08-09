---
title: '#90DaysOfDevOps - 全体像: DevOpsとプログラミング言語の学習 - 7日目'
published: false
description: 90DaysOfDevOps - 体像: DevOpsとプログラミング言語の学習
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048856
---
## 体像: DevOpsとプログラミング言語の学習

DevOpsエンジニアとして長期的に成功するためには、少なくとも1つのプログラミング言語を基礎的なレベルで知っておく必要がある、と言ってもいいと思います。最初のセッションでは、なぜこのような重要なスキルが必要なのかを探ります。この週やセクションが終わる頃には、なぜ、どのように、何をすれば学習を進めることができるのかについて、より深く理解できるようになっていることを願っています。

もし私がソーシャル上で「DevOps関連の職務にプログラミングスキルは必要ですか」と尋ねたら、その答えはおそらく「イエス」でしょう？そうでない場合は、私に教えてください。しかし、もっと大きな疑問があり、ここで明確な答えが得られないのが、どのプログラミング言語かということです。 ここで私が見た最も一般的な答えはPythonで、最近ではGolangやGoを学ぶべきという意見も増えています。

DevOpsで成功するためには、プログラミングスキルに関する十分な知識が必要だというのが、少なくとも私が得た教訓です。しかし、正しい道を選択するためには、なぜそれが必要なのかを理解する必要があります。

## なぜプログラミング言語を学ぶ必要があるのかを理解する

PythonとGoがDevOpsエンジニアに頻繁に推奨される理由は、DevOpsツールの多くがPythonかGoで書かれているからです。このことは、あなたが何を学ぶべきかを決定し、最も有益である可能性が高いので、重要なことです。もしあなたがDevOpsツールを作ろうとしている、あるいは作っているチームに参加しようとしているのなら、同じ言語を学ぶのは理にかなっていますし、もしあなたがKubernetesやコンテナに深く関わっているのなら、プログラミング言語としてGoを選びたいと思うかもしれませんね。私の場合、私が働いている会社（Kasten by Veeam）は、Kubernetesのデータ管理に焦点を当てたCloud-Nativeエコシステムで、すべてがGoで記述されています。

しかし、そのような明確な理由がない場合、学生であったり、キャリアの転換期であったりして、本当の意味での決断ができないかもしれません。そんなときは、自分がやりたいアプリケーションに共鳴し、フィットしそうなものを選べばいいと思います。

私はソフトウェア開発者になりたいわけではなく、プログラミング言語についてもう少し理解したいだけなのです。

また、Kasten K10やTerraform、HCLといったDevOpsツールとどのように連携するのかも重要です。これは設定ファイルと呼ばれるもので、DevOpsツールとどのようにやり取りして物事を実現するかというものです。(このセクションの最終日にYAMLについて少し掘り下げます。）

## プログラミング言語の習得を断念したのは、私自身のせい？

ほとんどの場合、あるいは職務によっては、エンジニアリングチームがワークフローにDevOpsを導入するのを支援し、アプリケーションに関する多くのテストを行い、構築されたワークフローが最初の数日間で述べたDevOpsの原則に合致していることを確認することになります。しかし実際には、アプリケーションのパフォーマンス問題などのトラブルシューティングを行うことが多いようです。これは、私が知っておくべきプログラミング言語は、コードが書かれている言語なのか、という最初の論点と理由に戻ります。アプリケーションがNodeJSで書かれている場合、GoやPythonのバッジを持っていても、あまり役に立ちません。

## なぜGoを習得するのか？

DevOpsのための次のプログラミング言語がGolangである理由近年、Goは非常に人気のあるプログラミング言語になっています。2021年のStackOverflow Surveyによると、Goは最も欲しいプログラミング、スクリプト、マークアップ言語の第4位で、Pythonがトップですが、私の話を聞いてください。[StackOverflow 2021 Developer Survey - Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)

また、Kubernetes、Docker、Grafana、Prometheusなど、最も有名なDevOpsツールやプラットフォームがGoで書かれていることも紹介した通りです。

DevOpsに最適なGoの特徴にはどのようなものがあるのでしょうか。

## Goプログラムのビルドとデプロイメント
Pythonのようなインタプリタ言語をDevOpsの役割で使用する利点は、Pythonプログラムを実行する前にコンパイルする必要がないことです。特に小規模な自動化タスクでは、コンパイルが必要なビルドプロセスで時間がかかるのは避けたいものです。 Goはコンパイルが速いことでも知られています。

## DevOpsにおけるGoとPythonの比較

Goプログラムは静的にリンクされています。つまり、Goプログラムをコンパイルすると、すべてが単一のバイナリ実行ファイルに含まれ、リモートマシンにインストールする必要のある外部依存は必要ありません。

Goはプラットフォームに依存しない言語です。つまり、Linux、Windows、macOSなど*すべてのオペレーティングシステム用のバイナリ実行ファイルを作成することができ、非常に簡単に実行することができます。Pythonでは、特定のオペレーティングシステム用のバイナリ実行ファイルを作成するのはそれほど簡単ではありません。

Goは非常にパフォーマンスの高い言語であり、特にPythonと比較してCPUやメモリなどのリソース使用量が少なく、高速なコンパイルと高速な実行が可能で、Go言語には数多くの最適化が実装されているため、非常にパフォーマンスが高い。(下記リソース) 

Pythonでは、特定のPythonプログラムを実装するためにサードパーティのライブラリを使用する必要がありますが、Goには標準ライブラリがあり、DevOpsに必要な機能の大部分が直接組み込まれています。これには、ファイル処理、HTTPウェブサービス、JSON処理、並行処理と並列処理のネイティブサポート、組み込みのテストなどの機能が含まれます。

これは決してPythonを投げつけるわけではなく、私がGoを選んだ理由を述べているのですが、上記のGo vs Pythonということではなく、私が働いている会社がGoでソフトウェアを開発しているので、一般的に理にかなっているから、というのが理由です。

この章をまだ読み終えていないのですが、最初のプログラミング言語を習得すると、他の言語への挑戦が容易になると言われています。おそらく、どこの会社でも、JavaScriptやNode JSアプリケーションの管理、アーキテクト、オーケストレーション、デバッグに携わらない仕事はないでしょう。

## リソース

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s) 
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I) 
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals) 
- [FreeCodeCamp -  Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s) 
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N) 

さて、このトピックの次の6日間、私の意図は、上記のリソースのいくつかに取り組んで、毎日の私のノートを文書化することです。もし、あなたが時間があれば、先に進み、時間が許す限り、それぞれを学習してください。

それでは、【8日目】(day08.md)でお会いしましょう。
