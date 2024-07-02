from src.models._base import *

class permissionResponse(hcSuccessModal):
    data: dict = Body(default={
        "permission_name_1": "Description_1", 
        "permission_name_2": "Description_2", 
        "permission_name_3": "Description_3", 
        "...": "..."
    })