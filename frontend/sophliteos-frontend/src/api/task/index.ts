import { defHttp } from '/@/utils/http/axios';
import { BasicApiResponse } from '../model/baseModel';
import { MediaServerParams } from './model/index';
enum Api {
  taskList = '/task/list',
  deleteTask = '/task/delete',
  startTask = '/task/start',
  addTask = '/task/add',
  modTask = '/task/modify',
  stopTask = '/task/stop',
  addAlgorithm = '/server/add',
  getAlgorithm = '/server/get',
}

export function getTaskList(params?: MediaServerParams) {
  return defHttp.post({ url: Api.taskList, params }, { apiUrl: 'algorithm' });
}
export function PostDeleteTask(params: any) {
  return defHttp.post({ url: Api.deleteTask, params }, { apiUrl: 'algorithm' });
}

export function StartTask(params: any) {
  return defHttp.post({ url: Api.startTask, params }, { apiUrl: 'algorithm' });
}

export function StopTask(params: any) {
  return defHttp.post({ url: Api.stopTask, params }, { apiUrl: 'algorithm' });
}

export function getAlgorithm() {
  return defHttp.get({ url: Api.getAlgorithm }, { apiUrl: 'algorithm' });
}
export function addAlgorithm(params: any) {
  return defHttp.post({ url: Api.addAlgorithm, params }, { apiUrl: '/algorithm' });
}
export function addTask(params: any) {
  return defHttp.post<BasicApiResponse>({ url: Api.addTask, params }, { apiUrl: '/algorithm' });
}
export function modTask(params: any) {
  return defHttp.post<BasicApiResponse>({ url: Api.modTask, params }, { apiUrl: '/algorithm' });
}
