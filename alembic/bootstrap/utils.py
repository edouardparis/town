import os
import warnings

from alembic import op
from sqlalchemy import MetaData, exc as sa_exc


migration_path = os.path.join(os.path.dirname(os.path.realpath(__file__)), os.pardir)


def get_sql(directory, name):
    with open(os.path.join(migration_path, directory, 'sql', '{}.sql'.format(name)), 'r') as f:
        return f.read()


def get_primary_key_sequences(table_obj):
    for column_name, column_obj in table_obj.primary_key.columns.items():
        try:
            server_default = column_obj.server_default.arg.text
        except AttributeError:
            continue
        else:
            # server_default = "nextval('auth_user_id_seq'::regclass)"
            if server_default and server_default.startswith("nextval('"):
                yield column_name, server_default.split("'")[1]


def get_metadata_table(table_name, bind):
    metadata = MetaData()

    with warnings.catch_warnings():
        warnings.simplefilter("ignore", category=sa_exc.SAWarning)
        metadata.reflect(bind, only=(table_name,))

    return metadata.tables[table_name]


def insert_data(table_name, data, bind=None):
    bind = bind or op.get_bind()

    table_obj = get_metadata_table(table_name, bind)

    if data and not isinstance(data, list):
        data = [data]

    op.bulk_insert(table_obj, data)

    # Update primary key sequence so it doesn't crash on next insert
    for column_name, sequence_name in get_primary_key_sequences(table_obj):
        sql_stmt = "SELECT setval('{}', (SELECT max({}) from {}), true);".format(sequence_name, column_name, table_name)
        op.execute(sql_stmt)
