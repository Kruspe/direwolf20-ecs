#!/bin/sh
set -e

aws ecs update-service --cluster "$ECS_CLUSTER_ARN" --service "$ECS_SERVICE_ARN" --desired-count 0
