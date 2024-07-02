import databases
import sqlalchemy

DATABASE_URL = "mysql+pymysql://aJznj4oBB9:Zk0gQAoT34@remotemysql.com/aJznj4oBB9"

database = databases.Database(DATABASE_URL)
metadata = sqlalchemy.MetaData()

notes = sqlalchemy.Table(
    "notes",
    metadata,
    sqlalchemy.Column("id", sqlalchemy.Integer, primary_key=True),
    sqlalchemy.Column("text", sqlalchemy.String(30)),
    sqlalchemy.Column("completed", sqlalchemy.Boolean),
)

engine = sqlalchemy.create_engine(DATABASE_URL)
metadata.create_all(engine)
