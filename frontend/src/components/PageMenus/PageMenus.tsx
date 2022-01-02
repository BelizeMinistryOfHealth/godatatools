import React from 'react';
import { useHistory } from 'react-router-dom';
import { Box } from 'grommet';

export const PageMenus = () => {
  const history = useHistory();

  return (
    <Box direction={'row'} gap={'small'}>
      <Box onClick={() => history.push('/export_tools')}>Export Tool</Box>
      <Box> | </Box>
      <Box onClick={() => history.push('/lab_test/results/search')}>Search</Box>
      <Box> | </Box>
      <Box onClick={() => history.push('/lab_test/exports')}>
        Lab Test Export
      </Box>
    </Box>
  );
};
