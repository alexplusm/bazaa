<template>
	<v-form v-model="valid">
		<v-text-field
			label="Логин"
			v-model="form.login"
			:rules="fieldRules"
			required
			outlined
		></v-text-field>

		<v-text-field
			label="Пароль"
			v-model="form.password"
			:rules="fieldRules"
			required
			outlined
		></v-text-field>

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
import { mapMutations } from 'vuex';
import { fieldRequiredFunc } from '../../utils/form-utils';
import { timeout } from '../../utils/test';

const requiredLogin = 'AzUseRadm';
const requiredPass = 'Qca05+Bz)3';

export default {
	name: 'AuthForm',
	data: () => ({
		valid: false,
		loading: false,
		form: {
			login: '',
			password: '',
		},
		fieldRules: [fieldRequiredFunc],
	}),
	methods: {
		...mapMutations(['authorize']),
		clearForm() {
			this.form.login = '';
			this.form.password = '';
		},
		submit() {
			if (!this.valid) {
				return;
			}

			this.loading = true;
			timeout(1500).then(() => {
				this.loading = false;

				if (this.form.login === requiredLogin && this.form.password === requiredPass) {
					this.authorize();
					this.$router.push('home');
				} else {
					this.clearForm();
				}
			});
		},
	},
};
</script>
