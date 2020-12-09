<template>
	<section>
		<v-combobox
			v-model="model"
			v-on:change="this.setCurrentExtSystem"
			:items="extSystems"
			label="Choose ext system"
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
		return {model: ''} // todo: remove model !!!
	},
	computed: {
		...mapGetters(['extSystems'])
	},
	methods: {
		...mapActions(['getExtSystemList', 'setCurrentExtSystem']),
	},
	mounted() {
		this.getExtSystemList();
	}
}
</script>

<style scoped>
</style>
