#!/usr/bin/env bash

#https://stackoverflow.com/questions/59895/how-to-get-the-source-directory-of-a-bash-script-from-within-the-script-itself
CA_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

cp ${CA_DIR}/*.crt /usr/local/share/ca-certificates/
chmod 644 /usr/local/share/ca-certificates/*.crt
update-ca-certificates
