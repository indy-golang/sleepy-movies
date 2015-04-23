#!/usr/bin/env bash
set -x
mongo sleepy-movies --eval "db.dropDatabase()"