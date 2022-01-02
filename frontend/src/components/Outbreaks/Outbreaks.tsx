import React, { useContext } from 'react';
import {
  Box,
  Button,
  Card,
  dark,
  DateInput,
  Form,
  FormField,
  Grid,
  Grommet,
  Heading,
  Layer,
  Main,
  ResponsiveContext,
  Text,
} from 'grommet';
import { useHttpApi } from '../../providers/HttpProvider';
import { Close, CloudDownload } from 'grommet-icons';
import Header from '../Header/Header';
import Spinner from '../Spinner/Spinner';
import { PageMenus } from '../PageMenus/PageMenus';
import { differenceInDays, format, isAfter, parseISO } from 'date-fns';

export interface Outbreak {
  _id: string;
  name: string;
}

export interface OutbreakData {
  status: 'loading' | 'start' | 'done' | 'error';
  outbreaks: Outbreak[];
  error?: Error;
}

const DATE_RANGE = ['2020-07-02', '2020-07-05'];

export const OutbreakGrid = (props: { outbreaks: Outbreak[] }) => {
  const size = useContext(ResponsiveContext);
  const [errors, setErrors] = React.useState<string | null>();
  const [dateRangeNoTZ, setDateRangeNoTZ] = React.useState<string[]>([]);
  const [
    selectedOutbreak,
    setSelectedOutbreak,
  ] = React.useState<Outbreak | null>(null);
  const { httpInstance } = useHttpApi();

  const onClose = () => {
    setSelectedOutbreak(null);
  };

  const onSubmit = () => {
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
    if (dayRange > 31) {
      setErrors('Date Range can only be a maximum of 31 days!');
      return;
    }
    const url = `${httpInstance.getBaseUrl()}/casesByOutbreak?outbreakId=${
      selectedOutbreak?._id
    }?startDate=${format(startDate, 'yyyy-MM-dd')}&endDate=${format(
      endDate,
      'yyyy-MM-dd'
    )}`;
    console.log('url: ', url);
    window.open(
      `${httpInstance.getBaseUrl()}/casesByOutbreak?outbreakId=${
        selectedOutbreak?._id
      }&startDate=${format(startDate, 'yyyy-MM-dd')}&endDate=${format(
        endDate,
        'yyyy-MM-dd'
      )}`,
      '_blank'
    );
    setDateRangeNoTZ([]);
    setSelectedOutbreak(null);
  };

  const onChangeRangeNoTZ = (event: { value: any }) => {
    const nextValue = event.value;
    setDateRangeNoTZ(nextValue);
  };

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
      {selectedOutbreak && (
        <Layer
          position={'center'}
          animation={'fadeIn'}
          onClickOutside={onClose}
          onEsc={onClose}
        >
          <Box overflow={'auto'} align={'center'}>
            <Box
              flex={true}
              direction={'row-responsive'}
              justify={'evenly'}
              pad={{ top: 'medium' }}
              align={'center'}
            >
              <Heading level={5} margin={'none'}>
                Select Dates for {selectedOutbreak?.name}
              </Heading>
              <Button
                icon={<Close color={'white'} size={'small'} />}
                onClick={onClose}
              />
            </Box>
            <Box flex={'grow'} overflow={'auto'} pad={{ vertical: 'medium' }}>
              <Form onSubmit={onSubmit}>
                {errors && (
                  <Box pad={{ left: 'small' }}>
                    <Text color={'accent-4'} size={'small'}>
                      {errors}
                    </Text>
                  </Box>
                )}
                <FormField label={'Date Ranges'}>
                  <DateInput
                    name={'dateRangeNoTZ'}
                    value={dateRangeNoTZ}
                    defaultValue={DATE_RANGE}
                    format={'yyyy/mm/dd-yyyy/mm/dd'}
                    onChange={onChangeRangeNoTZ}
                    inline
                  />
                </FormField>
                <Box justify={'start'} pad={'medium'}>
                  <Button label={'Submit'} type={'submit'} primary={true} />
                </Box>
              </Form>
            </Box>
          </Box>
        </Layer>
      )}
      <Grid columns={size !== 'small' ? 'small' : '100%'} gap={'small'}>
        {props.outbreaks.map((outbreak, index) => (
          <Card
            pad={'large'}
            key={index}
            onClick={() => setSelectedOutbreak(outbreak)}
          >
            <Text key={index}>{outbreak.name}</Text>
            <CloudDownload
              onClick={() => {
                setSelectedOutbreak(outbreak);
              }}
            />
          </Card>
        ))}
      </Grid>
    </Box>
  );
};

export const OutbreakPage = () => {
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
        <Box>
          <Header children={<PageMenus />} />
          <Box
            background={'light-1'}
            height={'large'}
            gap={'small'}
            pad={'small'}
            margin={{
              left: 'small',
              bottom: 'xxsmall',
              right: 'small',
              top: 'small',
            }}
          >
            <OutbreakGrid outbreaks={outbreakData.outbreaks} />
          </Box>
        </Box>
      </Main>
    </Grommet>
  );
};
