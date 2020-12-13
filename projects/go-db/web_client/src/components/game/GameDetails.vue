<template>
	<section>
		<v-progress-circular v-if="loading" indeterminate color="primary">
		</v-progress-circular>

		<v-row v-else>
			<v-col cols="5">
				<GameInfo :game="currentGame"/>
			</v-col>

			<v-col cols="4">
				<GameAddSources />
			</v-col>
		</v-row>
	</section>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import GameInfo from "./GameInfo";
import GameAddSources from "./GameAddSources"

export default {
	name: 'GameDetails',
	components: { GameInfo, GameAddSources },
	data: () => ({
		loading: true,
	}),
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
