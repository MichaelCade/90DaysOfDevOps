---
title: '#90DaysOfDevOps - Planejamento > Codificação > Build > Testes > Release > Deploy > Operação > Monitoramento > - Dia 5'
published: false
description: 90DaysOfDevOps - PlanPlanejamento > Codificação > Build > Testes > Release > Deploy > Operação > Monitoramento >
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048830
---

## Planejamento > Codificação > Build > Testes > Release > Deploy > Operação > Monitoramento

Hoje vamos nos concentrar nas etapas individuais do início ao fim e no ciclo contínuo de uma aplicação no mundo DevOps.
![DevOps](Images/Day5_DevOps8.png)

### Planejamento

Tudo começa com o processo de planejamento, é onde a equipe de desenvolvimento se reúne e descobre quais tipos de recursos e correções de bugs serão lançados em seu próximo sprint. Esta é uma oportunidade para você como engenheiro DevOps se envolver e aprender com que tipo de coisa virá ao seu caminho e também para  influenciar nas decisões deles. Ajudá-los a trabalhar na infraestrutura que você construiu ou orientá-los para algo que funcionará melhor para eles caso não estejam nesse caminho e, portanto, uma coisa importante a destacar aqui é que os desenvolvedores ou a equipe de engenharia de software são seu cliente como engenheiro DevOps, então esta é sua oportunidade de trabalhar com seu cliente antes que ele siga um caminho ruim.

### Codificação

Agora que a sessão de planejamento terminou, eles vão começar a escrever o código. Você pode não estar muito envolvido com isso, mas aqui é um dos lugares em que você pode se envolver com isso. Sempre que eles estiverem escrevendo código, você pode ajudá-los a entender melhor a infraestrutura, então, se eles souberem quais serviços estão disponíveis e como conversar melhor com esses serviços, eles farão isso e, quando terminarem, mesclarão esse código no repositório.

### Build

É aqui que iniciaremos o primeiro de nossos processos de automação, porque pegaremos o código deles e o construiremos juntos. Dependendo de qual linguagem eles estão usando, pode estar compilando ou pode estar criando uma imagem docker a partir desse código. De qualquer forma, vamos passar por esse processo usando nosso pipeline ci/cd.

## Testes

Depois da fase de construção (build), vamos executar alguns testes nele. Agora a equipe de desenvolvimento geralmente escreve o teste, você pode ter algumas informações sobre quais testes são escritos, mas precisamos executar esses testes. O teste é uma maneira de tentar minimizar a introdução de problemas em produção, obviamente não temos como garantir isso, mas queremos chegar o mais próximo possível de garantir que não serão inseridos novos bugs e não serão quebradas coisas que estavam funcionando.

## Release

Assim que os testes passarem, faremos o processo de release (lançamento) e, dependendo novamente do tipo de aplicação em que você está trabalhando, isso pode não ser uma etapa. Você sabe que o código pode estar no repositório do GitHub ou no repositório do git, ou onde quer que ele esteja, mas pode ser o processo de pegar seu código compilado ou a imagem do docker que você criou e colocá-la em um registro ou repositório onde está acessível por seus servidores de produção para o processo de deploy (implantação).

## Deploy

Deploy, é o que fazemos em seguida, porque a implantação é como o final do jogo de tudo isso, porque as implantações (deploys) são quando colocamos o código em produção. Enquanto não fazemos isso é que nossa empresa percebe o valor de todo o esforço e trabalho duro que você e a equipe de engenharia de software colocaram neste produto até este ponto.

## Operação

Uma vez implantado, vamos operá-lo. E operá-lo pode envolver algo como você começar a receber ligações de seus clientes que estão todos irritados porque o site está lento ou a aplicação está lenta, então você precisa descobrir por quê. Em seguida, possivelmente, construa o dimensionamento automático que você sabe lidar, aumente o número de servidores disponíveis durante os períodos de pico e diminua o número de servidores durante os períodos fora do pico de qualquer maneira. Outra coisa operacional que você faz é incluir como um loop de feedback da produção de volta para sua equipe de operações informando sobre os principais eventos que aconteceram na produção, fazer rollback do deploy. Isso pode ou não ser automatizado, dependendo do seu ambiente, o objetivo é sempre automatizá-lo quando possível. Existem alguns ambientes onde você possivelmente precisa fazer algumas etapas antes de estar pronto para fazer isso, mas idealmente você deseja implantar automaticamente como parte da sua automação. Porém caso você esteja fazendo isso, pode ser uma boa idéia incluir em suas etapas operacionais algum tipo de notificação para que sua equipe de operações saiba que uma implantação aconteceu.

## Monitoramento

Todas as partes acima levam à etapa final porque você precisa ter monitoramento, especialmente em torno de questões operacionais, solução de problemas de escalonamento automático, como se você não soubesse que há um problema se não tiver monitoramento instalado para informar que há um problema portanto, algumas das coisas para as quais você pode criar monitoramento são utilização de memória, utilização de espaço em disco, utilização da CPU, API endpoint, tempo de resposta, a rapidez com que esse endpoint está respondendo e uma grande parte disso também são logs. Os logs oferecem aos desenvolvedores a capacidade de ver o que está acontecendo sem precisar acessar os sistemas de produção.

## Limpe & Repita

Uma vez implementado, você volta ao início, ao estágio de planejamento, e repassa tudo novamente.

## Contínua

Muitas ferramentas nos ajudam a alcançar o processo contínuo acima, todo esse código e o objetivo final de ser completamente automatizado, a infraestrutura em nuvem ou qualquer ambiente é frequentemente descrito como Integração Contínua/Entrega Contínua/Implantação Contínua ou “CI/CD” para abreviar. Passaremos uma semana inteira sobre CI/CD mais tarde nos 90 dias com alguns exemplos e orientações para compreender os fundamentos.

### Entrega Contínua

Entrega Contínua = Planejar > Codar > Construir > Testar

### Integração Contínua

Este é efetivamente o resultado das fases de Entrega Contínua acima, mais o resultado da fase de Lançamento(Release). Esse é o caso tanto do fracasso quanto do sucesso, mas isso é realimentado na entrega contínua ou movido para a implantação contínua.

Integração Contínua = Plano > Código > Construir > Teste > Lançamento(Release)

### Implantação Contínua

Se você tiver uma versão bem-sucedida de sua integração contínua, mude para a implantação contínua, que traz as seguintes fases

O lançamento(release) da Integração Contínua foi bem-sucedida = Implantação Contínua = Implantar > Operar > Monitorar

Você pode ver essas três noções Contínuas acima como uma simples coleção de fases do ciclo de vida do DevOps.

Esta última parte foi uma recapitulação para mim no Dia 3, mas acho que isso torna as coisas mais claras para mim.

### Recursos

- [DevOps para desenvolvedores – Engenheiro de software ou DevOps?](https://www.youtube.com/watch?v=a0-uE3rOyeU)
- [Techworld with Nana -DevOps Roadmap 2022 - Como se tornar um engenheiro DevOps? O que é DevOps?](https://www.youtube.com/watch?v=9pZ2xmsSDdo&t=125s)
- [Como se tornar um engenheiro DevOps em 2021 - DevOps Roadmap](https://www.youtube.com/watch?v=5pxbp6FyTfk)

Se você chegou até aqui, saberá se é aqui que deseja estar ou não.

Vejo você no [Dia 6](day06.md).
