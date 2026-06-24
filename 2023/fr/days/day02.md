## Vue d'ensemble: DevSecOps

Bienvenue dans la 2e journée de cette édition 2023. Dans le premier module de ces 6 prochains jours, nous allons aborder les bases fondamentales de ce qu'est le DevSecOps.

### Qu'est ce que le DevSecOps? 

DevSecOps est une approche du développement logiciel dont le but est de réunir les équipes de développement, de sécurité et les équipes opérationnelles autour du développement et de la sécurisation d'une application.

L'approche est basée sur l'intégration, la livraison et le déploiement continu dont le but est de délivrer des mise à jours et des nouvelles fonctionnalités plus rapidement. 

Dans l'approche DevSecOps, la sécurisation se fait dès le développement de l'application, et non après coup. Concrètement, les tests de sécurité, la surveillance (monitoring) et les autres mesures de sécurisation sont développés dès le début du cycle de vie du développement logiciel, plutôt qu'ajouté après.

DevSecOps oeuvre pour améliorer la collaboration et la communication entre les équipes de développement, de sécurité et les équipes opérationnelles afin de créer des processus de développement logiciel plus efficace.

### DevSecOps vs DevOps 

J'utilise le terme "vs" ici de manière anecdotique. Le but n'étant pas de confronter les 2 termes, mais bien d'en comprendre les différences. Si vous vous rappelez bien de l'édition 2022, l'objectif du DevOps était d'améliorer la vitesse de traitement et de test, la fiabilité et la qualité générale des logiciels déployés.

DevSecOps est une extension de la philosophie DevOps. Le but est d'intégrer les bonnes pratiques en terme de sécurité lors du processus de dévelopement d'un logiciel. Ainsi, la sécurité du logiciel est intégralement partie prenante du développement d'un logiciel depuis le début plutôt qu'à la fin. Cela va aider à réduire les risques de vulnérabilités et à les rendre plus facile à identifier et à patcher. 

Le DevOps se concentre sur l'amélioration de la communication entre les équipes de développements et les équipes opérationnelles afin d'améliorer la rapidité de traitement, la fiabilité, et la qualité des logiciels délivré. Le DevSecOps quant à lui, se concentre sur l'intégration des bonnes pratiques de sécurité dans le processus de développement d'une application afin de réduire les risques de vulnérabilité et pour améliorer la sécurité du logiciel et des sytèmes d'informations en général. 

### Automatisation de la sécurité

L'automatisation de la sécurité fait référence a l'utilisation de différentes technologies, différents outil, pour lancer des tâches de sécurité sans intervention humaine.
Il existe différent moyen de sécurisé une application ou son SI, par exemple, l'utilisation d'un outil de monitoring de réseau pour détecter des menaces et les bloquer (IPS/IDS) ou l'utilisation d'outil basé sur des intelligence articielle pour analyser des paterns d'attaque pour identifier des activités inhabituelles. Les outils d'automatisation des sytèmes de sécurité ont été conçu pour faire en sorte que la sécurité d'une application, d'un SI ou autre se fasse de manière efficace et pour réduire les charges de travail redondantes des ingénieur spécialisé en cybersécurité.

La plus value qu'apporte le DevSecOps est la possibilité d'automatisé un grand nombre de tâches lorsqu'on développe et qu'on délivre une application. Lorsqu'on ajoute la partie sécurité dès le début du cycle de développement, il faut prendre en compte l'automatisation de cette partie sécurité.

### Security at Scale (Containers and Microservices)

La création et le déploiement de microservice et de container ont changé les manière de travailler de la plus part des entreprises, notamment grace à la grande scalabilité et dynamique que peuvent offrir ces nouveaux services.

C'est également pour cette raison que nous devons introduire l'automatisation des tâches de sécurité dans les principes fondamentaux du DevOps. Nous devons nous assurer que la sécurité des containers et des microservices est conforme aux guidelines de sécurité mise en place dans les entreprises. 

Pour être un peu plus précis, grâce aux technologies cloud-native, les entreprises ne peuvent plus se permettre de garder une politique de sécurité statique. Les modèles de sécurité des entreprises se doivent d'être aussi dynamique que les charges de travail et de comment elles tournent.

