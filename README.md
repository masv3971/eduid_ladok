# eduid_ladok

Integration between eduid and ladok

## Prerequisites

### build

* make
* docker
* docker-compose

### x509

1. Request from school with permission:
    * 11004:  kataloginformation.las
    * 21008:  kataloginformationbehorighet.behorigheter.allt.las
    * 21010:  kataloginformationbehorighet.anvandare.las
    * 51001:  studiedeltagande.las

## compile, run and stop

### docker-compose

* ```$ make container-build```
* ```$ make container-start```
* ```$ make container-stop```

### Local (Linux)

* ```$ make```
* ```$ ./start_deps.sh```
* ```$ ./bin/eduid_ladok```
