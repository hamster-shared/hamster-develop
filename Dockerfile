FROM hamstershare/debian_docker_cli:latest

COPY ./aline-test /usr/local/bin/aline-test

ENV PORT=8080
ENV DB_USER=root
ENV DB_PASSWORD=123456
ENV DB_HOST=127.0.0.1
ENV DB_PORT=3306
ENV DB_NAME=aline
EXPOSE ${PORT}
CMD /usr/local/bin/aline-test daemon -p ${PORT} --db_user ${DB_USER} --db_password ${DB_PASSWORD} --db_host ${DB_HOST} --db_port ${DB_PORT} --db_name ${DB_NAME}