Les équipes DevOps se doivent d'inclure des tâches automatique de gestion de la sécurité pour protéger l'environnement et les données dans son ensemble (Système d'information de manière générale) 

La liste suivante est tiré du post [RedHat](https://www.redhat.com/en/topics/devops/what-is-devsecops)

- Standardiser et automatiser l'environnement: chaque service doit avoir le moins de privilèges possible afin de minimiser les connections et les accès non autorisé.

- Centraliser les identités utilisateur et les RBAC: Des RBAC consciencieusement controlés et une centralisation de l'authentification sont des mécanismes essentiels pour sécuriser les microservices.

- Isoler les container, les microservices et le réseaux associé les uns des autres: Ce qui inclu également les donnés "chaudes", lesdonnées "froides" et les données en circulation. Toutes données représentent une grosses valeurs pour les attaquants.

- Chiffrer les données entre les apps et les services: Un orchestrateur de container (kub) avec une intégration de fonctionnalitées de sécurités aide a minimiser les chances d'accès non autorisé. 

- Créer des API Gateways sécurisé: Sécurisé les API augmente la visibilité sur les autorisations et le routing. En réduisant l'exposition des API, les entreprises et organisation peuvent réduire les surfaces d'attaques.

### La sécurité au centre de tout

Peu importe de votre parcours dans l'IT, vous ne pouvez pas être passé à coté du faite que la sécurité est un sujet très important dans toutes l'industries récemment. C'est notamment dû a l'apparition de breches de sécurité dans des grandes entreprises ou l'utilisation de mauvaises habitudes en matières de sécurités. De mon point de vu, la création et le développement de logiciels et d'application est beaucoup plus réalisable et accessible aujourd'hui qu'avant. Mais lors de la création d'application, l'exposition aux vulnérabilité est nettement en hausse. Ceci permets à de mauvaises personnes de voler des données, lancer des ransomware et causer la fermeture d'entreprises. Nous avons déjà beaucoup discuté de ce qu'est le DevSecOps, mais je crois fortement qu'explorer les vecteurs d'attaques et de comprendre pourquoi nous devons protéger notre cycle de développement est inavitable pour éviter les attaques informatiques, ou du moins, réduire les surfaces d'attaques.


### Cybersecurity vs DevSecOps

Il est important de noter les différences entre cybersécurité et DevSecOps afin de comprendre pourquoi la sécurité doit être intégré dans les process, les principes et la méthodologie DevOps.

La cybersécurité consiste à protéger le système d'information (données, systèmes, réseaux) d'attaquant malveillant, de voleur et de dommages physique ou virtuel. Il est important d'identifié et de comprendre les vulnérabilités, de mettre en place des mesures de sécurités, et de déployer des services de monitoring.

De l'autre côté, le DevSecOps et une combinaison des pratiques de développements, de sécurité et d'opérations. C'est une philosophie dont le but est d'intégré la sécurité lors du développement d'une application plutôt que de l'intégrer après-coup. Cela implique la collaboration entre les équipes de développement, de sécurité et les équipes opérationnelle durant le cycle de développement des systèmes d'informations.

Voici les différences notables entre la cybersécurité et la philosophie DevSecOps: 

**Focus**: Les équipes de cybersécurité sont principalement concentré sur la protection du systèmes d'information des menaces, alors que le DevSecOps se concentre sur l'intégration de la sécurité dans les process de développements.

**Scope**: La cybersécurité couvre une grande diversité de sujets, notamment la sécurité des réseaux, des données, des applications et bien plus encore. Le DevSecOps, de l'autre côté, se concentre sur l'amélioration de la sécurité dans le développement et le déploiement des applications.

**Approach**: Les équipes de cybersécurité implémentent des mesures de sécurités après que les processus de développement soient finis. L'approche DevSecOps est d'intégré la partie sécurisation dès le début des processus de développement.

**Collaboration**: La cybersécurité implique souvent la collaboration entre les équipes IT et les équipes sécurité, alors que le DevSecOps implique la communication entre les équipes de développement, de sécurité, et les équipes opérationnelles.


## Ressources 

Durant tous le projet, vous verrez apparaitre une liste de ressource qui vous permettrons de creuser un peu plus dans les différents sujets abordé.

- [TechWorld with Nana - What is DevSecOps? DevSecOps explained in 8 Mins](https://www.youtube.com/watch?v=nrhxNNH5lt0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=1&t=19s)

- [What is DevSecOps?](https://www.youtube.com/watch?v=J73MELGF6u0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=2&t=1s)

- [freeCodeCamp.org - Web App Vulnerabilities - DevSecOps Course for Beginners](https://www.youtube.com/watch?v=F5KJVuii0Yw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=3&t=67s)

- [The Importance of DevSecOps and 5 Steps to Doing it Properly (DevSecOps EXPLAINED)](https://www.youtube.com/watch?v=KaoPQLyWq_g&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=4&t=13s)

- [Continuous Delivery - What is DevSecOps?](https://www.youtube.com/watch?v=NdvMUcWNlFw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=5&t=6s)

- [Cloud Advocate - What is DevSecOps?](https://www.youtube.com/watch?v=a2y4Oj5wrZg&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=6)

- [Cloud Advocate - DevSecOps Pipeline CI Process - Real world example!](https://www.youtube.com/watch?v=ipe08lFQZU8&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=7&t=204s)

J'espère que cette journée à attiser votre curiosité sur les sujets abordé et que les ressources précédemment listé vous aiderons à creuser un peu plus dans les différents topics. Lors de la [Day 3](day03.md) nous essaierons de comprendre comment un attaquant peut réfléchir et ainsi comprendre pourquoi et comment nous pouvons protéger nos SI et application dès le début.