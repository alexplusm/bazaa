<template>
	<div>
		<ExtSystem />

		<hr>

		<GameList />
	</div>
</template>

<script>
import ExtSystem from "../components/ExtSystem";
import GameList from "../components/game/GameList";

export default {
	name: "GamesPage",
	components: {GameList, ExtSystem},
}
</script>

<style scoped>
</style>
