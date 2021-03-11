#!/bin/bash
#
# swagger.sh
#
# Copyright (c) 2019-2021 Junpei Kawamoto
#
# This software is released under the MIT License.
#
# http://opensource.org/licenses/mit-license.php
#
export GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w"

# Generate SiaStats client with Swagger
openapi-generator generate -i https://raw.githubusercontent.com/jkawamoto/siastats-swagger/master/swagger.yaml \
    -g go -o ./pkg/siastats --package-name siastats --git-user-id jkawamoto --git-repo-id go-siastats/pkg/siastats
