<template>

	<section>
		<v-progress-circular
			v-if="loading"
			indeterminate
			color="primary"
		>
		</v-progress-circular>

		<div v-else>
			{{ this.currentGame }}
		</div>
	</section>

</template>

<script>
import {mapActions, mapGetters} from 'vuex';

export default {
	name: "GameDetails",
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