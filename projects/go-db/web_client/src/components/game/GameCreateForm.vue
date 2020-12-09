<template>
	<section>
		<v-form v-model="valid">

			<v-text-field
				label="Name"
				v-model="form.name"
				:rules="fieldRules"
				required
			></v-text-field>

			<!-- Select ExtSystemId -->

			<v-text-field
				label="Question"
				v-model="form.question"
				:rules="fieldRules"
				required
			></v-text-field>

			<v-select
				label="Answer type"
				v-model="form.answerType"
				:items="answerTypes"
			></v-select>

			<!-- -->
			<v-row>
				<v-menu
					ref="menu"
					v-model="menu"
					:close-on-content-click="false"
					:return-value.sync="date"
					transition="scale-transition"
					offset-y
					min-width="290px"
				>
					<template v-slot:activator="{ on, attrs }">
						<v-text-field
							v-model="date"
							label="Picker in menu"
							prepend-icon="mdi-calendar"
							readonly
							v-bind="attrs"
							v-on="on"
						></v-text-field>
					</template>
					<v-date-picker
						v-model="date"
						no-title
						scrollable
					>
						<v-spacer></v-spacer>
						<v-btn
							text
							color="primary"
							@click="menu = false"
						>
							Cancel
						</v-btn>
						<v-btn
							text
							color="primary"
							@click="$refs.menu.save(date)"
						>
							OK
						</v-btn>
					</v-date-picker>
				</v-menu>
			</v-row>
			<!-- -->

			<v-row
				justify="space-around"
				align="center"
			>
				<v-col style="width: 350px; flex: 0 1 auto;">
					<h4>Start:</h4>
					<v-time-picker
						v-model="form.startDate"
						:max="form.endDate"
						format="24hr"
					></v-time-picker>
				</v-col>
				<v-col style="width: 350px; flex: 0 1 auto;">
					<h4>End:</h4>
					<v-time-picker
						v-model="form.endDate"
						:min="form.startDate"
						format="24hr"
					></v-time-picker>
				</v-col>
			</v-row>

		</v-form>
	</section>
</template>

<script>
import {fieldRequiredFunc} from "../../utils/form-utils";

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
			answerType: 0, // TODO: default value
			startDate: null,
			endDate: null,
		},

		menu: false, // TODO: K EEE K
		date: new Date().toISOString().substr(0, 10),

		fieldRules: [fieldRequiredFunc],
		answerTypes: [
			{value: 0, text: 'kek'},
			{value: 1, text: 'lol'},
			{value: 2, text: 'azaza'},
			{value: 3, text: 'lelelele'},
		]
	}),
}
</script>

<style scoped>
</style>