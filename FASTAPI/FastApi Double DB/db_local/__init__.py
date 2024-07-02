from distutils.log import debug
from db_local.config import *
from sqlalchemy.orm import sessionmaker, Session
from fastapi import Depends



cache_engine = create_engine(cache_DATABASE_URL, echo=True, connect_args={"check_same_thread": False})

cache_Base.metadata.create_all(cache_engine)


cache_SessionLocal  = sessionmaker(autocommit=False, autoflush=False,bind = cache_engine)

def cache_get_db():
    cache_db = cache_SessionLocal()
    try:
        yield cache_db
    finally:
        cache_db.close()

def cacheTblTimes(clintId, tblName, time=-1, cache_db=Depends(cache_get_db)):
    data = cache_db.query(cache_clientTables_Times).filter(cache_clientTables_Times.key == str(clintId)+"_"+tblName).first()
    if data is None or time != -1:
        time = datetime.datetime.utcnow() if time == -1 else time
        tmp = cache_clientTables_Times(key = str(clintId)+"_"+tblName, time = time)
        cache_db.add(tmp)
        cache_db.commit()
        return time
    return data.time
    