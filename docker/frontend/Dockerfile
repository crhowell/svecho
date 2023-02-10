# pull official base image
FROM node:19-alpine

# look for node_modules 1 folder above where app is mounted
# https://burnedikt.com/dockerized-node-development-and-mounting-node-volumes/
ENV PATH=/srv/node_modules:$PATH

# set work directory
WORKDIR /srv/
COPY ./frontend/package.json ./
COPY ./docker/frontend/entrypoint.sh ./entrypoint.sh

WORKDIR /srv/app/

COPY ./frontend/ ./

RUN chmod +x /srv/entrypoint.sh

ENTRYPOINT ["/srv/entrypoint.sh"]
