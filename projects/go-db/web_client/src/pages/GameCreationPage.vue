<template>
	<v-row>
		<v-col cols="5">
			<GameCreateForm />
		</v-col>
	</v-row>
</template>

<script>
import GameCreateForm from '../components/game/GameCreateForm';

export default {
	name: 'GameCreationPage',
	components: { GameCreateForm },
};
</script>

<style scoped></style>
