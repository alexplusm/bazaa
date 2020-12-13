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
			<GameUploadArchive />
		</div>

		<div v-if="selectedSourceType === sourceTypes[1].value">
			Расписание
		</div>

		<div v-if="selectedSourceType === sourceTypes[2].value">
			Другая игра
		</div>
	</section>
</template>

<script>
import GameUploadArchive from "./GameUploadArchive"

export default {
	name: "GameAddSources",
	components: {GameUploadArchive},
	data: () => ({
		selectedSourceType: 0,  // TODO: default value
		// TODO: sync with sourceTypes || TODO domain consts
		sourceTypes: [
			{ value: 0, text: 'Архив' },
			{ value: 1, text: 'Расписание' },
			{ value: 2, text: 'Результат другой игры' },
		],
	}),
}
</script>
