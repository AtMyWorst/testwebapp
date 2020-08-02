
const express = require('express');
const bodyParser = require('body-parser')
const tf = require('@tensorflow/tfjs');

const router = express.Router();
const jsonParser = bodyParser.json();
const textParser = bodyParser.text();

const ALPHA = 0.9;

let paramsMap = new Map();
let lossMap = new Map();

router.post('/update/:token', jsonParser, async function(req, res) {
  let shape = req.body.shape;
  let rawWeights = req.body.data;
  const token = req.params.token;

  let weights = tf.tensor(rawWeights, shape);

  if(paramsMap.has(token)) {
    let oldWeights = paramsMap.get(token);
    let newWeights = weights.mul(tf.scalar(ALPHA)).add(oldWeights.mul(tf.scalar(1 - ALPHA)));

    res.send({
      shape: newWeights.shape,
      data: await newWeights.flatten().array()
    });
  } else {
    paramsMap.set(token, weights);

    res.send({
      shape: weights.shape,
      data: await weights.flatten().array()
    });
  }
});
