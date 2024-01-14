#!/bin/sh

npm cache clean --force
npm install
npm run build

exec "$@"
