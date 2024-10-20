#!/bin/bash

vagrant destroy -f
rm -rf .vagrant
vagrant global-status --prune
vagrant up -vvvv