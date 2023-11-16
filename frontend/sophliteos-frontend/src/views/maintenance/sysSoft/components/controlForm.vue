<template>
  <a-skeleton :loading="loading" active>
    <div
      v-if="!deviceInfoStore.isSingleBoard"
      style="font-size: 16px; margin: 20px 0 0 20px; font-weight: 550"
      >{{
        isSoftware
          ? t('maintenance.systemUpdate.softwareUpdate')
          : t('maintenance.systemUpdate.systemUpdate')
      }}</div
    >
    <div v-else style="font-size: 16px; margin: 20px 0 0 20px; font-weight: 550">{{
      t('maintenance.systemUpdate.localUpdate1')
    }}</div>
    <div style="width: 20vw; margin-left: 20px">
      <a-form
        :model="formState"
        v-bind="formItemLayout"
        class="!my-4"
        @finish="handleUpload"
        labelAlign="left"
      >
        <a-form-item
          v-if="deviceInfoStore.isSingleBoard"
          :label="t('maintenance.systemUpdate.updateType')"
          :rules="[{ required: true }]"
        >
          {{ t('maintenance.systemUpdate.localUpdate') }}
        </a-form-item>
        <a-form-item
          :label="
            isSoftware
              ? t('maintenance.systemUpdate.currentSoftVersion')
              : isSsm
              ? t('maintenance.systemUpdate.currentSsmVersion')
              : t('maintenance.systemUpdate.currentVersion')
          "
          :rules="[{ required: true }]"
        >
          <span>{{
            isSoftware ? softVersion : isSsm ? deviceInfo.bmssmVersion : deviceInfo.sdkVersion
          }}</span>
        </a-form-item>
        <a-form-item
          name="file"
          :rules="[{ required: true, message: t(t('maintenance.systemUpdate.selectNeedFile')) }]"
          :label="t('sys.uploadFile.btnText')"
          v-model:value="formState.file"
        >
          <a-upload
            :file-list="fileList"
            :before-upload="beforeUpload"
            :progress="progress"
            @remove="handleRemove"
          >
            <a-button>
              <a-upload-outlined />
              {{ t('maintenance.systemUpdate.selectFile') }}
            </a-button>
            <br />
            <span class="tips">{{ t('maintenance.systemUpdate.fileFormat') }}</span>
          </a-upload>
        </a-form-item>
        <!-- <a-form-item :wrapper-col="{ span: 14, offset: 4 }" /> -->
        <a-button
          type="primary"
          :loading="uploading"
          html-type="submit"
          style="margin-top: 10px"
          :class="{ isSsmbtn: isSsm && !deviceInfoStore.isSingleBoard }"
        >
          {{
            uploading
              ? t('maintenance.systemUpdate.updaing')
              : t('maintenance.systemUpdate.startUpdate')
          }}
        </a-button>
      </a-form>
    </div>
  </a-skeleton>
