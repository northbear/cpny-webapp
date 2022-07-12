#!/usr/bin/env bash

## set -xv

BASE_DIR=$( cd `dirname $0`; pwd )
WORK_DIR=$( realpath "${BASE_DIR}/.." )

CNPY_REGION=${AWS_REGION:-eu-west-1}

# aws ecr get-login-password --region eu-west-1 | buildah login --username AWS --password-stdin 203511311076.dkr.ecr.eu-west-1.amazonaws.com
aws ecr get-login-password --region "${CNPY_REGION}" | \
    buildah login --username AWS --password-stdin "${CNPY_ECR_REPOSITORY}"
