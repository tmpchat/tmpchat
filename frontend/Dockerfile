FROM node:14-alpine

WORKDIR /app

# node-gyp needs Python.
# Please see https://github.com/nodejs/node-gyp/issues/1977
RUN apk add --no-cache make g++ python3

COPY package.json yarn.lock ./

RUN npm install

COPY . .

EXPOSE 8080

CMD ["npm", "run", "serve"]
