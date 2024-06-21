#!/bin/bash
rm -rf /tmp/nirvanad-temp

nirvanad --devnet --appdir=/tmp/nirvanad-temp --profile=6061 --loglevel=debug &
NIRVANAD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $NIRVANAD_PID

wait $NIRVANAD_PID
NIRVANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nirvanad exit code: $NIRVANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NIRVANAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
