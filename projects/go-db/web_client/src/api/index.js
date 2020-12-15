import axios from 'axios';

const processResponse = (response) => {
	const { data } = response;

	if (data.success) {
		return data.data;
	}
	throw new Error('resp error');
};

const extSystemList = ({ username, password }) => {
	const auth = { username, password };

	return axios
		.get('/api/v1/ext_system', { auth })
		.then(({ data }) => data)
		.then((data) => {
			if (!data.success) {
				// todo: throw error
			}
			return data.data;
		});
};

const extSystemCreate = (extSystem, { username, password }) => {
	const auth = { username, password };
	return axios
		.post('/api/v1/ext_system', extSystem, { auth })
		.then((data) => data);
};

const gameList = (extSystemId, { username, password }) => {
	const params = { extSystemId };
	const auth = { username, password };

	return axios.get('/api/v1/game', { params, auth }).then(processResponse);
};

const gameDetails = (gameId, extSystemId, { username, password }) => {
	const params = { extSystemId };
	const auth = { username, password };

	return axios
		.get('/api/v1/game/' + gameId, { params, auth })
		.then(processResponse);
};

const gameCreate = (game, { username, password }) => {
	const auth = { username, password };
	return axios.post('/api/v1/game', game, { auth }).then(processResponse);
};

const gameUpdateWithArchive = (
	gameId,
	file,
	onUploadProgress,
	{ username, password }
) => {
	const auth = { username, password };
	const config = { onUploadProgress, auth };
	const formData = new FormData();
	formData.append('archives', file);

	return axios.put('/api/v1/game/' + gameId, formData, config);
};

const checkCredentials = (credentials) => {
	const auth = {
		username: credentials.username,
		password: credentials.password,
	};

	return axios
		.get('/api/v1/check/alive', { auth })
		.then(() => true)
		.catch((err) => {
			if (err.response.status === 401) {
				return false;
			} else {
				console.error('auth error: ', err);
				return false;
			}
		});
};

export const api = {
	extSystem: {
		list: extSystemList,
		create: extSystemCreate,
	},
	game: {
		list: gameList,
		details: gameDetails,
		create: gameCreate,
		updateWithFile: gameUpdateWithArchive,
	},
	auth: {
		check: checkCredentials,
	},
};
