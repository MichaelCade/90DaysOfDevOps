---
title: '#90DaysOfDevOps - Foco na Aplicação - Dia 3'
published: false
description: 90DaysOfDevOps - Foco na Aplicação - Dia 3
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048825
---

## Ciclo de vida do DevOps - Foco na Aplicação

À medida que continuarmos nas próximas semanas, em 100% das vezes encontraremos esses títulos (Desenvolvimento Contínuo, Teste, Implantação, Monitoramento) repetidamente. Se você está indo para a função de engenheiro DevOps, a repetibilidade será algo com a qual você se acostumará, mas melhorar constantemente a cada vez é outra coisa que mantém as coisas interessantes.

Neste momento, vamos dar uma olhada na visão de alto nível da aplicação do início ao fim e depois voltar novamente como um loop constante.

### Desenvolvimento

Vamos dar um novo exemplo de uma aplicação. Para começar não temos nada criado e talvez como desenvolvedor, você tenha que discutir com seu cliente ou usuário final os requisitos e criar algum tipo de plano ou requisitos para sua aplicação. Em seguida, precisamos criar a partir dos requisitos nossa nova aplicação.

Com relação às ferramentas nesta etapa, não há nenhum requisito real aqui além de escolher sua IDE e a linguagem de programação que você deseja usar para escrever sua aplicação.

Como engenheiro DevOps, lembre-se de que provavelmente não é você quem está criando este plano ou codificando a aplicação para o usuário final, este será um desenvolvedor habilidoso.

Mas também não faria mal para você poder ler parte do código para poder tomar as melhores decisões de infraestrutura no futuro para sua aplicação.

Mencionamos anteriormente que esta aplicação pode ser escrito em qualquer linguagem. É importante que isso seja mantido usando um sistema de controle de versão, isso é algo que também abordaremos em detalhes mais adiante e, em particular, mergulharemos no **Git**.

Também é provável que não seja um desenvolvedor trabalhando neste projeto, embora esse possa ser o caso, mesmo que as práticas recomendadas exigissem um repositório de código para armazenar e colaborar no código, isso pode ser privado ou público e pode ser implantando de forma hospedada ou privada. De maneira geral, você teria o **GitHub ou GitLab** sendo usado como um repositório de código. Mais uma vez, abordaremos isso como parte de nossa seção sobre **Git** mais adiante.

### Teste

Nesta etapa, temos os nossos requisitos e temos a nossa aplicação a ser desenvolvida. Mas precisamos ter certeza de que estamos testando nosso código em todos os diferentes ambientes que temos disponíveis para nós ou talvez especificamente para a linguagem de programação escolhida.

Essa etapa permite que o time de QA execute teste de bugs. De forma cada vez mais frequente vemos contêineres sendo usados para simular o ambiente de teste que, em geral, pode melhorar as despesas gerais de infraestrutura física ou em nuvem.

Essa etapa provavelmente também será automatizada como parte da próxima área, que é a Integração Contínua.

A capacidade de automatizar esse teste contra 10,100 ou mesmo 1000 engenheiros de controle de qualidade que precisam fazer isso manualmente fala por si, esses engenheiros podem se concentrar em outra coisa dentro da pilha para garantir que você esteja se movendo mais rápido e desenvolvendo mais funcionalidades em vez de testar bugs e software que tende a ser o atraso na maioria dos lançamentos de software tradicionais que usam uma metodologia em cascata.

### Integração

Muito importante, a integração está no meio do ciclo de vida do DevOps. É a prática na qual os desenvolvedores precisam confirmar alterações no código-fonte com mais frequência. Isso pode ser diário ou semanal.

A cada commit, sua aplicação pode passar pelas fases de teste automatizadas e isso permite a detecção antecipada de problemas ou bugs antes da próxima fase.

Agora você pode estar dizendo: "Mas nós não criamos aplicações, nós os compramos de um fornecedor de software". Não se preocupe, muitas empresas fazem isso e continuarão a fazer isso e será o fornecedor de software que está se concentrando nas três fases acima, mas talvez você ainda queira adotar a fase final, pois isso permitirá implantações mais rápidas e eficientes de suas implantações prontas para uso.

Eu também sugeriria que apenas ter esse conhecimento acima é muito importante, pois você pode comprar software de prateleira hoje, mas e amanhã ou no futuro... talvez no próximo trabalho?

### Implantação

Ok, então temos nossa aplicação construída e testada de acordo com os requisitos de nosso usuário final e agora precisamos seguir em frente e implantar essa aplicação em produção para nossos usuários finais consumirem.

