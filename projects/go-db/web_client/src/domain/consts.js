// INFO: sync with backend
const textType = {
	value: 1,
	text: 'Текстовый',
};

const categoryType = {
	value: 2,
	text: 'Категориальный',
};

const rectangleCoordsType = {
	value: 3,
	text: 'Координаты прямоугольника',
};

const polygonCoordsType = {
	value: 4,
	text: 'Координаты точек полигонов',
};

export const answerTypesMap = {
	textType,
	categoryType,
	rectangleCoordsType,
	polygonCoordsType,
};

export const answerTypesArray = [
	textType,
	categoryType,
	rectangleCoordsType,
	polygonCoordsType,
];

const archiveSourceType = {
	value: 1,
	text: 'Архив',
}

const scheduleSourceType = {
	value: 2,
	text: 'Расписание'
}

const anotherGameSourceType = {
	value: 3,
	text: 'Другая игра'
}

export const sourceTypesMap = {
	archiveSourceType,
	scheduleSourceType,
	anotherGameSourceType
}

export const sourceTypesArray = [
	archiveSourceType,
	scheduleSourceType,
	anotherGameSourceType
]
