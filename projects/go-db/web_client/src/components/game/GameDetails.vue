<template>

	<section>
		<v-progress-circular
			v-if="loading"
			indeterminate
			color="primary"
		>
		</v-progress-circular>

		<div v-else>
			<div>Start date: {{ this.currentGame.startDate }}</div>
			<div>Finish date: {{ this.currentGame.finishDate }}</div>
			<div>Answer type: {{ this.currentGame.question.answerType }}</div>
			<div>Question text: {{ this.currentGame.question.text }}</div>
			<div>
				Options
				<div
					v-for="option in this.currentGame.question.options"
					:key="option.option"
				>
					{{ option.option }} | {{ option. text }}
				</div>
			</div>

			<div>
				Sources
				<div
					v-for="source in this.currentGame.sources"
					:key="source.sourceId"
				>
					{{ source.sourceId }} | {{ source.type }}
				</div>
			</div>

			<hr>

			AddSource (check box: true - archive | false - schedule)
			<GameUploadFile />
		</div>
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