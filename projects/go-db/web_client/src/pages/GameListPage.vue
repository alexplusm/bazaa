<template>
	<section>
		<v-row>
			<v-col cols="6">
				<ExtSystemSelectForm />
			</v-col>
		</v-row>
		<GameList />
	</section>
</template>

<script>
import GameList from '../components/game/GameList';
import ExtSystemSelectForm from '../components/ext-system/ExtSystemSelectForm';

export default {
	name: 'GamesPage',
	components: { GameList, ExtSystemSelectForm },
};
</script>
