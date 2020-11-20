import requests
import json

import sys
sys.path.append('../..')
from tests.config import get_url_start

url = get_url_start() + "/api/v1/ext_system"

data_without_id = {
    "description": "Активный гражданин",
    "postResultsUrl": "https://abc/lol.php"
}

data = {
    "extSystemId": "ext-id-1",
    "description": "some description",
    "postResultsUrl": "https://abc/lol.php"
}


def make_req(json_data):
    resp = requests.post(url=url, json=json_data)
    print(url, " | ", end="")
    resp_body = json.loads(resp.text)
    print(resp_body)
    return resp_body


def create_ext_system(ext_system):
    resp_body = make_req(ext_system)
    if resp_body["success"] and resp_body["data"] is not None:
        return resp_body["data"]["extSystemId"]
    # return error mgs else None
    return None


if __name__ == '__main__':
    create_ext_system(data)
