#!/usr/bin/env bash

cd /app/ && CYPRESS_defaultCommandTimeout=10000 cypress run -q
