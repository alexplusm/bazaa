import requests
import json
from datetime import timedelta, datetime

import sys
sys.path.append('../..')
from tests.config import get_url_start
from tests.utils import get_timestamp

url = get_url_start() + "/api/v1/game"

json_data = {
    "ext_system_id": "custom-ext_system-id",
    "name": "new game",
    "answer_type": 2,
    "start_date":   str(get_timestamp(timedelta(days=1))),
    "end_date":     str(get_timestamp(timedelta(days=3))),
    "question": "Choose answer",
    "options": "yep, nope"
}

json_data2 = {
    "ext_system_id": "custom-ext_system-id",
    "name": "new game",
    "answer_type": 2,
    "start_date":   str(get_timestamp(timedelta(days=1))),
    "end_date":     str(get_timestamp(timedelta(days=2))),
    "question": "Choose answer",
    "options": "yep, nope"
}

# TODO: past | future | far future
# TODO: options (required with some answer type)


def make_req(json_data_g):
    resp = requests.post(url=url, json=json_data_g)
    print(url, " | ", end="")
    resp_body = json.loads(resp.text)
    print(resp_body)
    return resp_body


def get_far_future_timestamps():
    year = datetime.today().year
    start = datetime(year=year + 11, month=1, day=1).timestamp()
    end = datetime(year=year + 11, month=1, day=3).timestamp()
    return round(start), round(end)


def create_game(game):
    resp_body = make_req(game)
    if resp_body["success"] and resp_body["data"] is not None:
        return resp_body["data"]["gameId"]
    # return error mgs else None
    return None


def far_future_create_game():
    start, end = get_far_future_timestamps()
    json_data["start_date"] = str(start)
    json_data["end_date"] = str(end)
    resp = requests.post(url=url, json=json_data)
    print("text", resp.text)

# -----------
# /Users/a.mogilevskiy/work/5gen/clean-city/archives/low1.zip"
# /Users/a.mogilevskiy/work/5gen/clean-city/archives/low2.zip"
# /Users/a.mogilevskiy/work/5gen/clean-city/archives/medium.zip"


def update_game(game_id):
    archive = "/Users/a.mogilevskiy/work/5gen/clean-city/archives/low1.zip"
    files = {'archives': open(archive,'rb')}
    upload_url = url + "/" + game_id
    resp = requests.put(url=upload_url, files=files)

    print(upload_url, " | ", end="")
    resp_body = json.loads(resp.text)
    print(resp_body)
    return resp_body


def prepare_game(game_id):
    prepare_url = url + "/prepare"
    prep_game = {"gameId": game_id}
    resp = requests.post(url=prepare_url, json=prep_game)

    print(prepare_url, " | ", end="")
    resp_body = json.loads(resp.text)
    print(resp_body)


if __name__ == '__main__':
    far_future_create_game()
