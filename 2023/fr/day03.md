## Penser comme un attaquant

Hier, nous avons abordé en détail ce qu'est le DevSecOps. Aujourd'hui, nous allons parler des caractéristiques d'un attaquant. Pour comprendre un attaquant, nous devons penser comme un attaquant.


### Les caractéristiques d'un attaquant

Pour commencer, toutes les entreprises et tous les software sont des vecteurs d'attaque pour un hacker. Il n'existe pas d'endroit 100% sûr. Tout ce qui est possible de faire, c'est de rendre un réseau, une SI, une entreprise, moins "attirante" aux hackers.

![](images/day03-2.jpg)
***[image from this source](https://www.trainerize.me/articles/outrun-bear/)***

En partant de ce principe, les attaquants sont des menaces constantes. 

Ces personnes malveillantes vont identifier les vulnérabilités d'un système en lançant des attaques dans un ordre spécific afin d'avoir accès au SI et extraire des data, ou lancer un ransomware, bref, pour remplir la mission qu'ils se sont attribué.

Les hackers peuvent être chanceux, mais ils travaillent toujours sur des attaques ciblé spécifiques.

Ils peuvent trouver des brêches rapidement, ou non. Toutes les attaques seront différentes.

### Leur motivation

Dans notre rôle d'ingénieur DevOps, nous allons provisionner des infrastructures, des softwares, ou autre. Nous allons probablement déployer tout ça sur différents environnements, différents cloud, différents types de virtualisation et de containerisation.

Nous devons nous poser les bonnes questions:

- **Comment** peuvent ils nous attaquer ?
- **Pourquoi** nous attaqueraient ils ?
- **Qu'avons** nous qui puissent avoir de la valeur pour un attaquant ?

Les motivations d'attaques sont différentes selon les attaquants. Ca peut également être pour s'amuser... Je pense que nous sommes tous passé par là à un moment donné de notre vie, à l'école par exemple, en allant un peu trop loin dans la découverte du réseau de notre fac ou lycée par exemple.

Mais nous avons pu voir dans les récentes attaques que celles-ci ont plutôt des objectifs pécunier ou politiques.

Par exemple, nous avons vu des workspace Kubernetes être utilisé par des attaquants pour se servir de la puissance de calcul disponible afin de miner de la crypto monnaie.

Dans le coeur des attaquants, leur objectif principal est la **DATA**.

Les données d'une entreprises sont extrèmement profitable sur le marché noir. C'est pourquoi nous mettons beaucoup d'effort à protéger nos données et s'assurer qu'elles soient sécurisées et chiffrées. 


### Attack Maps 

Dans le cadre d'une attaque planifié, les attaquants vont devoir mettre en place un plan en identifiant quels sont les services et les types de données ciblées.

Un "schéma d'attaque" (attack map) est une représentation visuel d'une attaque sur un réseau donné. Ce schéma montre les différentes étapes d'une attaques, les outils et techniques utilisé par un attaquants, ainsi que les différents point d'entrées et de sortis d'un réseau. Un schéma d'attaque peut être utilisé pour analyser les détails d'une attaque précédentes, identifier les vulnérabilités d'un réseau et planifier et construire les défenses contre de futures attaques. Elle peut également être utilisé pour communiquer des informations à des personnes non habitué à du langage techniques, comme des directeurs exécutifs, des managers, ou des équipes de juristes.

Vous pouvez voir dans la description ci-dessous qu'une schéma visuel d'attaque doit être créé par toutes les équipes. (red team et blue team. Sujet que nous couvrirons plus tard.)

Si vous souhaitiez construire un schéma d'attaque de votre réseau privé, il serait important de noter les points suivants:

- Construire un schéma de vos applications en y incluant les fluxs de communications et les technologies utilisées.

- Les listes des vulnératilitées et des surfaces d'attaques potentielles.

- Prendre conscience et schématiser la confidentialité, l'intégrité et la disponibilité des données pour chaque connections/intéraction avec des applications.

- Schématiser l'intégralité des attaques et vulnérabilités possible.

Un schéma d'attaque doit ressembler à ça:

![](images/day03-1.png)

En étudiant ce schéma, nous pouvons nous attendre à une attaque par déni de service (DOS) ou une attaque man-in-the-middle afin d'accéder au Bucket S3 pour éviter l'application de sauvegarder les données ou pour forcer l'application à sauvegarder de mauvaises données.

Ce schéma n'est jamais définitif. Pour la même raison que votre application va constamment évoluer en fonction des feedback, ce schéma d'attaque doit constamment évoluer  et être testé. Chaque test doit fournir des feedbacks afin de solidifier les défenses d'une application ou d'un système d'information. Nous pourrions appeler ça "Réponse Continue" dans la boucle des feedback liées à la sécurité.

Pour améliorer la sécurité, nous devons suivre 3 modèles différents:

- **Good** - Construire l'application selon un modèle "security by design" afin de réduire les attaques potentielles. 

- **Better** - Prioriser et construire des outils de sécurité pour les problèmes identifié lors du cycle de développement.

- **Best** - Construire et automatiser des scripts lors des déploiement pour détecter des soucis, faire des test unitaire, faire des tests de sécurité et des "Black Box" tests.

La sécurité peut être une contrainte lors de la conception d'un design.

## Resources 

- [devsecops.org](https://www.devsecops.org/)

- [TechWorld with Nana - What is DevSecOps? DevSecOps explained in 8 Mins](https://www.youtube.com/watch?v=nrhxNNH5lt0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=1&t=19s)

- [What is DevSecOps?](https://www.youtube.com/watch?v=J73MELGF6u0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=2&t=1s)

- [freeCodeCamp.org - Web App Vulnerabilities - DevSecOps Course for Beginners](https://www.youtube.com/watch?v=F5KJVuii0Yw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=3&t=67s)

- [The Importance of DevSecOps and 5 Steps to Doing it Properly (DevSecOps EXPLAINED)](https://www.youtube.com/watch?v=KaoPQLyWq_g&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=4&t=13s)

- [Continuous Delivery - What is DevSecOps?](https://www.youtube.com/watch?v=NdvMUcWNlFw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=5&t=6s)

- [Cloud Advocate - What is DevSecOps?](https://www.youtube.com/watch?v=a2y4Oj5wrZg&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=6)

- [Cloud Advocate - DevSecOps Pipeline CI Process - Real world example!](https://www.youtube.com/watch?v=ipe08lFQZU8&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=7&t=204s)

On se retrouver lors du [jour 4](day04.md) 
