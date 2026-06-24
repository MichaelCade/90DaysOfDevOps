## Open Source Security 

Les logiciels et solutions Open-Source ont été massivements adopté ces dernières années par les entreprises. Cela est dû notamment à l'approche communautaire et collaborative des projets.

L'appelation Open-Source fait référence à des logiciels/application dans le domaine publique qui peuvent être utilisé, modifié et partagé gratuitement.

La raison principale de cette adoption massive de solution Open Source vient du fait que ces solutions sont souvent utilisé en complément de solutions propriétaire, ce qui permet d'accélerer le time-to-market des entreprise. L'utilisation de solutions Open Source accélère le développement des applications et aide les entreprises à pousser leurs produits sur le marché plus rapidement.


### Qu'est-ce que la sécurité Open Source ?

La sécurité Open Source fait référence aux pratiques de sécurité mise en place pour sécuriser un système d'information qui utilise des logiciel open source.

Comme mentionné précédemment, les solutions open-source sont libre d'être utilisé, modifié et distribué gratuitement. Ces solutions sont souvent développer par une communauté de volontaire. Néanmoins, il y a de plus en plus de gros éditeur de logiciel qui contribuent à des projet open-source. Il n'y à qu'a regarder le repository Kubernetes pour se rendre compte des éditeurs qui ont décidé d'investir dans le projet.

Comme les solutions Open Source sont gratuite, elles sont facilement utilisé et étudié, permettant ainsi d'améliorer leur sécurité. Néanmoins, il est important de s'assurer que ces solutions sont utilisées de manière responsable et que les vulnérabilitées qui peuvent être trouvé sont communiqué rapidement pour maintenir un haut niveau de sécurité.


### Comprendre la supply chain de la sécurité des solutions Open Source

J'ai l'habitude de documenter mes trouvailles dans de petits paragraphes lorsque les sources vidéos sont assez longue. Néanmoins, cette vidéo là ne dure que 10 minutes, du coup, il est probablement plus simple que je vous fournisse le lien de la vidéo directement. [Understanding Open-Source Supply Chain Security](https://www.youtube.com/watch?v=pARGj6j0-ZY)

Que l'on soit une entreprise éditeur de logiciel qui utilise des solutions Open Source avec parcimonie, un projet communautaire qui utilise des packages Open Source ou autre, nous devons être conscient des enjeux de sécurité et proposer une meilleure visibilité entre les projets.


### Les 3 As de la sécurité des outils Open Sources

Une autre source que j'ai trouvé intéressant d'aborder ici nous viens d'IBM, le lien sera disponible dans la section ressource.

- **Assess** - Evaluer la santé du projet, si le repository est plutot actif, quel est le temps de réponse de la communauté qui maintiens le projet ? Si vous souhaitez utiliser une solution open source dont ces signaux sont mauvais, vous n'allez probablement pas être très content de la sécurité du projet.

On peut également jeter un oeil au modèle de sécurité mise en place sur le projet, comment sont fait les reviews du code, les validations des données, les tests spécialisé pour la sécurité, et surtout, où se situe le projet par rapport aux CVE déjà sortie ?

Quelles sont les dépendances du projet ? Vous devriez également regarder la santé générale des dépendants pour être sûr que l'intégralité de la stack de présente pas de risque de sécurité.

- **Adopt** - Si vous souhaitez utiliser une solution open source à l'intérieur de votre projet, ou comme une application Stand Alone, vous devez décider rapidement qui va manager l'application et la maintenir au sein de votre organisation ? Mettre en place une vrai gouvernance.

- **Act** - La sécurité est la responsabilité de tous, pas seulement la communauté de la solution. En tant qu'utilisateur, il est important de communiquer et de donner du feedback du projet.


### Vulnérabilité Log4j

En 2022, une important vulnérabilité à été mise en lumière (Log4j (CVE-2021-44228) RCE Vulnerability)

Log4j est une librairie assez commune de logging Java. Cette vulnérabilité a affecté des millions d'application basé sur du java.

Une personne malveillante pouvait utiliser cette vulnérabilité dans une application pour gagner accès à un système.

2 choses importantes que j'ai mentionné:

- Des **millions** d'application utilisent ce package.
- Des **acteurs malveillants** peuvent utiliser cette vulnérabilité pour avoir accès à un système ou pour déployer un malware dans un environnement.

Pourquoi je mentionne cet exemple ? La sécurité d'un système ne s'arrête jamais. La croissance de l'adoption d'outil open source à augmenter les vecteurs d'attaques sur des applications et c'est pourquoi nous devons tous faire un effort concernant la sécurité.


## Ressources 

- [Open Source Security Foundation](https://openssf.org/)
- [Snyk - State of open source security 2022](https://snyk.io/reports/open-source-security/)
- [IBM - The 3 A's of Open Source Security](https://www.youtube.com/watch?v=baZH6CX6Zno)
- [Log4j (CVE-2021-44228) RCE Vulnerability Explained](https://www.youtube.com/watch?v=0-abhd-CLwQ)

Rendez vous au [Day 6](day06.md).
