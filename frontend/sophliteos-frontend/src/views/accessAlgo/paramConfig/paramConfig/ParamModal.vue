<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="Submit">
    <BasicForm @register="registerForm">
      <template #width="{ model, field }">
        <a-input v-model:value="model[field]" placeholder="请输入数值">
          <template #addonBefore>
            <span>宽度</span>
          </template>
          <template #addonAfter>
            <span>px</span>
          </template>
        </a-input>
      </template>

      <template #height="{ model, field }">
        <a-input v-model:value="model[field]" placeholder="请输入数值">
          <template #addonBefore>
            <span>高度</span>
          </template>
          <template #addonAfter>
            <span>px</span>
          </template></a-input
        >
      </template>
      <template #interval="{ model, field }">
        <a-input v-model:value="model[field]" placeholder="请输入数值">
          <template #addonAfter>
            <span>秒</span>
          </template></a-input
        >
      </template>
      <template #threshold="{ model, field }">
        <a-input v-model:value="model[field]" placeholder="请输入数值">
          <template #addonAfter>
            <span>0-1</span>
          </template></a-input
        >
      </template>
    </BasicForm>
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { paramFormSchema } from './Data';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { algoConfigMod } from '/@/api/paramConfig/index';
  // import { Input } from 'ant-design-vue';
  const { t } = useI18n();
  const title = t('paramConfig.param.paramConfig');
  import { option, revOption } from '/@/components/Data/algoData';
  // const AInput = Input;
  const emit = defineEmits(['success', 'register']);
  const [registerForm, { setFieldsValue, resetFields, validate }] = useForm({
    labelWidth: 150,
    baseColProps: { span: 20 },
    schemas: paramFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 15,
    },
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    resetFields();
    setModalProps({ confirmLoading: false });
    setFieldsValue({
      ...data.record,
      ability: option[data.record.ability],
      width: data.record.minBox.width,
      height: data.record.minBox.height,
    });
  });
  async function Submit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      // TODO custom api
      const minBox = { width: 0, height: 0 };
      minBox.width = Number(values.width);
      minBox.height = Number(values.height);
      values.ability = revOption[values.ability];
      // console.log(values);
      await algoConfigMod({
        ability: values.ability,
        minBox: minBox,
        interval: Number(values.interval),
        threshold: Number(values.threshold),
      });
      emit('success');
      closeModal();
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
