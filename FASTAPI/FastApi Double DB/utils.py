from fastapi import APIRouter, Form, Security
from datetime import datetime, timedelta
from typing import Union

from fastapi import Depends, status
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm , SecurityScopes
from jose import JWTError, jwt
from passlib.context import CryptContext
from pydantic import BaseModel
from hcResponseBase import MeObj, hcCustomException, Me
import uuid

from db import *
from db_local import *
from config import *

# openssl rand -hex 32

pwd_context = CryptContext(schemes=["bcrypt"], deprecated="auto")
oauth2_scheme = OAuth2PasswordBearer(tokenUrl="api/auth/token_issue")

class TokenData(BaseModel):
    loginkey: Union[str, None] = None

class TokenResponse(BaseModel):
    access_token: str
    token_type: str
    expires_in: int

def random():
    return uuid.uuid1().hex[0:16]+uuid.uuid4().hex[0:16]

def verify_password(plain_password, hashed_password):
    return pwd_context.verify(plain_password, hashed_password)

def get_password_hash(password):
    return pwd_context.hash(password)


def create_access_token(data: dict, expires_delta: Union[timedelta, None] = None):
    to_encode = data.copy()
    if expires_delta:
        expire = datetime.datetime.utcnow() + expires_delta
    else:
        expire = datetime.datetime.utcnow() + timedelta(minutes=15)
    to_encode.update({"exp": expire})
    encoded_jwt = jwt.encode(to_encode, SECRET_KEY, algorithm=ALGORITHM)
    return encoded_jwt


async def get_current_user(security_scopes: SecurityScopes, token: str = Depends(oauth2_scheme), cache: Session = Depends(cache_get_db), db: Session = Depends(get_db)):
    credentials_exception = hcCustomException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        # detail="Could not validate credentials",
        headers={"WWW-Authenticate": "Bearer"},
    )
    cache.query(cache_login_data).filter(
            cache_login_data.lastActive_at < ( datetime.datetime.utcnow() - datetime.timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES) )
        ).delete()
    cache.commit()
    try:
        payload = jwt.decode(token, SECRET_KEY, algorithms=[ALGORITHM])
        loginkey: str = payload.get("sub")
        if loginkey is None:
            raise credentials_exception
        token_data = TokenData(loginkey=loginkey)
    except JWTError:
        raise credentials_exception
    user = cache.query(cache_login_data).filter_by(key=token_data.loginkey).first()
    if user is None:
        raise credentials_exception
    
    # Check user scope permissions
    print(str(cacheTblTimes(user.clientId, "PERMISSION_DATA",cache_db=cache))*10)
    if cacheTblTimes(user.clientId, "PERMISSION_DATA",cache_db=cache) > user.allowedRouteCacheTime:
        # Get user permissions
        print("Getting user permissions\n"*50)
        permissions = db.query(
                db_userPermissions,db_routesPermissions
            ).join(db_routesPermissions, db_routesPermissions.permissionId == db_userPermissions.permissionId).filter(db_userPermissions.userId == user.uid and db_routesPermissions.is_active == True and db_routesPermissions.is_del == False and db_routesPermissions.clientId == user.clientId
            ).join(db_routesClient, db_routesClient.id == db_routesPermissions.RouteClientId).filter(db_routesClient.is_active == True and db_routesClient.is_del == False and db_routesClient.clientId == user.clientId
            ).join(db_routesMaster, db_routesMaster.id == db_routesClient.routesMasterId).filter(db_routesMaster.is_active == True and db_routesMaster.is_del == False and db_routesMaster.clientId == user.clientId
            ).with_entities(db_routesMaster.id).all()
            
        permissions = [k["id"] for k in permissions]
        cache.merge(cache_login_data(
            key=loginkey,
            allowedRoutedArray=permissions,
            allowedRouteCacheTime = datetime.datetime.utcnow()
        ))
        cache.commit()
        cache.refresh(user)

    if user.uid != 0:
        for scope in security_scopes.scopes:
            if scopeId(scope) not in user.allowedRoutedArray: raise hcCustomException(status_code=403,headers={"WWW-Authenticate": "Bearer"})
    
    cache.query(cache_login_data).filter_by(key=token_data.loginkey).update(dict(lastActive_at=datetime.datetime.utcnow()))
    cache.commit()
    cache.refresh(user)
    u = db.query(db_UserLogin).filter_by(id=user.uid).with_entities(db_UserLogin.username, db_UserLogin.clientId, db_UserLogin.id).first()
    return MeObj(username=u.username, clientId=u.clientId, id=u.id, db=db, cache=cache)


router = APIRouter(
    prefix="/api/auth",
    tags=["Authorisation"],
    responses={
        status.HTTP_401_UNAUTHORIZED: {"description": "Incorrect username or password"},
        status.HTTP_422_UNPROCESSABLE_ENTITY: {},
        status.HTTP_429_TOO_MANY_REQUESTS: {},
    }
)


@router.post("/token_issue", responses={status.HTTP_200_OK: {"description": "Successful Response", "model": TokenResponse}})
async def token_generate(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(get_db), cache: Session = Depends(cache_get_db)):
    credentials_exception = hcCustomException(
        status_code=status.HTTP_401_UNAUTHORIZED,
        detail="Incorrect username or password",
        headers={"WWW-Authenticate": "Bearer"},
    )
    cache.query(cache_login_data).filter(
        cache_login_data.lastActive_at < ( datetime.datetime.utcnow() - datetime.timedelta(minutes=ACCESS_TOKEN_EXPIRE_MINUTES) )
    ).delete()
    cache.commit()

    user = db.query(db_UserLogin).filter(
                    db_UserLogin.username == form_data.username 
                    and db_UserLogin.is_active == 1
                    and db_UserLogin.is_del == 0
                    and db_UserLogin.disabled_by_admin == 0
                ).first()
    # raise credentials_exception
    if user is None:
        raise credentials_exception
    if not verify_password(form_data.password, user.password):
        raise credentials_exception
    access_token_expires = timedelta(minutes=365*24*60)  # 365 days
    key = str(user.id) + random()

    access_token = create_access_token(
        data={"sub": key}, expires_delta=access_token_expires
    )

    # permission = user.permissions
    # print(list(permission))

    
    # raise credentials_exception

    cache.add(cache_login_data(
        key=key,
        clientId=user.clientId,
        uid=user.id,
        allowedRouteCacheTime = datetime.datetime.utcfromtimestamp(0)
    ))
    cache.commit()
    return {"access_token": access_token, "token_type": "bearer", "expires_in": access_token_expires.total_seconds()}



