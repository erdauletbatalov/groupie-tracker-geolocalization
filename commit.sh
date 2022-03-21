#!/bin/bash
git add .
git commit -m "$1"
export GIT_SSL_NO_VERIFY=0
git push