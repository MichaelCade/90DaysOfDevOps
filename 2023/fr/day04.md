## <span style="color:red">Red Team</span> vs. <span style="color:blue">Blue Team</span>

J'aimerai revenir sur quelque chose que j'ai mentionné hier: <span style="color:red">**Red**</span> et <span style="color:blue">**Blue**</span> teams. 

Dans le domaine de la sécurité, les <span style="color:red">**Red**</span> teams et <span style="color:blue">**Blue**</span> teams travaillent comme attaquants et défenseurs pour améliorer la sécurité d'une entreprise et d'une organisation

Les 2 équipes travaillent pour améliorer la sécurité de l'entreprise, mais de manières différentes.

La <span style="color:red">**Red**</span> team à le rôle d'attaquant. Cette équipe essaye de trouver des vulnérabilitées dans de code d'une application ou dans l'infrastructure d'un système d'information. Ils essayent alors de casser les défenses et de s'introduire dans le système informatique de l'entreprise. 

La <span style="color:blue">**Blue**</span> team à plutôt un rôle de défenseur. Leur objectif et de défendre le SI contre les attaquants et de répondre aux incidents qui peuvent arriver.

![](images\day04-2.jpg)
***[image from this source](https://hackernoon.com/introducing-the-infosec-colour-wheel-blending-developers-with-red-and-blue-security-teams-6437c1a07700)***

### Les bénéfices

Une bonne façon d'amélioré la sécurité d'une entreprise peut être d'organiser des exercices entre les <span style="color:red">**Red**</span> team et <span style="color:blue">**Blue**</span> team. L'idée est de créer un scénario au plus proche de la réalité et d'une véritable attaque. Voici quelques-uns des domaines dans lesquels cette approche sera utile :

- Découverte des vulnérabilités
- Durcissement (hardening)
- Gagner de l'expérience dans la détection et l'isolement des attaques
- Contruire des plans de réponses et de remédiation (+ PRA/PCA)
- Faire prendre conscience de l'importance de la sécurité dans l'entreprise.

### <span style="color:red">Red Team</span> 

La NIST (national institute of standards and technology) décrit la <span style="color:red">**Red**</span> team de la manière suivante:

"Un groupe de personne organisé et autorisé a émuler les capacité d'attaque ou d'exploitation contre la sécurité d'une entreprise." 

Dans une scénario de simulation d'attaque, ils jouent les "méchants".

Lorsque l'on parle de <span style="color:red">**Red**</span> team et <span style="color:blue">**Blue**</span> team, il en va au delà des simples processus de DevSecOps et des principes des cycles de vie logiciel. Néanmoins, en apprendre un peu plus sur la sécurité de manière général ne fera de mal à personne. Qui peut le plus, peut le moins.

La tâche principal de la <span style="color:red">**Red**</span> team est de penser comme un attaquant (ce que nous avons vu lors de notre dernière session.) notamment en y ajoutant les principes de social engineering et y incluant toutes les équipes de l'entreprise afin de potentiellement les manipuler et avoir un accès au réseau et aux services informatiques dans son ensemble.

Une partie fondamentale de la <span style="color:red">**Red**</span> team est de comprendre les principes de fonctionnement du cycle de vie de développement d'un logiciel. En comprenant comment les applications sont constuites, il est plus facile d'identifier des faiblesses et des failles, de réécrire un programme et de procéder à une élévation de privilège ou de lancer un exploit. 

Vous devez avoir déjà entendu parler du terme "Penetration testing" ou "pen test". Pour une <span style="color:red">**Red**</span> team, cela signifie d'identifier et d'essayer d'exploiter des vulnérabilités connues dans un environnement données. Nous couvrirons un peu plus tard comment la monté en puissance des outils open-source peut contribuer à aider les membres des <span style="color:red">**Red**</span> teams.

### <span style="color:blue">Blue Team</span> 

La NIST (national institute of standards and technology) décrit la <span style="color:blue">**Blue**</span> team de la manière suivante:

"L'équipe responsable de la défense du système d'information d'une entreprise en maintenant un haut niveau de sécurité contre une équipe d'attaquant."

La <span style="color:blue">**Blue**</span> team joue la défense. Il ont la charge d'analyser les politiques de sécurité actuelle d'une entreprise et de mettre des actions en place pour empêcher ou ralentir les attaques. La <span style="color:blue">**Blue**</span> team est également en charge du monitoring continu (notion que nous avons couvert à la fin de l'édition 2022) de brêches et de vulnérabilité.

En tant que membre de la <span style="color:blue">**Blue**</span>, vous allez devoir comprendre quelles type d'asset vous devez protéger, et comment les protéger de la meilleure manière possible. Dans l'IT aujourd'hui, nous avons de nombreuses manière de faire tourner nos workload, nos applications, et d'héberger nos données.

- Etude de risque - L'étude de risque vous apportera une meilleure vision des asset les plus critiques de votre entreprise/business.

- Veille technologique - Quels sont les menaces ? Il y a des milliers de vulnérabilités, dont beaucoup qui ne sont pas patché. Comment atténuer les risques sans ralentir et interrompre le business de votre entreprise ? 

### Cybersecurity colour wheel 

Comme l'important de la sécurité informatique grandit, notamment à cause de l'impact que peuvent avoir les attaques sur les entreprises, il y a un besoin plus important que les simples <span style="color:red">**Red**</span> and <span style="color:blue">**Blue**</span> teams lorsque l'on parle de sécurité à l'intérieur d'un business.


![](images\day04-1.png)
***[image from this source](https://hackernoon.com/introducing-the-infosec-colour-wheel-blending-developers-with-red-and-blue-security-teams-6437c1a07700)***

- La <span style="color:yellow">**Yellow Team**</span> sont les "builders", les ingénieurs et les développeurs qui contruisent les applications.

"Nous avons les <span style="color:red">**Red**</span> and <span style="color:blue">**Blue**</span> comme nous avons toujours eu jusqu'à présent. Maintenant, avec la création de la <span style="color:yellow">**Yellow**</span> Team, nous avons d'autre équipe de couleur qui se créé (Orange, Green, et Purple) dédié à créer un esprit collaboratif et à mixer les skills entre les attaquants, les défenseurs, et les développeurs. Il est important de rendre le code plus sécurisé, et par extension, l'entreprise."

La citation au dessus est tiré de la première source listé à la fin du post. 

<span style="color:red">**Red**</span>, <span style="color:blue">**Blue**</span>, <span style="color:yellow">**Yellow**</span> sont les couleurs primaires. En les combinant, nous commencons a comprendre où interviennent les autres couleurs (ou couleurs secondaire).

- La <span style="color:purple">**Purple Team**</span> - The special team! Si vous prenez le <span style="color:blue">**Blue**</span> et le <span style="color:red">**Red**</span>, vous obtenez le <span style="color:purple">**Purple**</span>. Si ovus intégré les processus de défense dans les processus d'attaque et que vous collaborez en partageant le savoir des 2 équipes, vous obtiendrez une meilleure politique de sécurité dans son ensemble.

- La <span style="color:green">**Green Team**</span> - Feedback loop, la <span style="color:green">**Green**</span> team va récupérer et rassembler les feedback interne de la <span style="color:blue">**Blue**</span> team et travailler étroitement avec la <span style="color:yellow">**Yellow**</span> team pour être plus efficace. Mélanger la <span style="color:blue">**Blue**</span> et la <span style="color:green">**Green**</span> et vous <span style="color:purple">**obtenez**</span>?

- La <span style="color:orange">**Orange Team**</span>, comme la <span style="color:green">**Green**</span> team, travaille avec la <span style="color:blue">**Blue**</span> team pour les feedbacks, la <span style="color:orange">**Orange**</span> team, travaille avec la <span style="color:red">**Red**</span> team et partager les informations apprises avec la <span style="color:yellow">**Yellow**</span> team pour améliorer la sécurité du code.

Lorsque j'ai commencé a chercher ces notions, je me suis rendu compte que nous nous éloignons un peu des sujets DevOps habituel. Si une personne de la sphère DevSecOps lis ceci, est ce que ça vous parait correct ? Utils ? Et avez vous quelque chose à ajouter ?


## Resources 

- [Introducing the InfoSec colour wheel — blending developers with red and blue security teams.](https://hackernoon.com/introducing-the-infosec-colour-wheel-blending-developers-with-red-and-blue-security-teams-6437c1a07700)


On se retrouve au [Day 5](day05.md).

