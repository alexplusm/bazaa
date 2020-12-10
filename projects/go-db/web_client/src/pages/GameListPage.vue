<template>
	<section>
		<v-row>
			<v-col cols="6">
				<ExtSystemSelectBlock />
			</v-col>
		</v-row>
<!--		<ExtSystem />-->
		<GameList />
	</section>
</template>

<script>
// import ExtSystem from '../components/ext-system/ExtSystem';
import GameList from '../components/game/GameList';
import ExtSystemSelectBlock from '../components/ext-system/ExtSystemSelectBlock';

export default {
	name: 'GamesPage',
	components: { GameList,
		// ExtSystem ,
		ExtSystemSelectBlock},
};
</script>
