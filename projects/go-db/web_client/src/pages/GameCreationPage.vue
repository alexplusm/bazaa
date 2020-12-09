<template>
	<v-row>
		<v-col cols="4">
			<GameCreateForm />
		</v-col>
	</v-row>
</template>

<script>
import GameCreateForm from '../components/game/GameCreateForm';

export default {
	name: "GameCreationPage",
	components: {GameCreateForm},
}
</script>

<style scoped>
</style>