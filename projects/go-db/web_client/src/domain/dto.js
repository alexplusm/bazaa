import { answerTypesMap } from './consts';
import { dateToDto } from '../utils/date-utils';

export function createGameToDTO(game) {
	const result = {};

	result.name = game.name;
	result.question = game.question;
	result.extSystemId = game.extSystem.extSystemId;
	result.answerType = game.answerType;

	// TODO: test!
	const startDate = dateToDto(new Date(game.startDate));
	const endDate = dateToDto(new Date(game.endDate));

	result.startDate = startDate.toString();
	result.endDate = endDate.toString();

	if (game.answerType === answerTypesMap.categoryType.value) {
		result.options = game.options;
	}

	return result;
}
