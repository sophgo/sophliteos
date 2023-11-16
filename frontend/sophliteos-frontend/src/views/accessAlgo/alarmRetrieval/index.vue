<template>
  <div style="margin: 10px; display: flex; flex-direction: column">
    <div class="p-4 mb-2 bg-white">
      <BasicForm @register="registerForm">
        <template #alarms="{ model, field }">
          <a-select v-model:value="model[field]" :options="ableList" mode="multiple" />
        </template>
      </BasicForm>
    </div>
    <div class="p-2 bg-white">
      <div>
        <a-p v-if="total"
          >{{
            t('alarmRetrieval.alarm.totalNumber') + ': ' + total + t('alarmRetrieval.alarm.images')
          }}
        </a-p>
        <a-p v-else
          >{{ t('alarmRetrieval.alarm.totalNumber') + ': 0' + t('alarmRetrieval.alarm.images') }}
        </a-p>
        <a-p style="margin-left: 20px" v-if="maxSize"
          >{{
            t('alarmRetrieval.alarm.useSpace') +
            ' : ' +
            usedSize +
            '  ' +
            t('alarmRetrieval.alarm.maxSpace') +
            ' : ' +
            maxSize
          }}
        </a-p>
        <a-button
          v-if="usedSize"
          style="margin-left: 20px; color: blue; outline: none; border: none"
          @click="settingSpace()"
          >[{{ t('alarmRetrieval.alarm.setSpace') }}]
        </a-button>
      </div>
      <div style="width: 100%">
        <div style="min-height: calc(100vh - 325px); max-height: max-content">
          <a-card
            hoverable
            class="myCard"
            v-for="(item, index) in list"
            :key="index"
            @click="getImage(item)"
          >
            <template #cover>
              <div>
                <img :src="item.canvas" alt="" style="width: 310px; height: 220px" />
              </div>
            </template>
            <a-card-meta>
              <template #description>
                <div class="des_card">
                  <p style="color: blue">
                    <Icon icon="ri:time-line" />{{
                      ' ' + dayjs(item.time).format('YYYY-MM-DD HH:mm:ss')
                    }}</p
                  >
                  <p style="color: blue">
                    <Icon icon="material-symbols:where-to-vote-outline" />{{
                      ' ' + item.deviceName
                    }}</p
                  >
                </div>
              </template>
            </a-card-meta>
          </a-card>
        </div>
        <a-pagination
          v-model:current="pageNo"
          :total="total"
          :pageSize="pageSize"
          @change="onChange"
          size="small"
          show-quick-jumper
          :show-total="
            (total) =>
              ` ${t('alarmRetrieval.alarm.total')} ${total} ${t('alarmRetrieval.alarm.images')}`
          "
          style="text-align: right"
          class="bg-white"
        />
      </div>
    </div>

    <SettingSpace @register="registerModal" @success="handleSuccess" />
  </div>
</template>

