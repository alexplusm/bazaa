<template>
	<section>
		<h3 class="mb-8">Добавить источник</h3>

		<v-select
			label="Источник"
			v-model="selectedSourceType"
			:items="sourceTypes"
			dense
		></v-select>

		<div v-if="selectedSourceType === sourceTypes[0].value">
			<GameAttachArchives />
		</div>

		<div v-if="selectedSourceType === sourceTypes[1].value">Расписание</div>

		<div v-if="selectedSourceType === sourceTypes[2].value">
			<GameAttachAnotherGameResults />
		</div>
	</section>
</template>

<script>
import { sourceTypesMap, sourceTypesArray } from '../../domain/consts';
import GameAttachArchives from './GameAttachArchives';
import GameAttachAnotherGameResults from './GameAttachAnotherGameResults';

export default {
	name: 'GameAddSources',
	components: { GameAttachArchives, GameAttachAnotherGameResults },
	data: () => ({
		selectedSourceType: sourceTypesMap.archiveSourceType.value,
		sourceTypes: sourceTypesArray,
	}),
};
</script>
