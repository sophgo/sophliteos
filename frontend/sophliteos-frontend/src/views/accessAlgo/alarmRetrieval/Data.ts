import { FormSchema } from '/@/components/Form';
import { useI18n } from '/@/hooks/web/useI18n';
import { getVideosList } from '/@/api/dataSource/index';

const { t } = useI18n();
const colProps = {
  span: 6,
};

export const schemas: FormSchema[] = [
  {
    field: 'deviceName',
    label: t('alarmRetrieval.alarm.channel'),
    component: 'ApiSelect',
    componentProps: {
      api: getVideosList,
      labelField: 'name',
      valueField: 'name',
    },
    colProps,
    defaultValue: '',
  },
  {
    field: 'alarms',
    label: t('alarmRetrieval.alarm.search'),
    component: 'Select',
    slot: 'alarms',
    defaultValue: [],
    colProps,
  },
  {
    field: 'beginTime',
    component: 'DatePicker',
    componentProps: {
      valueFormat: 'YYYY-MM-DD',
    },
    label: t('alarmRetrieval.alarm.startTime'),
    colProps,
  },
  {
    field: 'endTime',
    component: 'DatePicker',
    label: t('alarmRetrieval.alarm.EndTime'),
    componentProps: {
      valueFormat: 'YYYY-MM-DD',
    },
    colProps,
  },
];
