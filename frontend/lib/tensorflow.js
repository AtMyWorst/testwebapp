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

  statsCa