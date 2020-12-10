<template>
	<ExtSystemSelect :items="extSystems" v-on:change="onChange" />
</template>

<script>
import {mapActions, mapGetters} from "vuex";
import ExtSystemSelect from './ExtSystemSelect';

export default {
	name: "ExtSystemSelectBlock",
	components: {ExtSystemSelect},
	computed: {
		...mapGetters(['extSystems'])
	},
	methods: {
		...mapActions(['getExtSystemList', 'setCurrentExtSystem']),
		onChange(selectedExtSystem) {
			console.log("SelectedExtSystem: ", selectedExtSystem);
			this.setCurrentExtSystem(selectedExtSystem);
		},
	},
	mounted() {
		this.getExtSystemList();
	},
}
</script>

<style scoped>
</style>
