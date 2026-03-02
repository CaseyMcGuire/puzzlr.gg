## Commands
To regenerate ent models:
```
go generate ./...
```

## Testing
Default tests (no Docker required):
```
go test ./...
```

Integration Tests:
- Docker must be installed. You can install directly from [here](https://docs.docker.com/get-docker/).

How to run integration tests:
```
./bin/test-integration.sh
```

To run a single integration test:
```
./bin/test-integration.sh -run TestCreateGameResolverSuccess
```

