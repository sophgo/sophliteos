<template>
  <section class="wrap">
    <div class="left">
      <div class="left-top">
        <div class="runStatus">{{ t('overview.PCIE.runState') }}</div>
        <div class="item-card" style="margin-left: 2.5vw">
          <div class="percent">
            <a-progress
              type="circle"
              :width="114"
              :stroke-color="{
                '0%': '#F87B13',
                '100%': '#FCB22B',
              }"
              :percent="
                originData.tpu.total
                  ? Number(((originData.tpu.used / originData.tpu.total) * 100).toFixed(1))
                  : 0
              "
            />
          </div>
          <div class="_title">TPU</div>
          <div class="_detail">
            {{ originData.tpu.used + 'TOPS / ' + originData.tpu.total + 'TOPS' }}
          </div>
        </div>
        <div class="item-card">
          <div class="percent">
            <a-progress
              type="circle"
              :width="114"
              :stroke-color="{
                '0%': '#1FE5FF',
                '100%': '#2AA8FF',
              }"
              :percent="originData.cpu.usage.toFixed(1) || 0"
            />
          </div>
          <div class="_title">CPU</div>
          <div class="_detail">
            {{
              originData.cpu.cores +
              t('overview.core') +
              ' ' +
              originData.cpu.frequency / 1000 +
              'GHz'
            }}
          </div>
        </div>
        <div class="item-card">
          <div class="percent">
            <a-progress
              type="circle"
              :width="114"
              :stroke-color="{
                '0%': '#16E4B8',
                '100%': '#3CB49A',
              }"
              :percent="originData.memory.usage.toFixed(1) || 0"
            />
          </div>
          <div class="_title">{{ t('overview.memory') }}</div>
          <div class="_detail">
            {{
              unitSize((originData.memory.usage * originData.memory.total) / 100) +
              '/' +
              unitSize(originData.memory.total)
            }}
          </div>
        </div>
        <div class="item-card" style="margin-right: 1.5vw">
          <div class="percent">
            <a-progress
              type="circle"
              :width="114"
              :stroke-color="{
                '0%': '#888EF0',
                '100%': '#515AFF',
              }"
              :percent="disk.total ? Number(((disk.used / disk.total) * 100).toFixed(1)) : 0"
            />
          </div>
          <div class="_title">
            <div>{{ t('overview.PCIE.disk') }}</div>
            <a-popover ref="popover" placement="bottomRight">
              <template #content>
                <div
                  v-for="(item, index) in originData.disk"
                  :key="index"
                  style="margin: 20px 0 10px 0"
                >
                  <div style="font-size: 12px; color: #323233">
                    {{ t('overview.PCIE.diskName') + ':' + item.diskName }}
                  </div>
                  <div style="font-size: 12px; color: #323233"
                    >{{ t('overview.PCIE.onMount') + ':' + item.mountOn }}
                  </div>
                  <div style="font-size: 12px; color: #323233">
                    {{ t('overview.PCIE.diskCapacity') + ':' + unitSize(item.total - item.free) }} /
                    {{ unitSize(item.total) }}
                  </div>
                  <a-progress
                    :percent="parseInt(((item.total - item.free) * 100) / item.total) || 0"
                    color="#0960BD"
                    :strokeWidth="10"
                    define-back-color="#fff"
                    :showInfo="false"
                  />
                  <span style="color: #0960bd; font-size: 12px; font-weight: 500"
                    >{{ parseInt(((item.total - item.free) / item.total) * 100) || 0 }}%
                  </span>
                </div>
              </template>
              <a-button class="disks" v-if="originData.disk.length > 1">{{
                originData.disk.length + t('overview.ge')
              }}</a-button>
            </a-popover>
          </div>
          <div class="_detail">{{ unitSize(disk.used) + '/' + unitSize(disk.total) }}</div>
        </div>
      </div>
      <div class="left-bottom" v-if="originData.coreComputingUnit">
        <div class="dashboard">
          <div class="dashboard_title">{{ t('overview.PCIE.board') }}</div>
          <div class="boards">
            <a-tabs
              v-model:activeKey="activeKey"
              type="card"
              :tabBarGutter="20"
              :tabBarStyle="{ borderBottom: 'unset' }"
            >
              <a-tab-pane
                :tab="t('overview.PCIE.card') + String(index + 1)"
                v-for="(item, index) in originData.coreComputingUnit.board"
                :key="String(index)"
              >
                <div class="boards_inner">
                  <div class="item-card" style="margin-left: 2.5vw">
                    <img :src="getImg(item.boardType)" style="width: 194px; height: 94px" alt="" />
                    <div class="_title">{{ item.boardType }}</div>
                    <div class="_detail" style="min-width: 164px">{{ 'SN: ' + item.boardSn }}</div>
                  </div>
                  <div class="item-card">
                    <div class="percent">
                      <a-progress
                        type="circle"
                        :width="90"
                        strokeColor="#A8C3EC"
                        :percent="
                          item.tpuTotal
                            ? Number(((item.tpuUsed / item.tpuTotal) * 100).toFixed(1))
                            : 0
                        "
                      />
                    </div>
                    <div class="_title">TPU</div>
                    <div class="_detail" style="min-width: 97px">
                      {{ item.tpuUsed + 'TOPS /' + item.tpuTotal }}TOPS
                    </div>
                  </div>
                  <div class="item-card">
                    <div class="percent">
                      <a-progress
                        type="circle"
                        :width="90"
                        strokeColor="#A8C3EC"
                        :percent="
                          item.tpuTotal
                            ? Number(((item.memUsed / item.memTotal) * 100).toFixed(1))
                            : 0
                        "
                      />
                    </div>
                    <div class="_title">{{ t('overview.PCIE.TPUmemory') }}</div>
                    <div class="_detail">
                      {{ unitSize(item.memUsed) + '/' + unitSize(item.memTotal) }}
                    </div>
                  </div>
                  <div
                    class="item-card"
                    style="min-width: 75px; margin-right: 1.5vw; margin-bottom: 18px"
                  >
                    <i style="margin-left: 63px; width: 37px; height: 18px; font-size: 16px">
                      {{ item.temperature + '°C' }}</i
                    >
                    <img
                      src="../../assets/images/temperatrue.png"
                      alt=""
                      style="margin-left: 12px; height: 72px"
                    />
                    <div class="_title">{{ t('overview.temperature') }}</div>
                    <div class="_detail"></div>
                  </div>
                </div>
                <div class="chips">
                  <div class="chip_title">{{ 'AI' + t('overview.chip') }}</div>
                  <span>{{ t('overview.zong') + item.chip.length + t('overview.ge') }}</span>
                </div>
                <div class="chipTable">
                  <table style="width: 100%; text-align: center">
                    <tr>
                      <th style="width: 10%">{{ t('overview.chip') }}</th>
                      <th style="width: 10%">{{ t('overview.device.name') }}</th>
                      <th style="width: 20%">{{ t('overview.PCIE.slot') }}</th>
                      <th>TPU</th>
                      <th>{{ t('overview.temperature') }}</th>
                      <th>{{ t('overview.PCIE.TPUmemory') }}</th>
                    </tr>
                    <tr v-for="(items, i) in item.chip" :key="i">
                      <td>
                        <img
                          :src="getImg(chipList[items.chipType])"
                          alt=""
                          style="height: 50px; width: 72px; margin: 0 auto"
                        />
                      </td>
                      <td>{{ chipList[items.chipType] }}</td>
                      <td>{{ items.slot }}</td>
                      <td>
                        <div style="display: flex; flex-direction: column; align-items: center">
                          <span style="display: flex">{{ items.tpuUtililizationRate + '%' }}</span
                          ><span style="display: flex; color: #606266">{{
                            (
                              (items.tpuUtililizationRate * items.theoretialCalculationCapacity) /
                              100
                            ).toFixed(1) +
                            'TOPS/' +
                            items.theoretialCalculationCapacity +
                            'TOPS'
                          }}</span>
                        </div>
                      </td>
                      <td>
                        <div style="display: flex; flex-direction: column; align-items: center">
                          <span style="display: flex"> {{ items.temperature + '°C' }}</span>
                        </div>
                      </td>
                      <td>
                        <div style="display: flex; flex-direction: column; align-items: center">
                          <span style="display: flex">
                            {{
                              ((items.memoryUsedBytes / items.memoryTotalBytes) * 100).toFixed(1) +
                              '%'
                            }}</span
                          ><span style="display: flex; color: #606266"
                            >{{ unitSize(items.memoryUsedBytes) }}
                            /
                            {{ unitSize(items.memoryTotalBytes) }}
                          </span>
                        </div>
                      </td>
                    </tr>
                  </table>
                </div>
              </a-tab-pane>
            </a-tabs>
          </div>
        </div>
      </div>
      <div v-else class="empty-board">暂无板卡信息</div>
    </div>
    <div class="right">
      <right-info />
    </div>
  </section>
