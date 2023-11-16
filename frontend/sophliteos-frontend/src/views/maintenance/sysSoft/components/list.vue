<template>
  <BasicTable @register="registerTable" :title="t('maintenance.systemUpdate.updateList')">
    <template #bodyCell="{ column, record }">
      <template v-if="column.key === 'status'">
        {{ t(`maintenance.systemUpdate.status.${record.status}`) }}
      </template>
      <template v-if="column.key === 'type'">
        {{ t(`maintenance.systemUpdate.type.${record.type}`) }}
      </template>
      <template v-if="column.key === 'strategy'">
        {{ t(`maintenance.systemUpdate.strategy.${record.strategy}`) }}
      </template>
    </template>
  </BasicTable>
</template>
<script lang="ts" setup>
  import { BasicTable, useTable } from '/@/components/Table';
  import { getBasicColumns } from './tableData';
  import { upgradeStatusApi } from '/@/api/maintenance/index';
  import { useI18n } from '/@/hooks/web/useI18n';
  const { t } = useI18n();

  const [registerTable] = useTable({
    title: t('maintenance.systemUpdate.updateStatus'),
    api: upgradeStatusApi,
    columns: getBasicColumns(),
    showTableSetting: true,
    tableSetting: { fullScreen: true },
    showIndexColumn: true,
    indexColumnProps: {
      width: 60,
    },
    rowKey: 'name',
  });
</script>
