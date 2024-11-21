# Groupie Trackers

Groupie Trackers est un projet qui utilise une API pour manipuler et afficher des informations sur différents groupes et artistes.

## API

L'API se compose de quatre parties :

1. **Artistes** : Contient des informations sur certains groupes et artistes comme leur(s) nom(s), image, l'année où ils ont commencé leur activité, la date de leur premier album et les membres.
2. **Lieux** : Comprend leurs derniers lieux de concert et/ou les lieux des prochains concerts.
3. **Dates** : Comprend leurs dernières dates de concert et/ou les prochaines dates de concert.
4. **Relation** : Fait le lien entre toutes les autres parties, artistes, dates et lieux.

## Objectif

L'objectif est de construire un site web convivial où vous pouvez afficher les informations des groupes à travers plusieurs visualisations de données (exemples : blocs, cartes, tableaux, listes, pages, graphiques, etc). C'est à vous de décider comment vous allez l'afficher.

Ce projet se concentre également sur la création d'événements/actions et sur leur visualisation.

## Événements / Actions

L'événement/action que nous voulons que vous fassiez est connu sous le nom d'appel client au serveur (client-serveur). On peut dire que c'est une fonctionnalité de votre choix qui doit déclencher une action. Cette action doit communiquer avec le serveur afin de recevoir des informations, ([request-response])(https://en.wikipedia.org/wiki/Request%E2%80%93response)

## Instructions

- Le backend doit être écrit en Go.
- Le site et le serveur ne peuvent pas planter à aucun moment.
- Toutes les pages doivent fonctionner correctement et vous devez prendre soin de toutes les erreurs.
- Le code doit respecter les bonnes pratiques.
- Il est recommandé d'avoir des fichiers de test pour les tests unitaires.

## Packages autorisés

Seuls les packages standard Go sont autorisés.

## Utilisation

Vous pouvez voir un exemple d'une API RESTful ici

## Ce que ce projet nous a apporté :

- Manipulation et stockage des données.
- Fichiers et format JSON.
- HTML.
- Création et affichage d'événements.
- Client-serveur.

## Docker

Docker a été ajouté à ce projet pour faciliter le déploiement et assurer la cohérence de l'environnement de développement. Docker permet de créer un conteneur qui contient toutes les dépendances nécessaires pour exécuter l'application, garantissant ainsi que l'application fonctionne de la même manière sur toutes les machines.

### Pour construire l'image Docker, utilisez la commande suivante :

docker build -t groupie .

### Pour exécuter l'application à l'aide de Docker, utilisez la commande suivante :

docker run -p 3000:8080 -d groupie
