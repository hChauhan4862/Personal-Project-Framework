from src.models._base import *

class UserModel(BaseModel):
    user_id: Optional[int] = None
    first_name: str
    last_name: str
    email: str
    father_name: Optional[str] = Field(
        None, title="The father name of the User", max_length=300
    )
    age: float = Field(..., gt=0,
                       description="The age must be greater than zero")

class userResponse(hcSuccessModal):
    class dataClass(BaseModel):
        msg: int = Body( description="Always False")
    data: List[dataClass]
    pass