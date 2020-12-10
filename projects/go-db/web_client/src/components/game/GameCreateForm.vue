<template>
	<section>
		<v-form v-model="valid">
			<v-text-field
				label="Название игры"
				v-model="form.name"
				:rules="fieldRules"
				required
			></v-text-field>

			<!-- Select ExtSystemId -->

			<v-text-field
				label="Вопрос"
				v-model="form.question"
				:rules="fieldRules"
				required
			></v-text-field>

			<v-select
				label="Тип ответов"
				v-model="form.answerType"
				:items="answerTypes"
			></v-select>

			<v-row justify="space-around">
				<div>
					<label for="start-date">
						Начало игры
					</label>
					<datetime
						id="start-date"
						input-id="start-date"
						v-model="form.startDate"
						type="datetime"
						:input-class="'rounded-lg pl-4'"
						:input-style="'border: #9E9E9E 1px solid;'"
					></datetime>
				</div>
				<div>
					<label for="end-date">
						Конец игры
					</label>
					<datetime
						id="end-date"
						input-id="end-date"
						v-model="form.endDate"
						type="datetime"
						:input-class="'rounded-lg pl-4'"
						:input-style="'border: #9E9E9E 1px solid;'"
					></datetime>
				</div>

				<div class="mt-4" v-if="!!datesError">
					<v-alert
						dense
						outlined
						type="error"
					>
						{{ datesError }}
					</v-alert>
				</div>
			</v-row>

			<v-row class="mt-12" justify="center">
				<v-btn
					color="success"
					@click="submit"
				>
					Создать игру
				</v-btn>
			</v-row>
		</v-form>
	</section>
</template>

<script>
import {fieldRequiredFunc} from "../../utils/form-utils";
import {answerTypesMap, answerTypesArray} from "../../domain/consts";

// ExtSystemID string `json:"extSystemId"`
// AnswerType  int    `json:"answerType"`
// Options     string `json:"options"`
// StartDate   string `json:"startDate"`
// EndDate     string `json:"endDate"`

export default {
	name: "GameCreateForm",
	data: () => ({
		valid: false,
		form: {
			name: '',
			question: '',
			answerType: answerTypesMap.categoryType.value,
			startDate: null,
			endDate: null,
		},
		datesError: '',
		fieldRules: [fieldRequiredFunc],
		answerTypes: answerTypesArray,
	}),
	methods: {
		submit() {
			console.log("form: ", this.form);
			console.log("start: ", this.form.startDate, typeof(this.form.startDate));
			console.log("end: ", this.form.endDate, typeof(this.form.endDate));

			const startDate = new Date(this.form.startDate);
			console.log("s", startDate);
		}
	}
}
</script>

<style scoped>
</style>