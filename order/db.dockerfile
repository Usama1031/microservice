FROM postgres:17.5

COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]