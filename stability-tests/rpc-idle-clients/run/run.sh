#!/bin/bash
rm -rf /tmp/nirvanad-temp

NUM_CLIENTS=128
nirvanad --devnet --appdir=/tmp/nirvanad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
NIRVANAD_PID=$!
NIRVANAD_KILLED=0
function killNirvanadIfNotKilled() {
  if [ $NIRVANAD_KILLED -eq 0 ]; then
    kill $NIRVANAD_PID
  fi
}
trap "killNirvanadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $NIRVANAD_PID

wait $NIRVANAD_PID
NIRVANAD_EXIT_CODE=$?
NIRVANAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "Nirvanad exit code: $NIRVANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NIRVANAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
