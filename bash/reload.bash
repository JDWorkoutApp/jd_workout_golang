#!/bin/bash

PID=$(pgrep -f "go run cmd/main.go")

if [ -n "$PID" ]; then
  echo "PID: $PID"
  kill $(lsof -t -i:80)
else
  echo "Pid not found"
fi