from src.endpoints._base import *
from src.models.student import *

#APIRouter creates path operations for item module
router = APIRouter(
    prefix="/student",
    tags=["Student"]
)

def ed_api(db, studentId):
    try:
        import requests, json
        header = {
            "Authorization": "Bearer KVRxkxcShKvjNvzNwJ3gJM3nBzR2O7mszd1e_QnSjl9h_R8jnGrYnlgEU8yVCc_Bu6UGrbJaZLbYiZYCLJ4xkLvupdUV7-pZzwfof-JdF7rUtEB4TjgpoZ9TSRlIsxHMZ4fr-4H-JIILkJxgmD7Q-iwL6WFIgTnYSyBaP4__348tEc0yzyUM1UmjaaP_cJy6Z5oYUXXqdoXC0PObPbqxVVES-VuIsj23Iz-VhH_QON1zIe-Lp7nJFGYUOzIeD4cgsg1jGf9j98B9iLcp1NPScELSGjt7gwlkrMMVmOOeVsvWiYllQFS1yIhhZIu97Eo--MjNpVsJX39q-iYyw-McPQ",
            "X-ContextId": "194",
            # "X-RX": "3",
            # "X-UserId": "218140",
            # "X-WB": "1"
        }
        ED = requests.get("https://beta.edumarshal.com/api/SearchStudentByEmployee/GetStudentByEmployee?admissionNo="+studentId, headers=header).json()
        if len(ED) == 0 or ED[0]["admissionNumber"].strip() != studentId.strip():
            return False
        ED = ED[0]
        userDetails = json.loads(ED["userDetails"])
        print(userDetails)
        db.add(StudentEdRecord(
            userId = ED["userId"],
            loginName = ED["loginName"],
            firstName = ED["firstName"],
            middleName = ED["middleName"],
            lastName = ED["lastName"],
            courseId = ED["courseId"],
            batchId = ED["batchId"],
            admissionNumber = ED["admissionNumber"],
            rollNumber = ED["rollNumber"],
            email = ED["email"],
            userDetails = userDetails,
            profilePictureId = ED["profilePictureId"],
            userStatus = ED["userStatus"],
            aadhaarNumber = ED["aadhaarNumber"],
            smsMobileNumber = ED["smsMobileNumber"],
            parentMobileNumber = ED["parentMobileNumber"],
            fatherName = ED["fatherName"],
            motherName = ED["motherName"],
            address = ED["address"],
            gender = ED["gender"],
            isDeleted = ED["isDeleted"],

            selectedBatch = userDetails["selectedBatch"] if "selectedBatch" in userDetails else None,
            selectedCourse = userDetails["selectedCourse"] if "selectedCourse" in userDetails else None,
            selectedBranch = userDetails["selectedBranch"] if "selectedBranch" in userDetails else None,
            selectedSemester = userDetails["selectedSemester"] if "selectedSemester" in userDetails else None,
            
        ))
        db.commit()
        return True
    except Exception as e:
        # print("HC_ERROR",e)
        raise hcCustomException(status_code=404, detail="Please update your details in Edumarshal")


@router.get("/fetchEDdata/{studentId}", responses={200: {"model": hcSuccessModal}})
async def read_root(studentId: str = Path(regex = STUDENT_REGEX), db: Session = Depends(get_db) ):
    studentId = studentId.upper()

    RECORD = db.query(StudentEdRecord).filter(StudentEdRecord.admissionNumber == studentId).with_entities(
        StudentEdRecord.admissionNumber,
        (StudentEdRecord.firstName + " " + StudentEdRecord.middleName + " " + StudentEdRecord.lastName).label("name"),
        StudentEdRecord.email,
        StudentEdRecord.smsMobileNumber,
        StudentEdRecord.selectedBranch, 
        StudentEdRecord.selectedBatch, 
        StudentEdRecord.selectedCourse,
        # StudentEdRecord.selectedSemester, 
        StudentEdRecord.isDeleted, 
        StudentEdRecord.userStatus
    ).first()
    if RECORD is None:
        if ed_api(db, studentId) == False:
            raise hcCustomException(status_code=404, detail="Student Not Found")
        RECORD = db.query(StudentEdRecord).filter(StudentEdRecord.admissionNumber == studentId).with_entities(
            StudentEdRecord.admissionNumber,
            (StudentEdRecord.firstName + " " + StudentEdRecord.middleName + " " + StudentEdRecord.lastName).label("name"),
            StudentEdRecord.email,
            StudentEdRecord.smsMobileNumber,
            StudentEdRecord.selectedBranch, 
            StudentEdRecord.selectedBatch, 
            StudentEdRecord.selectedCourse,
            # StudentEdRecord.selectedSemester, 
            StudentEdRecord.isDeleted, 
            StudentEdRecord.userStatus
        ).first()
        
    if RECORD:
        if RECORD.isDeleted or not(RECORD.userStatus):
            raise hcCustomException(403, detail="Student is Blocked or Deleted")
        return hcRes(data=RECORD)

    
    raise hcCustomException(status_code=404, detail="Something Went Wrong")


@router.post("/signup", responses={200: {"model": hcSuccessModal}})
async def read_root(form : SignUpForm, db: Session = Depends(get_db) ):
    if db.query(authInfo).filter(authInfo.username == form.studentId).first() is not None:
        raise hcCustomException(status_code=422, detail="Student already exists")
    
    ED_RECORD = db.query(StudentEdRecord).filter(StudentEdRecord.admissionNumber == form.studentId).with_entities(
        StudentEdRecord.admissionNumber,
        (StudentEdRecord.firstName + " " + StudentEdRecord.middleName + " " + StudentEdRecord.lastName).label("name"),
        StudentEdRecord.email,
        StudentEdRecord.smsMobileNumber.label("mobile"),
        StudentEdRecord.address,
        StudentEdRecord.rollNumber,
        StudentEdRecord.aadhaarNumber,
        StudentEdRecord.isDeleted, 
        StudentEdRecord.userStatus
    ).first()

    if ED_RECORD == None or ED_RECORD.isDeleted or not(ED_RECORD.userStatus) or not(ED_RECORD.aadhaarNumber == None or ED_RECORD.aadhaarNumber == form.aadhaarNumber) or not(ED_RECORD.rollNumber == None or ED_RECORD.rollNumber == form.rollNo):
        raise hcCustomException(status_code=422, detail="Student Data Mismatched or Fetch Again...")
    
    U = authInfo(username=form.studentId, password=get_password_hash(form.password), userRole="STUDENT", userStatus = True)
    db.add(U)
    db.flush()
    db.refresh(U)
    db.add(userInfo(uid = U.uid, name=ED_RECORD.name, email=ED_RECORD.email, phone=ED_RECORD.mobile, address=ED_RECORD.address))
    db.commit()
    
    return hcSuccessModal(data={"msg": "Success"})

