# Proxy for address suggestion requests

## What is it

This service solves the problem of caching responses from address suggestion services. Services that provide address suggestions based on partial input have limitations on the number of requests, and when the limits are exceeded, they require switching to a paid subscription plan.

This service acts as a proxy by searching for data in its cache. If the data is not found, it sends requests to all available services and caches the responses for an extended period. This approach significantly reduces the number of requests and improves overall performance.

Available providers:
- Yandex Suggestion

Planned providers:
- Dadata
- 2Gis

## How to deploy

- Create and edit config file `cp config.example.toml config.toml`
- Edit docker-compose.yml
- `docker compose up`

## Contribute
You may feel free to join the development or make feature requests for this service. I have created this service to resolve my own needs, but I have designed it with extensibility in mind, so everything should be easy.

If you want to contribute but don't know where to start, here are some implementation ideas:

- Add other suggestion providers, such as [Datata](https://dadata.ru/suggestions/#address) or [2GIS](https://docs.2gis.com/en/api/search/suggest/reference/3.0/suggests#/paths/~13.0~1suggests/get)
- Include additional caching services besides Redis, such as [MongoDB](https://www.mongodb.com/) or [BadgerDB](https://dgraph.io/docs/badger/)
- Devise an efficient method to select the best data from different providers
- Enhance the external API requests to take user location into account

## How to develop

Create `config.toml` file and edit settings

```shell
cp config.example.toml config.toml
```

You can start cache in docker and run application on your host:
```shell
docker compose up cache
clear; go run ./
```

Or can star all via docker compose. Don't forget edit redis host in config.toml, if you run application via docker compose host must be "cache" 
```shell
make build-docker && docker compose up
```