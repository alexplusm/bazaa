<template>
	<section>
		<v-form v-model="valid">
			<v-text-field
				label="Название игры"
				v-model="form.name"
				:rules="fieldRules"
				required
				outlined
			></v-text-field>

			<ExtSystemSelect :items="extSystems" v-model="form.extSystem" />

			<v-text-field
				label="Вопрос"
				v-model="form.question"
				:rules="fieldRules"
				required
				outlined
			></v-text-field>

			<v-select
				label="Тип ответов"
				v-model="form.answerType"
				:items="answerTypes"
				outlined
			></v-select>

			<div v-if="optionsRequired">
				<v-alert
					type="info"
					outlined
					dense
				>
					Введите возможные ответы через <strong>запятую</strong>.
				</v-alert>

				<v-text-field
					label="Возможные ответы"
					v-model="form.options"
					:rules="fieldRules"
					required
					outlined
				></v-text-field>
			</div>

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
			</v-row>

			<v-alert
				v-if="!!errorMessage"
				class="mt-8"
				type="error"
				outlined
				dense
			>
				{{ errorMessage }}
			</v-alert>

			<v-row class="mt-8" justify="center">
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
import {mapActions, mapGetters} from "vuex";
import {isValid, isAfter} from 'date-fns'
import ExtSystemSelect from '../../components/ext-system/ExtSystemSelect';
import {fieldRequiredFunc} from "../../utils/form-utils";
import {answerTypesMap, answerTypesArray} from "../../domain/consts";
import {createGameToDTO} from "../../domain/dto";


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
	components: {ExtSystemSelect},
	data: () => ({
		valid: false,
		form: {
			name: '',
			question: '',
			extSystem: null,
			answerType: answerTypesMap.categoryType.value,
			options: '',
			startDate: null,
			endDate: null,
		},
		errorMessage: '',
		fieldRules: [fieldRequiredFunc],
		answerTypes: answerTypesArray,
	}),
	computed: {
		...mapGetters(['extSystems']),
		optionsRequired() {
			return this.form.answerType === answerTypesMap.categoryType.value
		}
	},
	methods: {
		...mapActions(['getExtSystemList', 'setCurrentExtSystem', 'createGame']),
		clearForm() {
			// TODO: clear Form
		},
		submit() {
			const options = this.form.options.split(',').map(s => s.trim()).filter(s => s !== "");

			if (this.optionsRequired && options.length === 0) {
				this.errorMessage = 'Возможные ответы введены не корректно';
				return;
			}
			this.form.options = options.join(',');

			const res = processDates(this.form.startDate, this.form.endDate);
			this.errorMessage = res.error;

			if (res.error != null) return;

			const dto = createGameToDTO({...this.form})

			console.log("123", dto);

			this.createGame(dto)
				.then(() => this.clearForm())
				.then(resp => console.log("RESSSSP: ", resp));
		}
	},
	mounted() {
		this.getExtSystemList();
	},
}
</script>
