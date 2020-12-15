# INFO: build stage
FROM node:15.2-alpine3.10 as build-stage

# INFO: for production on server
ENV https_proxy http://10.200.81.56:3128/

LABEL maintainer="Alex Mogilevskiy <dark.qvark@gmail.com>"

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# INFO: production stage
FROM nginx:1.19-alpine as production-stage

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY --from=build-stage /app/dist /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]
