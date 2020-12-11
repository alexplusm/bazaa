<template>
	<v-row justify="center">
		<v-col cols="3" align="center">
			<h3 class="mb-6">Портал Администратора</h3>
			<AuthForm />
		</v-col>
	</v-row>
</template>

<script>
import AuthForm from '../components/auth/AuthForm';

export default {
	name: 'AuthPage',
	components: { AuthForm },
};
</script>
