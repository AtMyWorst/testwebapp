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
	  this.bas