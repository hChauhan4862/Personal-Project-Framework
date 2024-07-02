from db.config import *


class tbl_clients_info(Base):
   __tablename__ = 'clients_info'
   
   id = Column(Integer, primary_key = True, autoincrement = True, nullable = False )
   short_code = Column(String(16) , unique = True, nullable = False, index = True)
   name = Column(String(50), unique = True)
   address = Column(Text)
   is_active = Column(SmallInteger)
   is_del = Column(SmallInteger)
   created_at = Column(DateTime, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, default = datetime.datetime.utcnow)