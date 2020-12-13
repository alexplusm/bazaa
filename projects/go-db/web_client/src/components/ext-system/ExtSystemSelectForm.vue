<template>
	<ExtSystemSelect
		v-model="selectedExtSystem"
		:items="extSystems"
		@input="onInput"
	/>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import ExtSystemSelect from './ExtSystemSelect';

export default {
	name: 'ExtSystemSelectForm',
	data: () => ({
		selectedExtSystem: null,
	}),
	components: { ExtSystemSelect },
	computed: {
		...mapGetters(['extSystems', 'currentExtSystem']),
	},
	methods: {
		...mapActions(['getExtSystemList', 'setCurrentExtSystem']),
		onInput(selectedExtSystem) {
			this.setCurrentExtSystem(selectedExtSystem);
		},
	},
	mounted() {
		this.selectedExtSystem = this.currentExtSystem;
		this.getExtSystemList();
	},
};
</script>
