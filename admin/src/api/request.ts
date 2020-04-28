import axios from "axios";

import authState from "../state/auth";

const createHistory = require("history").createHashHistory;
const history = createHistory();
const request = axios.create({
  baseURL: process.env.REACT_APP_Base_Url,
  headers: {
    "content-type": "application/json",
    Authorization: authState.authModel?.token
  },
  validateStatus: () => true,
});

request.interceptors.request.use(
  (config) => {
    if (authState.authModel && authState.authModel.token.length > 0) {
      config.url = encodeURI(config.url ?? "")
      config.headers.Authorization = authState.authModel?.token;
    }
    return config;
  },
  (err) => {
    return Promise.reject(err);
  }
);

request.interceptors.response.use(
  (resp) => {
    if (resp.status === 401) {
      history.replace("/login");
    }
    return resp;
  },
  (err) => {
    return Promise.reject(err);
  }
);

export default request;
