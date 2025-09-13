# Wizard API

Hello! This is a simple project for me to learn Kubernetes & sqlc.
It may not have the best practices, but it works! and thats whats important!

## Technologies

1. [sqlc](https://sqlc.dev/) - For generating type-safe code for interacting with the database
2. [Go](https://go.dev/) - The language this simple API is coded in
3. [Docker](https://docker.io/) - For deploying this app

## Endpoints

1. `GET /wizard` - Get all wizards
2. `GET /wizard/:id` - Return one individual wizard
3. `POST /wizard` - Put a wizard into the database
4. `DELETE /wizard/:id` - Delete a wizard from the database

## Deployment

### Kubernetes

It's as simple as running `kubectl apply -f ./deploy/k8s`!

### Docker

It's as simple as running `docker compose -f ./deploy/docker/compose.yml up`!