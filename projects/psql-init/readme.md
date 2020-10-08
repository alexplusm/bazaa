### Source
https://www.postgresqltutorial.com/


Run this in resource directory

* unzip dvdrental.zip
* mkdir .tmp
* tar -xf dvdrental.tar -C .tmp/


### Install PostgreSQL on ubuntu 20.04
https://www.postgresqltutorial.com/install-postgresql-linux/


### Install PostgreSQL on macOS
https://www.postgresqltutorial.com/install-postgresql-macos/


### Access the postgreSQL
// switch over postgres account
* sudo -i -u postgres

// access the postgreSQL
* psql

// logout postgres account
* exit


### Load sample database (ubuntu 20.04)

* sudo -i -u postgres
* psql
* create database dvdrental;
* \q
// move to with dvdrental.tar
* pg_restore --dbname=dvdrental --verbose dvdrental.tar

database filled

check data

* psql
* \c dvdrental
* select count(*) from film; // 1000
