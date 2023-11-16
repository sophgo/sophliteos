<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="submit">
    <div class="config">
      <BasicForm @register="registerForm" style="padding-top: 60px" />
    </div>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { algoConfigFormSchema } from './taskData';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { addAlgorithm } from '/@/api/task/index';
  import { toRaw } from 'vue';
  const { t } = useI18n();
  const title = t('taskList.taskList.algoConfigTitle');
  const emit = defineEmits(['configsuccess', 'register']);
  const [registerForm, { resetFields, validate, setFieldsValue }] = useForm({
    labelWidth: 80,
    baseColProps: { span: 21 },
    schemas: algoConfigFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
  });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    console.log(toRaw(data));

    setFieldsValue({
      ...data,
    });
    setModalProps({ confirmLoading: false });
  });
  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      await addAlgorithm({ ip: values.ip, port: Number(values.port) });
      emit('configsuccess');
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less"></style>
