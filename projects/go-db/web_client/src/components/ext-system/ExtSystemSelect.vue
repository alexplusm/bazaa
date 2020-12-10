<template>
	<section>
		<v-combobox
			label="Выберите внешнюю систему"
			v-bind:selected="selected"
			v-on:change="onChange"
			:items="items"
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

export default {
	name: "ExtSystemSelect",
	model: {
		prop: 'selected',
		event: 'change'
	},
	props: {
		selected: Object,
		items: Array
	},
	methods: {
		onChange(selected) {
			this.$emit('change', {...selected});
		}
	}
}
</script>

<style scoped>
</style>
