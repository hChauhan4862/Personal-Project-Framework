from fastapi import APIRouter, Request
from typing import List, Optional
from fastapi import Query

from utils import scopes_config, scope_ids, get_current_user, Security
from hcResponseBase import hcSuccessModal, hcCustomException, Me, hcRes
from db import *
from db_local import *