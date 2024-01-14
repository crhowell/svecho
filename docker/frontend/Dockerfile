# pull official base image
FROM node:20-alpine

# set work directory
WORKDIR /srv/
COPY ./docker/frontend/entrypoint.sh ./entrypoint.sh
RUN chmod +x /srv/entrypoint.sh

WORKDIR /srv/app/
COPY ./frontend/ ./

ENTRYPOINT ["/srv/entrypoint.sh"]
