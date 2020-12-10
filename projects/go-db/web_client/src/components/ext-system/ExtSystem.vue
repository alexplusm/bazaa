<template>
<!--	<v-container>-->
		<v-row>
			<v-col cols="6">
				<ExtSystemSelectBlock />
			</v-col>
		</v-row>
<!--	</v-container>-->
</template>

<script>
import ExtSystemSelectBlock from './ExtSystemSelectBlock';

export default {
	name: 'ExtSystem',
	components: { ExtSystemSelectBlock },
};
</script>

<style scoped></style>
