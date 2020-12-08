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

export const api = {
    extSystem: {
        list: extSystemList,
        create: extSystemCreate,
    }
}