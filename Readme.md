## Intro

This repository is to save small projects then I made to pratice my knowledges.

## DOCKER

```cmd
docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=Pass1234" `-p 1433:1433 --name sql1 --hostname sql1 -d mcr.microsoft.com/mssql/server:2022-latest

docker exec -it <container_id> /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P Pass1234
```

## NOTE

- For each project it has a separate branch.
- The branch name needs to be named with the programming language as suffix

##
Icons images by: [IconScout](https://iconscout.com/)