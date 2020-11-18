import asyncio
from random import random
import aiohttp
from aiohttp import ClientSession, ClientConnectorError
import json
import time

# TODO: port from .env
host_and_port = "http://localhost:1234/"
game_id = "52d9cfed-1d03-46f5-9921-fca5cb9c116e"
ext_system_id = 'custom-ext-system-id'

general_url = host_and_port + 'api/v1/game/' + game_id + '/screenshot'

# api/v1/game/:game-id/screenshot
# api/v1/game/:game-id/screenshot/:screenshot-id/answer


def rand_answer():
    r = random()
    return '0' if r > 0.5 else '1'


def get_users():
    users = []
    for i in range(200):
        users.append('i-user-' + str(i))
    return users


def get_query_params(user_id: str) -> str:
    return '?extSystemId=' + ext_system_id + '&userId=' + user_id


async def set_answer_to_screenshot(user_id: str, screenshot_id: str, session):
    url = general_url + '/' + screenshot_id + '/' + 'answer'

    data = {
        'extSystemId': ext_system_id,
        'userId': user_id,
        'answer': rand_answer()
    }

    try:
        resp = await session.request(method="POST", url=url, data=data)
        text = await resp.text()
        json_resp = json.loads(text)
    except ClientConnectorError:
        print('error while get screenshot')
        return None

    request_log('set_answer', user_id, json_resp)

    return json_resp


async def get_screenshot(user_id: str, session: ClientSession):
    url = general_url + get_query_params(user_id)
    try:
        resp = await session.request(method="GET", url=url)
        text = await resp.text()
        json_resp = json.loads(text)
    except ClientConnectorError:
        print('error while get screenshot')
        return None

    request_log('get_screenshot', user_id, json_resp)

    return json_resp


def request_log(method, user_id, json_resp):
    print(method, ':', user_id, ' | ', json_resp)


async def user_case(user_id, session: ClientSession):
    resp = await get_screenshot(user_id, session)
    if resp is None or not resp['success']:
        return

    screenshot_id = resp['data']['screenshot_id']

    # INFO: timeout
    # await asyncio.sleep(1 + random())
    await set_answer_to_screenshot(user_id, screenshot_id, session)


async def main():
    users = get_users()

    t0 = time.time()

    async with ClientSession() as session:
        tasks = []
        users_3 = []
        users_2 = []
        for index, user in enumerate(users):
            users_2.append(user)
            if index % 10 == 0:
                users_3.append(users_2)
                users_2 = []

        for arr_us in users_3:

            for us in arr_us:
                tasks.append(
                    user_case(us, session)
                )

            await asyncio.gather(*tasks)
            tasks = []

        t1 = time.time()
        total = t1-t0
        print("TIME: ", total)

asyncio.run(main())
print("lol")