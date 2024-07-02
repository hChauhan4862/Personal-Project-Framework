import enum
from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship
from sqlalchemy import String,Text, Unicode, UnicodeText, Integer, SmallInteger, BigInteger, Numeric, Float, DateTime, Date, Time
from sqlalchemy import LargeBinary, Enum, Boolean, JSON, ARRAY, TIMESTAMP, ForeignKey
from sqlalchemy import Column
import datetime

DATABASE_URL = "mysql+pymysql://aJznj4oBB9:Zk0gQAoT34@remotemysql.com/aJznj4oBB9"
Base = declarative_base()


# Main Data Tables
class db_clientsInfo(Base):
   __tablename__ = 'clients_info'
   
   id         = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   short_code = Column(String(16) , unique = True, nullable = False, index = True)
   name       = Column(String(50), unique = True, nullable = False)
   address    = Column(Text)
   is_active  = Column(Boolean, default = True)
   is_del     = Column(Boolean, default = False)
   created_at = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, default = datetime.datetime.utcnow)

class db_menuMaster(Base):
   __tablename__ = 'menumaster'
   id         = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   title      = Column(String(30), nullable = False)
   menuType   = Column(String(10), nullable = False)
   is_active  = Column(Boolean, default = True)
   is_del     = Column(Boolean, default = False)
   created_at = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, default = datetime.datetime.utcnow)


# Client Specific Data Tables
class db_MenuClient(Base):
   __tablename__ = 'menuclient'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   MenuMasterId  = Column(Integer, nullable = False)
   clientId      = Column(Integer, nullable = False)
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)

class db_permissionGroups(Base):
   __tablename__ = 'permissiongroups'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   title         = Column(String(30), nullable = False)
   clientId      = Column(Integer, nullable = False)
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)

class db_menuPermission(Base):
   __tablename__ = 'menupermission'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   MenuClientId  = Column(Integer, nullable = False)
   permissionId  = Column(Integer, nullable = False)
   clientId      = Column(Integer, nullable = False)
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)

class db_UserLogin(Base):
   __tablename__ = 'user_logins'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   username      = Column(String(30), nullable = False)
   password      = Column(String(30), nullable = False)
   clientId      = Column(Integer, nullable = False)
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)