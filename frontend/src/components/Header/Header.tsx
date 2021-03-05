import React, { ReactNode } from 'react';
import { Box, Heading } from 'grommet';

interface HeaderProps {
  children?: ReactNode;
}

const Header = (props: HeaderProps) => {
  return (
    <Box
      background={'dark-1'}
      responsive={true}
      flex={false}
      gap={'medium'}
      pad={'small'}
      height={'small'}
      align={'center'}
      margin={{
        left: 'small',
        bottom: 'xxsmall',
        right: 'small',
        top: 'small',
      }}
    >
      <Heading level={1}>GoData Tools</Heading>
      {props.children}
    </Box>
  );
};

export default Header;
