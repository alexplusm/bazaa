<template>
	<section>
		<v-file-input
			label="Прикрепленный файл"
			v-model="archive"
			:disabled="loading"
			show-size
			counter
			required
		></v-file-input>

		<v-btn :disabled="!archive" @click="upload"> Загрузить файл </v-btn>
		<v-progress-circular
			v-if="loading"
			:rotate="360"
			:size="50"
			:width="8"
			:value="progress"
			color="primary"
		>
			{{ progress }}
		</v-progress-circular>
	</section>
</template>

<script>
import { mapActions } from 'vuex';

function progressCallback(progressEvent) {
	const percentCompleted = Math.round(
		(progressEvent.loaded * 100) / progressEvent.total
	);
	if (percentCompleted <= 95) {
		this.progress = percentCompleted;
	}
}

export default {
	name: 'GameAttachArchive',
	data: () => ({
		archive: null,
		loading: false,
		progress: 0,
	}),
	methods: {
		...mapActions(['updateGameWithArchive']),
		upload() {
			const { id } = this.$route.params;
			const skip = this.loading || !this.archive || !id;

			if (skip) return;

			const payload = {
				gameId: id,
				file: this.archive,
				progressCallback: progressCallback.bind(this),
			};

			this.loading = true;
			this.updateGameWithArchive(payload)
				.then(() => {
					this.archive = null;
				})
				.catch((err) => console.log('err: ', err))
				.finally(() => {
					this.loading = false;
					this.progress = 0;
				});
		},
	},
};
</script>
