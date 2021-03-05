import React from 'react';
import { Box, dark, Grommet, Main } from 'grommet';
import Header from '../Header/Header';
import LabTestsSearch from '../LabTests/Search';
import { PageMenus } from '../PageMenus/PageMenus';

const Home = () => {
  return (
    <Grommet theme={dark} full={true}>
      <Main justify={'evenly'} background={'dark-3'} responsive={true}>
        <Box>
          <Header children={<PageMenus />} />
          <Box
            background={'light-1'}
            height={'large'}
            align={'center'}
            justify={'center'}
          >
            <LabTestsSearch />
          </Box>
        </Box>
      </Main>
    </Grommet>
  );
};

export default Home;
