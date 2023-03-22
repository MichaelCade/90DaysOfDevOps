---
title: '#90DaysOfDevOps - Zadania DevOps Engineera - Dzień 2'
published: false
description: 90DaysOfDevOps - Zadania DevOps Engineera - Dzień 2
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048699
date: '2022-04-17T21:15:34Z'
---
## Zadania DevOps Engineera - Dzień 2

Mam nadzieję, że wracasz tutaj po przerobieniu tekstu z poprzedniego dnia [Dzień 1 #90DaysOfDevOps](day02.md)

W pierwszym poście poruszyliśmy podstawowe zagadnienia, lecz teraz musimy przyjrzeć się dokładniej poszczególnym konceptom i zrozumieć, że na wytwarzanie oprogramowania składają się dwa obszary. **Development (rozwój)**, gdzie programiści(developerzy) tworzą i testują aplikację. Oraz **Operations (utrzymanie)**, gdzie aplikacja jest wdrażana i utrzymywana na serwerze.

## DevOps to pomost pomiędzy nimi

Aby zrozumieć DevOps i zadania, za które Inżynier DevOps jest odpowiedzialny, musimy zrozumieć narzędzia lub procesy związane z tym, jak te dwa światy się łączą.

Wszystko zaczyna się od aplikacji! W trakcie nauki zrozumiesz, że tak naprawdę wszystko co dotyczy DevOpsu tyczy się aplikacji.

Developerzy tworzą oprogramowanie, mogą to robić z użyciem wielu różnych stacków technologicznych, w późniejszych rozdziałach poświęcimy temu procesowi więcej uwagi. Wytwarzanie oprogramowania może wiązać się z użyciem wielu różnych języków programowania, narzędzi do budowania aplikacji, repozytoriów kodu itd.

Jako DevOps engineer nie będziesz programować samej aplikacji, ale dobre zrozumienie konceptów, z którymi pracują developerzy, oraz zrozumienie systemów, narzędzi i procesów z jakich korzystają jest kluczem do osiągnięcia sukcesu.

Wysokopoziomowo, będziesz potrzebować wiedzieć o tym jak aplikacja jest skonfigurowana, by móc rozmawiać z wszystkimi wymaganymi serwisami, oraz znać wymogi tego, jak taką komunikację testować.

Aplikacja musi być gdzieś wdrożona, dla uproszczenia przyjmijmy, że jest to serwer, nieważne jaki, po prostu serwer. Zakładamy również, że ten serwer powinien być dostępny dla klienta lub użytkowanika końcowego aplikacji.

Ten serwer musi istnieć w jakiejś określonej konfiguracji – on premises, w publicznym cloudzie, serverless (Ok, wybiegamy tutaj trochę za daleko, ten projekt nie pokrywa serverless, jednak jest to jedna z możliwości i coraz więcej firm wybiera ten model). Ktoś musi te serwery stworzyć i skonfigurować, aby przygotować je do hostowania aplikacji. Ten obszar, skonfigurowania i wdrożenia tych serwerów, może leżeć w zakresie twoich obowiązków jako DevOps Engineera.

Na tych serwerach uruchomiony jest jakiś system operacyjny, zazwyczaj jest to Linux, poświęcimy całą sekcję/tydzień na pokrycie podstawowej wiedzy, jaką powinienieś/aś zdobyć w tym obszarze.

Prawdopodobne jest również, że dany serwer będzie potrzebował komunikować się z innymi serwisami w naszej sieci lub środowisku, więc będziemy również potrzebować odpowiedniej wiedzy z zakresu sieci i ich konfiguracji, jako, że ten obszar również będzie znajdować się w kometencjach DevOps Engineera. Bardziej szczegółowe informacje znajdziesz w odpowiednich sekcjach, gdzie będizemy mówić o wszystkim związanym z DNSami, DHCP, Load Balancigiem itd.

## Jack of all trades, Master of none 

Na tym etapie należałoby zaznaczyć, że nie potrzebujesz eksperckiej wiedzy z zakresu sieci, czy infrastruktury. Potrzebujesz jedynie podstawowej wiedzy na temat tego, jak skonfigurować wszystko co jest potrzebne, żeby poszczególne zasoby mogły działać i rozmawiać ze sobą nawzajem. Podobnie jest z wiedzą w zakresie języków programowania. Nie musisz być developerem. Jednak możesz przejść do DevOpsu jako specjalista w danym obszarze i możę to być świetna podstawa do zaadaptowania do innych obszarów.

Prawdopodobnie również nie przejmiesz odpowiedzialności w kwestii codziennego zarządzania serwerami lub aplikacjami.

Mówiliśmy o sewerach, ale prawdopodobnie twoja aplikacja będzie tworzona z myślą o działaniu w kontenerach, które również zazwyczaj działają na serwerze, ale w tym przypadku również musisz rozumieć koncepty takie jak wirtualizacja, chmurowe Infrastructure as a Service (IaaS), ale również konteneryzacji. Podczas tych 90 dni skupimy się przede wszystkim na kontenerach.

## Podsumowanie wysokopoziomowe

Z jednej strony mamy developerów, tworzących nowe ficzery i funkcjonalności (jak i naprawiających bugi) dla danej aplikacji.

Z drugiej strony, mamy jakiegoś rodzaju środowisko, infrastrukturę i serwery, które należy skonfigurować i którymi trzeba zarządzać, aby były w stanie obsługiwać tę aplikację i komunikować się ze wszystkimi wymaganymi usługami i serwisami

W takiej sytuacji należy zadać sobie pytanie, w jaki sposób chcemy zaimplementować te funkcjonalności i poprawki do naszego produktu i sprawić, by były dostępne dla użytkowników końcowych?

W jaki sposób chcemy wypuścić nową wersję aplikacji? To jedno z głównych zadań inżyniera DevOps, co ważne, nie polega to jedynie na znalezieniu jednorazowego rozwiązania, ale musimy robić to w sposób ciągły zautomatyzowany i efektywny, oraz włączyć do całego procesu testowanie tego oprogramowania.

Na tym etapie zakończymy dzisiejszy dzień nauki, mając nadzieję, że ta wiedza była przydatna. Następne kilka dni poświęcimy na głębsze spojrzenie na poszczególne obszary DevOps oraz spojrzymy dokładniej na narzędzia, procesy i korzyści z takich rozwiązań.

## Dodatkowe źródła

Zawsze jestem otwarty na dodawanie dodatkowych źródeł do tych plików readme, jako, że służą one pogłębieniu wiedzy.

Moja rada to obejrzenie wszystkich poniższych filmów i mam nadzieję, że również wyciągnęliście jakieś użyteczne informacje z powyższego tekstu

- [What is DevOps? - TechWorld with Nana](https://www.youtube.com/watch?v=0yWAtQ6wYNM)
- [What is DevOps? - GitHub YouTube](https://www.youtube.com/watch?v=kBV8gPVZNEE)
- [What is DevOps? - IBM YouTube](https://www.youtube.com/watch?v=UbtB4sMaaNM)
- [What is DevOps? - AWS ](https://aws.amazon.com/devops/what-is-devops/)
- [What is DevOps? - Microsoft](https://docs.microsoft.com/en-us/devops/what-is-devops)

Jeżeli dotarłeś/aś do tego momentu, już wiesz, czy chcesz iść dalej tą drogą. Do zoabaczenia w [Dniu 3](day03.md).  
