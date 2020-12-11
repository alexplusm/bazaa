<template>
	<section>
		<v-combobox
			label="Внешняя система"
			v-bind:selected="selected"
			v-on:change="onChange"
			:items="items"
			outlined
		>
			<template v-slot:selection="{ item }">
				{{ item.extSystemId }}
			</template>
			<template v-slot:item="{ item }">
				<v-col>
					<strong>{{ item.extSystemId }}</strong>
					<div>
						{{ item.description }}
						<i>({{ item.postResultsUrl }})</i>
					</div>
					<v-divider></v-divider>
				</v-col>
			</template>
		</v-combobox>
	</section>
</template>

<script>
// TODO: init value: must be select current from store
export default {
	name: 'ExtSystemSelect',
	model: {
		prop: 'selected',
		event: 'change',
	},
	props: {
		selected: Object,
		items: Array,
	},
	methods: {
		onChange(selected) {
			this.$emit('change', { ...selected });
		},
	},
};
</script>