</template>
<script setup lang="ts">
  import { computed, reactive, ref } from 'vue';
  import { storeToRefs } from 'pinia';
  import { Progress, Tabs, Popover, Button } from 'ant-design-vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import rightInfo from './components/rightInfo.vue';
  import { useDeviceInfo } from '/@/store/modules/overview';
  const AProgress = Progress;
  const APopover = Popover;
  const ATabs = Tabs;
  const ATabPane = Tabs.TabPane;
  const AButton = Button;

  const activeKey = ref('0');
  const { t } = useI18n();
  const deviceStore = useDeviceInfo();
  const { originData } = storeToRefs(deviceStore);
  const chipList = {
    1: 'BM1684',
    2: 'BM1684X',
  };
  const disk = reactive({
    total: 0,
    used: 0,
  });
  originData.value.disk.forEach((item) => {
    disk.total += item.total;
    disk.used += item.total - item.free;
  });
  function getImg(url) {
    return new URL(`../../assets/images/${url}.png`, import.meta.url).href;
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
</script>
<style scoped lang="less">
  .wrap {
    margin-top: 18px;
    height: 100%;
    width: 100%;
    display: flex;
    flex-direction: row;

    .left {
      width: 100%;
      height: 100%;
      margin-left: 16px;
    }

    .right {
      min-width: 250px;
      background: #ffffff;
      border-radius: 10px;
      margin: 0 16px;
      display: flex;
      flex-direction: column;
    }
  }

  .empty-board {
    flex: 1;
    background: #fff;
    text-align: center;
    font-size: 20px;
    margin-top: 20px;
    padding: 50px;
    border-radius: 10px;
    min-height: 646px;
  }

  .left-top {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    height: 244px;
    width: 100%;
    background-color: #fff;
    position: relative;
    border-radius: 10px;

    .item-card {
      min-width: 138px;
    }

    .runStatus {
      height: 25px;
      font-size: 16px;
      font-weight: 500;
      color: #323233;
      position: absolute;
      top: 14px;
      left: 24px;
    }

    .disks {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 39px;
      height: 20px;
      border-radius: 10px;
      border: 1px solid #0d60bb;
      font-size: 14px;
      color: #095ebc;
      background: #fff;
    }

    .disks:hover {
      background: #fff !important;
      border: 1px solid #0d60bb !important;
    }
  }

  .item-card {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    font-size: 12px;
    padding: 12px;
    box-sizing: border-box;
    border-radius: 10px;
    margin-top: 20px;

    & + .item-card {
      margin-left: 15px;
    }
    // .percent {
    //   margin-top: 25px;
    // }

    ._title /deep/ {
      display: flex;
      flex-direction: row;
      line-height: 20px;
      font-size: 14px;
      font-weight: 500;
      color: #323233;
      margin: 14px 0 0 4px;
    }

    ._detail {
      font-weight: 400;
      color: #606266;
      font-size: 14px;
    }
  }

  .left-bottom {
    margin-top: 18px;
    display: flex;
    border-radius: 10px;

    .dashboard {
      height: 646px;
      height: 100%;
      width: 100%;
      flex: 1;
      background: white;
      border-radius: 10px;
    }

    .boards {
      .boards_inner {
        height: 200px;
        background: #f9fbfd;
        margin: 18px 24px;
        border-radius: 9px;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
      }
    }

    .chipTable {
      // max-height: 300px;
      // overflow: auto;
      // scroll-behavior: smooth;
      margin: 18px 23px 0 23px;

      tr {
        th {
          height: 46px;
          background: #f9fbfd;
          font-size: 14px;
          font-weight: 400;
          color: #323233;
        }

        td {
          max-width: 0.1vw;
          min-width: 100px;
          height: 55px;
          font-size: 14px;
          font-weight: 400;
          color: #323233;
        }
      }
    }
  }

  .chips {
    margin: 18px 0 0 23px;
    display: flex;
    flex-direction: row;

    .chip_title {
      margin-left: 7px;
      height: 25px;
      font-weight: 500;
      color: #333333;
      font-size: 18px;
    }

    span {
      display: flex;
      justify-content: center;
      align-items: center;
      width: 55px;
      height: 25px;
      background: #d3e7fb;
      border-radius: 10px;
      font-weight: 400;
      color: #095ebc;
      font-size: 14px;
      margin-left: 6px;
    }
  }

  .dashboard_title {
    margin: 18px 0 0 30px;
    height: 25px;
    font-size: 18px;
    font-weight: 500;
    color: #323233;
  }
</style>
<style lang="less">
  .boards /deep/ {
    .ant-tabs > .ant-tabs-nav {
      position: unset;
      margin: 0;
      //tabs取消下划线
      .ant-tabs-nav-wrap {
        margin: 18px 0 0 31px;
      }

      .ant-tabs-tab {
        width: 74px;
        height: 36px;
        background: #f9fbfd;
        border-radius: 6px;
        border: 1px solid #e6e8eb;
        font-size: 14px;
        font-weight: 400;
        color: #8c8c8c;
        display: flex;
        justify-content: center;
        align-items: center;
      }

      .ant-tabs-tab-active {
        background: #e6f2ff;
        color: #0960bd;
        border: 1px solid #0d60bb;
      }
    }
  }
  @media screen and (max-width: 1660px) {
    .boards /deep/ {
      .ant-tabs > .ant-tabs-nav {
        .ant-tabs-tab {
          width: 52px;
          height: 30px;
        }
      }
    }
  }
</style>
