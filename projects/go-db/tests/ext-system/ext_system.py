import sys
import requests

sys.path.append('../..')
from tests.config import get_url_start

url = get_url_start() + "/api/v1/ext-system"

data_without_id = {
    "description": "Активный гражданин",
    "postResultsUrl": "https://abc/lol.php"
}


def create_ext_system():
    data = {
        "extSystemId": "keeekak",
        "description": "some description",
        "postResultsUrl": "https://abc/lol.php"
    }
    resp = requests.post(url=url, json=data)
    print(resp.text)


if __name__ == '__main__':
    create_ext_system()
