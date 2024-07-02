from src.endpoints._base import *
from src.models.admin import *

#APIRouter creates path operations for item module
router = APIRouter(
    prefix="/admin",
    tags=["Admin"]
)

@router.post("/add_teacher", responses={200: {"model": hcSuccessModal}})
async def read_root(form : TeacherAdd, me = Security(get_current_user, scopes=["TEACHER::WRITE"])):
    db = me.db
    if db.query(authInfo).filter(authInfo.username == form.username).first() is not None:
        raise hcCustomException(status_code=422, detail="Username already exists")
    U = authInfo(username=form.username, password=get_password_hash(form.password), userRole="TEACHER")
    db.add(U)
    db.flush()
    db.refresh(U)
    db.add(userInfo(uid = U.uid, name=form.name, email=form.email, phone=form.phone, address=form.address))
    db.commit()
    
    return hcSuccessModal(data={"msg": "Success"})
