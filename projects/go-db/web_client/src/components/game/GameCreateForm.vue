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
import {isValid, isAfter} from 'date-fns'

// ExtSystemID string `json:"extSystemId"`
// AnswerType  int    `json:"answerType"`
// Options     string `json:"options"`
// StartDate   string `json:"startDate"`
// EndDate     string `json:"endDate"`

// TODO: in utils !!!
function processDates(start, end) {
	const result = { start: null, end: null,  error: null };

	if (!start) {
		result.error = 'Необходимо ввести начало игры';
		return result;
	}
	if (!end) {
		result.error = 'Необходимо ввести конец игры';
		return result;
	}
	const startDate = new Date(start);
	const endDate = new Date(end);
	if (!isValid(startDate)) {
		result.error = 'Не правильный формат начала игры';
		return result;
	}
	if (!isValid(endDate)) {
		result.error = 'Не правильный формат конца игры';
		return result;
	}

	const now = new Date();

	if (isAfter(now, startDate)) {
		result.error = 'Начало игры не может быть в прошлом';
		return result;
	}
	if (isAfter(startDate, endDate)) {
		result.error = 'Начало игры не может после конца игры';
		return result;
	}
	result.start = startDate;
	result.end = endDate;

	return result;
}

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

			const res = processDates(this.form.startDate, this.form.endDate);
			this.datesError = res.error;

			console.log("res: ", res);

			if (res.error != null) return;

			console.log("123");
		}
	}
}
</script>

<style scoped>
</style>