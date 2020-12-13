<template>
	<section>
		<v-combobox
			label="Внешняя система"
			:value="selected"
			:items="items"
			item-value="extSystemId"
			@input="onInput"
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
export default {
	name: 'ExtSystemSelect',
	model: {
		prop: 'selected',
		event: 'input',
	},
	props: {
		selected: Object||null,
		items: {
			type: Array,
			default: () => []
		},
	},
	methods: {
		onInput(selected) {
			this.$emit('input', { ...selected });
		},
	},
};
</script>
