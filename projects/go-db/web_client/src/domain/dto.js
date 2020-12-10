import {answerTypesMap} from "./consts";

export function createGameToDTO(game) {
    const result = {};

    result.name = game.name;
    result.question = game.question;
    result.extSystemId = game.extSystem.extSystemId;
    result.answerType = game.answerType;

    const startDate = Number(new Date(game.startDate)) / 1000;
    const endDate = Number(new Date(game.endDate)) / 1000;

    result.startDate = startDate.toString()
    result.endDate = endDate.toString()

    if (game.answerType === answerTypesMap.categoryType.value) {
        result.options = game.options;
    }

    return result;
}
