<template>
	<section>
		<v-progress-circular v-if="loading" indeterminate color="primary">
		</v-progress-circular>

		<v-row v-else>
			<!-- TODO: GameDetailsInfo component -->
			<v-col cols="5">
				<div>Start date: {{ this.currentGame.startDate }}</div>
				<div>Finish date: {{ this.currentGame.finishDate }}</div>
				<div>
					Answer type: {{ this.currentGame.question.answerType }}
				</div>
				<div>Question: {{ this.currentGame.question.text }}</div>

				<ul>
					Options
					<li
						v-for="option in this.currentGame.question.options"
						:key="option.option"
					>
						text: <strong>{{ option.text }}</strong> (option:
						<strong>{{ option.option }}</strong
						>)
					</li>
				</ul>

				<ul>
					Sources
					<li
						v-for="source in this.currentGame.sources"
						:key="source.sourceId"
					>
						{{ source.sourceId }} | {{ source.type }}
					</li>
				</ul>
			</v-col>

			<v-col cols="4">
				<h3>Добавить источник</h3>

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
			</v-col>
		</v-row>
	</section>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import GameUploadFile from './GameUploadFile';

export default {
	name: 'GameDetails',
	components: { GameUploadFile },
	data() {
		return {
			loading: true,
			selectedSourceType: 0,
			// TODO: sync with sourceTypes
			sourceTypes: [
				{ value: 0, text: 'Archive' },
				{ value: 1, text: 'Schedule' },
				{ value: 2, text: 'Game results' },
			],
		};
	},
	methods: {
		...mapActions(['getGameDetails']),
	},
	computed: {
		...mapGetters(['currentGame']),
	},
	mounted() {
		const { id } = this.$route.params;
		this.getGameDetails(id).finally(() => (this.loading = false));
	},
};
</script>

<style scoped>
h3 {
	padding-bottom: 20px;
/*	TODO: use vue class pt-16 ? */
}
</style>
