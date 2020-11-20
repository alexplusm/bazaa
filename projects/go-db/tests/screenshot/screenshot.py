import requests


# --- async requests
# https://stackoverflow.com/questions/2632520/what-is-the-fastest-way-to-send-100-000-http-requests-in-python

# TODO: port from .env
host_and_port = "http://localhost:1234/"

game_id = "baabf15b-3a05-4592-9935-101637c12d67"
screenshot_id = "2111a89a-4527-494b-817f-1fc32836b9d1"
ext_system_id = 'custom-ext_system-id'

request_count = 10


def get_answer():
    url = host_and_port + "api/v1/game/"+ game_id \
          + "/screenshot?extSystemId=" + ext_system_id \
          + "&userId="

    for i in range(request_count):
        user_id = 'user-my-id-' + str(i)

        response = requests.get(url + user_id)
        print("Get Answer: " + str(i) + " | ", response.text, url + user_id)


def set_answer():
    url = host_and_port + "api/v1/game/" + game_id + "/screenshot/" + screenshot_id + "/answer"

    for i in range(request_count):
        json_data = {
            'extSystemId': ext_system_id,
            'userId': 'user-my-id-' + str(i),
            'answer': 'my-answer' + str(i)
        }

        response = requests.post(url, json=json_data)
        print("Set Answer: " + str(i) + " | ", response.text)


# get_answer()
set_answer()
