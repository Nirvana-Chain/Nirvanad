#!/bin/bash
rm -rf /tmp/nirvanad-temp

nirvanad --simnet --appdir=/tmp/nirvanad-temp --profile=6061 &
NIRVANAD_PID=$!

sleep 1

orphans --simnet -alocalhost:42511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $NIRVANAD_PID

wait $NIRVANAD_PID
NIRVANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nirvanad exit code: $NIRVANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NIRVANAD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
