import React from 'react';
import { Box, dark, Grommet, Heading, Main } from 'grommet';
import Header from '../Header/Header';
import { PageMenus } from '../PageMenus/PageMenus';
import LabTestsSearch, { LabTestSearchFormState } from './Search';
import { FormExtendedEvent } from 'grommet/components/Form';
import { useHttpApi } from '../../providers/HttpProvider';
import { LabTest } from '../../models/labTest';
import Spinner from '../Spinner/Spinner';
import LabTests from './LabTests';

const LabTestResultSearch = () => {
  const [formState, setFormState] = React.useState<LabTestSearchFormState>();
  const [labTests, setLabTests] = React.useState<LabTest[]>();
  const [loading, setLoading] = React.useState<boolean>(false);
  const { httpInstance } = useHttpApi();
  const onSubmit = (event: FormExtendedEvent<LabTestSearchFormState>) => {
    setLoading(true);
    setFormState(event.value);
  };

  React.useEffect(() => {
    // console.log('set state: ', formState);
    const getLabTests = () => {
      console.dir({ formState });
      httpInstance
        .get(
          `${httpInstance.getBaseUrl()}/labTests/searchByName?firstName=${
            formState?.firstName
          }&lastName=${formState?.lastName}`
        )
        .then((result) => {
          setLabTests(result.data);
          setLoading(false);
        })
        .catch((error) => {
          console.error(error);
          setLoading(false);
        });
    };

    if (loading) {
      getLabTests();
    }
  }, [formState, loading, httpInstance]);

  return (
    <Grommet theme={dark} full={true}>
      <Main justify={'evenly'} background={'dark-3'} responsive={true}>
        <Box>
          <Header children={<PageMenus />} />
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
            {loading && (
              <Box>
                <Spinner size={228} />
              </Box>
            )}
            {!loading && <LabTestsSearch onSubmit={onSubmit} />}
            {!loading && labTests && <LabTests labTests={labTests} />}
            {!loading && labTests?.length == 0 && formState && (
              <Heading level={4}>No Results Found.</Heading>
            )}
          </Box>
        </Box>
      </Main>
    </Grommet>
  );
};

export default LabTestResultSearch;