</template>
<script lang="ts" setup>
  import { reactive, ref, onMounted } from 'vue';
  import type { UnwrapRef } from 'vue';
  import { storeToRefs } from 'pinia';
  import { Upload } from 'ant-design-vue';
  import { UploadOutlined } from '@ant-design/icons-vue';

  import { upgradeApi, upgradeSoftApi, upgradeSsmApi } from '/@/api/maintenance/index';
  import { useDeviceInfo } from '/@/store/modules/overview';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { buildUUID } from '/@/utils/uuid';
  import { AxiosCanceler } from '/@/utils/http/axios/axiosCancel';
  import { useGlobSetting } from '/@/hooks/setting';
  import { getSoftwareInfoApi } from '/@/api/overview/index';

  import { useI18n } from '/@/hooks/web/useI18n';
  const { t } = useI18n();

  const props = defineProps({
    isSoftware: {
      type: Boolean,
      default: false,
    },
    isSsm: {
      type: Boolean,
      default: false,
    },
  });

  // 获取当前软件版本
  const softVersion = ref('');
  const getSoftInfor = async () => {
    const result = await getSoftwareInfoApi();
    softVersion.value = result.buildname;
  };
  onMounted(async () => {
    if (props.isSoftware) {
      getSoftInfor();
    }
  });
  const { uploadUrl = '' } = useGlobSetting();

  const axiosCanceler = new AxiosCanceler();
  const { createMessage } = useMessage();
  const AUpload = Upload;
  const AUploadOutlined = UploadOutlined;

  const deviceInfoStore = useDeviceInfo();
  const { deviceInfo } = storeToRefs(deviceInfoStore);

  const loading = ref(false);

  interface FormState {
    type: string;
    ota: string;
    file: any;
  }

  const formState: UnwrapRef<FormState> = reactive({
    type: 'local',
    ota: '',
    file: '',
  });

  const formItemLayout = {
    labelCol: { span: 10 },
    wrapperCol: { span: 14 },
  };

  // 上传逻辑
  const fileList = ref<any>([]);
  const uploading = ref(false);
  const progress = {
    strokeColor: {
      '0%': '#108ee9',
      '100%': '#87d068',
    },
    strokeWidth: 3,
    format: (percent) => `${parseFloat(percent.toFixed(2))}%`,
    class: 'test',
  };
  const handleRemove = (file) => {
    const index = fileList.value.indexOf(file);
    const newFileList = fileList.value.slice();
    newFileList.splice(index, 1);
    fileList.value = newFileList;
    // 取消请求
    axiosCanceler.removePending({ method: 'post', url: uploadUrl });
    uploading.value = false;
  };
  // const { createMessage } = useMessage();
  const beforeUpload = (file) => {
    const isTgz = file.name.endsWith('.tgz');
    if (!isTgz) {
      createMessage.error('抱歉，仅支持上传.tgz格式文件，请重新选择文件！');
      return false;
    }
    const { size, name } = file;
    const commonItem = {
      uuid: buildUUID(),
      file,
      size,
      name,
      percent: 0,
      type: name.split('.').pop(),
    };
    fileList.value = [commonItem];
    formState.file = commonItem;
    return false;
  };
  const handleUpload = async () => {
    try {
      uploading.value = true;
      const data = await Promise.all(
        fileList.value.map((item) => {
          return uploadApiByItem(item);
        }),
      );
      uploading.value = false;
      // 生产环境:抛出错误
      const errorList = data.filter((item: any) => !item.success);
      if (errorList.length > 0) {
        throw errorList;
      } else {
        const text = props.isSoftware
          ? t('maintenance.systemUpdate.softwareUpdateStatus')
          : t('maintenance.systemUpdate.controlStartUpload');
        createMessage.success(text);
      }
    } catch (e) {
      uploading.value = false;
      // throw e;
    }
  };

  async function uploadApiByItem(item) {
    try {
      item.status = 'uploading';
      const currentApi = props.isSoftware
        ? upgradeSoftApi
        : props.isSsm
        ? upgradeSsmApi
        : upgradeApi;
      const dataParams = props.isSoftware ? {} : { module: 'ctrl' };
      const { data } = await currentApi(
        {
          data: dataParams,
          file: item.file,
        },
        function onUploadProgress(progressEvent: ProgressEvent) {
          const complete = ((progressEvent.loaded / progressEvent.total) * 100) | 0;
          item.percent = complete;
        },
      );
      item.status = 'success';
      item.responseData = data;
      return {
        success: true,
        error: null,
      };
    } catch (e) {
      item.status = 'error';
      return {
        success: false,
        error: e,
      };
    }
  }
</script>
<style scoped lang="less">
  .ant-form-item .ant-upload .tips {
    font-size: 12px;
    color: red;
    position: relative;
    top: 4px;
    opacity: 0.8;
  }

  .isSsmbtn {
    margin-top: 105px !important;
  }

  :deep(.ant-upload-list-item-progress) {
    bottom: -20px;
  }
</style>
