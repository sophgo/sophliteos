<template>
  <BasicModal v-bind="$attrs" @register="registerModal" @ok="submit" width="80vw">
    <BasicForm @register="registerForm">
      <template #urlList="{ model, field }">
        <a-select v-model:value="model[field]" :options="urlList" />
      </template>
      <template #abilitie="{ model, field }">
        <div class="task">
          <a-checkbox-group v-model:value="model[field]" :options="options" />
        </div>
      </template>
    </BasicForm>
    <template #title>
      <div style="display: flex; flex-direction: row">
        <div>{{ message1 }}</div>
        <Alert :message="message2" show-icon class="alert" />
      </div>
    </template>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { addTaskSchema } from './taskData';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { addTask } from '/@/api/task/index';
  import { Alert, CheckboxGroup, Select } from 'ant-design-vue';
  import { options } from '/@/components/Data/algoData';
  import { getVideosList } from '/@/api/dataSource/index';
  import { ref } from 'vue';
  const { t } = useI18n();
  const message1 = t('taskList.taskList.addTask');
  const message2 = t('taskList.taskList.message');
  const ACheckboxGroup = CheckboxGroup;
  const ASelect = Select;
  const urlList = ref();
  const emit = defineEmits(['addsuccess', 'register']);
  const [registerForm, { resetFields, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 10 },
    schemas: addTaskSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
  });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    const res = await getVideosList();
    urlList.value = res.map((item) => {
      //@ts-ignore
      const { name, deviceId, url } = item;
      return { label: `${name}-${deviceId}-${url}`, value: `${name}-${deviceId}-${url}` };
    });
    setModalProps({ confirmLoading: false });
    urlList.value = urlList.value.filter(
      (item) => !data.usedUrl.includes(item.label.split('-')[0]),
    );
  });

  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      // TODO custom api
      const deviceName = values.urlList.split('-')[0];
      const deviceId = values.urlList.split('-')[1];
      const url = values.urlList.split('-')[2];
      await addTask({
        taskName: values.taskName,
        abilities: values.abilities,
        deviceName: deviceName,
        deviceId: deviceId,
        url: url,
      });
      emit('addsuccess');
      resetFields();
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less" scoped>
  .alert {
    position: relative;
    left: 22%;
    top: 50%;
    width: 50%;
    height: 50px;
    font-size: 14px;
  }

  .ant-checkbox-group {
    margin-top: 30px;
  }
</style>
<style lang="less">
  .task {
    display: flex;
    flex-direction: row;

    .ant-checkbox-group-item {
      width: 190px;
    }
  }
</style>
