<template>
	<section>
		<v-form v-model="valid">
			<v-text-field
				label="ID"
				v-model="form.extSystemId"
				outlined
			></v-text-field>

			<v-text-field
				label="Описание"
				v-model="form.description"
				:rules="fieldRules"
				required
				outlined
			></v-text-field>

			<v-text-field
				label="Url для отправки результатов"
				v-model="form.postResultsUrl"
				:rules="fieldRules"
				required
				outlined
			></v-text-field>

			<v-row justify="center">
				<v-btn :disabled="!valid" color="success" @click="submit">
					Создать
				</v-btn>
			</v-row>
		</v-form>
	</section>
</template>

<script>
import { mapActions } from 'vuex';
import { fieldRequiredFunc } from '../../utils/form-utils';

export default {
	name: 'ExtSystemCreate',
	data() {
		return {
			valid: false,
			form: {
				extSystemId: '',
				description: '',
				postResultsUrl: '',
			},
			fieldRules: [fieldRequiredFunc],
		};
	},
	methods: {
		...mapActions(['createExtSystem', 'getExtSystemList']),
		clearForm() {
			this.form.extSystemId = '';
			this.form.description = '';
			this.form.postResultsUrl = '';
		},
		submit() {
			const data = {
				extSystemId: this.form.extSystemId,
				description: this.form.description,
				postResultsUrl: this.form.postResultsUrl,
			};

			// TODO: resp process in actions -> return only true|false
			this.createExtSystem(data)
				.then((resp) => console.log('REEESP: ', resp))
				.then(() => this.getExtSystemList()) // TODO: process result?
				.then(() => this.clearForm()); // TODO: form validation error
		},
	},
};
</script>
