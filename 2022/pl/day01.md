---
title: '#90DaysOfDevOps - Introduction - Day 1'
published: true
description: 90DaysOfDevOps - Introduction
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048731
date: '2022-04-17T10:12:40Z'
---
## Wstęp - dzień 1 

Dzień 1 naszej 90-dniowej przygody, której celem jest poznanie podstaw DevOps i narzędzi, które pomagają w przyjęciu postawy DevOps. 

Moja podróż z DevOps rozpoczęła się dla mnie kilka lat temu, ale wówczas koncentrowałem się na platformach wirtualizacyjnych i technologiach opartych na chmurze, zajmowałem się głównie Infrastructure as Code i zarządzaniem konfiguracją aplikacji za pomocą Terraform i Chef. 

W marcu 2021 r. otrzymałem wspaniałą możliwość skoncentrowania swoich wysiłków na strategii Cloud Native w firmie Kasten by Veeam. Miało to oznaczać ogromny nacisk na Kubernetes i DevOps oraz społeczność związaną z tymi technologiami. Wtedy też zacząłem rozmawiać ze społecznością i dowiadywać się coraz więcej o kulturze, narzędziach i procesach DevOps, dlatego zacząłem publicznie dokumentować niektóre obszary, których chciałem się nauczyć. 


[So you want to learn DevOps?](https://blog.kasten.io/devops-learning-curve)

## Rozpocznijmy podróż

Jeśli przeczytasz powyższy blog, zobaczysz, że jest to ogolny przeglad przebiegu mojej podróży edukacyjnej i zaznaczajac, że w tym momencie nigdzie nie jestem ekspertem w żadnej z tych sekcji, chciałem podzielić się pewnymi materialami, zarówno DARMOWYMI, jak i płatnymi, ale z opcją dla obu, ponieważ wszyscy mamy różne okoliczności.  

W ciągu następnych 90 dni chcę udokumentować te materialy i omówić podstawowe obszary. Bardzo bym chciał, aby społeczność również się w to zaangażowała. Podziel się swoją podróżą i zasobami, abyśmy mogli uczyć się publicznie i pomagać sobie nawzajem. 

W pliku Readme w repozytorium projektu mozesz zauwazyc, że podzieliłem projekt na części i jest to 12 tygodni plus 6 dni. Przez pierwsze 6 dni będziemy poznawać podstawy DevOps w ogóle, a następnie zagłębimy się w niektóre z konkretnych obszarów. Lista ta w żadnym wypadku nie jest wyczerpująca i chciałbym, aby społeczność pomogła mi stworzyc z niej użyteczne źródło wiedzy.

Kolejnym źródłem, którym chciałbym się podzielić i które, jak sądzę, każdy powinien dobrze przejrzeć, być może stworzyć mapę myśli dla siebie, swoich zainteresowań i pozycji, jest poniższa publikacja:

[DevOps Roadmap](https://roadmap.sh/devops)

Uważam, że jest to świetne źródło informacji, bylo niezwykle przydatne gdy tworzyłem swoją pierwszą listę i wpis na blogu na ten temat. Możesz również zobaczyć inne obszary, w których omówiono dużo więcej szczegółów poza 12 tematami, które wymieniłem w tym repozytorium. 

## Pierwsze kroki - Czym jest DevOps? 

Jest tak wiele artykułów na blogach i filmów na YouTube, że nie sposób ich tutaj wymienić, ale ponieważ zaczynamy 90-dniowe wyzwanie i skupiamy się na poświęceniu około godziny dziennie na nauczenie się czegoś nowego lub związanego z DevOps, pomyślałem, że dobrze byłoby na początek przybliżyć niektóre z najważniejszych informacji na temat tego, czym jest DevOps. 

Po pierwsze, DevOps nie jest narzędziem. Nie można go kupić, nie jest to SKU oprogramowania ani repozytorium open source na GitHubie, które można pobrać. Nie jest to także język programowania, nie jest to też jakaś czarna magia. 

DevOps to sposób na robienie mądrzejszych rzeczy w procesie tworzenia oprogramowania. - Chwila... Ale jeśli nie jesteś programistą, to czy powinieneś się teraz odwrócić i nie brać udziału w tym projekcie? Nie. W żadnym wypadku. Zostań... Ponieważ DevOps łączy w sobie rozwój oprogramowania i operacje. Wspomniałem wcześniej, że jestem bardziej po stronie maszyn wirtualnych, a to z reguły podlega pod stronę operacyjną, ale w społeczności są ludzie z różnych środowisk, dla których DevOps będzie w 100% korzystny. Programiści, inżynierowie operacyjni i inżynierowie ds. zapewnienia jakości mogą w rówiez  nauczyć się najlepszych rozwiazan, lepiej rozumiejąc DevOps. 

DevOps to zbior praktyk, które pomagają osiągnąć cel: skracaja czas między fazą powstawania pomysłu na produkt a jego wydaniem w formie produkcyjnej użytkownikowi końcowemu lub komukolwiek innemu, np. zespołowi wewnętrznemu lub klientowi. 

Kolejnym obszarem, w który zagłębimy się w tym pierwszym tygodniu, jest **metodologia Agile**. 
DevOps i Agile są powszechnie stosowane razem, aby osiągnąć ciągłość dostarczania (continuous delivery) **aplikacji**. 

The high-level takeaway is that a DevOps mindset or culture is about shrinking the long, drawn out software release process from potentially years to being able to drop smaller releases more frequently. The other key fundamental point to understand here is the responsibility of a DevOps engineer to break down silos between the teams I previously mentioned: Developers, Operations and QA.

Najważniejsze jest to, że mentalność i kultura DevOps polega na skróceniu długiego, ciągnącego się latami procesu wydawania oprogramowania do możliwości częstszego wydawania mniejszych wersji. Innym kluczowym, fundamentalnym punktem, który należy tutaj zrozumieć, jest odpowiedzialność inżyniera DevOps za przełamanie silosów pomiędzy zespołami, o których wspomniałem wcześniej: Developers, Operations i QA. 


Z perspektywy DevOps, **Rozwój, Testowanie i Wdrażanie** wszystkie te elementy są częścią zespołu DevOps.

Ostatnią kwestią, którą chciałbym poruszyć, jest to, że aby uczynić procesy tak skuteczne i efektywne, jak to tylko możliwe, musimy wykorzystać **Automatyzację**.

## Materiały

Jestem zawsze otwarty na dodawanie dodatkowych zasobów do tych plików readme, ponieważ służą one jako narzędzie do nauki.  

Radzę obejrzeć wszystkie poniższe filmy i mam nadzieję, że wyciągnęliście również wnioski z powyższych tekstów i objaśnień. 

- [DevOps in 5 Minutes](https://www.youtube.com/watch?v=Xrgk023l4lI)
- [What is DevOps? Easy Way](https://www.youtube.com/watch?v=_Gpe1Zn-1fE&t=43s)
- [DevOps roadmap 2022 | Success Roadmap 2022](https://www.youtube.com/watch?v=7l_n97Mt0ko)

Jeśli dotarłeś tak daleko, będziesz wiedział, czy to jest miejsce, w którym chcesz się znaleźć, czy nie. Do zobaczenia w [Day 2](day02.md).  
