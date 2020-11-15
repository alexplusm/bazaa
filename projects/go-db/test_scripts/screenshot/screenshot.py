import requests

host_and_port = "http://localhost:1234/"

game_id = "my-game_id"
screenshot_id = "SCR_123"

url = host_and_port + "api/v1/game/" + game_id + "/screenshot/" + screenshot_id + "/answer"

jsonData = {
    'extSystemId': 'ext-ID-228',
    'userId': 'user-my-id-888',
    'answer': 'my-answer'
}

response = requests.post(url, json=jsonData)

print(response.text)
