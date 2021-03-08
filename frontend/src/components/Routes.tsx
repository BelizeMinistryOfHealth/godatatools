import React from 'react';
import { Route, Switch } from 'react-router-dom';
import { OutbreakPage } from './Outbreaks/Outbreaks';
import LabTestResultSearch from './LabTests/LabTestResultSearch';

const Routes = () => {
  return (
    <Switch>
      <Route exact={true} path={'/'} component={LabTestResultSearch} />
      <Route exact={true} path={'/export_tools'} component={OutbreakPage} />
      <Route exact={true} path={'/outbreaks'} component={OutbreakPage} />
      <Route
        exact={true}
        path={'/lab_test/results/search'}
        component={LabTestResultSearch}
      />
    </Switch>
  );
};

export default Routes;
