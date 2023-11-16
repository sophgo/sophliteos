<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    :showOkBtn="false"
    width="40vw"
    :height="600"
    @cancel="player.destroy()"
  >
    <video ref="video" controls></video>
  </BasicModal>
</template>
<script ts setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import mpegts from 'mpegts.js';
  import { ref } from 'vue';
  const title = ref('');
  const url = ref();
  const video = ref();
  let player = null;
  const [registerModal, { setModalProps }] = useModalInner((data) => {
    title.value = data.record.name;
    url.value = data.res;

    if (mpegts.getFeatureList().mseLivePlayback) {
      player = mpegts.createPlayer({
        type: 'mse', // could also be mpegts, m2ts, flv
        isLive: true,
        url: url.value,
      });
      player.attachMediaElement(video.value);
      player.load();
      player.play();
    }
    setModalProps({ confirmLoading: false });
  });
</script>
<style scoped>
  video {
    height: 100%;
    width: 100%;
  }
</style>
