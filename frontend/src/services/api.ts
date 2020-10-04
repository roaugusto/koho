import axios from 'axios';

export const baseURL = process.env.REACT_APP_BASE_URL;
console.log('baseURL', baseURL);

const api = axios.create({
  baseURL: baseURL,
});

export default api;
