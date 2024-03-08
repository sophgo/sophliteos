<template>
  <div class="info">
    <div class="left-box">
      <a-descriptions
        :column="1"
        layout="horizontal"
        :title="t('alarmRetrieval.alarm.alarmDetail')"
      >
        <a-descriptions-item>
          <img :src="image" alt="" class="img" />
        </a-descriptions-item>
      </a-descriptions>
    </div>
    <div class="right-box">
      <a-descriptions :column="1" layout="horizontal" :title="t('alarmRetrieval.alarm.detailInfo')">
        <a-descriptions-item :label="t('alarmRetrieval.alarm.alarmType')">{{
          option[iamgeInfo.alarmType]
        }}</a-descriptions-item>
        <a-descriptions-item :label="t('alarmRetrieval.alarm.captureTime')">{{
          dayjs(iamgeInfo.time).format('YYYY-MM-DD HH:mm:ss')
        }}</a-descriptions-item>
        <a-descriptions-item :label="t('dataSource.videoManage.deviceName')">{{
          iamgeInfo.deviceName
        }}</a-descriptions-item>
        <!-- <a-descriptions-item :label="t('alarmRetrieval.alarm.confidenceLevel')">{{
          iamgeInfo.itemsInBox[0].confidence
        }}</a-descriptions-item> -->
      </a-descriptions>
    </div>
  </div>
</template>
<script setup>
  import { Descriptions, DescriptionsItem } from 'ant-design-vue';
  import { onMounted, ref } from 'vue';
  import { useRoute } from 'vue-router';
  import { getAlarmImage } from '/@/api/alrmRetrieval/index';
  import { storeToRefs } from 'pinia';
  import { alarmInfo } from '/@/store/modules/alrmRetrieval';
  import { option } from '/@/components/Data/algoData';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { drawRectOnImage } from '/@/utils/image/index';
  import dayjs from 'dayjs';
  const { t } = useI18n();
  const ADescriptions = Descriptions;
  const ADescriptionsItem = DescriptionsItem;
  const route = useRoute();
  const url = ref(route.params?.image);
  const image = ref();

  const store = alarmInfo();
  const { iamgeInfo } = storeToRefs(store);
  onMounted(() => {
    getImage(url);
  });

  async function getImage(url) {
    await getAlarmImage(url.value).then((res) => {
      const myBlob = new window.Blob([res.data], { type: 'image/jpeg' });
      drawRectOnImage(
        window.URL.createObjectURL(myBlob),
        iamgeInfo.value.boxes,
        iamgeInfo.value.itemsInBox,
      ).then((res) => {
        image.value = res;
      });
    });
  }
</script>
<style lang="less" scoped>
  .info {
    display: flex;
    margin: 20px 10px;
  }

  .left-box {
    flex: 8; /* 8:2 宽度比例 */
    padding: 20px; /* 可以根据需要添加内边距 */
  }

  .right-box {
    flex: 2;
    border-left: 1px solid #e8e8e8; /* 可以根据需要添加边框样式 */
  }

  .img {
    margin: 0 auto;
  }
</style>
