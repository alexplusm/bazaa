<template>
	<v-simple-table>
		<template v-slot:default>
			<thead>
				<tr>
					<th class="text-left">ID</th>
					<th class="text-left">Название</th>
					<th class="text-left">Статус</th>
					<th class="text-left">Начало</th>
					<th class="text-left">Конец</th>
				</tr>
			</thead>
			<tbody>
				<tr v-for="item in games" :key="item.gameId">
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
import { mapGetters } from 'vuex';

export default {
	name: 'GameTable',
	computed: {
		...mapGetters(['games']),
	},
};
</script>
