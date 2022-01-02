import React from 'react';
import {
  Box,
  Button,
  dark,
  DateInput,
  Form,
  FormField,
  Grommet,
  Heading,
  Main,
} from 'grommet';
import { PageMenus } from '../PageMenus/PageMenus';
import Header from '../Header/Header';
import { FormExtendedEvent } from 'grommet/components/Form';
import { differenceInDays, format, isAfter, parseISO } from 'date-fns';
import { useHttpApi } from '../../providers/HttpProvider';

interface LabTestFormState {
  startDate?: string;
  endDate?: string;
}

const DATE_RANGE = ['2020-07-02', '2020-07-05'];

const LabExports = (): JSX.Element => {
  const [errors, setErrors] = React.useState<string | null>();
  const [dateRangeNoTZ, setDateRangeNoTZ] = React.useState<string[]>([]);
  const { httpInstance } = useHttpApi();

  const onChangeRangeNoTZ = (event: { value: any }) => {
    const nextValue = event.value;
    console.log('onChange', nextValue);
    setDateRangeNoTZ(nextValue);
  };

  const onSubmit = (event: any) => {
    setErrors(null);

    if (dateRangeNoTZ.length === 0) {
      setErrors('Please select a date or range of dates');
      return;
    }
    const startDate = parseISO(dateRangeNoTZ[0]);
    const endDate = parseISO(dateRangeNoTZ[1]);

    const today = new Date();
    if (isAfter(startDate, today) || isAfter(endDate, today)) {
      setErrors('Dates can not be in the future');
      return;
    }
    const dayRange = differenceInDays(endDate, startDate);
    if (dayRange > 7) {
      setErrors('Date Range can only be a maximum of 7 days!');
      return;
    }

    window.open(
      `${httpInstance.getBaseUrl()}/labtestreport/csv?startDate=${format(
        startDate,
        'yyyy-MM-dd'
      )}&endDate=${format(endDate, 'yyyy-MM-dd')}`,
      '_blank'
    );
  };

  return (
    <Grommet theme={dark} full>
      <Main justify={'evenly'} background={'dark-3'} responsive>
        <Box>
          <Header children={<PageMenus />} />
        </Box>
        <Box
          background={'light-1'}
          height={'large'}
          align={'center'}
          justify={'start'}
          gap={'large'}
          pad={'small'}
          margin={{
            left: 'small',
            bottom: 'xxsmall',
            right: 'small',
            top: 'small',
          }}
        >
          <Box
            direction={'column'}
            flex
            gap={'medium'}
            pad={'small'}
            round={'large'}
            background={'dark-3'}
          >
            <Heading level={2} color={'white'} margin={{ left: 'large' }}>
              Export Lab Tests
            </Heading>

            <Form onSubmit={onSubmit}>
              <Box gap={'large'} margin={{ left: 'large', right: 'large' }}>
                {errors && errors.length > 0 && <Box>{errors}</Box>}

                <FormField
                  name={'dateRangeNoTZ'}
                  label={'Select Dates'}
                  color={'white'}
                >
                  <DateInput
                    name={'dateRangeNoTZ'}
                    value={dateRangeNoTZ}
                    defaultValue={DATE_RANGE}
                    format={'yyyy/mm/dd-yyyy/mm/dd'}
                    onChange={onChangeRangeNoTZ}
                    inline
                  />
                </FormField>
              </Box>
              <Box justify={'start'}>
                <Button
                  label={'Submit'}
                  margin={{ left: 'large' }}
                  gap={'large'}
                  type={'submit'}
                  primary={true}
                />
              </Box>
            </Form>
          </Box>
        </Box>
      </Main>
    </Grommet>
  );
};

LabExports.displayName = 'LabExports';
export default LabExports;
