---
title: '#90DaysOfDevOps - Responsabilidades de um engenheiro de DevOps - Dia 2'
published: false
description: 90DaysOfDevOps - Responsabilidades de um engenheiro de DevOps 
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048699
date: '2022-10-13T11:17:00Z'
---

## Responsabilidades de um engenheiro de DevOps

Espero que você esteja chegando aqui depois de analisar os recursos postados no [Dia1 de #90DiasDeDevOps](dia01.md)

Isso foi brevemente abordado no primeiro post, mas agora devemos nos aprofundar nesse conceito e entender que existem duas partes principais ao criar uma aplicação. Temos a parte de **Desenvolvimento** onde os desenvolvedores de software programam a aplicação e o testam. Em seguida, temos a parte de **Operações** em que a aplicação é implantada e mantida em um servidor.

## DevOps é o elo entre os dois

Para entender DevOps ou as tarefas que um engenheiro de DevOps estaria realizando, precisamos entender as ferramentas ou o processo e a visão geral delas e como elas se unem.

Tudo começa com a aplicação! Você verá que é tudo sobre a aplicação quando se trata de DevOps.

Os desenvolvedores criarão uma aplicação, isso pode ser feito com muitas pilhas de tecnologia diferentes e vamos deixar isso para a imaginação por enquanto, pois entraremos nisso mais tarde. Isso também pode envolver muitas linguagens de programação diferentes, ferramentas de construção, repositórios de código etc.

Como engenheiro de DevOps, você não estará programando a aplicação, mas ter uma boa compreensão dos conceitos de como um desenvolvedor trabalha e dos sistemas, ferramentas e processos que estão usando é a chave para o sucesso.

Em um nível muito alto, você precisará saber como o aplicativo está configurado para se comunicar com todos os seus serviços ou serviços de dados necessários e, em seguida, incluir um requisito de como isso pode ou deve ser testado.

A aplicação precisará ser implantada em algum lugar, vamos simplificar aqui e dizer que será um servidor, não importa onde, mas um servidor. Espera-se que isso seja acessado pelo cliente ou usuário final, dependendo da aplicação que foi criada.

Este servidor precisa ser executado em algum lugar, on-premises, em uma nuvem pública, sem servidor (Ok, eu fui longe demais, não cobriremos sem servidor, mas é uma opção e mais e mais empresas estão indo nessa direção). Alguém precisa criar e configurar esses servidores e prepará-los para a execução do aplicativo. Agora, esse elemento pode chegar a você como engenheiro DevOps para implantar e configurar esses servidores.

Esses servidores executam um sistema operacional e, em geral, isso será Linux, mas temos uma seção ou semana inteira onde abordamos alguns dos conhecimentos fundamentais que você deve obter aqui.

Também é provável que precisemos nos comunicar com outros serviços em nossa rede ou ambiente, portanto, também precisamos ter esse nível de conhecimento sobre rede e configuração, isso pode, em algum grau, também cair aos pés do engenheiro DevOps. Novamente, abordaremos isso com mais detalhes em uma seção dedicada, falando sobre todas as coisas DNS, DHCP, balanceamento de carga etc.

## Jack of all trades, Master of none
_Essa é uma figura de linguagem americana, para se referir a uma pessoa que tem habilidades em diversas atividades, ao invés de estar focada em apenas uma. A melhor tradução ao pé da letra para Português seria "A pessoa que faz tudo e não é especialista em nada"_

Eu vou ser duro nesse ponto, você não precisa ser um especialista em rede ou infraestrutura, você precisa de um conhecimento básico de como colocar as coisas em funcionamento e conversando entre si, da mesma forma que talvez ter um conhecimento básico de uma linguagem de programação sem precisar ser um desenvolvedor. No entanto, você pode estar entrando nisso como um especialista em uma área e isso é uma ótima base para se adaptar a outras áreas.

Você também provavelmente não assumirá o gerenciamento desses servidores ou da aplicação diariamente.

Temos falado sobre servidores, mas a probabilidade é que sua aplicação seja desenvolvida para ser executada como contêineres, que ainda rodam em um servidor na maioria das vezes, mas você também precisará entender não apenas de virtualização, infraestrutura de nuvem como serviço (IaaS ), mas também a conteinerização. O foco nestes 90 dias será mais voltado para os contêineres.

## Visão geral de alto nível

De um lado temos nossos desenvolvedores criando novos recursos e funcionalidades (assim como correções de bugs) para a aplicação.

Do outro lado, temos algum tipo de ambiente, infraestrutura ou servidores que são configurados e gerenciados para executar esta aplicação e se comunicar com todos os seus serviços necessários.

A grande questão é como obtemos esses recursos e correções de bugs em nossos produtos e os disponibilizamos para esses usuários finais?

Como lançamos a nova versão da aplicação? Esta é uma das principais tarefas de um engenheiro DevOps, e o importante aqui não é apenas descobrir como fazer isso uma vez, mas precisamos fazer isso continuamente e de maneira automatizada e eficiente, que também precisa incluir testes!

É aqui que vamos terminar este dia de aprendizado, espero que tenha sido útil. Nos próximos dias, vamos nos aprofundar um pouco mais em mais algumas áreas do DevOps e, em seguida, nas seções que se aprofundam nas ferramentas, nos processos e nos benefícios deles.

## Recursos

Estou sempre aberto a adicionar recursos adicionais a esses arquivos README, pois estão aqui como uma ferramenta de aprendizado.

Meu conselho é assistir a todos os itens abaixo e espero que você também tenha captado algo do texto e das explicações acima.

- [O que é DevOps? - TechWorld with Nana](https://www.youtube.com/watch?v=0yWAtQ6wYNM)
- [O que é DevOps? - GitHub YouTube](https://www.youtube.com/watch?v=kBV8gPVZNEE)
- [O que é DevOps? - IBM YouTube](https://www.youtube.com/watch?v=UbtB4sMaaNM)
- [O que é DevOps? - AWS](https://aws.amazon.com/devops/what-is-devops/)
- [O que é DevOps? - Microsoft](https://docs.microsoft.com/en-us/devops/what-is-devops)

Se você chegou até aqui, saberá se é aqui que você quer estar ou não. Vejo você no [Dia 3](dia03.md).
