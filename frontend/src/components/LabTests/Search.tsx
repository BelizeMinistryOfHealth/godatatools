import React from 'react';
import { Box, Button, Form, FormField, Heading, TextInput } from 'grommet';
import { FormExtendedEvent } from 'grommet/components/Form';

export interface LabTestSearchFormState {
  firstName?: string;
  lastName?: string;
}

export interface LabTestsSearchProps<T> {
  onSubmit?: (event: FormExtendedEvent<T>) => void;
}

const LabTestsSearch = (props: LabTestsSearchProps<LabTestSearchFormState>) => {
  const { onSubmit } = props;
  const [value, setValue] = React.useState<LabTestSearchFormState>({});

  return (
    <Box
      direction={'column'}
      gap={'medium'}
      pad={'small'}
      round={'large'}
      background={'dark-3'}
    >
      <Heading level={2} color={'white'} margin={{ left: 'large' }}>
        Search For Lab Tests
      </Heading>

      <Form
        value={value}
        onChange={(nextValue: LabTestSearchFormState, { touched }: any) => {
          setValue(nextValue);
        }}
        onSubmit={(event) => {
          if (onSubmit) {
            onSubmit(event);
          }
        }}
      >
        <Box
          direction={'row'}
          gap={'large'}
          margin={{ left: 'large', right: 'large' }}
        >
          <FormField
            name={'firstName'}
            label={'First Name'}
            required={true}
            color={'white'}
          >
            <TextInput color={'white'} name={'firstName'} />
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
              type={'submit'}
            />
          </Box>
        </Box>
      </Form>
    </Box>
  );
};

export default LabTestsSearch;
