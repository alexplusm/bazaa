<template>
	<div id="app">
		<MenuBar></MenuBar>
		<router-view></router-view>
	</div>
</template>

<script>
import MenuBar from "./components/MenuBar";
import axios from "axios"

export default {
    name: 'App',
	components: { MenuBar },
    mounted () {
        axios
            .get('/api/check/alive')
            .then(response => console.log("response: ", response))
    }
}
</script>
