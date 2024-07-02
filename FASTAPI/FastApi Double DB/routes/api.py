from fastapi import APIRouter, status
from src.endpoints import admin,userSelf
from hcResponseBase import hcCustomError



router = APIRouter(prefix="/api/v1", responses={
            status.HTTP_200_OK : {},
            status.HTTP_400_BAD_REQUEST: {"model": hcCustomError},
            status.HTTP_401_UNAUTHORIZED: {"description": "Unauthorized Access | Token expired or Not valid"},
            status.HTTP_403_FORBIDDEN: {"description": "Forbidden Access | Scope Not Found"},
            status.HTTP_422_UNPROCESSABLE_ENTITY: {},
            status.HTTP_429_TOO_MANY_REQUESTS: {},
        }
    )

router.include_router(admin.router)
router.include_router(userSelf.router)
