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
    // Load the next batch