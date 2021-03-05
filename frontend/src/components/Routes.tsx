import React from 'react';
import Home from './Home/Home';
import { Route, Switch } from 'react-router-dom';
import { OutbreakPage } from './Outbreaks/Outbreaks';

const Routes = () => {
  return (
    <Switch>
      <Route exact={true} path={'/outbreaks'} component={Home} />
      <Route exact={true} path={'/'} component={Home} />
      <Route exact={true} path={'/export_tools'} component={OutbreakPage} />
    </Switch>
  );
};

export default Routes;
