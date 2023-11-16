import { defineStore } from 'pinia';

export const alarmInfo = defineStore('info', {
  state: () => ({
    iamgeInfo: {
      alarmType: '',
      time: '',
      itemsInBox: [{ confidence: '' }],
      deviceName: '',
      boxes: [
        {
          width: 0,
          height: 0,
          x: 0,
          y: 0,
        },
      ],
    },
  }),
  actions: {
    setInfo(value) {
      this.iamgeInfo = value;
    },
  },
  getters: {
    // 获取数据的 getter
    getInfo(): any {
      return this.iamgeInfo;
    },
  },
});
