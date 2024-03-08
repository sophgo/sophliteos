import { defHttp } from '/@/utils/http/axios';
import { MediaServerParams } from './model/index';

enum Api {
  algoConfigMod = '/config/mod',
  getalgoConfig = '/config/get',
}

export function getalgoConfig(params: MediaServerParams) {
  return defHttp.get({ url: Api.getalgoConfig, params }, { apiUrl: 'algorithm' });
}

export function algoConfigMod(params: any) {
  return defHttp.post({ url: Api.algoConfigMod, params }, { apiUrl: 'algorithm' });
}
