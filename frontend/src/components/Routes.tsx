import React from 'react';
import Home from './Home/Home';
import { Route, Switch } from 'react-router-dom';

const Routes = () => {
  return (
    <Switch>
      <Route exact={true} path={'/outbreaks'} component={Home} />
      <Route exact={true} path={'/'} component={Home} />
    </Switch>
  );
};

export default Routes;
