import React, { ReactNode, useRef } from "react";
import HttpApi from "../api/http";

interface HttpApiProviderContext {
  httpInstance: HttpApi;
  baseUrl: string;
}

export const HttpApiContext = React.createContext<HttpApiProviderContext>(
  ({} as unknown) as HttpApiProviderContext
);

export const useHttpApi = () => React.useContext(HttpApiContext);

export interface HttpApiProviderArg {
  baseUrl: string;
  children: ReactNode;
}

const HttpApiProvider = (arg: HttpApiProviderArg) => {
  const { baseUrl, children } = arg;
  const httpInstance = useRef<HttpApi>();

  httpInstance.current = new HttpApi(baseUrl);

  const contextValue = { httpInstance: httpInstance.current, baseUrl };
  if (httpInstance.current) {
    return (
      <HttpApiContext.Provider value={contextValue}>
        {children}
      </HttpApiContext.Provider>
    );
  } else {
    return <>Loading...</>;
  }
};

export default HttpApiProvider;
