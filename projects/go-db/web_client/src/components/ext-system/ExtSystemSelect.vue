<template>
	<section>
		<v-select
			label="Внешняя система"
			:value="selectedItem"
			:items="items"
			item-value="extSystemId"
			@change="onChange"
			clearable
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
				</v-col>
			</template>
		</v-select>
	</section>
</template>

<script>
export default {
	name: 'ExtSystemSelect',
	model: {
		prop: 'selected',
		event: 'change',
	},
	props: {
		selected: Object || null,
		items: {
			type: Array,
			default: () => [],
		},
	},
	computed: {
		selectedItem() {
			return this.selected ? this.selected.extSystemId : null;
		},
	},
	methods: {
		onChange(selected) {
			const value = this.items.find(
				(item) => item.extSystemId === selected
			);
			this.$emit('change', value);
		},
	},
};
</script>
