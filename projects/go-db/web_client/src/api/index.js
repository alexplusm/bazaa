import axios from 'axios';

const processResponse = (response) => {
	const { data } = response;

	if (data.success) {
		return data.data;
	}
	throw new Error('resp error');
};

const extSystemList = () => {
	return axios
		.get('/api/v1/ext_system')
		.then(({ data }) => data)
		.then((data) => {
			if (!data.success) {
				// todo: throw error
			}
			return data.data;
		});
};

const extSystemCreate = (extSystem) => {
	return axios.post('/api/v1/ext_system', extSystem).then((data) => data);
};

const gameList = (extSystemId) => {
	const params = { extSystemId };

	return axios.get('/api/v1/game', { params }).then(processResponse);
};

const gameDetails = (gameId, extSystemId) => {
	const params = { extSystemId };

	return axios
		.get('/api/v1/game/' + gameId, { params })
		.then(processResponse);
};

const gameCreate = (game) => {
	return axios.post('/api/v1/game/', game).then(processResponse);
};

const gameUpdateWithArchive = (gameId, file, onUploadProgress) => {
	const config = { onUploadProgress };
	const formData = new FormData();
	formData.append('archives', file);

	return axios.put('/api/v1/game/' + gameId, formData, config);
};

const checkCredentials = (credentials) => {
	const auth = {
		username: credentials.username,
		password: credentials.password,
	}

	return axios.get('/api/v1/check/alive', {auth})
		.then(() => true)
		.catch(err => {
			if (err.response.status === 401) {
				return false;
			} else {
				console.error("auth error: ", err);
				return false;
			}
		});
}

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
		check: checkCredentials
	}
};
