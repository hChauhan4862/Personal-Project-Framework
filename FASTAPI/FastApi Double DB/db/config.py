from sqlalchemy import create_engine 
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship
from sqlalchemy import String,Text, Unicode, UnicodeText, Integer, SmallInteger, BigInteger, Numeric, Float, DateTime, Date, Time
from sqlalchemy import LargeBinary, Enum, Boolean, JSON, ARRAY, TIMESTAMP, ForeignKey
from sqlalchemy import Column
import datetime
from config import getenv

DATABASE_URL = "mysql+pymysql://root:@localhost/storemanager"
if getenv('DETA_RUNTIME'):
   DATABASE_URL = "mysql+pymysql://aJznj4oBB9:Zk0gQAoT34@remotemysql.com/aJznj4oBB9"

DATABASE_URL = "mysql+pymysql://root:Hps202132@54.226.224.14/storemanager"
# DATABASE_URL = "mysql+pymysql://aJznj4oBB9:Zk0gQAoT34@remotemysql.com/aJznj4oBB9"
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

   # logins = relationship("db_UserLogin", back_populates = "clientInfo")
   # menus = relationship("db_routesClient", back_populates = "clientInfo")
   # permissions_master = relationship("db_permissionsMaster", back_populates = "clientInfo")
   # permissionMenus  = relationship("db_routesPermissions", back_populates = "clientInfo")


class db_routesMaster(Base):
   __tablename__ = 'routes_master'
   id         = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   title      = Column(String(30), nullable = False)
   menuType   = Column(String(10), nullable = False)
   detail   = Column(Text, nullable = False)
   is_active  = Column(Boolean, default = True)
   is_del     = Column(Boolean, default = False)
   created_at = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, default = datetime.datetime.utcnow)


# Client Specific Data Tables
class db_routesClient(Base):
   __tablename__ = 'routes_client'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   routesMasterId  = Column(Integer, ForeignKey('routes_master.id'))
   clientId      = Column(Integer, ForeignKey('clients_info.id'))
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)

   # clientInfo = relationship("db_clientsInfo", back_populates = "menus")
   # routes_permissionss = relationship("db_routesPermissions", back_populates = "clientMenus")

class db_permissionsMaster(Base):
   __tablename__ = 'permissions_master'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   title         = Column(String(30), nullable = False)
   clientId      = Column(Integer, ForeignKey('clients_info.id'))
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)

   # clientInfo = relationship("db_clientsInfo", back_populates = "permissions_master")
   # menus = relationship("db_routesPermissions", back_populates = "permissions")

class db_routesPermissions(Base):
   __tablename__ = 'routes_permissions'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   RouteClientId  = Column(Integer, ForeignKey('routes_client.id'), nullable = False)
   permissionId  = Column(Integer, ForeignKey('permissions_master.id'))
   clientId      = Column(Integer, ForeignKey('clients_info.id'))
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)

   # permissions = relationship("db_permissionsMaster", back_populates = "menus")
   # clientMenus = relationship("db_routesClient", back_populates = "routes_permissionss")
   # clientInfo = relationship("db_clientsInfo", back_populates = "permissionMenus")

class db_UserLogin(Base):
   __tablename__ = 'user_logins'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   username      = Column(String(30), nullable = False, unique = True, index = True)
   password      = Column(String(30), nullable = False)
   clientId      = Column(Integer, ForeignKey('clients_info.id'))
   disabled_by_admin     = Column(Boolean, default = False)
   is_active     = Column(Boolean, default = True)
   is_del        = Column(Boolean, default = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow)

   # clientInfo    = relationship('db_clientsInfo', backref='login')
   # permissions    = relationship("db_userPermissions")


class db_userPermissions(Base):
   __tablename__ = 'user_permissions'
   id            = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   userId        = Column(Integer, ForeignKey('user_logins.id'), nullable = False)
   permissionId  = Column(Integer, ForeignKey('permissions_master.id'), nullable = False)
   clientId      = Column(Integer, ForeignKey('clients_info.id'), nullable = False)
   is_active     = Column(Boolean, default = True, nullable = False)
   is_del        = Column(Boolean, default = False, nullable = False)
   created_at    = Column(DateTime, default = datetime.datetime.utcnow, nullable = False)
   updated_at    = Column(DateTime, default = datetime.datetime.utcnow, nullable = False)

