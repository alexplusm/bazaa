<template>
	<v-container>
		<v-row justify="center">
			<ExtSystemCreate class="form"/>
		</v-row>
	</v-container>
</template>

<script>
import ExtSystemCreate from "../components/ExtSystemCreate";

export default {
	name: "ExtSystemCreatePage",
	components: {ExtSystemCreate},
}
</script>

<style scoped>
.form {
	padding-top: 20px;
	width: 400px;
}
</style>