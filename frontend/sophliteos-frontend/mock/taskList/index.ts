import { MockMethod } from 'vite-plugin-mock';
import { resultSuccess } from '../_util';
const res = [
  {
    id: 'video1',
  },
  {
    id: 'video2',
  },
  {
    id: 'video3',
  },
  {
    id: 'video4',
  },
];
const list = {
  total: 2,
  pageSize: 20,
  pageCount: 1,
  pageNo: 1,
  items: [
    {
      taskName: 'smoking and helmet',
      deviceName: 'test', //视频通道名称
      status: 1, //任务状态：1在线，0离线
      errorReason: '', //任务状态离线原因
      abilities: ['Smoking', 'WithoutHelmetOnSite'], //任务对应的算法能力
    },
    {
      taskName: '10001',
      deviceName: 'test2', //视频通道名称
      status: 0, //任务状态：1在线，0离线
      errorReason: 'Parameter error', //任务状态离线原因
      abilities: ['Smoking'], //任务对应的算法能力
    },
  ],
};
const msg = {
  code: 0, //0成功，其他失败
  msg: 'ok', //消息描述
};
export default [
  {
    url: '/api/taskList/videoSource',
    statusCode: 200,
    method: 'get',
    response: () => {
      return resultSuccess(res);
    },
  },
  {
    url: '/api/algorithm/task/list',
    statusCode: 200,
    method: 'get',
    response: () => {
      return resultSuccess(list);
    },
  },
  {
    url: '/api/algorithm/task/delete',
    statusCode: 200,
    method: 'post',
    response: () => {
      return resultSuccess(msg);
    },
  },
  {
    url: '/api/algorithm/task/start',
    statusCode: 200,
    method: 'post',
    response: () => {
      return resultSuccess(msg);
    },
  },
  {
    url: '/api/algorithm/task/stop',
    statusCode: 200,
    method: 'post',
    response: () => {
      return resultSuccess(msg);
    },
  },
] as MockMethod[];
