import { NodeDataModel } from '../node-data-model';
import { NodeData } from '../../interfaces/node-data';

export class VppTopoBvi extends NodeDataModel implements NodeData {

  constructor(data?: NodeData) {
    super(data);
  }
}
