import sys
import requests
from datetime import timedelta, datetime

sys.path.append('../..')
from tests.config import get_url_start
from tests.utils import get_timestamp

url = get_url_start() + "/api/v1/game"

json_data = {
    "ext_system_id": "custom-ext-system-id",
    "name": "new game",
    "answer_type": 2,
    "start_date":   str(get_timestamp(timedelta(days=1))),
    "end_date":     str(get_timestamp(timedelta(days=3))),
    "question": "Choose answer",
    "options": "yep, nope"
}

json_data2 = {
    "ext_system_id": "custom-ext-system-id",
    "name": "new game",
    "answer_type": 2,
    "start_date":   str(get_timestamp(timedelta(days=1))),
    "end_date":     str(get_timestamp(timedelta(days=2))),
    "question": "Choose answer",
    "options": "yep, nope"
}

# TODO: past | future | far future
# TODO: options (required with some answer type)


def get_far_future_timestamps():
    year = datetime.today().year
    start = datetime(year=year + 11, month=1, day=1).timestamp()
    end = datetime(year=year + 11, month=1, day=3).timestamp()
    return round(start), round(end)


def create_game():
    resp = requests.post(url=url, json=json_data)
    print("text", resp.text)
    assert resp.status_code == 200


def far_future_create_game():
    start, end = get_far_future_timestamps()
    json_data["start_date"] = str(start)
    json_data["end_date"] = str(end)
    resp = requests.post(url=url, json=json_data)
    print("text", resp.text)


if __name__ == '__main__':
    create_game()
    far_future_create_game()
