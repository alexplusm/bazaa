<template>
	<v-navigation-drawer permanent app width="280px">
		<v-list-item>
			<v-list-item-content>
				<v-list-item-title class="title">
					Портал Администратора
				</v-list-item-title>
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
				{ title: 'Главная', href: '/home/game' },
				{
					title: 'Создать внешнюю систему',
					href: '/home/ext-system/create',
				},
				{ title: 'Создать игру', href: '/home/game-create' },
			],
		};
	},
};
</script>
