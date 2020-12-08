<template>
	<v-container>
		<v-row justify="center">

			<v-col cols="3">

				<v-row justify="center">
					<h3>Админка</h3>
				</v-row>

				<!-- TODO: FORM Into other component -->
				<v-form v-model="valid">

					<v-text-field
						label="Login"
						required
						v-model="login"
						:rules="loginRules"
					></v-text-field>

					<v-text-field
						label="Password"
						required
						v-model="password"
						:rules="passwordRules"
					></v-text-field>

					<v-row justify="center">
						<v-progress-circular
							v-if="loading"
							indeterminate
							color="primary"
						>
						</v-progress-circular>
						<v-btn
							v-else
							:disabled="!valid"
							color="success"
							@click="submit"
						>
							Log in
						</v-btn>
					</v-row>
				</v-form>
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
import {fieldRequiredFunc} from "../utils/form-utils";

function timeout(ms) {
	return new Promise(resolve => setTimeout(resolve, ms));
}

export default {
	name: "AuthPage",
	data: () => ({
		valid: false,
		loading: false,
		login: '',
		password: '',
		loginRules: [fieldRequiredFunc],
		passwordRules: [fieldRequiredFunc],
	}),
	methods: {
		authorize() {
			this.$store.commit('authorize');
		},
		submit() {
			console.log("valid: ", this.valid)
			if (!this.valid) {
				return;
			}

			this.loading = true;

			timeout(3)
				.then(() => {
					console.log(this.password, this.login)
				})
				.then(() => {
					this.loading = false;
					this.authorize();
					this.$router.push('home');
				});
		},
	}
}
</script>

<style scoped>
</style>
