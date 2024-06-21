#!/bin/bash

APPDIR=/tmp/nirvanad-temp
NIRVANAD_RPC_PORT=29587

rm -rf "${APPDIR}"

nirvanad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${NIRVANAD_RPC_PORT}" --profile=6061 &
NIRVANAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${NIRVANAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $NIRVANAD_PID

wait $NIRVANAD_PID
NIRVANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nirvanad exit code: $NIRVANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NIRVANAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
