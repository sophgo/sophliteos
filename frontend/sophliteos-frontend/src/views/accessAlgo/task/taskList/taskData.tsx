import { BasicColumn, FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
// import { options } from '/@/components/Data/algoData';
const { t } = useI18n();

export const columns: BasicColumn[] = [
  {
    title: t('taskList.taskList.id'),
    dataIndex: 'taskName',
    width: 120,
  },
  {
    title: t('taskList.taskList.videoSource'),
    dataIndex: 'deviceName',
    width: 120,
  },
  {
    title: t('taskList.taskList.algoInfo'),
    dataIndex: 'abilities',
    width: 200,
  },
  {
    title: t('taskList.taskList.status'),
    dataIndex: 'status',
    width: 100,
    customRender: ({ record }) => {
      const status = record.status;
      const enable = ~~status === 1;
      const color = enable ? 'green' : 'red';
      const text = enable ? '运行中' : '已停止';
      return h(Tag, { color: color }, () => text);
    },
  },
];

export const taskFormSchema: FormSchema[] = [
  {
    field: 'taskName',
    label: t('taskList.taskList.id'),
    component: 'Input',
    componentProps: {
      disabled: true,
    },
    required: true,
  },
  {
    field: 'urlList',
    label: t('taskList.taskList.videoSource'),
    component: 'Select',
    slot: 'urlList',
    required: true,
  },
  {
    field: 'abilities',
    label: '算法设置',
    component: 'CheckboxGroup',
    helpMessage: ['算法设置'],
    // required: true,
    // componentProps: {
    //   options: options,
    // },
    slot: 'abilitie',
    colProps: {
      span: 24,
    },
  },
];

export const addTaskSchema: FormSchema[] = [
  {
    field: 'taskName',
    label: t('taskList.taskList.id'),
    component: 'Input',
    required: true,
  },
  {
    field: 'urlList',
    label: t('taskList.taskList.videoSource'),
    component: 'Select',
    slot: 'urlList',
    required: true,
  },
  {
    field: 'abilities',
    label: '算法设置',
    component: 'CheckboxGroup',
    helpMessage: ['算法设置'],
    // required: true,
    // componentProps: {
    //   options: options,
    // },
    slot: 'abilitie',
    colProps: {
      span: 24,
    },
  },
];

export const algoConfigFormSchema: FormSchema[] = [
  {
    label: t('dataSource.mediaServers.serverIp'),
    field: 'ip',
    component: 'Input',
    // colProps: { span: 8 },
    required: true,
  },
  {
    label: t('dataSource.mediaServers.serverPort'),
    field: 'port',
    component: 'Input',
    required: true,
  },
];
