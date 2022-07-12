#!/usr/bin/env bash

## set -xv

BASE_DIR=$( cd `dirname $0`; pwd )
WORK_DIR=$( realpath "${BASE_DIR}/.." )

CNPY_TF_CONF=$( realpath "${WORK_DIR}/../cnpy-terraform-infra" )

[[ -d "${CNPY_TF_CONF}" ]] || { echo "Error: a TF configuration doesn't exist: ${CNPY_RF_CONF}"; exit 1; }

ECR_REGISTRY=$( cd "${CNPY_TF_CONF}"; terraform output -json | \
                    jq -r '.cnpy_ecr_urls.value[0]' | \
                    awk -F '/' '{ print $1}' )
REGION=$( echo ${ECR_REGISTRY} | awk -F . '{ print $4 }' )

echo """\
export CNPY_ECR_REPOSITORY=${ECR_REGISTRY}
export CNPY_REGION=${REGION}
export CNPY_APP_NAME=webapp
"""