<script setup>
  // import { useMessage } from '/@/hooks/web/useMessage';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useModal } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form';
  import SettingSpace from './component/SettingSpace.vue';
  import { Card, CardMeta, Pagination, Select } from 'ant-design-vue';
  import { Icon } from '/@/components/Icon';
  import { AlarmList } from '/@/api/alrmRetrieval/index';
  import { onMounted, ref, watchEffect } from 'vue';
  import { useGo } from '/@/hooks/web/usePage';
  import { getTaskList } from '/@/api/task/index';
  import { option, revOption } from '/@/components/Data/algoData';
  import { alarmInfo } from '/@/store/modules/alrmRetrieval';
  import { cropperImageByRect } from '/@/utils/image/index';
  import dayjs from 'dayjs';
  import { schemas } from './Data';
  const { t } = useI18n();
  let params = {
    beginTime: 0,
    endTime: 0,
    pageNo: 1,
    pageSize: 20,
    deviceName: '',
    alarms: [],
  };
  const ACard = Card;
  const ACardMeta = CardMeta;
  const ASelect = Select;
  const APagination = Pagination;
  const go = useGo();
  const list = ref(); //卡片列表
  const total = ref();
  const pageNo = ref();
  const pageSize = ref();
  let maxSize = ref('');
  let usedSize = ref('');
  const ableList = ref(); //检索列表
  onMounted(() => {
    getList(params);
    ablelist();
  });
  async function ablelist() {
    ableList.value = await getTaskList({ pageNo: 1, pageSize: 10 }).then((result) => {
      return [...new Set(result.items.map((item) => item.abilities).flat())].map((item) => {
        return { lable: item, value: option[item] };
      });
    });
  }
  async function getList(params) {
    const res = await AlarmList(params);
    list.value = res.items;
    total.value = res.total;
    pageNo.value = res.pageNo;
    pageSize.value = res.pageSize;
    maxSize.value = convert(res.maxSize);
    usedSize.value = res.usedSize;
  }
  const [registerModal, { openModal }] = useModal();
  function settingSpace() {
    openModal(true, { max, usedSize });
  }
  function handleSuccess() {
    getList(params);
  }
  const [registerForm, { validate }] = useForm({
    schemas: schemas,
    labelWidth: 100,
    labelCol: {
      sm: { span: 24 }, // 小型平板电脑
      md: { span: 12 }, // 中型平板电脑
      lg: { span: 7 }, // 大屏幕桌面
    },
    baseColProps: { span: 6 },
    actionColOptions: { span: 24 },
    autoSubmitOnEnter: true,
    submitFunc: handleSubmit,
  });
  async function handleSubmit() {
    const data = await validate();
    const alarmList = Object.keys(data.alarms).map((key) => revOption[data.alarms[key]]);
    data.beginTime = data.beginTime ? `${data.beginTime} 00:00:00` : 0;
    data.endTime = data.endTime ? `${data.endTime} 23:59:59` : 0;
    params = {
      beginTime: new Date(data.beginTime).getTime(),
      endTime: new Date(data.endTime).getTime(),
      pageNo: 1,
      pageSize: 20,
      deviceName: data.deviceName,
      alarms: alarmList,
    };
    getList(params);
  }
  async function onChange(p, size) {
    const data = await validate();
    data.beginTime = data.beginTime ? `${data.beginTime} 00:00:00` : data.beginTime;
    data.endTime = data.endTime ? `${data.endTime} 23:59:59` : data.endTime;
    const alarmList = Object.keys(data.alarms).map((key) => revOption[data.alarms[key]]);
    params = {
      beginTime: new Date(data.beginTime).getTime(),
      endTime: new Date(data.endTime).getTime(),
      pageNo: p,
      pageSize: size,
      deviceName: data.deviceName,
      alarms: alarmList,
    };
    getList(params);
  }
  function getImage(item) {
    const param = item.image.split('?')[1];
    const store = alarmInfo();
    store.setInfo(item);
    go('/accessAlgo/AlarmRetrieval/AlarmDetail/' + param);
  }
  watchEffect(() => {
    if (list.value) {
      list.value.map((item) => {
        cropperImageByRect(item.image, { width: '300', height: '220' }, item.boxes[0]).then(
          (res) => {
            item.canvas = res;
          },
        );
      });
    }
  });
  const max = ref(0);
  function convert(sizeInMB) {
    max.value = sizeInMB;
    if (sizeInMB < 1024) {
      return sizeInMB + ' MB';
    } else if (sizeInMB < 1024 * 1024) {
      const sizeInGB = sizeInMB / 1024;
      return sizeInGB.toFixed(2) + ' GB';
    } else {
      const sizeInTB = sizeInMB / (1024 * 1024);
      return sizeInTB.toFixed(2) + ' TB';
    }
  }
</script>

<style lang="less">
  .des_card {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    height: 20px;
  }

  .myCard {
    display: inline-block;
    margin-left: 15px !important;
    width: 310px;

    .ant-card-body {
      padding: 15px !important;
    }
  }
</style>
