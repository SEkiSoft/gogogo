// Copyright (c) 2016 David Lu
// See License.txt

import React from 'react';
import ReactDOM from 'react-dom';
import {Router, browserHistory} from 'react-router/es6';

import rRoot from 'routes/route_root.jsx';

function renderRootComponent() {
    ReactDOM.render((
        <Router
            history={browserHistory}
            routes={rRoot}
        />
    ),
    document.getElementById('root'));
}

global.window.setup_root = () => {
    renderRootComponent();
};
