<template>

	<section>
		<v-progress-circular
			v-if="loading"
			indeterminate
			color="primary"
		>
		</v-progress-circular>

		<v-row v-else>
			<v-col cols="5">
				<div>Start date: {{ this.currentGame.startDate }}</div>
				<div>Finish date: {{ this.currentGame.finishDate }}</div>
				<div>Answer type: {{ this.currentGame.question.answerType }}</div>
				<div>Question: {{ this.currentGame.question.text }}</div>

				<ul>Options
					<li
						v-for="option in this.currentGame.question.options"
						:key="option.option"
					>
						text: <strong>{{ option.text }}</strong> (option: <strong>{{ option.option }}</strong>)
					</li>
				</ul>

				<ul>Sources
					<li
						v-for="source in this.currentGame.sources"
						:key="source.sourceId"
					>
						{{ source.sourceId }} | {{ source.type }}
					</li>
				</ul>
			</v-col>

			<v-col cols="4">
				AddSource (check box: true - archive | false - schedule)
				<GameUploadFile />
			</v-col>
		</v-row>
	</section>
</template>

<script>
import {mapActions, mapGetters} from 'vuex';
import GameUploadFile from "./GameUploadFile";

export default {
	name: "GameDetails",
	components: {GameUploadFile},
	data() {
		return {
			loading: true,
		}
	},
	methods: {
		...mapActions(['getGameDetails'])
	},
	computed: {
		...mapGetters(['currentGame'])
	},
	mounted() {
		const {id} = this.$route.params;
		this.getGameDetails(id)
			.finally(() => this.loading = false);
	}
}
</script>

<style scoped>
</style>