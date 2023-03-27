FROM flyway/flyway:7.15.0
COPY ./.docker/db/migrations/*.sql /flyway/sql/
COPY ./flyway.conf /flyway/conf/
ENTRYPOINT [ "flyway", "migrate" ]