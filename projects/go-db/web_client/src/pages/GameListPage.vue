<template>
	<section>
		<v-row>
			<v-col cols="6">
				<ExtSystemSelectForm />
			</v-col>
		</v-row>

		<h4 v-if="!this.currentExtSystem">Внешняя система не выбрана</h4>
		<GameTable v-else />
	</section>
</template>

<script>
import { mapGetters } from 'vuex';
import GameTable from '../components/game/GameTable';
import ExtSystemSelectForm from '../components/ext-system/ExtSystemSelectForm';

export default {
	name: 'GamesPage',
	components: { GameTable, ExtSystemSelectForm },
	computed: {
		...mapGetters(['currentExtSystem']),
	},
};
</script>
