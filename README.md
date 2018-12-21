# Lizards Town

# install

# Install virtualenv

```
pip install virtualenv
virtualenv .env
source .env/bin/activate
```

`pip install alembic`

# Database installation

Database installation

To install postgresql on linux, you should use the following instructions

Linux

`sudo apt-get install libgeos-dev libxml2-dev libgdal-dev`

OSX

`brew install postgresql`

At this point your database server should be started.

## Create a new user

Linux

`sudo -u postgres psql`

OSX

`psql postgres`

BOTH

Then create the user

`CREATE USER lizardstown WITH PASSWORD 'lizardstown' CREATEDB;`

 ## Create a new database

Linux

`sudo -u postgres psql`

OSX

`psql postgres`

Then create the database

`CREATE DATABASE lizardstown WITH OWNER lizardstown;`

