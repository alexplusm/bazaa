<template>
	<div>
		GamesPage

		<ExtSystem />


<!--	select	-->
<!--		<v-select-->
<!--			:items="this.EXT_SYSTEMS"-->
<!--			:item-value="'extSystemId'"-->
<!--			:item-text="this.getText"-->
<!--			label="ExtSystem"-->
<!--		></v-select>-->

	</div>
</template>

<script>
import ExtSystem from "../components/ExtSystem";


export default {
	name: "GamesPage",
	components: {ExtSystem},
	computed: {
		// ...mapGetters(['EXT_SYSTEMS'])
	},
	methods: {
		// ...mapActions([
		// 	'GET_EXT_SYSTEMS_FROM_API'
		// ]),
		// getText(extSys) {
		// 	return extSys.extSystemId + ' | ' + extSys.description + ' | ' + extSys.postResultsUrl
		// }
	},
	mounted() {
		// this.GET_EXT_SYSTEMS_FROM_API();
	}
}
</script>

<style scoped>
</style>
