<template>
	<ExtSystemSelect
		v-model="selectedExtSystem"
		:items="extSystems"
		@input="onChange"
	/>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import ExtSystemSelect from './ExtSystemSelect';

export default {
	name: 'ExtSystemSelectBlock',
	data: () => ({
		selectedExtSystem: null,
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
		this.selectedExtSystem = this.currentExtSystem;
		this.getExtSystemList();
	},
};
</script>
