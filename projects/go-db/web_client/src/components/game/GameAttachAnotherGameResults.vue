<template>
	<section>
		{{ currentGames }}
	</section>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';

export default {
	name: "GameAttachAnotherGameResults",
	data: () => ({
		loading: true,
	}),
	computed: {
		...mapGetters(['games']),
		currentGames() {
			console.log("GAMES !!!");
			return this.games;
		}
	},
	methods: {
		...mapActions(['getGameList']),
	},
	mounted() {
		this.getGameList().finally(() => this.loading = false)

	}
}
</script>
