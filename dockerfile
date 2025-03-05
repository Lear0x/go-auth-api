# Étape 1 : Construction de l'application
FROM golang:1.20-alpine as builder

# Définir le répertoire de travail
WORKDIR /app

# Copier les fichiers de module Go
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source de l'application
COPY . .

# Compiler l'application
RUN go build -o main .

# Étape 2 : Créer l'image finale
FROM alpine:latest

# Installer les dépendances nécessaires
RUN apk --no-cache add ca-certificates

# Définir le répertoire de travail
WORKDIR /root/

# Copier l'exécutable depuis l'étape de build
COPY --from=builder /app/main .

# Copier le fichier .env (qui doit être dans le répertoire racine)
COPY .env .env

# Exposer le port sur lequel l'API sera accessible
EXPOSE 8080

# Démarrer l'application
CMD ["./main"]
