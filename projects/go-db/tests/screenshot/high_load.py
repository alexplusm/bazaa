import asyncio
from random import random
import aiohttp
from aiohttp import ClientSession, ClientConnectorError
import json
import time

import sys
sys.path.append('../..')
from tests.config import get_url_start

# TODO: from config
# host_and_port = "http://localhost:8080/"
host_and_port = get_url_start() + '/'

# game_id = "52d9cfed-1d03-46f5-9921-fca5cb9c116e"
# ext_system_id = 'custom-ext_system-id'


def rand_answer():
    r = random()
    return '0' if r > 0.2 else '1'


def get_users():
    count = 100000
    if len(sys.argv) > 1:
        count = int(sys.argv[1])

    users = []
    for i in range(1, count + 1):
        users.append('i-user-' + str(i))
    return users


def get_query_params(user_id: str, ext_system_id) -> str:
    return '?extSystemId=' + ext_system_id + '&userId=' + user_id


async def set_answer_to_screenshot(ext_system_id, game_id, user_id: str, screenshot_id: str, session):
    general_url = host_and_port + 'api/v1/game/' + game_id + '/screenshot'
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


async def get_screenshot(ext_system_id, game_id, user_id: str, session: ClientSession):
    general_url = host_and_port + 'api/v1/game/' + game_id + '/screenshot'
    url = general_url + get_query_params(user_id, ext_system_id)
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
    if json_resp is not None and json_resp['success'] and json_resp['data'] is not None:
        print(method, ':', user_id, ' | ', json_resp['data'])
    else:
        print(method, ':', user_id, ' | ', json_resp)
    print('---')


async def user_case(ext_system_id, game_id, user_id, session: ClientSession):
    resp = await get_screenshot(ext_system_id, game_id, user_id, session)
    if resp is None or not resp['success']:
        return 0

    # resp['error']['message'] = "game is finished"

    screenshot_id = resp['data']['screenshot_id']

    # INFO: timeout
    # await asyncio.sleep(1 + random() * 2)
    await set_answer_to_screenshot(ext_system_id, game_id, user_id, screenshot_id, session)
    return 1


async def main(ext_system_id, game_id):
    users = get_users()

    t0 = time.time()

    async with ClientSession() as session:
        tasks = []
        list_list_users = []
        inner_users = []

        if len(users) < 5:
            list_list_users.append(users)

        for user in users:
            inner_users.append(user)
            if len(inner_users) == 5:
                list_list_users.append(inner_users)
                inner_users = []

        for list_users in list_list_users:
            for user in list_users:
                tasks.append(
                    user_case(ext_system_id, game_id, user, session)
                )

            results = await asyncio.gather(*tasks)
            tasks = []
            if 0 in results:
                break

        t1 = time.time()
        total = t1-t0
        print("TIME: ", total)


def run_high_load(ext_system_id, game_id):
    asyncio.run(main(ext_system_id, game_id))


if __name__ == "__main__":
    # run_high_load()
    pass
