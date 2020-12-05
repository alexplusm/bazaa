<template>
	<div>
		<div>ext system list</div>
		<div v-for="extSystem in extSystems" :key="extSystem.extSystemId">
			{{ extSystem.extSystemId }} | {{ extSystem.description }} | {{ extSystem.postResultsUrl }}
			<hr>
		</div>
	</div>
</template>

<script>
import axios from "axios"

export default {
	name: "ExtSystemList",
	data() {
		return {
			extSystems: []
		}
	},
	methods: {

		getExtSystemList: function () {
			// TODO: не работает
			axios
				.get('/api/v1/ext-system')
				.then(resp => {
					console.log("resp: ", resp)
					if (resp.data.success) {
						console.log(resp.data.data);
						this.extSystems = resp.data.data
					}
				});
		}
	},
	mounted() {
		this.getExtSystemList()
	}
}
</script>

<style scoped>
</style>
