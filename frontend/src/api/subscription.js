import axios from 'axios';

const api = axios.create({
  baseURL: '/api',
});

export const importSubscription = async (url) => {
  const response = await api.post('/import', { url });
  return response.data;
};
