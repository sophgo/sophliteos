import type { AppRouteModule } from '/@/router/types';
import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';
// import { useDeviceInfo } from '/@/store/modules/overview';
// const deviceStore = useDeviceInfo();
// console.log(deviceStore.isSingleBoard);
// setTimeout(() => {
//   console.log(deviceStore.isSingleBoard);
// }, 3000);
const maintenance: AppRouteModule = {
  path: '/maintenance',
  name: 'Maintenance',
  component: LAYOUT,
  redirect: '/maintenance/sysSoft',
  meta: {
    orderNo: 2,
    icon: 'bx:server',
    title: t('routes.dashboard.maintenance'),
    // hideMenu: true,
  },
  children: [
    {
      path: 'sysSoft',
      name: 'SysSoft',
      component: () => import('/@/views/maintenance/sysSoft/index.vue'),
      meta: {
        title: t('routes.dashboard.sysSoft'),
      },
    },

    {
      path: 'softUpdate',
      name: 'softUpdate',
      redirect: '/maintenance/softUpdate/ssmUpdate',
      meta: {
        title: t('routes.dashboard.softUpdate'),
      },
      children: [
        {
          path: 'ssmUpdate',
          name: 'ssmUpdate',
          component: () => import('/@/views/maintenance/softUpdate/ssmUpdate/index.vue'),
          meta: {
            title: t('routes.dashboard.ssmUpdate'),
          },
        },
        {
          path: 'LiteOSUpdate',
          name: 'LiteOSUpdate',
          component: () => import('/@/views/maintenance/softUpdate/index.vue'),
          meta: {
            title: t('routes.dashboard.liteOsUpdate'),
          },
        },
      ],
    },
    {
      path: 'coreBoardMap',
      name: 'coreBoardMap',
      component: () => import('/@/views/maintenance/coreBoardMap/index.vue'),
      meta: {
        title: t('routes.dashboard.coreBoardMap'),
        hideMenu: true,
      },
    },
    {
      path: 'networkSetting',
      name: 'NetworkSetting',
      component: () => import('../../../views/maintenance/networkSetting/index.vue'),
      meta: {
        title: t('routes.dashboard.networkSetting'),
      },
    },
    {
      path: 'threshold',
      name: 'Threshold',
      component: () => import('/@/views/maintenance/threshold/index.vue'),
      meta: {
        title: t('routes.dashboard.threshold'),
      },
    },
  ],
};

export default maintenance;
