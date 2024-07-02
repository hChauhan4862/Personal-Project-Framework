from http import client
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import String,Text, Unicode, UnicodeText, Integer, SmallInteger, BigInteger, Numeric, Float, DateTime, Date, Time, LargeBinary, Enum, Boolean, JSON, ARRAY, TIMESTAMP
from sqlalchemy import Column
import datetime
from config import getenv

cache_DATABASE_URL = "sqlite:///./localdb.db"
if getenv('DETA_RUNTIME'):
   cache_DATABASE_URL = "sqlite:////tmp/localdb.db"

cache_Base = declarative_base()

class cache_login_data(cache_Base):
   __tablename__ = 'loginTokens'
   key = Column(String(40) ,primary_key = True, unique = True, nullable = False, index = True)
   clientId = Column(Integer, nullable = False)
   uid = Column(Integer, nullable = False)
   allowedRoutedArray = Column(JSON, nullable = False, default = [])
   allowedRouteCacheTime = Column(DateTime)
   login_at = Column(DateTime, default = datetime.datetime.utcnow, nullable = False)
   lastActive_at = Column(DateTime, default = datetime.datetime.utcnow, nullable = False)

class cache_clientTables_Times(cache_Base):
   __tablename__ = 'cacheTraceTable'
   key = Column(String(40) ,primary_key = True, unique = True, nullable = False, index = True)
   time = Column(DateTime, default = datetime.datetime.utcnow, nullable = False)