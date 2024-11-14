# API de Usuarios Aleatorios

Este proyecto es una API en Go que consume el servicio de [RandomUser.me](https://randomuser.me/) para obtener datos aleatorios de usuarios. La API filtra y devuelve información específica de 15000 usuarios únicos en formato JSON.

## Características

- Consume datos de usuarios de la API de RandomUser.me
- Filtra la respuesta para devolver solo:
  - Género
  - Primer nombre
  - Primer apellido
  - Email
  - UUID (único para cada usuario)
- Elimina duplicados en función del UUID y garantiza que solo se devuelvan 15 usuarios únicos
- La API responde en menos de 2.25 segundos

## Tecnologías

- Go
- Gin Web Framework