Este é o estágio em que o código é implantado nos servidores de produção, agora é onde as coisas ficam extremamente interessantes e é onde o restante de nossos 86 dias se aprofunda nessas áreas. Porque aplicações diferentes exigem hardware ou configurações diferentes. É aqui que o **Gerenciamento de configuração de aplicações** e a **Infraestrutura como Código** podem desempenhar um papel fundamental no ciclo de vida do DevOps. Pode ser que sua aplicação seja **Containerizada**, mas também esteja disponível para execução em uma máquina virtual. Isso também nos leva a plataformas como o **Kubernetes**, que orquestrariam esses contêineres e garantiriam que você tivesse o estado desejado disponível para seus usuários finais.

Desses tópicos ousados, entraremos em mais detalhes nas próximas semanas para obter um melhor conhecimento básico sobre o que são e quando usá-los.

### Monitoramento

As coisas estão se movendo rapidamente aqui e temos nossa aplicação que estamos atualizando continuamente com novos recursos e funcionalidades. E ainda temos nossos testes para garantir que nenhum gremlin seja encontrado. Temos a aplicação em execução em nosso ambiente que pode manter continuamente a configuração e o desempenho necessários.

Mas agora precisamos ter certeza de que nossos usuários finais estão obtendo a experiência de que precisam. Aqui precisamos ter certeza de que o desempenho das nossas aplicações está sendo monitorado continuamente, esta etapa permitirá que seus desenvolvedores tomem melhores decisões sobre melhorias na aplicação em versões futuras para melhor atender os usuários finais.

Esta seção também é onde vamos capturar o ciclo de feedback sobre os recursos que foram implementados e como os usuários finais gostariam de torná-los melhores para eles.

A confiabilidade também é um fator chave aqui, no final das contas, queremos que nossa aplicação esteja disponível o tempo que for necessário. Isso leva a outras áreas de **observabilidade, segurança e gerenciamento de dados** que devem ser monitoradas continuamente e o feedback sempre pode ser usado para melhorar, atualizar e liberar a aplicação continuamente.

Algumas contribuições da comunidade aqui especificamente [@\_ediri](https://twitter.com/_ediri) mencionam que também fazem parte desse processo contínuo termos as equipes de FinOps envolvidas. As aplicações e dados estão sendo executados e armazenados em algum lugar que você deve monitorar continuamente para garantir que, se as coisas mudarem do ponto de vista dos recursos, seus custos não estejam causando grandes problemas financeiros em suas contas de Cloud.

Eu acho que também é um bom momento para trazer à tona o "Engenheiro DevOps" mencionado acima. Embora existam muitas posições de Engenheiro DevOps que as pessoas ocupam, essa não é a maneira ideal de posicionar o processo de DevOps. O que quero dizer é que, ao falar com outras pessoas da comunidade, o título de Engenheiro DevOps não deve ser o objetivo de ninguém, porque realmente qualquer posição deve adotar os processos de DevOps e a cultura explicada aqui. O DevOps deve ser usado em muitas posições diferentes, como engenheiro/arquiteto de tecnologias cloud native, administrador de virtualização, arquiteto/engenheiro de nuvem e administrador de infraestrutura. Isso é para citar alguns, mas o motivo para usar o Engenheiro DevOps acima foi realmente destacar o escopo do processo usado por qualquer uma das posições acima e muito mais.

## Recursos

Estou sempre aberto a adicionar recursos adicionais a esses arquivos README, pois está aqui como uma ferramenta de aprendizado.

Meu conselho é assistir a todos os itens abaixo e espero que você também tenha captado algo do texto e das explicações acima.

- [Desenvolvimento Contínuo](https://www.youtube.com/watch?v=UnjwVYAN7Ns) Também acrescentarei que isso é focado na indústria, mas a cultura enxuta pode ser seguida de perto com o DevOps.
- [Testes Contínuos - IBM YouTube](https://www.youtube.com/watch?v=RYQbmjLgubM)
- [Integração Contínua - IBM YouTube](https://www.youtube.com/watch?v=1er2cjUq1UI)
- [Monitoramento Contínuo](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [O fluxo remoto](https://www.notion.so/The-Remote-Flow-d90982e77a144f4f990c135f115f41c6)
- [Fundação FinOps - O que é FinOps](https://www.finops.org/introduction/what-is-finops/)
- [**Não gratuito** O projeto Fênix: Um romance sobre TI, DevOps e como ajudar sua empresa a vencer](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/1942788290/)

Se você chegou até aqui, saberá se é aqui que você quer estar ou não. Vejo você no [dia 4](day04.md).
