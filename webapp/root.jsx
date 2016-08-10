// Copyright (c) 2016 David Lu
// See License.txt

import React from 'react';
import ReactDOM from 'react-dom';
import {Router, browserHistory} from 'react-router/es6';

function renderRootComponent () {
  ReactDOM.render((
    <Router
      history={browserHistory}
      routes={null}
    />
  ),
  document.getElementById('root'));
}

global.window.setup_root = () => {
  renderRootComponent();
};
