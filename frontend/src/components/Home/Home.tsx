import React, { useContext } from 'react';
import { useHttpApi } from '../../providers/HttpProvider';
import Spinner from '../Spinner/Spinner';
import {
  Box,
  Card,
  dark,
  Grid,
  Grommet,
  Heading,
  Main,
  ResponsiveContext,
  Text,
} from 'grommet';
import { CloudDownload } from 'grommet-icons';

export interface Outbreak {
  _id: string;
  name: string;
}

interface OutbreakData {
  status: 'loading' | 'start' | 'done' | 'error';
  outbreaks: Outbreak[];
  error?: Error;
}

const OutbreakGrid = (props: { outbreaks: Outbreak[] }) => {
  const size = useContext(ResponsiveContext);
  const { httpInstance } = useHttpApi();
  return (
    <Box
      pad={'small'}
      background={{ color: 'light-2' }}
      gap={'medium'}
      flex={false}
      direction={'column'}
      justify={'center'}
      margin={{ top: 'small', left: 'small', right: 'small' }}
    >
      <Grid columns={size !== 'small' ? 'small' : '100%'} gap={'small'}>
        {props.outbreaks.map((outbreak, index) => (
          <Card pad={'large'} key={index}>
            <Text key={index}>{outbreak.name}</Text>
            <CloudDownload
              onClick={() => {
                window.open(
                  `${httpInstance.getBaseUrl()}/casesByOutbreak?outbreakId=${
                    outbreak._id
                  }`,
                  '_blank'
                );
              }}
            />
          </Card>
        ))}
      </Grid>
    </Box>
  );
};

const Home = () => {
  const [outbreakData, setOutbreakData] = React.useState<OutbreakData>({
    status: 'start',
    outbreaks: [],
  });
  const { httpInstance } = useHttpApi();

  React.useEffect(() => {
    const fetchOutbreaks = () => {
      httpInstance
        .get('/outbreaks')
        .then((resp) => {
          const outbreaks = resp.data.map((o: Outbreak) => {
            const name = o.name.split('_');
            return {
              _id: o._id,
              name: name[name.length - 1],
            };
          });
          setOutbreakData({
            outbreaks: outbreaks,
            status: 'done',
          });
        })
        .catch((err) => {
          console.error('error: ', err);
          setOutbreakData({ outbreaks: [], status: 'error', error: err });
        });
    };

    if (outbreakData.status === 'loading') {
      fetchOutbreaks();
    }
  }, [outbreakData, httpInstance]);

  if (outbreakData.status === 'start') {
    setOutbreakData({ status: 'loading', outbreaks: [] });
  }

  if (outbreakData.status === 'loading') {
    return <Spinner size={228} />;
  }

  return (
    <Grommet theme={dark} full={true}>
      <Main justify={'evenly'} background={'dark-3'} responsive={true}>
        <Box
          background={'light-1'}
          responsive={true}
          flex={false}
          gap={'medium'}
          pad={'small'}
          margin={{ left: 'small', bottom: 'xxsmall', right: 'small' }}
        >
          <Heading level={3}>Select an Outbreak to download cases</Heading>
        </Box>
        <OutbreakGrid outbreaks={outbreakData.outbreaks} />
      </Main>
    </Grommet>
  );
};

export default Home;
