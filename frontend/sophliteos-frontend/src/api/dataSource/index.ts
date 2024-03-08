import { defHttp } from '/@/utils/http/axios';

import { BasicApiResponse, VideoApiResponse } from '../model/baseModel';

enum Api {
  getMediaServer = '/media/get',
  addMediaServer = '/media/add',
  livePreview = '/media/live?',
  videosList = '/media/list', //视频资源获取
  delDevice = '/media/dev/del',
  modDevice = '/media/dev/mod',
  addDevice = '/media/dev/add',
  deviceCheck = '/media/check',
}

export function getMediaServer() {
  return defHttp.get({ url: Api.getMediaServer }, { apiUrl: 'algorithm' });
}
export function PostAddMediaServer(params: any) {
  return defHttp.post<BasicApiResponse>(
    { url: Api.addMediaServer, params },
    { apiUrl: 'algorithm' },
  );
}
export function LivePreview(params: any) {
  return defHttp.get<BasicApiResponse>(
    { url: Api.livePreview, params },
    { apiUrl: 'algorithm', joinTime: false },
  );
}
export function getVideosList() {
  const res = defHttp
    .post<VideoApiResponse>({ url: Api.videosList }, { apiUrl: 'algorithm' })
    .then((res) => {
      return res.device;
    });
  return res;
}
export function DelDevice(params: any) {
  return defHttp.post({ url: Api.delDevice, params }, { apiUrl: 'algorithm' });
}
export function ModDevice(params: any) {
  return defHttp.post({ url: Api.modDevice, params }, { apiUrl: 'algorithm' });
}

export function AddDevice(params: any) {
  return defHttp.post({ url: Api.addDevice, params }, { apiUrl: 'algorithm' });
}

export function DeviceCheck(params: any) {
  return defHttp.post({ url: Api.deviceCheck, params }, { apiUrl: 'algorithm' });
}
