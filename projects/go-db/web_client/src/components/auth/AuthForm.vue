<template>
	<v-form v-model="valid">
		<v-text-field
			label="Логин"
			v-model="form.username"
			:rules="fieldRules"
			required
			outlined
		></v-text-field>

		<v-text-field
			label="Пароль"
			v-model="form.password"
			:rules="fieldRules"
			:type="showPassword ? 'text' : 'password'"
			:append-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
			@click:append="showPassword = !showPassword"
			required
			outlined
		></v-text-field>

		<v-alert
			v-if="!!form.error"
			dense
			outlined
			type="error"
		>
			{{ form.error }}
		</v-alert>

		<v-row justify="center">
			<v-progress-circular v-if="loading" indeterminate color="primary">
			</v-progress-circular>
			<v-btn v-else :disabled="!valid" color="success" @click="submit">
				Войти
			</v-btn>
		</v-row>
	</v-form>
</template>

<script>
import { mapActions } from 'vuex';
import { fieldRequiredFunc } from '../../utils/form-utils';

export default {
	name: 'AuthForm',
	data: () => ({
		valid: false,
		loading: false,
		form: {
			username: null,
			password: null,
			error: null,
		},
		showPassword: false,
		fieldRules: [fieldRequiredFunc],
	}),
	methods: {
		...mapActions(['authorize']),
		submit() {
			if (!this.valid) {
				return;
			}
			this.form.error = null;

			this.loading = true;
			this.authorize(this.form)
				.then(value => {
					if (value) {
						this.$router.push('home');
					} else {
						this.form.error = 'Неверные логин и пароль';
					}
				})
				.finally(() => this.loading = false);
		},
	},
};
</script>
