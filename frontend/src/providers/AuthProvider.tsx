import React, { ReactNode } from "react";

interface AuthProviderContext {
  token: string;
}

export const AuthContext = React.createContext<AuthProviderContext>({
  token: "",
});

export interface AuthProviderArgs {
  token: string;
  children: ReactNode;
}

const AuthProvider = (args: AuthProviderArgs) => {
  const { children } = args;
  return (
    <AuthContext.Provider value={{ token: args.token }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
