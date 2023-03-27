FROM flyway/flyway
COPY ./.docker/db/migrations/*.sql /flyway/sql/
COPY ./flyway.conf /flyway/conf/
ENTRYPOINT [ "flyway", "migrate" ]