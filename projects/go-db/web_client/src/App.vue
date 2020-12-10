<template>
	<v-app>
		<v-main>
			<section>
				<router-view></router-view>
			</section>
		</v-main>
	</v-app>
</template>

<script>
export default {
	name: 'App',
};
</script>

<style scoped>
section {
	padding: 10px;
}
</style>
