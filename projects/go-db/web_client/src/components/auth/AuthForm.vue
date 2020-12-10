<template>
	<v-form v-model="valid">
		<v-text-field
			label="Login"
			required
			v-model="login"
			:rules="fieldRules"
			outlined
		></v-text-field>

		<v-text-field
			label="Password"
			required
			v-model="password"
			:rules="fieldRules"
			outlined
		></v-text-field>

		<v-row justify="center">
			<v-progress-circular v-if="loading" indeterminate color="primary">
			</v-progress-circular>
			<v-btn v-else :disabled="!valid" color="success" @click="submit">
				Log in
			</v-btn>
		</v-row>
	</v-form>
</template>

<script>
import { mapMutations } from 'vuex';
import { fieldRequiredFunc } from '../../utils/form-utils';
import { timeout } from '../../utils/test';

export default {
	name: 'AuthForm',
	data: () => ({
		valid: false,
		loading: false,
		login: 'a',
		password: 'b',
		fieldRules: [fieldRequiredFunc],
	}),
	methods: {
		...mapMutations(['authorize']),
		submit() {
			if (!this.valid) {
				return;
			}

			this.loading = true;
			timeout(3).then(() => {
				this.loading = false;
				this.authorize();
				this.$router.push('home');
			});
		},
	},
};
</script>
