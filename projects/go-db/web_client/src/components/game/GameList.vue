<template>
	<section>
		{{ this.games }}
	</section>
</template>

<script>
import {mapGetters} from 'vuex'

export default {
	name: "GameList",
	computed: {
		...mapGetters(['games'])
	}
}
</script>

<style scoped>
</style>
