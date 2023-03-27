FROM flyway/flyway:8.5.13
COPY ./.docker/db/migrations/*.sql /flyway/sql/
COPY ./flyway.conf /flyway/conf/
ENTRYPOINT [ "flyway", "migrate" ]