# TODO: setup venv
# TODO: modules/packages

from datetime import timedelta

import sys
sys.path.append('..')

from tests.ext_system.ext_system import create_ext_system
from tests.game.game import create_game, update_game, prepare_game
from tests.screenshot.high_load import run_high_load
from tests.utils import get_timestamp
from tests.config import get_url_start

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


def main():
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


if __name__ == "__main__":
    main()
