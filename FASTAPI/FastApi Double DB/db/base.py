from db.config import *
from sqlalchemy.orm import sessionmaker, Session
from fastapi import Depends



from db.tbl_clients_info import *



engine = create_engine(DATABASE_URL, echo=True)

engine.connect()


# Base.metadata.create_all(engine)


SessionLocal  = sessionmaker(autocommit=False, autoflush=False,bind = engine)

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
