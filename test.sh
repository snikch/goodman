#!/bin/bash

go build -o bin/dredd-hooks-go

bundle exec cucumber $1
