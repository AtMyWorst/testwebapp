const express = require('express');
const tf = require('@tensorflow/tfjs');

const paramsRouter = require('./routes/params');
const adminRouter = require('./routes/params');

// Configuring port
const port = process.env.PORT || '10201';

//  Creating the app
const app = express();

// Regis