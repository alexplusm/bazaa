<template>
	<v-row>
		<v-col cols="4">
			<ExtSystemCreate />
		</v-col>
	</v-row>
</template>

<script>
import ExtSystemCreate from "../components/ext-system/ExtSystemCreate";

export default {
	name: "ExtSystemCreatePage",
	components: {ExtSystemCreate},
}
</script>

<style scoped>
</style>