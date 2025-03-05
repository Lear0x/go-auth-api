
# README - Dockerisation et Accès à PostgreSQL via pgAdmin

## Prérequis

- Docker et Docker Compose installés.

## Étapes

1. **Lancer les services** :
   ```bash
   docker-compose up --build
   ```
   - Application Go : `http://localhost:8080`
   - pgAdmin : `http://localhost:5050`
   - PostgreSQL : `localhost:5432`

2. **Accéder à pgAdmin** :
   - Connexion :  
     **Email** : `admin@admin.com`  
     **Mot de passe** : `admin`
   
3. **Ajouter un serveur PostgreSQL dans pgAdmin** :
   - **Host** : `db`
   - **Port** : `5432`
   - **Nom d'utilisateur** : `admin`
   - **Mot de passe** : `pass`
   - **Base de données** : `authdb`

4. **Arrêter les services** :
   ```bash
   docker-compose down
   ```

---
