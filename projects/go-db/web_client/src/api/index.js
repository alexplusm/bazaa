import axios from 'axios';

const processResponse = response => {
    const {data} = response;

    if (data.success) {
        return data.data;
    }
    throw new Error('resp error');
}

const extSystemList = () => {
    return axios.get('/api/v1/ext-system')
        .then(({data}) => data)
        .then(data => {
            if (!data.success) {
                // todo: throw error
            }
            return data.data;
        });
}

const extSystemCreate = extSystem => {
    return axios.post('/api/v1/ext-system', extSystem)
        .then(data => data)
}

const gameList = extSystemId => {
    const params = {extSystemId};

    return axios.get('/api/v1/game', {params})
        .then(processResponse);
}

const gameDetails = (gameId, extSystemId) => {
    const params = {extSystemId};

    return axios.get('/api/v1/game/' + gameId, {params})
        .then(processResponse);
}

const gameCreate = game => {
    return axios.post('/api/v1/game/', game)
        .then(processResponse);
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
    }
}