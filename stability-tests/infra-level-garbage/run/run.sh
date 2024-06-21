#!/bin/bash
rm -rf /tmp/nirvanad-temp

nirvanad --devnet --appdir=/tmp/nirvanad-temp --profile=6061 &
NIRVANAD_PID=$!

sleep 1

infra-level-garbage --devnet -alocalhost:42611 -m messages.dat --profile=7000
TEST_EXIT_CODE=$?

kill $NIRVANAD_PID

wait $NIRVANAD_PID
NIRVANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nirvanad exit code: $NIRVANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NIRVANAD_EXIT_CODE -eq 0 ]; then
  echo "infra-level-garbage test: PASSED"
  exit 0
fi
echo "infra-level-garbage test: FAILED"
exit 1
