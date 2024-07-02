from db.config import Base, create_engine, DATABASE_URL
from db.config import db_clientsInfo, db_routesMaster, db_routesClient, db_permissionsMaster, db_routesPermissions, db_UserLogin, db_userPermissions
from sqlalchemy.orm import sessionmaker, Session
from fastapi import Depends

engine = create_engine(DATABASE_URL, echo=True)

# Base.metadata.create_all(engine)


SessionLocal  = sessionmaker(autocommit=False, autoflush=False,bind = engine)

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
