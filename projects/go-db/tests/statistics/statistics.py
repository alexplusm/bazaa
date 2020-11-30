import requests
import json

import sys
sys.path.append('../..')
from tests.config import get_url_start


def get_statistics_user(user_id, params):
    url = get_url_start() + "/api/v1/statistics/user/" + user_id
    resp = requests.get(url, params=params)
    resp_body = json.loads(resp.text)
    return resp_body
