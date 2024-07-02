from src.endpoints._base import *
from src.models.userSelf import *

#APIRouter creates path operations for item module
router = APIRouter(
    prefix="/me",
    tags=["Authorisation"]
)

@router.get("/myPermissions", responses={200: {"model": permissionResponse}})
async def get_route_permission_list(request: Request, me: Me = Security(get_current_user, scopes=["ROUTES::WRITE"])):
    permissions = me.db.query(
        db_userPermissions,db_routesPermissions
    ).join(db_routesPermissions, db_routesPermissions.permissionId == db_userPermissions.permissionId).filter(db_userPermissions.userId == me.id and db_routesPermissions.is_active == True and db_routesPermissions.is_del == False and db_routesPermissions.clientId == me.clientId
    ).join(db_routesClient, db_routesClient.id == db_routesPermissions.RouteClientId).filter(db_routesClient.is_active == True and db_routesClient.is_del == False and db_routesClient.clientId == me.clientId
    ).join(db_routesMaster, db_routesMaster.id == db_routesClient.routesMasterId).filter(db_routesMaster.is_active == True and db_routesMaster.is_del == False and db_routesMaster.clientId == me.clientId
    ).with_entities(db_routesMaster.id, db_routesMaster.title, db_routesMaster.detail).all()

    permissions = {k["title"]:k["detail"] for k in permissions}
    return hcRes(data=permissions)