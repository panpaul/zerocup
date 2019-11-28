FROM node:10-alpine

RUN mkdir -p /src/nuxt-app
WORKDIR /src/nuxt-app

COPY . /src/nuxt-app/
RUN npm config set registry https://registry.npm.taobao.org
RUN npm install

RUN npm run build

EXPOSE 3000
ENV NUXT_HOST 0.0.0.0
ENV NUXT_PORT 3000
CMD [ "npm", "run", "start"]