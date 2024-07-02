from db_local.config import *
from sqlalchemy.orm import sessionmaker, Session
from fastapi import Depends



from db_local.tbl_clients_info import *



localDb_engine = create_engine(DATABASE_URL_LOCAL, echo=True)

Base_local.metadata.create_all(localDb_engine)


localDb_SessionLocal  = sessionmaker(autocommit=False, autoflush=False,bind = localDb_engine)

def get_db_local():
    db_local = localDb_SessionLocal()
    try:
        yield db_local
    finally:
        db_local.close()
