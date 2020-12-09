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
import axios from 'axios';

export default {
    name: 'App',
    mounted() {
        axios.get('/api/check/alive')
            .then(response => console.log("response: ", response))
    }
}
</script>

<style scoped>
section {
	padding: 10px;
}
</style>
