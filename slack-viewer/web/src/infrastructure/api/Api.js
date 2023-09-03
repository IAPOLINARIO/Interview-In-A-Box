import axios from 'axios'

const instance = axios.create({
    baseURL:  process.env.REACT_APP_API_BASE_URL
});

instance.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem("okta-token-storage");
      if (token) {
        config.headers["Authorization"] = "Bearer " + JSON.parse(token).accessToken.accessToken;
      }
      return config;
    },
    (error) => {
      return Promise.reject(error);
    }
  );

export class API {
    static async delete({ path, config }) {
        const response = await instance.delete(path, config);
        return response.data;
    }

    static async get({path, config}) {
        const response = await instance.get(path, config);
        return response.data;
    }

    static async patch({path, body, config}) {
        const response = await instance.patch(path, body, config);
        return response.data;
    }

    static async post({path, body, config}) {
        const response = await instance.post(path, body, config);
        return response.data;
    }

    static async put({path, body, config}) {
        const response = await instance.put(path,body, config);
        return response.data;
    }
}
