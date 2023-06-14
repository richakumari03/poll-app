import React from 'react';
import ReactDOM from 'react-dom';

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min';

import App from './App';
import { BrowserRouter } from 'react-router-dom';


ReactDOM.render(
        <BrowserRouter><App /></BrowserRouter>, document.getElementById('root'));

