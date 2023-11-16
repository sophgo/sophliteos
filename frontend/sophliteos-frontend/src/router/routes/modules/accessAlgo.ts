import type { AppRouteModule } from '/@/router/types';
import { t } from '/@/hooks/web/useI18n';
import { LAYOUT } from '/@/router/constant';

const accessAlgo: AppRouteModule = {
  path: '/accessAlgo',
  name: 'accessAlgo',
  component: LAYOUT,
  redirect: '/accessAlgo/videos',
  meta: {
    orderNo: 10,
    // hideChildrenInMenu: true,
    icon: 'tdesign:task',
    title: t('routes.dashboard.accessAlgo'),
  },
  children: [
    {
      path: 'videoManage',
      name: 'videoManage',
      component: () => import('/@/views/accessAlgo/dataSource/videoManage/videos.vue'),
      meta: {
        title: t('routes.dashboard.videoManage'),
      },
    },
    {
      path: 'taskList',
      name: 'TaskList',
      component: () => import('/@/views/accessAlgo/task/taskList/index.vue'),
      meta: {
        title: t('routes.dashboard.task'),
      },
    },
    {
      path: 'AlarmRetrieval',
      name: 'AlarmRetrieval',
      component: () => import('/@/views/accessAlgo/alarmRetrieval/index.vue'),
      meta: {
        title: t('routes.dashboard.alarmRetrieval'),
        hideChildrenInMenu: true,
      },
      children: [
        {
          path: 'AlarmDetail/:image',
          name: 'AlarmDetail',
          component: () => import('/@/views/accessAlgo/alarmRetrieval/alarmDetail.vue'),
          meta: {
            title: t('routes.dashboard.alarmDetail'),
            // hideMenu: true,
            // hideBreadcrumb: true, // 在隐藏面包屑中，隐藏当前菜单
            // hideTab: true,
          },
        },
      ],
    },

    {
      path: 'paramList',
      name: 'paramList',
      component: () => import('/@/views/accessAlgo/paramConfig/paramConfig/index.vue'),
      meta: {
        title: t('routes.dashboard.AlgoParamConfig'),
      },
    },
  ],
};

export default accessAlgo;
