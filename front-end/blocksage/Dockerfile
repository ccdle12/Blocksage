# This Dockerfile is suitable for DEVELOPMENT ONLY!
# Package.json was edited to include: webpack-dev-server --host 0.0.0.0
# ^ This is NOT SUITABLE FOR PRODUCTION!

# FROM node:9.11.1-alpine
FROM node:latest

WORKDIR /app/blocksage
COPY ./ .

RUN npm install \
    && npm run build

CMD npm run dev