import * as tf from '@tensorflow/tfjs';

/*
* Params for logging
*  -> Epoch
*  -> Loss
*  -> Accuracy
*  -> Batch num
*  -> Runtime
*/

class DistTensorflow {
  modelId;
  model;
  batchSize;
  batchNo = 0;
  stopped = false;

  statsCallback;

  constructor(modelId, statsCallback, base1, base2) {
    this.modelId = modelId;
    this.statsCallback = statsCallback;
	this.base1 = base1;
	  this.base2 = base2;

  tf.loadLayersModel(`${this.base2}/model?id=${this.modelId}`).then(function (model) {
    this.model = model;

    // Compile the model with default optimizer and loss
    this.model.compile({
      optimizer: tf.train.adam(),
      loss: 'categoricalCrossentropy',
      metrics: ['accuracy'],
    });
  });
  }

  async loadNextBatch() {
    // Load the next batch from the backend
    let res = await fetch(`${base2}/?model=${this.modelId}`, {
      method: 'GET',
      redirect: 'follow',
    });

    const batchShape = res.data.batch;
    const labelShape = res.data.label;

    // Set batch size
    this.batchSize = batchShape[0];

    // Load the minibatch data
    res = await fetch(`${base2}/data/batch?model=${this.modelId}`, {
      method: 'GET',
      redirect: 'follow',
    });

    let batchArray = new UInt8Array(await res.arrayBuffer());

    // Load the minibatch labels
    