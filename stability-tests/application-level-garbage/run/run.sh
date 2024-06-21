#!/bin/bash
rm -rf /tmp/nirvanad-temp

nirvanad --devnet --appdir=/tmp/nirvanad-temp --profile=6061 --loglevel=debug &
NIRVANAD_PID=$!
NIRVANAD_KILLED=0
function killNirvanadIfNotKilled() {
    if [ $NIRVANAD_KILLED -eq 0 ]; then
      kill $NIRVANAD_PID
    fi
}
trap "killNirvanadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:42611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $NIRVANAD_PID

wait $NIRVANAD_PID
NIRVANAD_KILLED=1
NIRVANAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "Nirvanad exit code: $NIRVANAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $NIRVANAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
