# build stage
FROM node:15.2-alpine3.10 as build-stage

# INFO: for server
ENV https_proxy http://10.200.81.56:3128/

LABEL maintainer="Alex Mogilevskiy <dark.qvark@gmail.com>"

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# production stage
FROM nginx:1.19.4-alpine as production-stage

COPY --from=build-stage /app/dist /usr/share/nginx/html

#EXPOSE 8000

CMD ["nginx", "-g", "daemon off;"]
