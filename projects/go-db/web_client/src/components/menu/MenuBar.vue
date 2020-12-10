<template>
	<v-navigation-drawer permanent app width="200px">
		<v-list-item>
			<v-list-item-content>
				<v-list-item-title class="title">
					Admin site
				</v-list-item-title>
				<v-list-item-subtitle> admin </v-list-item-subtitle>
			</v-list-item-content>
		</v-list-item>

		<v-divider></v-divider>

		<v-list dense nav>
			<v-list-item
				v-for="item in items"
				:key="item.title"
				:to="item.href"
			>
				<v-list-item-content>
					<v-list-item-title>{{ item.title }}</v-list-item-title>
				</v-list-item-content>
			</v-list-item>
		</v-list>
	</v-navigation-drawer>
</template>

<script>
export default {
	name: 'MenuBar',
	data() {
		return {
			items: [
				{ title: 'Home', href: '/home/game' },
				{ title: 'Create ext system', href: '/home/ext-system/create' },
				{ title: 'Create game', href: '/home/game-create' },
			],
		};
	},
};
</script>

<style scoped></style>
