FROM postgres:14.10-alpine3.19

# You can configure these envs and adjust to what are in params/.env file
ENV POSTGRES_DB=db-test
ENV POSTGRES_USER=username-test
ENV POSTGRES_PASSWORD=password-test

# Comment this line below to disable auto employees table creation & it seed if you want to use migration from the app for employees table
COPY migrations/seeds/init.sql /docker-entrypoint-initdb.d

EXPOSE 5432