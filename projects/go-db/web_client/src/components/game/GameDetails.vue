<template>
	<section>
		<v-progress-circular v-if="loading" indeterminate color="primary">
		</v-progress-circular>

		<v-row v-else>
			<v-col cols="5">
				<GameInfo :game="currentGame"/>
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
import GameInfo from "./GameInfo";

export default {
	name: 'GameDetails',
	components: { GameUploadFile, GameInfo },
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
