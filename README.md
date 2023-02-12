# jamify-service-template
Base template for creating new services in Jamify

Please follow the checklist when creating new services

- [ ] Create the repo from the template
- [ ] Rename the container name in the `docker-compose.yaml`
- [ ] Change the port exposed in the `Dockerfile`
- [ ] Change the ports in the  `docker-compose.yaml`
- [ ] Change the module in `go.mod` to point to your new service instead of the template
- [ ] Change the prog name in the `Makefile`
- [ ] Change name of package in `.vscode/launch.json`
- [ ] Add branch protection rules to your new repo (require 1 approval, no pushing to main, etc)
