import sys
import requests
from datetime import timedelta

sys.path.append('../..')
from tests.config import get_url_start
from tests.utils import get_timestamp

url = get_url_start() + "/api/v1/game"

json_data = {
    "ext_system_id": "custom-ext-system-id",
    "name": "new game",
    "answer_type": 2,
    "start_date":   str(get_timestamp(timedelta(days=2))),
    "end_date":     str(get_timestamp(timedelta(days=3))),
    "question": "Choose answer",
    "options": "yep, nope"
}

# TODO: past | future | far future
# TODO: options (required with some answer type)


def create_game():
    resp = requests.post(url=url, json=json_data)
    print("text", resp.text)
    assert resp.status_code == 200


if __name__ == '__main__':
    create_game()
