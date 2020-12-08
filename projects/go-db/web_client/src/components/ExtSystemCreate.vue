<template>
	<section>
		<v-form>
			<v-text-field
				label="ID"
				v-model="extSystemId"
			></v-text-field>

			<v-text-field
				label="Description"
				v-model="description"
				:rules="fieldRules"
				required
			></v-text-field>

			<v-text-field
				label="Post Results Url"
				v-model="postResultsUrl"
				:rules="fieldRules"
				required
			></v-text-field>

		</v-form>
	</section>
</template>

<script>
import {fieldRequiredFunc} from "../utils/formUtils";

export default {
	name: "ExtSystemCreate",
	data() {
		return {
			extSystemId: '',
			description: '',
			postResultsUrl: '',
			fieldRules: [fieldRequiredFunc]
		}
	}
}
</script>

<style scoped>

</style>