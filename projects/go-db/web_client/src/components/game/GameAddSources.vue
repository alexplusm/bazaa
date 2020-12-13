<template>
	<section>

		<h3>Добавить источник ккк</h3>

		<v-select
			v-model="selectedSourceType"
			:items="sourceTypes"
			label="Source type"
			dense
		></v-select>

		<div v-if="selectedSourceType === sourceTypes[0].value">
			<GameUploadFile />
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
import GameUploadFile from "./GameUploadFile"

export default {
	name: "GameAddSources",
	components: {
		GameUploadFile
	},
	data() {
		return {
			selectedSourceType: 0,
			// TODO: sync with sourceTypes
			sourceTypes: [
				{ value: 0, text: 'Архив' },
				{ value: 1, text: 'Расписание' },
				{ value: 2, text: 'Результат другой игры' },
			],
		};
	},
}
</script>

<style scoped>
</style>