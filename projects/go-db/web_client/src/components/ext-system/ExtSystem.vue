<template>
	<v-container>
		<v-row>
			<v-col cols="6">
				<ExtSystemSelect />
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
import ExtSystemSelect from "./ExtSystemSelect";

export default {
	name: "ExtSystem",
	components: {ExtSystemSelect}
}
</script>

<style scoped>
</style>
