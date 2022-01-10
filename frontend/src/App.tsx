import React from 'react';
import Login from './components/Login/Login';
import HttpApiProvider from './providers/HttpProvider';
import { dark, Grommet, Main } from 'grommet';
import { BrowserRouter } from 'react-router-dom';
import Routes from './components/Routes';

const REACT_APP_API_URL = process.env['REACT_APP_API_URL']
  ? process.env['REACT_APP_API_URL']
  : '';

interface Token {
  token: string;
  loading: boolean;
  error?: Error;
}

function App() {
  const [token, setToken] = React.useState<Token>({
    loading: false,
    token: '',
  });

  // Display Login form if token is empty
  if (token.token.trim() === '') {
    return (
      <HttpApiProvider baseUrl={REACT_APP_API_URL}>
        <Login onLogin={(token) => setToken({ token, loading: false })} />;
      </HttpApiProvider>
    );
  }

  return (
    <Grommet theme={dark} full>
      <BrowserRouter>
        <Main direction={'column'} flex={false} responsive={true}>
          <HttpApiProvider baseUrl={REACT_APP_API_URL}>
            <Routes />
          </HttpApiProvider>
        </Main>
      </BrowserRouter>
    </Grommet>
  );
}

export default App;
