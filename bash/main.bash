#!/bin/bash

while true; do
  go run cmd/main.go

  if [ $? -eq 0 ]; then
    echo "executing..."
  else
    echo "interrupted"
    sleep 1
  fi
done