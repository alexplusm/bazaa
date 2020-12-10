<template>
	<ExtSystemSelect
		:items="extSystems"
		v-on:change="onChange"
	/>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import ExtSystemSelect from './ExtSystemSelect';

export default {
	name: 'ExtSystemSelectBlock',
	data: () => ({
		val: null
	}),
	components: { ExtSystemSelect },
	computed: {
		...mapGetters(['extSystems', 'currentExtSystem']),
	},
	methods: {
		...mapActions(['getExtSystemList', 'setCurrentExtSystem']),
		onChange(selectedExtSystem) {
			this.setCurrentExtSystem(selectedExtSystem);
		},
	},
	mounted() {
		// this.val = {...this.currentExtSystem};
		//
		// console.log("currentExtSystem: ", this.currentExtSystem);

		this.getExtSystemList();
	},
};
</script>
