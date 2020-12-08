<template>
	<v-container>
		<v-row>
			<v-col cols="6">
				<ExtSystemSelect />
			</v-col>
			<v-col cols="6">
				<ExtSystemCreate />
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
import ExtSystemSelect from "./ExtSystemSelect";
import ExtSystemCreate from "./ExtSystemCreate";

export default {
	name: "ExtSystem",
	components: {
		ExtSystemCreate,
		ExtSystemSelect
	}
}
</script>

<style scoped>
</style>
