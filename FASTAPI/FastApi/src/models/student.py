from src.models._base import *

class SignUpForm(BaseModel):
    studentId: str = Field(..., title="The username of the Student", max_length=30, regex=STUDENT_REGEX)
    password: str = Field(..., title="The password of the Student", max_length=50, regex="^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*(_|[^\w])).+$")
    rollNo: str = Field(..., title="The Roll Number of the Student", max_length=13, regex="^2[0-9]{12}$")
    aadhaarNumber: str = Field(..., title="The Aadhaar Number of the Student", max_length=12, regex="^[0-9]{12}$")
