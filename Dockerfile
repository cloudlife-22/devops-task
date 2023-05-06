# syntax=docker/dockerfile:1
# https://docs.docker.com/language/golang/build-images/
# https://snyk.io/blog/containerizing-go-applications-with-docker/

##
## Build the application from source
##

FROM golang:1.20 AS build

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /devops-task
# RUN CGO_ENABLED=0 GOOS=linux go build -o -ldflags=-X=main.version=$(git rev-parse --short HEAD) /devops-task
# RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go

##
## Run the tests in the container
##

FROM build AS run-test-stage

RUN go vet -v
RUN go test -v
RUN go test -cover

# #
# Deploy the application binary into a lean image
# #

FROM gcr.io/distroless/static-debian11
COPY --from=build /devops-task /devops-task

EXPOSE 8080
EXPOSE 2112

# distroless nonroot user https://github.com/GoogleContainerTools/distroless/tree/main/examples/nonroot
USER nonroot:nonroot


ENTRYPOINT ["/devops-task"]

# ENV PATH="/go/bin:${PATH}"
# CMD ["/main"]