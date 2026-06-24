---
title: '#90DaysOfDevOps - Panorama geral: aprendendo uma linguagem de programação - Day 7'
published: false
description: 90DaysOfDevOps - Panorama geral: aprendendo uma linguagem de programação
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048856
---

## Panorama geral: DevOps e aprendizado de uma linguagem de programação

Acho justo dizer que para ter sucesso no longo prazo como engenheiro de DevOps você precisa conhecer pelo menos uma linguagem de programação em um nível básico. Quero aproveitar esta primeira sessão desta seção para explorar por que essa é uma habilidade tão importante de se ter e, esperançosamente, até o final desta semana ou seção, você terá uma melhor compreensão do porquê, como e o que fazer. fazer para progredir em sua jornada de aprendizado.

Acho que se eu perguntasse nas redes sociais, você precisa ter habilidades de programação para funções relacionadas a DevOps, a resposta provavelmente seria um duro sim? Deixe-me saber se você pensa o contrário? Ok, mas então uma questão maior e é aqui que você não obterá uma resposta tão clara sobre qual linguagem de programação? A resposta mais comum que vi aqui foi Python ou cada vez mais, vemos que Golang ou Go deveria ser a linguagem que você aprenderá.

Para ter sucesso no DevOps, você precisa ter um bom conhecimento de habilidades de programação, pelo menos o que concluí. Mas temos que entender por que precisamos disso para escolher o caminho certo.

## Entenda por que você precisa aprender uma linguagem de programação.

A razão pela qual Python e Go são recomendados com tanta frequência para engenheiros de DevOps é que muitas das ferramentas de DevOps são escritas em Python ou Go, o que faz sentido se você pretende criar ferramentas de DevOps. Agora, isso é importante porque determinará realmente o que você deve aprender e o que provavelmente seria mais benéfico. Se você pretende construir ferramentas DevOps ou está se juntando a uma equipe que o faz, então faria sentido aprender a mesma linguagem. Se você estiver fortemente envolvido em Kubernetes ou Containers, é mais do que provável que você queira escolha Go como sua linguagem de programação. Para mim, a empresa em que trabalho (Kasten by Veeam) está no ecossistema Cloud-Native focado em gerenciamento de dados para Kubernetes e tudo é escrito em Go.

Mas então você pode não ter um raciocínio claro como esse para escolher se você pode ser um estudante ou estar em transição de carreira sem nenhuma decisão real tomada por você. Acho que nesta situação você deve escolher aquele que parece ressoar e se adequar aos aplicativos com os quais deseja trabalhar.

Lembre-se de que não pretendo me tornar um desenvolvedor de software aqui, só quero entender um pouco mais sobre a linguagem de programação para poder ler e entender o que essas ferramentas estão fazendo e, então, isso possivelmente nos levará a como podemos ajudar a melhorar as coisas.

Também é importante saber como você interage com as ferramentas DevOps que podem ser Kasten K10 ou Terraform e HCL. Isso é o que chamaremos de arquivos de configuração e é assim que você interage com essas ferramentas DevOps para fazer as coisas acontecerem, geralmente serão YAML. (Podemos usar o último dia desta seção para mergulhar um pouco no YAML)

## Acabei de me convencer a não aprender uma linguagem de programação?

Na maioria das vezes ou dependendo da função, você ajudará as equipes de engenharia a implementar DevOps em seu fluxo de trabalho, fazendo muitos testes em torno do aplicativo e garantindo que o fluxo de trabalho criado esteja alinhado aos princípios de DevOps que mencionamos nos primeiros dias. . Mas, na realidade, muitas vezes será necessário solucionar um problema de desempenho do aplicativo ou algo parecido. Isso volta ao meu argumento e raciocínio originais: a linguagem de programação que preciso saber é aquela em que o código está escrito? Se o aplicativo for escrito em NodeJS, não ajudará muito se você tiver um emblema Go ou Python.

## Por que Go?

