<template>
	<div>
		<MenuBar></MenuBar>
		<hr>
		<router-view></router-view>
	</div>
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