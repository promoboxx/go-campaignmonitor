#!/bin/bash

if [ -z ${1+x} ]; then echo "package name is not set.\n sh init.sh <packageName>\n"; exit 1; else echo "using package name '$1'"; fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"

echo "replacing <packageName> with $1..."
LC_CTYPE=C && LANG=C && find . -type f -print0 | xargs -0 sed -i '' -e "s/<packageName>/$1/g"

govendor init

echo "\nFinished.  Don't forget to update the README.md.\n"

# deletes this script
rm -- "$0"