Por que Golang é a próxima linguagem de programação para DevOps? Go se tornou uma linguagem de programação muito popular nos últimos anos. De acordo com a pesquisa StackOverflow para 2021, Go ficou em quarto lugar nas linguagens de programação, script e marcação mais procuradas, com Python sendo o principal, mas me escute. [Pesquisa de desenvolvedores StackOverflow 2021 – link mais procurado](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)

Como também mencionei, algumas das ferramentas e plataformas DevOps mais conhecidas são escritas em Go, como Kubernetes, Docker, Grafana e Prometheus.

Quais são algumas das características do Go que o tornam excelente para DevOps?

## Construção e implantação de programas Go

Uma vantagem de usar uma linguagem como Python interpretada em uma função DevOps é que você não precisa compilar um programa python antes de executá-lo. Especialmente para tarefas de automação menores, você não quer ser retardado por um processo de construção que requer compilação, embora Go seja uma linguagem de programação compilada, **Go compila diretamente em código de máquina**. Go também é conhecido por tempos de compilação rápidos.

## Go vs Python para DevOps

Os programas Go são vinculados estaticamente, isso significa que quando você compila um programa go, tudo é incluído em um único executável binário e nenhuma dependência externa será necessária para ser instalada na máquina remota, o que facilita a implantação de programas go, em comparação com o programa python que usa bibliotecas externas, você deve ter certeza de que todas essas bibliotecas estão instaladas na máquina remota na qual deseja executar.

Go é uma linguagem independente de plataforma, o que significa que você pode produzir executáveis binários para todos os sistemas operacionais, Linux, Windows, macOS etc. e é muito fácil de fazer. Com Python, não é tão fácil criar esses executáveis binários para sistemas operacionais específicos.

Go é uma linguagem de muito desempenho, tem compilação rápida e tempo de execução rápido com menor uso de recursos como CPU e memória, especialmente em comparação com python, inúmeras otimizações foram implementadas na linguagem Go que a torna tão eficiente. (Recursos abaixo)

Ao contrário do Python, que geralmente requer o uso de bibliotecas de terceiros para implementar um programa Python específico, go inclui uma biblioteca padrão que possui a maioria das funcionalidades necessárias para DevOps integradas diretamente nela. Isso inclui processamento de arquivos de funcionalidade, serviços web HTTP, processamento JSON, suporte nativo para simultaneidade e paralelismo, bem como testes integrados.

Isso não é de forma alguma jogar o Python debaixo do ônibus. Estou apenas dando meus motivos para escolher Go, mas eles não são os Go vs Python acima, geralmente porque faz sentido, já que a empresa para a qual trabalho desenvolve software em Go, e é por isso.

Direi que, uma vez que você tenha feito isso, ou pelo menos me disseram que não estou lendo muitas páginas deste capítulo agora, é que depois que você aprende sua primeira linguagem de programação, fica mais fácil aprender outras linguagens. Você provavelmente nunca terá um único emprego em qualquer empresa em qualquer lugar onde não precise lidar com gerenciamento, arquitetura, orquestração e depuração de aplicativos JavaScript e Node JS.

## Resources

- [Pesquisa de desenvolvedores StackOverflow 2021](https://insights.stackoverflow.com/survey/2021)
- [Por que estamos escolhendo Golang para aprender](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Aprenda em 12 minutos](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Curso completo de Golang - 3 horas e 24 minutos](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**PAGO** Nigel Poulton Pluralsight - Fundamentos do Go - 3 horas e 26 minutos](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Aprenda Programação Go - Tutorial Golang para Iniciantes](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - lista de reprodução completa](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Agora, nos próximos 6 dias deste tópico, pretendo trabalhar com alguns dos recursos listados acima e documentar minhas anotações de cada dia. Você notará que geralmente duram cerca de 3 horas como um curso completo. Queria compartilhar minha lista completa para que, se você tiver tempo, siga em frente e trabalhe em cada um se o tempo permitir, vou me ater à minha hora de aprendizado cada dia.

Vejo você no [Dia 8](day08.md).
