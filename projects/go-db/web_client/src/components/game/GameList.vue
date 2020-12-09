<template>
	<v-simple-table>
		<template v-slot:default>
			<thead>
			<tr>
				<th class="text-left">ID</th>
				<th class="text-left">Name</th>
				<th class="text-left">Status</th>
				<th class="text-left">From</th>
				<th class="text-left">To</th>
			</tr>
			</thead>
			<tbody>
			<tr
				v-for="item in games"
				:key="item.gameId"
			>
				<td>
					<router-link :to="{ path: item.gameId }" append>
						{{ item.gameId }}
					</router-link>
				</td>
				<td>{{ item.name }}</td>
				<td>{{ item.status }}</td>
				<td>{{ item.from }}</td>
				<td>{{ item.to }}</td>
			</tr>
			</tbody>
		</template>
	</v-simple-table>
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
