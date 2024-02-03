#!/bin/bash

npm run build
scp -r dist/* root@th7mo.com:/var/www/garden.th7mo/
