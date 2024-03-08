<template>
  <div>
    <BasicTable @register="registerTable">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <a-button @click="handleEdit(record)" type="primary">{{
            t('paramConfig.param.paramConfig')
          }}</a-button>
        </template>
        <template v-if="column.key === 'minBox'">
          {{ '宽度: ' + record.minBox.width + 'px ' + '高度: ' + record.minBox.height + 'px ' }}
        </template>
        <template v-if="column.key === 'ability'">
          {{ option[record.ability] }}
        </template>
      </template>
    </BasicTable>
    <ParamModal @register="registerModal" @success="handleSuccess" />
  </div>
</template>

<script lang="ts" setup>
  import { BasicTable, useTable } from '/@/components/Table';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { getalgoConfig } from '/@/api/paramConfig/index';
  import { getBasicColumns } from './Data';
  import { useModal } from '/@/components/Modal';
  import { option } from '/@/components/Data/algoData';
  import ParamModal from './ParamModal.vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  const { createMessage } = useMessage();

  const { t } = useI18n();
  const [registerTable, { reload }] = useTable({
    api: getalgoConfig,
    rowKey: 'ability',
    columns: getBasicColumns(),
    showTableSetting: true,
    bordered: true,
    tableSetting: { fullScreen: true },
    showIndexColumn: true,
    indexColumnProps: {
      width: 160,
    },
    actionColumn: {
      width: 120,
      title: t('component.cropper.btn_action'),
      dataIndex: 'action',
    },
  });
  const [registerModal, { openModal }] = useModal();
  function handleEdit(record: Recordable) {
    openModal(true, {
      record,
    });
  }
  function handleSuccess() {
    createMessage.success('操作成功');
    reload();
  }
</script>
