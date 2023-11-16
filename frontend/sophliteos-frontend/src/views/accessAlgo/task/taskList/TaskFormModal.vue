<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="submit" width="80vw">
    <BasicForm @register="registerForm">
      <template #urlList="{ model, field }">
        <a-select v-model:value="model[field]" :options="urlList" />
      </template>
      <template #abilitie="{ model, field }">
        <div class="editTask">
          <a-checkbox-group v-model:value="model[field]" :options="options" />
        </div>
      </template>
    </BasicForm>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { taskFormSchema } from './taskData';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { modTask } from '/@/api/task/index';
  import { options } from '/@/components/Data/algoData';
  import { CheckboxGroup } from 'ant-design-vue';
  import { getVideosList } from '/@/api/dataSource/index';
  import { ref } from 'vue';
  const urlList = ref();
  const { t } = useI18n();
  const ACheckboxGroup = CheckboxGroup;
  const title = t('taskList.taskList.editTask');
  const emit = defineEmits(['success', 'register']);
  const [registerForm, { setFieldsValue, resetFields, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 10 },
    schemas: taskFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
  });
  // onMounted(async () => {
  //   const res = await getVideosList();
  //   urlList.value = res.map((item) => {
  //     //@ts-ignore
  //     const { name, deviceId, url } = item;
  //     return { label: `${name}-${deviceId}-${url}`, value: `${name}-${deviceId}-${url}` };
  //   });
  // });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setModalProps({ confirmLoading: false });
    setFieldsValue({
      ...data.record,
      urlList: data.record.deviceName,
    });
    const res = await getVideosList();
    urlList.value = res.map((item) => {
      //@ts-ignore
      const { name, deviceId, url } = item;
      return { label: `${name}-${deviceId}-${url}`, value: `${name}-${deviceId}-${url}` };
    });
    data.usedUrl = data.usedUrl.filter((name) => name !== data.record.deviceName);
    urlList.value = urlList.value.filter(
      (item) => !data.usedUrl.includes(item.label.split('-')[0]),
    );
  });
  async function submit() {
    try {
      const values = await validate();
      let deviceId = '';
      for (const item of urlList.value) {
        // 使用正则表达式来匹配name值
        const parts = item.label.split('-');

        if (parts[0] === values.urlList.split('-')[0]) {
          // 如果找到匹配的name值，提取deviceId并赋给deviceId变量
          deviceId = parts[1];
          break; // 找到匹配后可以退出循环
        }
      }

      setModalProps({ confirmLoading: true });
      // TODO custom api
      await modTask({ taskName: values.taskName, deviceId: deviceId, abilities: values.abilities });
      emit('success');
      deviceId = '';
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>

<style lang="less" scoped>
  .ant-checkbox-group {
    margin-top: 30px;
  }
</style>
<style lang="less">
  .editTask {
    .ant-checkbox-group-item {
      width: 190px;
    }
  }
</style>
