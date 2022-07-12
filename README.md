# cnpy-webapp

It's a demo web applicaiton built on micro-service archtecture approach.

## Content

It has two services `web` and `api` in respective directories. A directory `bin/`
contains additional tools required for setting up a build environment and interaction
with cloud infrastructure.

## Prerequisites

It requires:
* Linux tools/packages `curl`, `jq`, base development packages including `gmake`
* `buildah`, version ~> 1.24
* `awscli2`, version ~> 2.7
* `golang` compiler, version ~> 1.16 (Optional)
* Terraform, version ~> 1.2

The service are written by Go and as development dependency required Golang compiler.
For building working images the compiler isn't required. the `buildah` uses a golang
container image during a multi-stage image building.

*NB* It's supposed that the cloud infrastructure `cnpy-terraform-infra` is already deployed
and its directory are cloned in the same directory where cloned the current project.

## Build and Deploy The Services/Images

Before building the images of the services, a build enviroment should be preconfigured.
To configure the environment execute in working terminal a command:
```
$ eval $(bin/cnpy-build-environment.sh)
```
It set up required environment variables with values taken from terraform output.

Now, You can go to respective service directory (`web/` and `api/`) and start building.

An application verison are provided in `Makefile`. The version is injected to service
binaries and used for tagging container images. To raise up the version just update
a value of variable APP_VERSION in `Makefile`.

To build for development needs, use command `make`. It builds a binary and place it
in `build/` directory. But it requires a Golang compiler installed in your system.

To test building images, use command `make build-image`. If the image is built up
successfully you can build and push image to an ECRs provisioned by terraform
configuration `cnpy-terraform-infra`. To make it:
1. Make login to AWS ECR by command `../bin/cnpy-ecr-login.sh` (single time action. Usually
once per day)
2. Build the image and push it to the ECR by command `make push-image`

## Deployment Process

Proposed initial CI/CD process requires a presence images version 0.0.1
at respective ECRs. To provide this version:
1. Make check out a commit tagged 0.0.1 as a branch with command `git checkout -b initial 0.0.1` 
for the project
2. Ensure that the build environment is defined properly (it has environment variables started
with `CNPY_*` in `printenv` output) or use `eval $(bin/cnpy-build-environment.sh)` to update it
3. Build and push all images from respective directories with command `make push-image`

## Troubleshooting

if a pushing an image fails, ensure that you are logged in to the ECR by command 
`bin/cnpy-ecr-login.sh`

