import { MockMethod } from 'vite-plugin-mock';
import { resultSuccess } from '../_util';
const result = [
  {
    name: '1',
    serverIp: '2',
    serverPort: '3',
    remark: '4',
  },
  {
    name: '11',
    serverIp: '22',
    serverPort: '33',
    remark: '44',
  },
];
const result2 = [
  {
    codec: 'h264',
    deviceId: '32028101001310000004',
    name: 'test3',
    protocol: 2,
    resolution: '1280*720',
    status: 'ON',
    type: 'camera',
    url: 'rtsp://172.28.8.119:26644/live/123/test3.mp4',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
  {
    codec: 'h265',
    deviceId: '32028101001310000005',
    name: 'test1',
    protocol: 2,
    resolution: '4096*2160',
    status: 'ON',
    type: 'camera',
    url: 'rtsp://172.28.8.119:26644/live/123/test1.mp4',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
  {
    codec: 'h265',
    deviceId: '32028101001310000006',
    name: 'test2',
    protocol: 2,
    resolution: '2048*1536',
    status: 'ON',
    type: 'camera',
    url: 'rtsp://172.28.8.119:26644/live/123/test2.mp4',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
  {
    codec: 'unknown',
    deviceId: '32028101001310000007',
    name: 'test4',
    protocol: 2,
    resolution: 'unknown',
    status: 'OFF',
    type: 'camera',
    url: 'rtsp://172.28.8.119:26644/live/123/test4.mp4',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
  {
    codec: 'h265',
    deviceId: '32028101001310000008',
    name: '02010002414000000.mp4',
    protocol: 2,
    resolution: '1920*1080',
    status: 'ON',
    type: 'camera',
    url: 'rtsp://172.28.8.119:26644/live/123/02010002414000000.mp4',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
  {
    codec: 'h264',
    deviceId: '32028101001310000009',
    name: 'ca-Person',
    protocol: 2,
    resolution: '1920*1080',
    status: 'ON',
    type: 'camera',
    url: 'rtsp://172.28.8.119:26644/live/123/004_car_person.mp4',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
  {
    codec: '',
    deviceId: '32028101002150000001',
    name: '32028101002150000001',
    protocol: 0,
    resolution: '',
    status: '',
    type: 'bizGroup',
    url: '',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
  {
    codec: '',
    deviceId: '32028101002169999999',
    name: 'default_group',
    protocol: 0,
    resolution: '',
    status: '',
    type: 'bizGroup',
    url: '',
    mediaServer: '172.28.8.119:26080',
    mediaPull: 0,
    isNeedDetect: 1,
  },
];

const getMediaServer = {
  code: 0, //0成功，其他失败
  msg: 'ok', //消息描述
  result: {
    name: '服务名称',
    ip: '172.28.8.119',
    port: 26080,
  }, //返回流媒体ip和端口
};
const addMediaServer = {
  code: 0, //0成功，其他失败
  msg: 'ok', //消息描述
};
export default [
  {
    url: '/api/middleground-ops/mediaServer/list',
    statusCode: 200,
    method: 'get',
    response: () => {
      return resultSuccess(result);
    },
  },
  {
    url: '/api/algorithm/media/list',
    statusCode: 200,
    method: 'get',
    response: () => {
      return resultSuccess(result2);
    },
  },
  {
    url: '/api/algorithm/media/get',
    statusCode: 200,
    method: 'get',
    response: () => {
      return resultSuccess(getMediaServer);
    },
  },
  {
    url: '/api/algorithm/media/add',
    statusCode: 200,
    method: 'post',
    response: () => {
      return resultSuccess(addMediaServer);
    },
  },
] as MockMethod[];
