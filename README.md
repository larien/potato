# potato

This project is a Go skelethon to be used to develop new microservices. I'm still not sure what I'm gonna do with this in the future; the goal for now is to finish it and write down all the topics to be learn in order to understand what takes to build a Go microservice. Enjoy!

## TODO

- [X] 1. Receive an answer from an endpoint
- [X] 2. Run tests locally
- [X] 3. Run app locally
- [X] 4. Submit aplication to container registry
- [X] 5. Improve configuration
- [X] 6. Create router
- [X] 7. Expose CRUD endpoints
- [X] 8. Certify server health
- [X] 9. Get parameters from request
- [ ] 10. Isolate business logic
- [ ] 11. Store data
- [ ] 12. Enable admin permissions
- [ ] 13. Enable server debugging
- [ ] 14. Monitor requests
- [ ] ?. Prevent other services from explosion
- [ ] ?. Enable profiling
- [ ] ?. Improve response performance

## Learning topics

1. Receive an answer from an endpoint
   - Understanding the HTTP server implementation
2. Run tests locally
   - Table driven tests
   - httptest
   - Makefile
   - go mod
3. Run app locally
   - N/A
4. Deploy application
   - Container registry
   - GitHub Actions
   - Docker
   - YAML
   - lint
5. Improve configuration
   - Configuration files
   - Viper
   - YAML/TOML/JSON
   - flag package
6. Create router
   - Mux
   - Handler
   - Why and how to build a router
   - Obtain route variables
7. Expose CRUD endpoint
   - CRUD
   - Lock
   - HTTP methods, requests, responses
8. Certify server health
   - Healthcheck
9. Get parameters from request
   - Query params