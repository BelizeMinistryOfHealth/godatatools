import React from 'react';
import { Box, Button, Form, FormField, TextInput } from 'grommet';

const LabTestsSearch = () => {
  return (
    <Box direction={'row'}>
      <Form>
        <FormField name={'firstName'} label={'First Name'} required={true}>
          <TextInput name={'firstName'} />
        </FormField>
        <FormField name={'lastName'} label={'Last Name'} required={true}>
          <TextInput name={'lastName'} />
        </FormField>
        <Button label={'Search'} />
      </Form>
    </Box>
  );
};

export default LabTestsSearch;
