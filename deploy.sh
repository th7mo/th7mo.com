#!/bin/bash

go run src/build.go
scp -r dist/* root@th7mo.com:/var/www/garden.th7mo/
