import http from "../utils/http-client";

const login = (data) => {
    return http.post('/login', data, {
        transformResponse: [(result) => {
            const parsed = JSON.parse(result);
            localStorage.setItem('authUser', JSON.stringify(parsed));
            return parsed;
        }]
    });
}

const register = (data) => {
    return http.post('/signup', JSON.stringify(data))
}

const profile = () => {
    return http.get('/user');
}

const logout = () => {
    return http.get('/logout', null, {
        transformResponse: [(result) => {
            localStorage.removeItem('authUser');
            return JSON.parse(result);
        }]
    });
}

const createPoll = (data) => {
    return http.post('/createPoll', JSON.stringify(data));
}

const getPolls = () => {
    return http.get('/polls', null, {
        transformResponse: [(result) => {
            return JSON.parse(result);
        }]
    });
}

const updateVote = (data) => {
    return http.post('/updateVote', JSON.stringify(data));
}

const getAuthUser = () => {
    return JSON.parse(localStorage.getItem('authUser'));
}  

const methods = { 
    login,
    register,
    profile,
    logout,
    getAuthUser,
    createPoll,
    getPolls,
    updateVote
}

export default methods;
