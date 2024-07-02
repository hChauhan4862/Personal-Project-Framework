import time
from traceback import print_tb
from fastapi import FastAPI, Request, status, Security
from fastapi.openapi.utils import get_openapi
from fastapi.middleware.cors import CORSMiddleware

from fastapi.exceptions import RequestValidationError, FastAPIError
from fastapi.responses import JSONResponse, HTMLResponse, FileResponse
import uvicorn
from sqlalchemy.exc import SQLAlchemyError


from routes.api import router as api_router
from swagger_ui import router as swagger_ui_router
from hcResponseBase import Me, hcCustomError, hcCustomException, description

from db import *
from db_local import *
from utils import get_current_user, router as token_router


app = FastAPI(
    title="Store Manager API - HC", 
    # description=description,
    description="",
    version="V1",
    # swagger_ui_parameters={"defaultModelsExpandDepth": -1, "persistAuthorization": True},
    terms_of_service="http://hcitsolutions.in/storemanager/terms-of-service.html",
    contact={
        "name": "Harendra Chauhan",
        "url": "http://hcitsolutions.in/",
        "email": "harendra_chauhan@yahoo.com",
    },
    license_info={
        "name": "Apache 2.0",
        "url": "https://www.apache.org/licenses/LICENSE-2.0.html",
    },
    responses={200:{}},
    # openapi_url=None,
    docs_url=None

)

@app.exception_handler(hcCustomException)
async def unicorn_exception_handler(request: Request, exc: hcCustomException):
    if exc.detail == "" and exc.status_code in [401,403,422,429]:
        return HTMLResponse(status_code=exc.status_code, content="")

    return JSONResponse(
        status_code=exc.status_code,
        content={"error": True, "error_code": exc.error_code, "detail": exc.detail, "data" : exc.data},
        headers=exc.headers
    )

@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request, exc):
    # print(exc.errors())
    try:
        variable = exc.errors()[0]["loc"][1]
    except Exception as e:
        variable = ""
    error_msg = {"error": True, "error_code": 422 , "detail" : "Validation Error in variable "+ variable, "data" : {}}
    return JSONResponse(error_msg, status_code=status.HTTP_400_BAD_REQUEST)

@app.exception_handler(FastAPIError)
def handle_fastapiError(request, exc):
    raise hcCustomException(status.HTTP_500_INTERNAL_SERVER_ERROR, detail="Internal Server Error - Error")

@app.exception_handler(SQLAlchemyError)
def handle_sqlalchemy(request: Request, error):
    print(str(error))
    raise hcCustomException(status.HTTP_500_INTERNAL_SERVER_ERROR, detail="Internal Server Error - Error"+str(error))

# Middle Ware for all Routes
@app.middleware("http")
async def add_process_time_header(request: Request, call_next):
    start_time = time.time()
    response = await call_next(request)
    process_time = time.time() - start_time
    response.headers["X-Process-Time"] = str(process_time)
    return response

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:8000","*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Base routes
@app.get('/favicon.ico', include_in_schema=False)
async def favicon():
    return FileResponse("favicon.ico")
@app.get("/")
async def read_root(db: Session = Depends(get_db)):
    return {"apiStatus": "live"}

# Security Test Route
@app.get("/test")
async def test(request: Request, me: Me = Security(get_current_user, scopes=[])):
    permissions = me.db.query(
        db_userPermissions,db_routesPermissions
    ).join(db_routesPermissions, db_routesPermissions.permissionId == db_userPermissions.permissionId).filter(db_userPermissions.userId == me.id and db_routesPermissions.is_active == True and db_routesPermissions.is_del == False and db_routesPermissions.clientId == me.clientId
    ).join(db_routesClient, db_routesClient.id == db_routesPermissions.RouteClientId).filter(db_routesClient.is_active == True and db_routesClient.is_del == False and db_routesClient.clientId == me.clientId
    ).join(db_routesMaster, db_routesMaster.id == db_routesClient.routesMasterId).filter(db_routesMaster.is_active == True and db_routesMaster.is_del == False and db_routesMaster.clientId == me.clientId
    ).with_entities(db_routesMaster.id).all()

    # permissions = me.db.query(db_userPermissions).join(
    #     db_routesPermissions, db_userPermissions.permissionId == db_routesPermissions.permissionId
    # ).join(
    #     db_routesClient, db_routesPermissions.RouteClientId == db_routesClient.id
    # ).join(
    #     db_routesMaster, db_routesClient.routesMasterId == db_routesMaster.id
    # ).filter(
    #     db_userPermissions, db_userPermissions.userId == me.id
    # )

    # permissions = me.db.query(db_routesMaster).join(
    #     db_routesMaster.is_active == True and db_routesMaster.is_del == False
    # ).join(
    #     db_routesClient, db_routesClient.routesMasterId == db_routesMaster.id
    # ).join(
    #     db_routesPermissions, db_routesPermissions.RouteClientId == db_routesClient.id
    # ).join(
    #     db_userPermissions, db_userPermissions.permissionId == db_routesPermissions.id
    # ).filter(
    #     db_userPermissions, db_userPermissions.userId == me.id
    # )

    # permissions = me.db.query(db_routesMaster).join(
    #     db_routesClient).join(db_routesPermissions).join(db_userPermissions).filter(
    #         db_routesMaster.is_active == True and db_routesMaster.is_del == False and
    #         db_routesClient.routesMasterId == db_routesMaster.id and db_routesClient.is_active == True and db_routesClient.is_del == False and db_routesClient.clientId == me.clientId and
    #         db_routesPermissions.RouteClientId == db_routesClient.id and db_routesPermissions.is_active == True and db_routesPermissions.is_del == False and db_routesPermissions.clientId == me.clientId and
    #         db_userPermissions.permissionId == db_routesPermissions.id and db_userPermissions.is_active == True and db_userPermissions.is_del == False and db_userPermissions.clientId == me.clientId and
    #         db_userPermissions.userId == me.id
    #         # me.id == db_userPermissions.userId and db_userPermissions.is_active == True and db_userPermissions.is_del == False and
    #         # db_routesPermissions.id == db_userPermissions.permissionId and db_routesPermissions.is_active == True and db_routesPermissions.is_del == False and
    #         # db_routesClient.id == db_routesPermissions.RouteClientId and db_routesClient.is_active == True and db_routesClient.is_del == False and
    #     )
    # permissions = str(permissions)

    permissions = [k["id"] for k in permissions]
    return {"permissions": permissions}
    
    user.permissions
    del user.password
    del user.db
    del user.cache
    return {"current_user": user}


app.include_router(token_router)
app.include_router(api_router)
app.include_router(swagger_ui_router)

if __name__ == '__main__':
    uvicorn.run("main:app", host='0.0.0.0', port=8000, log_level="info", reload=True, server_header=False)
    print("running")