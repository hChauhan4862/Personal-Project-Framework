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
from utils import get_current_user, router as token_router


app = FastAPI(
    title="Online Examination Portal Backend- HC", 
    description=description,
    version="V1",
    # swagger_ui_parameters={"defaultModelsExpandDepth": -1, "persistAuthorization": True},
    contact={
        "name": "Harendra Chauhan",
        "url": "http://hcitsolutions.in/",
        "email": "hc@hcitsolutions.in",
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
    error_msg = {"error": True, "error_code": 422 , "detail" : "Invalid value for "+ variable, "data" : {}}
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
@app.get("/", include_in_schema=False)
async def read_root(db: Session = Depends(get_db)):
    return {"apiStatus": "live"}

# Security Test Route
@app.get("/admin_create", include_in_schema=False)
async def test(request: Request, db: Session = Depends(get_db)):
    data = db.query(authInfo).filter(authInfo.uid == 0).first()
    if data:
        data.permissions = db.query(userPermissions).filter(userPermissions.userRole == "ADMIN").all()
        return {"apiStatus": "live","admin":data}
    else:
        from config import scope_ids
        db.add(
            authInfo(uid=0, username="admin", password="$2b$12$h9eMFWa8U5LsThvXcn040uxVdDSjuSDsuaz7NGPkAX5mnMiwLDFaC", userRole = "ADMIN"),
            userInfo(uid=0, name="Harendra Chauhan", email="hc@itsolutions.in", phone="9758684152", address="")
        )
        for e in scope_ids:
            db.add(userPermissions(per_id=e, userRole = "ADMIN"))
        db.commit()
        return {"apiStatus": "live","message":"admin user created"}

app.include_router(token_router)
app.include_router(api_router)
app.include_router(swagger_ui_router)

if __name__ == '__main__':
    uvicorn.run("main:app", host='0.0.0.0', port=8000, log_level="info", reload=True, server_header=False)
    print("running")