import React from 'react';
import {
  Box,
  Button,
  Card,
  CardBody,
  dark,
  Form,
  FormField,
  Grommet,
  Heading,
  Main,
  TextInput,
} from 'grommet';
import { useHttpApi } from '../../providers/HttpProvider';

interface Creds {
  username: string;
  password: string;
}

export interface LoginProps {
  onLogin: (token: string) => void;
}

const Login = (props: LoginProps) => {
  // Form status: START -> SUBMIT -> ERROR -> SUCCESS
  const [status, setStatus] = React.useState('START');
  const [creds, setCreds] = React.useState<Creds>();
  const [token, setToken] = React.useState<string>('');
  const { httpInstance } = useHttpApi();

  const onSubmit = (e: { value: any }) => {
    setCreds(e.value);
    setStatus('SUBMIT');
  };

  React.useEffect(() => {
    const login = () => {
      httpInstance
        .post('/auth', creds)
        .then((resp) => {
          console.log('response: ', resp.data);
          setToken(resp.data.token);
          setStatus('SUCCESS');
        })
        .catch((error) => {
          console.dir({ error });
          setStatus('ERROR');
        });
    };
    if (status === 'SUBMIT') {
      login();
    }
  }, [status, creds, httpInstance]);

  React.useEffect(() => {
    if (status === 'SUCCESS') {
      props.onLogin(token);
    }
  }, [status, token]);

  return (
    <Grommet theme={dark} full>
      <Main direction={'column'} background={'dark-1'}>
        <Box
          flex={false}
          direction={'column'}
          justify={'center'}
          align={'center'}
          gap={'medium'}
          pad={'medium'}
        >
          <Heading>EPI's Godata Tools</Heading>
          <Card
            background={'white'}
            pad={'large'}
            align={'center'}
            margin={{ top: 'xlarge' }}
          >
            <CardBody justify={'evenly'} responsive={true}>
              <Form onSubmit={onSubmit}>
                <FormField
                  name={'username'}
                  label={'Username'}
                  required={true}
                  width={'medium'}
                >
                  <TextInput name={'username'} />
                </FormField>
                <FormField name={'password'} label={'Password'} required={true}>
                  <TextInput name={'password'} type={'password'} />
                </FormField>
                <Box
                  direction={'row'}
                  justify={'center'}
                  margin={{ top: 'medium' }}
                >
                  <Button label={'Login'} type={'submit'} primary />
                </Box>
              </Form>
            </CardBody>
          </Card>
        </Box>
      </Main>
    </Grommet>
  );
};

export default Login;
