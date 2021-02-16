import React from 'react';
import { Box } from 'grommet';
import { ReactComponent as SpinnerSvg } from './spinner.svg';

export interface SpinnerProps {
  size: number;
}

const Spinner = (props: SpinnerProps) => (
  <Box align='center' justify='center' flex={'grow'}>
    <SpinnerSvg height={props.size} width={props.size} />
  </Box>
);

export default Spinner;
