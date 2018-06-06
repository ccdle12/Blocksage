FROM node:alpine

COPY ./ /api

WORKDIR ./api

RUN apk add --update bash \ 
    && rm -rf node_modules \
    && npm i \
    && npm i nodemon -g

CMD nodemon index.js