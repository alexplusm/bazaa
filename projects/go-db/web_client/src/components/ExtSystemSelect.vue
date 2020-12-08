<template>
	<section>
		<v-combobox
			v-model="model"
			v-on:change="this.setCurrentExtSystem"
			:items="extSystems"
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
import {mapActions, mapGetters, mapMutations} from "vuex";

export default {
	name: "ExtSystemSelect",
	data() {
		return {model: ''}
	},
	computed: {
		...mapGetters(['extSystems'])
	},
	methods: {
		...mapActions(['getExtSystemList']),
		...mapMutations(['setCurrentExtSystem']),
	},
	mounted() {
		this.getExtSystemList();
	}
}
</script>

<style scoped>
</style>
