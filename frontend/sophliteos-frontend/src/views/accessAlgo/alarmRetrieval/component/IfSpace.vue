<template>
  <BasicModal
    v-bind="$attrs"
    @register="SpaceModal"
    @ok="Submit"
    width="30vw"
    :okText="t('alarmRetrieval.alarm.saveSetting')"
  >
    <Alert :message="message" show-icon type="warning" style="margin-top: 100px">
      <template #icon><WarningOutlined /></template
    ></Alert>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { Alert } from 'ant-design-vue';
  import { WarningOutlined } from '@ant-design/icons-vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { modSize } from '/@/api/alrmRetrieval/index';
  import { ref } from 'vue';
  const { t } = useI18n();
  const space = ref();
  const emit = defineEmits(['success', 'register']);
  const message = t('alarmRetrieval.alarm.notice');
  const [SpaceModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    setModalProps({ confirmLoading: false });
    space.value = data.sliderValue;
  });
  async function Submit() {
    await modSize({ maxSize: space.value });
    emit('success');
    closeModal();
  }
</script>
