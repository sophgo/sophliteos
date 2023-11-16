<template>
  <a-tabs v-model:activeKey="activeKey" class="!m-4 !p-4 bg-white">
    <a-tab-pane key="wan" :tab="t('maintenance.threshold.title')">
      <a-skeleton :loading="pageLoading" active>
        <a-form
          :model="formState"
          v-bind="formItemLayout"
          class="flex flex-wrap justify-around"
          size="large"
        >
          <a-form-item
            v-for="item of formItemList"
            :key="item.field"
            :label="item.label"
            class="w-2/5"
          >
            <a-input-number
              v-model:value="formState[item.field]"
              :placeholder="item.placeholder || placeholder"
              :min="0"
              :max="item.max || 100"
            >
              <template #addonAfter>
                <span v-if="item.unit" :style="{ color }">{{ item.unit }}</span>
                <percentage-outlined v-else :style="{ color }" />
              </template>
            </a-input-number>
          </a-form-item>
          <a-form-item class="w-2/5" />
          <a-form-item :wrapper-col="{ offset: 8, span: 16 }" class="w-2/5 !mr-1/2">
            <a-button type="primary" @click="submitForm" :loading="loading">{{
              t('sys.btn.confirm')
            }}</a-button>
          </a-form-item>
        </a-form>
      </a-skeleton>
    </a-tab-pane>
  </a-tabs>
</template>
<script lang="ts" setup>
  import { reactive, ref, onMounted } from 'vue';
  import { PercentageOutlined } from '@ant-design/icons-vue';
  import type { UnwrapRef } from 'vue';
  import { setAlarm, getAlarm } from '/@/api/maintenance/index';
  import { Tabs } from 'ant-design-vue';

  import { useI18n } from '/@/hooks/web/useI18n';
  // import { useMessage } from '/@/hooks/web/useMessage';

  import { AlarmParams } from '/@/api/maintenance/model/index';

  // const { createErrorModal } = useMessage();
  const { t } = useI18n();

  const ATabs = Tabs;
  const ATabPane = Tabs.TabPane;

  const placeholder = t('sys.form.phNumber');
  const color = '#0960BD';

  const activeKey = ref('wan');
  const formItemList = [
    {
      label: t('maintenance.threshold.fanSpeed'),
      field: 'fanSpeed',
      unit: 'r/min',
      placeholder: t('sys.form.placeholder'),
      max: 1000000,
    },
    {
      label: t('maintenance.threshold.boardTemperature'),
      field: 'boardTemperature',
      unit: '°C',
    },
    {
      label: t('maintenance.threshold.coreTemperature'),
      field: 'coreTemperature',
      unit: '°C',
    },
    {
      label: t('maintenance.threshold.cpuRate'),
      field: 'cpuRate',
    },
    {
      label: t('maintenance.threshold.totalMemoryScale'),
      field: 'totalMemoryScale',
    },
    // {
    //   label: t('maintenance.threshold.memoryScale'),
    //   field: 'systemScale',
    // },
    // {
    //   label: t('maintenance.threshold.videoScale'),
    //   field: 'videoScale',
    // },
    {
      label: t('maintenance.threshold.tpuScale'),
      field: 'tpuScale',
    },
    // {
    //   label: t('maintenance.threshold.externalHardDiskRate'),
    //   field: 'externalHardDiskRate',
    // },
    {
      label: t('maintenance.threshold.diskRate'),
      field: 'diskRate',
    },
  ];
  const formState: UnwrapRef<AlarmParams> = reactive({
    fanSpeed: 0,
    boardTemperature: 0,
    coreTemperature: 0,
    cpuRate: 90,
    totalMemoryScale: 90,
    // systemScale: 90,
    // videoScale: 90,
    tpuScale: 90,
    // externalHardDiskRate: 90,
    diskRate: 90,
  });
  const formItemLayout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
  };

  const pageLoading = ref(true);
  const init = async () => {
    const result = await getAlarm();
    pageLoading.value = false;
    if (result) {
      formState.fanSpeed = result.fanSpeed;
      formState.boardTemperature = result.boardTemperature;
      formState.coreTemperature = result.coreTemperature;
      formState.cpuRate = result.cpuRate * 100;
      formState.totalMemoryScale = result.totalMemoryScale * 100;
      // formState.systemScale = result.systemScale * 100;
      // formState.videoScale = result.videoScale * 100;
      formState.tpuScale = result.tpuScale * 100;
      // formState.externalHardDiskRate = result.externalHardDiskRate * 100;
      formState.diskRate = result.diskRate * 100;
    }
  };
  onMounted(() => {
    init();
  });
  const loading = ref(false);
  const submitForm = async () => {
    try {
      loading.value = true;
      const params = {
        ...formState,
      };
      // 不需要转变的字段
      const staticFields = ['fanSpeed', 'boardTemperature', 'coreTemperature'];
      Object.keys(params).forEach((key) => {
        if (!staticFields.includes(key)) {
          params[key] = params[key] / 100;
        }
      });
      await setAlarm(params);
    } catch (error) {
      // createErrorModal({
      //   title: t('sys.api.errorTip'),
      //   content: (error as unknown as Error).message || t('sys.api.networkExceptionMsg'),
      //   getContainer: () => document.body,
      // });
    } finally {
      loading.value = false;
    }
  };
</script>
