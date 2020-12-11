<template>
	<section>
		<v-row>
			<v-col cols="6">
				<ExtSystemSelectBlock />
			</v-col>
		</v-row>
		<GameList />
	</section>
</template>

<script>
import GameList from '../components/game/GameList';
import ExtSystemSelectBlock from '../components/ext-system/ExtSystemSelectBlock';

export default {
	name: 'GamesPage',
	components: { GameList, ExtSystemSelectBlock },
};
</script>
