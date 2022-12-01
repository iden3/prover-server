#!/usr/bin/env sh

if [ $# -ne 0 ];
then
  # some command provided - execute it instead of prover
  exec "$@"
else
  PROVER_CMD="/home/app/prover"

  # detect support for adx instruction set by current cpu
  HAS_ADX=`cat /proc/cpuinfo | grep adx | wc -l`

  if [ $HAS_ADX == "0" ];
  then
    PROVER_CMD="/home/app/prover_noasm"
    echo "No adx support detected, falling back to prover without asm optimizations"
  fi
  exec "$PROVER_CMD"
fi
