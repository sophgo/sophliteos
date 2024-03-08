<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    @ok="submit"
    :okText="t('dataSource.mediaServers.save')"
  >
    <BasicForm @register="registerForm" style="padding-top: 60px" />
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { schemas } from './tableData';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { PostAddMediaServer } from '/@/api/dataSource/index';
  const { t } = useI18n();
  const title = t('dataSource.mediaServers.configService');
  const emit = defineEmits(['success', 'register']);
  const [registerForm, { resetFields, validate, setFieldsValue }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 21 },
    schemas: schemas,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
  });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setModalProps({ confirmLoading: false });
    setFieldsValue({
      ...data.res,
    });
  });
  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      await PostAddMediaServer({ ip: values.ip, port: Number(values.port) });
      emit('success');
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less"></style>
