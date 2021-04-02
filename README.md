# gqlgen-todos

## GraphQL server in golang

### Steps: 

1.- generate schema from gqlgen-todos/graph/schema.graphqls
```
gqlgen generate
```

2.- generate resolvers: over graph/resolver.go run command
```
go:generate ./...
```