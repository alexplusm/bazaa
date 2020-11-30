# TODO: setup venv
# TODO: modules/packages

from datetime import timedelta

import sys
sys.path.append('..')

from tests.ext_system.ext_system import create_ext_system
from tests.game.game import (
    create_game, update_game, prepare_game, get_games
)
from tests.screenshot.high_load import run_high_load
from tests.utils import get_timestamp
from tests.config import get_url_start
from tests.statistics.statistics import (
    get_statistics_user
)

# prepare game

ext_system = {
    # "extSystemId": "ext-id-5",
    "description": "some description",
    "postResultsUrl": "https://abc/lol.php"
}

game = {
    "extSystemId": None,
    "name": "new game",
    "answerType": 2,
    "startDate":   str(get_timestamp(timedelta(days=1))),
    "endDate":     str(get_timestamp(timedelta(days=2))),
    "question": "Choose answer",
    "options": "yep, nope"
}


def complete_test():
    print("Config: ", get_url_start())
    ext_system_id = create_ext_system(ext_system)
    if ext_system_id is None:
        ext_system_id = ext_system["extSystemId"]

    game["extSystemId"] = ext_system_id
    game_id = create_game(game)
    if game_id is None:
        print("error while game creation")
    update_game(game_id)
    prepare_game(game_id)

    run_high_load(ext_system_id, game_id)


def test_statistics():
    user_id = "user-id-1"
    game_id = "5c7713c7-3960-4c0d-ae7f-c27417ed234d"
    ext_system_id = "e23d6e88-7a4d-4945-aeb0-eaed9c6761eb"
    params = {
        "extSystemId": ext_system_id,
        "gameIds": game_id + "," + game_id + "asd",
        "totalOnly": "TrUe",
        "from": str(get_timestamp(timedelta(days=-10))),
        "to": str(get_timestamp(timedelta(days=1))),
    }
    res = get_statistics_user(user_id, params)
    print("result: ", res)


def main():
    test_statistics()

    # complete_test()

    # ext_system_id = "ext-id-3"
    # get_games(ext_system_id)


if __name__ == "__main__":
    main()