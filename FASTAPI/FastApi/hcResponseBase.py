from typing import Any, Callable, Optional, Sequence, Union
from fastapi import Body, params
from pydantic import BaseModel
from sqlalchemy.orm import Session

# class Token(BaseModel):
#     access_token: str
#     token_type: str

def Security(dependency: Optional[Callable[..., Any]] = None,*,scopes: Optional[Union[Sequence[int],Sequence[str]]] = None,use_cache: bool = True):
    if scopes and len(scopes) > 0 and type(scopes[0]) == int:
        for i in range(len(scopes)):
            scopes[i] = str(scopes[i])
    return params.Security(dependency=dependency, scopes=scopes, use_cache=use_cache)

class hcCustomError(BaseModel):
    error: bool = Body(default=True)
    error_code : int = Body(default=400)
    detail : str = Body(default="Error Description")
    data : dict = Body(example={})

class hcCustomException(Exception):
    def __init__(self, status_code: int = 400, detail: str = "", error_code: int = -1, headers: dict = None):
        self.detail = detail
        self.status_code = status_code
        self.error_code = error_code if error_code != -1 else status_code
        self.headers = headers
        self.data = {}

class hcSuccessModal(BaseModel):
    error: int = Body( example=False, default=False)
    error_code: int = Body( example=200, default=200)
    detail : str = Body(default="Success")
    # data: list = list[dataClass]

class Me(BaseModel):
    uid: int
    username: str
    role: str
    scopes: list
    db: Session
    
    class Config:
        arbitrary_types_allowed = True

class MeObj(object):
    def __init__(self, username, uid,role,scopes, db):
        self.username = username
        self.uid = uid
        self.role = role
        self.db = db
        self.scopes = scopes

class hcRes(object):
    def __init__(self, error: bool = False, error_code: int = 200, detail: str = "Success", data: dict = {}):
        self.error = error
        self.error_code = error_code
        self.detail = detail
        self.data = data

# class hcSuccessModal2(hcSuccessModal):
#     class dataClass(BaseModel):
#         msg: int = Body( description="Always False")
#     data: List[dataClass]

# class userResponse(hcSuccessModal):
#     class dataClass(BaseModel):
#         msg: str = Body( description="Always False")
#     data: List[dataClass]
#     pass



description = """
Welcome to the API Documentation

API helps you do awesome stuff. 🚀

Online Examination Portal is a project which will help teachers to automatic evaluate answers given by students and generate report with the help of intelligence.

### Project Developer Team:
* Harendra Chauhan
* YashKant Tyagi

### Technologies:
	Backend Rest API:	Python
	WebSocket:		    Python
	AI SYSTEM		    Python
	DATABASE		    MySQL (RDBMS)
	FRONTEND		    HTML, CSS (BootStrap), JAVASCRIPT (JQuery)

### AI SYSTEM:
	AI SYSTEM will help teacher in evaluation.
	AI Chat Bot for help desk
	Monitor students activity in exam windows e.g. detection of opening new windows, minimization of exam windows

### Security in Online Examination Portal:

	Use of middleware
	Disallow unauthorized api access
	Smart use of CORS Policies
	Filtering each input given by user for preventing XSS attack
	Allow only single login at same time
	Different Permissions for each roles
	There will be only single Admin who will manage all multiple teachers, students, helpdesk logins.

### Overview on Roles and Responsibilities:
	admin panel:
		Add Teacher
		Add Student
	teacher panel:
		Add / Remove/ Manage Exams
		Add / Remove/ Manage Question
		Assign Exam to class or Student
		Start Exam
		Check Result	
		Help AI to evaluate by adding Answers
	Help Desk Panel:
		Resolving Queries.
	student panel:
		Attempts Exams
		HelpDesk
        ChatBot
"""