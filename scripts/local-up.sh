#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

ROOT=$(dirname "${BASH_SOURCE}")/..

# clean up local env.
function local-cleanup {
  test -d "${TMPDIR}" && rm -rf $TMPDIR
  test -n "${MONGO_PID-}" && ps -p $MONGO_PID > /dev/null && kill $MONGO_PID
  test -n "${TBB_PID-}" && ps -p $TBB_PID > /dev/null && kill $TBB_PID
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
godep go build -race .
./the-big-brother-is-watching-you -mock-path=./test/mockfile.txt -poll-period=1 &
TBB_PID=$!

while true; do sleep 10; done
