#!/bin/bash

redis-server ./redis.conf --port 7001 --cluster-config-file node1.conf &
pid1=$!
redis-server ./redis.conf --port 7002 --cluster-config-file node2.conf &
pid2=$!

trap "kill $pid1 $pid2" EXIT

redis-server ./redis.conf --port 7003 --cluster-config-file node3.conf
