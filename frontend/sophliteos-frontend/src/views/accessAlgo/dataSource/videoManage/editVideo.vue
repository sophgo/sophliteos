<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="submit">
    <BasicForm @register="registerForm" style="padding-top: 20px" />
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { video, isAdd } from './tableData';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { ModDevice, AddDevice } from '/@/api/dataSource/index';
  import { ref } from 'vue';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register']);

  const title = ref();
  const deviceId = ref();
  const [registerForm, { resetFields, validate, setFieldsValue }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 21 },
    schemas: video,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 20,
    },
    submitButtonOptions: {
      text: t('dataSource.mediaServers.save'),
    },
  });
  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setModalProps({ confirmLoading: false });
    if (data.record === 'add') {
      title.value = t('dataSource.videoManage.addDevice');
      isAdd(false);
    } else {
      title.value = t('dataSource.videoManage.editDevice');
      isAdd(true);
      setFieldsValue({
        ...data.record,
      });

      deviceId.value = data.record.deviceId;
    }
    console.log(video);
  });
  async function submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });

      if (title.value == t('dataSource.videoManage.editDevice')) {
        await ModDevice({
          name: values.name,
          protocol: values.protocol,
          ptzType: values.ptzType,
          url: values.url,
          deviceId: deviceId.value,
        });
      } else {
        await AddDevice(values);
      }

      emit('success');
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
<style lang="less"></style>
