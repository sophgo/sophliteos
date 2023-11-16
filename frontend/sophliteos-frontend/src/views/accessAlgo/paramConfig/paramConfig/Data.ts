import { BasicColumn, FormSchema } from '/@/components/Table';
import { useI18n } from '/@/hooks/web/useI18n';
const { t } = useI18n();
export function getBasicColumns(): BasicColumn[] {
  return [
    {
      title: t('paramConfig.param.algoType'),
      dataIndex: 'ability',
      width: 150,
      align: 'left',
    },
    {
      title: t('paramConfig.param.minDetectRange'),
      dataIndex: 'minBox',
      width: 150,
      align: 'left',
    },
    {
      title: t('paramConfig.param.alarmInterval'),
      dataIndex: 'interval',
      width: 150,
      align: 'left',
    },
    {
      title: t('paramConfig.param.algoThreshold'),
      dataIndex: 'threshold',
      width: 150,
      align: 'left',
    },
  ];
}

export const paramFormSchema: FormSchema[] = [
  {
    field: 'ability',
    label: t('paramConfig.param.algoType'),
    component: 'Input',
    componentProps: {
      disabled: true,
    },
  },
  // {
  //   field: 'minBox',
  //   label: t('paramConfig.param.minDetectRange'),
  //   component: 'Input',
  //   slot: 'minBox',
  // },
  {
    field: 'width',
    label: t('paramConfig.param.minDetectRange'),
    component: 'Input',
    slot: 'width',
  },
  {
    field: 'height',
    label: ' ',
    component: 'Input',
    slot: 'height',
  },
  {
    field: 'interval',
    label: t('paramConfig.param.alarmInterval'),
    component: 'Input',
    slot: 'interval',
  },
  {
    field: 'threshold',
    label: t('paramConfig.param.algoThreshold'),
    component: 'Input',
    slot: 'threshold',
  },
];
