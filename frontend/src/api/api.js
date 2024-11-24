import axios from 'axios';

const api = axios.create({
  baseURL: '/api',
});

export const getNodes = async () => {
  const response = await api.get('/nodes');
  return response.data.nodes;
};

export const addNode = async (node) => {
  const response = await api.post('/nodes', node);
  return response.data;
};

export const updateNode = async (id, node) => {
  const response = await api.put(`/nodes/${id}`, node);
  return response.data;
};

export const deleteNode = async (id) => {
  const response = await api.delete(`/nodes/${id}`);
  return response.data;
};

export const generateSubscription = async (userKey) => {
  const response = await api.get(`/subscribe/${userKey}`);
  return response.data;
};

export const importSubscription = async (url) => {
  const response = await api.post('/import', { url });
  return response.data;
};
