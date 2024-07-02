from sqlalchemy import create_engine 
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import relationship
from sqlalchemy import String,Text, Unicode, UnicodeText, Integer, SmallInteger, BigInteger, Numeric, Float, DateTime, Date, Time
from sqlalchemy import LargeBinary, Enum, Boolean, JSON, ARRAY, TIMESTAMP, ForeignKey
from sqlalchemy import Column
import platform
import datetime

DATABASE_URL = "sqlite:///./localdb.db"
if platform.system() != "Windows":
   DATABASE_URL = "sqlite:////tmp/localdb.db"

Base = declarative_base()

class authInfo(Base):
   __tablename__ = 'auth_info'
   uid = Column(Integer,primary_key = True, unique = True, nullable = False, index = True, autoincrement = True)
   username = Column(String(50), nullable = False, unique = True)
   password = Column(String(64), nullable = False,)
   token_secret = Column(String(40) , unique = True)
   userRole = Column(String(10), nullable = False)
   userStatus = Column(Boolean, nullable = False, default = True)
   login_at = Column(DateTime)
   lastActive_at = Column(DateTime)
   created_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)

   info = relationship("userInfo")

class userInfo(Base):
   __tablename__ = 'user_info'
   uid = Column(Integer,ForeignKey('auth_info.uid'), primary_key = True, unique = True, nullable = False, index = True, autoincrement = True)
   name = Column(String(50), nullable = False)
   email = Column(String(50), nullable = False, unique = True)
   phone = Column(String(10), nullable = False, unique = True)
   address = Column(String(100), nullable = False)
   created_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)

class userPermissions(Base):
   __tablename__ = 'user_permissions'
   id = Column(Integer,primary_key = True, unique = True, nullable = False, index = True, autoincrement = True)
   per_id = Column(Integer, nullable = False, index = True)
   userRole = Column(String(10), nullable = False)
   allowed = Column(Boolean, nullable = False, default = True)
   created_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)

class StudentEdRecord(Base):
   __tablename__ = 'student_ed_record'
   userId = Column(Integer, primary_key = True, unique = True, nullable = False, index = True)
   loginName = Column(String(50), nullable = False, unique = True)
   firstName = Column(String(50), nullable = False)
   middleName = Column(String(50))
   lastName = Column(String(50))
   courseId = Column(Integer, nullable = False)
   batchId = Column(Integer, nullable = False)
   admissionNumber = Column(String(50), nullable = False, unique = True)
   rollNumber = Column(String(50))
   email = Column(String(50), nullable = False, unique = True)
   userDetails = Column(JSON)
   # dob = Column(DateTime)
   profilePictureId = Column(Integer, nullable = False)
   userStatus = Column(Boolean, nullable = False)
   # inactiveDate = Column(DateTime)
   aadhaarNumber = Column(String(50))
   smsMobileNumber = Column(String(50))
   parentMobileNumber = Column(String(50))
   fatherName = Column(String(50), nullable = False)
   motherName = Column(String(50))
   address = Column(String(100))
   gender = Column(String(10), nullable = False)
   isDeleted = Column(Boolean, nullable = False, default = False)

   selectedBatch = Column(String(60))
   selectedCourse = Column(String(60))
   selectedBranch = Column(String(60))
   selectedSemester = Column(String(60))

   created_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)
   updated_at = Column(DateTime, nullable = False, default = datetime.datetime.utcnow)