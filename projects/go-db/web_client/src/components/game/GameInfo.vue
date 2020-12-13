<template>
	<section>
		<div>Start date: {{ game.startDate }}</div>
		<div>Finish date: {{ game.finishDate }}</div>
		<div>
			Answer type: {{ game.question.answerType }}
		</div>
		<div>Question: {{ game.question.text }}</div>

		<ul>
			Options
			<li
				v-for="option in game.question.options"
				:key="option.option"
			>
				text: <strong>{{ option.text }}</strong> (option:
				<strong>{{ option.option }}</strong
				>)
			</li>
		</ul>

		<ul>
			Sources
			<li
				v-for="source in game.sources"
				:key="source.sourceId"
			>
				{{ source.sourceId }} | {{ source.type }}
			</li>
		</ul>
	</section>
</template>

<script>
export default {
	name: "GameInfo",
	props: {
		game: Object
	}
}
</script>

<style scoped>

</style>