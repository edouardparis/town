create-database:
	sudo psql -U postgres -c "DROP DATABASE IF EXISTS lizardstown"
	sudo psql -U postgres -c "CREATE DATABASE lizardstown WITH OWNER lizardstown"

bootstrap:
	alembic -c alembic/alembic.ini -x with-bootstrap upgrade head
