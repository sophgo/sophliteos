<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button type="primary" @click="addTask">{{ t('taskList.taskList.addTask') }}</a-button>
        <a-button type="primary" @click="algoConfig">{{
          t('taskList.taskList.algoConfigTitle')
        }}</a-button>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'mdi:success-circle-outline',
                tooltip: t('component.table.start'),
                color: 'success',
                onClick: handle.bind(null, record),
                ifShow: !Boolean(record.status),
              },
              {
                icon: 'nimbus:stop',
                tooltip: t('component.table.stop'),
                color: 'warning',
                onClick: handle.bind(null, record),
                ifShow: Boolean(record.status),
              },
              {
                icon: 'clarity:note-edit-line',
                tooltip: t('component.cropper.btn_edit'),
                onClick: handleEdit.bind(null, record),
              },
              {
                icon: 'ant-design:delete-outlined',
                color: 'error',
                tooltip: t('component.cropper.btn_delete'),
                popConfirm: {
                  title: '是否确认删除',
                  placement: 'left',
                  confirm: handleDelete.bind(null, record.taskName),
                },
              },
            ]"
          />
        </template>
        <template v-if="column.key === 'abilities'">
          <a-tooltip placement="bottom">
            <template #title>
              <div
                v-for="(item, index) in record.abilities"
                :key="index"
                style="display: inline-block"
              >
                <span>{{ option[item] }}</span>
                <span v-if="index !== record.abilities.length - 1">;&nbsp;</span>
              </div>
            </template>
            <div
              v-for="(item, index) in record.abilities"
              :key="index"
              style="display: inline-block"
            >
              <span>{{ option[item] }}</span>
              <span v-if="index !== record.abilities.length - 1">;&nbsp;</span>
            </div>
          </a-tooltip>
        </template>
      </template>
    </BasicTable>
    <TaskFormModal @register="registerModal" @success="handleSuccess" />
    <AddTaskModal @register="addModal" @addsuccess="addSuccess" />
    <AlgoConfigModal @register="algoConfigModal" @configsuccess="configSuccess" />
  </div>
</template>

<script lang="ts" setup>
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    getTaskList,
    PostDeleteTask,
    StartTask,
    StopTask,
    getAlgorithm,
  } from '/@/api/task/index';
  import { useModal } from '/@/components/Modal';
  import { columns } from './taskData';
  import TaskFormModal from './TaskFormModal.vue';
  import AddTaskModal from './AddTaskModal.vue';
  import AlgoConfigModal from './AlgoConfigModal.vue';
  // import { useRedo } from '/@/hooks/web/usePage';
  import { option } from '/@/components/Data/algoData';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { Tooltip } from 'ant-design-vue';
  const { createMessage } = useMessage();
  // const redo = useRedo();
  const { t } = useI18n();
  const ATooltip = Tooltip;
  const [registerTable, { reload, getDataSource }] = useTable({
    api: getTaskList,
    rowKey: 'taskName',
    columns,
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

  const [registerModal, { openModal: editOpenModal }] = useModal();
  function handleEdit(record: Recordable) {
    const usedUrl = getDataSource().map((item) => item.deviceName) || [];
    editOpenModal(true, {
      record,
      usedUrl,
    });
  }
  async function handleDelete(id: string) {
    await PostDeleteTask({ taskName: id });
    reload();
  }
  async function handle(record: any) {
    if (record.status === 0) {
      await StartTask({ taskName: record.taskName }).then();
      createMessage.success('任务已启动');
      record.status = 1;
    } else {
      await StopTask({ taskName: record.taskName, deviceName: record.taskName }).then();
      createMessage.success('任务已停止');
      record.status = 0;
    }
    setTimeout(() => {
      reload();
    }, 1000);
  }
  function handleSuccess() {
    createMessage.success('操作成功');
    reload();
  }
  const [addModal, { openModal: addOpenModal }] = useModal();
  function addTask() {
    const usedUrl = getDataSource() ? getDataSource().map((item) => item.deviceName) : [];
    addOpenModal(true, { usedUrl });
  }
  function addSuccess() {
    createMessage.success('新建成功');
    reload();
  }
  const [algoConfigModal, { openModal: algoConfigOpenModal }] = useModal();
  async function algoConfig() {
    const res = await getAlgorithm();
    algoConfigOpenModal(true, res);
  }
  function configSuccess() {
    createMessage.success('配置成功');
    reload();
  }
</script>
