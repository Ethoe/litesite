FROM flyway/flyway:8.0.4
COPY ./.docker/db/migrations/*.sql /flyway/sql/
COPY ./flyway.conf /flyway/conf/
ENTRYPOINT [ "flyway", "migrate" ]