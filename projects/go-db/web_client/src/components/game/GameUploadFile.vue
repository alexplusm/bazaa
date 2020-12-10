<template>
	<section>
		<v-file-input
			label="Attache archive"
			v-model="archive"
			:disabled="loading"
			show-size
			counter
			required
		></v-file-input>

		<v-btn :disabled="!archive" @click="upload"> Upload archive </v-btn>
		<v-progress-circular v-if="loading" indeterminate color="primary" />
	</section>
</template>

<script>
import { mapActions } from 'vuex';

export default {
	name: 'GameUploadFile',
	data: () => ({
		archive: null,
		loading: false,
	}),
	methods: {
		...mapActions(['updateGameWithArchive']),
		upload() {
			const { id } = this.$route.params;
			const skip = this.loading || !this.archive || !id;

			if (skip) return;

			const payload = { gameId: id, file: this.archive };

			this.loading = true;
			this.updateGameWithArchive(payload)
				.then(() => {
					this.archive = null;
				})
				.catch((err) => console.log('err: ', err))
				.finally(() => (this.loading = false));
		},
	},
};
</script>
