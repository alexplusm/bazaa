<template>
	<section>
		<v-combobox
			v-model="model"
			v-on:change="this.updateCurrentExtSystem"
			:items="EXT_SYSTEMS"
			label="ExtSystems"
			solo
		>
			<template v-slot:selection="{ item }">
				{{ item.extSystemId }}
			</template>
			<template v-slot:item="{ item }">
				<div>
					{{ item.extSystemId}}
					<p>{{ item.description}}</p>
					<div>{{ item.postResultsUrl }}</div>
				</div>
			</template>
		</v-combobox>
	</section>
</template>

<script>
import {mapActions, mapGetters} from "vuex";

export default {
	name: "ExtSystemSelect",
	data() {
		return {model: ''}
	},
	computed: {
		...mapGetters(['EXT_SYSTEMS'])
	},
	methods: {
		...mapActions([
			'GET_EXT_SYSTEMS_FROM_API'
		]),
		updateCurrentExtSystem(extSystem) {
			console.log("extSystem", extSystem);
			// TODO: dispatch action
		}
	},
	mounted() {
		this.GET_EXT_SYSTEMS_FROM_API();
	}
}
</script>

<style scoped>
</style>
