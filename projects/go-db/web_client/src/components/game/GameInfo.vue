<template>
	<section>
		<div>Начало игры: {{ game.startDate }}</div>
		<div>Конец игры: {{ game.finishDate }}</div>
		<div>Тип ответов: {{ game.question.answerType }}</div>
		<div>Вопрос: {{ game.question.text }}</div>

		<ul>
			Опции
			<li v-for="option in game.question.options" :key="option.option">
				text: <strong>{{ option.text }}</strong> (option:
				<strong>{{ option.option }}</strong
				>)
			</li>
		</ul>

		<ul>
			Источники
			<li v-for="source in game.sources" :key="source.sourceId">
				{{ source.sourceId }} | {{ source.type }}
			</li>
		</ul>
	</section>
</template>

<script>
export default {
	name: 'GameInfo',
	props: {
		game: Object,
	},
};
</script>
