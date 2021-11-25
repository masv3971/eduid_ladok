# eduid_ladok
Integration between eduid and ladok

## Prerequisites
1. Create account in schools local IDP
2. Ask Ladok to create a certificate for the above user
    * The certificate needs to have at least the following permissions:
    * 


## Build and run in Docker-compose
1. docker-compose build
2. docker-compose up -d
## Build outside Docker
```
# Linux
$ make

# mac m1
$ make mac_m1
```

## Flow
This is the general flow for each school using this.

```
graph LR;
    ladok -->|atom| eduid_ladok;
    ladok -->|rest| eduid_ladok;
    eduid_ladok -->|rest| ladok;

    eduid_ladok --> |scim|EduID;
```     