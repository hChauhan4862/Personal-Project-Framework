from sqlalchemy import create_engine 
DATABASE_URL = "mysql+pymysql://root:Hps202132@54.226.224.14/storemanager"
DATABASE_URL = "mysql+pymysql://aJznj4oBB9:Zk0gQAoT34@remotemysql.com/aJznj4oBB9"
DATABASE_URL = "mysql+pymysql://root:@localhost/storemanager"
DATABASE_URL = "mysql+pymysql://rishi:@54.226.224.14/storemanager"

engine = create_engine(DATABASE_URL, echo=True)
try:
    engine.connect()
except Exception as e:
    print(e)


