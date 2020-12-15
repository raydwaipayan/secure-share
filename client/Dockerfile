FROM node:lts-alpine

WORKDIR /app
COPY package.json ./
RUN apk add --no-cache curl git && cd /tmp && \
    curl -#L https://github.com/tj/node-prune/releases/download/v1.0.1/node-prune_1.0.1_linux_amd64.tar.gz | tar -xvzf- && \
    mv -v node-prune /usr/local/bin && rm -rvf

RUN npm install && node-prune && npm cache clean --force
RUN npm i -g serve

COPY . ./
RUN yarn build

ENV HOST 0.0.0.0
EXPOSE 3000
CMD ["serve", "-l", "3000", "-s", "dist"]