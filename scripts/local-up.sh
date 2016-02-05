#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(dirname "${BASH_SOURCE}")/..

# clean up local env.
function local-cleanup {
  echo "rm tmp dir"
  test -d "${TMPDIR}" && rm -rf $TMPDIR
  echo "kill mongod process"
  test -n "${MONGO_PID-}" && ps -p $MONGO_PID > /dev/null && kill $MONGO_PID
  echo "kill tbb process"
  test -n "${TBB_PID-}" && ps -p $TBB_PID > /dev/null && kill $TBB_PID
  echo "rm all the tmp and processes"
}

trap local-cleanup INT EXIT

# Run mongo.
TMPDIR=`mktemp -d /tmp/tbb.XXXXXXXXXX`
mkdir ${TMPDIR}/db
mongod -dbpath "${TMPDIR}/db" > ${TMPDIR}/tbb-mongo.log 2>&1 &
MONGO_PID=$!

echo "-> mongodb log in ${TMPDIR}/tbb-mongo.log"

# godep go build -race .
cd $ROOT
./the-big-brother-is-watching-you -mock-path=./test/mockfile.txt &
TBB_PID=$!

while true; do sleep 10; done
