<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="Submit" width="30%">
    <Alert :message="message" show-icon class="alert" type="warning" />
    <div class="slider">
      {{ t('alarmRetrieval.alarm.setting') }}：<a-input-number
        v-model:value="sliderValue"
        style="width: 60% !important"
      >
        <template #addonAfter>
          <span>MB</span>
        </template>
      </a-input-number>
    </div>
  </BasicModal>
  <IfSpace @register="SpaceModal" @success="success" />
</template>
<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { ref } from 'vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { Alert } from 'ant-design-vue';
  import { useModal } from '/@/components/Modal';
  import { modSize } from '/@/api/alrmRetrieval/index';
  import IfSpace from './IfSpace.vue';
  import { useMessage } from '/@/hooks/web/useMessage';
  const { createMessage } = useMessage();
  const { t } = useI18n();
  const title = t('alarmRetrieval.alarm.setSpace');
  const message = t('alarmRetrieval.alarm.notice');
  const sliderValue = ref();
  const usedSize = ref();
  const maxSize = ref();
  const emit = defineEmits(['success', 'register']);

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    setModalProps({ confirmLoading: false });
    maxSize.value = +data.max;
    usedSize.value = Number(data.usedSize.split(' ')[0]);
    sliderValue.value = maxSize.value; //当前设置
  });
  const [SpaceModal, { openModal }] = useModal();
  async function Submit() {
    if (!Number.isInteger(sliderValue.value)) {
      createMessage.error('请输入一个整数');
    } else if (sliderValue.value < usedSize.value) {
      openModal(true, { sliderValue });
      closeModal();
    } else {
      try {
        setModalProps({ confirmLoading: true });
        // TODO custom api
        await modSize({ maxSize: sliderValue.value });
        emit('success');
        closeModal();
      } finally {
        setModalProps({ confirmLoading: false });
      }
    }
  }
  function success() {
    emit('success');
  }
</script>
<style lang="less" scoped>
  .alert {
    font-size: 14px;
    margin: 0 auto;
  }

  .slider {
    margin-top: 70px;
    position: relative;
    padding: 0px 30px;
    display: flex;
    align-items: center;
  }
</style>
