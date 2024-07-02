from src.endpoints._base import *

#APIRouter creates path operations for item module
router = APIRouter(
    prefix="/admin",
    tags=["Admin"]
)

@router.get("/create_routes", responses={200: {"model": hcSuccessModal}})
# async def read_root(me = Security(get_current_user, scopes=["ROUTES::WRITE"])):
    # db = me.db
async def read_root(db = Depends(get_db)):
    for e in scopes_config:
        id = e["id"]
        title = e["name"]
        type = title.split("::")[1]
        desc = e["desc"]

        t = db_routesMaster(id=id,title=title, menuType=type, detail = desc, is_active = 1, is_del = 0)
        db.merge(t)
    
    db.commit()
    return hcSuccessModal(data={"msg": "Success"})
