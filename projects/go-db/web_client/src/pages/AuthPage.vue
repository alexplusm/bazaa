<template>
	<v-container>
		<v-row justify="center">
			<v-col cols="3">
				<v-row justify="center">
					<h3>Админка</h3>
				</v-row>

				<AuthForm />
			</v-col>
		</v-row>
	</v-container>
</template>

<script>
import AuthForm from '../components/auth/AuthForm';

export default {
	name: 'AuthPage',
	components: { AuthForm },
};
</script>

<style scoped></style>
