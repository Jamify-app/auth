# jamify-service-template
Base template for creating new services in Jamify

Please follow the checklist when creating new services

- [x] Create the repo from the template
- [x] Rename the container name in the `docker-compose.yaml`
- [x] Change the port exposed in the `Dockerfile`
- [x] Change the ports in the  `docker-compose.yaml`
- [x] Change the module in `go.mod` to point to your new service instead of the template
- [x] Change the prog name in the `Makefile`
- [x] Change name of package in `.vscode/launch.json`
- [x] Add branch protection rules to your new repo (require 1 approval, no pushing to main, etc)
