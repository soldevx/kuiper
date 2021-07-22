FROM nginx:1.21.1

LABEL maintainer="EPI"
LABEL description="image supporting web server nginx"
LABEL version="1.0"

# Copy custom nginx config
COPY ./websrv/nginx.conf /etc/nginx/nginx.conf
COPY ./websrv/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80 443

ENTRYPOINT ["nginx"]
CMD ["-g", "daemon off;"]