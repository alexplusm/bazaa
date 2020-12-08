<template>
	<section>
		<v-form v-model="valid">
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

			<v-row justify="center">
<!--				<v-progress-circular-->
<!--					v-if="loading"-->
<!--					indeterminate-->
<!--					color="primary"-->
<!--				>-->
<!--				</v-progress-circular>-->
<!--				v-else-->
				<v-btn
					:disabled="!valid"
					color="success"
					@click="submit"
				>
					Create
				</v-btn>
			</v-row>

		</v-form>
	</section>
</template>

<script>
import {fieldRequiredFunc} from "../utils/formUtils";
import {mapActions} from 'vuex'

export default {
	name: "ExtSystemCreate",
	data() {
		return {
			valid: false,
			extSystemId: '',
			description: '',
			postResultsUrl: '',
			fieldRules: [fieldRequiredFunc]
		}
	},
	methods: {
		...mapActions(['createExtSystem']),
		submit() {
			const data = {
				extSystemId: this.extSystemId,
				description: this.description,
				postResultsUrl: this.postResultsUrl,
			}

			// TODO: resp process in actions -> return only true|false
			// если ошибка -> показываем сообщение об ошибке
			// если ок -> сообщение что ОК и рефрешим список систем (другим экшеном?)
			// 		-> очищаем форму
			this.createExtSystem(data).then(resp => console.log("REEESP: ", resp));
		}
	}
}
</script>

<style scoped>

</style>