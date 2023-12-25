<template>
  <div class="device_info">
    <span>{{ t('overview.PCIE.deviceInfo') }}</span>
    <div class="device_image">
      <div class="img">
        <img
          :src="getImg(originData.deviceType)"
          alt=""
          style="display: block; height: 97px; width: 118px; margin: 0 auto"
        />
      </div>
      <div class="editName">
        <p
          v-if="!edit"
          style="font-size: 14px; color: #323233; font-weight: 500; padding-bottom: 2px; margin: 0"
          >{{ deviceInfo.deviceName }}</p
        >
        <a-input
          v-if="edit"
          v-model:value="deviceName"
          ref="deviceNameInput"
          @blur="handleBlur"
          @keyup.enter="handleBlur"
        />
        <a-tooltip :title="t('overview.device.editType')" placement="bottom" :visible="editType">
          <EditOutlined
            @click="toggleEdit"
            v-if="!edit"
            @mouseenter="editType = true"
            @mouseleave="editType = false"
          />
        </a-tooltip>
      </div>
      <p style="font-size: 12px; color: #8c8c8c; font-weight: 400"
        >SN:{{ ' ' + originData.deviceSn }}</p
      >
    </div>
    <table>
      <tr>
        <th>{{ t('overview.computePower') }}</th>
      </tr>
      <tr>
        <td>TPU</td>
        <td>{{ originData.tpu.total || 0 }}TOPS</td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.TPUmemory') }}</td>
        <td>{{ unitSize(originData.tpu.memTotal) }}</td>
      </tr>
    </table>

    <div style="border: 1px solid #e6f2ff"></div>
    <table>
      <tr>
        <th>{{ t('overview.PCIE.hostInfo') }}</th>
      </tr>
      <tr>
        <td>CPU</td>
        <td>
          {{ originData.cpu.type }}
          {{ originData.cpu.cores + t('overview.core') }}
        </td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.sdkVersion') }}</td>
        <td>{{ originData.sdkVersion }}</td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.ssmVersion') }}</td>
        <td>{{ originData.bmssmVersion }}</td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.runTime') }}</td>
        <td>{{ originData.runTime }}</td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.operateSys') }}</td>
        <td>{{ originData.operatingSystem }}</td>
      </tr>
    </table>

    <div style="border: 1px solid #e6f2ff"></div>
    <table>
      <tr>
        <th>{{ t('overview.PCIE.netInfo') }}</th>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.ipAdress') }}</td>
        <td>{{ originData.deviceIp }}</td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.netName') }}</td>
        <td>{{ netCardInfo[0].name }}</td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.bindWidth') }}</td>
        <td>{{ netCardInfo[0].bandwidth + 'Mbps' }}</td>
      </tr>
      <tr>
        <td>{{ t('overview.PCIE.macAdress') }}</td>
        <td>{{ netCardInfo[0].mac }}</td>
      </tr>
    </table>
  </div>
</template>
<script lang="ts" setup>
  import { useI18n } from '/@/hooks/web/useI18n';
  import { storeToRefs } from 'pinia';
  import { Tooltip } from 'ant-design-vue';
  import { useDeviceInfo } from '/@/store/modules/overview';
  import { computed, ref, nextTick } from 'vue';
  import { EditOutlined } from '@ant-design/icons-vue';
  import { setDeviceInfoApi } from '/@/api/overview/index';
  const deviceStore = useDeviceInfo();
  const { originData, deviceInfo } = storeToRefs(deviceStore);
  const { t } = useI18n();
  const ATooltip = Tooltip;
  const netCardInfo = computed(() => {
    return originData.value.netCard.filter((el) => {
      return el.ip === originData.value.deviceIp;
    });
  });
  function getImg(url) {
    return new URL(`../../../assets/images/${url}.png`, import.meta.url).href;
  }
  const unitSize = computed(() => {
    return function (size, suffix = '') {
      if (!size) return 0;
      const unitStep = ['M', 'G', 'T', 'P', 'E'];
      let step = 0;
      while (size >= 1024) {
        size = size / 1024;
        step++;
      }
      return size.toFixed(1) + unitStep[step] + suffix;
    };
  });
  // 输入框失去焦点逻辑
  const edit = ref(false);
  const handleBlur = async () => {
    edit.value = false;
    const deviceNameTrim = deviceName.value.trim();

    if (deviceNameTrim === '') return;
    const params = {
      deviceName: deviceNameTrim,
      deviceType: originData.value.deviceType,
    };
    const result = await setDeviceInfoApi(params);
    if (result && result.code === 0) {
      deviceStore.updateDevice('deviceName', deviceNameTrim);
    }
  };
  const deviceName = ref('');
  const editType = ref(false); //是否展示tooltip
  const deviceNameInput = ref();
  const toggleEdit = async () => {
    editType.value = false;
    edit.value = true;
    await nextTick();
    deviceNameInput.value.focus();
  };
</script>
<style scoped lang="less">
  .device_info {
    padding: 24px 21px;
    min-width: 250px;

    span {
      height: 20px;
      margin-left: 5px;
      font-size: 18px;
      font-weight: 520;
      color: #323233;
    }

    .device_image {
      background-color: rgba(249, 251, 253, 1);
      border-radius: 9px;
      height: 213px;
      margin-top: 25px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;

      .img {
        width: 237px;
        height: 97px;
        margin: 20px;

        img {
          display: block;
          height: 97px;
          width: 200px !important;
          margin: 0 auto;
        }
      }

      .editName {
        display: flex;
      }

      .device_run {
        border-radius: 10px;
        height: 20px;
        width: 70px;
        position: relative;
        text-align: center;
        left: -63.5px;

        span {
          width: 42px;
          height: 20px;
          font-size: 14px;
          font-weight: 400;
          line-height: 20px;
        }
      }
    }

    table {
      width: 100%;
      text-align: left;

      margin: 20px auto 15px 3px;

      tr {
        height: 28px;
        margin: 14px auto;

        th {
          width: 64px;
          height: 22px;
          font-size: 16px;
          font-weight: 400;
          color: #323233;
        }

        td {
          height: 20px;
          font-size: 14px;

          font-weight: 400;
          color: #323233;
        }

        td:first-child {
          width: 90px;
          color: #8c8c8c;
        }

        td:last-child {
          min-width: 140px;
        }
      }
    }

    table:last-of-type {
      border-bottom: 0;
      margin-bottom: 0;
    }
  }
</style>
