<template>
	<section>
		<v-progress-circular
			v-if="loading"
			indeterminate
			color="primary"
		>
		</v-progress-circular>

		<v-form v-model="form.valid">
			<v-select
				label="Игра источник"
				v-model="form.model.sourceGameId"
				:items="currentGames"
				item-value="gameId"
				clearable
			>
				<template v-slot:selection="{ item }">
					{{ item.gameId }}
				</template>
				<template v-slot:item="{ item }">
					{{ item.gameId }}
<!--					<v-col>-->
						<!-- TODO: details -->
<!--					</v-col>-->
				</template>
			</v-select>

			<v-select
				v-if="!!form.model.sourceGameId"
				label="Ответ, по которому выбираются скриншоты"
				v-model="form.model.answer"
				:items="currentGameDetails.question.options"
				clearable
			>
			</v-select>

			<v-alert v-if="!!this.error" dense outlined type="error">
				{{ this.error }}
			</v-alert>

			<v-btn :disabled="!formIsValid" @click="submit()"> Отправить </v-btn>
		</v-form>
	</section>
</template>

<script>
import { mapActions, mapGetters } from 'vuex';
import { answerTypesMap } from '../../domain/consts';

export default {
	name: "GameAttachAnotherGameResults",
	data: () => ({
		loading: true,
		form: {
			model: {
				sourceGameId: null,
				answer: null,
			}
		},
		error: null
	}),
	computed: {
		...mapGetters(['games']),
		currentGames() {
			return this.games.filter(game =>
				game.question.answerType === answerTypesMap.categoryType.value &&
				this.$route.params.id !== game.gameId
			);
		},
		currentGameDetails() {
			console.log("games: ", this.games);
			return this.games.find(game => game.gameId === this.form.model.sourceGameId);
		},
		formIsValid() {
			return !!this.form.model.answer && !!this.form.model.sourceGameId
		},
	},
	methods: {
		...mapActions(['getGameList', 'attachAnotherGameResultsToGame']),
		clearForm() {
			this.form.model.sourceGameId = null;
			this.form.model.answer = null;
		},
		submit() {
			if (!this.formIsValid) {
				return;
			}

			const data = {
				gameId: this.$route.params.id,
				formValue: this.form.model,
			};

			this.attachAnotherGameResultsToGame(data).then(
				result => {
					if (result.hasError) {
						this.error = result.errorMessage;
					} else {
						this.clearForm();
					}
				}
			)
		}
	},
	mounted() {
		this.getGameList().finally(() => this.loading = false)
	}
}
</script>
