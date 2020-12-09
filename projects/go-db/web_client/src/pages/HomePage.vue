<template>
	<section>
		<MenuBar></MenuBar>

		<router-view></router-view>
	</section>
</template>

<script>
import MenuBar from "../components/MenuBar";

export default {
	name: "HomePage",
	components: {
		MenuBar
	}
}
</script>

<style scoped>
</style>