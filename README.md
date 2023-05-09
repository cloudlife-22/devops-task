# devops-task
Exercise showcasing some technical abilities


## Requirements 
1. The source code should be available on a public github repository or a zip/tarball of a git repository via email. √
2. The repository must have README.md highlighting how to run the solution. √
3. Create a 'hello world' http server that returns 200 OK on port 8080. You can use something simple like a [golang example](https://yourbasic.org/golang/http-server-example/) or just the default nginx docker container. √
4. Package this into a Docker container https://github.com/cloudlife-22/fleet-infra √
5. Push it to any container registry √
6. Deploy the container to a orchestration platform of your choosing (GKE) √

## Bonus Points
1. Container security scanning √
2. CD pipeline or gitops tooling √

## Solution Overview

For this task I have created a basic Go app using [golang example](https://yourbasic.org/golang/http-server-example/). I have also written some basic tests alongside the application and exposed some metrics using prometheus.

The solution builds using github actions and the pipeline is triggered when there is a PR or Commit to the main branch (it is ignored if changes are made to the contents of the kustomize directory as these are picked up by fluxcd). If a PR to the main branch is created only the build-and-test job will be run, which build the go app, run tests and print a test coverage report. If a commit to main occurs a build-image-test-scan-and-push job is run in parallel to the build-and-test job, in this case a docker image is built using a multistage builds and finally packaged into a distroless image, before being checked for vulnerabilities using Trivy (and of course ignoring them :/) before pushing them to dockerhub: https://hub.docker.com/repository/docker/khanosal/devops-task/general. 

Once a new image is pushed it is automatically deployed to the k8s cluster. Any changes to the contents of the kustomize directory are picked up via fluxcd which is configured here: https://github.com/cloudlife-22/fleet-infra (This is still WIP and will likely be updated).

By no means does this solution demonstrate a production ready deployment pipeline but rather serves to demonstrate multiple features of a potential solution. 

## How to run this solution 

make changes to the go app main.go file create a pr to main, watch your pipeline build here: https://github.com/cloudlife-22/devops-task/actions. Once your PR is merged watch the new pipeline in the same place until docker image is pushed, once docker image is pushed it will be deployed to k8s. If all goes well the likely url will be https://devops-task.app otherwise I flopped `\_("/)_/` or just didn't have enough time :P. This is not a complete solution and there may be further changes.  