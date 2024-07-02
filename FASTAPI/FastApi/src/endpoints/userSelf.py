from src.endpoints._base import *
from src.models.userSelf import *

#APIRouter creates path operations for item module
router = APIRouter(
    prefix="/me",
    tags=["Authorisation"]
)

@router.get("/myPermissions", responses={200: {"model": permissionResponse}})
async def get_route_permission_list(request: Request, me: Me = Security(get_current_user, scopes=["ROUTES::READ"])):
    permissions = {k["name"]:k["desc"] for k in scopes_config if k["id"] in me.scopes}
    return hcRes(data=permissions)