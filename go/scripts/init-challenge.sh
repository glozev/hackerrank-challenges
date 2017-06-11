#!/usr/bin/env bash

challenge=`echo $1 | tr '-' '_'`
current_dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
mkdir -p $challenge

pushd $challenge
    cat > "$challenge.go" << EOF
package main

import (
	"io"
	"os"
)

func Submission(stdin io.Reader, stdout io.Writer) {

}

func main() {
	Submission(os.Stdin, os.Stdout)
}
EOF
    ginkgo bootstrap -template $current_dir/test_suite_template.bootstrap
popd
