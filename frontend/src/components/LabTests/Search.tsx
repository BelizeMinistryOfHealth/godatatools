import React from 'react';
import { Box, Button, Form, FormField, TextInput } from 'grommet';

const LabTestsSearch = () => {
  return (
    <Box direction={'row'} gap={'medium'} pad={'small'}>
      <Form>
        <Box direction={'row'} gap={'large'}>
          <FormField name={'firstName'} label={'First Name'} required={true}>
            <TextInput name={'firstName'} />
          </FormField>
          <FormField name={'lastName'} label={'Last Name'} required={true}>
            <TextInput name={'lastName'} />
          </FormField>
          <Box justify={'center'}>
            <Button
              label={'Search'}
              margin={{ left: 'large' }}
              gap={'large'}
              primary={true}
            />
          </Box>
        </Box>
      </Form>
    </Box>
  );
};

export default LabTestsSearch;
