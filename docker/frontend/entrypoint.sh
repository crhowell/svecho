#!/bin/sh

cd ..
npm install && npm cache clean --force

cd app
npm run build

exec "$@"
