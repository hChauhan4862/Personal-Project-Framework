SECRET_KEY = "40cc439e264811edeafa09bdca7d4d99422f9e57264811ed3bd5849b35b54ed1"
ALGORITHM = "HS256"
ACCESS_TOKEN_EXPIRE_MINUTES = 24 * 60

DEV_MODE = True

scopes_config = [
    {"id": 1, "name": "ROUTES::WRITE", "desc": "Create or Update Routes"},
    {"id": 2, "name": "ROUTES::READ", "desc": "Read Routes"},
    {"id": 3, "name": "ROUTES::DELETE", "desc": "Delete Routes"},
    {"id": 4, "name": "TEACHER::WRITE", "desc": "Create or Update a Teacher Record"},
]

STUDENT_REGEX = "^2[0-9]{6}(([-][A-Z])|([A-Z]))?$"

scope_ids = tuple(set([e["id"] for e in scopes_config]))
scope_names = tuple(set([e["name"] for e in scopes_config]))
assert len(scope_ids) == len(scope_names) 
def scopeId(scope: str):
    if scope not in scope_names:
        assert False, "Scope not found"
    return scope_ids[scope_names.index(scope)]