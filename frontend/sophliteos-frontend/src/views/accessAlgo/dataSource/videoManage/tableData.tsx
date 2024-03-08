import { BasicColumn, FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { ref } from 'vue';

const { t } = useI18n();

export function getBasicColumns(): BasicColumn[] {
  return [
    {
      title: t('dataSource.videoManage.deviceName'),
      dataIndex: 'name',
      width: 150,
      align: 'center',
      ellipsis: true,
    },
    {
      title: t('dataSource.videoManage.mediaServerName'),
      dataIndex: 'mediaServer',
      width: 150,
      align: 'left',
    },
    {
      title: t('dataSource.videoManage.type'),
      dataIndex: 'ptzType',
      width: 100,
      align: 'left',
      ellipsis: true,
    },
    {
      title: t('dataSource.videoManage.status'),
      dataIndex: 'status',
      width: 80,
      align: 'left',
      customRender: ({ record }) => {
        const status = record.status;
        const enable = status === 'ON';
        const color = enable ? 'green' : 'red';
        const text = enable ? '在线' : '离线';
        return h(Tag, { color: color }, () => text);
      },
    },
    {
      title: t('dataSource.videoManage.protocol'),
      dataIndex: 'protocol',
      width: 120,
      align: 'left',
      ellipsis: true,
    },
    {
      title: t('dataSource.videoManage.mediaPull'),
      dataIndex: 'mediaPull',
      width: 120,
      align: 'left',
    },
    {
      title: t('dataSource.videoManage.isNeedDetect'),
      dataIndex: 'isNeedDetect',
      width: 120,
      align: 'left',
    },
    {
      title: t('dataSource.videoManage.encodeFormat'),
      dataIndex: 'codec',
      width: 150,
      align: 'left',
    },
    {
      title: t('dataSource.videoManage.resolution'),
      dataIndex: 'resolution',
      width: 100,
      align: 'left',
    },
  ];
}
export const schemas: FormSchema[] = [
  {
    field: 'ip',
    component: 'Input',
    label: t('dataSource.mediaServers.serverIp'),

    componentProps: {
      placeholder: '请输入服务IP',
    },
    required: true,
  },
  {
    field: 'port',
    component: 'Input',

    label: t('dataSource.mediaServers.serverPort'),
    componentProps: {
      placeholder: '请输入服务端口',
    },
    required: true,
  },
];
const disable = ref();
export function isAdd(param) {
  disable.value = param;
}

export const video: FormSchema[] = [
  {
    field: 'protocol',
    component: 'Select',
    label: t('dataSource.videoManage.flowType'),
    componentProps: {
      options: [
        {
          label: 'RTSP',
          value: 2,
        },
      ],
    },
    required: true,
  },
  {
    field: 'name',
    component: 'Input',
    label: t('dataSource.videoManage.deviceName'),
    componentProps: {
      disabled: disable,
    },
    required: true,
  },
  {
    field: 'url',
    component: 'Input',
    label: t('dataSource.videoManage.RTSP'),
    required: true,
  },
  {
    field: 'ptzType',
    component: 'Select',
    componentProps: {
      options: [
        {
          label: '未知',
          value: 0,
        },
        {
          label: '球机',
          value: 1,
        },
        {
          label: '半球',
          value: 2,
        },
        {
          label: '枪机',
          value: 3,
        },
        {
          label: '遥控枪机',
          value: 4,
        },
      ],
    },
    label: t('dataSource.videoManage.type'),
  },
];
