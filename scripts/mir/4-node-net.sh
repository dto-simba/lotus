#!/usr/bin/env bash

n=`date '+%Y-%m-%dT%T'`
export MIR_INTERCEPTOR_OUTPUT="mir-event-logs/run-$n"

tmux new-session -d -s "mir" \; \
  new-window   -t "mir" \; \
  split-window -t "mir:0" -v \; \
  split-window -t "mir:0.0" -h \; \
  split-window -t "mir:0.2" -h \; \
  \
  split-window -t "mir:1" -v \; \
  split-window -t "mir:1.0" -h \; \
  split-window -t "mir:1.2" -h \; \
  \
  send-keys -t "mir:0.0" "
        ./scripts/mir/daemon.sh 0" Enter \; \
  send-keys -t "mir:0.1" "
        ./scripts/mir/daemon.sh 1" Enter \; \
  send-keys -t "mir:0.2" "
        ./scripts/mir/daemon.sh 2" Enter \; \
  send-keys -t "mir:0.3" "
        ./scripts/mir/daemon.sh 3" Enter \; \
  \
  send-keys -t "mir:1.0" "
        ./scripts/mir/validator.sh 0 2>&1 | tee validator_0.log" Enter \; \
  send-keys -t "mir:1.1" "
        ./scripts/mir/validator.sh 1 2>&1 | tee validator_1.log" Enter \; \
  send-keys -t "mir:1.2" "
        ./scripts/mir/validator.sh 2 2>&1 | tee validator_2.log" Enter \; \
  send-keys -t "mir:1.3" "
        ./scripts/mir/validator.sh 3 2>&1 | tee validator_3.log" Enter \; \
 attach-session -t "mir:0.3"

 pkill -f lotus
