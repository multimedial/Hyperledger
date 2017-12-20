FROM mysql/mysql-server
MAINTAINER SSE WZE <johannes.wenzel@schoenhofer.de>

COPY ./fabricexplorer.sql /docker-entrypoint-initdb.d

ENTRYPOINT ["/entrypoint.sh"]
ENTRYPOINT ["/entrypoint.sh"]
HEALTHCHECK CMD /healthcheck.sh
EXPOSE 3306 33060
CMD ["mysqld"]
