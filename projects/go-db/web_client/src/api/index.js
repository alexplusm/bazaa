import axios from 'axios';

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
        .then(resp => console.log("GAME LIST: ", resp));
}

export const api = {
    extSystem: {
        list: extSystemList,
        create: extSystemCreate,
    },
    game: {
        list: gameList
    }
}