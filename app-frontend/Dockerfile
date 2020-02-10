FROM nginx:1.17.8-alpine

COPY ./build/nginx.conf /etc/nginx/conf.d/default.conf

WORKDIR /var/www/creeptoo
COPY ./src/* ./
