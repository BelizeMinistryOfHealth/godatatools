import Axios from 'axios';
import { useState } from 'react';

class HttpApi {
  protected baseUrl: string;
  axiosInstance;
  headers: {
    Accept: string;
    'Content-Type': string;
    responseType?: string;
  };
  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
    this.headers = {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    };
    this.axiosInstance = Axios.create({
      baseURL: this.baseUrl,
      headers: this.headers,
    });
  }

  async get(path: string) {
    return this.axiosInstance.get(path);
  }

  async stream(path: string) {
    return this.axiosInstance.get(path, {
      headers: {
        ...this.headers,
        Accept: 'application/json',
        'Content-Type': 'application/pdf',
        responseType: 'blob',
      },
    });
  }

  async put(path: string, body: any) {
    return this.axiosInstance.put(path, JSON.stringify(body));
  }

  async post(path: string, body: any) {
    return this.axiosInstance.post(path, JSON.stringify(body));
  }

  async delete(path: string) {
    return this.axiosInstance.delete(path);
  }

  public getBaseUrl(): string {
    return this.baseUrl;
  }
}

export const useHttpApi = (baseUrl: string) => {
  const [instance, setInstance] = useState<HttpApi>();
  const i = new HttpApi(baseUrl);
  setInstance(i);
  return instance;
};

export default HttpApi;
